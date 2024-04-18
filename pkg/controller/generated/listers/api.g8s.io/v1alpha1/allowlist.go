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

// AllowlistLister helps list Allowlists.
// All objects returned here must be treated as read-only.
type AllowlistLister interface {
	// List lists all Allowlists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Allowlist, err error)
	// Get retrieves the Allowlist from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Allowlist, error)
	AllowlistListerExpansion
}

// allowlistLister implements the AllowlistLister interface.
type allowlistLister struct {
	indexer cache.Indexer
}

// NewAllowlistLister returns a new AllowlistLister.
func NewAllowlistLister(indexer cache.Indexer) AllowlistLister {
	return &allowlistLister{indexer: indexer}
}

// List lists all Allowlists in the indexer.
func (s *allowlistLister) List(selector labels.Selector) (ret []*v1alpha1.Allowlist, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Allowlist))
	})
	return ret, err
}

// Get retrieves the Allowlist from the index for a given name.
func (s *allowlistLister) Get(name string) (*v1alpha1.Allowlist, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("allowlist"), name)
	}
	return obj.(*v1alpha1.Allowlist), nil
}