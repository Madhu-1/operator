/*
Copyright 2021.

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

package fake

import (
	replicationlib "github.com/kube-storage/spec/lib/go/replication"
)

// ReplicationClient to fake replication operations.
type ReplicationClient struct {
	// EnableVolumeReplicationMock mocks EnableVolumeReplication RPC call.
	EnableVolumeReplicationMock func(volumeID string, secrets, parameters map[string]string) (*replicationlib.EnableVolumeReplicationResponse, error)
	// DisableVolumeReplicationMock mocks DisableVolumeReplication RPC call.
	DisableVolumeReplicationMock func(volumeID string, secrets, parameters map[string]string) (*replicationlib.DisableVolumeReplicationResponse, error)
	// PromoteVolumeMock mocks PromoteVolume RPC call.
	PromoteVolumeMock func(volumeID string, force bool, secrets, parameters map[string]string) (*replicationlib.PromoteVolumeResponse, error)
	// DemoteVolumeMock mocks DemoteVolume RPC call.
	DemoteVolumeMock func(volumeID string, secrets, parameters map[string]string) (*replicationlib.DemoteVolumeResponse, error)
	// ResyncVolumeMock mocks ResyncVolume RPC call.
	ResyncVolumeMock func(volumeID string, secrets, parameters map[string]string) (*replicationlib.ResyncVolumeResponse, error)
}

// EnableVolumeReplication calls EnableVolumeReplicationMock mock function.
func (rc *ReplicationClient) EnableVolumeReplication(volumeID string, secrets, parameters map[string]string) (*replicationlib.EnableVolumeReplicationResponse, error) {
	return rc.EnableVolumeReplicationMock(volumeID, secrets, parameters)
}

// DisableVolumeReplication calls DisableVolumeReplicationMock mock function.
func (rc *ReplicationClient) DisableVolumeReplication(volumeID string, secrets, parameters map[string]string) (*replicationlib.DisableVolumeReplicationResponse, error) {
	return rc.DisableVolumeReplicationMock(volumeID, secrets, parameters)
}

// PromoteVolume calls PromoteVolumeMock mock function.
func (rc *ReplicationClient) PromoteVolume(volumeID string, force bool, secrets, parameters map[string]string) (*replicationlib.PromoteVolumeResponse, error) {
	return rc.PromoteVolumeMock(volumeID, force, secrets, parameters)
}

// DemoteVolume calls DemoteVolumeMock mock function.
func (rc *ReplicationClient) DemoteVolume(volumeID string, secrets, parameters map[string]string) (*replicationlib.DemoteVolumeResponse, error) {
	return rc.DemoteVolumeMock(volumeID, secrets, parameters)
}

// ResyncVolume calls ResyncVolumeMock function.
func (rc *ReplicationClient) ResyncVolume(volumeID string, secrets, parameters map[string]string) (*replicationlib.ResyncVolumeResponse, error) {
	return rc.ResyncVolumeMock(volumeID, secrets, parameters)
}
