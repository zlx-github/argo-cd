// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAuthConfig) DeepCopyInto(out *AWSAuthConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAuthConfig.
func (in *AWSAuthConfig) DeepCopy() *AWSAuthConfig {
	if in == nil {
		return nil
	}
	out := new(AWSAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppProject) DeepCopyInto(out *AppProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppProject.
func (in *AppProject) DeepCopy() *AppProject {
	if in == nil {
		return nil
	}
	out := new(AppProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppProjectList) DeepCopyInto(out *AppProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppProjectList.
func (in *AppProjectList) DeepCopy() *AppProjectList {
	if in == nil {
		return nil
	}
	out := new(AppProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppProjectSpec) DeepCopyInto(out *AppProjectSpec) {
	*out = *in
	if in.SourceRepos != nil {
		in, out := &in.SourceRepos, &out.SourceRepos
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]ApplicationDestination, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]ProjectRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterResourceWhitelist != nil {
		in, out := &in.ClusterResourceWhitelist, &out.ClusterResourceWhitelist
		*out = make([]v1.GroupKind, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceResourceBlacklist != nil {
		in, out := &in.NamespaceResourceBlacklist, &out.NamespaceResourceBlacklist
		*out = make([]v1.GroupKind, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppProjectSpec.
func (in *AppProjectSpec) DeepCopy() *AppProjectSpec {
	if in == nil {
		return nil
	}
	out := new(AppProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		if *in == nil {
			*out = nil
		} else {
			*out = new(Operation)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationCondition) DeepCopyInto(out *ApplicationCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationCondition.
func (in *ApplicationCondition) DeepCopy() *ApplicationCondition {
	if in == nil {
		return nil
	}
	out := new(ApplicationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDestination) DeepCopyInto(out *ApplicationDestination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDestination.
func (in *ApplicationDestination) DeepCopy() *ApplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ApplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSource) DeepCopyInto(out *ApplicationSource) {
	*out = *in
	if in.ComponentParameterOverrides != nil {
		in, out := &in.ComponentParameterOverrides, &out.ComponentParameterOverrides
		*out = make([]ComponentParameter, len(*in))
		copy(*out, *in)
	}
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationSourceHelm)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Kustomize != nil {
		in, out := &in.Kustomize, &out.Kustomize
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationSourceKustomize)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Ksonnet != nil {
		in, out := &in.Ksonnet, &out.Ksonnet
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationSourceKsonnet)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Directory != nil {
		in, out := &in.Directory, &out.Directory
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationSourceDirectory)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Plugin != nil {
		in, out := &in.Plugin, &out.Plugin
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationSourcePlugin)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSource.
func (in *ApplicationSource) DeepCopy() *ApplicationSource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourceDirectory) DeepCopyInto(out *ApplicationSourceDirectory) {
	*out = *in
	in.Jsonnet.DeepCopyInto(&out.Jsonnet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceDirectory.
func (in *ApplicationSourceDirectory) DeepCopy() *ApplicationSourceDirectory {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourceHelm) DeepCopyInto(out *ApplicationSourceHelm) {
	*out = *in
	if in.ValueFiles != nil {
		in, out := &in.ValueFiles, &out.ValueFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]HelmParameter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceHelm.
func (in *ApplicationSourceHelm) DeepCopy() *ApplicationSourceHelm {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceHelm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourceJsonnet) DeepCopyInto(out *ApplicationSourceJsonnet) {
	*out = *in
	if in.ExtVars != nil {
		in, out := &in.ExtVars, &out.ExtVars
		*out = make([]JsonnetVar, len(*in))
		copy(*out, *in)
	}
	if in.TLAs != nil {
		in, out := &in.TLAs, &out.TLAs
		*out = make([]JsonnetVar, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceJsonnet.
func (in *ApplicationSourceJsonnet) DeepCopy() *ApplicationSourceJsonnet {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceJsonnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourceKsonnet) DeepCopyInto(out *ApplicationSourceKsonnet) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]KsonnetParameter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceKsonnet.
func (in *ApplicationSourceKsonnet) DeepCopy() *ApplicationSourceKsonnet {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceKsonnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourceKustomize) DeepCopyInto(out *ApplicationSourceKustomize) {
	*out = *in
	if in.ImageTags != nil {
		in, out := &in.ImageTags, &out.ImageTags
		*out = make([]KustomizeImageTag, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceKustomize.
func (in *ApplicationSourceKustomize) DeepCopy() *ApplicationSourceKustomize {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceKustomize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourcePlugin) DeepCopyInto(out *ApplicationSourcePlugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourcePlugin.
func (in *ApplicationSourcePlugin) DeepCopy() *ApplicationSourcePlugin {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourcePlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	out.Destination = in.Destination
	if in.SyncPolicy != nil {
		in, out := &in.SyncPolicy, &out.SyncPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncPolicy)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.IgnoreDifferences != nil {
		in, out := &in.IgnoreDifferences, &out.IgnoreDifferences
		*out = make([]ResourceIgnoreDifferences, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceStatus, len(*in))
		copy(*out, *in)
	}
	in.Sync.DeepCopyInto(&out.Sync)
	out.Health = in.Health
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]RevisionHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ApplicationCondition, len(*in))
		copy(*out, *in)
	}
	in.ObservedAt.DeepCopyInto(&out.ObservedAt)
	if in.OperationState != nil {
		in, out := &in.OperationState, &out.OperationState
		if *in == nil {
			*out = nil
		} else {
			*out = new(OperationState)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationWatchEvent) DeepCopyInto(out *ApplicationWatchEvent) {
	*out = *in
	in.Application.DeepCopyInto(&out.Application)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationWatchEvent.
func (in *ApplicationWatchEvent) DeepCopy() *ApplicationWatchEvent {
	if in == nil {
		return nil
	}
	out := new(ApplicationWatchEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	in.ConnectionState.DeepCopyInto(&out.ConnectionState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	in.TLSClientConfig.DeepCopyInto(&out.TLSClientConfig)
	if in.AWSAuthConfig != nil {
		in, out := &in.AWSAuthConfig, &out.AWSAuthConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSAuthConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Command) DeepCopyInto(out *Command) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Command.
func (in *Command) DeepCopy() *Command {
	if in == nil {
		return nil
	}
	out := new(Command)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComparedTo) DeepCopyInto(out *ComparedTo) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	out.Destination = in.Destination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComparedTo.
func (in *ComparedTo) DeepCopy() *ComparedTo {
	if in == nil {
		return nil
	}
	out := new(ComparedTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentParameter) DeepCopyInto(out *ComponentParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentParameter.
func (in *ComponentParameter) DeepCopy() *ComponentParameter {
	if in == nil {
		return nil
	}
	out := new(ComponentParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigManagementPlugin) DeepCopyInto(out *ConfigManagementPlugin) {
	*out = *in
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		if *in == nil {
			*out = nil
		} else {
			*out = new(Command)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Generate.DeepCopyInto(&out.Generate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigManagementPlugin.
func (in *ConfigManagementPlugin) DeepCopy() *ConfigManagementPlugin {
	if in == nil {
		return nil
	}
	out := new(ConfigManagementPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionState) DeepCopyInto(out *ConnectionState) {
	*out = *in
	if in.ModifiedAt != nil {
		in, out := &in.ModifiedAt, &out.ModifiedAt
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionState.
func (in *ConnectionState) DeepCopy() *ConnectionState {
	if in == nil {
		return nil
	}
	out := new(ConnectionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthStatus) DeepCopyInto(out *HealthStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthStatus.
func (in *HealthStatus) DeepCopy() *HealthStatus {
	if in == nil {
		return nil
	}
	out := new(HealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmParameter) DeepCopyInto(out *HelmParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmParameter.
func (in *HelmParameter) DeepCopy() *HelmParameter {
	if in == nil {
		return nil
	}
	out := new(HelmParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepository) DeepCopyInto(out *HelmRepository) {
	*out = *in
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepository.
func (in *HelmRepository) DeepCopy() *HelmRepository {
	if in == nil {
		return nil
	}
	out := new(HelmRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfoItem) DeepCopyInto(out *InfoItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfoItem.
func (in *InfoItem) DeepCopy() *InfoItem {
	if in == nil {
		return nil
	}
	out := new(InfoItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTToken) DeepCopyInto(out *JWTToken) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTToken.
func (in *JWTToken) DeepCopy() *JWTToken {
	if in == nil {
		return nil
	}
	out := new(JWTToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonnetVar) DeepCopyInto(out *JsonnetVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonnetVar.
func (in *JsonnetVar) DeepCopy() *JsonnetVar {
	if in == nil {
		return nil
	}
	out := new(JsonnetVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KsonnetParameter) DeepCopyInto(out *KsonnetParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KsonnetParameter.
func (in *KsonnetParameter) DeepCopy() *KsonnetParameter {
	if in == nil {
		return nil
	}
	out := new(KsonnetParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizeImageTag) DeepCopyInto(out *KustomizeImageTag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizeImageTag.
func (in *KustomizeImageTag) DeepCopy() *KustomizeImageTag {
	if in == nil {
		return nil
	}
	out := new(KustomizeImageTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncOperation)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationState) DeepCopyInto(out *OperationState) {
	*out = *in
	in.Operation.DeepCopyInto(&out.Operation)
	if in.SyncResult != nil {
		in, out := &in.SyncResult, &out.SyncResult
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncOperationResult)
			(*in).DeepCopyInto(*out)
		}
	}
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	if in.FinishedAt != nil {
		in, out := &in.FinishedAt, &out.FinishedAt
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationState.
func (in *OperationState) DeepCopy() *OperationState {
	if in == nil {
		return nil
	}
	out := new(OperationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectRole) DeepCopyInto(out *ProjectRole) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.JWTTokens != nil {
		in, out := &in.JWTTokens, &out.JWTTokens
		*out = make([]JWTToken, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectRole.
func (in *ProjectRole) DeepCopy() *ProjectRole {
	if in == nil {
		return nil
	}
	out := new(ProjectRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	in.ConnectionState.DeepCopyInto(&out.ConnectionState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDiff) DeepCopyInto(out *ResourceDiff) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDiff.
func (in *ResourceDiff) DeepCopy() *ResourceDiff {
	if in == nil {
		return nil
	}
	out := new(ResourceDiff)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIgnoreDifferences) DeepCopyInto(out *ResourceIgnoreDifferences) {
	*out = *in
	if in.JSONPointers != nil {
		in, out := &in.JSONPointers, &out.JSONPointers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIgnoreDifferences.
func (in *ResourceIgnoreDifferences) DeepCopy() *ResourceIgnoreDifferences {
	if in == nil {
		return nil
	}
	out := new(ResourceIgnoreDifferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceNode) DeepCopyInto(out *ResourceNode) {
	*out = *in
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = make([]InfoItem, len(*in))
		copy(*out, *in)
	}
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]ResourceNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceNode.
func (in *ResourceNode) DeepCopy() *ResourceNode {
	if in == nil {
		return nil
	}
	out := new(ResourceNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceResult) DeepCopyInto(out *ResourceResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceResult.
func (in *ResourceResult) DeepCopy() *ResourceResult {
	if in == nil {
		return nil
	}
	out := new(ResourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
	out.Health = in.Health
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionHistory) DeepCopyInto(out *RevisionHistory) {
	*out = *in
	in.DeployedAt.DeepCopyInto(&out.DeployedAt)
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionHistory.
func (in *RevisionHistory) DeepCopy() *RevisionHistory {
	if in == nil {
		return nil
	}
	out := new(RevisionHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncOperation) DeepCopyInto(out *SyncOperation) {
	*out = *in
	if in.SyncStrategy != nil {
		in, out := &in.SyncStrategy, &out.SyncStrategy
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncStrategy)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]SyncOperationResource, len(*in))
		copy(*out, *in)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplicationSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncOperation.
func (in *SyncOperation) DeepCopy() *SyncOperation {
	if in == nil {
		return nil
	}
	out := new(SyncOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncOperationResource) DeepCopyInto(out *SyncOperationResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncOperationResource.
func (in *SyncOperationResource) DeepCopy() *SyncOperationResource {
	if in == nil {
		return nil
	}
	out := new(SyncOperationResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncOperationResult) DeepCopyInto(out *SyncOperationResult) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*ResourceResult, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ResourceResult)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncOperationResult.
func (in *SyncOperationResult) DeepCopy() *SyncOperationResult {
	if in == nil {
		return nil
	}
	out := new(SyncOperationResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncPolicy) DeepCopyInto(out *SyncPolicy) {
	*out = *in
	if in.Automated != nil {
		in, out := &in.Automated, &out.Automated
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncPolicyAutomated)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncPolicy.
func (in *SyncPolicy) DeepCopy() *SyncPolicy {
	if in == nil {
		return nil
	}
	out := new(SyncPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncPolicyAutomated) DeepCopyInto(out *SyncPolicyAutomated) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncPolicyAutomated.
func (in *SyncPolicyAutomated) DeepCopy() *SyncPolicyAutomated {
	if in == nil {
		return nil
	}
	out := new(SyncPolicyAutomated)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncStatus) DeepCopyInto(out *SyncStatus) {
	*out = *in
	in.ComparedTo.DeepCopyInto(&out.ComparedTo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncStatus.
func (in *SyncStatus) DeepCopy() *SyncStatus {
	if in == nil {
		return nil
	}
	out := new(SyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncStrategy) DeepCopyInto(out *SyncStrategy) {
	*out = *in
	if in.Apply != nil {
		in, out := &in.Apply, &out.Apply
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncStrategyApply)
			**out = **in
		}
	}
	if in.Hook != nil {
		in, out := &in.Hook, &out.Hook
		if *in == nil {
			*out = nil
		} else {
			*out = new(SyncStrategyHook)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncStrategy.
func (in *SyncStrategy) DeepCopy() *SyncStrategy {
	if in == nil {
		return nil
	}
	out := new(SyncStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncStrategyApply) DeepCopyInto(out *SyncStrategyApply) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncStrategyApply.
func (in *SyncStrategyApply) DeepCopy() *SyncStrategyApply {
	if in == nil {
		return nil
	}
	out := new(SyncStrategyApply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncStrategyHook) DeepCopyInto(out *SyncStrategyHook) {
	*out = *in
	out.SyncStrategyApply = in.SyncStrategyApply
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncStrategyHook.
func (in *SyncStrategyHook) DeepCopy() *SyncStrategyHook {
	if in == nil {
		return nil
	}
	out := new(SyncStrategyHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSClientConfig) DeepCopyInto(out *TLSClientConfig) {
	*out = *in
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSClientConfig.
func (in *TLSClientConfig) DeepCopy() *TLSClientConfig {
	if in == nil {
		return nil
	}
	out := new(TLSClientConfig)
	in.DeepCopyInto(out)
	return out
}
