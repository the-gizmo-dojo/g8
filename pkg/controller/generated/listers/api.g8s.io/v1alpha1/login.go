/*
Copyright 2024 James Riley O'Donnell.

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
	v1alpha1 "github.com/jrodonnell/g8s/pkg/controller/apis/api.g8s.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LoginLister helps list Logins.
// All objects returned here must be treated as read-only.
type LoginLister interface {
	// List lists all Logins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Login, err error)
	// Logins returns an object that can list and get Logins.
	Logins(namespace string) LoginNamespaceLister
	LoginListerExpansion
}

// loginLister implements the LoginLister interface.
type loginLister struct {
	indexer cache.Indexer
}

// NewLoginLister returns a new LoginLister.
func NewLoginLister(indexer cache.Indexer) LoginLister {
	return &loginLister{indexer: indexer}
}

// List lists all Logins in the indexer.
func (s *loginLister) List(selector labels.Selector) (ret []*v1alpha1.Login, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Login))
	})
	return ret, err
}

// Logins returns an object that can list and get Logins.
func (s *loginLister) Logins(namespace string) LoginNamespaceLister {
	return loginNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LoginNamespaceLister helps list and get Logins.
// All objects returned here must be treated as read-only.
type LoginNamespaceLister interface {
	// List lists all Logins in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Login, err error)
	// Get retrieves the Login from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Login, error)
	LoginNamespaceListerExpansion
}

// loginNamespaceLister implements the LoginNamespaceLister
// interface.
type loginNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Logins in the indexer for a given namespace.
func (s loginNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Login, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Login))
	})
	return ret, err
}

// Get retrieves the Login from the indexer for a given namespace and name.
func (s loginNamespaceLister) Get(name string) (*v1alpha1.Login, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("login"), name)
	}
	return obj.(*v1alpha1.Login), nil
}