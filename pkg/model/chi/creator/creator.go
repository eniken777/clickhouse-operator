// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package creator

import (
	log "github.com/altinity/clickhouse-operator/pkg/announcer"
	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	model "github.com/altinity/clickhouse-operator/pkg/model/chi"
	"github.com/altinity/clickhouse-operator/pkg/model/chi/config"
	core "k8s.io/api/core/v1"
)

type labeler interface {
	GetConfigMapCHICommon() map[string]string
	GetConfigMapCHICommonUsers() map[string]string
	GetConfigMapHost(host *api.ChiHost) map[string]string
	GetClusterScope(cluster api.ICluster) map[string]string
	GetPV(pv *core.PersistentVolume, host *api.ChiHost) map[string]string
	GetHostScope(host *api.ChiHost, applySupplementaryServiceLabels bool) map[string]string
	GetPVC(
		pvc *core.PersistentVolumeClaim,
		host *api.ChiHost,
		template *api.VolumeClaimTemplate,
	) map[string]string
	GetServiceCHI(chi api.IChi) map[string]string
	GetSelectorCHIScopeReady() map[string]string
	GetServiceCluster(cluster api.ICluster) map[string]string
	GetServiceHost(host *api.ChiHost) map[string]string
	GetServiceShard(shard api.IShard) map[string]string
	GetHostScopeReady(host *api.ChiHost, applySupplementaryServiceLabels bool) map[string]string
}

type annotator interface {
	GetConfigMapCHICommon() map[string]string
	GetConfigMapCHICommonUsers() map[string]string
	GetConfigMapHost(host *api.ChiHost) map[string]string
	GetClusterScope(cluster api.ICluster) map[string]string
	GetPV(pv *core.PersistentVolume, host *api.ChiHost) map[string]string
	GetHostScope(host *api.ChiHost) map[string]string
	GetPVC(
		pvc *core.PersistentVolumeClaim,
		host *api.ChiHost,
		template *api.VolumeClaimTemplate,
	) map[string]string
	GetServiceCHI(chi api.IChi) map[string]string
	GetServiceCluster(cluster api.ICluster) map[string]string
	GetServiceHost(host *api.ChiHost) map[string]string
	GetServiceShard(shard api.IShard) map[string]string
}

// Creator specifies creator object
type Creator struct {
	chi                  api.IChi
	configFilesGenerator *config.ClickHouseConfigFilesGenerator
	labels               labeler
	annotations          annotator
	a                    log.Announcer
}

// NewCreator creates new Creator object
func NewCreator(chi api.IChi, configFilesGenerator *config.ClickHouseConfigFilesGenerator) *Creator {
	return &Creator{
		chi:                  chi,
		configFilesGenerator: configFilesGenerator,
		labels:               model.NewLabeler(chi),
		annotations:          model.NewAnnotator(chi),
		a:                    log.M(chi),
	}
}
