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

package config

// FilesGeneratorOptionsClickHouse specifies options for clickhouse configuration generator
type FilesGeneratorOptionsClickHouse struct {
	RemoteServersOptions *RemoteServersOptions
}

// NewConfigFilesGeneratorOptionsClickHouse creates new options for clickhouse configuration generator
func NewConfigFilesGeneratorOptionsClickHouse() *FilesGeneratorOptionsClickHouse {
	return &FilesGeneratorOptionsClickHouse{}
}

// GetRemoteServersOptions gets remote-servers generator options
func (o *FilesGeneratorOptionsClickHouse) GetRemoteServersOptions() *RemoteServersOptions {
	if o == nil {
		return nil
	}
	return o.RemoteServersOptions
}

// SetRemoteServersOptions sets remote-servers generator options
func (o *FilesGeneratorOptionsClickHouse) SetRemoteServersOptions(opts *RemoteServersOptions) *FilesGeneratorOptionsClickHouse {
	if o == nil {
		return nil
	}
	o.RemoteServersOptions = opts

	return o
}

// defaultConfigFilesGeneratorOptionsClickHouse creates new default options for clickhouse config generator
func defaultConfigFilesGeneratorOptionsClickHouse() *FilesGeneratorOptionsClickHouse {
	return NewConfigFilesGeneratorOptionsClickHouse()
}
