// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartRepository) DeepCopyInto(out *HelmChartRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartRepository.
func (in *HelmChartRepository) DeepCopy() *HelmChartRepository {
	if in == nil {
		return nil
	}
	out := new(HelmChartRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmChartRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartRepositoryList) DeepCopyInto(out *HelmChartRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmChartRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartRepositoryList.
func (in *HelmChartRepositoryList) DeepCopy() *HelmChartRepositoryList {
	if in == nil {
		return nil
	}
	out := new(HelmChartRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmChartRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartRepositorySpec) DeepCopyInto(out *HelmChartRepositorySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartRepositorySpec.
func (in *HelmChartRepositorySpec) DeepCopy() *HelmChartRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(HelmChartRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartRepositoryStatus) DeepCopyInto(out *HelmChartRepositoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartRepositoryStatus.
func (in *HelmChartRepositoryStatus) DeepCopy() *HelmChartRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(HelmChartRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
