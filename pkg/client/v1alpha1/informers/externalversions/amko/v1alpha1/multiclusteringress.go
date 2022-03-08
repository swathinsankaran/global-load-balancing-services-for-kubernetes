/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	amkov1alpha1 "github.com/vmware/global-load-balancing-services-for-kubernetes/pkg/apis/amko/v1alpha1"
	versioned "github.com/vmware/global-load-balancing-services-for-kubernetes/pkg/client/v1alpha1/clientset/versioned"
	internalinterfaces "github.com/vmware/global-load-balancing-services-for-kubernetes/pkg/client/v1alpha1/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vmware/global-load-balancing-services-for-kubernetes/pkg/client/v1alpha1/listers/amko/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MultiClusterIngressInformer provides access to a shared informer and lister for
// MultiClusterIngresses.
type MultiClusterIngressInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MultiClusterIngressLister
}

type multiClusterIngressInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMultiClusterIngressInformer constructs a new informer for MultiClusterIngress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMultiClusterIngressInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMultiClusterIngressInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMultiClusterIngressInformer constructs a new informer for MultiClusterIngress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMultiClusterIngressInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AmkoV1alpha1().MultiClusterIngresses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AmkoV1alpha1().MultiClusterIngresses(namespace).Watch(context.TODO(), options)
			},
		},
		&amkov1alpha1.MultiClusterIngress{},
		resyncPeriod,
		indexers,
	)
}

func (f *multiClusterIngressInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMultiClusterIngressInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *multiClusterIngressInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&amkov1alpha1.MultiClusterIngress{}, f.defaultInformer)
}

func (f *multiClusterIngressInformer) Lister() v1alpha1.MultiClusterIngressLister {
	return v1alpha1.NewMultiClusterIngressLister(f.Informer().GetIndexer())
}