//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManager) DeepCopyInto(out *ClusterManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManager.
func (in *ClusterManager) DeepCopy() *ClusterManager {
	if in == nil {
		return nil
	}
	out := new(ClusterManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerDeployOption) DeepCopyInto(out *ClusterManagerDeployOption) {
	*out = *in
	if in.Hosted != nil {
		in, out := &in.Hosted, &out.Hosted
		*out = new(HostedClusterManagerConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerDeployOption.
func (in *ClusterManagerDeployOption) DeepCopy() *ClusterManagerDeployOption {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerDeployOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerList) DeepCopyInto(out *ClusterManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerList.
func (in *ClusterManagerList) DeepCopy() *ClusterManagerList {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerSpec) DeepCopyInto(out *ClusterManagerSpec) {
	*out = *in
	in.NodePlacement.DeepCopyInto(&out.NodePlacement)
	in.DeployOption.DeepCopyInto(&out.DeployOption)
	if in.RegistrationConfiguration != nil {
		in, out := &in.RegistrationConfiguration, &out.RegistrationConfiguration
		*out = new(RegistrationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerSpec.
func (in *ClusterManagerSpec) DeepCopy() *ClusterManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerStatus) DeepCopyInto(out *ClusterManagerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Generations != nil {
		in, out := &in.Generations, &out.Generations
		*out = make([]GenerationStatus, len(*in))
		copy(*out, *in)
	}
	if in.RelatedResources != nil {
		in, out := &in.RelatedResources, &out.RelatedResources
		*out = make([]RelatedResourceMeta, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerStatus.
func (in *ClusterManagerStatus) DeepCopy() *ClusterManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGate) DeepCopyInto(out *FeatureGate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGate.
func (in *FeatureGate) DeepCopy() *FeatureGate {
	if in == nil {
		return nil
	}
	out := new(FeatureGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenerationStatus) DeepCopyInto(out *GenerationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenerationStatus.
func (in *GenerationStatus) DeepCopy() *GenerationStatus {
	if in == nil {
		return nil
	}
	out := new(GenerationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedClusterManagerConfiguration) DeepCopyInto(out *HostedClusterManagerConfiguration) {
	*out = *in
	out.RegistrationWebhookConfiguration = in.RegistrationWebhookConfiguration
	out.WorkWebhookConfiguration = in.WorkWebhookConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedClusterManagerConfiguration.
func (in *HostedClusterManagerConfiguration) DeepCopy() *HostedClusterManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(HostedClusterManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Klusterlet) DeepCopyInto(out *Klusterlet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Klusterlet.
func (in *Klusterlet) DeepCopy() *Klusterlet {
	if in == nil {
		return nil
	}
	out := new(Klusterlet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Klusterlet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletDeployOption) DeepCopyInto(out *KlusterletDeployOption) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletDeployOption.
func (in *KlusterletDeployOption) DeepCopy() *KlusterletDeployOption {
	if in == nil {
		return nil
	}
	out := new(KlusterletDeployOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletList) DeepCopyInto(out *KlusterletList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Klusterlet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletList.
func (in *KlusterletList) DeepCopy() *KlusterletList {
	if in == nil {
		return nil
	}
	out := new(KlusterletList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KlusterletList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletSpec) DeepCopyInto(out *KlusterletSpec) {
	*out = *in
	if in.ExternalServerURLs != nil {
		in, out := &in.ExternalServerURLs, &out.ExternalServerURLs
		*out = make([]ServerURL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NodePlacement.DeepCopyInto(&out.NodePlacement)
	out.DeployOption = in.DeployOption
	if in.RegistrationConfiguration != nil {
		in, out := &in.RegistrationConfiguration, &out.RegistrationConfiguration
		*out = new(RegistrationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletSpec.
func (in *KlusterletSpec) DeepCopy() *KlusterletSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletStatus) DeepCopyInto(out *KlusterletStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Generations != nil {
		in, out := &in.Generations, &out.Generations
		*out = make([]GenerationStatus, len(*in))
		copy(*out, *in)
	}
	if in.RelatedResources != nil {
		in, out := &in.RelatedResources, &out.RelatedResources
		*out = make([]RelatedResourceMeta, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletStatus.
func (in *KlusterletStatus) DeepCopy() *KlusterletStatus {
	if in == nil {
		return nil
	}
	out := new(KlusterletStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePlacement) DeepCopyInto(out *NodePlacement) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePlacement.
func (in *NodePlacement) DeepCopy() *NodePlacement {
	if in == nil {
		return nil
	}
	out := new(NodePlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrationConfiguration) DeepCopyInto(out *RegistrationConfiguration) {
	*out = *in
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make([]FeatureGate, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrationConfiguration.
func (in *RegistrationConfiguration) DeepCopy() *RegistrationConfiguration {
	if in == nil {
		return nil
	}
	out := new(RegistrationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelatedResourceMeta) DeepCopyInto(out *RelatedResourceMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelatedResourceMeta.
func (in *RelatedResourceMeta) DeepCopy() *RelatedResourceMeta {
	if in == nil {
		return nil
	}
	out := new(RelatedResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerURL) DeepCopyInto(out *ServerURL) {
	*out = *in
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerURL.
func (in *ServerURL) DeepCopy() *ServerURL {
	if in == nil {
		return nil
	}
	out := new(ServerURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfiguration) DeepCopyInto(out *WebhookConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfiguration.
func (in *WebhookConfiguration) DeepCopy() *WebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}
