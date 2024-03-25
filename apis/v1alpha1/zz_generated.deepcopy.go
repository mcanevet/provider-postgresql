//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfig) DeepCopyInto(out *StoreConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfig.
func (in *StoreConfig) DeepCopy() *StoreConfig {
	if in == nil {
		return nil
	}
	out := new(StoreConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigList) DeepCopyInto(out *StoreConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoreConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigList.
func (in *StoreConfigList) DeepCopy() *StoreConfigList {
	if in == nil {
		return nil
	}
	out := new(StoreConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigSpec) DeepCopyInto(out *StoreConfigSpec) {
	*out = *in
	in.SecretStoreConfig.DeepCopyInto(&out.SecretStoreConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigSpec.
func (in *StoreConfigSpec) DeepCopy() *StoreConfigSpec {
	if in == nil {
		return nil
	}
	out := new(StoreConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreConfigStatus) DeepCopyInto(out *StoreConfigStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreConfigStatus.
func (in *StoreConfigStatus) DeepCopy() *StoreConfigStatus {
	if in == nil {
		return nil
	}
	out := new(StoreConfigStatus)
	in.DeepCopyInto(out)
	return out
}
