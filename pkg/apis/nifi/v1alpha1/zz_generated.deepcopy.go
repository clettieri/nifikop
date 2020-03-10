// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	metav1 "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapProperties) DeepCopyInto(out *BootstrapProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapProperties.
func (in *BootstrapProperties) DeepCopy() *BootstrapProperties {
	if in == nil {
		return nil
	}
	out := new(BootstrapProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReference) DeepCopyInto(out *ClusterReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReference.
func (in *ClusterReference) DeepCopy() *ClusterReference {
	if in == nil {
		return nil
	}
	out := new(ClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalListenerConfig) DeepCopyInto(out *ExternalListenerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalListenerConfig.
func (in *ExternalListenerConfig) DeepCopy() *ExternalListenerConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalListenerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GracefulActionState) DeepCopyInto(out *GracefulActionState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GracefulActionState.
func (in *GracefulActionState) DeepCopy() *GracefulActionState {
	if in == nil {
		return nil
	}
	out := new(GracefulActionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalListenerConfig) DeepCopyInto(out *InternalListenerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalListenerConfig.
func (in *InternalListenerConfig) DeepCopy() *InternalListenerConfig {
	if in == nil {
		return nil
	}
	out := new(InternalListenerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LdapConfiguration) DeepCopyInto(out *LdapConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LdapConfiguration.
func (in *LdapConfiguration) DeepCopy() *LdapConfiguration {
	if in == nil {
		return nil
	}
	out := new(LdapConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenersConfig) DeepCopyInto(out *ListenersConfig) {
	*out = *in
	if in.InternalListeners != nil {
		in, out := &in.InternalListeners, &out.InternalListeners
		*out = make([]InternalListenerConfig, len(*in))
		copy(*out, *in)
	}
	if in.SSLSecrets != nil {
		in, out := &in.SSLSecrets, &out.SSLSecrets
		*out = new(SSLSecrets)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenersConfig.
func (in *ListenersConfig) DeepCopy() *ListenersConfig {
	if in == nil {
		return nil
	}
	out := new(ListenersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiCluster) DeepCopyInto(out *NifiCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiCluster.
func (in *NifiCluster) DeepCopy() *NifiCluster {
	if in == nil {
		return nil
	}
	out := new(NifiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NifiCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiClusterList) DeepCopyInto(out *NifiClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NifiCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiClusterList.
func (in *NifiClusterList) DeepCopy() *NifiClusterList {
	if in == nil {
		return nil
	}
	out := new(NifiClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NifiClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiClusterSpec) DeepCopyInto(out *NifiClusterSpec) {
	*out = *in
	in.ListenersConfig.DeepCopyInto(&out.ListenersConfig)
	out.ReadOnlyConfig = in.ReadOnlyConfig
	if in.NodeConfigGroups != nil {
		in, out := &in.NodeConfigGroups, &out.NodeConfigGroups
		*out = make(map[string]NodeConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LdapConfiguration = in.LdapConfiguration
	out.VaultConfig = in.VaultConfig
	out.NifiClusterTaskSpec = in.NifiClusterTaskSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiClusterSpec.
func (in *NifiClusterSpec) DeepCopy() *NifiClusterSpec {
	if in == nil {
		return nil
	}
	out := new(NifiClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiClusterStatus) DeepCopyInto(out *NifiClusterStatus) {
	*out = *in
	if in.NodesState != nil {
		in, out := &in.NodesState, &out.NodesState
		*out = make(map[string]NodeState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.RollingUpgrade = in.RollingUpgrade
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiClusterStatus.
func (in *NifiClusterStatus) DeepCopy() *NifiClusterStatus {
	if in == nil {
		return nil
	}
	out := new(NifiClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiClusterTaskSpec) DeepCopyInto(out *NifiClusterTaskSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiClusterTaskSpec.
func (in *NifiClusterTaskSpec) DeepCopy() *NifiClusterTaskSpec {
	if in == nil {
		return nil
	}
	out := new(NifiClusterTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiProperties) DeepCopyInto(out *NifiProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiProperties.
func (in *NifiProperties) DeepCopy() *NifiProperties {
	if in == nil {
		return nil
	}
	out := new(NifiProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiUser) DeepCopyInto(out *NifiUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiUser.
func (in *NifiUser) DeepCopy() *NifiUser {
	if in == nil {
		return nil
	}
	out := new(NifiUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NifiUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiUserList) DeepCopyInto(out *NifiUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NifiUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiUserList.
func (in *NifiUserList) DeepCopy() *NifiUserList {
	if in == nil {
		return nil
	}
	out := new(NifiUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NifiUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiUserSpec) DeepCopyInto(out *NifiUserSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiUserSpec.
func (in *NifiUserSpec) DeepCopy() *NifiUserSpec {
	if in == nil {
		return nil
	}
	out := new(NifiUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NifiUserStatus) DeepCopyInto(out *NifiUserStatus) {
	*out = *in
	if in.ACLs != nil {
		in, out := &in.ACLs, &out.ACLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NifiUserStatus.
func (in *NifiUserStatus) DeepCopy() *NifiUserStatus {
	if in == nil {
		return nil
	}
	out := new(NifiUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	if in.ReadOnlyConfig != nil {
		in, out := &in.ReadOnlyConfig, &out.ReadOnlyConfig
		*out = new(ReadOnlyConfig)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = new(NodeConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	if in.RunAsUser != nil {
		in, out := &in.RunAsUser, &out.RunAsUser
		*out = new(int64)
		**out = **in
	}
	if in.IsNode != nil {
		in, out := &in.IsNode, &out.IsNode
		*out = new(bool)
		**out = **in
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageConfigs != nil {
		in, out := &in.StorageConfigs, &out.StorageConfigs
		*out = make([]StorageConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourcesRequirements != nil {
		in, out := &in.ResourcesRequirements, &out.ResourcesRequirements
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeAnnotations != nil {
		in, out := &in.NodeAnnotations, &out.NodeAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeState) DeepCopyInto(out *NodeState) {
	*out = *in
	out.GracefulActionState = in.GracefulActionState
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeState.
func (in *NodeState) DeepCopy() *NodeState {
	if in == nil {
		return nil
	}
	out := new(NodeState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadOnlyConfig) DeepCopyInto(out *ReadOnlyConfig) {
	*out = *in
	out.NifiProperties = in.NifiProperties
	out.ZookeeperProperties = in.ZookeeperProperties
	out.BootstrapProperties = in.BootstrapProperties
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadOnlyConfig.
func (in *ReadOnlyConfig) DeepCopy() *ReadOnlyConfig {
	if in == nil {
		return nil
	}
	out := new(ReadOnlyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpgradeStatus) DeepCopyInto(out *RollingUpgradeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpgradeStatus.
func (in *RollingUpgradeStatus) DeepCopy() *RollingUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(RollingUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLSecrets) DeepCopyInto(out *SSLSecrets) {
	*out = *in
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(metav1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLSecrets.
func (in *SSLSecrets) DeepCopy() *SSLSecrets {
	if in == nil {
		return nil
	}
	out := new(SSLSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.PVCSpec != nil {
		in, out := &in.PVCSpec, &out.PVCSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConfig) DeepCopyInto(out *VaultConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConfig.
func (in *VaultConfig) DeepCopy() *VaultConfig {
	if in == nil {
		return nil
	}
	out := new(VaultConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperProperties) DeepCopyInto(out *ZookeeperProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperProperties.
func (in *ZookeeperProperties) DeepCopy() *ZookeeperProperties {
	if in == nil {
		return nil
	}
	out := new(ZookeeperProperties)
	in.DeepCopyInto(out)
	return out
}
