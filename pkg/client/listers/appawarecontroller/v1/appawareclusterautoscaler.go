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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/application-aware-controller/pkg/apis/appawarecontroller/v1"
	"k8s.io/client-go/tools/cache"
)

// AppawareClusterAutoscalerLister helps list AppawareClusterAutoscalers.
// All objects returned here must be treated as read-only.
type AppawareClusterAutoscalerLister interface {
	// List lists all AppawareClusterAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AppawareClusterAutoscaler, err error)
	// AppawareClusterAutoscalers returns an object that can list and get AppawareClusterAutoscalers.
	AppawareClusterAutoscalers(namespace string) AppawareClusterAutoscalerNamespaceLister
	AppawareClusterAutoscalerListerExpansion
}

// appawareClusterAutoscalerLister implements the AppawareClusterAutoscalerLister interface.
type appawareClusterAutoscalerLister struct {
	indexer cache.Indexer
}

// NewAppawareClusterAutoscalerLister returns a new AppawareClusterAutoscalerLister.
func NewAppawareClusterAutoscalerLister(indexer cache.Indexer) AppawareClusterAutoscalerLister {
	return &appawareClusterAutoscalerLister{indexer: indexer}
}

// List lists all AppawareClusterAutoscalers in the indexer.
func (s *appawareClusterAutoscalerLister) List(selector labels.Selector) (ret []*v1.AppawareClusterAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppawareClusterAutoscaler))
	})
	return ret, err
}

// AppawareClusterAutoscalers returns an object that can list and get AppawareClusterAutoscalers.
func (s *appawareClusterAutoscalerLister) AppawareClusterAutoscalers(namespace string) AppawareClusterAutoscalerNamespaceLister {
	return appawareClusterAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppawareClusterAutoscalerNamespaceLister helps list and get AppawareClusterAutoscalers.
// All objects returned here must be treated as read-only.
type AppawareClusterAutoscalerNamespaceLister interface {
	// List lists all AppawareClusterAutoscalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AppawareClusterAutoscaler, err error)
	// Get retrieves the AppawareClusterAutoscaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AppawareClusterAutoscaler, error)
	AppawareClusterAutoscalerNamespaceListerExpansion
}

// appawareClusterAutoscalerNamespaceLister implements the AppawareClusterAutoscalerNamespaceLister
// interface.
type appawareClusterAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppawareClusterAutoscalers in the indexer for a given namespace.
func (s appawareClusterAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1.AppawareClusterAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppawareClusterAutoscaler))
	})
	return ret, err
}

// Get retrieves the AppawareClusterAutoscaler from the indexer for a given namespace and name.
func (s appawareClusterAutoscalerNamespaceLister) Get(name string) (*v1.AppawareClusterAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("appawareclusterautoscaler"), name)
	}
	return obj.(*v1.AppawareClusterAutoscaler), nil
}
