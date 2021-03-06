// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMetadataSpec) DeepCopyInto(out *AppMetadataSpec) {
	*out = *in
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]IconSpec, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]ContributorSpec, len(*in))
		copy(*out, *in)
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]ContributorSpec, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]StackInstallSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMetadataSpec.
func (in *AppMetadataSpec) DeepCopy() *AppMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(AppMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Behavior) DeepCopyInto(out *Behavior) {
	*out = *in
	out.CRD = in.CRD
	in.Engine.DeepCopyInto(&out.Engine)
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Behavior.
func (in *Behavior) DeepCopy() *Behavior {
	if in == nil {
		return nil
	}
	out := new(Behavior)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BehaviorCRD) DeepCopyInto(out *BehaviorCRD) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BehaviorCRD.
func (in *BehaviorCRD) DeepCopy() *BehaviorCRD {
	if in == nil {
		return nil
	}
	out := new(BehaviorCRD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CRDList) DeepCopyInto(out *CRDList) {
	{
		in := &in
		*out = make(CRDList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDList.
func (in CRDList) DeepCopy() CRDList {
	if in == nil {
		return nil
	}
	out := new(CRDList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackInstall) DeepCopyInto(out *ClusterStackInstall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackInstall.
func (in *ClusterStackInstall) DeepCopy() *ClusterStackInstall {
	if in == nil {
		return nil
	}
	out := new(ClusterStackInstall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStackInstall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackInstallList) DeepCopyInto(out *ClusterStackInstallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterStackInstall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackInstallList.
func (in *ClusterStackInstallList) DeepCopy() *ClusterStackInstallList {
	if in == nil {
		return nil
	}
	out := new(ClusterStackInstallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStackInstallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContributorSpec) DeepCopyInto(out *ContributorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContributorSpec.
func (in *ContributorSpec) DeepCopy() *ContributorSpec {
	if in == nil {
		return nil
	}
	out := new(ContributorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerDeployment) DeepCopyInto(out *ControllerDeployment) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerDeployment.
func (in *ControllerDeployment) DeepCopy() *ControllerDeployment {
	if in == nil {
		return nil
	}
	out := new(ControllerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerSpec) DeepCopyInto(out *ControllerSpec) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(ControllerDeployment)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerSpec.
func (in *ControllerSpec) DeepCopy() *ControllerSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldBinding) DeepCopyInto(out *FieldBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldBinding.
func (in *FieldBinding) DeepCopy() *FieldBinding {
	if in == nil {
		return nil
	}
	out := new(FieldBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IconSpec) DeepCopyInto(out *IconSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IconSpec.
func (in *IconSpec) DeepCopy() *IconSpec {
	if in == nil {
		return nil
	}
	out := new(IconSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizeEngineConfiguration) DeepCopyInto(out *KustomizeEngineConfiguration) {
	*out = *in
	if in.Overlays != nil {
		in, out := &in.Overlays, &out.Overlays
		*out = make([]KustomizeEngineOverlay, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kustomization != nil {
		in, out := &in.Kustomization, &out.Kustomization
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizeEngineConfiguration.
func (in *KustomizeEngineConfiguration) DeepCopy() *KustomizeEngineConfiguration {
	if in == nil {
		return nil
	}
	out := new(KustomizeEngineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizeEngineOverlay) DeepCopyInto(out *KustomizeEngineOverlay) {
	*out = *in
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]FieldBinding, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizeEngineOverlay.
func (in *KustomizeEngineOverlay) DeepCopy() *KustomizeEngineOverlay {
	if in == nil {
		return nil
	}
	out := new(KustomizeEngineOverlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsSpec) DeepCopyInto(out *PermissionsSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsSpec.
func (in *PermissionsSpec) DeepCopy() *PermissionsSpec {
	if in == nil {
		return nil
	}
	out := new(PermissionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stack) DeepCopyInto(out *Stack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stack.
func (in *Stack) DeepCopy() *Stack {
	if in == nil {
		return nil
	}
	out := new(Stack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackDefinition) DeepCopyInto(out *StackDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackDefinition.
func (in *StackDefinition) DeepCopy() *StackDefinition {
	if in == nil {
		return nil
	}
	out := new(StackDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackDefinitionList) DeepCopyInto(out *StackDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StackDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackDefinitionList.
func (in *StackDefinitionList) DeepCopy() *StackDefinitionList {
	if in == nil {
		return nil
	}
	out := new(StackDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackDefinitionSource) DeepCopyInto(out *StackDefinitionSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackDefinitionSource.
func (in *StackDefinitionSource) DeepCopy() *StackDefinitionSource {
	if in == nil {
		return nil
	}
	out := new(StackDefinitionSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackDefinitionSpec) DeepCopyInto(out *StackDefinitionSpec) {
	*out = *in
	in.AppMetadataSpec.DeepCopyInto(&out.AppMetadataSpec)
	if in.CRDs != nil {
		in, out := &in.CRDs, &out.CRDs
		*out = make(CRDList, len(*in))
		copy(*out, *in)
	}
	in.Controller.DeepCopyInto(&out.Controller)
	in.Permissions.DeepCopyInto(&out.Permissions)
	in.Behavior.DeepCopyInto(&out.Behavior)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackDefinitionSpec.
func (in *StackDefinitionSpec) DeepCopy() *StackDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(StackDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackDefinitionStatus) DeepCopyInto(out *StackDefinitionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackDefinitionStatus.
func (in *StackDefinitionStatus) DeepCopy() *StackDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(StackDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackInstall) DeepCopyInto(out *StackInstall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackInstall.
func (in *StackInstall) DeepCopy() *StackInstall {
	if in == nil {
		return nil
	}
	out := new(StackInstall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackInstall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackInstallList) DeepCopyInto(out *StackInstallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StackInstall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackInstallList.
func (in *StackInstallList) DeepCopy() *StackInstallList {
	if in == nil {
		return nil
	}
	out := new(StackInstallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackInstallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackInstallSpec) DeepCopyInto(out *StackInstallSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackInstallSpec.
func (in *StackInstallSpec) DeepCopy() *StackInstallSpec {
	if in == nil {
		return nil
	}
	out := new(StackInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackInstallStatus) DeepCopyInto(out *StackInstallStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.InstallJob != nil {
		in, out := &in.InstallJob, &out.InstallJob
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.StackRecord != nil {
		in, out := &in.StackRecord, &out.StackRecord
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackInstallStatus.
func (in *StackInstallStatus) DeepCopy() *StackInstallStatus {
	if in == nil {
		return nil
	}
	out := new(StackInstallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackList) DeepCopyInto(out *StackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackList.
func (in *StackList) DeepCopy() *StackList {
	if in == nil {
		return nil
	}
	out := new(StackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackResourceEngineConfiguration) DeepCopyInto(out *StackResourceEngineConfiguration) {
	*out = *in
	if in.Kustomize != nil {
		in, out := &in.Kustomize, &out.Kustomize
		*out = new(KustomizeEngineConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackResourceEngineConfiguration.
func (in *StackResourceEngineConfiguration) DeepCopy() *StackResourceEngineConfiguration {
	if in == nil {
		return nil
	}
	out := new(StackResourceEngineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	in.AppMetadataSpec.DeepCopyInto(&out.AppMetadataSpec)
	if in.CRDs != nil {
		in, out := &in.CRDs, &out.CRDs
		*out = make(CRDList, len(*in))
		copy(*out, *in)
	}
	in.Controller.DeepCopyInto(&out.Controller)
	in.Permissions.DeepCopyInto(&out.Permissions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.ControllerRef != nil {
		in, out := &in.ControllerRef, &out.ControllerRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}
