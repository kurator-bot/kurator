/*
Copyright Kurator Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	fleetv1alpha1 "kurator.dev/kurator/pkg/apis/fleet/v1alpha1"
	versioned "kurator.dev/kurator/pkg/client-go/generated/clientset/versioned"
	internalinterfaces "kurator.dev/kurator/pkg/client-go/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "kurator.dev/kurator/pkg/client-go/generated/listers/fleet/v1alpha1"
)

// FleetInformer provides access to a shared informer and lister for
// Fleets.
type FleetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FleetLister
}

type fleetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFleetInformer constructs a new informer for Fleet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFleetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFleetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFleetInformer constructs a new informer for Fleet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFleetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FleetV1alpha1().Fleets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FleetV1alpha1().Fleets(namespace).Watch(context.TODO(), options)
			},
		},
		&fleetv1alpha1.Fleet{},
		resyncPeriod,
		indexers,
	)
}

func (f *fleetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFleetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fleetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fleetv1alpha1.Fleet{}, f.defaultInformer)
}

func (f *fleetInformer) Lister() v1alpha1.FleetLister {
	return v1alpha1.NewFleetLister(f.Informer().GetIndexer())
}