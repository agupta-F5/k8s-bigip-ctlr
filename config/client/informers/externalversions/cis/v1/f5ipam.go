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

package v1

import (
	time "time"

	cisv1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	versioned "github.com/F5Networks/k8s-bigip-ctlr/config/client/clientset/versioned"
	internalinterfaces "github.com/F5Networks/k8s-bigip-ctlr/config/client/informers/externalversions/internalinterfaces"
	v1 "github.com/F5Networks/k8s-bigip-ctlr/config/client/listers/cis/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// F5IPAMInformer provides access to a shared informer and lister for
// F5IPAMs.
type F5IPAMInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.F5IPAMLister
}

type f5IPAMInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewF5IPAMInformer constructs a new informer for F5IPAM type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewF5IPAMInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredF5IPAMInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredF5IPAMInformer constructs a new informer for F5IPAM type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredF5IPAMInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().F5IPAMs(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().F5IPAMs(namespace).Watch(options)
			},
		},
		&cisv1.F5IPAM{},
		resyncPeriod,
		indexers,
	)
}

func (f *f5IPAMInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredF5IPAMInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *f5IPAMInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cisv1.F5IPAM{}, f.defaultInformer)
}

func (f *f5IPAMInformer) Lister() v1.F5IPAMLister {
	return v1.NewF5IPAMLister(f.Informer().GetIndexer())
}
