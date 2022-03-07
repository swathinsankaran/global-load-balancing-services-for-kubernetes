#!/usr/bin/env bash

# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

if [ $# -ne 1 ]; then
	echo "Codegen needs a version parameter, e.g.: v1alpha1"
	exit 1
fi

echo "Generating clientsets and informers for $1"

# configurable envs
readonly AMKOCRD_VERSION=$1
readonly AMKO_PACKAGE=github.com/vmware/global-load-balancing-services-for-kubernetes

###

readonly SCRIPT_ROOT="$(cd "$(dirname "${BASH_SOURCE}")"/.. && pwd)"

readonly GO111MODULE="on"
readonly GOFLAGS="-mod=mod"
readonly GOPATH="$(mktemp -d)"

export GO111MODULE GOFLAGS GOPATH

# Even when modules are enabled, the code-generator tools always write to
# a traditional GOPATH directory, so fake on up to point to the current
# workspace.
mkdir -p "$GOPATH/src/github.com/vmware"
ln -s "${SCRIPT_ROOT}" "$GOPATH/src/${AMKO_PACKAGE}"

readonly OUTPUT_PKG="${AMKO_PACKAGE}/pkg/client/${AMKOCRD_VERSION}"
readonly FQ_APIS="${AMKO_PACKAGE}/pkg/apis/amko/${AMKOCRD_VERSION}"
readonly APIS_PKG=AMKO_PACKAGE
readonly CLIENTSET_NAME=versioned
readonly CLIENTSET_PKG_NAME=clientset

if [[ "${VERIFY_CODEGEN:-}" == "true" ]]; then
  echo "Running in verification mode"
  readonly VERIFY_FLAG="--verify-only"
fi

readonly COMMON_FLAGS="${VERIFY_FLAG:-} --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt"

echo "Generating deepcopy funcs"
go run k8s.io/code-generator/cmd/deepcopy-gen \
        --input-dirs "${FQ_APIS}" \
        -O zz_generated.deepcopy \
        --bounding-dirs "${APIS_PKG}" \
        ${COMMON_FLAGS}

echo "Generating clientset at ${OUTPUT_PKG}/${CLIENTSET_PKG_NAME}"
go run k8s.io/code-generator/cmd/client-gen \
        --clientset-name "${CLIENTSET_NAME}" \
        --input-base "" \
        --input "${FQ_APIS}" \
        --output-package "${OUTPUT_PKG}/${CLIENTSET_PKG_NAME}" \
        ${COMMON_FLAGS}

echo "Generating listers at ${OUTPUT_PKG}/listers"
go run k8s.io/code-generator/cmd/lister-gen \
        --input-dirs "${FQ_APIS}" \
        --output-package "${OUTPUT_PKG}/listers" \
        ${COMMON_FLAGS}

echo "Generating informers at ${OUTPUT_PKG}/informers"
go run k8s.io/code-generator/cmd/informer-gen \
         --input-dirs "${FQ_APIS}" \
         --versioned-clientset-package "${OUTPUT_PKG}/${CLIENTSET_PKG_NAME}/${CLIENTSET_NAME}" \
         --listers-package "${OUTPUT_PKG}/listers" \
         --output-package "${OUTPUT_PKG}/informers" \
         ${COMMON_FLAGS}

echo "Generating register at ${FQ_APIS}"
go run k8s.io/code-generator/cmd/register-gen \
        --input-dirs "${FQ_APIS}" \
        --output-package "${FQ_APIS}" \
        ${COMMON_FLAGS}
