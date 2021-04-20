// Copyright 2021 Google LLC. All Rights Reserved.
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
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Instance) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.MaxDiskSize) {
		if err := r.MaxDiskSize.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CurrentDiskSize) {
		if err := r.CurrentDiskSize.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DiskEncryptionConfiguration) {
		if err := r.DiskEncryptionConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.FailoverReplica) {
		if err := r.FailoverReplica.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MasterInstance) {
		if err := r.MasterInstance.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReplicaConfiguration) {
		if err := r.ReplicaConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ScheduledMaintenance) {
		if err := r.ScheduledMaintenance.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Settings) {
		if err := r.Settings.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ServerCaCert) {
		if err := r.ServerCaCert.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.OnPremisesConfiguration) {
		if err := r.OnPremisesConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DiskEncryptionStatus) {
		if err := r.DiskEncryptionStatus.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceMaxDiskSize) validate() error {
	return nil
}
func (r *InstanceCurrentDiskSize) validate() error {
	return nil
}
func (r *InstanceDiskEncryptionConfiguration) validate() error {
	return nil
}
func (r *InstanceFailoverReplica) validate() error {
	if !dcl.IsEmptyValueIndirect(r.FailoverInstance) {
		if err := r.FailoverInstance.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceFailoverReplicaFailoverInstance) validate() error {
	return nil
}
func (r *InstanceIPAddresses) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TimeToRetire) {
		if err := r.TimeToRetire.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceIPAddressesTimeToRetire) validate() error {
	return nil
}
func (r *InstanceMasterInstance) validate() error {
	return nil
}
func (r *InstanceReplicaConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MysqlReplicaConfiguration) {
		if err := r.MysqlReplicaConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReplicaPoolConfiguration) {
		if err := r.ReplicaPoolConfiguration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceReplicaConfigurationMysqlReplicaConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MasterHeartbeatPeriod) {
		if err := r.MasterHeartbeatPeriod.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) validate() error {
	return nil
}
func (r *InstanceReplicaConfigurationReplicaPoolConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StaticPoolConfiguration) {
		if err := r.StaticPoolConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AutoscalingPoolConfiguration) {
		if err := r.AutoscalingPoolConfiguration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) validate() error {
	return nil
}
func (r *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) validate() error {
	return nil
}
func (r *InstanceScheduledMaintenance) validate() error {
	if !dcl.IsEmptyValueIndirect(r.StartTime) {
		if err := r.StartTime.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceScheduledMaintenanceStartTime) validate() error {
	return nil
}
func (r *InstanceSettings) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SettingsVersion) {
		if err := r.SettingsVersion.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StorageAutoResizeLimit) {
		if err := r.StorageAutoResizeLimit.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IPConfiguration) {
		if err := r.IPConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LocationPreference) {
		if err := r.LocationPreference.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MaintenanceWindow) {
		if err := r.MaintenanceWindow.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BackupConfiguration) {
		if err := r.BackupConfiguration.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DataDiskSizeGb) {
		if err := r.DataDiskSizeGb.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ActiveDirectoryConfig) {
		if err := r.ActiveDirectoryConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.InsightsConfig) {
		if err := r.InsightsConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceSettingsSettingsVersion) validate() error {
	return nil
}
func (r *InstanceSettingsStorageAutoResizeLimit) validate() error {
	return nil
}
func (r *InstanceSettingsIPConfiguration) validate() error {
	return nil
}
func (r *InstanceSettingsIPConfigurationAuthorizedNetworks) validate() error {
	return nil
}
func (r *InstanceSettingsLocationPreference) validate() error {
	return nil
}
func (r *InstanceSettingsDatabaseFlags) validate() error {
	return nil
}
func (r *InstanceSettingsMaintenanceWindow) validate() error {
	return nil
}
func (r *InstanceSettingsBackupConfiguration) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BackupRetentionSettings) {
		if err := r.BackupRetentionSettings.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InstanceSettingsBackupConfigurationBackupRetentionSettings) validate() error {
	return nil
}
func (r *InstanceSettingsDataDiskSizeGb) validate() error {
	return nil
}
func (r *InstanceSettingsActiveDirectoryConfig) validate() error {
	return nil
}
func (r *InstanceSettingsDenyMaintenancePeriods) validate() error {
	return nil
}
func (r *InstanceSettingsInsightsConfig) validate() error {
	return nil
}
func (r *InstanceReplicaInstances) validate() error {
	return nil
}
func (r *InstanceServerCaCert) validate() error {
	return nil
}
func (r *InstanceOnPremisesConfiguration) validate() error {
	return nil
}
func (r *InstanceDiskEncryptionStatus) validate() error {
	return nil
}

func instanceGetURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

func instanceListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances/", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func instanceCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/instances", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil

}

func instanceDeleteURL(userBasePath string, r *Instance) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, params), nil
}

// instanceApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type instanceApiOperation interface {
	do(context.Context, *Instance, *Client) error
}

// newUpdateInstanceUpdateRequest creates a request for an
// Instance resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInstanceUpdateRequest(ctx context.Context, f *Instance, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		req["state"] = v
	}
	if v, err := expandInstanceReplicaInstancesSlice(c, f.ReplicaInstances); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaInstances into replicaInstances: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["replicaInstances"] = v
	}
	if v, err := expandInstanceServerCaCert(c, f.ServerCaCert); err != nil {
		return nil, fmt.Errorf("error expanding ServerCaCert into serverCaCert: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["serverCaCert"] = v
	}
	if v := f.IPv6Address; !dcl.IsEmptyValueIndirect(v) {
		req["ipv6Address"] = v
	}
	if v := f.ServiceAccountEmailAddress; !dcl.IsEmptyValueIndirect(v) {
		req["serviceAccountEmailAddress"] = v
	}
	if v, err := expandInstanceOnPremisesConfiguration(c, f.OnPremisesConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding OnPremisesConfiguration into onPremisesConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["onPremisesConfiguration"] = v
	}
	if v := f.SuspensionReason; !dcl.IsEmptyValueIndirect(v) {
		req["suspensionReason"] = v
	}
	if v, err := expandInstanceDiskEncryptionStatus(c, f.DiskEncryptionStatus); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionStatus into diskEncryptionStatus: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["diskEncryptionStatus"] = v
	}
	if v := f.InstanceUid; !dcl.IsEmptyValueIndirect(v) {
		req["instanceUid"] = v
	}
	return req, nil
}

// marshalUpdateInstanceUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateInstanceUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInstanceUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateInstanceUpdateOperation) do(ctx context.Context, r *Instance, c *Client) error {
	_, err := c.GetInstance(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateInstanceUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateInstanceUpdateRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listInstanceRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := instanceListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InstanceMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listInstanceOperation struct {
	Instances []map[string]interface{} `json:"instances"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listInstance(ctx context.Context, project, pageToken string, pageSize int32) ([]*Instance, string, error) {
	b, err := c.listInstanceRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInstanceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Instance
	for _, v := range m.Instances {
		res := flattenInstance(c, v)
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInstance(ctx context.Context, f func(*Instance) bool, resources []*Instance) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInstance(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteInstanceOperation struct{}

func (op *deleteInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {

	_, err := c.GetInstance(ctx, r.urlNormalized())

	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("Instance not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetInstance checking for existence. error: %v", err)
		return err
	}

	u, err := instanceDeleteURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetInstance(ctx, r.urlNormalized())
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createInstanceOperation struct {
	response map[string]interface{}
}

func (op *createInstanceOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInstanceOperation) do(ctx context.Context, r *Instance, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := instanceCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.SQLOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(ctx, c.Config, "https://www.googleapis.com/sql/v1beta4/", "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetInstance(ctx, r.urlNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInstanceRaw(ctx context.Context, r *Instance) ([]byte, error) {

	u, err := instanceGetURL(c.Config.BasePath, r.urlNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) instanceDiffsForRawDesired(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (initial, desired *Instance, diffs []instanceDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Instance
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Instance); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Instance, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInstance(ctx, fetchState.urlNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Instance resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Instance resource: %v", err)
		}
		c.Config.Logger.Info("Found that Instance resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}

	c.Config.Logger.Infof("Found initial state for Instance: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Instance: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInstanceInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Instance: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInstanceDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Instance: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInstance(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInstanceInitialState(rawInitial, rawDesired *Instance) (*Instance, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInstanceDesiredState(rawDesired, rawInitial *Instance, opts ...dcl.ApplyOption) (*Instance, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.MaxDiskSize = canonicalizeInstanceMaxDiskSize(rawDesired.MaxDiskSize, nil, opts...)
		rawDesired.CurrentDiskSize = canonicalizeInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, nil, opts...)
		rawDesired.DiskEncryptionConfiguration = canonicalizeInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, nil, opts...)
		rawDesired.FailoverReplica = canonicalizeInstanceFailoverReplica(rawDesired.FailoverReplica, nil, opts...)
		rawDesired.MasterInstance = canonicalizeInstanceMasterInstance(rawDesired.MasterInstance, nil, opts...)
		rawDesired.ReplicaConfiguration = canonicalizeInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, nil, opts...)
		rawDesired.ScheduledMaintenance = canonicalizeInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, nil, opts...)
		rawDesired.Settings = canonicalizeInstanceSettings(rawDesired.Settings, nil, opts...)
		rawDesired.ServerCaCert = canonicalizeInstanceServerCaCert(rawDesired.ServerCaCert, nil, opts...)
		rawDesired.OnPremisesConfiguration = canonicalizeInstanceOnPremisesConfiguration(rawDesired.OnPremisesConfiguration, nil, opts...)
		rawDesired.DiskEncryptionStatus = canonicalizeInstanceDiskEncryptionStatus(rawDesired.DiskEncryptionStatus, nil, opts...)

		return rawDesired, nil
	}
	if dcl.IsZeroValue(rawDesired.BackendType) {
		rawDesired.BackendType = rawInitial.BackendType
	}
	if dcl.StringCanonicalize(rawDesired.ConnectionName, rawInitial.ConnectionName) {
		rawDesired.ConnectionName = rawInitial.ConnectionName
	}
	if dcl.IsZeroValue(rawDesired.DatabaseVersion) {
		rawDesired.DatabaseVersion = rawInitial.DatabaseVersion
	}
	if dcl.StringCanonicalize(rawDesired.Etag, rawInitial.Etag) {
		rawDesired.Etag = rawInitial.Etag
	}
	if dcl.StringCanonicalize(rawDesired.GceZone, rawInitial.GceZone) {
		rawDesired.GceZone = rawInitial.GceZone
	}
	if dcl.IsZeroValue(rawDesired.InstanceType) {
		rawDesired.InstanceType = rawInitial.InstanceType
	}
	if dcl.StringCanonicalize(rawDesired.MasterInstanceName, rawInitial.MasterInstanceName) {
		rawDesired.MasterInstanceName = rawInitial.MasterInstanceName
	}
	rawDesired.MaxDiskSize = canonicalizeInstanceMaxDiskSize(rawDesired.MaxDiskSize, rawInitial.MaxDiskSize, opts...)
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.StringCanonicalize(rawDesired.Region, rawInitial.Region) {
		rawDesired.Region = rawInitial.Region
	}
	if dcl.StringCanonicalize(rawDesired.RootPassword, rawInitial.RootPassword) {
		rawDesired.RootPassword = rawInitial.RootPassword
	}
	rawDesired.CurrentDiskSize = canonicalizeInstanceCurrentDiskSize(rawDesired.CurrentDiskSize, rawInitial.CurrentDiskSize, opts...)
	rawDesired.DiskEncryptionConfiguration = canonicalizeInstanceDiskEncryptionConfiguration(rawDesired.DiskEncryptionConfiguration, rawInitial.DiskEncryptionConfiguration, opts...)
	rawDesired.FailoverReplica = canonicalizeInstanceFailoverReplica(rawDesired.FailoverReplica, rawInitial.FailoverReplica, opts...)
	if dcl.IsZeroValue(rawDesired.IPAddresses) {
		rawDesired.IPAddresses = rawInitial.IPAddresses
	}
	rawDesired.MasterInstance = canonicalizeInstanceMasterInstance(rawDesired.MasterInstance, rawInitial.MasterInstance, opts...)
	rawDesired.ReplicaConfiguration = canonicalizeInstanceReplicaConfiguration(rawDesired.ReplicaConfiguration, rawInitial.ReplicaConfiguration, opts...)
	rawDesired.ScheduledMaintenance = canonicalizeInstanceScheduledMaintenance(rawDesired.ScheduledMaintenance, rawInitial.ScheduledMaintenance, opts...)
	rawDesired.Settings = canonicalizeInstanceSettings(rawDesired.Settings, rawInitial.Settings, opts...)
	if dcl.StringCanonicalize(rawDesired.State, rawInitial.State) {
		rawDesired.State = rawInitial.State
	}
	if dcl.IsZeroValue(rawDesired.ReplicaInstances) {
		rawDesired.ReplicaInstances = rawInitial.ReplicaInstances
	}
	rawDesired.ServerCaCert = canonicalizeInstanceServerCaCert(rawDesired.ServerCaCert, rawInitial.ServerCaCert, opts...)
	if dcl.StringCanonicalize(rawDesired.IPv6Address, rawInitial.IPv6Address) {
		rawDesired.IPv6Address = rawInitial.IPv6Address
	}
	if dcl.StringCanonicalize(rawDesired.ServiceAccountEmailAddress, rawInitial.ServiceAccountEmailAddress) {
		rawDesired.ServiceAccountEmailAddress = rawInitial.ServiceAccountEmailAddress
	}
	rawDesired.OnPremisesConfiguration = canonicalizeInstanceOnPremisesConfiguration(rawDesired.OnPremisesConfiguration, rawInitial.OnPremisesConfiguration, opts...)
	if dcl.IsZeroValue(rawDesired.SuspensionReason) {
		rawDesired.SuspensionReason = rawInitial.SuspensionReason
	}
	rawDesired.DiskEncryptionStatus = canonicalizeInstanceDiskEncryptionStatus(rawDesired.DiskEncryptionStatus, rawInitial.DiskEncryptionStatus, opts...)
	if dcl.StringCanonicalize(rawDesired.InstanceUid, rawInitial.InstanceUid) {
		rawDesired.InstanceUid = rawInitial.InstanceUid
	}

	return rawDesired, nil
}

func canonicalizeInstanceNewState(c *Client, rawNew, rawDesired *Instance) (*Instance, error) {

	if dcl.IsEmptyValueIndirect(rawNew.BackendType) && dcl.IsEmptyValueIndirect(rawDesired.BackendType) {
		rawNew.BackendType = rawDesired.BackendType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConnectionName) && dcl.IsEmptyValueIndirect(rawDesired.ConnectionName) {
		rawNew.ConnectionName = rawDesired.ConnectionName
	} else {
		if dcl.StringCanonicalize(rawDesired.ConnectionName, rawNew.ConnectionName) {
			rawNew.ConnectionName = rawDesired.ConnectionName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DatabaseVersion) && dcl.IsEmptyValueIndirect(rawDesired.DatabaseVersion) {
		rawNew.DatabaseVersion = rawDesired.DatabaseVersion
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GceZone) && dcl.IsEmptyValueIndirect(rawDesired.GceZone) {
		rawNew.GceZone = rawDesired.GceZone
	} else {
		if dcl.StringCanonicalize(rawDesired.GceZone, rawNew.GceZone) {
			rawNew.GceZone = rawDesired.GceZone
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceType) && dcl.IsEmptyValueIndirect(rawDesired.InstanceType) {
		rawNew.InstanceType = rawDesired.InstanceType
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterInstanceName) && dcl.IsEmptyValueIndirect(rawDesired.MasterInstanceName) {
		rawNew.MasterInstanceName = rawDesired.MasterInstanceName
	} else {
		if dcl.StringCanonicalize(rawDesired.MasterInstanceName, rawNew.MasterInstanceName) {
			rawNew.MasterInstanceName = rawDesired.MasterInstanceName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxDiskSize) && dcl.IsEmptyValueIndirect(rawDesired.MaxDiskSize) {
		rawNew.MaxDiskSize = rawDesired.MaxDiskSize
	} else {
		rawNew.MaxDiskSize = canonicalizeNewInstanceMaxDiskSize(c, rawDesired.MaxDiskSize, rawNew.MaxDiskSize)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Region) && dcl.IsEmptyValueIndirect(rawDesired.Region) {
		rawNew.Region = rawDesired.Region
	} else {
		if dcl.StringCanonicalize(rawDesired.Region, rawNew.Region) {
			rawNew.Region = rawDesired.Region
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.RootPassword) && dcl.IsEmptyValueIndirect(rawDesired.RootPassword) {
		rawNew.RootPassword = rawDesired.RootPassword
	} else {
		rawNew.RootPassword = rawDesired.RootPassword
	}

	if dcl.IsEmptyValueIndirect(rawNew.CurrentDiskSize) && dcl.IsEmptyValueIndirect(rawDesired.CurrentDiskSize) {
		rawNew.CurrentDiskSize = rawDesired.CurrentDiskSize
	} else {
		rawNew.CurrentDiskSize = canonicalizeNewInstanceCurrentDiskSize(c, rawDesired.CurrentDiskSize, rawNew.CurrentDiskSize)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskEncryptionConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.DiskEncryptionConfiguration) {
		rawNew.DiskEncryptionConfiguration = rawDesired.DiskEncryptionConfiguration
	} else {
		rawNew.DiskEncryptionConfiguration = canonicalizeNewInstanceDiskEncryptionConfiguration(c, rawDesired.DiskEncryptionConfiguration, rawNew.DiskEncryptionConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.FailoverReplica) && dcl.IsEmptyValueIndirect(rawDesired.FailoverReplica) {
		rawNew.FailoverReplica = rawDesired.FailoverReplica
	} else {
		rawNew.FailoverReplica = canonicalizeNewInstanceFailoverReplica(c, rawDesired.FailoverReplica, rawNew.FailoverReplica)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPAddresses) && dcl.IsEmptyValueIndirect(rawDesired.IPAddresses) {
		rawNew.IPAddresses = rawDesired.IPAddresses
	} else {
		rawNew.IPAddresses = canonicalizeNewInstanceIPAddressesSlice(c, rawDesired.IPAddresses, rawNew.IPAddresses)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MasterInstance) && dcl.IsEmptyValueIndirect(rawDesired.MasterInstance) {
		rawNew.MasterInstance = rawDesired.MasterInstance
	} else {
		rawNew.MasterInstance = canonicalizeNewInstanceMasterInstance(c, rawDesired.MasterInstance, rawNew.MasterInstance)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReplicaConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.ReplicaConfiguration) {
		rawNew.ReplicaConfiguration = rawDesired.ReplicaConfiguration
	} else {
		rawNew.ReplicaConfiguration = canonicalizeNewInstanceReplicaConfiguration(c, rawDesired.ReplicaConfiguration, rawNew.ReplicaConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ScheduledMaintenance) && dcl.IsEmptyValueIndirect(rawDesired.ScheduledMaintenance) {
		rawNew.ScheduledMaintenance = rawDesired.ScheduledMaintenance
	} else {
		rawNew.ScheduledMaintenance = canonicalizeNewInstanceScheduledMaintenance(c, rawDesired.ScheduledMaintenance, rawNew.ScheduledMaintenance)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Settings) && dcl.IsEmptyValueIndirect(rawDesired.Settings) {
		rawNew.Settings = rawDesired.Settings
	} else {
		rawNew.Settings = canonicalizeNewInstanceSettings(c, rawDesired.Settings, rawNew.Settings)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		if dcl.StringCanonicalize(rawDesired.State, rawNew.State) {
			rawNew.State = rawDesired.State
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ReplicaInstances) && dcl.IsEmptyValueIndirect(rawDesired.ReplicaInstances) {
		rawNew.ReplicaInstances = rawDesired.ReplicaInstances
	} else {
		rawNew.ReplicaInstances = canonicalizeNewInstanceReplicaInstancesSlice(c, rawDesired.ReplicaInstances, rawNew.ReplicaInstances)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServerCaCert) && dcl.IsEmptyValueIndirect(rawDesired.ServerCaCert) {
		rawNew.ServerCaCert = rawDesired.ServerCaCert
	} else {
		rawNew.ServerCaCert = canonicalizeNewInstanceServerCaCert(c, rawDesired.ServerCaCert, rawNew.ServerCaCert)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPv6Address) && dcl.IsEmptyValueIndirect(rawDesired.IPv6Address) {
		rawNew.IPv6Address = rawDesired.IPv6Address
	} else {
		if dcl.StringCanonicalize(rawDesired.IPv6Address, rawNew.IPv6Address) {
			rawNew.IPv6Address = rawDesired.IPv6Address
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ServiceAccountEmailAddress) && dcl.IsEmptyValueIndirect(rawDesired.ServiceAccountEmailAddress) {
		rawNew.ServiceAccountEmailAddress = rawDesired.ServiceAccountEmailAddress
	} else {
		if dcl.StringCanonicalize(rawDesired.ServiceAccountEmailAddress, rawNew.ServiceAccountEmailAddress) {
			rawNew.ServiceAccountEmailAddress = rawDesired.ServiceAccountEmailAddress
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.OnPremisesConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.OnPremisesConfiguration) {
		rawNew.OnPremisesConfiguration = rawDesired.OnPremisesConfiguration
	} else {
		rawNew.OnPremisesConfiguration = canonicalizeNewInstanceOnPremisesConfiguration(c, rawDesired.OnPremisesConfiguration, rawNew.OnPremisesConfiguration)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SuspensionReason) && dcl.IsEmptyValueIndirect(rawDesired.SuspensionReason) {
		rawNew.SuspensionReason = rawDesired.SuspensionReason
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DiskEncryptionStatus) && dcl.IsEmptyValueIndirect(rawDesired.DiskEncryptionStatus) {
		rawNew.DiskEncryptionStatus = rawDesired.DiskEncryptionStatus
	} else {
		rawNew.DiskEncryptionStatus = canonicalizeNewInstanceDiskEncryptionStatus(c, rawDesired.DiskEncryptionStatus, rawNew.DiskEncryptionStatus)
	}

	if dcl.IsEmptyValueIndirect(rawNew.InstanceUid) && dcl.IsEmptyValueIndirect(rawDesired.InstanceUid) {
		rawNew.InstanceUid = rawDesired.InstanceUid
	} else {
		if dcl.StringCanonicalize(rawDesired.InstanceUid, rawNew.InstanceUid) {
			rawNew.InstanceUid = rawDesired.InstanceUid
		}
	}

	return rawNew, nil
}

func canonicalizeInstanceMaxDiskSize(des, initial *InstanceMaxDiskSize, opts ...dcl.ApplyOption) *InstanceMaxDiskSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceMaxDiskSize(c *Client, des, nw *InstanceMaxDiskSize) *InstanceMaxDiskSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceMaxDiskSizeSet(c *Client, des, nw []InstanceMaxDiskSize) []InstanceMaxDiskSize {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMaxDiskSize
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMaxDiskSize(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceMaxDiskSizeSlice(c *Client, des, nw []InstanceMaxDiskSize) []InstanceMaxDiskSize {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceMaxDiskSize
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMaxDiskSize(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceCurrentDiskSize(des, initial *InstanceCurrentDiskSize, opts ...dcl.ApplyOption) *InstanceCurrentDiskSize {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceCurrentDiskSize(c *Client, des, nw *InstanceCurrentDiskSize) *InstanceCurrentDiskSize {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceCurrentDiskSizeSet(c *Client, des, nw []InstanceCurrentDiskSize) []InstanceCurrentDiskSize {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceCurrentDiskSize
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceCurrentDiskSize(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceCurrentDiskSizeSlice(c *Client, des, nw []InstanceCurrentDiskSize) []InstanceCurrentDiskSize {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceCurrentDiskSize
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceCurrentDiskSize(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceDiskEncryptionConfiguration(des, initial *InstanceDiskEncryptionConfiguration, opts ...dcl.ApplyOption) *InstanceDiskEncryptionConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.KmsKeyName, initial.KmsKeyName) || dcl.IsZeroValue(des.KmsKeyName) {
		des.KmsKeyName = initial.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceDiskEncryptionConfiguration(c *Client, des, nw *InstanceDiskEncryptionConfiguration) *InstanceDiskEncryptionConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyName, nw.KmsKeyName) {
		nw.KmsKeyName = des.KmsKeyName
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceDiskEncryptionConfigurationSet(c *Client, des, nw []InstanceDiskEncryptionConfiguration) []InstanceDiskEncryptionConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDiskEncryptionConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDiskEncryptionConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceDiskEncryptionConfigurationSlice(c *Client, des, nw []InstanceDiskEncryptionConfiguration) []InstanceDiskEncryptionConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceDiskEncryptionConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceDiskEncryptionConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceFailoverReplica(des, initial *InstanceFailoverReplica, opts ...dcl.ApplyOption) *InstanceFailoverReplica {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.BoolCanonicalize(des.Available, initial.Available) || dcl.IsZeroValue(des.Available) {
		des.Available = initial.Available
	}
	des.FailoverInstance = canonicalizeInstanceFailoverReplicaFailoverInstance(des.FailoverInstance, initial.FailoverInstance, opts...)

	return des
}

func canonicalizeNewInstanceFailoverReplica(c *Client, des, nw *InstanceFailoverReplica) *InstanceFailoverReplica {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.BoolCanonicalize(des.Available, nw.Available) {
		nw.Available = des.Available
	}
	nw.FailoverInstance = canonicalizeNewInstanceFailoverReplicaFailoverInstance(c, des.FailoverInstance, nw.FailoverInstance)

	return nw
}

func canonicalizeNewInstanceFailoverReplicaSet(c *Client, des, nw []InstanceFailoverReplica) []InstanceFailoverReplica {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceFailoverReplica
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceFailoverReplica(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceFailoverReplicaSlice(c *Client, des, nw []InstanceFailoverReplica) []InstanceFailoverReplica {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceFailoverReplica
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceFailoverReplica(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceFailoverReplicaFailoverInstance(des, initial *InstanceFailoverReplicaFailoverInstance, opts ...dcl.ApplyOption) *InstanceFailoverReplicaFailoverInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Region, initial.Region) || dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewInstanceFailoverReplicaFailoverInstance(c *Client, des, nw *InstanceFailoverReplicaFailoverInstance) *InstanceFailoverReplicaFailoverInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Region, nw.Region) {
		nw.Region = des.Region
	}

	return nw
}

func canonicalizeNewInstanceFailoverReplicaFailoverInstanceSet(c *Client, des, nw []InstanceFailoverReplicaFailoverInstance) []InstanceFailoverReplicaFailoverInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceFailoverReplicaFailoverInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceFailoverReplicaFailoverInstance(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceFailoverReplicaFailoverInstanceSlice(c *Client, des, nw []InstanceFailoverReplicaFailoverInstance) []InstanceFailoverReplicaFailoverInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceFailoverReplicaFailoverInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceFailoverReplicaFailoverInstance(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceIPAddresses(des, initial *InstanceIPAddresses, opts ...dcl.ApplyOption) *InstanceIPAddresses {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}
	if dcl.StringCanonicalize(des.IPAddress, initial.IPAddress) || dcl.IsZeroValue(des.IPAddress) {
		des.IPAddress = initial.IPAddress
	}
	des.TimeToRetire = canonicalizeInstanceIPAddressesTimeToRetire(des.TimeToRetire, initial.TimeToRetire, opts...)

	return des
}

func canonicalizeNewInstanceIPAddresses(c *Client, des, nw *InstanceIPAddresses) *InstanceIPAddresses {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.IPAddress, nw.IPAddress) {
		nw.IPAddress = des.IPAddress
	}
	nw.TimeToRetire = canonicalizeNewInstanceIPAddressesTimeToRetire(c, des.TimeToRetire, nw.TimeToRetire)

	return nw
}

func canonicalizeNewInstanceIPAddressesSet(c *Client, des, nw []InstanceIPAddresses) []InstanceIPAddresses {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceIPAddresses
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceIPAddresses(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceIPAddressesSlice(c *Client, des, nw []InstanceIPAddresses) []InstanceIPAddresses {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceIPAddresses
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceIPAddresses(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceIPAddressesTimeToRetire(des, initial *InstanceIPAddressesTimeToRetire, opts ...dcl.ApplyOption) *InstanceIPAddressesTimeToRetire {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewInstanceIPAddressesTimeToRetire(c *Client, des, nw *InstanceIPAddressesTimeToRetire) *InstanceIPAddressesTimeToRetire {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceIPAddressesTimeToRetireSet(c *Client, des, nw []InstanceIPAddressesTimeToRetire) []InstanceIPAddressesTimeToRetire {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceIPAddressesTimeToRetire
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceIPAddressesTimeToRetire(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceIPAddressesTimeToRetireSlice(c *Client, des, nw []InstanceIPAddressesTimeToRetire) []InstanceIPAddressesTimeToRetire {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceIPAddressesTimeToRetire
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceIPAddressesTimeToRetire(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceMasterInstance(des, initial *InstanceMasterInstance, opts ...dcl.ApplyOption) *InstanceMasterInstance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Region, initial.Region) || dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewInstanceMasterInstance(c *Client, des, nw *InstanceMasterInstance) *InstanceMasterInstance {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Region, nw.Region) {
		nw.Region = des.Region
	}

	return nw
}

func canonicalizeNewInstanceMasterInstanceSet(c *Client, des, nw []InstanceMasterInstance) []InstanceMasterInstance {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceMasterInstance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceMasterInstance(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceMasterInstanceSlice(c *Client, des, nw []InstanceMasterInstance) []InstanceMasterInstance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceMasterInstance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceMasterInstance(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaConfiguration(des, initial *InstanceReplicaConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	des.MysqlReplicaConfiguration = canonicalizeInstanceReplicaConfigurationMysqlReplicaConfiguration(des.MysqlReplicaConfiguration, initial.MysqlReplicaConfiguration, opts...)
	if dcl.BoolCanonicalize(des.FailoverTarget, initial.FailoverTarget) || dcl.IsZeroValue(des.FailoverTarget) {
		des.FailoverTarget = initial.FailoverTarget
	}
	des.ReplicaPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfiguration(des.ReplicaPoolConfiguration, initial.ReplicaPoolConfiguration, opts...)

	return des
}

func canonicalizeNewInstanceReplicaConfiguration(c *Client, des, nw *InstanceReplicaConfiguration) *InstanceReplicaConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	nw.MysqlReplicaConfiguration = canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(c, des.MysqlReplicaConfiguration, nw.MysqlReplicaConfiguration)
	if dcl.BoolCanonicalize(des.FailoverTarget, nw.FailoverTarget) {
		nw.FailoverTarget = des.FailoverTarget
	}
	nw.ReplicaPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(c, des.ReplicaPoolConfiguration, nw.ReplicaPoolConfiguration)

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationSet(c *Client, des, nw []InstanceReplicaConfiguration) []InstanceReplicaConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaConfigurationSlice(c *Client, des, nw []InstanceReplicaConfiguration) []InstanceReplicaConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaConfigurationMysqlReplicaConfiguration(des, initial *InstanceReplicaConfigurationMysqlReplicaConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.DumpFilePath, initial.DumpFilePath) || dcl.IsZeroValue(des.DumpFilePath) {
		des.DumpFilePath = initial.DumpFilePath
	}
	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.StringCanonicalize(des.Password, initial.Password) || dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}
	if dcl.IsZeroValue(des.ConnectRetryInterval) {
		des.ConnectRetryInterval = initial.ConnectRetryInterval
	}
	des.MasterHeartbeatPeriod = canonicalizeInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des.MasterHeartbeatPeriod, initial.MasterHeartbeatPeriod, opts...)
	if dcl.StringCanonicalize(des.CaCertificate, initial.CaCertificate) || dcl.IsZeroValue(des.CaCertificate) {
		des.CaCertificate = initial.CaCertificate
	}
	if dcl.StringCanonicalize(des.ClientCertificate, initial.ClientCertificate) || dcl.IsZeroValue(des.ClientCertificate) {
		des.ClientCertificate = initial.ClientCertificate
	}
	if dcl.StringCanonicalize(des.ClientKey, initial.ClientKey) || dcl.IsZeroValue(des.ClientKey) {
		des.ClientKey = initial.ClientKey
	}
	if dcl.StringCanonicalize(des.SslCipher, initial.SslCipher) || dcl.IsZeroValue(des.SslCipher) {
		des.SslCipher = initial.SslCipher
	}
	if dcl.BoolCanonicalize(des.VerifyServerCertificate, initial.VerifyServerCertificate) || dcl.IsZeroValue(des.VerifyServerCertificate) {
		des.VerifyServerCertificate = initial.VerifyServerCertificate
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, des, nw *InstanceReplicaConfigurationMysqlReplicaConfiguration) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.DumpFilePath, nw.DumpFilePath) {
		nw.DumpFilePath = des.DumpFilePath
	}
	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}
	if dcl.StringCanonicalize(des.Password, nw.Password) {
		nw.Password = des.Password
	}
	nw.MasterHeartbeatPeriod = canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, des.MasterHeartbeatPeriod, nw.MasterHeartbeatPeriod)
	if dcl.StringCanonicalize(des.CaCertificate, nw.CaCertificate) {
		nw.CaCertificate = des.CaCertificate
	}
	if dcl.StringCanonicalize(des.ClientCertificate, nw.ClientCertificate) {
		nw.ClientCertificate = des.ClientCertificate
	}
	if dcl.StringCanonicalize(des.ClientKey, nw.ClientKey) {
		nw.ClientKey = des.ClientKey
	}
	if dcl.StringCanonicalize(des.SslCipher, nw.SslCipher) {
		nw.SslCipher = des.SslCipher
	}
	if dcl.BoolCanonicalize(des.VerifyServerCertificate, nw.VerifyServerCertificate) {
		nw.VerifyServerCertificate = des.VerifyServerCertificate
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationMysqlReplicaConfiguration) []InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationMysqlReplicaConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, des, nw []InstanceReplicaConfigurationMysqlReplicaConfiguration) []InstanceReplicaConfigurationMysqlReplicaConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaConfigurationMysqlReplicaConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(des, initial *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, des, nw *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSet(c *Client, des, nw []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, des, nw []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	des.StaticPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des.StaticPoolConfiguration, initial.StaticPoolConfiguration, opts...)
	des.AutoscalingPoolConfiguration = canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des.AutoscalingPoolConfiguration, initial.AutoscalingPoolConfiguration, opts...)
	if dcl.IsZeroValue(des.ReplicaCount) {
		des.ReplicaCount = initial.ReplicaCount
	}
	if dcl.BoolCanonicalize(des.ExposeReplicaIP, initial.ExposeReplicaIP) || dcl.IsZeroValue(des.ExposeReplicaIP) {
		des.ExposeReplicaIP = initial.ExposeReplicaIP
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, des, nw *InstanceReplicaConfigurationReplicaPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	nw.StaticPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, des.StaticPoolConfiguration, nw.StaticPoolConfiguration)
	nw.AutoscalingPoolConfiguration = canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, des.AutoscalingPoolConfiguration, nw.AutoscalingPoolConfiguration)
	if dcl.BoolCanonicalize(des.ExposeReplicaIP, nw.ExposeReplicaIP) {
		nw.ExposeReplicaIP = des.ExposeReplicaIP
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationReplicaPoolConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaConfigurationReplicaPoolConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.ReplicaCount) {
		des.ReplicaCount = initial.ReplicaCount
	}
	if dcl.BoolCanonicalize(des.ExposeReplicaIP, initial.ExposeReplicaIP) || dcl.IsZeroValue(des.ExposeReplicaIP) {
		des.ExposeReplicaIP = initial.ExposeReplicaIP
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, des, nw *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.BoolCanonicalize(des.ExposeReplicaIP, nw.ExposeReplicaIP) {
		nw.ExposeReplicaIP = des.ExposeReplicaIP
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(des, initial *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, opts ...dcl.ApplyOption) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.MinReplicaCount) {
		des.MinReplicaCount = initial.MinReplicaCount
	}
	if dcl.IsZeroValue(des.MaxReplicaCount) {
		des.MaxReplicaCount = initial.MaxReplicaCount
	}
	if dcl.IsZeroValue(des.TargetCpuUtil) {
		des.TargetCpuUtil = initial.TargetCpuUtil
	}

	return des
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, des, nw *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSet(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, des, nw []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceScheduledMaintenance(des, initial *InstanceScheduledMaintenance, opts ...dcl.ApplyOption) *InstanceScheduledMaintenance {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.StartTime = canonicalizeInstanceScheduledMaintenanceStartTime(des.StartTime, initial.StartTime, opts...)
	if dcl.BoolCanonicalize(des.CanDefer, initial.CanDefer) || dcl.IsZeroValue(des.CanDefer) {
		des.CanDefer = initial.CanDefer
	}
	if dcl.BoolCanonicalize(des.CanReschedule, initial.CanReschedule) || dcl.IsZeroValue(des.CanReschedule) {
		des.CanReschedule = initial.CanReschedule
	}

	return des
}

func canonicalizeNewInstanceScheduledMaintenance(c *Client, des, nw *InstanceScheduledMaintenance) *InstanceScheduledMaintenance {
	if des == nil || nw == nil {
		return nw
	}

	nw.StartTime = canonicalizeNewInstanceScheduledMaintenanceStartTime(c, des.StartTime, nw.StartTime)
	if dcl.BoolCanonicalize(des.CanDefer, nw.CanDefer) {
		nw.CanDefer = des.CanDefer
	}
	if dcl.BoolCanonicalize(des.CanReschedule, nw.CanReschedule) {
		nw.CanReschedule = des.CanReschedule
	}

	return nw
}

func canonicalizeNewInstanceScheduledMaintenanceSet(c *Client, des, nw []InstanceScheduledMaintenance) []InstanceScheduledMaintenance {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceScheduledMaintenance
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceScheduledMaintenance(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceScheduledMaintenanceSlice(c *Client, des, nw []InstanceScheduledMaintenance) []InstanceScheduledMaintenance {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceScheduledMaintenance
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceScheduledMaintenance(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceScheduledMaintenanceStartTime(des, initial *InstanceScheduledMaintenanceStartTime, opts ...dcl.ApplyOption) *InstanceScheduledMaintenanceStartTime {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Seconds) {
		des.Seconds = initial.Seconds
	}
	if dcl.IsZeroValue(des.Nanos) {
		des.Nanos = initial.Nanos
	}

	return des
}

func canonicalizeNewInstanceScheduledMaintenanceStartTime(c *Client, des, nw *InstanceScheduledMaintenanceStartTime) *InstanceScheduledMaintenanceStartTime {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceScheduledMaintenanceStartTimeSet(c *Client, des, nw []InstanceScheduledMaintenanceStartTime) []InstanceScheduledMaintenanceStartTime {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceScheduledMaintenanceStartTime
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceScheduledMaintenanceStartTime(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceScheduledMaintenanceStartTimeSlice(c *Client, des, nw []InstanceScheduledMaintenanceStartTime) []InstanceScheduledMaintenanceStartTime {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceScheduledMaintenanceStartTime
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceScheduledMaintenanceStartTime(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettings(des, initial *InstanceSettings, opts ...dcl.ApplyOption) *InstanceSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AuthorizedGaeApplications) {
		des.AuthorizedGaeApplications = initial.AuthorizedGaeApplications
	}
	if dcl.StringCanonicalize(des.Tier, initial.Tier) || dcl.IsZeroValue(des.Tier) {
		des.Tier = initial.Tier
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.IsZeroValue(des.AvailabilityType) {
		des.AvailabilityType = initial.AvailabilityType
	}
	if dcl.IsZeroValue(des.PricingPlan) {
		des.PricingPlan = initial.PricingPlan
	}
	if dcl.IsZeroValue(des.ReplicationType) {
		des.ReplicationType = initial.ReplicationType
	}
	if dcl.IsZeroValue(des.ActivationPolicy) {
		des.ActivationPolicy = initial.ActivationPolicy
	}
	if dcl.BoolCanonicalize(des.StorageAutoResize, initial.StorageAutoResize) || dcl.IsZeroValue(des.StorageAutoResize) {
		des.StorageAutoResize = initial.StorageAutoResize
	}
	if dcl.IsZeroValue(des.DataDiskType) {
		des.DataDiskType = initial.DataDiskType
	}
	if dcl.BoolCanonicalize(des.DatabaseReplicationEnabled, initial.DatabaseReplicationEnabled) || dcl.IsZeroValue(des.DatabaseReplicationEnabled) {
		des.DatabaseReplicationEnabled = initial.DatabaseReplicationEnabled
	}
	if dcl.BoolCanonicalize(des.CrashSafeReplicationEnabled, initial.CrashSafeReplicationEnabled) || dcl.IsZeroValue(des.CrashSafeReplicationEnabled) {
		des.CrashSafeReplicationEnabled = initial.CrashSafeReplicationEnabled
	}
	des.SettingsVersion = canonicalizeInstanceSettingsSettingsVersion(des.SettingsVersion, initial.SettingsVersion, opts...)
	if dcl.IsZeroValue(des.UserLabels) {
		des.UserLabels = initial.UserLabels
	}
	des.StorageAutoResizeLimit = canonicalizeInstanceSettingsStorageAutoResizeLimit(des.StorageAutoResizeLimit, initial.StorageAutoResizeLimit, opts...)
	des.IPConfiguration = canonicalizeInstanceSettingsIPConfiguration(des.IPConfiguration, initial.IPConfiguration, opts...)
	des.LocationPreference = canonicalizeInstanceSettingsLocationPreference(des.LocationPreference, initial.LocationPreference, opts...)
	if dcl.IsZeroValue(des.DatabaseFlags) {
		des.DatabaseFlags = initial.DatabaseFlags
	}
	des.MaintenanceWindow = canonicalizeInstanceSettingsMaintenanceWindow(des.MaintenanceWindow, initial.MaintenanceWindow, opts...)
	des.BackupConfiguration = canonicalizeInstanceSettingsBackupConfiguration(des.BackupConfiguration, initial.BackupConfiguration, opts...)
	des.DataDiskSizeGb = canonicalizeInstanceSettingsDataDiskSizeGb(des.DataDiskSizeGb, initial.DataDiskSizeGb, opts...)
	des.ActiveDirectoryConfig = canonicalizeInstanceSettingsActiveDirectoryConfig(des.ActiveDirectoryConfig, initial.ActiveDirectoryConfig, opts...)
	if dcl.StringCanonicalize(des.Collation, initial.Collation) || dcl.IsZeroValue(des.Collation) {
		des.Collation = initial.Collation
	}
	if dcl.IsZeroValue(des.DenyMaintenancePeriods) {
		des.DenyMaintenancePeriods = initial.DenyMaintenancePeriods
	}
	des.InsightsConfig = canonicalizeInstanceSettingsInsightsConfig(des.InsightsConfig, initial.InsightsConfig, opts...)

	return des
}

func canonicalizeNewInstanceSettings(c *Client, des, nw *InstanceSettings) *InstanceSettings {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Tier, nw.Tier) {
		nw.Tier = des.Tier
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.BoolCanonicalize(des.StorageAutoResize, nw.StorageAutoResize) {
		nw.StorageAutoResize = des.StorageAutoResize
	}
	if dcl.BoolCanonicalize(des.DatabaseReplicationEnabled, nw.DatabaseReplicationEnabled) {
		nw.DatabaseReplicationEnabled = des.DatabaseReplicationEnabled
	}
	if dcl.BoolCanonicalize(des.CrashSafeReplicationEnabled, nw.CrashSafeReplicationEnabled) {
		nw.CrashSafeReplicationEnabled = des.CrashSafeReplicationEnabled
	}
	nw.SettingsVersion = canonicalizeNewInstanceSettingsSettingsVersion(c, des.SettingsVersion, nw.SettingsVersion)
	nw.StorageAutoResizeLimit = canonicalizeNewInstanceSettingsStorageAutoResizeLimit(c, des.StorageAutoResizeLimit, nw.StorageAutoResizeLimit)
	nw.IPConfiguration = canonicalizeNewInstanceSettingsIPConfiguration(c, des.IPConfiguration, nw.IPConfiguration)
	nw.LocationPreference = canonicalizeNewInstanceSettingsLocationPreference(c, des.LocationPreference, nw.LocationPreference)
	nw.DatabaseFlags = canonicalizeNewInstanceSettingsDatabaseFlagsSlice(c, des.DatabaseFlags, nw.DatabaseFlags)
	nw.MaintenanceWindow = canonicalizeNewInstanceSettingsMaintenanceWindow(c, des.MaintenanceWindow, nw.MaintenanceWindow)
	nw.BackupConfiguration = canonicalizeNewInstanceSettingsBackupConfiguration(c, des.BackupConfiguration, nw.BackupConfiguration)
	nw.DataDiskSizeGb = canonicalizeNewInstanceSettingsDataDiskSizeGb(c, des.DataDiskSizeGb, nw.DataDiskSizeGb)
	nw.ActiveDirectoryConfig = canonicalizeNewInstanceSettingsActiveDirectoryConfig(c, des.ActiveDirectoryConfig, nw.ActiveDirectoryConfig)
	if dcl.StringCanonicalize(des.Collation, nw.Collation) {
		nw.Collation = des.Collation
	}
	nw.DenyMaintenancePeriods = canonicalizeNewInstanceSettingsDenyMaintenancePeriodsSlice(c, des.DenyMaintenancePeriods, nw.DenyMaintenancePeriods)
	nw.InsightsConfig = canonicalizeNewInstanceSettingsInsightsConfig(c, des.InsightsConfig, nw.InsightsConfig)

	return nw
}

func canonicalizeNewInstanceSettingsSet(c *Client, des, nw []InstanceSettings) []InstanceSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettings(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsSlice(c *Client, des, nw []InstanceSettings) []InstanceSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettings(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsSettingsVersion(des, initial *InstanceSettingsSettingsVersion, opts ...dcl.ApplyOption) *InstanceSettingsSettingsVersion {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsSettingsVersion(c *Client, des, nw *InstanceSettingsSettingsVersion) *InstanceSettingsSettingsVersion {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsSettingsVersionSet(c *Client, des, nw []InstanceSettingsSettingsVersion) []InstanceSettingsSettingsVersion {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsSettingsVersion
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsSettingsVersion(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsSettingsVersionSlice(c *Client, des, nw []InstanceSettingsSettingsVersion) []InstanceSettingsSettingsVersion {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsSettingsVersion
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsSettingsVersion(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsStorageAutoResizeLimit(des, initial *InstanceSettingsStorageAutoResizeLimit, opts ...dcl.ApplyOption) *InstanceSettingsStorageAutoResizeLimit {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsStorageAutoResizeLimit(c *Client, des, nw *InstanceSettingsStorageAutoResizeLimit) *InstanceSettingsStorageAutoResizeLimit {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsStorageAutoResizeLimitSet(c *Client, des, nw []InstanceSettingsStorageAutoResizeLimit) []InstanceSettingsStorageAutoResizeLimit {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsStorageAutoResizeLimit
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsStorageAutoResizeLimit(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsStorageAutoResizeLimitSlice(c *Client, des, nw []InstanceSettingsStorageAutoResizeLimit) []InstanceSettingsStorageAutoResizeLimit {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsStorageAutoResizeLimit
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsStorageAutoResizeLimit(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsIPConfiguration(des, initial *InstanceSettingsIPConfiguration, opts ...dcl.ApplyOption) *InstanceSettingsIPConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IPv4Enabled, initial.IPv4Enabled) || dcl.IsZeroValue(des.IPv4Enabled) {
		des.IPv4Enabled = initial.IPv4Enabled
	}
	if dcl.NameToSelfLink(des.PrivateNetwork, initial.PrivateNetwork) || dcl.IsZeroValue(des.PrivateNetwork) {
		des.PrivateNetwork = initial.PrivateNetwork
	}
	if dcl.BoolCanonicalize(des.RequireSsl, initial.RequireSsl) || dcl.IsZeroValue(des.RequireSsl) {
		des.RequireSsl = initial.RequireSsl
	}
	if dcl.IsZeroValue(des.AuthorizedNetworks) {
		des.AuthorizedNetworks = initial.AuthorizedNetworks
	}

	return des
}

func canonicalizeNewInstanceSettingsIPConfiguration(c *Client, des, nw *InstanceSettingsIPConfiguration) *InstanceSettingsIPConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IPv4Enabled, nw.IPv4Enabled) {
		nw.IPv4Enabled = des.IPv4Enabled
	}
	if dcl.NameToSelfLink(des.PrivateNetwork, nw.PrivateNetwork) {
		nw.PrivateNetwork = des.PrivateNetwork
	}
	if dcl.BoolCanonicalize(des.RequireSsl, nw.RequireSsl) {
		nw.RequireSsl = des.RequireSsl
	}
	nw.AuthorizedNetworks = canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, des.AuthorizedNetworks, nw.AuthorizedNetworks)

	return nw
}

func canonicalizeNewInstanceSettingsIPConfigurationSet(c *Client, des, nw []InstanceSettingsIPConfiguration) []InstanceSettingsIPConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsIPConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsIPConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsIPConfigurationSlice(c *Client, des, nw []InstanceSettingsIPConfiguration) []InstanceSettingsIPConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsIPConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsIPConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsIPConfigurationAuthorizedNetworks(des, initial *InstanceSettingsIPConfigurationAuthorizedNetworks, opts ...dcl.ApplyOption) *InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}
	if dcl.IsZeroValue(des.ExpirationTime) {
		des.ExpirationTime = initial.ExpirationTime
	}
	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, des, nw *InstanceSettingsIPConfigurationAuthorizedNetworks) *InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}
	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworksSet(c *Client, des, nw []InstanceSettingsIPConfigurationAuthorizedNetworks) []InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsIPConfigurationAuthorizedNetworks
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsIPConfigurationAuthorizedNetworks(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, des, nw []InstanceSettingsIPConfigurationAuthorizedNetworks) []InstanceSettingsIPConfigurationAuthorizedNetworks {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsIPConfigurationAuthorizedNetworks
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsIPConfigurationAuthorizedNetworks(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsLocationPreference(des, initial *InstanceSettingsLocationPreference, opts ...dcl.ApplyOption) *InstanceSettingsLocationPreference {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Zone, initial.Zone) || dcl.IsZeroValue(des.Zone) {
		des.Zone = initial.Zone
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceSettingsLocationPreference(c *Client, des, nw *InstanceSettingsLocationPreference) *InstanceSettingsLocationPreference {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Zone, nw.Zone) {
		nw.Zone = des.Zone
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceSettingsLocationPreferenceSet(c *Client, des, nw []InstanceSettingsLocationPreference) []InstanceSettingsLocationPreference {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsLocationPreference
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsLocationPreference(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsLocationPreferenceSlice(c *Client, des, nw []InstanceSettingsLocationPreference) []InstanceSettingsLocationPreference {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsLocationPreference
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsLocationPreference(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsDatabaseFlags(des, initial *InstanceSettingsDatabaseFlags, opts ...dcl.ApplyOption) *InstanceSettingsDatabaseFlags {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsDatabaseFlags(c *Client, des, nw *InstanceSettingsDatabaseFlags) *InstanceSettingsDatabaseFlags {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewInstanceSettingsDatabaseFlagsSet(c *Client, des, nw []InstanceSettingsDatabaseFlags) []InstanceSettingsDatabaseFlags {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsDatabaseFlags
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsDatabaseFlags(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsDatabaseFlagsSlice(c *Client, des, nw []InstanceSettingsDatabaseFlags) []InstanceSettingsDatabaseFlags {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsDatabaseFlags
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsDatabaseFlags(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsMaintenanceWindow(des, initial *InstanceSettingsMaintenanceWindow, opts ...dcl.ApplyOption) *InstanceSettingsMaintenanceWindow {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Hour) {
		des.Hour = initial.Hour
	}
	if dcl.IsZeroValue(des.Day) {
		des.Day = initial.Day
	}
	if dcl.IsZeroValue(des.UpdateTrack) {
		des.UpdateTrack = initial.UpdateTrack
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceSettingsMaintenanceWindow(c *Client, des, nw *InstanceSettingsMaintenanceWindow) *InstanceSettingsMaintenanceWindow {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceSettingsMaintenanceWindowSet(c *Client, des, nw []InstanceSettingsMaintenanceWindow) []InstanceSettingsMaintenanceWindow {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsMaintenanceWindow
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsMaintenanceWindow(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsMaintenanceWindowSlice(c *Client, des, nw []InstanceSettingsMaintenanceWindow) []InstanceSettingsMaintenanceWindow {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsMaintenanceWindow
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsMaintenanceWindow(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsBackupConfiguration(des, initial *InstanceSettingsBackupConfiguration, opts ...dcl.ApplyOption) *InstanceSettingsBackupConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.StartTime, initial.StartTime) || dcl.IsZeroValue(des.StartTime) {
		des.StartTime = initial.StartTime
	}
	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		des.Enabled = initial.Enabled
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.BoolCanonicalize(des.BinaryLogEnabled, initial.BinaryLogEnabled) || dcl.IsZeroValue(des.BinaryLogEnabled) {
		des.BinaryLogEnabled = initial.BinaryLogEnabled
	}
	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		des.Location = initial.Location
	}
	des.BackupRetentionSettings = canonicalizeInstanceSettingsBackupConfigurationBackupRetentionSettings(des.BackupRetentionSettings, initial.BackupRetentionSettings, opts...)
	if dcl.IsZeroValue(des.TransactionLogRetentionDays) {
		des.TransactionLogRetentionDays = initial.TransactionLogRetentionDays
	}

	return des
}

func canonicalizeNewInstanceSettingsBackupConfiguration(c *Client, des, nw *InstanceSettingsBackupConfiguration) *InstanceSettingsBackupConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StartTime, nw.StartTime) {
		nw.StartTime = des.StartTime
	}
	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.BoolCanonicalize(des.BinaryLogEnabled, nw.BinaryLogEnabled) {
		nw.BinaryLogEnabled = des.BinaryLogEnabled
	}
	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}
	nw.BackupRetentionSettings = canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettings(c, des.BackupRetentionSettings, nw.BackupRetentionSettings)

	return nw
}

func canonicalizeNewInstanceSettingsBackupConfigurationSet(c *Client, des, nw []InstanceSettingsBackupConfiguration) []InstanceSettingsBackupConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsBackupConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsBackupConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsBackupConfigurationSlice(c *Client, des, nw []InstanceSettingsBackupConfiguration) []InstanceSettingsBackupConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsBackupConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsBackupConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsBackupConfigurationBackupRetentionSettings(des, initial *InstanceSettingsBackupConfigurationBackupRetentionSettings, opts ...dcl.ApplyOption) *InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.RetentionUnit) {
		des.RetentionUnit = initial.RetentionUnit
	}
	if dcl.IsZeroValue(des.RetainedBackups) {
		des.RetainedBackups = initial.RetainedBackups
	}

	return des
}

func canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, des, nw *InstanceSettingsBackupConfigurationBackupRetentionSettings) *InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettingsSet(c *Client, des, nw []InstanceSettingsBackupConfigurationBackupRetentionSettings) []InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsBackupConfigurationBackupRetentionSettings
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, des, nw []InstanceSettingsBackupConfigurationBackupRetentionSettings) []InstanceSettingsBackupConfigurationBackupRetentionSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsBackupConfigurationBackupRetentionSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsDataDiskSizeGb(des, initial *InstanceSettingsDataDiskSizeGb, opts ...dcl.ApplyOption) *InstanceSettingsDataDiskSizeGb {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewInstanceSettingsDataDiskSizeGb(c *Client, des, nw *InstanceSettingsDataDiskSizeGb) *InstanceSettingsDataDiskSizeGb {
	if des == nil || nw == nil {
		return nw
	}

	return nw
}

func canonicalizeNewInstanceSettingsDataDiskSizeGbSet(c *Client, des, nw []InstanceSettingsDataDiskSizeGb) []InstanceSettingsDataDiskSizeGb {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsDataDiskSizeGb
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsDataDiskSizeGb(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsDataDiskSizeGbSlice(c *Client, des, nw []InstanceSettingsDataDiskSizeGb) []InstanceSettingsDataDiskSizeGb {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsDataDiskSizeGb
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsDataDiskSizeGb(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsActiveDirectoryConfig(des, initial *InstanceSettingsActiveDirectoryConfig, opts ...dcl.ApplyOption) *InstanceSettingsActiveDirectoryConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.Domain, initial.Domain) || dcl.IsZeroValue(des.Domain) {
		des.Domain = initial.Domain
	}

	return des
}

func canonicalizeNewInstanceSettingsActiveDirectoryConfig(c *Client, des, nw *InstanceSettingsActiveDirectoryConfig) *InstanceSettingsActiveDirectoryConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.Domain, nw.Domain) {
		nw.Domain = des.Domain
	}

	return nw
}

func canonicalizeNewInstanceSettingsActiveDirectoryConfigSet(c *Client, des, nw []InstanceSettingsActiveDirectoryConfig) []InstanceSettingsActiveDirectoryConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsActiveDirectoryConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsActiveDirectoryConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsActiveDirectoryConfigSlice(c *Client, des, nw []InstanceSettingsActiveDirectoryConfig) []InstanceSettingsActiveDirectoryConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsActiveDirectoryConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsActiveDirectoryConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsDenyMaintenancePeriods(des, initial *InstanceSettingsDenyMaintenancePeriods, opts ...dcl.ApplyOption) *InstanceSettingsDenyMaintenancePeriods {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.StartDate, initial.StartDate) || dcl.IsZeroValue(des.StartDate) {
		des.StartDate = initial.StartDate
	}
	if dcl.StringCanonicalize(des.EndDate, initial.EndDate) || dcl.IsZeroValue(des.EndDate) {
		des.EndDate = initial.EndDate
	}
	if dcl.StringCanonicalize(des.Time, initial.Time) || dcl.IsZeroValue(des.Time) {
		des.Time = initial.Time
	}

	return des
}

func canonicalizeNewInstanceSettingsDenyMaintenancePeriods(c *Client, des, nw *InstanceSettingsDenyMaintenancePeriods) *InstanceSettingsDenyMaintenancePeriods {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.StartDate, nw.StartDate) {
		nw.StartDate = des.StartDate
	}
	if dcl.StringCanonicalize(des.EndDate, nw.EndDate) {
		nw.EndDate = des.EndDate
	}
	if dcl.StringCanonicalize(des.Time, nw.Time) {
		nw.Time = des.Time
	}

	return nw
}

func canonicalizeNewInstanceSettingsDenyMaintenancePeriodsSet(c *Client, des, nw []InstanceSettingsDenyMaintenancePeriods) []InstanceSettingsDenyMaintenancePeriods {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsDenyMaintenancePeriods
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsDenyMaintenancePeriods(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, des, nw []InstanceSettingsDenyMaintenancePeriods) []InstanceSettingsDenyMaintenancePeriods {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsDenyMaintenancePeriods
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsDenyMaintenancePeriods(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceSettingsInsightsConfig(des, initial *InstanceSettingsInsightsConfig, opts ...dcl.ApplyOption) *InstanceSettingsInsightsConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.QueryInsightsEnabled, initial.QueryInsightsEnabled) || dcl.IsZeroValue(des.QueryInsightsEnabled) {
		des.QueryInsightsEnabled = initial.QueryInsightsEnabled
	}
	if dcl.BoolCanonicalize(des.RecordClientAddress, initial.RecordClientAddress) || dcl.IsZeroValue(des.RecordClientAddress) {
		des.RecordClientAddress = initial.RecordClientAddress
	}
	if dcl.BoolCanonicalize(des.RecordApplicationTags, initial.RecordApplicationTags) || dcl.IsZeroValue(des.RecordApplicationTags) {
		des.RecordApplicationTags = initial.RecordApplicationTags
	}
	if dcl.IsZeroValue(des.QueryStringLength) {
		des.QueryStringLength = initial.QueryStringLength
	}

	return des
}

func canonicalizeNewInstanceSettingsInsightsConfig(c *Client, des, nw *InstanceSettingsInsightsConfig) *InstanceSettingsInsightsConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.QueryInsightsEnabled, nw.QueryInsightsEnabled) {
		nw.QueryInsightsEnabled = des.QueryInsightsEnabled
	}
	if dcl.BoolCanonicalize(des.RecordClientAddress, nw.RecordClientAddress) {
		nw.RecordClientAddress = des.RecordClientAddress
	}
	if dcl.BoolCanonicalize(des.RecordApplicationTags, nw.RecordApplicationTags) {
		nw.RecordApplicationTags = des.RecordApplicationTags
	}

	return nw
}

func canonicalizeNewInstanceSettingsInsightsConfigSet(c *Client, des, nw []InstanceSettingsInsightsConfig) []InstanceSettingsInsightsConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceSettingsInsightsConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceSettingsInsightsConfig(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceSettingsInsightsConfigSlice(c *Client, des, nw []InstanceSettingsInsightsConfig) []InstanceSettingsInsightsConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceSettingsInsightsConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceSettingsInsightsConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceReplicaInstances(des, initial *InstanceReplicaInstances, opts ...dcl.ApplyOption) *InstanceReplicaInstances {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		des.Name = initial.Name
	}
	if dcl.StringCanonicalize(des.Region, initial.Region) || dcl.IsZeroValue(des.Region) {
		des.Region = initial.Region
	}

	return des
}

func canonicalizeNewInstanceReplicaInstances(c *Client, des, nw *InstanceReplicaInstances) *InstanceReplicaInstances {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Region, nw.Region) {
		nw.Region = des.Region
	}

	return nw
}

func canonicalizeNewInstanceReplicaInstancesSet(c *Client, des, nw []InstanceReplicaInstances) []InstanceReplicaInstances {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceReplicaInstances
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceReplicaInstances(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceReplicaInstancesSlice(c *Client, des, nw []InstanceReplicaInstances) []InstanceReplicaInstances {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceReplicaInstances
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceReplicaInstances(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceServerCaCert(des, initial *InstanceServerCaCert, opts ...dcl.ApplyOption) *InstanceServerCaCert {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.CertSerialNumber, initial.CertSerialNumber) || dcl.IsZeroValue(des.CertSerialNumber) {
		des.CertSerialNumber = initial.CertSerialNumber
	}
	if dcl.StringCanonicalize(des.Cert, initial.Cert) || dcl.IsZeroValue(des.Cert) {
		des.Cert = initial.Cert
	}
	if dcl.IsZeroValue(des.CreateTime) {
		des.CreateTime = initial.CreateTime
	}
	if dcl.StringCanonicalize(des.CommonName, initial.CommonName) || dcl.IsZeroValue(des.CommonName) {
		des.CommonName = initial.CommonName
	}
	if dcl.IsZeroValue(des.ExpirationTime) {
		des.ExpirationTime = initial.ExpirationTime
	}
	if dcl.StringCanonicalize(des.Sha1Fingerprint, initial.Sha1Fingerprint) || dcl.IsZeroValue(des.Sha1Fingerprint) {
		des.Sha1Fingerprint = initial.Sha1Fingerprint
	}
	if dcl.StringCanonicalize(des.Instance, initial.Instance) || dcl.IsZeroValue(des.Instance) {
		des.Instance = initial.Instance
	}

	return des
}

func canonicalizeNewInstanceServerCaCert(c *Client, des, nw *InstanceServerCaCert) *InstanceServerCaCert {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.CertSerialNumber, nw.CertSerialNumber) {
		nw.CertSerialNumber = des.CertSerialNumber
	}
	if dcl.StringCanonicalize(des.Cert, nw.Cert) {
		nw.Cert = des.Cert
	}
	if dcl.StringCanonicalize(des.CommonName, nw.CommonName) {
		nw.CommonName = des.CommonName
	}
	if dcl.StringCanonicalize(des.Sha1Fingerprint, nw.Sha1Fingerprint) {
		nw.Sha1Fingerprint = des.Sha1Fingerprint
	}
	if dcl.StringCanonicalize(des.Instance, nw.Instance) {
		nw.Instance = des.Instance
	}

	return nw
}

func canonicalizeNewInstanceServerCaCertSet(c *Client, des, nw []InstanceServerCaCert) []InstanceServerCaCert {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceServerCaCert
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceServerCaCert(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceServerCaCertSlice(c *Client, des, nw []InstanceServerCaCert) []InstanceServerCaCert {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceServerCaCert
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceServerCaCert(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceOnPremisesConfiguration(des, initial *InstanceOnPremisesConfiguration, opts ...dcl.ApplyOption) *InstanceOnPremisesConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.HostPort, initial.HostPort) || dcl.IsZeroValue(des.HostPort) {
		des.HostPort = initial.HostPort
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}
	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		des.Username = initial.Username
	}
	if dcl.StringCanonicalize(des.Password, initial.Password) || dcl.IsZeroValue(des.Password) {
		des.Password = initial.Password
	}
	if dcl.StringCanonicalize(des.CaCertificate, initial.CaCertificate) || dcl.IsZeroValue(des.CaCertificate) {
		des.CaCertificate = initial.CaCertificate
	}
	if dcl.StringCanonicalize(des.ClientCertificate, initial.ClientCertificate) || dcl.IsZeroValue(des.ClientCertificate) {
		des.ClientCertificate = initial.ClientCertificate
	}
	if dcl.StringCanonicalize(des.ClientKey, initial.ClientKey) || dcl.IsZeroValue(des.ClientKey) {
		des.ClientKey = initial.ClientKey
	}
	if dcl.StringCanonicalize(des.DumpFilePath, initial.DumpFilePath) || dcl.IsZeroValue(des.DumpFilePath) {
		des.DumpFilePath = initial.DumpFilePath
	}
	if dcl.StringCanonicalize(des.Database, initial.Database) || dcl.IsZeroValue(des.Database) {
		des.Database = initial.Database
	}
	if dcl.IsZeroValue(des.ReplicatedDatabases) {
		des.ReplicatedDatabases = initial.ReplicatedDatabases
	}

	return des
}

func canonicalizeNewInstanceOnPremisesConfiguration(c *Client, des, nw *InstanceOnPremisesConfiguration) *InstanceOnPremisesConfiguration {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.HostPort, nw.HostPort) {
		nw.HostPort = des.HostPort
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}
	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}
	if dcl.StringCanonicalize(des.Password, nw.Password) {
		nw.Password = des.Password
	}
	if dcl.StringCanonicalize(des.CaCertificate, nw.CaCertificate) {
		nw.CaCertificate = des.CaCertificate
	}
	if dcl.StringCanonicalize(des.ClientCertificate, nw.ClientCertificate) {
		nw.ClientCertificate = des.ClientCertificate
	}
	if dcl.StringCanonicalize(des.ClientKey, nw.ClientKey) {
		nw.ClientKey = des.ClientKey
	}
	if dcl.StringCanonicalize(des.DumpFilePath, nw.DumpFilePath) {
		nw.DumpFilePath = des.DumpFilePath
	}
	if dcl.StringCanonicalize(des.Database, nw.Database) {
		nw.Database = des.Database
	}

	return nw
}

func canonicalizeNewInstanceOnPremisesConfigurationSet(c *Client, des, nw []InstanceOnPremisesConfiguration) []InstanceOnPremisesConfiguration {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceOnPremisesConfiguration
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceOnPremisesConfiguration(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceOnPremisesConfigurationSlice(c *Client, des, nw []InstanceOnPremisesConfiguration) []InstanceOnPremisesConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceOnPremisesConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceOnPremisesConfiguration(c, &d, &n))
	}

	return items
}

func canonicalizeInstanceDiskEncryptionStatus(des, initial *InstanceDiskEncryptionStatus, opts ...dcl.ApplyOption) *InstanceDiskEncryptionStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.KmsKeyVersionName, initial.KmsKeyVersionName) || dcl.IsZeroValue(des.KmsKeyVersionName) {
		des.KmsKeyVersionName = initial.KmsKeyVersionName
	}
	if dcl.StringCanonicalize(des.Kind, initial.Kind) || dcl.IsZeroValue(des.Kind) {
		des.Kind = initial.Kind
	}

	return des
}

func canonicalizeNewInstanceDiskEncryptionStatus(c *Client, des, nw *InstanceDiskEncryptionStatus) *InstanceDiskEncryptionStatus {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KmsKeyVersionName, nw.KmsKeyVersionName) {
		nw.KmsKeyVersionName = des.KmsKeyVersionName
	}
	if dcl.StringCanonicalize(des.Kind, nw.Kind) {
		nw.Kind = des.Kind
	}

	return nw
}

func canonicalizeNewInstanceDiskEncryptionStatusSet(c *Client, des, nw []InstanceDiskEncryptionStatus) []InstanceDiskEncryptionStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []InstanceDiskEncryptionStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if !compareInstanceDiskEncryptionStatus(c, &d, &n) {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewInstanceDiskEncryptionStatusSlice(c *Client, des, nw []InstanceDiskEncryptionStatus) []InstanceDiskEncryptionStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return des
	}

	var items []InstanceDiskEncryptionStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInstanceDiskEncryptionStatus(c, &d, &n))
	}

	return items
}

type instanceDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         instanceApiOperation
	Diffs            []*dcl.FieldDiff
	// This is for reporting only.
	FieldName string
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffInstance(c *Client, desired, actual *Instance, opts ...dcl.ApplyOption) ([]instanceDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var diffs []instanceDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.BackendType, actual.BackendType, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "EnumType", FieldName: "backend_type"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.ConnectionName, actual.ConnectionName, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "connection_name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.DatabaseVersion, actual.DatabaseVersion, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "EnumType", FieldName: "database_version"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "etag"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.GceZone, actual.GceZone, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "gce_zone"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.InstanceType, actual.InstanceType, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "EnumType", FieldName: "instance_type"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.MasterInstanceName, actual.MasterInstanceName, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "master_instance_name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "name"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "project"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.Region, actual.Region, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "region"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.RootPassword, actual.RootPassword, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "root_password"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{RequiresRecreate: true, Diffs: ds})
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "state"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.IPv6Address, actual.IPv6Address, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "ipv6_address"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.ServiceAccountEmailAddress, actual.ServiceAccountEmailAddress, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "service_account_email_address"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.SuspensionReason, actual.SuspensionReason, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "suspension_reason"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceUpdateOperation{}, Diffs: ds,
		})
	}

	if ds, err := dcl.Diff(desired.InstanceUid, actual.InstanceUid, dcl.Info{Ignore: false, OutputOnly: false, IgnoredPrefixes: []string(nil), Type: "", FieldName: "instance_uid"}); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, instanceDiff{
			UpdateOp: &updateInstanceUpdateOperation{}, Diffs: ds,
		})
	}

	if compareInstanceMaxDiskSize(c, desired.MaxDiskSize, actual.MaxDiskSize) {
		c.Config.Logger.Infof("Detected diff in MaxDiskSize.\nDESIRED: %v\nACTUAL: %v", desired.MaxDiskSize, actual.MaxDiskSize)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MaxDiskSize",
		})
	}
	if compareInstanceCurrentDiskSize(c, desired.CurrentDiskSize, actual.CurrentDiskSize) {
		c.Config.Logger.Infof("Detected diff in CurrentDiskSize.\nDESIRED: %v\nACTUAL: %v", desired.CurrentDiskSize, actual.CurrentDiskSize)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "CurrentDiskSize",
		})
	}
	if compareInstanceDiskEncryptionConfiguration(c, desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration) {
		c.Config.Logger.Infof("Detected diff in DiskEncryptionConfiguration.\nDESIRED: %v\nACTUAL: %v", desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "DiskEncryptionConfiguration",
		})
	}
	if compareInstanceFailoverReplica(c, desired.FailoverReplica, actual.FailoverReplica) {
		c.Config.Logger.Infof("Detected diff in FailoverReplica.\nDESIRED: %v\nACTUAL: %v", desired.FailoverReplica, actual.FailoverReplica)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "FailoverReplica",
		})
	}
	if compareInstanceIPAddressesSlice(c, desired.IPAddresses, actual.IPAddresses) {
		c.Config.Logger.Infof("Detected diff in IPAddresses.\nDESIRED: %v\nACTUAL: %v", desired.IPAddresses, actual.IPAddresses)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "IPAddresses",
		})
	}
	if compareInstanceMasterInstance(c, desired.MasterInstance, actual.MasterInstance) {
		c.Config.Logger.Infof("Detected diff in MasterInstance.\nDESIRED: %v\nACTUAL: %v", desired.MasterInstance, actual.MasterInstance)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "MasterInstance",
		})
	}
	if compareInstanceReplicaConfiguration(c, desired.ReplicaConfiguration, actual.ReplicaConfiguration) {
		c.Config.Logger.Infof("Detected diff in ReplicaConfiguration.\nDESIRED: %v\nACTUAL: %v", desired.ReplicaConfiguration, actual.ReplicaConfiguration)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ReplicaConfiguration",
		})
	}
	if compareInstanceScheduledMaintenance(c, desired.ScheduledMaintenance, actual.ScheduledMaintenance) {
		c.Config.Logger.Infof("Detected diff in ScheduledMaintenance.\nDESIRED: %v\nACTUAL: %v", desired.ScheduledMaintenance, actual.ScheduledMaintenance)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "ScheduledMaintenance",
		})
	}
	if compareInstanceSettings(c, desired.Settings, actual.Settings) {
		c.Config.Logger.Infof("Detected diff in Settings.\nDESIRED: %v\nACTUAL: %v", desired.Settings, actual.Settings)
		diffs = append(diffs, instanceDiff{
			RequiresRecreate: true,
			FieldName:        "Settings",
		})
	}
	if compareInstanceReplicaInstancesSlice(c, desired.ReplicaInstances, actual.ReplicaInstances) {
		c.Config.Logger.Infof("Detected diff in ReplicaInstances.\nDESIRED: %v\nACTUAL: %v", desired.ReplicaInstances, actual.ReplicaInstances)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "ReplicaInstances",
		})

	}
	if compareInstanceServerCaCert(c, desired.ServerCaCert, actual.ServerCaCert) {
		c.Config.Logger.Infof("Detected diff in ServerCaCert.\nDESIRED: %v\nACTUAL: %v", desired.ServerCaCert, actual.ServerCaCert)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "ServerCaCert",
		})

	}
	if compareInstanceOnPremisesConfiguration(c, desired.OnPremisesConfiguration, actual.OnPremisesConfiguration) {
		c.Config.Logger.Infof("Detected diff in OnPremisesConfiguration.\nDESIRED: %v\nACTUAL: %v", desired.OnPremisesConfiguration, actual.OnPremisesConfiguration)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "OnPremisesConfiguration",
		})

	}
	if compareInstanceDiskEncryptionStatus(c, desired.DiskEncryptionStatus, actual.DiskEncryptionStatus) {
		c.Config.Logger.Infof("Detected diff in DiskEncryptionStatus.\nDESIRED: %v\nACTUAL: %v", desired.DiskEncryptionStatus, actual.DiskEncryptionStatus)

		diffs = append(diffs, instanceDiff{
			UpdateOp:  &updateInstanceUpdateOperation{},
			FieldName: "DiskEncryptionStatus",
		})

	}
	// We need to ensure that this list does not contain identical operations *most of the time*.
	// There may be some cases where we will need multiple copies of the same operation - for instance,
	// if a resource has multiple prerequisite-containing fields.  For now, we don't know of any
	// such examples and so we deduplicate unconditionally.

	// The best way for us to do this is to iterate through the list
	// and remove any copies of operations which are identical to a previous operation.
	// This is O(n^2) in the number of operations, but n will always be very small,
	// even 10 would be an extremely high number.
	var opTypes []string
	var deduped []instanceDiff
	for _, d := range diffs {
		// Two operations are considered identical if they have the same type.
		// The type of an operation is derived from the name of the update method.
		if !dcl.StringSliceContains(fmt.Sprintf("%T", d.UpdateOp), opTypes) {
			deduped = append(deduped, d)
			opTypes = append(opTypes, fmt.Sprintf("%T", d.UpdateOp))
		} else {
			c.Config.Logger.Infof("Omitting planned operation of type %T since once is already scheduled.", d.UpdateOp)
		}
	}

	return deduped, nil
}
func compareInstanceMaxDiskSize(c *Client, desired, actual *InstanceMaxDiskSize) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceMaxDiskSizeSlice(c *Client, desired, actual []InstanceMaxDiskSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaxDiskSize, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMaxDiskSize(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMaxDiskSize, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMaxDiskSizeMap(c *Client, desired, actual map[string]InstanceMaxDiskSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMaxDiskSize, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceMaxDiskSize, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceMaxDiskSize(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceMaxDiskSize, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceCurrentDiskSize(c *Client, desired, actual *InstanceCurrentDiskSize) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceCurrentDiskSizeSlice(c *Client, desired, actual []InstanceCurrentDiskSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceCurrentDiskSize, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceCurrentDiskSize(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceCurrentDiskSize, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceCurrentDiskSizeMap(c *Client, desired, actual map[string]InstanceCurrentDiskSize) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceCurrentDiskSize, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceCurrentDiskSize, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceCurrentDiskSize(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceCurrentDiskSize, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceDiskEncryptionConfiguration(c *Client, desired, actual *InstanceDiskEncryptionConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.KmsKeyName, actual.KmsKeyName) && !dcl.IsZeroValue(desired.KmsKeyName) {
		c.Config.Logger.Infof("Diff in KmsKeyName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyName), dcl.SprintResource(actual.KmsKeyName))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareInstanceDiskEncryptionConfigurationSlice(c *Client, desired, actual []InstanceDiskEncryptionConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDiskEncryptionConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDiskEncryptionConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDiskEncryptionConfigurationMap(c *Client, desired, actual map[string]InstanceDiskEncryptionConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDiskEncryptionConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceDiskEncryptionConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceFailoverReplica(c *Client, desired, actual *InstanceFailoverReplica) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Available, actual.Available) && !dcl.IsZeroValue(desired.Available) {
		c.Config.Logger.Infof("Diff in Available.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Available), dcl.SprintResource(actual.Available))
		return true
	}
	if compareInstanceFailoverReplicaFailoverInstance(c, desired.FailoverInstance, actual.FailoverInstance) && !dcl.IsZeroValue(desired.FailoverInstance) {
		c.Config.Logger.Infof("Diff in FailoverInstance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailoverInstance), dcl.SprintResource(actual.FailoverInstance))
		return true
	}
	return false
}

func compareInstanceFailoverReplicaSlice(c *Client, desired, actual []InstanceFailoverReplica) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceFailoverReplica, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceFailoverReplica(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplica, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceFailoverReplicaMap(c *Client, desired, actual map[string]InstanceFailoverReplica) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceFailoverReplica, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplica, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceFailoverReplica(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplica, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceFailoverReplicaFailoverInstance(c *Client, desired, actual *InstanceFailoverReplicaFailoverInstance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) {
		c.Config.Logger.Infof("Diff in Region.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}

func compareInstanceFailoverReplicaFailoverInstanceSlice(c *Client, desired, actual []InstanceFailoverReplicaFailoverInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceFailoverReplicaFailoverInstance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceFailoverReplicaFailoverInstance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplicaFailoverInstance, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceFailoverReplicaFailoverInstanceMap(c *Client, desired, actual map[string]InstanceFailoverReplicaFailoverInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceFailoverReplicaFailoverInstance, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplicaFailoverInstance, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceFailoverReplicaFailoverInstance(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceFailoverReplicaFailoverInstance, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceIPAddresses(c *Client, desired, actual *InstanceIPAddresses) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Type, actual.Type) && !dcl.IsZeroValue(desired.Type) {
		c.Config.Logger.Infof("Diff in Type.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Type), dcl.SprintResource(actual.Type))
		return true
	}
	if !dcl.StringCanonicalize(desired.IPAddress, actual.IPAddress) && !dcl.IsZeroValue(desired.IPAddress) {
		c.Config.Logger.Infof("Diff in IPAddress.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPAddress), dcl.SprintResource(actual.IPAddress))
		return true
	}
	if compareInstanceIPAddressesTimeToRetire(c, desired.TimeToRetire, actual.TimeToRetire) && !dcl.IsZeroValue(desired.TimeToRetire) {
		c.Config.Logger.Infof("Diff in TimeToRetire.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TimeToRetire), dcl.SprintResource(actual.TimeToRetire))
		return true
	}
	return false
}

func compareInstanceIPAddressesSlice(c *Client, desired, actual []InstanceIPAddresses) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddresses, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceIPAddresses(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceIPAddresses, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceIPAddressesMap(c *Client, desired, actual map[string]InstanceIPAddresses) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddresses, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceIPAddresses, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceIPAddresses(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceIPAddresses, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceIPAddressesTimeToRetire(c *Client, desired, actual *InstanceIPAddressesTimeToRetire) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareInstanceIPAddressesTimeToRetireSlice(c *Client, desired, actual []InstanceIPAddressesTimeToRetire) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddressesTimeToRetire, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceIPAddressesTimeToRetire(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceIPAddressesTimeToRetire, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceIPAddressesTimeToRetireMap(c *Client, desired, actual map[string]InstanceIPAddressesTimeToRetire) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddressesTimeToRetire, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceIPAddressesTimeToRetire, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceIPAddressesTimeToRetire(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceIPAddressesTimeToRetire, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceMasterInstance(c *Client, desired, actual *InstanceMasterInstance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) {
		c.Config.Logger.Infof("Diff in Region.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}

func compareInstanceMasterInstanceSlice(c *Client, desired, actual []InstanceMasterInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMasterInstance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceMasterInstance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceMasterInstance, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceMasterInstanceMap(c *Client, desired, actual map[string]InstanceMasterInstance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceMasterInstance, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceMasterInstance, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceMasterInstance(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceMasterInstance, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfiguration(c *Client, desired, actual *InstanceReplicaConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, desired.MysqlReplicaConfiguration, actual.MysqlReplicaConfiguration) && !dcl.IsZeroValue(desired.MysqlReplicaConfiguration) {
		c.Config.Logger.Infof("Diff in MysqlReplicaConfiguration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MysqlReplicaConfiguration), dcl.SprintResource(actual.MysqlReplicaConfiguration))
		return true
	}
	if !dcl.BoolCanonicalize(desired.FailoverTarget, actual.FailoverTarget) && !dcl.IsZeroValue(desired.FailoverTarget) {
		c.Config.Logger.Infof("Diff in FailoverTarget.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.FailoverTarget), dcl.SprintResource(actual.FailoverTarget))
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, desired.ReplicaPoolConfiguration, actual.ReplicaPoolConfiguration) && !dcl.IsZeroValue(desired.ReplicaPoolConfiguration) {
		c.Config.Logger.Infof("Diff in ReplicaPoolConfiguration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaPoolConfiguration), dcl.SprintResource(actual.ReplicaPoolConfiguration))
		return true
	}
	return false
}

func compareInstanceReplicaConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMap(c *Client, desired, actual map[string]InstanceReplicaConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationMysqlReplicaConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.DumpFilePath, actual.DumpFilePath) && !dcl.IsZeroValue(desired.DumpFilePath) {
		c.Config.Logger.Infof("Diff in DumpFilePath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DumpFilePath), dcl.SprintResource(actual.DumpFilePath))
		return true
	}
	if !dcl.StringCanonicalize(desired.Username, actual.Username) && !dcl.IsZeroValue(desired.Username) {
		c.Config.Logger.Infof("Diff in Username.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Username), dcl.SprintResource(actual.Username))
		return true
	}
	if !dcl.StringCanonicalize(desired.Password, actual.Password) && !dcl.IsZeroValue(desired.Password) {
		c.Config.Logger.Infof("Diff in Password.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Password), dcl.SprintResource(actual.Password))
		return true
	}
	if !reflect.DeepEqual(desired.ConnectRetryInterval, actual.ConnectRetryInterval) && !dcl.IsZeroValue(desired.ConnectRetryInterval) {
		c.Config.Logger.Infof("Diff in ConnectRetryInterval.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ConnectRetryInterval), dcl.SprintResource(actual.ConnectRetryInterval))
		return true
	}
	if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, desired.MasterHeartbeatPeriod, actual.MasterHeartbeatPeriod) && !dcl.IsZeroValue(desired.MasterHeartbeatPeriod) {
		c.Config.Logger.Infof("Diff in MasterHeartbeatPeriod.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MasterHeartbeatPeriod), dcl.SprintResource(actual.MasterHeartbeatPeriod))
		return true
	}
	if !dcl.StringCanonicalize(desired.CaCertificate, actual.CaCertificate) && !dcl.IsZeroValue(desired.CaCertificate) {
		c.Config.Logger.Infof("Diff in CaCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CaCertificate), dcl.SprintResource(actual.CaCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClientCertificate, actual.ClientCertificate) && !dcl.IsZeroValue(desired.ClientCertificate) {
		c.Config.Logger.Infof("Diff in ClientCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificate), dcl.SprintResource(actual.ClientCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClientKey, actual.ClientKey) && !dcl.IsZeroValue(desired.ClientKey) {
		c.Config.Logger.Infof("Diff in ClientKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientKey), dcl.SprintResource(actual.ClientKey))
		return true
	}
	if !dcl.StringCanonicalize(desired.SslCipher, actual.SslCipher) && !dcl.IsZeroValue(desired.SslCipher) {
		c.Config.Logger.Infof("Diff in SslCipher.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SslCipher), dcl.SprintResource(actual.SslCipher))
		return true
	}
	if !dcl.BoolCanonicalize(desired.VerifyServerCertificate, actual.VerifyServerCertificate) && !dcl.IsZeroValue(desired.VerifyServerCertificate) {
		c.Config.Logger.Infof("Diff in VerifyServerCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.VerifyServerCertificate), dcl.SprintResource(actual.VerifyServerCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationMysqlReplicaConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfigurationMap(c *Client, desired, actual map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, desired, actual *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, desired, actual []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap(c *Client, desired, actual map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationReplicaPoolConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, desired.StaticPoolConfiguration, actual.StaticPoolConfiguration) && !dcl.IsZeroValue(desired.StaticPoolConfiguration) {
		c.Config.Logger.Infof("Diff in StaticPoolConfiguration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StaticPoolConfiguration), dcl.SprintResource(actual.StaticPoolConfiguration))
		return true
	}
	if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, desired.AutoscalingPoolConfiguration, actual.AutoscalingPoolConfiguration) && !dcl.IsZeroValue(desired.AutoscalingPoolConfiguration) {
		c.Config.Logger.Infof("Diff in AutoscalingPoolConfiguration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AutoscalingPoolConfiguration), dcl.SprintResource(actual.AutoscalingPoolConfiguration))
		return true
	}
	if !reflect.DeepEqual(desired.ReplicaCount, actual.ReplicaCount) && !dcl.IsZeroValue(desired.ReplicaCount) {
		c.Config.Logger.Infof("Diff in ReplicaCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaCount), dcl.SprintResource(actual.ReplicaCount))
		return true
	}
	if !dcl.BoolCanonicalize(desired.ExposeReplicaIP, actual.ExposeReplicaIP) && !dcl.IsZeroValue(desired.ExposeReplicaIP) {
		c.Config.Logger.Infof("Diff in ExposeReplicaIP.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeReplicaIP), dcl.SprintResource(actual.ExposeReplicaIP))
		return true
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationReplicaPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationMap(c *Client, desired, actual map[string]InstanceReplicaConfigurationReplicaPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaConfigurationReplicaPoolConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.ReplicaCount, actual.ReplicaCount) && !dcl.IsZeroValue(desired.ReplicaCount) {
		c.Config.Logger.Infof("Diff in ReplicaCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicaCount), dcl.SprintResource(actual.ReplicaCount))
		return true
	}
	if !dcl.BoolCanonicalize(desired.ExposeReplicaIP, actual.ExposeReplicaIP) && !dcl.IsZeroValue(desired.ExposeReplicaIP) {
		c.Config.Logger.Infof("Diff in ExposeReplicaIP.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExposeReplicaIP), dcl.SprintResource(actual.ExposeReplicaIP))
		return true
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap(c *Client, desired, actual map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, desired, actual *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.MinReplicaCount, actual.MinReplicaCount) && !dcl.IsZeroValue(desired.MinReplicaCount) {
		c.Config.Logger.Infof("Diff in MinReplicaCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MinReplicaCount), dcl.SprintResource(actual.MinReplicaCount))
		return true
	}
	if !reflect.DeepEqual(desired.MaxReplicaCount, actual.MaxReplicaCount) && !dcl.IsZeroValue(desired.MaxReplicaCount) {
		c.Config.Logger.Infof("Diff in MaxReplicaCount.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaxReplicaCount), dcl.SprintResource(actual.MaxReplicaCount))
		return true
	}
	if !reflect.DeepEqual(desired.TargetCpuUtil, actual.TargetCpuUtil) && !dcl.IsZeroValue(desired.TargetCpuUtil) {
		c.Config.Logger.Infof("Diff in TargetCpuUtil.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TargetCpuUtil), dcl.SprintResource(actual.TargetCpuUtil))
		return true
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, desired, actual []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap(c *Client, desired, actual map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceScheduledMaintenance(c *Client, desired, actual *InstanceScheduledMaintenance) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if compareInstanceScheduledMaintenanceStartTime(c, desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if !dcl.BoolCanonicalize(desired.CanDefer, actual.CanDefer) && !dcl.IsZeroValue(desired.CanDefer) {
		c.Config.Logger.Infof("Diff in CanDefer.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanDefer), dcl.SprintResource(actual.CanDefer))
		return true
	}
	if !dcl.BoolCanonicalize(desired.CanReschedule, actual.CanReschedule) && !dcl.IsZeroValue(desired.CanReschedule) {
		c.Config.Logger.Infof("Diff in CanReschedule.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CanReschedule), dcl.SprintResource(actual.CanReschedule))
		return true
	}
	return false
}

func compareInstanceScheduledMaintenanceSlice(c *Client, desired, actual []InstanceScheduledMaintenance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduledMaintenance, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceScheduledMaintenance(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenance, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceScheduledMaintenanceMap(c *Client, desired, actual map[string]InstanceScheduledMaintenance) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduledMaintenance, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenance, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceScheduledMaintenance(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenance, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceScheduledMaintenanceStartTime(c *Client, desired, actual *InstanceScheduledMaintenanceStartTime) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Seconds, actual.Seconds) && !dcl.IsZeroValue(desired.Seconds) {
		c.Config.Logger.Infof("Diff in Seconds.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Seconds), dcl.SprintResource(actual.Seconds))
		return true
	}
	if !reflect.DeepEqual(desired.Nanos, actual.Nanos) && !dcl.IsZeroValue(desired.Nanos) {
		c.Config.Logger.Infof("Diff in Nanos.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Nanos), dcl.SprintResource(actual.Nanos))
		return true
	}
	return false
}

func compareInstanceScheduledMaintenanceStartTimeSlice(c *Client, desired, actual []InstanceScheduledMaintenanceStartTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduledMaintenanceStartTime, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceScheduledMaintenanceStartTime(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenanceStartTime, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceScheduledMaintenanceStartTimeMap(c *Client, desired, actual map[string]InstanceScheduledMaintenanceStartTime) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceScheduledMaintenanceStartTime, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenanceStartTime, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceScheduledMaintenanceStartTime(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceScheduledMaintenanceStartTime, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettings(c *Client, desired, actual *InstanceSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringSliceEquals(desired.AuthorizedGaeApplications, actual.AuthorizedGaeApplications) && !dcl.IsZeroValue(desired.AuthorizedGaeApplications) {
		c.Config.Logger.Infof("Diff in AuthorizedGaeApplications.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AuthorizedGaeApplications), dcl.SprintResource(actual.AuthorizedGaeApplications))
		return true
	}
	if !dcl.StringCanonicalize(desired.Tier, actual.Tier) && !dcl.IsZeroValue(desired.Tier) {
		c.Config.Logger.Infof("Diff in Tier.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Tier), dcl.SprintResource(actual.Tier))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !reflect.DeepEqual(desired.AvailabilityType, actual.AvailabilityType) && !dcl.IsZeroValue(desired.AvailabilityType) {
		c.Config.Logger.Infof("Diff in AvailabilityType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AvailabilityType), dcl.SprintResource(actual.AvailabilityType))
		return true
	}
	if !reflect.DeepEqual(desired.PricingPlan, actual.PricingPlan) && !dcl.IsZeroValue(desired.PricingPlan) {
		c.Config.Logger.Infof("Diff in PricingPlan.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PricingPlan), dcl.SprintResource(actual.PricingPlan))
		return true
	}
	if !reflect.DeepEqual(desired.ReplicationType, actual.ReplicationType) && !dcl.IsZeroValue(desired.ReplicationType) {
		c.Config.Logger.Infof("Diff in ReplicationType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicationType), dcl.SprintResource(actual.ReplicationType))
		return true
	}
	if !reflect.DeepEqual(desired.ActivationPolicy, actual.ActivationPolicy) && !dcl.IsZeroValue(desired.ActivationPolicy) {
		c.Config.Logger.Infof("Diff in ActivationPolicy.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ActivationPolicy), dcl.SprintResource(actual.ActivationPolicy))
		return true
	}
	if !dcl.BoolCanonicalize(desired.StorageAutoResize, actual.StorageAutoResize) && !dcl.IsZeroValue(desired.StorageAutoResize) {
		c.Config.Logger.Infof("Diff in StorageAutoResize.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StorageAutoResize), dcl.SprintResource(actual.StorageAutoResize))
		return true
	}
	if !reflect.DeepEqual(desired.DataDiskType, actual.DataDiskType) && !dcl.IsZeroValue(desired.DataDiskType) {
		c.Config.Logger.Infof("Diff in DataDiskType.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DataDiskType), dcl.SprintResource(actual.DataDiskType))
		return true
	}
	if !dcl.BoolCanonicalize(desired.DatabaseReplicationEnabled, actual.DatabaseReplicationEnabled) && !dcl.IsZeroValue(desired.DatabaseReplicationEnabled) {
		c.Config.Logger.Infof("Diff in DatabaseReplicationEnabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatabaseReplicationEnabled), dcl.SprintResource(actual.DatabaseReplicationEnabled))
		return true
	}
	if !dcl.BoolCanonicalize(desired.CrashSafeReplicationEnabled, actual.CrashSafeReplicationEnabled) && !dcl.IsZeroValue(desired.CrashSafeReplicationEnabled) {
		c.Config.Logger.Infof("Diff in CrashSafeReplicationEnabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CrashSafeReplicationEnabled), dcl.SprintResource(actual.CrashSafeReplicationEnabled))
		return true
	}
	if compareInstanceSettingsSettingsVersion(c, desired.SettingsVersion, actual.SettingsVersion) && !dcl.IsZeroValue(desired.SettingsVersion) {
		c.Config.Logger.Infof("Diff in SettingsVersion.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.SettingsVersion), dcl.SprintResource(actual.SettingsVersion))
		return true
	}
	if !dcl.MapEquals(desired.UserLabels, actual.UserLabels, []string(nil)) && !dcl.IsZeroValue(desired.UserLabels) {
		c.Config.Logger.Infof("Diff in UserLabels.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UserLabels), dcl.SprintResource(actual.UserLabels))
		return true
	}
	if compareInstanceSettingsStorageAutoResizeLimit(c, desired.StorageAutoResizeLimit, actual.StorageAutoResizeLimit) && !dcl.IsZeroValue(desired.StorageAutoResizeLimit) {
		c.Config.Logger.Infof("Diff in StorageAutoResizeLimit.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StorageAutoResizeLimit), dcl.SprintResource(actual.StorageAutoResizeLimit))
		return true
	}
	if compareInstanceSettingsIPConfiguration(c, desired.IPConfiguration, actual.IPConfiguration) && !dcl.IsZeroValue(desired.IPConfiguration) {
		c.Config.Logger.Infof("Diff in IPConfiguration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPConfiguration), dcl.SprintResource(actual.IPConfiguration))
		return true
	}
	if compareInstanceSettingsLocationPreference(c, desired.LocationPreference, actual.LocationPreference) && !dcl.IsZeroValue(desired.LocationPreference) {
		c.Config.Logger.Infof("Diff in LocationPreference.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.LocationPreference), dcl.SprintResource(actual.LocationPreference))
		return true
	}
	if compareInstanceSettingsDatabaseFlagsSlice(c, desired.DatabaseFlags, actual.DatabaseFlags) && !dcl.IsZeroValue(desired.DatabaseFlags) {
		c.Config.Logger.Infof("Diff in DatabaseFlags.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DatabaseFlags), dcl.SprintResource(actual.DatabaseFlags))
		return true
	}
	if compareInstanceSettingsMaintenanceWindow(c, desired.MaintenanceWindow, actual.MaintenanceWindow) && !dcl.IsZeroValue(desired.MaintenanceWindow) {
		c.Config.Logger.Infof("Diff in MaintenanceWindow.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.MaintenanceWindow), dcl.SprintResource(actual.MaintenanceWindow))
		return true
	}
	if compareInstanceSettingsBackupConfiguration(c, desired.BackupConfiguration, actual.BackupConfiguration) && !dcl.IsZeroValue(desired.BackupConfiguration) {
		c.Config.Logger.Infof("Diff in BackupConfiguration.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackupConfiguration), dcl.SprintResource(actual.BackupConfiguration))
		return true
	}
	if compareInstanceSettingsDataDiskSizeGb(c, desired.DataDiskSizeGb, actual.DataDiskSizeGb) && !dcl.IsZeroValue(desired.DataDiskSizeGb) {
		c.Config.Logger.Infof("Diff in DataDiskSizeGb.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DataDiskSizeGb), dcl.SprintResource(actual.DataDiskSizeGb))
		return true
	}
	if compareInstanceSettingsActiveDirectoryConfig(c, desired.ActiveDirectoryConfig, actual.ActiveDirectoryConfig) && !dcl.IsZeroValue(desired.ActiveDirectoryConfig) {
		c.Config.Logger.Infof("Diff in ActiveDirectoryConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ActiveDirectoryConfig), dcl.SprintResource(actual.ActiveDirectoryConfig))
		return true
	}
	if !dcl.StringCanonicalize(desired.Collation, actual.Collation) && !dcl.IsZeroValue(desired.Collation) {
		c.Config.Logger.Infof("Diff in Collation.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Collation), dcl.SprintResource(actual.Collation))
		return true
	}
	if compareInstanceSettingsDenyMaintenancePeriodsSlice(c, desired.DenyMaintenancePeriods, actual.DenyMaintenancePeriods) && !dcl.IsZeroValue(desired.DenyMaintenancePeriods) {
		c.Config.Logger.Infof("Diff in DenyMaintenancePeriods.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DenyMaintenancePeriods), dcl.SprintResource(actual.DenyMaintenancePeriods))
		return true
	}
	if compareInstanceSettingsInsightsConfig(c, desired.InsightsConfig, actual.InsightsConfig) && !dcl.IsZeroValue(desired.InsightsConfig) {
		c.Config.Logger.Infof("Diff in InsightsConfig.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.InsightsConfig), dcl.SprintResource(actual.InsightsConfig))
		return true
	}
	return false
}

func compareInstanceSettingsSlice(c *Client, desired, actual []InstanceSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettings, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsMap(c *Client, desired, actual map[string]InstanceSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettings, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettings, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettings(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettings, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsSettingsVersion(c *Client, desired, actual *InstanceSettingsSettingsVersion) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceSettingsSettingsVersionSlice(c *Client, desired, actual []InstanceSettingsSettingsVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsSettingsVersion, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsSettingsVersion(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsSettingsVersion, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsSettingsVersionMap(c *Client, desired, actual map[string]InstanceSettingsSettingsVersion) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsSettingsVersion, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsSettingsVersion, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsSettingsVersion(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsSettingsVersion, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsStorageAutoResizeLimit(c *Client, desired, actual *InstanceSettingsStorageAutoResizeLimit) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceSettingsStorageAutoResizeLimitSlice(c *Client, desired, actual []InstanceSettingsStorageAutoResizeLimit) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsStorageAutoResizeLimit, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsStorageAutoResizeLimit(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsStorageAutoResizeLimit, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsStorageAutoResizeLimitMap(c *Client, desired, actual map[string]InstanceSettingsStorageAutoResizeLimit) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsStorageAutoResizeLimit, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsStorageAutoResizeLimit, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsStorageAutoResizeLimit(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsStorageAutoResizeLimit, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsIPConfiguration(c *Client, desired, actual *InstanceSettingsIPConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.IPv4Enabled, actual.IPv4Enabled) && !dcl.IsZeroValue(desired.IPv4Enabled) {
		c.Config.Logger.Infof("Diff in IPv4Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.IPv4Enabled), dcl.SprintResource(actual.IPv4Enabled))
		return true
	}
	if !dcl.NameToSelfLink(desired.PrivateNetwork, actual.PrivateNetwork) && !dcl.IsZeroValue(desired.PrivateNetwork) {
		c.Config.Logger.Infof("Diff in PrivateNetwork.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.PrivateNetwork), dcl.SprintResource(actual.PrivateNetwork))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RequireSsl, actual.RequireSsl) && !dcl.IsZeroValue(desired.RequireSsl) {
		c.Config.Logger.Infof("Diff in RequireSsl.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RequireSsl), dcl.SprintResource(actual.RequireSsl))
		return true
	}
	if compareInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, desired.AuthorizedNetworks, actual.AuthorizedNetworks) && !dcl.IsZeroValue(desired.AuthorizedNetworks) {
		c.Config.Logger.Infof("Diff in AuthorizedNetworks.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.AuthorizedNetworks), dcl.SprintResource(actual.AuthorizedNetworks))
		return true
	}
	return false
}

func compareInstanceSettingsIPConfigurationSlice(c *Client, desired, actual []InstanceSettingsIPConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsIPConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsIPConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsIPConfigurationMap(c *Client, desired, actual map[string]InstanceSettingsIPConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsIPConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsIPConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, desired, actual *InstanceSettingsIPConfigurationAuthorizedNetworks) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	if !reflect.DeepEqual(desired.ExpirationTime, actual.ExpirationTime) && !dcl.IsZeroValue(desired.ExpirationTime) {
		c.Config.Logger.Infof("Diff in ExpirationTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExpirationTime), dcl.SprintResource(actual.ExpirationTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, desired, actual []InstanceSettingsIPConfigurationAuthorizedNetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsIPConfigurationAuthorizedNetworks(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsIPConfigurationAuthorizedNetworksMap(c *Client, desired, actual map[string]InstanceSettingsIPConfigurationAuthorizedNetworks) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsIPConfigurationAuthorizedNetworks(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsIPConfigurationAuthorizedNetworks, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsLocationPreference(c *Client, desired, actual *InstanceSettingsLocationPreference) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Zone, actual.Zone) && !dcl.IsZeroValue(desired.Zone) {
		c.Config.Logger.Infof("Diff in Zone.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Zone), dcl.SprintResource(actual.Zone))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareInstanceSettingsLocationPreferenceSlice(c *Client, desired, actual []InstanceSettingsLocationPreference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsLocationPreference, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsLocationPreference(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsLocationPreference, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsLocationPreferenceMap(c *Client, desired, actual map[string]InstanceSettingsLocationPreference) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsLocationPreference, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsLocationPreference, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsLocationPreference(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsLocationPreference, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDatabaseFlags(c *Client, desired, actual *InstanceSettingsDatabaseFlags) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceSettingsDatabaseFlagsSlice(c *Client, desired, actual []InstanceSettingsDatabaseFlags) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDatabaseFlags, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDatabaseFlags(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDatabaseFlags, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDatabaseFlagsMap(c *Client, desired, actual map[string]InstanceSettingsDatabaseFlags) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDatabaseFlags, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsDatabaseFlags, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsDatabaseFlags(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDatabaseFlags, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsMaintenanceWindow(c *Client, desired, actual *InstanceSettingsMaintenanceWindow) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Hour, actual.Hour) && !dcl.IsZeroValue(desired.Hour) {
		c.Config.Logger.Infof("Diff in Hour.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Hour), dcl.SprintResource(actual.Hour))
		return true
	}
	if !reflect.DeepEqual(desired.Day, actual.Day) && !dcl.IsZeroValue(desired.Day) {
		c.Config.Logger.Infof("Diff in Day.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Day), dcl.SprintResource(actual.Day))
		return true
	}
	if !reflect.DeepEqual(desired.UpdateTrack, actual.UpdateTrack) && !dcl.IsZeroValue(desired.UpdateTrack) {
		c.Config.Logger.Infof("Diff in UpdateTrack.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.UpdateTrack), dcl.SprintResource(actual.UpdateTrack))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareInstanceSettingsMaintenanceWindowSlice(c *Client, desired, actual []InstanceSettingsMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsMaintenanceWindow, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsMaintenanceWindow(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsMaintenanceWindow, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsMaintenanceWindowMap(c *Client, desired, actual map[string]InstanceSettingsMaintenanceWindow) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsMaintenanceWindow, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsMaintenanceWindow, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsMaintenanceWindow(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsMaintenanceWindow, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfiguration(c *Client, desired, actual *InstanceSettingsBackupConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.StartTime, actual.StartTime) && !dcl.IsZeroValue(desired.StartTime) {
		c.Config.Logger.Infof("Diff in StartTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartTime), dcl.SprintResource(actual.StartTime))
		return true
	}
	if !dcl.BoolCanonicalize(desired.Enabled, actual.Enabled) && !dcl.IsZeroValue(desired.Enabled) {
		c.Config.Logger.Infof("Diff in Enabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Enabled), dcl.SprintResource(actual.Enabled))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !dcl.BoolCanonicalize(desired.BinaryLogEnabled, actual.BinaryLogEnabled) && !dcl.IsZeroValue(desired.BinaryLogEnabled) {
		c.Config.Logger.Infof("Diff in BinaryLogEnabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BinaryLogEnabled), dcl.SprintResource(actual.BinaryLogEnabled))
		return true
	}
	if !dcl.StringCanonicalize(desired.Location, actual.Location) && !dcl.IsZeroValue(desired.Location) {
		c.Config.Logger.Infof("Diff in Location.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Location), dcl.SprintResource(actual.Location))
		return true
	}
	if compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, desired.BackupRetentionSettings, actual.BackupRetentionSettings) && !dcl.IsZeroValue(desired.BackupRetentionSettings) {
		c.Config.Logger.Infof("Diff in BackupRetentionSettings.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.BackupRetentionSettings), dcl.SprintResource(actual.BackupRetentionSettings))
		return true
	}
	if !reflect.DeepEqual(desired.TransactionLogRetentionDays, actual.TransactionLogRetentionDays) && !dcl.IsZeroValue(desired.TransactionLogRetentionDays) {
		c.Config.Logger.Infof("Diff in TransactionLogRetentionDays.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.TransactionLogRetentionDays), dcl.SprintResource(actual.TransactionLogRetentionDays))
		return true
	}
	return false
}

func compareInstanceSettingsBackupConfigurationSlice(c *Client, desired, actual []InstanceSettingsBackupConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsBackupConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfigurationMap(c *Client, desired, actual map[string]InstanceSettingsBackupConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsBackupConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, desired, actual *InstanceSettingsBackupConfigurationBackupRetentionSettings) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.RetentionUnit, actual.RetentionUnit) && !dcl.IsZeroValue(desired.RetentionUnit) {
		c.Config.Logger.Infof("Diff in RetentionUnit.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetentionUnit), dcl.SprintResource(actual.RetentionUnit))
		return true
	}
	if !reflect.DeepEqual(desired.RetainedBackups, actual.RetainedBackups) && !dcl.IsZeroValue(desired.RetainedBackups) {
		c.Config.Logger.Infof("Diff in RetainedBackups.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RetainedBackups), dcl.SprintResource(actual.RetainedBackups))
		return true
	}
	return false
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, desired, actual []InstanceSettingsBackupConfigurationBackupRetentionSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsMap(c *Client, desired, actual map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettings, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDataDiskSizeGb(c *Client, desired, actual *InstanceSettingsDataDiskSizeGb) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !reflect.DeepEqual(desired.Value, actual.Value) && !dcl.IsZeroValue(desired.Value) {
		c.Config.Logger.Infof("Diff in Value.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Value), dcl.SprintResource(actual.Value))
		return true
	}
	return false
}

func compareInstanceSettingsDataDiskSizeGbSlice(c *Client, desired, actual []InstanceSettingsDataDiskSizeGb) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDataDiskSizeGb, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDataDiskSizeGb(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDataDiskSizeGb, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDataDiskSizeGbMap(c *Client, desired, actual map[string]InstanceSettingsDataDiskSizeGb) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDataDiskSizeGb, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsDataDiskSizeGb, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsDataDiskSizeGb(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDataDiskSizeGb, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsActiveDirectoryConfig(c *Client, desired, actual *InstanceSettingsActiveDirectoryConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !dcl.StringCanonicalize(desired.Domain, actual.Domain) && !dcl.IsZeroValue(desired.Domain) {
		c.Config.Logger.Infof("Diff in Domain.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Domain), dcl.SprintResource(actual.Domain))
		return true
	}
	return false
}

func compareInstanceSettingsActiveDirectoryConfigSlice(c *Client, desired, actual []InstanceSettingsActiveDirectoryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsActiveDirectoryConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsActiveDirectoryConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsActiveDirectoryConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsActiveDirectoryConfigMap(c *Client, desired, actual map[string]InstanceSettingsActiveDirectoryConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsActiveDirectoryConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsActiveDirectoryConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsActiveDirectoryConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsActiveDirectoryConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDenyMaintenancePeriods(c *Client, desired, actual *InstanceSettingsDenyMaintenancePeriods) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.StartDate, actual.StartDate) && !dcl.IsZeroValue(desired.StartDate) {
		c.Config.Logger.Infof("Diff in StartDate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.StartDate), dcl.SprintResource(actual.StartDate))
		return true
	}
	if !dcl.StringCanonicalize(desired.EndDate, actual.EndDate) && !dcl.IsZeroValue(desired.EndDate) {
		c.Config.Logger.Infof("Diff in EndDate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.EndDate), dcl.SprintResource(actual.EndDate))
		return true
	}
	if !dcl.StringCanonicalize(desired.Time, actual.Time) && !dcl.IsZeroValue(desired.Time) {
		c.Config.Logger.Infof("Diff in Time.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Time), dcl.SprintResource(actual.Time))
		return true
	}
	return false
}

func compareInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, desired, actual []InstanceSettingsDenyMaintenancePeriods) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDenyMaintenancePeriods, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDenyMaintenancePeriods(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDenyMaintenancePeriods, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDenyMaintenancePeriodsMap(c *Client, desired, actual map[string]InstanceSettingsDenyMaintenancePeriods) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDenyMaintenancePeriods, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsDenyMaintenancePeriods, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsDenyMaintenancePeriods(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDenyMaintenancePeriods, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceSettingsInsightsConfig(c *Client, desired, actual *InstanceSettingsInsightsConfig) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.BoolCanonicalize(desired.QueryInsightsEnabled, actual.QueryInsightsEnabled) && !dcl.IsZeroValue(desired.QueryInsightsEnabled) {
		c.Config.Logger.Infof("Diff in QueryInsightsEnabled.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryInsightsEnabled), dcl.SprintResource(actual.QueryInsightsEnabled))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RecordClientAddress, actual.RecordClientAddress) && !dcl.IsZeroValue(desired.RecordClientAddress) {
		c.Config.Logger.Infof("Diff in RecordClientAddress.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RecordClientAddress), dcl.SprintResource(actual.RecordClientAddress))
		return true
	}
	if !dcl.BoolCanonicalize(desired.RecordApplicationTags, actual.RecordApplicationTags) && !dcl.IsZeroValue(desired.RecordApplicationTags) {
		c.Config.Logger.Infof("Diff in RecordApplicationTags.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.RecordApplicationTags), dcl.SprintResource(actual.RecordApplicationTags))
		return true
	}
	if !reflect.DeepEqual(desired.QueryStringLength, actual.QueryStringLength) && !dcl.IsZeroValue(desired.QueryStringLength) {
		c.Config.Logger.Infof("Diff in QueryStringLength.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.QueryStringLength), dcl.SprintResource(actual.QueryStringLength))
		return true
	}
	return false
}

func compareInstanceSettingsInsightsConfigSlice(c *Client, desired, actual []InstanceSettingsInsightsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsInsightsConfig, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsInsightsConfig(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsInsightsConfig, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsInsightsConfigMap(c *Client, desired, actual map[string]InstanceSettingsInsightsConfig) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsInsightsConfig, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceSettingsInsightsConfig, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceSettingsInsightsConfig(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceSettingsInsightsConfig, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceReplicaInstances(c *Client, desired, actual *InstanceReplicaInstances) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Name, actual.Name) && !dcl.IsZeroValue(desired.Name) {
		c.Config.Logger.Infof("Diff in Name.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Name), dcl.SprintResource(actual.Name))
		return true
	}
	if !dcl.StringCanonicalize(desired.Region, actual.Region) && !dcl.IsZeroValue(desired.Region) {
		c.Config.Logger.Infof("Diff in Region.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Region), dcl.SprintResource(actual.Region))
		return true
	}
	return false
}

func compareInstanceReplicaInstancesSlice(c *Client, desired, actual []InstanceReplicaInstances) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaInstances, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceReplicaInstances(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceReplicaInstances, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceReplicaInstancesMap(c *Client, desired, actual map[string]InstanceReplicaInstances) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceReplicaInstances, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceReplicaInstances, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceReplicaInstances(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceReplicaInstances, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceServerCaCert(c *Client, desired, actual *InstanceServerCaCert) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !dcl.StringCanonicalize(desired.CertSerialNumber, actual.CertSerialNumber) && !dcl.IsZeroValue(desired.CertSerialNumber) {
		c.Config.Logger.Infof("Diff in CertSerialNumber.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CertSerialNumber), dcl.SprintResource(actual.CertSerialNumber))
		return true
	}
	if !dcl.StringCanonicalize(desired.Cert, actual.Cert) && !dcl.IsZeroValue(desired.Cert) {
		c.Config.Logger.Infof("Diff in Cert.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Cert), dcl.SprintResource(actual.Cert))
		return true
	}
	if !reflect.DeepEqual(desired.CreateTime, actual.CreateTime) && !dcl.IsZeroValue(desired.CreateTime) {
		c.Config.Logger.Infof("Diff in CreateTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CreateTime), dcl.SprintResource(actual.CreateTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.CommonName, actual.CommonName) && !dcl.IsZeroValue(desired.CommonName) {
		c.Config.Logger.Infof("Diff in CommonName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CommonName), dcl.SprintResource(actual.CommonName))
		return true
	}
	if !reflect.DeepEqual(desired.ExpirationTime, actual.ExpirationTime) && !dcl.IsZeroValue(desired.ExpirationTime) {
		c.Config.Logger.Infof("Diff in ExpirationTime.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ExpirationTime), dcl.SprintResource(actual.ExpirationTime))
		return true
	}
	if !dcl.StringCanonicalize(desired.Sha1Fingerprint, actual.Sha1Fingerprint) && !dcl.IsZeroValue(desired.Sha1Fingerprint) {
		c.Config.Logger.Infof("Diff in Sha1Fingerprint.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Sha1Fingerprint), dcl.SprintResource(actual.Sha1Fingerprint))
		return true
	}
	if !dcl.StringCanonicalize(desired.Instance, actual.Instance) && !dcl.IsZeroValue(desired.Instance) {
		c.Config.Logger.Infof("Diff in Instance.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Instance), dcl.SprintResource(actual.Instance))
		return true
	}
	return false
}

func compareInstanceServerCaCertSlice(c *Client, desired, actual []InstanceServerCaCert) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServerCaCert, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceServerCaCert(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceServerCaCert, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceServerCaCertMap(c *Client, desired, actual map[string]InstanceServerCaCert) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceServerCaCert, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceServerCaCert, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceServerCaCert(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceServerCaCert, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceOnPremisesConfiguration(c *Client, desired, actual *InstanceOnPremisesConfiguration) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.HostPort, actual.HostPort) && !dcl.IsZeroValue(desired.HostPort) {
		c.Config.Logger.Infof("Diff in HostPort.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.HostPort), dcl.SprintResource(actual.HostPort))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	if !dcl.StringCanonicalize(desired.Username, actual.Username) && !dcl.IsZeroValue(desired.Username) {
		c.Config.Logger.Infof("Diff in Username.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Username), dcl.SprintResource(actual.Username))
		return true
	}
	if !dcl.StringCanonicalize(desired.Password, actual.Password) && !dcl.IsZeroValue(desired.Password) {
		c.Config.Logger.Infof("Diff in Password.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Password), dcl.SprintResource(actual.Password))
		return true
	}
	if !dcl.StringCanonicalize(desired.CaCertificate, actual.CaCertificate) && !dcl.IsZeroValue(desired.CaCertificate) {
		c.Config.Logger.Infof("Diff in CaCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.CaCertificate), dcl.SprintResource(actual.CaCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClientCertificate, actual.ClientCertificate) && !dcl.IsZeroValue(desired.ClientCertificate) {
		c.Config.Logger.Infof("Diff in ClientCertificate.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientCertificate), dcl.SprintResource(actual.ClientCertificate))
		return true
	}
	if !dcl.StringCanonicalize(desired.ClientKey, actual.ClientKey) && !dcl.IsZeroValue(desired.ClientKey) {
		c.Config.Logger.Infof("Diff in ClientKey.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ClientKey), dcl.SprintResource(actual.ClientKey))
		return true
	}
	if !dcl.StringCanonicalize(desired.DumpFilePath, actual.DumpFilePath) && !dcl.IsZeroValue(desired.DumpFilePath) {
		c.Config.Logger.Infof("Diff in DumpFilePath.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.DumpFilePath), dcl.SprintResource(actual.DumpFilePath))
		return true
	}
	if !dcl.StringCanonicalize(desired.Database, actual.Database) && !dcl.IsZeroValue(desired.Database) {
		c.Config.Logger.Infof("Diff in Database.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Database), dcl.SprintResource(actual.Database))
		return true
	}
	if !dcl.StringSliceEquals(desired.ReplicatedDatabases, actual.ReplicatedDatabases) && !dcl.IsZeroValue(desired.ReplicatedDatabases) {
		c.Config.Logger.Infof("Diff in ReplicatedDatabases.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.ReplicatedDatabases), dcl.SprintResource(actual.ReplicatedDatabases))
		return true
	}
	return false
}

func compareInstanceOnPremisesConfigurationSlice(c *Client, desired, actual []InstanceOnPremisesConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceOnPremisesConfiguration, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceOnPremisesConfiguration(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceOnPremisesConfiguration, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceOnPremisesConfigurationMap(c *Client, desired, actual map[string]InstanceOnPremisesConfiguration) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceOnPremisesConfiguration, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceOnPremisesConfiguration, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceOnPremisesConfiguration(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceOnPremisesConfiguration, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceDiskEncryptionStatus(c *Client, desired, actual *InstanceDiskEncryptionStatus) bool {
	if desired == nil {
		return false
	}
	if actual == nil {
		return true
	}
	if !dcl.StringCanonicalize(desired.KmsKeyVersionName, actual.KmsKeyVersionName) && !dcl.IsZeroValue(desired.KmsKeyVersionName) {
		c.Config.Logger.Infof("Diff in KmsKeyVersionName.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.KmsKeyVersionName), dcl.SprintResource(actual.KmsKeyVersionName))
		return true
	}
	if !dcl.StringCanonicalize(desired.Kind, actual.Kind) && !dcl.IsZeroValue(desired.Kind) {
		c.Config.Logger.Infof("Diff in Kind.\nDESIRED: %s\nACTUAL: %s\n", dcl.SprintResource(desired.Kind), dcl.SprintResource(actual.Kind))
		return true
	}
	return false
}

func compareInstanceDiskEncryptionStatusSlice(c *Client, desired, actual []InstanceDiskEncryptionStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDiskEncryptionStatus, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDiskEncryptionStatus(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionStatus, element %d.\nDESIRED: %s\nACTUAL: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDiskEncryptionStatusMap(c *Client, desired, actual map[string]InstanceDiskEncryptionStatus) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDiskEncryptionStatus, lengths unequal.")
		return true
	}
	for k, desiredValue := range desired {
		actualValue, ok := actual[k]
		if !ok {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionStatus, key %s not found in ACTUAL.\n", k)
			return true
		}
		if compareInstanceDiskEncryptionStatus(c, &desiredValue, &actualValue) {
			c.Config.Logger.Infof("Diff in InstanceDiskEncryptionStatus, key %s.\nDESIRED: %s\nACTUAL: %s\n", k, dcl.SprintResource(desiredValue), dcl.SprintResource(actualValue))
			return true
		}
	}
	return false
}

func compareInstanceBackendTypeEnumSlice(c *Client, desired, actual []InstanceBackendTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceBackendTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceBackendTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceBackendTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceBackendTypeEnum(c *Client, desired, actual *InstanceBackendTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceDatabaseVersionEnumSlice(c *Client, desired, actual []InstanceDatabaseVersionEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceDatabaseVersionEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceDatabaseVersionEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceDatabaseVersionEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceDatabaseVersionEnum(c *Client, desired, actual *InstanceDatabaseVersionEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceInstanceTypeEnumSlice(c *Client, desired, actual []InstanceInstanceTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceInstanceTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceInstanceTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceInstanceTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceInstanceTypeEnum(c *Client, desired, actual *InstanceInstanceTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceIPAddressesTypeEnumSlice(c *Client, desired, actual []InstanceIPAddressesTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceIPAddressesTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceIPAddressesTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceIPAddressesTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceIPAddressesTypeEnum(c *Client, desired, actual *InstanceIPAddressesTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsAvailabilityTypeEnumSlice(c *Client, desired, actual []InstanceSettingsAvailabilityTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsAvailabilityTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsAvailabilityTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsAvailabilityTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsAvailabilityTypeEnum(c *Client, desired, actual *InstanceSettingsAvailabilityTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsPricingPlanEnumSlice(c *Client, desired, actual []InstanceSettingsPricingPlanEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsPricingPlanEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsPricingPlanEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsPricingPlanEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsPricingPlanEnum(c *Client, desired, actual *InstanceSettingsPricingPlanEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsReplicationTypeEnumSlice(c *Client, desired, actual []InstanceSettingsReplicationTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsReplicationTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsReplicationTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsReplicationTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsReplicationTypeEnum(c *Client, desired, actual *InstanceSettingsReplicationTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsActivationPolicyEnumSlice(c *Client, desired, actual []InstanceSettingsActivationPolicyEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsActivationPolicyEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsActivationPolicyEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsActivationPolicyEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsActivationPolicyEnum(c *Client, desired, actual *InstanceSettingsActivationPolicyEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsDataDiskTypeEnumSlice(c *Client, desired, actual []InstanceSettingsDataDiskTypeEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsDataDiskTypeEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsDataDiskTypeEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsDataDiskTypeEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsDataDiskTypeEnum(c *Client, desired, actual *InstanceSettingsDataDiskTypeEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsMaintenanceWindowUpdateTrackEnumSlice(c *Client, desired, actual []InstanceSettingsMaintenanceWindowUpdateTrackEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsMaintenanceWindowUpdateTrackEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsMaintenanceWindowUpdateTrackEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsMaintenanceWindowUpdateTrackEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsMaintenanceWindowUpdateTrackEnum(c *Client, desired, actual *InstanceSettingsMaintenanceWindowUpdateTrackEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumSlice(c *Client, desired, actual []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) bool {
	if len(desired) != len(actual) {
		c.Config.Logger.Info("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum, lengths unequal.")
		return true
	}
	for i := 0; i < len(desired); i++ {
		if compareInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(c, &desired[i], &actual[i]) {
			c.Config.Logger.Infof("Diff in InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum, element %d.\nOLD: %s\nNEW: %s\n", i, dcl.SprintResource(desired[i]), dcl.SprintResource(actual[i]))
			return true
		}
	}
	return false
}

func compareInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(c *Client, desired, actual *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum) bool {
	return !reflect.DeepEqual(desired, actual)
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Instance) urlNormalized() *Instance {
	normalized := dcl.Copy(*r).(Instance)
	normalized.ConnectionName = dcl.SelfLinkToName(r.ConnectionName)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.GceZone = dcl.SelfLinkToName(r.GceZone)
	normalized.MasterInstanceName = dcl.SelfLinkToName(r.MasterInstanceName)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Region = dcl.SelfLinkToName(r.Region)
	normalized.RootPassword = dcl.SelfLinkToName(r.RootPassword)
	normalized.State = dcl.SelfLinkToName(r.State)
	normalized.IPv6Address = dcl.SelfLinkToName(r.IPv6Address)
	normalized.ServiceAccountEmailAddress = dcl.SelfLinkToName(r.ServiceAccountEmailAddress)
	normalized.InstanceUid = dcl.SelfLinkToName(r.InstanceUid)
	return &normalized
}

func (r *Instance) getFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) createFields() string {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Instance) deleteFields() (string, string) {
	n := r.urlNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Instance) updateURL(userBasePath, updateName string) (string, error) {
	n := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/instances/{{name}}", "https://www.googleapis.com/sql/v1beta4/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Instance resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Instance) marshal(c *Client) ([]byte, error) {
	m, err := expandInstance(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Instance: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalInstance decodes JSON responses into the Instance resource schema.
func unmarshalInstance(b []byte, c *Client) (*Instance, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInstance(m, c)
}

func unmarshalMapInstance(m map[string]interface{}, c *Client) (*Instance, error) {

	return flattenInstance(c, m), nil
}

// expandInstance expands Instance into a JSON request object.
func expandInstance(c *Client, f *Instance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.BackendType; !dcl.IsEmptyValueIndirect(v) {
		m["backendType"] = v
	}
	if v := f.ConnectionName; !dcl.IsEmptyValueIndirect(v) {
		m["connectionName"] = v
	}
	if v := f.DatabaseVersion; !dcl.IsEmptyValueIndirect(v) {
		m["databaseVersion"] = v
	}
	if v := f.Etag; !dcl.IsEmptyValueIndirect(v) {
		m["etag"] = v
	}
	if v := f.GceZone; !dcl.IsEmptyValueIndirect(v) {
		m["gceZone"] = v
	}
	if v := f.InstanceType; !dcl.IsEmptyValueIndirect(v) {
		m["instanceType"] = v
	}
	if v := f.MasterInstanceName; !dcl.IsEmptyValueIndirect(v) {
		m["masterInstanceName"] = v
	}
	if v, err := expandInstanceMaxDiskSize(c, f.MaxDiskSize); err != nil {
		return nil, fmt.Errorf("error expanding MaxDiskSize into maxDiskSize: %w", err)
	} else if v != nil {
		m["maxDiskSize"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}
	if v := f.RootPassword; !dcl.IsEmptyValueIndirect(v) {
		m["rootPassword"] = v
	}
	if v, err := expandInstanceCurrentDiskSize(c, f.CurrentDiskSize); err != nil {
		return nil, fmt.Errorf("error expanding CurrentDiskSize into currentDiskSize: %w", err)
	} else if v != nil {
		m["currentDiskSize"] = v
	}
	if v, err := expandInstanceDiskEncryptionConfiguration(c, f.DiskEncryptionConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionConfiguration into diskEncryptionConfiguration: %w", err)
	} else if v != nil {
		m["diskEncryptionConfiguration"] = v
	}
	if v, err := expandInstanceFailoverReplica(c, f.FailoverReplica); err != nil {
		return nil, fmt.Errorf("error expanding FailoverReplica into failoverReplica: %w", err)
	} else if v != nil {
		m["failoverReplica"] = v
	}
	if v, err := expandInstanceIPAddressesSlice(c, f.IPAddresses); err != nil {
		return nil, fmt.Errorf("error expanding IPAddresses into ipAddresses: %w", err)
	} else if v != nil {
		m["ipAddresses"] = v
	}
	if v, err := expandInstanceMasterInstance(c, f.MasterInstance); err != nil {
		return nil, fmt.Errorf("error expanding MasterInstance into masterInstance: %w", err)
	} else if v != nil {
		m["masterInstance"] = v
	}
	if v, err := expandInstanceReplicaConfiguration(c, f.ReplicaConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaConfiguration into replicaConfiguration: %w", err)
	} else if v != nil {
		m["replicaConfiguration"] = v
	}
	if v, err := expandInstanceScheduledMaintenance(c, f.ScheduledMaintenance); err != nil {
		return nil, fmt.Errorf("error expanding ScheduledMaintenance into scheduledMaintenance: %w", err)
	} else if v != nil {
		m["scheduledMaintenance"] = v
	}
	if v, err := expandInstanceSettings(c, f.Settings); err != nil {
		return nil, fmt.Errorf("error expanding Settings into settings: %w", err)
	} else if v != nil {
		m["settings"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v, err := expandInstanceReplicaInstancesSlice(c, f.ReplicaInstances); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaInstances into replicaInstances: %w", err)
	} else if v != nil {
		m["replicaInstances"] = v
	}
	if v, err := expandInstanceServerCaCert(c, f.ServerCaCert); err != nil {
		return nil, fmt.Errorf("error expanding ServerCaCert into serverCaCert: %w", err)
	} else if v != nil {
		m["serverCaCert"] = v
	}
	if v := f.IPv6Address; !dcl.IsEmptyValueIndirect(v) {
		m["ipv6Address"] = v
	}
	if v := f.ServiceAccountEmailAddress; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccountEmailAddress"] = v
	}
	if v, err := expandInstanceOnPremisesConfiguration(c, f.OnPremisesConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding OnPremisesConfiguration into onPremisesConfiguration: %w", err)
	} else if v != nil {
		m["onPremisesConfiguration"] = v
	}
	if v := f.SuspensionReason; !dcl.IsEmptyValueIndirect(v) {
		m["suspensionReason"] = v
	}
	if v, err := expandInstanceDiskEncryptionStatus(c, f.DiskEncryptionStatus); err != nil {
		return nil, fmt.Errorf("error expanding DiskEncryptionStatus into diskEncryptionStatus: %w", err)
	} else if v != nil {
		m["diskEncryptionStatus"] = v
	}
	if v := f.InstanceUid; !dcl.IsEmptyValueIndirect(v) {
		m["instanceUid"] = v
	}

	return m, nil
}

// flattenInstance flattens Instance from a JSON request object into the
// Instance type.
func flattenInstance(c *Client, i interface{}) *Instance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	r := &Instance{}
	r.BackendType = flattenInstanceBackendTypeEnum(m["backendType"])
	r.ConnectionName = dcl.FlattenString(m["connectionName"])
	r.DatabaseVersion = flattenInstanceDatabaseVersionEnum(m["databaseVersion"])
	r.Etag = dcl.FlattenString(m["etag"])
	r.GceZone = dcl.FlattenString(m["gceZone"])
	r.InstanceType = flattenInstanceInstanceTypeEnum(m["instanceType"])
	r.MasterInstanceName = dcl.FlattenString(m["masterInstanceName"])
	r.MaxDiskSize = flattenInstanceMaxDiskSize(c, m["maxDiskSize"])
	r.Name = dcl.FlattenString(m["name"])
	r.Project = dcl.FlattenString(m["project"])
	r.Region = dcl.FlattenString(m["region"])
	r.RootPassword = dcl.FlattenString(m["rootPassword"])
	r.CurrentDiskSize = flattenInstanceCurrentDiskSize(c, m["currentDiskSize"])
	r.DiskEncryptionConfiguration = flattenInstanceDiskEncryptionConfiguration(c, m["diskEncryptionConfiguration"])
	r.FailoverReplica = flattenInstanceFailoverReplica(c, m["failoverReplica"])
	r.IPAddresses = flattenInstanceIPAddressesSlice(c, m["ipAddresses"])
	r.MasterInstance = flattenInstanceMasterInstance(c, m["masterInstance"])
	r.ReplicaConfiguration = flattenInstanceReplicaConfiguration(c, m["replicaConfiguration"])
	r.ScheduledMaintenance = flattenInstanceScheduledMaintenance(c, m["scheduledMaintenance"])
	r.Settings = flattenInstanceSettings(c, m["settings"])
	r.State = dcl.FlattenString(m["state"])
	r.ReplicaInstances = flattenInstanceReplicaInstancesSlice(c, m["replicaInstances"])
	r.ServerCaCert = flattenInstanceServerCaCert(c, m["serverCaCert"])
	r.IPv6Address = dcl.FlattenString(m["ipv6Address"])
	r.ServiceAccountEmailAddress = dcl.FlattenString(m["serviceAccountEmailAddress"])
	r.OnPremisesConfiguration = flattenInstanceOnPremisesConfiguration(c, m["onPremisesConfiguration"])
	r.SuspensionReason = dcl.FlattenStringSlice(m["suspensionReason"])
	r.DiskEncryptionStatus = flattenInstanceDiskEncryptionStatus(c, m["diskEncryptionStatus"])
	r.InstanceUid = dcl.FlattenString(m["instanceUid"])

	return r
}

// expandInstanceMaxDiskSizeMap expands the contents of InstanceMaxDiskSize into a JSON
// request object.
func expandInstanceMaxDiskSizeMap(c *Client, f map[string]InstanceMaxDiskSize) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMaxDiskSize(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMaxDiskSizeSlice expands the contents of InstanceMaxDiskSize into a JSON
// request object.
func expandInstanceMaxDiskSizeSlice(c *Client, f []InstanceMaxDiskSize) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMaxDiskSize(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMaxDiskSizeMap flattens the contents of InstanceMaxDiskSize from a JSON
// response object.
func flattenInstanceMaxDiskSizeMap(c *Client, i interface{}) map[string]InstanceMaxDiskSize {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMaxDiskSize{}
	}

	if len(a) == 0 {
		return map[string]InstanceMaxDiskSize{}
	}

	items := make(map[string]InstanceMaxDiskSize)
	for k, item := range a {
		items[k] = *flattenInstanceMaxDiskSize(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMaxDiskSizeSlice flattens the contents of InstanceMaxDiskSize from a JSON
// response object.
func flattenInstanceMaxDiskSizeSlice(c *Client, i interface{}) []InstanceMaxDiskSize {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMaxDiskSize{}
	}

	if len(a) == 0 {
		return []InstanceMaxDiskSize{}
	}

	items := make([]InstanceMaxDiskSize, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMaxDiskSize(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMaxDiskSize expands an instance of InstanceMaxDiskSize into a JSON
// request object.
func expandInstanceMaxDiskSize(c *Client, f *InstanceMaxDiskSize) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceMaxDiskSize flattens an instance of InstanceMaxDiskSize from a JSON
// response object.
func flattenInstanceMaxDiskSize(c *Client, i interface{}) *InstanceMaxDiskSize {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMaxDiskSize{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceCurrentDiskSizeMap expands the contents of InstanceCurrentDiskSize into a JSON
// request object.
func expandInstanceCurrentDiskSizeMap(c *Client, f map[string]InstanceCurrentDiskSize) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceCurrentDiskSize(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceCurrentDiskSizeSlice expands the contents of InstanceCurrentDiskSize into a JSON
// request object.
func expandInstanceCurrentDiskSizeSlice(c *Client, f []InstanceCurrentDiskSize) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceCurrentDiskSize(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceCurrentDiskSizeMap flattens the contents of InstanceCurrentDiskSize from a JSON
// response object.
func flattenInstanceCurrentDiskSizeMap(c *Client, i interface{}) map[string]InstanceCurrentDiskSize {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceCurrentDiskSize{}
	}

	if len(a) == 0 {
		return map[string]InstanceCurrentDiskSize{}
	}

	items := make(map[string]InstanceCurrentDiskSize)
	for k, item := range a {
		items[k] = *flattenInstanceCurrentDiskSize(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceCurrentDiskSizeSlice flattens the contents of InstanceCurrentDiskSize from a JSON
// response object.
func flattenInstanceCurrentDiskSizeSlice(c *Client, i interface{}) []InstanceCurrentDiskSize {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceCurrentDiskSize{}
	}

	if len(a) == 0 {
		return []InstanceCurrentDiskSize{}
	}

	items := make([]InstanceCurrentDiskSize, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceCurrentDiskSize(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceCurrentDiskSize expands an instance of InstanceCurrentDiskSize into a JSON
// request object.
func expandInstanceCurrentDiskSize(c *Client, f *InstanceCurrentDiskSize) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceCurrentDiskSize flattens an instance of InstanceCurrentDiskSize from a JSON
// response object.
func flattenInstanceCurrentDiskSize(c *Client, i interface{}) *InstanceCurrentDiskSize {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceCurrentDiskSize{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceDiskEncryptionConfigurationMap expands the contents of InstanceDiskEncryptionConfiguration into a JSON
// request object.
func expandInstanceDiskEncryptionConfigurationMap(c *Client, f map[string]InstanceDiskEncryptionConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDiskEncryptionConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDiskEncryptionConfigurationSlice expands the contents of InstanceDiskEncryptionConfiguration into a JSON
// request object.
func expandInstanceDiskEncryptionConfigurationSlice(c *Client, f []InstanceDiskEncryptionConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDiskEncryptionConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDiskEncryptionConfigurationMap flattens the contents of InstanceDiskEncryptionConfiguration from a JSON
// response object.
func flattenInstanceDiskEncryptionConfigurationMap(c *Client, i interface{}) map[string]InstanceDiskEncryptionConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDiskEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceDiskEncryptionConfiguration{}
	}

	items := make(map[string]InstanceDiskEncryptionConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceDiskEncryptionConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDiskEncryptionConfigurationSlice flattens the contents of InstanceDiskEncryptionConfiguration from a JSON
// response object.
func flattenInstanceDiskEncryptionConfigurationSlice(c *Client, i interface{}) []InstanceDiskEncryptionConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDiskEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceDiskEncryptionConfiguration{}
	}

	items := make([]InstanceDiskEncryptionConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDiskEncryptionConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDiskEncryptionConfiguration expands an instance of InstanceDiskEncryptionConfiguration into a JSON
// request object.
func expandInstanceDiskEncryptionConfiguration(c *Client, f *InstanceDiskEncryptionConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceDiskEncryptionConfiguration flattens an instance of InstanceDiskEncryptionConfiguration from a JSON
// response object.
func flattenInstanceDiskEncryptionConfiguration(c *Client, i interface{}) *InstanceDiskEncryptionConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDiskEncryptionConfiguration{}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceFailoverReplicaMap expands the contents of InstanceFailoverReplica into a JSON
// request object.
func expandInstanceFailoverReplicaMap(c *Client, f map[string]InstanceFailoverReplica) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFailoverReplica(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFailoverReplicaSlice expands the contents of InstanceFailoverReplica into a JSON
// request object.
func expandInstanceFailoverReplicaSlice(c *Client, f []InstanceFailoverReplica) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFailoverReplica(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFailoverReplicaMap flattens the contents of InstanceFailoverReplica from a JSON
// response object.
func flattenInstanceFailoverReplicaMap(c *Client, i interface{}) map[string]InstanceFailoverReplica {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFailoverReplica{}
	}

	if len(a) == 0 {
		return map[string]InstanceFailoverReplica{}
	}

	items := make(map[string]InstanceFailoverReplica)
	for k, item := range a {
		items[k] = *flattenInstanceFailoverReplica(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceFailoverReplicaSlice flattens the contents of InstanceFailoverReplica from a JSON
// response object.
func flattenInstanceFailoverReplicaSlice(c *Client, i interface{}) []InstanceFailoverReplica {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFailoverReplica{}
	}

	if len(a) == 0 {
		return []InstanceFailoverReplica{}
	}

	items := make([]InstanceFailoverReplica, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFailoverReplica(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceFailoverReplica expands an instance of InstanceFailoverReplica into a JSON
// request object.
func expandInstanceFailoverReplica(c *Client, f *InstanceFailoverReplica) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Available; !dcl.IsEmptyValueIndirect(v) {
		m["available"] = v
	}
	if v, err := expandInstanceFailoverReplicaFailoverInstance(c, f.FailoverInstance); err != nil {
		return nil, fmt.Errorf("error expanding FailoverInstance into failoverInstance: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["failoverInstance"] = v
	}

	return m, nil
}

// flattenInstanceFailoverReplica flattens an instance of InstanceFailoverReplica from a JSON
// response object.
func flattenInstanceFailoverReplica(c *Client, i interface{}) *InstanceFailoverReplica {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFailoverReplica{}
	r.Name = dcl.FlattenString(m["name"])
	r.Available = dcl.FlattenBool(m["available"])
	r.FailoverInstance = flattenInstanceFailoverReplicaFailoverInstance(c, m["failoverInstance"])

	return r
}

// expandInstanceFailoverReplicaFailoverInstanceMap expands the contents of InstanceFailoverReplicaFailoverInstance into a JSON
// request object.
func expandInstanceFailoverReplicaFailoverInstanceMap(c *Client, f map[string]InstanceFailoverReplicaFailoverInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceFailoverReplicaFailoverInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceFailoverReplicaFailoverInstanceSlice expands the contents of InstanceFailoverReplicaFailoverInstance into a JSON
// request object.
func expandInstanceFailoverReplicaFailoverInstanceSlice(c *Client, f []InstanceFailoverReplicaFailoverInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceFailoverReplicaFailoverInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceFailoverReplicaFailoverInstanceMap flattens the contents of InstanceFailoverReplicaFailoverInstance from a JSON
// response object.
func flattenInstanceFailoverReplicaFailoverInstanceMap(c *Client, i interface{}) map[string]InstanceFailoverReplicaFailoverInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceFailoverReplicaFailoverInstance{}
	}

	if len(a) == 0 {
		return map[string]InstanceFailoverReplicaFailoverInstance{}
	}

	items := make(map[string]InstanceFailoverReplicaFailoverInstance)
	for k, item := range a {
		items[k] = *flattenInstanceFailoverReplicaFailoverInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceFailoverReplicaFailoverInstanceSlice flattens the contents of InstanceFailoverReplicaFailoverInstance from a JSON
// response object.
func flattenInstanceFailoverReplicaFailoverInstanceSlice(c *Client, i interface{}) []InstanceFailoverReplicaFailoverInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceFailoverReplicaFailoverInstance{}
	}

	if len(a) == 0 {
		return []InstanceFailoverReplicaFailoverInstance{}
	}

	items := make([]InstanceFailoverReplicaFailoverInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceFailoverReplicaFailoverInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceFailoverReplicaFailoverInstance expands an instance of InstanceFailoverReplicaFailoverInstance into a JSON
// request object.
func expandInstanceFailoverReplicaFailoverInstance(c *Client, f *InstanceFailoverReplicaFailoverInstance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenInstanceFailoverReplicaFailoverInstance flattens an instance of InstanceFailoverReplicaFailoverInstance from a JSON
// response object.
func flattenInstanceFailoverReplicaFailoverInstance(c *Client, i interface{}) *InstanceFailoverReplicaFailoverInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceFailoverReplicaFailoverInstance{}
	r.Name = dcl.FlattenString(m["name"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceIPAddressesMap expands the contents of InstanceIPAddresses into a JSON
// request object.
func expandInstanceIPAddressesMap(c *Client, f map[string]InstanceIPAddresses) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceIPAddresses(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceIPAddressesSlice expands the contents of InstanceIPAddresses into a JSON
// request object.
func expandInstanceIPAddressesSlice(c *Client, f []InstanceIPAddresses) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceIPAddresses(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceIPAddressesMap flattens the contents of InstanceIPAddresses from a JSON
// response object.
func flattenInstanceIPAddressesMap(c *Client, i interface{}) map[string]InstanceIPAddresses {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceIPAddresses{}
	}

	if len(a) == 0 {
		return map[string]InstanceIPAddresses{}
	}

	items := make(map[string]InstanceIPAddresses)
	for k, item := range a {
		items[k] = *flattenInstanceIPAddresses(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceIPAddressesSlice flattens the contents of InstanceIPAddresses from a JSON
// response object.
func flattenInstanceIPAddressesSlice(c *Client, i interface{}) []InstanceIPAddresses {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceIPAddresses{}
	}

	if len(a) == 0 {
		return []InstanceIPAddresses{}
	}

	items := make([]InstanceIPAddresses, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceIPAddresses(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceIPAddresses expands an instance of InstanceIPAddresses into a JSON
// request object.
func expandInstanceIPAddresses(c *Client, f *InstanceIPAddresses) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v := f.IPAddress; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddress"] = v
	}
	if v, err := expandInstanceIPAddressesTimeToRetire(c, f.TimeToRetire); err != nil {
		return nil, fmt.Errorf("error expanding TimeToRetire into timeToRetire: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timeToRetire"] = v
	}

	return m, nil
}

// flattenInstanceIPAddresses flattens an instance of InstanceIPAddresses from a JSON
// response object.
func flattenInstanceIPAddresses(c *Client, i interface{}) *InstanceIPAddresses {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceIPAddresses{}
	r.Type = flattenInstanceIPAddressesTypeEnum(m["type"])
	r.IPAddress = dcl.FlattenString(m["ipAddress"])
	r.TimeToRetire = flattenInstanceIPAddressesTimeToRetire(c, m["timeToRetire"])

	return r
}

// expandInstanceIPAddressesTimeToRetireMap expands the contents of InstanceIPAddressesTimeToRetire into a JSON
// request object.
func expandInstanceIPAddressesTimeToRetireMap(c *Client, f map[string]InstanceIPAddressesTimeToRetire) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceIPAddressesTimeToRetire(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceIPAddressesTimeToRetireSlice expands the contents of InstanceIPAddressesTimeToRetire into a JSON
// request object.
func expandInstanceIPAddressesTimeToRetireSlice(c *Client, f []InstanceIPAddressesTimeToRetire) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceIPAddressesTimeToRetire(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceIPAddressesTimeToRetireMap flattens the contents of InstanceIPAddressesTimeToRetire from a JSON
// response object.
func flattenInstanceIPAddressesTimeToRetireMap(c *Client, i interface{}) map[string]InstanceIPAddressesTimeToRetire {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceIPAddressesTimeToRetire{}
	}

	if len(a) == 0 {
		return map[string]InstanceIPAddressesTimeToRetire{}
	}

	items := make(map[string]InstanceIPAddressesTimeToRetire)
	for k, item := range a {
		items[k] = *flattenInstanceIPAddressesTimeToRetire(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceIPAddressesTimeToRetireSlice flattens the contents of InstanceIPAddressesTimeToRetire from a JSON
// response object.
func flattenInstanceIPAddressesTimeToRetireSlice(c *Client, i interface{}) []InstanceIPAddressesTimeToRetire {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceIPAddressesTimeToRetire{}
	}

	if len(a) == 0 {
		return []InstanceIPAddressesTimeToRetire{}
	}

	items := make([]InstanceIPAddressesTimeToRetire, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceIPAddressesTimeToRetire(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceIPAddressesTimeToRetire expands an instance of InstanceIPAddressesTimeToRetire into a JSON
// request object.
func expandInstanceIPAddressesTimeToRetire(c *Client, f *InstanceIPAddressesTimeToRetire) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenInstanceIPAddressesTimeToRetire flattens an instance of InstanceIPAddressesTimeToRetire from a JSON
// response object.
func flattenInstanceIPAddressesTimeToRetire(c *Client, i interface{}) *InstanceIPAddressesTimeToRetire {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceIPAddressesTimeToRetire{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandInstanceMasterInstanceMap expands the contents of InstanceMasterInstance into a JSON
// request object.
func expandInstanceMasterInstanceMap(c *Client, f map[string]InstanceMasterInstance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceMasterInstance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceMasterInstanceSlice expands the contents of InstanceMasterInstance into a JSON
// request object.
func expandInstanceMasterInstanceSlice(c *Client, f []InstanceMasterInstance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceMasterInstance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceMasterInstanceMap flattens the contents of InstanceMasterInstance from a JSON
// response object.
func flattenInstanceMasterInstanceMap(c *Client, i interface{}) map[string]InstanceMasterInstance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceMasterInstance{}
	}

	if len(a) == 0 {
		return map[string]InstanceMasterInstance{}
	}

	items := make(map[string]InstanceMasterInstance)
	for k, item := range a {
		items[k] = *flattenInstanceMasterInstance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceMasterInstanceSlice flattens the contents of InstanceMasterInstance from a JSON
// response object.
func flattenInstanceMasterInstanceSlice(c *Client, i interface{}) []InstanceMasterInstance {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceMasterInstance{}
	}

	if len(a) == 0 {
		return []InstanceMasterInstance{}
	}

	items := make([]InstanceMasterInstance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceMasterInstance(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceMasterInstance expands an instance of InstanceMasterInstance into a JSON
// request object.
func expandInstanceMasterInstance(c *Client, f *InstanceMasterInstance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenInstanceMasterInstance flattens an instance of InstanceMasterInstance from a JSON
// response object.
func flattenInstanceMasterInstance(c *Client, i interface{}) *InstanceMasterInstance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceMasterInstance{}
	r.Name = dcl.FlattenString(m["name"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceReplicaConfigurationMap expands the contents of InstanceReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMap(c *Client, f map[string]InstanceReplicaConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationSlice expands the contents of InstanceReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationSlice(c *Client, f []InstanceReplicaConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationMap flattens the contents of InstanceReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationSlice flattens the contents of InstanceReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfiguration{}
	}

	items := make([]InstanceReplicaConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfiguration expands an instance of InstanceReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfiguration(c *Client, f *InstanceReplicaConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v, err := expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c, f.MysqlReplicaConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding MysqlReplicaConfiguration into mysqlReplicaConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mysqlReplicaConfiguration"] = v
	}
	if v := f.FailoverTarget; !dcl.IsEmptyValueIndirect(v) {
		m["failoverTarget"] = v
	}
	if v, err := expandInstanceReplicaConfigurationReplicaPoolConfiguration(c, f.ReplicaPoolConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding ReplicaPoolConfiguration into replicaPoolConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["replicaPoolConfiguration"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfiguration flattens an instance of InstanceReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfiguration(c *Client, i interface{}) *InstanceReplicaConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.MysqlReplicaConfiguration = flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c, m["mysqlReplicaConfiguration"])
	r.FailoverTarget = dcl.FlattenBool(m["failoverTarget"])
	r.ReplicaPoolConfiguration = flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c, m["replicaPoolConfiguration"])

	return r
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMap expands the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationSlice expands the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, f []InstanceReplicaConfigurationMysqlReplicaConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMap flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationMysqlReplicaConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationSlice flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationMysqlReplicaConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationMysqlReplicaConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationMysqlReplicaConfiguration expands an instance of InstanceReplicaConfigurationMysqlReplicaConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, f *InstanceReplicaConfigurationMysqlReplicaConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.DumpFilePath; !dcl.IsEmptyValueIndirect(v) {
		m["dumpFilePath"] = v
	}
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v := f.ConnectRetryInterval; !dcl.IsEmptyValueIndirect(v) {
		m["connectRetryInterval"] = v
	}
	if v, err := expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, f.MasterHeartbeatPeriod); err != nil {
		return nil, fmt.Errorf("error expanding MasterHeartbeatPeriod into masterHeartbeatPeriod: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["masterHeartbeatPeriod"] = v
	}
	if v := f.CaCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["caCertificate"] = v
	}
	if v := f.ClientCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["clientCertificate"] = v
	}
	if v := f.ClientKey; !dcl.IsEmptyValueIndirect(v) {
		m["clientKey"] = v
	}
	if v := f.SslCipher; !dcl.IsEmptyValueIndirect(v) {
		m["sslCipher"] = v
	}
	if v := f.VerifyServerCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["verifyServerCertificate"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfiguration flattens an instance of InstanceReplicaConfigurationMysqlReplicaConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationMysqlReplicaConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationMysqlReplicaConfiguration{}
	r.DumpFilePath = dcl.FlattenString(m["dumpFilePath"])
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])
	r.ConnectRetryInterval = dcl.FlattenInteger(m["connectRetryInterval"])
	r.MasterHeartbeatPeriod = flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, m["masterHeartbeatPeriod"])
	r.CaCertificate = dcl.FlattenString(m["caCertificate"])
	r.ClientCertificate = dcl.FlattenString(m["clientCertificate"])
	r.ClientKey = dcl.FlattenString(m["clientKey"])
	r.SslCipher = dcl.FlattenString(m["sslCipher"])
	r.VerifyServerCertificate = dcl.FlattenBool(m["verifyServerCertificate"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap expands the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap(c *Client, f map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice expands the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, f []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	items := make(map[string]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice flattens the contents of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriodSlice(c *Client, i interface{}) []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	}

	items := make([]InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod expands an instance of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod into a JSON
// request object.
func expandInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, f *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod flattens an instance of InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod from a JSON
// response object.
func flattenInstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod(c *Client, i interface{}) *InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationMysqlReplicaConfigurationMasterHeartbeatPeriod{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationMap expands the contents of InstanceReplicaConfigurationReplicaPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationReplicaPoolConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationSlice expands the contents of InstanceReplicaConfigurationReplicaPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, f []InstanceReplicaConfigurationReplicaPoolConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationMap flattens the contents of InstanceReplicaConfigurationReplicaPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationReplicaPoolConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationReplicaPoolConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationSlice flattens the contents of InstanceReplicaConfigurationReplicaPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationReplicaPoolConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationReplicaPoolConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationReplicaPoolConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationReplicaPoolConfiguration expands an instance of InstanceReplicaConfigurationReplicaPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, f *InstanceReplicaConfigurationReplicaPoolConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, f.StaticPoolConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding StaticPoolConfiguration into staticPoolConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["staticPoolConfiguration"] = v
	}
	if v, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, f.AutoscalingPoolConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding AutoscalingPoolConfiguration into autoscalingPoolConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["autoscalingPoolConfiguration"] = v
	}
	if v := f.ReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["replicaCount"] = v
	}
	if v := f.ExposeReplicaIP; !dcl.IsEmptyValueIndirect(v) {
		m["exposeReplicaIp"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfiguration flattens an instance of InstanceReplicaConfigurationReplicaPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationReplicaPoolConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationReplicaPoolConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.StaticPoolConfiguration = flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, m["staticPoolConfiguration"])
	r.AutoscalingPoolConfiguration = flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, m["autoscalingPoolConfiguration"])
	r.ReplicaCount = dcl.FlattenInteger(m["replicaCount"])
	r.ExposeReplicaIP = dcl.FlattenBool(m["exposeReplicaIp"])

	return r
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, f []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration expands an instance of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, f *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.ReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["replicaCount"] = v
	}
	if v := f.ExposeReplicaIP; !dcl.IsEmptyValueIndirect(v) {
		m["exposeReplicaIP"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration flattens an instance of InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationReplicaPoolConfigurationStaticPoolConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.ReplicaCount = dcl.FlattenInteger(m["replicaCount"])
	r.ExposeReplicaIP = dcl.FlattenBool(m["exposeReplicaIP"])

	return r
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap(c *Client, f map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice expands the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, f []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationMap(c *Client, i interface{}) map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	items := make(map[string]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice flattens the contents of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfigurationSlice(c *Client, i interface{}) []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	}

	items := make([]InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration expands an instance of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration into a JSON
// request object.
func expandInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, f *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.MinReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["minReplicaCount"] = v
	}
	if v := f.MaxReplicaCount; !dcl.IsEmptyValueIndirect(v) {
		m["maxReplicaCount"] = v
	}
	if v := f.TargetCpuUtil; !dcl.IsEmptyValueIndirect(v) {
		m["targetCpuUtil"] = v
	}

	return m, nil
}

// flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration flattens an instance of InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration from a JSON
// response object.
func flattenInstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration(c *Client, i interface{}) *InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaConfigurationReplicaPoolConfigurationAutoscalingPoolConfiguration{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.MinReplicaCount = dcl.FlattenInteger(m["minReplicaCount"])
	r.MaxReplicaCount = dcl.FlattenInteger(m["maxReplicaCount"])
	r.TargetCpuUtil = dcl.FlattenDouble(m["targetCpuUtil"])

	return r
}

// expandInstanceScheduledMaintenanceMap expands the contents of InstanceScheduledMaintenance into a JSON
// request object.
func expandInstanceScheduledMaintenanceMap(c *Client, f map[string]InstanceScheduledMaintenance) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceScheduledMaintenance(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceScheduledMaintenanceSlice expands the contents of InstanceScheduledMaintenance into a JSON
// request object.
func expandInstanceScheduledMaintenanceSlice(c *Client, f []InstanceScheduledMaintenance) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceScheduledMaintenance(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceScheduledMaintenanceMap flattens the contents of InstanceScheduledMaintenance from a JSON
// response object.
func flattenInstanceScheduledMaintenanceMap(c *Client, i interface{}) map[string]InstanceScheduledMaintenance {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceScheduledMaintenance{}
	}

	if len(a) == 0 {
		return map[string]InstanceScheduledMaintenance{}
	}

	items := make(map[string]InstanceScheduledMaintenance)
	for k, item := range a {
		items[k] = *flattenInstanceScheduledMaintenance(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceScheduledMaintenanceSlice flattens the contents of InstanceScheduledMaintenance from a JSON
// response object.
func flattenInstanceScheduledMaintenanceSlice(c *Client, i interface{}) []InstanceScheduledMaintenance {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceScheduledMaintenance{}
	}

	if len(a) == 0 {
		return []InstanceScheduledMaintenance{}
	}

	items := make([]InstanceScheduledMaintenance, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceScheduledMaintenance(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceScheduledMaintenance expands an instance of InstanceScheduledMaintenance into a JSON
// request object.
func expandInstanceScheduledMaintenance(c *Client, f *InstanceScheduledMaintenance) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v, err := expandInstanceScheduledMaintenanceStartTime(c, f.StartTime); err != nil {
		return nil, fmt.Errorf("error expanding StartTime into startTime: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.CanDefer; !dcl.IsEmptyValueIndirect(v) {
		m["canDefer"] = v
	}
	if v := f.CanReschedule; !dcl.IsEmptyValueIndirect(v) {
		m["canReschedule"] = v
	}

	return m, nil
}

// flattenInstanceScheduledMaintenance flattens an instance of InstanceScheduledMaintenance from a JSON
// response object.
func flattenInstanceScheduledMaintenance(c *Client, i interface{}) *InstanceScheduledMaintenance {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceScheduledMaintenance{}
	r.StartTime = flattenInstanceScheduledMaintenanceStartTime(c, m["startTime"])
	r.CanDefer = dcl.FlattenBool(m["canDefer"])
	r.CanReschedule = dcl.FlattenBool(m["canReschedule"])

	return r
}

// expandInstanceScheduledMaintenanceStartTimeMap expands the contents of InstanceScheduledMaintenanceStartTime into a JSON
// request object.
func expandInstanceScheduledMaintenanceStartTimeMap(c *Client, f map[string]InstanceScheduledMaintenanceStartTime) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceScheduledMaintenanceStartTime(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceScheduledMaintenanceStartTimeSlice expands the contents of InstanceScheduledMaintenanceStartTime into a JSON
// request object.
func expandInstanceScheduledMaintenanceStartTimeSlice(c *Client, f []InstanceScheduledMaintenanceStartTime) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceScheduledMaintenanceStartTime(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceScheduledMaintenanceStartTimeMap flattens the contents of InstanceScheduledMaintenanceStartTime from a JSON
// response object.
func flattenInstanceScheduledMaintenanceStartTimeMap(c *Client, i interface{}) map[string]InstanceScheduledMaintenanceStartTime {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceScheduledMaintenanceStartTime{}
	}

	if len(a) == 0 {
		return map[string]InstanceScheduledMaintenanceStartTime{}
	}

	items := make(map[string]InstanceScheduledMaintenanceStartTime)
	for k, item := range a {
		items[k] = *flattenInstanceScheduledMaintenanceStartTime(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceScheduledMaintenanceStartTimeSlice flattens the contents of InstanceScheduledMaintenanceStartTime from a JSON
// response object.
func flattenInstanceScheduledMaintenanceStartTimeSlice(c *Client, i interface{}) []InstanceScheduledMaintenanceStartTime {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceScheduledMaintenanceStartTime{}
	}

	if len(a) == 0 {
		return []InstanceScheduledMaintenanceStartTime{}
	}

	items := make([]InstanceScheduledMaintenanceStartTime, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceScheduledMaintenanceStartTime(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceScheduledMaintenanceStartTime expands an instance of InstanceScheduledMaintenanceStartTime into a JSON
// request object.
func expandInstanceScheduledMaintenanceStartTime(c *Client, f *InstanceScheduledMaintenanceStartTime) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Seconds; !dcl.IsEmptyValueIndirect(v) {
		m["seconds"] = v
	}
	if v := f.Nanos; !dcl.IsEmptyValueIndirect(v) {
		m["nanos"] = v
	}

	return m, nil
}

// flattenInstanceScheduledMaintenanceStartTime flattens an instance of InstanceScheduledMaintenanceStartTime from a JSON
// response object.
func flattenInstanceScheduledMaintenanceStartTime(c *Client, i interface{}) *InstanceScheduledMaintenanceStartTime {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceScheduledMaintenanceStartTime{}
	r.Seconds = dcl.FlattenInteger(m["seconds"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandInstanceSettingsMap expands the contents of InstanceSettings into a JSON
// request object.
func expandInstanceSettingsMap(c *Client, f map[string]InstanceSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsSlice expands the contents of InstanceSettings into a JSON
// request object.
func expandInstanceSettingsSlice(c *Client, f []InstanceSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsMap flattens the contents of InstanceSettings from a JSON
// response object.
func flattenInstanceSettingsMap(c *Client, i interface{}) map[string]InstanceSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettings{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettings{}
	}

	items := make(map[string]InstanceSettings)
	for k, item := range a {
		items[k] = *flattenInstanceSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsSlice flattens the contents of InstanceSettings from a JSON
// response object.
func flattenInstanceSettingsSlice(c *Client, i interface{}) []InstanceSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettings{}
	}

	if len(a) == 0 {
		return []InstanceSettings{}
	}

	items := make([]InstanceSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettings expands an instance of InstanceSettings into a JSON
// request object.
func expandInstanceSettings(c *Client, f *InstanceSettings) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.AuthorizedGaeApplications; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedGaeApplications"] = v
	}
	if v := f.Tier; !dcl.IsEmptyValueIndirect(v) {
		m["tier"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.AvailabilityType; !dcl.IsEmptyValueIndirect(v) {
		m["availabilityType"] = v
	}
	if v := f.PricingPlan; !dcl.IsEmptyValueIndirect(v) {
		m["pricingPlan"] = v
	}
	if v := f.ReplicationType; !dcl.IsEmptyValueIndirect(v) {
		m["replicationType"] = v
	}
	if v := f.ActivationPolicy; !dcl.IsEmptyValueIndirect(v) {
		m["activationPolicy"] = v
	}
	if v := f.StorageAutoResize; !dcl.IsEmptyValueIndirect(v) {
		m["storageAutoResize"] = v
	}
	if v := f.DataDiskType; !dcl.IsEmptyValueIndirect(v) {
		m["dataDiskType"] = v
	}
	if v := f.DatabaseReplicationEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["databaseReplicationEnabled"] = v
	}
	if v := f.CrashSafeReplicationEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["crashSafeReplicationEnabled"] = v
	}
	if v, err := expandInstanceSettingsSettingsVersion(c, f.SettingsVersion); err != nil {
		return nil, fmt.Errorf("error expanding SettingsVersion into settingsVersion: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["settingsVersion"] = v
	}
	if v := f.UserLabels; !dcl.IsEmptyValueIndirect(v) {
		m["userLabels"] = v
	}
	if v, err := expandInstanceSettingsStorageAutoResizeLimit(c, f.StorageAutoResizeLimit); err != nil {
		return nil, fmt.Errorf("error expanding StorageAutoResizeLimit into storageAutoResizeLimit: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["storageAutoResizeLimit"] = v
	}
	if v, err := expandInstanceSettingsIPConfiguration(c, f.IPConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding IPConfiguration into ipConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["ipConfiguration"] = v
	}
	if v, err := expandInstanceSettingsLocationPreference(c, f.LocationPreference); err != nil {
		return nil, fmt.Errorf("error expanding LocationPreference into locationPreference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["locationPreference"] = v
	}
	if v, err := expandInstanceSettingsDatabaseFlagsSlice(c, f.DatabaseFlags); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseFlags into databaseFlags: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseFlags"] = v
	}
	if v, err := expandInstanceSettingsMaintenanceWindow(c, f.MaintenanceWindow); err != nil {
		return nil, fmt.Errorf("error expanding MaintenanceWindow into maintenanceWindow: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["maintenanceWindow"] = v
	}
	if v, err := expandInstanceSettingsBackupConfiguration(c, f.BackupConfiguration); err != nil {
		return nil, fmt.Errorf("error expanding BackupConfiguration into backupConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["backupConfiguration"] = v
	}
	if v, err := expandInstanceSettingsDataDiskSizeGb(c, f.DataDiskSizeGb); err != nil {
		return nil, fmt.Errorf("error expanding DataDiskSizeGb into dataDiskSizeGb: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dataDiskSizeGb"] = v
	}
	if v, err := expandInstanceSettingsActiveDirectoryConfig(c, f.ActiveDirectoryConfig); err != nil {
		return nil, fmt.Errorf("error expanding ActiveDirectoryConfig into activeDirectoryConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["activeDirectoryConfig"] = v
	}
	if v := f.Collation; !dcl.IsEmptyValueIndirect(v) {
		m["collation"] = v
	}
	if v, err := expandInstanceSettingsDenyMaintenancePeriodsSlice(c, f.DenyMaintenancePeriods); err != nil {
		return nil, fmt.Errorf("error expanding DenyMaintenancePeriods into denyMaintenancePeriods: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["denyMaintenancePeriods"] = v
	}
	if v, err := expandInstanceSettingsInsightsConfig(c, f.InsightsConfig); err != nil {
		return nil, fmt.Errorf("error expanding InsightsConfig into insightsConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["insightsConfig"] = v
	}

	return m, nil
}

// flattenInstanceSettings flattens an instance of InstanceSettings from a JSON
// response object.
func flattenInstanceSettings(c *Client, i interface{}) *InstanceSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettings{}
	r.AuthorizedGaeApplications = dcl.FlattenStringSlice(m["authorizedGaeApplications"])
	r.Tier = dcl.FlattenString(m["tier"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.AvailabilityType = flattenInstanceSettingsAvailabilityTypeEnum(m["availabilityType"])
	r.PricingPlan = flattenInstanceSettingsPricingPlanEnum(m["pricingPlan"])
	r.ReplicationType = flattenInstanceSettingsReplicationTypeEnum(m["replicationType"])
	r.ActivationPolicy = flattenInstanceSettingsActivationPolicyEnum(m["activationPolicy"])
	r.StorageAutoResize = dcl.FlattenBool(m["storageAutoResize"])
	r.DataDiskType = flattenInstanceSettingsDataDiskTypeEnum(m["dataDiskType"])
	r.DatabaseReplicationEnabled = dcl.FlattenBool(m["databaseReplicationEnabled"])
	r.CrashSafeReplicationEnabled = dcl.FlattenBool(m["crashSafeReplicationEnabled"])
	r.SettingsVersion = flattenInstanceSettingsSettingsVersion(c, m["settingsVersion"])
	r.UserLabels = dcl.FlattenKeyValuePairs(m["userLabels"])
	r.StorageAutoResizeLimit = flattenInstanceSettingsStorageAutoResizeLimit(c, m["storageAutoResizeLimit"])
	r.IPConfiguration = flattenInstanceSettingsIPConfiguration(c, m["ipConfiguration"])
	r.LocationPreference = flattenInstanceSettingsLocationPreference(c, m["locationPreference"])
	r.DatabaseFlags = flattenInstanceSettingsDatabaseFlagsSlice(c, m["databaseFlags"])
	r.MaintenanceWindow = flattenInstanceSettingsMaintenanceWindow(c, m["maintenanceWindow"])
	r.BackupConfiguration = flattenInstanceSettingsBackupConfiguration(c, m["backupConfiguration"])
	r.DataDiskSizeGb = flattenInstanceSettingsDataDiskSizeGb(c, m["dataDiskSizeGb"])
	r.ActiveDirectoryConfig = flattenInstanceSettingsActiveDirectoryConfig(c, m["activeDirectoryConfig"])
	r.Collation = dcl.FlattenString(m["collation"])
	r.DenyMaintenancePeriods = flattenInstanceSettingsDenyMaintenancePeriodsSlice(c, m["denyMaintenancePeriods"])
	r.InsightsConfig = flattenInstanceSettingsInsightsConfig(c, m["insightsConfig"])

	return r
}

// expandInstanceSettingsSettingsVersionMap expands the contents of InstanceSettingsSettingsVersion into a JSON
// request object.
func expandInstanceSettingsSettingsVersionMap(c *Client, f map[string]InstanceSettingsSettingsVersion) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsSettingsVersion(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsSettingsVersionSlice expands the contents of InstanceSettingsSettingsVersion into a JSON
// request object.
func expandInstanceSettingsSettingsVersionSlice(c *Client, f []InstanceSettingsSettingsVersion) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsSettingsVersion(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsSettingsVersionMap flattens the contents of InstanceSettingsSettingsVersion from a JSON
// response object.
func flattenInstanceSettingsSettingsVersionMap(c *Client, i interface{}) map[string]InstanceSettingsSettingsVersion {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsSettingsVersion{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsSettingsVersion{}
	}

	items := make(map[string]InstanceSettingsSettingsVersion)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsSettingsVersion(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsSettingsVersionSlice flattens the contents of InstanceSettingsSettingsVersion from a JSON
// response object.
func flattenInstanceSettingsSettingsVersionSlice(c *Client, i interface{}) []InstanceSettingsSettingsVersion {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsSettingsVersion{}
	}

	if len(a) == 0 {
		return []InstanceSettingsSettingsVersion{}
	}

	items := make([]InstanceSettingsSettingsVersion, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsSettingsVersion(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsSettingsVersion expands an instance of InstanceSettingsSettingsVersion into a JSON
// request object.
func expandInstanceSettingsSettingsVersion(c *Client, f *InstanceSettingsSettingsVersion) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsSettingsVersion flattens an instance of InstanceSettingsSettingsVersion from a JSON
// response object.
func flattenInstanceSettingsSettingsVersion(c *Client, i interface{}) *InstanceSettingsSettingsVersion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsSettingsVersion{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceSettingsStorageAutoResizeLimitMap expands the contents of InstanceSettingsStorageAutoResizeLimit into a JSON
// request object.
func expandInstanceSettingsStorageAutoResizeLimitMap(c *Client, f map[string]InstanceSettingsStorageAutoResizeLimit) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsStorageAutoResizeLimit(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsStorageAutoResizeLimitSlice expands the contents of InstanceSettingsStorageAutoResizeLimit into a JSON
// request object.
func expandInstanceSettingsStorageAutoResizeLimitSlice(c *Client, f []InstanceSettingsStorageAutoResizeLimit) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsStorageAutoResizeLimit(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsStorageAutoResizeLimitMap flattens the contents of InstanceSettingsStorageAutoResizeLimit from a JSON
// response object.
func flattenInstanceSettingsStorageAutoResizeLimitMap(c *Client, i interface{}) map[string]InstanceSettingsStorageAutoResizeLimit {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsStorageAutoResizeLimit{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsStorageAutoResizeLimit{}
	}

	items := make(map[string]InstanceSettingsStorageAutoResizeLimit)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsStorageAutoResizeLimit(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsStorageAutoResizeLimitSlice flattens the contents of InstanceSettingsStorageAutoResizeLimit from a JSON
// response object.
func flattenInstanceSettingsStorageAutoResizeLimitSlice(c *Client, i interface{}) []InstanceSettingsStorageAutoResizeLimit {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsStorageAutoResizeLimit{}
	}

	if len(a) == 0 {
		return []InstanceSettingsStorageAutoResizeLimit{}
	}

	items := make([]InstanceSettingsStorageAutoResizeLimit, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsStorageAutoResizeLimit(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsStorageAutoResizeLimit expands an instance of InstanceSettingsStorageAutoResizeLimit into a JSON
// request object.
func expandInstanceSettingsStorageAutoResizeLimit(c *Client, f *InstanceSettingsStorageAutoResizeLimit) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsStorageAutoResizeLimit flattens an instance of InstanceSettingsStorageAutoResizeLimit from a JSON
// response object.
func flattenInstanceSettingsStorageAutoResizeLimit(c *Client, i interface{}) *InstanceSettingsStorageAutoResizeLimit {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsStorageAutoResizeLimit{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceSettingsIPConfigurationMap expands the contents of InstanceSettingsIPConfiguration into a JSON
// request object.
func expandInstanceSettingsIPConfigurationMap(c *Client, f map[string]InstanceSettingsIPConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsIPConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsIPConfigurationSlice expands the contents of InstanceSettingsIPConfiguration into a JSON
// request object.
func expandInstanceSettingsIPConfigurationSlice(c *Client, f []InstanceSettingsIPConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsIPConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsIPConfigurationMap flattens the contents of InstanceSettingsIPConfiguration from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationMap(c *Client, i interface{}) map[string]InstanceSettingsIPConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsIPConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsIPConfiguration{}
	}

	items := make(map[string]InstanceSettingsIPConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsIPConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsIPConfigurationSlice flattens the contents of InstanceSettingsIPConfiguration from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationSlice(c *Client, i interface{}) []InstanceSettingsIPConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsIPConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceSettingsIPConfiguration{}
	}

	items := make([]InstanceSettingsIPConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsIPConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsIPConfiguration expands an instance of InstanceSettingsIPConfiguration into a JSON
// request object.
func expandInstanceSettingsIPConfiguration(c *Client, f *InstanceSettingsIPConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.IPv4Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["ipv4Enabled"] = v
	}
	if v := f.PrivateNetwork; !dcl.IsEmptyValueIndirect(v) {
		m["privateNetwork"] = v
	}
	if v := f.RequireSsl; !dcl.IsEmptyValueIndirect(v) {
		m["requireSsl"] = v
	}
	if v, err := expandInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, f.AuthorizedNetworks); err != nil {
		return nil, fmt.Errorf("error expanding AuthorizedNetworks into authorizedNetworks: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authorizedNetworks"] = v
	}

	return m, nil
}

// flattenInstanceSettingsIPConfiguration flattens an instance of InstanceSettingsIPConfiguration from a JSON
// response object.
func flattenInstanceSettingsIPConfiguration(c *Client, i interface{}) *InstanceSettingsIPConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsIPConfiguration{}
	r.IPv4Enabled = dcl.FlattenBool(m["ipv4Enabled"])
	r.PrivateNetwork = dcl.FlattenString(m["privateNetwork"])
	r.RequireSsl = dcl.FlattenBool(m["requireSsl"])
	r.AuthorizedNetworks = flattenInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c, m["authorizedNetworks"])

	return r
}

// expandInstanceSettingsIPConfigurationAuthorizedNetworksMap expands the contents of InstanceSettingsIPConfigurationAuthorizedNetworks into a JSON
// request object.
func expandInstanceSettingsIPConfigurationAuthorizedNetworksMap(c *Client, f map[string]InstanceSettingsIPConfigurationAuthorizedNetworks) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsIPConfigurationAuthorizedNetworks(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsIPConfigurationAuthorizedNetworksSlice expands the contents of InstanceSettingsIPConfigurationAuthorizedNetworks into a JSON
// request object.
func expandInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, f []InstanceSettingsIPConfigurationAuthorizedNetworks) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsIPConfigurationAuthorizedNetworks(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsIPConfigurationAuthorizedNetworksMap flattens the contents of InstanceSettingsIPConfigurationAuthorizedNetworks from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationAuthorizedNetworksMap(c *Client, i interface{}) map[string]InstanceSettingsIPConfigurationAuthorizedNetworks {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	items := make(map[string]InstanceSettingsIPConfigurationAuthorizedNetworks)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsIPConfigurationAuthorizedNetworks(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsIPConfigurationAuthorizedNetworksSlice flattens the contents of InstanceSettingsIPConfigurationAuthorizedNetworks from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationAuthorizedNetworksSlice(c *Client, i interface{}) []InstanceSettingsIPConfigurationAuthorizedNetworks {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	if len(a) == 0 {
		return []InstanceSettingsIPConfigurationAuthorizedNetworks{}
	}

	items := make([]InstanceSettingsIPConfigurationAuthorizedNetworks, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsIPConfigurationAuthorizedNetworks(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsIPConfigurationAuthorizedNetworks expands an instance of InstanceSettingsIPConfigurationAuthorizedNetworks into a JSON
// request object.
func expandInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, f *InstanceSettingsIPConfigurationAuthorizedNetworks) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["expirationTime"] = v
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceSettingsIPConfigurationAuthorizedNetworks flattens an instance of InstanceSettingsIPConfigurationAuthorizedNetworks from a JSON
// response object.
func flattenInstanceSettingsIPConfigurationAuthorizedNetworks(c *Client, i interface{}) *InstanceSettingsIPConfigurationAuthorizedNetworks {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsIPConfigurationAuthorizedNetworks{}
	r.Value = dcl.FlattenString(m["value"])
	r.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	r.Name = dcl.FlattenString(m["name"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceSettingsLocationPreferenceMap expands the contents of InstanceSettingsLocationPreference into a JSON
// request object.
func expandInstanceSettingsLocationPreferenceMap(c *Client, f map[string]InstanceSettingsLocationPreference) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsLocationPreference(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsLocationPreferenceSlice expands the contents of InstanceSettingsLocationPreference into a JSON
// request object.
func expandInstanceSettingsLocationPreferenceSlice(c *Client, f []InstanceSettingsLocationPreference) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsLocationPreference(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsLocationPreferenceMap flattens the contents of InstanceSettingsLocationPreference from a JSON
// response object.
func flattenInstanceSettingsLocationPreferenceMap(c *Client, i interface{}) map[string]InstanceSettingsLocationPreference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsLocationPreference{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsLocationPreference{}
	}

	items := make(map[string]InstanceSettingsLocationPreference)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsLocationPreference(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsLocationPreferenceSlice flattens the contents of InstanceSettingsLocationPreference from a JSON
// response object.
func flattenInstanceSettingsLocationPreferenceSlice(c *Client, i interface{}) []InstanceSettingsLocationPreference {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsLocationPreference{}
	}

	if len(a) == 0 {
		return []InstanceSettingsLocationPreference{}
	}

	items := make([]InstanceSettingsLocationPreference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsLocationPreference(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsLocationPreference expands an instance of InstanceSettingsLocationPreference into a JSON
// request object.
func expandInstanceSettingsLocationPreference(c *Client, f *InstanceSettingsLocationPreference) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Zone; !dcl.IsEmptyValueIndirect(v) {
		m["zone"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceSettingsLocationPreference flattens an instance of InstanceSettingsLocationPreference from a JSON
// response object.
func flattenInstanceSettingsLocationPreference(c *Client, i interface{}) *InstanceSettingsLocationPreference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsLocationPreference{}
	r.Zone = dcl.FlattenString(m["zone"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceSettingsDatabaseFlagsMap expands the contents of InstanceSettingsDatabaseFlags into a JSON
// request object.
func expandInstanceSettingsDatabaseFlagsMap(c *Client, f map[string]InstanceSettingsDatabaseFlags) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsDatabaseFlags(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsDatabaseFlagsSlice expands the contents of InstanceSettingsDatabaseFlags into a JSON
// request object.
func expandInstanceSettingsDatabaseFlagsSlice(c *Client, f []InstanceSettingsDatabaseFlags) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsDatabaseFlags(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsDatabaseFlagsMap flattens the contents of InstanceSettingsDatabaseFlags from a JSON
// response object.
func flattenInstanceSettingsDatabaseFlagsMap(c *Client, i interface{}) map[string]InstanceSettingsDatabaseFlags {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsDatabaseFlags{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsDatabaseFlags{}
	}

	items := make(map[string]InstanceSettingsDatabaseFlags)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsDatabaseFlags(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsDatabaseFlagsSlice flattens the contents of InstanceSettingsDatabaseFlags from a JSON
// response object.
func flattenInstanceSettingsDatabaseFlagsSlice(c *Client, i interface{}) []InstanceSettingsDatabaseFlags {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDatabaseFlags{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDatabaseFlags{}
	}

	items := make([]InstanceSettingsDatabaseFlags, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDatabaseFlags(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsDatabaseFlags expands an instance of InstanceSettingsDatabaseFlags into a JSON
// request object.
func expandInstanceSettingsDatabaseFlags(c *Client, f *InstanceSettingsDatabaseFlags) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsDatabaseFlags flattens an instance of InstanceSettingsDatabaseFlags from a JSON
// response object.
func flattenInstanceSettingsDatabaseFlags(c *Client, i interface{}) *InstanceSettingsDatabaseFlags {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsDatabaseFlags{}
	r.Name = dcl.FlattenString(m["name"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandInstanceSettingsMaintenanceWindowMap expands the contents of InstanceSettingsMaintenanceWindow into a JSON
// request object.
func expandInstanceSettingsMaintenanceWindowMap(c *Client, f map[string]InstanceSettingsMaintenanceWindow) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsMaintenanceWindowSlice expands the contents of InstanceSettingsMaintenanceWindow into a JSON
// request object.
func expandInstanceSettingsMaintenanceWindowSlice(c *Client, f []InstanceSettingsMaintenanceWindow) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsMaintenanceWindow(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsMaintenanceWindowMap flattens the contents of InstanceSettingsMaintenanceWindow from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindowMap(c *Client, i interface{}) map[string]InstanceSettingsMaintenanceWindow {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsMaintenanceWindow{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsMaintenanceWindow{}
	}

	items := make(map[string]InstanceSettingsMaintenanceWindow)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsMaintenanceWindow(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsMaintenanceWindowSlice flattens the contents of InstanceSettingsMaintenanceWindow from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindowSlice(c *Client, i interface{}) []InstanceSettingsMaintenanceWindow {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsMaintenanceWindow{}
	}

	if len(a) == 0 {
		return []InstanceSettingsMaintenanceWindow{}
	}

	items := make([]InstanceSettingsMaintenanceWindow, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsMaintenanceWindow(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsMaintenanceWindow expands an instance of InstanceSettingsMaintenanceWindow into a JSON
// request object.
func expandInstanceSettingsMaintenanceWindow(c *Client, f *InstanceSettingsMaintenanceWindow) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Hour; !dcl.IsEmptyValueIndirect(v) {
		m["hour"] = v
	}
	if v := f.Day; !dcl.IsEmptyValueIndirect(v) {
		m["day"] = v
	}
	if v := f.UpdateTrack; !dcl.IsEmptyValueIndirect(v) {
		m["updateTrack"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceSettingsMaintenanceWindow flattens an instance of InstanceSettingsMaintenanceWindow from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindow(c *Client, i interface{}) *InstanceSettingsMaintenanceWindow {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsMaintenanceWindow{}
	r.Hour = dcl.FlattenInteger(m["hour"])
	r.Day = dcl.FlattenInteger(m["day"])
	r.UpdateTrack = flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum(m["updateTrack"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// expandInstanceSettingsBackupConfigurationMap expands the contents of InstanceSettingsBackupConfiguration into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationMap(c *Client, f map[string]InstanceSettingsBackupConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsBackupConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsBackupConfigurationSlice expands the contents of InstanceSettingsBackupConfiguration into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationSlice(c *Client, f []InstanceSettingsBackupConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsBackupConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsBackupConfigurationMap flattens the contents of InstanceSettingsBackupConfiguration from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationMap(c *Client, i interface{}) map[string]InstanceSettingsBackupConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsBackupConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsBackupConfiguration{}
	}

	items := make(map[string]InstanceSettingsBackupConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsBackupConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsBackupConfigurationSlice flattens the contents of InstanceSettingsBackupConfiguration from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationSlice(c *Client, i interface{}) []InstanceSettingsBackupConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsBackupConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceSettingsBackupConfiguration{}
	}

	items := make([]InstanceSettingsBackupConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsBackupConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsBackupConfiguration expands an instance of InstanceSettingsBackupConfiguration into a JSON
// request object.
func expandInstanceSettingsBackupConfiguration(c *Client, f *InstanceSettingsBackupConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.BinaryLogEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["binaryLogEnabled"] = v
	}
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c, f.BackupRetentionSettings); err != nil {
		return nil, fmt.Errorf("error expanding BackupRetentionSettings into backupRetentionSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["backupRetentionSettings"] = v
	}
	if v := f.TransactionLogRetentionDays; !dcl.IsEmptyValueIndirect(v) {
		m["transactionLogRetentionDays"] = v
	}

	return m, nil
}

// flattenInstanceSettingsBackupConfiguration flattens an instance of InstanceSettingsBackupConfiguration from a JSON
// response object.
func flattenInstanceSettingsBackupConfiguration(c *Client, i interface{}) *InstanceSettingsBackupConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsBackupConfiguration{}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.BinaryLogEnabled = dcl.FlattenBool(m["binaryLogEnabled"])
	r.Location = dcl.FlattenString(m["location"])
	r.BackupRetentionSettings = flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c, m["backupRetentionSettings"])
	r.TransactionLogRetentionDays = dcl.FlattenInteger(m["transactionLogRetentionDays"])

	return r
}

// expandInstanceSettingsBackupConfigurationBackupRetentionSettingsMap expands the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationBackupRetentionSettingsMap(c *Client, f map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice expands the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, f []InstanceSettingsBackupConfigurationBackupRetentionSettings) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsMap flattens the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsMap(c *Client, i interface{}) map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	items := make(map[string]InstanceSettingsBackupConfigurationBackupRetentionSettings)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice flattens the contents of InstanceSettingsBackupConfigurationBackupRetentionSettings from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsSlice(c *Client, i interface{}) []InstanceSettingsBackupConfigurationBackupRetentionSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	if len(a) == 0 {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	}

	items := make([]InstanceSettingsBackupConfigurationBackupRetentionSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsBackupConfigurationBackupRetentionSettings expands an instance of InstanceSettingsBackupConfigurationBackupRetentionSettings into a JSON
// request object.
func expandInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, f *InstanceSettingsBackupConfigurationBackupRetentionSettings) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.RetentionUnit; !dcl.IsEmptyValueIndirect(v) {
		m["retentionUnit"] = v
	}
	if v := f.RetainedBackups; !dcl.IsEmptyValueIndirect(v) {
		m["retainedBackups"] = v
	}

	return m, nil
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettings flattens an instance of InstanceSettingsBackupConfigurationBackupRetentionSettings from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettings(c *Client, i interface{}) *InstanceSettingsBackupConfigurationBackupRetentionSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsBackupConfigurationBackupRetentionSettings{}
	r.RetentionUnit = flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(m["retentionUnit"])
	r.RetainedBackups = dcl.FlattenInteger(m["retainedBackups"])

	return r
}

// expandInstanceSettingsDataDiskSizeGbMap expands the contents of InstanceSettingsDataDiskSizeGb into a JSON
// request object.
func expandInstanceSettingsDataDiskSizeGbMap(c *Client, f map[string]InstanceSettingsDataDiskSizeGb) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsDataDiskSizeGb(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsDataDiskSizeGbSlice expands the contents of InstanceSettingsDataDiskSizeGb into a JSON
// request object.
func expandInstanceSettingsDataDiskSizeGbSlice(c *Client, f []InstanceSettingsDataDiskSizeGb) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsDataDiskSizeGb(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsDataDiskSizeGbMap flattens the contents of InstanceSettingsDataDiskSizeGb from a JSON
// response object.
func flattenInstanceSettingsDataDiskSizeGbMap(c *Client, i interface{}) map[string]InstanceSettingsDataDiskSizeGb {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsDataDiskSizeGb{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsDataDiskSizeGb{}
	}

	items := make(map[string]InstanceSettingsDataDiskSizeGb)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsDataDiskSizeGb(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsDataDiskSizeGbSlice flattens the contents of InstanceSettingsDataDiskSizeGb from a JSON
// response object.
func flattenInstanceSettingsDataDiskSizeGbSlice(c *Client, i interface{}) []InstanceSettingsDataDiskSizeGb {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDataDiskSizeGb{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDataDiskSizeGb{}
	}

	items := make([]InstanceSettingsDataDiskSizeGb, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDataDiskSizeGb(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsDataDiskSizeGb expands an instance of InstanceSettingsDataDiskSizeGb into a JSON
// request object.
func expandInstanceSettingsDataDiskSizeGb(c *Client, f *InstanceSettingsDataDiskSizeGb) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenInstanceSettingsDataDiskSizeGb flattens an instance of InstanceSettingsDataDiskSizeGb from a JSON
// response object.
func flattenInstanceSettingsDataDiskSizeGb(c *Client, i interface{}) *InstanceSettingsDataDiskSizeGb {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsDataDiskSizeGb{}
	r.Value = dcl.FlattenInteger(m["value"])

	return r
}

// expandInstanceSettingsActiveDirectoryConfigMap expands the contents of InstanceSettingsActiveDirectoryConfig into a JSON
// request object.
func expandInstanceSettingsActiveDirectoryConfigMap(c *Client, f map[string]InstanceSettingsActiveDirectoryConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsActiveDirectoryConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsActiveDirectoryConfigSlice expands the contents of InstanceSettingsActiveDirectoryConfig into a JSON
// request object.
func expandInstanceSettingsActiveDirectoryConfigSlice(c *Client, f []InstanceSettingsActiveDirectoryConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsActiveDirectoryConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsActiveDirectoryConfigMap flattens the contents of InstanceSettingsActiveDirectoryConfig from a JSON
// response object.
func flattenInstanceSettingsActiveDirectoryConfigMap(c *Client, i interface{}) map[string]InstanceSettingsActiveDirectoryConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsActiveDirectoryConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsActiveDirectoryConfig{}
	}

	items := make(map[string]InstanceSettingsActiveDirectoryConfig)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsActiveDirectoryConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsActiveDirectoryConfigSlice flattens the contents of InstanceSettingsActiveDirectoryConfig from a JSON
// response object.
func flattenInstanceSettingsActiveDirectoryConfigSlice(c *Client, i interface{}) []InstanceSettingsActiveDirectoryConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsActiveDirectoryConfig{}
	}

	if len(a) == 0 {
		return []InstanceSettingsActiveDirectoryConfig{}
	}

	items := make([]InstanceSettingsActiveDirectoryConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsActiveDirectoryConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsActiveDirectoryConfig expands an instance of InstanceSettingsActiveDirectoryConfig into a JSON
// request object.
func expandInstanceSettingsActiveDirectoryConfig(c *Client, f *InstanceSettingsActiveDirectoryConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.Domain; !dcl.IsEmptyValueIndirect(v) {
		m["domain"] = v
	}

	return m, nil
}

// flattenInstanceSettingsActiveDirectoryConfig flattens an instance of InstanceSettingsActiveDirectoryConfig from a JSON
// response object.
func flattenInstanceSettingsActiveDirectoryConfig(c *Client, i interface{}) *InstanceSettingsActiveDirectoryConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsActiveDirectoryConfig{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.Domain = dcl.FlattenString(m["domain"])

	return r
}

// expandInstanceSettingsDenyMaintenancePeriodsMap expands the contents of InstanceSettingsDenyMaintenancePeriods into a JSON
// request object.
func expandInstanceSettingsDenyMaintenancePeriodsMap(c *Client, f map[string]InstanceSettingsDenyMaintenancePeriods) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsDenyMaintenancePeriods(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsDenyMaintenancePeriodsSlice expands the contents of InstanceSettingsDenyMaintenancePeriods into a JSON
// request object.
func expandInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, f []InstanceSettingsDenyMaintenancePeriods) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsDenyMaintenancePeriods(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsDenyMaintenancePeriodsMap flattens the contents of InstanceSettingsDenyMaintenancePeriods from a JSON
// response object.
func flattenInstanceSettingsDenyMaintenancePeriodsMap(c *Client, i interface{}) map[string]InstanceSettingsDenyMaintenancePeriods {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsDenyMaintenancePeriods{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsDenyMaintenancePeriods{}
	}

	items := make(map[string]InstanceSettingsDenyMaintenancePeriods)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsDenyMaintenancePeriods(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsDenyMaintenancePeriodsSlice flattens the contents of InstanceSettingsDenyMaintenancePeriods from a JSON
// response object.
func flattenInstanceSettingsDenyMaintenancePeriodsSlice(c *Client, i interface{}) []InstanceSettingsDenyMaintenancePeriods {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDenyMaintenancePeriods{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDenyMaintenancePeriods{}
	}

	items := make([]InstanceSettingsDenyMaintenancePeriods, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDenyMaintenancePeriods(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsDenyMaintenancePeriods expands an instance of InstanceSettingsDenyMaintenancePeriods into a JSON
// request object.
func expandInstanceSettingsDenyMaintenancePeriods(c *Client, f *InstanceSettingsDenyMaintenancePeriods) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.StartDate; !dcl.IsEmptyValueIndirect(v) {
		m["startDate"] = v
	}
	if v := f.EndDate; !dcl.IsEmptyValueIndirect(v) {
		m["endDate"] = v
	}
	if v := f.Time; !dcl.IsEmptyValueIndirect(v) {
		m["time"] = v
	}

	return m, nil
}

// flattenInstanceSettingsDenyMaintenancePeriods flattens an instance of InstanceSettingsDenyMaintenancePeriods from a JSON
// response object.
func flattenInstanceSettingsDenyMaintenancePeriods(c *Client, i interface{}) *InstanceSettingsDenyMaintenancePeriods {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsDenyMaintenancePeriods{}
	r.StartDate = dcl.FlattenString(m["startDate"])
	r.EndDate = dcl.FlattenString(m["endDate"])
	r.Time = dcl.FlattenString(m["time"])

	return r
}

// expandInstanceSettingsInsightsConfigMap expands the contents of InstanceSettingsInsightsConfig into a JSON
// request object.
func expandInstanceSettingsInsightsConfigMap(c *Client, f map[string]InstanceSettingsInsightsConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceSettingsInsightsConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceSettingsInsightsConfigSlice expands the contents of InstanceSettingsInsightsConfig into a JSON
// request object.
func expandInstanceSettingsInsightsConfigSlice(c *Client, f []InstanceSettingsInsightsConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceSettingsInsightsConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceSettingsInsightsConfigMap flattens the contents of InstanceSettingsInsightsConfig from a JSON
// response object.
func flattenInstanceSettingsInsightsConfigMap(c *Client, i interface{}) map[string]InstanceSettingsInsightsConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceSettingsInsightsConfig{}
	}

	if len(a) == 0 {
		return map[string]InstanceSettingsInsightsConfig{}
	}

	items := make(map[string]InstanceSettingsInsightsConfig)
	for k, item := range a {
		items[k] = *flattenInstanceSettingsInsightsConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceSettingsInsightsConfigSlice flattens the contents of InstanceSettingsInsightsConfig from a JSON
// response object.
func flattenInstanceSettingsInsightsConfigSlice(c *Client, i interface{}) []InstanceSettingsInsightsConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsInsightsConfig{}
	}

	if len(a) == 0 {
		return []InstanceSettingsInsightsConfig{}
	}

	items := make([]InstanceSettingsInsightsConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsInsightsConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceSettingsInsightsConfig expands an instance of InstanceSettingsInsightsConfig into a JSON
// request object.
func expandInstanceSettingsInsightsConfig(c *Client, f *InstanceSettingsInsightsConfig) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.QueryInsightsEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["queryInsightsEnabled"] = v
	}
	if v := f.RecordClientAddress; !dcl.IsEmptyValueIndirect(v) {
		m["recordClientAddress"] = v
	}
	if v := f.RecordApplicationTags; !dcl.IsEmptyValueIndirect(v) {
		m["recordApplicationTags"] = v
	}
	if v := f.QueryStringLength; !dcl.IsEmptyValueIndirect(v) {
		m["queryStringLength"] = v
	}

	return m, nil
}

// flattenInstanceSettingsInsightsConfig flattens an instance of InstanceSettingsInsightsConfig from a JSON
// response object.
func flattenInstanceSettingsInsightsConfig(c *Client, i interface{}) *InstanceSettingsInsightsConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceSettingsInsightsConfig{}
	r.QueryInsightsEnabled = dcl.FlattenBool(m["queryInsightsEnabled"])
	r.RecordClientAddress = dcl.FlattenBool(m["recordClientAddress"])
	r.RecordApplicationTags = dcl.FlattenBool(m["recordApplicationTags"])
	r.QueryStringLength = dcl.FlattenInteger(m["queryStringLength"])

	return r
}

// expandInstanceReplicaInstancesMap expands the contents of InstanceReplicaInstances into a JSON
// request object.
func expandInstanceReplicaInstancesMap(c *Client, f map[string]InstanceReplicaInstances) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceReplicaInstances(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceReplicaInstancesSlice expands the contents of InstanceReplicaInstances into a JSON
// request object.
func expandInstanceReplicaInstancesSlice(c *Client, f []InstanceReplicaInstances) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceReplicaInstances(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceReplicaInstancesMap flattens the contents of InstanceReplicaInstances from a JSON
// response object.
func flattenInstanceReplicaInstancesMap(c *Client, i interface{}) map[string]InstanceReplicaInstances {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceReplicaInstances{}
	}

	if len(a) == 0 {
		return map[string]InstanceReplicaInstances{}
	}

	items := make(map[string]InstanceReplicaInstances)
	for k, item := range a {
		items[k] = *flattenInstanceReplicaInstances(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceReplicaInstancesSlice flattens the contents of InstanceReplicaInstances from a JSON
// response object.
func flattenInstanceReplicaInstancesSlice(c *Client, i interface{}) []InstanceReplicaInstances {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceReplicaInstances{}
	}

	if len(a) == 0 {
		return []InstanceReplicaInstances{}
	}

	items := make([]InstanceReplicaInstances, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceReplicaInstances(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceReplicaInstances expands an instance of InstanceReplicaInstances into a JSON
// request object.
func expandInstanceReplicaInstances(c *Client, f *InstanceReplicaInstances) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Region; !dcl.IsEmptyValueIndirect(v) {
		m["region"] = v
	}

	return m, nil
}

// flattenInstanceReplicaInstances flattens an instance of InstanceReplicaInstances from a JSON
// response object.
func flattenInstanceReplicaInstances(c *Client, i interface{}) *InstanceReplicaInstances {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceReplicaInstances{}
	r.Name = dcl.FlattenString(m["name"])
	r.Region = dcl.FlattenString(m["region"])

	return r
}

// expandInstanceServerCaCertMap expands the contents of InstanceServerCaCert into a JSON
// request object.
func expandInstanceServerCaCertMap(c *Client, f map[string]InstanceServerCaCert) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceServerCaCert(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceServerCaCertSlice expands the contents of InstanceServerCaCert into a JSON
// request object.
func expandInstanceServerCaCertSlice(c *Client, f []InstanceServerCaCert) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceServerCaCert(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceServerCaCertMap flattens the contents of InstanceServerCaCert from a JSON
// response object.
func flattenInstanceServerCaCertMap(c *Client, i interface{}) map[string]InstanceServerCaCert {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceServerCaCert{}
	}

	if len(a) == 0 {
		return map[string]InstanceServerCaCert{}
	}

	items := make(map[string]InstanceServerCaCert)
	for k, item := range a {
		items[k] = *flattenInstanceServerCaCert(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceServerCaCertSlice flattens the contents of InstanceServerCaCert from a JSON
// response object.
func flattenInstanceServerCaCertSlice(c *Client, i interface{}) []InstanceServerCaCert {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceServerCaCert{}
	}

	if len(a) == 0 {
		return []InstanceServerCaCert{}
	}

	items := make([]InstanceServerCaCert, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceServerCaCert(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceServerCaCert expands an instance of InstanceServerCaCert into a JSON
// request object.
func expandInstanceServerCaCert(c *Client, f *InstanceServerCaCert) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.CertSerialNumber; !dcl.IsEmptyValueIndirect(v) {
		m["certSerialNumber"] = v
	}
	if v := f.Cert; !dcl.IsEmptyValueIndirect(v) {
		m["cert"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.CommonName; !dcl.IsEmptyValueIndirect(v) {
		m["commonName"] = v
	}
	if v := f.ExpirationTime; !dcl.IsEmptyValueIndirect(v) {
		m["expirationTime"] = v
	}
	if v := f.Sha1Fingerprint; !dcl.IsEmptyValueIndirect(v) {
		m["sha1Fingerprint"] = v
	}
	if v := f.Instance; !dcl.IsEmptyValueIndirect(v) {
		m["instance"] = v
	}

	return m, nil
}

// flattenInstanceServerCaCert flattens an instance of InstanceServerCaCert from a JSON
// response object.
func flattenInstanceServerCaCert(c *Client, i interface{}) *InstanceServerCaCert {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceServerCaCert{}
	r.Kind = dcl.FlattenString(m["kind"])
	r.CertSerialNumber = dcl.FlattenString(m["certSerialNumber"])
	r.Cert = dcl.FlattenString(m["cert"])
	r.CreateTime = dcl.FlattenString(m["createTime"])
	r.CommonName = dcl.FlattenString(m["commonName"])
	r.ExpirationTime = dcl.FlattenString(m["expirationTime"])
	r.Sha1Fingerprint = dcl.FlattenString(m["sha1Fingerprint"])
	r.Instance = dcl.FlattenString(m["instance"])

	return r
}

// expandInstanceOnPremisesConfigurationMap expands the contents of InstanceOnPremisesConfiguration into a JSON
// request object.
func expandInstanceOnPremisesConfigurationMap(c *Client, f map[string]InstanceOnPremisesConfiguration) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceOnPremisesConfiguration(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceOnPremisesConfigurationSlice expands the contents of InstanceOnPremisesConfiguration into a JSON
// request object.
func expandInstanceOnPremisesConfigurationSlice(c *Client, f []InstanceOnPremisesConfiguration) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceOnPremisesConfiguration(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceOnPremisesConfigurationMap flattens the contents of InstanceOnPremisesConfiguration from a JSON
// response object.
func flattenInstanceOnPremisesConfigurationMap(c *Client, i interface{}) map[string]InstanceOnPremisesConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceOnPremisesConfiguration{}
	}

	if len(a) == 0 {
		return map[string]InstanceOnPremisesConfiguration{}
	}

	items := make(map[string]InstanceOnPremisesConfiguration)
	for k, item := range a {
		items[k] = *flattenInstanceOnPremisesConfiguration(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceOnPremisesConfigurationSlice flattens the contents of InstanceOnPremisesConfiguration from a JSON
// response object.
func flattenInstanceOnPremisesConfigurationSlice(c *Client, i interface{}) []InstanceOnPremisesConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceOnPremisesConfiguration{}
	}

	if len(a) == 0 {
		return []InstanceOnPremisesConfiguration{}
	}

	items := make([]InstanceOnPremisesConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceOnPremisesConfiguration(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceOnPremisesConfiguration expands an instance of InstanceOnPremisesConfiguration into a JSON
// request object.
func expandInstanceOnPremisesConfiguration(c *Client, f *InstanceOnPremisesConfiguration) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.HostPort; !dcl.IsEmptyValueIndirect(v) {
		m["hostPort"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v := f.CaCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["caCertificate"] = v
	}
	if v := f.ClientCertificate; !dcl.IsEmptyValueIndirect(v) {
		m["clientCertificate"] = v
	}
	if v := f.ClientKey; !dcl.IsEmptyValueIndirect(v) {
		m["clientKey"] = v
	}
	if v := f.DumpFilePath; !dcl.IsEmptyValueIndirect(v) {
		m["dumpFilePath"] = v
	}
	if v := f.Database; !dcl.IsEmptyValueIndirect(v) {
		m["database"] = v
	}
	if v := f.ReplicatedDatabases; !dcl.IsEmptyValueIndirect(v) {
		m["replicatedDatabases"] = v
	}

	return m, nil
}

// flattenInstanceOnPremisesConfiguration flattens an instance of InstanceOnPremisesConfiguration from a JSON
// response object.
func flattenInstanceOnPremisesConfiguration(c *Client, i interface{}) *InstanceOnPremisesConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceOnPremisesConfiguration{}
	r.HostPort = dcl.FlattenString(m["hostPort"])
	r.Kind = dcl.FlattenString(m["kind"])
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])
	r.CaCertificate = dcl.FlattenString(m["caCertificate"])
	r.ClientCertificate = dcl.FlattenString(m["clientCertificate"])
	r.ClientKey = dcl.FlattenString(m["clientKey"])
	r.DumpFilePath = dcl.FlattenString(m["dumpFilePath"])
	r.Database = dcl.FlattenString(m["database"])
	r.ReplicatedDatabases = dcl.FlattenStringSlice(m["replicatedDatabases"])

	return r
}

// expandInstanceDiskEncryptionStatusMap expands the contents of InstanceDiskEncryptionStatus into a JSON
// request object.
func expandInstanceDiskEncryptionStatusMap(c *Client, f map[string]InstanceDiskEncryptionStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInstanceDiskEncryptionStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInstanceDiskEncryptionStatusSlice expands the contents of InstanceDiskEncryptionStatus into a JSON
// request object.
func expandInstanceDiskEncryptionStatusSlice(c *Client, f []InstanceDiskEncryptionStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInstanceDiskEncryptionStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInstanceDiskEncryptionStatusMap flattens the contents of InstanceDiskEncryptionStatus from a JSON
// response object.
func flattenInstanceDiskEncryptionStatusMap(c *Client, i interface{}) map[string]InstanceDiskEncryptionStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InstanceDiskEncryptionStatus{}
	}

	if len(a) == 0 {
		return map[string]InstanceDiskEncryptionStatus{}
	}

	items := make(map[string]InstanceDiskEncryptionStatus)
	for k, item := range a {
		items[k] = *flattenInstanceDiskEncryptionStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenInstanceDiskEncryptionStatusSlice flattens the contents of InstanceDiskEncryptionStatus from a JSON
// response object.
func flattenInstanceDiskEncryptionStatusSlice(c *Client, i interface{}) []InstanceDiskEncryptionStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDiskEncryptionStatus{}
	}

	if len(a) == 0 {
		return []InstanceDiskEncryptionStatus{}
	}

	items := make([]InstanceDiskEncryptionStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDiskEncryptionStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandInstanceDiskEncryptionStatus expands an instance of InstanceDiskEncryptionStatus into a JSON
// request object.
func expandInstanceDiskEncryptionStatus(c *Client, f *InstanceDiskEncryptionStatus) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}
	if v := f.KmsKeyVersionName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyVersionName"] = v
	}
	if v := f.Kind; !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenInstanceDiskEncryptionStatus flattens an instance of InstanceDiskEncryptionStatus from a JSON
// response object.
func flattenInstanceDiskEncryptionStatus(c *Client, i interface{}) *InstanceDiskEncryptionStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InstanceDiskEncryptionStatus{}
	r.KmsKeyVersionName = dcl.FlattenString(m["kmsKeyVersionName"])
	r.Kind = dcl.FlattenString(m["kind"])

	return r
}

// flattenInstanceBackendTypeEnumSlice flattens the contents of InstanceBackendTypeEnum from a JSON
// response object.
func flattenInstanceBackendTypeEnumSlice(c *Client, i interface{}) []InstanceBackendTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceBackendTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceBackendTypeEnum{}
	}

	items := make([]InstanceBackendTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceBackendTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceBackendTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceBackendTypeEnum with the same value as that string.
func flattenInstanceBackendTypeEnum(i interface{}) *InstanceBackendTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceBackendTypeEnumRef("")
	}

	return InstanceBackendTypeEnumRef(s)
}

// flattenInstanceDatabaseVersionEnumSlice flattens the contents of InstanceDatabaseVersionEnum from a JSON
// response object.
func flattenInstanceDatabaseVersionEnumSlice(c *Client, i interface{}) []InstanceDatabaseVersionEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceDatabaseVersionEnum{}
	}

	if len(a) == 0 {
		return []InstanceDatabaseVersionEnum{}
	}

	items := make([]InstanceDatabaseVersionEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceDatabaseVersionEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceDatabaseVersionEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceDatabaseVersionEnum with the same value as that string.
func flattenInstanceDatabaseVersionEnum(i interface{}) *InstanceDatabaseVersionEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceDatabaseVersionEnumRef("")
	}

	return InstanceDatabaseVersionEnumRef(s)
}

// flattenInstanceInstanceTypeEnumSlice flattens the contents of InstanceInstanceTypeEnum from a JSON
// response object.
func flattenInstanceInstanceTypeEnumSlice(c *Client, i interface{}) []InstanceInstanceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceInstanceTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceInstanceTypeEnum{}
	}

	items := make([]InstanceInstanceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceInstanceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceInstanceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceInstanceTypeEnum with the same value as that string.
func flattenInstanceInstanceTypeEnum(i interface{}) *InstanceInstanceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceInstanceTypeEnumRef("")
	}

	return InstanceInstanceTypeEnumRef(s)
}

// flattenInstanceIPAddressesTypeEnumSlice flattens the contents of InstanceIPAddressesTypeEnum from a JSON
// response object.
func flattenInstanceIPAddressesTypeEnumSlice(c *Client, i interface{}) []InstanceIPAddressesTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceIPAddressesTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceIPAddressesTypeEnum{}
	}

	items := make([]InstanceIPAddressesTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceIPAddressesTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceIPAddressesTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceIPAddressesTypeEnum with the same value as that string.
func flattenInstanceIPAddressesTypeEnum(i interface{}) *InstanceIPAddressesTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceIPAddressesTypeEnumRef("")
	}

	return InstanceIPAddressesTypeEnumRef(s)
}

// flattenInstanceSettingsAvailabilityTypeEnumSlice flattens the contents of InstanceSettingsAvailabilityTypeEnum from a JSON
// response object.
func flattenInstanceSettingsAvailabilityTypeEnumSlice(c *Client, i interface{}) []InstanceSettingsAvailabilityTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsAvailabilityTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsAvailabilityTypeEnum{}
	}

	items := make([]InstanceSettingsAvailabilityTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsAvailabilityTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsAvailabilityTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsAvailabilityTypeEnum with the same value as that string.
func flattenInstanceSettingsAvailabilityTypeEnum(i interface{}) *InstanceSettingsAvailabilityTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsAvailabilityTypeEnumRef("")
	}

	return InstanceSettingsAvailabilityTypeEnumRef(s)
}

// flattenInstanceSettingsPricingPlanEnumSlice flattens the contents of InstanceSettingsPricingPlanEnum from a JSON
// response object.
func flattenInstanceSettingsPricingPlanEnumSlice(c *Client, i interface{}) []InstanceSettingsPricingPlanEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsPricingPlanEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsPricingPlanEnum{}
	}

	items := make([]InstanceSettingsPricingPlanEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsPricingPlanEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsPricingPlanEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsPricingPlanEnum with the same value as that string.
func flattenInstanceSettingsPricingPlanEnum(i interface{}) *InstanceSettingsPricingPlanEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsPricingPlanEnumRef("")
	}

	return InstanceSettingsPricingPlanEnumRef(s)
}

// flattenInstanceSettingsReplicationTypeEnumSlice flattens the contents of InstanceSettingsReplicationTypeEnum from a JSON
// response object.
func flattenInstanceSettingsReplicationTypeEnumSlice(c *Client, i interface{}) []InstanceSettingsReplicationTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsReplicationTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsReplicationTypeEnum{}
	}

	items := make([]InstanceSettingsReplicationTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsReplicationTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsReplicationTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsReplicationTypeEnum with the same value as that string.
func flattenInstanceSettingsReplicationTypeEnum(i interface{}) *InstanceSettingsReplicationTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsReplicationTypeEnumRef("")
	}

	return InstanceSettingsReplicationTypeEnumRef(s)
}

// flattenInstanceSettingsActivationPolicyEnumSlice flattens the contents of InstanceSettingsActivationPolicyEnum from a JSON
// response object.
func flattenInstanceSettingsActivationPolicyEnumSlice(c *Client, i interface{}) []InstanceSettingsActivationPolicyEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsActivationPolicyEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsActivationPolicyEnum{}
	}

	items := make([]InstanceSettingsActivationPolicyEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsActivationPolicyEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsActivationPolicyEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsActivationPolicyEnum with the same value as that string.
func flattenInstanceSettingsActivationPolicyEnum(i interface{}) *InstanceSettingsActivationPolicyEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsActivationPolicyEnumRef("")
	}

	return InstanceSettingsActivationPolicyEnumRef(s)
}

// flattenInstanceSettingsDataDiskTypeEnumSlice flattens the contents of InstanceSettingsDataDiskTypeEnum from a JSON
// response object.
func flattenInstanceSettingsDataDiskTypeEnumSlice(c *Client, i interface{}) []InstanceSettingsDataDiskTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsDataDiskTypeEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsDataDiskTypeEnum{}
	}

	items := make([]InstanceSettingsDataDiskTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsDataDiskTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsDataDiskTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsDataDiskTypeEnum with the same value as that string.
func flattenInstanceSettingsDataDiskTypeEnum(i interface{}) *InstanceSettingsDataDiskTypeEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsDataDiskTypeEnumRef("")
	}

	return InstanceSettingsDataDiskTypeEnumRef(s)
}

// flattenInstanceSettingsMaintenanceWindowUpdateTrackEnumSlice flattens the contents of InstanceSettingsMaintenanceWindowUpdateTrackEnum from a JSON
// response object.
func flattenInstanceSettingsMaintenanceWindowUpdateTrackEnumSlice(c *Client, i interface{}) []InstanceSettingsMaintenanceWindowUpdateTrackEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsMaintenanceWindowUpdateTrackEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsMaintenanceWindowUpdateTrackEnum{}
	}

	items := make([]InstanceSettingsMaintenanceWindowUpdateTrackEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsMaintenanceWindowUpdateTrackEnum with the same value as that string.
func flattenInstanceSettingsMaintenanceWindowUpdateTrackEnum(i interface{}) *InstanceSettingsMaintenanceWindowUpdateTrackEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsMaintenanceWindowUpdateTrackEnumRef("")
	}

	return InstanceSettingsMaintenanceWindowUpdateTrackEnumRef(s)
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumSlice flattens the contents of InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum from a JSON
// response object.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumSlice(c *Client, i interface{}) []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum{}
	}

	if len(a) == 0 {
		return []InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum{}
	}

	items := make([]InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(item.(interface{})))
	}

	return items
}

// flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum asserts that an interface is a string, and returns a
// pointer to a *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum with the same value as that string.
func flattenInstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum(i interface{}) *InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnum {
	s, ok := i.(string)
	if !ok {
		return InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumRef("")
	}

	return InstanceSettingsBackupConfigurationBackupRetentionSettingsRetentionUnitEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Instance) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInstance(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}
