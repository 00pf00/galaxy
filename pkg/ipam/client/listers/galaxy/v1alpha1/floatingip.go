/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "tkestack.io/galaxy/pkg/ipam/apis/galaxy/v1alpha1"
)

// FloatingIPLister helps list FloatingIPs.
// All objects returned here must be treated as read-only.
type FloatingIPLister interface {
	// List lists all FloatingIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error)
	// Get retrieves the FloatingIP from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FloatingIP, error)
	FloatingIPListerExpansion
}

// floatingIPLister implements the FloatingIPLister interface.
type floatingIPLister struct {
	indexer cache.Indexer
}

// NewFloatingIPLister returns a new FloatingIPLister.
func NewFloatingIPLister(indexer cache.Indexer) FloatingIPLister {
	return &floatingIPLister{indexer: indexer}
}

// List lists all FloatingIPs in the indexer.
func (s *floatingIPLister) List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FloatingIP))
	})
	return ret, err
}

// Get retrieves the FloatingIP from the index for a given name.
func (s *floatingIPLister) Get(name string) (*v1alpha1.FloatingIP, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("floatingip"), name)
	}
	return obj.(*v1alpha1.FloatingIP), nil
}
