/*
Copyright 2020 The Kubernetes Authors

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
	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkAttachmentDefinitionLister helps list NetworkAttachmentDefinitions.
type NetworkAttachmentDefinitionLister interface {
	// List lists all NetworkAttachmentDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1.NetworkAttachmentDefinition, err error)
	// NetworkAttachmentDefinitions returns an object that can list and get NetworkAttachmentDefinitions.
	NetworkAttachmentDefinitions(namespace string) NetworkAttachmentDefinitionNamespaceLister
	NetworkAttachmentDefinitionListerExpansion
}

// networkAttachmentDefinitionLister implements the NetworkAttachmentDefinitionLister interface.
type networkAttachmentDefinitionLister struct {
	indexer cache.Indexer
}

// NewNetworkAttachmentDefinitionLister returns a new NetworkAttachmentDefinitionLister.
func NewNetworkAttachmentDefinitionLister(indexer cache.Indexer) NetworkAttachmentDefinitionLister {
	return &networkAttachmentDefinitionLister{indexer: indexer}
}

// List lists all NetworkAttachmentDefinitions in the indexer.
func (s *networkAttachmentDefinitionLister) List(selector labels.Selector) (ret []*v1.NetworkAttachmentDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkAttachmentDefinition))
	})
	return ret, err
}

// NetworkAttachmentDefinitions returns an object that can list and get NetworkAttachmentDefinitions.
func (s *networkAttachmentDefinitionLister) NetworkAttachmentDefinitions(namespace string) NetworkAttachmentDefinitionNamespaceLister {
	return networkAttachmentDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkAttachmentDefinitionNamespaceLister helps list and get NetworkAttachmentDefinitions.
type NetworkAttachmentDefinitionNamespaceLister interface {
	// List lists all NetworkAttachmentDefinitions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.NetworkAttachmentDefinition, err error)
	// Get retrieves the NetworkAttachmentDefinition from the indexer for a given namespace and name.
	Get(name string) (*v1.NetworkAttachmentDefinition, error)
	NetworkAttachmentDefinitionNamespaceListerExpansion
}

// networkAttachmentDefinitionNamespaceLister implements the NetworkAttachmentDefinitionNamespaceLister
// interface.
type networkAttachmentDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkAttachmentDefinitions in the indexer for a given namespace.
func (s networkAttachmentDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1.NetworkAttachmentDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkAttachmentDefinition))
	})
	return ret, err
}

// Get retrieves the NetworkAttachmentDefinition from the indexer for a given namespace and name.
func (s networkAttachmentDefinitionNamespaceLister) Get(name string) (*v1.NetworkAttachmentDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("networkattachmentdefinition"), name)
	}
	return obj.(*v1.NetworkAttachmentDefinition), nil
}
