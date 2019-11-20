/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "knative.dev/eventing-contrib/couchdb/source/pkg/apis/sources/v1alpha1"
)

// CouchDbSourceLister helps list CouchDbSources.
type CouchDbSourceLister interface {
	// List lists all CouchDbSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CouchDbSource, err error)
	// CouchDbSources returns an object that can list and get CouchDbSources.
	CouchDbSources(namespace string) CouchDbSourceNamespaceLister
	CouchDbSourceListerExpansion
}

// couchDbSourceLister implements the CouchDbSourceLister interface.
type couchDbSourceLister struct {
	indexer cache.Indexer
}

// NewCouchDbSourceLister returns a new CouchDbSourceLister.
func NewCouchDbSourceLister(indexer cache.Indexer) CouchDbSourceLister {
	return &couchDbSourceLister{indexer: indexer}
}

// List lists all CouchDbSources in the indexer.
func (s *couchDbSourceLister) List(selector labels.Selector) (ret []*v1alpha1.CouchDbSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CouchDbSource))
	})
	return ret, err
}

// CouchDbSources returns an object that can list and get CouchDbSources.
func (s *couchDbSourceLister) CouchDbSources(namespace string) CouchDbSourceNamespaceLister {
	return couchDbSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CouchDbSourceNamespaceLister helps list and get CouchDbSources.
type CouchDbSourceNamespaceLister interface {
	// List lists all CouchDbSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CouchDbSource, err error)
	// Get retrieves the CouchDbSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CouchDbSource, error)
	CouchDbSourceNamespaceListerExpansion
}

// couchDbSourceNamespaceLister implements the CouchDbSourceNamespaceLister
// interface.
type couchDbSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CouchDbSources in the indexer for a given namespace.
func (s couchDbSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CouchDbSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CouchDbSource))
	})
	return ret, err
}

// Get retrieves the CouchDbSource from the indexer for a given namespace and name.
func (s couchDbSourceNamespaceLister) Get(name string) (*v1alpha1.CouchDbSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("couchdbsource"), name)
	}
	return obj.(*v1alpha1.CouchDbSource), nil
}