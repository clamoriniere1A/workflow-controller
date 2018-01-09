// Generated file, do not modify manually!

// This file was automatically generated by informer-gen

package v1

import (
	time "time"

	workflow_v1 "github.com/amadeusitgroup/workflow-controller/pkg/api/workflow/v1"
	internalinterfaces "github.com/amadeusitgroup/workflow-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/amadeusitgroup/workflow-controller/pkg/client/listers/workflow/v1"
	versioned "github.com/amadeusitgroup/workflow-controller/pkg/client/versioned"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WorkflowInformer provides access to a shared informer and lister for
// Workflows.
type WorkflowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.WorkflowLister
}

type workflowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkflowInformer constructs a new informer for Workflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkflowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkflowInformer constructs a new informer for Workflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkflowV1().Workflows(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkflowV1().Workflows(namespace).Watch(options)
			},
		},
		&workflow_v1.Workflow{},
		resyncPeriod,
		indexers,
	)
}

func (f *workflowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkflowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workflowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workflow_v1.Workflow{}, f.defaultInformer)
}

func (f *workflowInformer) Lister() v1.WorkflowLister {
	return v1.NewWorkflowLister(f.Informer().GetIndexer())
}
