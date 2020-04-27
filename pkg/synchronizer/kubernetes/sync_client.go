// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"time"

	dplv1alpha1 "github.com/open-cluster-management/multicloud-operators-deployable/pkg/apis/apps/v1"
	gerr "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
)

const (
	syncTimeout = time.Second * 2
)

type DplUnit struct {
	Dpl *dplv1alpha1.Deployable
	Gvk schema.GroupVersionKind
}

type resourceOrder struct {
	subType string
	hostSub types.NamespacedName
	dpls    []DplUnit
	err     chan error
}

type SyncSource interface {
	GetInterval() int
	GetValidatedGVK(schema.GroupVersionKind) *schema.GroupVersionKind
	IsResourceNamespaced(schema.GroupVersionKind) bool
	AddTemplates(string, types.NamespacedName, []DplUnit) error
	CleanupByHost(types.NamespacedName, string)
}

func (sync *KubeSynchronizer) GetInterval() int {
	return sync.Interval
}

// GetValidatedGVK return right gvk from original
func (sync *KubeSynchronizer) GetValidatedGVK(org schema.GroupVersionKind) *schema.GroupVersionKind {
	valid := &org

	if _, ok := internalReplacedGroupVersionKind[org]; ok {
		valid = internalReplacedGroupVersionKind[org]
	}

	gk := schema.GroupKind{Group: valid.Group, Kind: valid.Kind}

	klog.V(5).Infof("gk: %#v, valid: %#v ", gk, valid)

	if _, ok := internalIgnoredGroupKind[gk]; ok {
		return nil
	}

	if sync.Extension != nil && sync.Extension.IsIgnoredGroupKind(gk) {
		return nil
	}

	found := false

	var regGvk schema.GroupVersionKind

	// return the right version of gv
	for gvk := range sync.KubeResources {
		if valid.GroupKind() == gvk.GroupKind() {
			if valid.Version == gvk.Version {
				return &gvk
			}

			found = true
			regGvk = gvk
		}
	}

	// if there's a GK ready served by the k8s, then we are going to registry
	// the incoming unknown version for this GK as well, since k8s would handle
	// the version conversion, if there isn't any version conversion on k8s,
	// user would get deploy failed error, which would aligned with the kubectl
	// behavior
	if found {
		kubeResourceAddVersionToGK(sync.KubeResources, regGvk, *valid)
		return valid
	}

	return nil
}

func (sync *KubeSynchronizer) IsResourceNamespaced(gvk schema.GroupVersionKind) bool {
	return sync.KubeResources[gvk].Namespaced
}

func (sync *KubeSynchronizer) AddTemplates(subType string, hostSub types.NamespacedName, dpls []DplUnit) error {
	rsOrder := resourceOrder{
		subType: subType,
		hostSub: hostSub,
		dpls:    dpls,
		err:     make(chan error, 1),
	}

	select {
	case sync.tplCh <- rsOrder:
		klog.Info("write resource request/order to cache")
	case <-time.After(syncTimeout):
		return gerr.New("timeout on writing templates to syncrhonizer")
	}

	var err error

	select {
	case serr := <-rsOrder.err:
		if serr != nil {
			return gerr.Wrap(err, "failed to add templates")
		}
	case <-time.After(syncTimeout):
		return gerr.New("timeout on waiting templates write result from syncrhonizer")
	}

	return nil
}

// CleanupByHost returns initialized validator struct
func (sync *KubeSynchronizer) CleanupByHost(host types.NamespacedName, syncsource string) {
	if err := sync.AddTemplates(syncsource, host, []DplUnit{}); err != nil {
		klog.Errorf("failed to clean up by host, having err: %v", err)
	}
}
