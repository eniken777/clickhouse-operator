// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	core_v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiCluster) DeepCopyInto(out *ChiCluster) {
	*out = *in
	in.Layout.DeepCopyInto(&out.Layout)
	in.Deployment.DeepCopyInto(&out.Deployment)
	out.Address = in.Address
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiCluster.
func (in *ChiCluster) DeepCopy() *ChiCluster {
	if in == nil {
		return nil
	}
	out := new(ChiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterAddress) DeepCopyInto(out *ChiClusterAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterAddress.
func (in *ChiClusterAddress) DeepCopy() *ChiClusterAddress {
	if in == nil {
		return nil
	}
	out := new(ChiClusterAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterLayout) DeepCopyInto(out *ChiClusterLayout) {
	*out = *in
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make([]ChiClusterLayoutShard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterLayout.
func (in *ChiClusterLayout) DeepCopy() *ChiClusterLayout {
	if in == nil {
		return nil
	}
	out := new(ChiClusterLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterLayoutShard) DeepCopyInto(out *ChiClusterLayoutShard) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ChiClusterLayoutShardReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Address = in.Address
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterLayoutShard.
func (in *ChiClusterLayoutShard) DeepCopy() *ChiClusterLayoutShard {
	if in == nil {
		return nil
	}
	out := new(ChiClusterLayoutShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterLayoutShardAddress) DeepCopyInto(out *ChiClusterLayoutShardAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterLayoutShardAddress.
func (in *ChiClusterLayoutShardAddress) DeepCopy() *ChiClusterLayoutShardAddress {
	if in == nil {
		return nil
	}
	out := new(ChiClusterLayoutShardAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterLayoutShardReplica) DeepCopyInto(out *ChiClusterLayoutShardReplica) {
	*out = *in
	in.Deployment.DeepCopyInto(&out.Deployment)
	out.Address = in.Address
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterLayoutShardReplica.
func (in *ChiClusterLayoutShardReplica) DeepCopy() *ChiClusterLayoutShardReplica {
	if in == nil {
		return nil
	}
	out := new(ChiClusterLayoutShardReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterLayoutShardReplicaAddress) DeepCopyInto(out *ChiClusterLayoutShardReplicaAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterLayoutShardReplicaAddress.
func (in *ChiClusterLayoutShardReplicaAddress) DeepCopy() *ChiClusterLayoutShardReplicaAddress {
	if in == nil {
		return nil
	}
	out := new(ChiClusterLayoutShardReplicaAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiConfiguration) DeepCopyInto(out *ChiConfiguration) {
	*out = *in
	in.Zookeeper.DeepCopyInto(&out.Zookeeper)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ChiCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiConfiguration.
func (in *ChiConfiguration) DeepCopy() *ChiConfiguration {
	if in == nil {
		return nil
	}
	out := new(ChiConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiConfigurationZookeeper) DeepCopyInto(out *ChiConfigurationZookeeper) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ChiConfigurationZookeeperNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiConfigurationZookeeper.
func (in *ChiConfigurationZookeeper) DeepCopy() *ChiConfigurationZookeeper {
	if in == nil {
		return nil
	}
	out := new(ChiConfigurationZookeeper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiConfigurationZookeeperNode) DeepCopyInto(out *ChiConfigurationZookeeperNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiConfigurationZookeeperNode.
func (in *ChiConfigurationZookeeperNode) DeepCopy() *ChiConfigurationZookeeperNode {
	if in == nil {
		return nil
	}
	out := new(ChiConfigurationZookeeperNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDefaults) DeepCopyInto(out *ChiDefaults) {
	*out = *in
	out.DistributedDDL = in.DistributedDDL
	in.Deployment.DeepCopyInto(&out.Deployment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDefaults.
func (in *ChiDefaults) DeepCopy() *ChiDefaults {
	if in == nil {
		return nil
	}
	out := new(ChiDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDeployment) DeepCopyInto(out *ChiDeployment) {
	*out = *in
	in.Zone.DeepCopyInto(&out.Zone)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDeployment.
func (in *ChiDeployment) DeepCopy() *ChiDeployment {
	if in == nil {
		return nil
	}
	out := new(ChiDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDeploymentZone) DeepCopyInto(out *ChiDeploymentZone) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDeploymentZone.
func (in *ChiDeploymentZone) DeepCopy() *ChiDeploymentZone {
	if in == nil {
		return nil
	}
	out := new(ChiDeploymentZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDistributedDDL) DeepCopyInto(out *ChiDistributedDDL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDistributedDDL.
func (in *ChiDistributedDDL) DeepCopy() *ChiDistributedDDL {
	if in == nil {
		return nil
	}
	out := new(ChiDistributedDDL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodTemplate) DeepCopyInto(out *ChiPodTemplate) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]core_v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]core_v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodTemplate.
func (in *ChiPodTemplate) DeepCopy() *ChiPodTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiSpec) DeepCopyInto(out *ChiSpec) {
	*out = *in
	in.Defaults.DeepCopyInto(&out.Defaults)
	in.Configuration.DeepCopyInto(&out.Configuration)
	in.Templates.DeepCopyInto(&out.Templates)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiSpec.
func (in *ChiSpec) DeepCopy() *ChiSpec {
	if in == nil {
		return nil
	}
	out := new(ChiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiStatus) DeepCopyInto(out *ChiStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiStatus.
func (in *ChiStatus) DeepCopy() *ChiStatus {
	if in == nil {
		return nil
	}
	out := new(ChiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplates) DeepCopyInto(out *ChiTemplates) {
	*out = *in
	if in.PodTemplates != nil {
		in, out := &in.PodTemplates, &out.PodTemplates
		*out = make([]ChiPodTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]ChiVolumeClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplates.
func (in *ChiTemplates) DeepCopy() *ChiTemplates {
	if in == nil {
		return nil
	}
	out := new(ChiTemplates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiVolumeClaimTemplate) DeepCopyInto(out *ChiVolumeClaimTemplate) {
	*out = *in
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiVolumeClaimTemplate.
func (in *ChiVolumeClaimTemplate) DeepCopy() *ChiVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallation) DeepCopyInto(out *ClickHouseInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallation.
func (in *ClickHouseInstallation) DeepCopy() *ClickHouseInstallation {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationList) DeepCopyInto(out *ClickHouseInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationList.
func (in *ClickHouseInstallationList) DeepCopy() *ClickHouseInstallationList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
