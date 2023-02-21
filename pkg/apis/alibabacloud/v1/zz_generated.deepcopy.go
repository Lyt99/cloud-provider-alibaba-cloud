//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessLogConfig) DeepCopyInto(out *AccessLogConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogConfig.
func (in *AccessLogConfig) DeepCopy() *AccessLogConfig {
	if in == nil {
		return nil
	}
	out := new(AccessLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessLogTracingConfig) DeepCopyInto(out *AccessLogTracingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogTracingConfig.
func (in *AccessLogTracingConfig) DeepCopy() *AccessLogTracingConfig {
	if in == nil {
		return nil
	}
	out := new(AccessLogTracingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	if in.FixedResponseConfig != nil {
		in, out := &in.FixedResponseConfig, &out.FixedResponseConfig
		*out = new(FixedResponseActionConfig)
		**out = **in
	}
	if in.RedirectConfig != nil {
		in, out := &in.RedirectConfig, &out.RedirectConfig
		*out = new(RedirectActionConfig)
		**out = **in
	}
	if in.ForwardConfig != nil {
		in, out := &in.ForwardConfig, &out.ForwardConfig
		*out = new(ForwardActionConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbConfig) DeepCopyInto(out *AlbConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbConfig.
func (in *AlbConfig) DeepCopy() *AlbConfig {
	if in == nil {
		return nil
	}
	out := new(AlbConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlbConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbConfigList) DeepCopyInto(out *AlbConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlbConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbConfigList.
func (in *AlbConfigList) DeepCopy() *AlbConfigList {
	if in == nil {
		return nil
	}
	out := new(AlbConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlbConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlbConfigSpec) DeepCopyInto(out *AlbConfigSpec) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]*ListenerSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ListenerSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlbConfigSpec.
func (in *AlbConfigSpec) DeepCopy() *AlbConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AlbConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingConfig) DeepCopyInto(out *BillingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingConfig.
func (in *BillingConfig) DeepCopy() *BillingConfig {
	if in == nil {
		return nil
	}
	out := new(BillingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeletionProtectionConfig) DeepCopyInto(out *DeletionProtectionConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeletionProtectionConfig.
func (in *DeletionProtectionConfig) DeepCopy() *DeletionProtectionConfig {
	if in == nil {
		return nil
	}
	out := new(DeletionProtectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedResponseActionConfig) DeepCopyInto(out *FixedResponseActionConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedResponseActionConfig.
func (in *FixedResponseActionConfig) DeepCopy() *FixedResponseActionConfig {
	if in == nil {
		return nil
	}
	out := new(FixedResponseActionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardActionConfig) DeepCopyInto(out *ForwardActionConfig) {
	*out = *in
	if in.TargetGroups != nil {
		in, out := &in.TargetGroups, &out.TargetGroups
		*out = make([]TargetGroupTuple, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardActionConfig.
func (in *ForwardActionConfig) DeepCopy() *ForwardActionConfig {
	if in == nil {
		return nil
	}
	out := new(ForwardActionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressStatus) DeepCopyInto(out *IngressStatus) {
	*out = *in
	out.LoadBalancer = in.LoadBalancer
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressStatus.
func (in *IngressStatus) DeepCopy() *IngressStatus {
	if in == nil {
		return nil
	}
	out := new(IngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerSpec) DeepCopyInto(out *ListenerSpec) {
	*out = *in
	if in.GzipEnabled != nil {
		in, out := &in.GzipEnabled, &out.GzipEnabled
		*out = new(bool)
		**out = **in
	}
	out.QuicConfig = in.QuicConfig
	if in.Http2Enabled != nil {
		in, out := &in.Http2Enabled, &out.Http2Enabled
		*out = new(bool)
		**out = **in
	}
	if in.DefaultActions != nil {
		in, out := &in.DefaultActions, &out.DefaultActions
		*out = make([]Action, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Port = in.Port
	if in.CaCertificates != nil {
		in, out := &in.CaCertificates, &out.CaCertificates
		*out = make([]Certificate, len(*in))
		copy(*out, *in)
	}
	out.XForwardedForConfig = in.XForwardedForConfig
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = make([]Certificate, len(*in))
		copy(*out, *in)
	}
	out.LogConfig = in.LogConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerSpec.
func (in *ListenerSpec) DeepCopy() *ListenerSpec {
	if in == nil {
		return nil
	}
	out := new(ListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.ZoneMappings != nil {
		in, out := &in.ZoneMappings, &out.ZoneMappings
		*out = make([]ZoneMapping, len(*in))
		copy(*out, *in)
	}
	out.AccessLogConfig = in.AccessLogConfig
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	out.BillingConfig = in.BillingConfig
	if in.ForceOverride != nil {
		in, out := &in.ForceOverride, &out.ForceOverride
		*out = new(bool)
		**out = **in
	}
	out.ModificationProtectionConfig = in.ModificationProtectionConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfig) DeepCopyInto(out *LogConfig) {
	*out = *in
	out.AccessLogTracingConfig = in.AccessLogTracingConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfig.
func (in *LogConfig) DeepCopy() *LogConfig {
	if in == nil {
		return nil
	}
	out := new(LogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModificationProtectionConfig) DeepCopyInto(out *ModificationProtectionConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModificationProtectionConfig.
func (in *ModificationProtectionConfig) DeepCopy() *ModificationProtectionConfig {
	if in == nil {
		return nil
	}
	out := new(ModificationProtectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuicConfig) DeepCopyInto(out *QuicConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuicConfig.
func (in *QuicConfig) DeepCopy() *QuicConfig {
	if in == nil {
		return nil
	}
	out := new(QuicConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedirectActionConfig) DeepCopyInto(out *RedirectActionConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedirectActionConfig.
func (in *RedirectActionConfig) DeepCopy() *RedirectActionConfig {
	if in == nil {
		return nil
	}
	out := new(RedirectActionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupTuple) DeepCopyInto(out *TargetGroupTuple) {
	*out = *in
	out.ServicePort = in.ServicePort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupTuple.
func (in *TargetGroupTuple) DeepCopy() *TargetGroupTuple {
	if in == nil {
		return nil
	}
	out := new(TargetGroupTuple)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XForwardedForConfig) DeepCopyInto(out *XForwardedForConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XForwardedForConfig.
func (in *XForwardedForConfig) DeepCopy() *XForwardedForConfig {
	if in == nil {
		return nil
	}
	out := new(XForwardedForConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneMapping) DeepCopyInto(out *ZoneMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneMapping.
func (in *ZoneMapping) DeepCopy() *ZoneMapping {
	if in == nil {
		return nil
	}
	out := new(ZoneMapping)
	in.DeepCopyInto(out)
	return out
}
