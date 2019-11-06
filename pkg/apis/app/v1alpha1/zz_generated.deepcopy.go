// +build !ignore_autogenerated

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

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	appv1alpha1 "github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1"
	pkgapisappv1alpha1 "github.com/IBM/multicloud-operators-deployable/pkg/apis/app/v1alpha1"
	apisappv1alpha1 "github.com/IBM/multicloud-operators-placementrule/pkg/apis/app/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HourRange) DeepCopyInto(out *HourRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HourRange.
func (in *HourRange) DeepCopy() *HourRange {
	if in == nil {
		return nil
	}
	out := new(HourRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overrides) DeepCopyInto(out *Overrides) {
	*out = *in
	if in.PackageOverrides != nil {
		in, out := &in.PackageOverrides, &out.PackageOverrides
		*out = make([]PackageOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overrides.
func (in *Overrides) DeepCopy() *Overrides {
	if in == nil {
		return nil
	}
	out := new(Overrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageFilter) DeepCopyInto(out *PackageFilter) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FilterRef != nil {
		in, out := &in.FilterRef, &out.FilterRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageFilter.
func (in *PackageFilter) DeepCopy() *PackageFilter {
	if in == nil {
		return nil
	}
	out := new(PackageFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageOverride) DeepCopyInto(out *PackageOverride) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageOverride.
func (in *PackageOverride) DeepCopy() *PackageOverride {
	if in == nil {
		return nil
	}
	out := new(PackageOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberItem) DeepCopyInto(out *SubscriberItem) {
	*out = *in
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(Subscription)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriptionConfigMap != nil {
		in, out := &in.SubscriptionConfigMap, &out.SubscriptionConfigMap
		*out = new(corev1.ConfigMap)
		(*in).DeepCopyInto(*out)
	}
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(appv1alpha1.Channel)
		(*in).DeepCopyInto(*out)
	}
	if in.ChannelSecret != nil {
		in, out := &in.ChannelSecret, &out.ChannelSecret
		*out = new(corev1.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.ChannelConfigMap != nil {
		in, out := &in.ChannelConfigMap, &out.ChannelConfigMap
		*out = new(corev1.ConfigMap)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberItem.
func (in *SubscriberItem) DeepCopy() *SubscriberItem {
	if in == nil {
		return nil
	}
	out := new(SubscriberItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SubscriptionClusterStatusMap) DeepCopyInto(out *SubscriptionClusterStatusMap) {
	{
		in := &in
		*out = make(SubscriptionClusterStatusMap, len(*in))
		for key, val := range *in {
			var outVal *SubscriptionPerClusterStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SubscriptionPerClusterStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionClusterStatusMap.
func (in SubscriptionClusterStatusMap) DeepCopy() SubscriptionClusterStatusMap {
	if in == nil {
		return nil
	}
	out := new(SubscriptionClusterStatusMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionPerClusterStatus) DeepCopyInto(out *SubscriptionPerClusterStatus) {
	*out = *in
	if in.SubscriptionPackageStatus != nil {
		in, out := &in.SubscriptionPackageStatus, &out.SubscriptionPackageStatus
		*out = make(map[string]*SubscriptionUnitStatus, len(*in))
		for key, val := range *in {
			var outVal *SubscriptionUnitStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SubscriptionUnitStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionPerClusterStatus.
func (in *SubscriptionPerClusterStatus) DeepCopy() *SubscriptionPerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionPerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	if in.PackageFilter != nil {
		in, out := &in.PackageFilter, &out.PackageFilter
		*out = new(PackageFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.PackageOverrides != nil {
		in, out := &in.PackageOverrides, &out.PackageOverrides
		*out = make([]*Overrides, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Overrides)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(apisappv1alpha1.Placement)
		(*in).DeepCopyInto(*out)
	}
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]pkgapisappv1alpha1.Overrides, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeWindow != nil {
		in, out := &in.TimeWindow, &out.TimeWindow
		*out = new(TimeWindow)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.Statuses != nil {
		in, out := &in.Statuses, &out.Statuses
		*out = make(SubscriptionClusterStatusMap, len(*in))
		for key, val := range *in {
			var outVal *SubscriptionPerClusterStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SubscriptionPerClusterStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionUnitStatus) DeepCopyInto(out *SubscriptionUnitStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.ResourceStatus != nil {
		in, out := &in.ResourceStatus, &out.ResourceStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionUnitStatus.
func (in *SubscriptionUnitStatus) DeepCopy() *SubscriptionUnitStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionUnitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeWindow) DeepCopyInto(out *TimeWindow) {
	*out = *in
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]time.Weekday, len(*in))
		copy(*out, *in)
	}
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = make([]HourRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeWindow.
func (in *TimeWindow) DeepCopy() *TimeWindow {
	if in == nil {
		return nil
	}
	out := new(TimeWindow)
	in.DeepCopyInto(out)
	return out
}
