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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Cluster) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "azureRegion"); err != nil {
		return err
	}
	if err := dcl.Required(r, "resourceGroupId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "client"); err != nil {
		return err
	}
	if err := dcl.Required(r, "networking"); err != nil {
		return err
	}
	if err := dcl.Required(r, "controlPlane"); err != nil {
		return err
	}
	if err := dcl.Required(r, "authorization"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Networking) {
		if err := r.Networking.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ControlPlane) {
		if err := r.ControlPlane.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Authorization) {
		if err := r.Authorization.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.WorkloadIdentityConfig) {
		if err := r.WorkloadIdentityConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterNetworking) validate() error {
	if err := dcl.Required(r, "virtualNetworkId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "podAddressCidrBlocks"); err != nil {
		return err
	}
	if err := dcl.Required(r, "serviceAddressCidrBlocks"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterControlPlane) validate() error {
	if err := dcl.Required(r, "version"); err != nil {
		return err
	}
	if err := dcl.Required(r, "subnetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "sshConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SshConfig) {
		if err := r.SshConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RootVolume) {
		if err := r.RootVolume.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MainVolume) {
		if err := r.MainVolume.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DatabaseEncryption) {
		if err := r.DatabaseEncryption.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ClusterControlPlaneSshConfig) validate() error {
	if err := dcl.Required(r, "authorizedKey"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterControlPlaneRootVolume) validate() error {
	return nil
}
func (r *ClusterControlPlaneMainVolume) validate() error {
	return nil
}
func (r *ClusterControlPlaneDatabaseEncryption) validate() error {
	if err := dcl.Required(r, "resourceGroupId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "kmsKeyIdentifier"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterAuthorization) validate() error {
	if err := dcl.Required(r, "adminUsers"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterAuthorizationAdminUsers) validate() error {
	if err := dcl.Required(r, "username"); err != nil {
		return err
	}
	return nil
}
func (r *ClusterWorkloadIdentityConfig) validate() error {
	return nil
}
func (r *Cluster) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-gkemulticloud.googleapis.com/v1", params)
}

func (r *Cluster) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Cluster) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters", nr.basePath(), userBasePath, params), nil

}

func (r *Cluster) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters?azureClusterId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Cluster) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/azureClusters/{{name}}", nr.basePath(), userBasePath, params), nil
}

// clusterApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type clusterApiOperation interface {
	do(context.Context, *Cluster, *Client) error
}

func (c *Client) listClusterRaw(ctx context.Context, r *Cluster, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ClusterMaxPage {
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

type listClusterOperation struct {
	AzureClusters []map[string]interface{} `json:"azureClusters"`
	Token         string                   `json:"nextPageToken"`
}

func (c *Client) listCluster(ctx context.Context, r *Cluster, pageToken string, pageSize int32) ([]*Cluster, string, error) {
	b, err := c.listClusterRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listClusterOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Cluster
	for _, v := range m.AzureClusters {
		res, err := unmarshalMapCluster(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCluster(ctx context.Context, f func(*Cluster) bool, resources []*Cluster) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCluster(ctx, res)
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

type deleteClusterOperation struct{}

func (op *deleteClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	r, err := c.GetCluster(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Cluster not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetCluster checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetCluster(ctx, r)
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
type createClusterOperation struct {
	response map[string]interface{}
}

func (op *createClusterOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createClusterOperation) do(ctx context.Context, r *Cluster, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
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
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetCluster(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getClusterRaw(ctx context.Context, r *Cluster) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
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

func (c *Client) clusterDiffsForRawDesired(ctx context.Context, rawDesired *Cluster, opts ...dcl.ApplyOption) (initial, desired *Cluster, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Cluster
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Cluster); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Cluster, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCluster(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Cluster resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Cluster resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Cluster resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Cluster: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Cluster: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeClusterInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Cluster: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeClusterDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Cluster: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCluster(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeClusterInitialState(rawInitial, rawDesired *Cluster) (*Cluster, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeClusterDesiredState(rawDesired, rawInitial *Cluster, opts ...dcl.ApplyOption) (*Cluster, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Networking = canonicalizeClusterNetworking(rawDesired.Networking, nil, opts...)
		rawDesired.ControlPlane = canonicalizeClusterControlPlane(rawDesired.ControlPlane, nil, opts...)
		rawDesired.Authorization = canonicalizeClusterAuthorization(rawDesired.Authorization, nil, opts...)
		rawDesired.WorkloadIdentityConfig = canonicalizeClusterWorkloadIdentityConfig(rawDesired.WorkloadIdentityConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Cluster{}
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.AzureRegion, rawInitial.AzureRegion) {
		canonicalDesired.AzureRegion = rawInitial.AzureRegion
	} else {
		canonicalDesired.AzureRegion = rawDesired.AzureRegion
	}
	if dcl.StringCanonicalize(rawDesired.ResourceGroupId, rawInitial.ResourceGroupId) {
		canonicalDesired.ResourceGroupId = rawInitial.ResourceGroupId
	} else {
		canonicalDesired.ResourceGroupId = rawDesired.ResourceGroupId
	}
	if dcl.NameToSelfLink(rawDesired.Client, rawInitial.Client) {
		canonicalDesired.Client = rawInitial.Client
	} else {
		canonicalDesired.Client = rawDesired.Client
	}
	canonicalDesired.Networking = canonicalizeClusterNetworking(rawDesired.Networking, rawInitial.Networking, opts...)
	canonicalDesired.ControlPlane = canonicalizeClusterControlPlane(rawDesired.ControlPlane, rawInitial.ControlPlane, opts...)
	canonicalDesired.Authorization = canonicalizeClusterAuthorization(rawDesired.Authorization, rawInitial.Authorization, opts...)
	if dcl.IsZeroValue(rawDesired.Annotations) {
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}

	return canonicalDesired, nil
}

func canonicalizeClusterNewState(c *Client, rawNew, rawDesired *Cluster) (*Cluster, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.AzureRegion) && dcl.IsNotReturnedByServer(rawDesired.AzureRegion) {
		rawNew.AzureRegion = rawDesired.AzureRegion
	} else {
		if dcl.StringCanonicalize(rawDesired.AzureRegion, rawNew.AzureRegion) {
			rawNew.AzureRegion = rawDesired.AzureRegion
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.ResourceGroupId) && dcl.IsNotReturnedByServer(rawDesired.ResourceGroupId) {
		rawNew.ResourceGroupId = rawDesired.ResourceGroupId
	} else {
		if dcl.StringCanonicalize(rawDesired.ResourceGroupId, rawNew.ResourceGroupId) {
			rawNew.ResourceGroupId = rawDesired.ResourceGroupId
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Client) && dcl.IsNotReturnedByServer(rawDesired.Client) {
		rawNew.Client = rawDesired.Client
	} else {
		if dcl.NameToSelfLink(rawDesired.Client, rawNew.Client) {
			rawNew.Client = rawDesired.Client
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Networking) && dcl.IsNotReturnedByServer(rawDesired.Networking) {
		rawNew.Networking = rawDesired.Networking
	} else {
		rawNew.Networking = canonicalizeNewClusterNetworking(c, rawDesired.Networking, rawNew.Networking)
	}

	if dcl.IsNotReturnedByServer(rawNew.ControlPlane) && dcl.IsNotReturnedByServer(rawDesired.ControlPlane) {
		rawNew.ControlPlane = rawDesired.ControlPlane
	} else {
		rawNew.ControlPlane = canonicalizeNewClusterControlPlane(c, rawDesired.ControlPlane, rawNew.ControlPlane)
	}

	if dcl.IsNotReturnedByServer(rawNew.Authorization) && dcl.IsNotReturnedByServer(rawDesired.Authorization) {
		rawNew.Authorization = rawDesired.Authorization
	} else {
		rawNew.Authorization = canonicalizeNewClusterAuthorization(c, rawDesired.Authorization, rawNew.Authorization)
	}

	if dcl.IsNotReturnedByServer(rawNew.State) && dcl.IsNotReturnedByServer(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Endpoint) && dcl.IsNotReturnedByServer(rawDesired.Endpoint) {
		rawNew.Endpoint = rawDesired.Endpoint
	} else {
		if dcl.StringCanonicalize(rawDesired.Endpoint, rawNew.Endpoint) {
			rawNew.Endpoint = rawDesired.Endpoint
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Uid) && dcl.IsNotReturnedByServer(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Reconciling) && dcl.IsNotReturnedByServer(rawDesired.Reconciling) {
		rawNew.Reconciling = rawDesired.Reconciling
	} else {
		if dcl.BoolCanonicalize(rawDesired.Reconciling, rawNew.Reconciling) {
			rawNew.Reconciling = rawDesired.Reconciling
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Etag) && dcl.IsNotReturnedByServer(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Annotations) && dcl.IsNotReturnedByServer(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.WorkloadIdentityConfig) && dcl.IsNotReturnedByServer(rawDesired.WorkloadIdentityConfig) {
		rawNew.WorkloadIdentityConfig = rawDesired.WorkloadIdentityConfig
	} else {
		rawNew.WorkloadIdentityConfig = canonicalizeNewClusterWorkloadIdentityConfig(c, rawDesired.WorkloadIdentityConfig, rawNew.WorkloadIdentityConfig)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeClusterNetworking(des, initial *ClusterNetworking, opts ...dcl.ApplyOption) *ClusterNetworking {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterNetworking{}

	if dcl.StringCanonicalize(des.VirtualNetworkId, initial.VirtualNetworkId) || dcl.IsZeroValue(des.VirtualNetworkId) {
		cDes.VirtualNetworkId = initial.VirtualNetworkId
	} else {
		cDes.VirtualNetworkId = des.VirtualNetworkId
	}
	if dcl.IsZeroValue(des.PodAddressCidrBlocks) {
		des.PodAddressCidrBlocks = initial.PodAddressCidrBlocks
	} else {
		cDes.PodAddressCidrBlocks = des.PodAddressCidrBlocks
	}
	if dcl.IsZeroValue(des.ServiceAddressCidrBlocks) {
		des.ServiceAddressCidrBlocks = initial.ServiceAddressCidrBlocks
	} else {
		cDes.ServiceAddressCidrBlocks = des.ServiceAddressCidrBlocks
	}

	return cDes
}

func canonicalizeClusterNetworkingSlice(des, initial []ClusterNetworking, opts ...dcl.ApplyOption) []ClusterNetworking {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterNetworking, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterNetworking(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterNetworking, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterNetworking(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterNetworking(c *Client, des, nw *ClusterNetworking) *ClusterNetworking {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterNetworking while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.VirtualNetworkId, nw.VirtualNetworkId) {
		nw.VirtualNetworkId = des.VirtualNetworkId
	}

	return nw
}

func canonicalizeNewClusterNetworkingSet(c *Client, des, nw []ClusterNetworking) []ClusterNetworking {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterNetworking
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterNetworkingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterNetworkingSlice(c *Client, des, nw []ClusterNetworking) []ClusterNetworking {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterNetworking
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterNetworking(c, &d, &n))
	}

	return items
}

func canonicalizeClusterControlPlane(des, initial *ClusterControlPlane, opts ...dcl.ApplyOption) *ClusterControlPlane {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterControlPlane{}

	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	if dcl.StringCanonicalize(des.SubnetId, initial.SubnetId) || dcl.IsZeroValue(des.SubnetId) {
		cDes.SubnetId = initial.SubnetId
	} else {
		cDes.SubnetId = des.SubnetId
	}
	if dcl.StringCanonicalize(des.VmSize, initial.VmSize) || dcl.IsZeroValue(des.VmSize) {
		cDes.VmSize = initial.VmSize
	} else {
		cDes.VmSize = des.VmSize
	}
	cDes.SshConfig = canonicalizeClusterControlPlaneSshConfig(des.SshConfig, initial.SshConfig, opts...)
	cDes.RootVolume = canonicalizeClusterControlPlaneRootVolume(des.RootVolume, initial.RootVolume, opts...)
	cDes.MainVolume = canonicalizeClusterControlPlaneMainVolume(des.MainVolume, initial.MainVolume, opts...)
	cDes.DatabaseEncryption = canonicalizeClusterControlPlaneDatabaseEncryption(des.DatabaseEncryption, initial.DatabaseEncryption, opts...)
	if dcl.IsZeroValue(des.Tags) {
		des.Tags = initial.Tags
	} else {
		cDes.Tags = des.Tags
	}

	return cDes
}

func canonicalizeClusterControlPlaneSlice(des, initial []ClusterControlPlane, opts ...dcl.ApplyOption) []ClusterControlPlane {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterControlPlane, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterControlPlane(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterControlPlane, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterControlPlane(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterControlPlane(c *Client, des, nw *ClusterControlPlane) *ClusterControlPlane {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterControlPlane while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	if dcl.StringCanonicalize(des.SubnetId, nw.SubnetId) {
		nw.SubnetId = des.SubnetId
	}
	if dcl.StringCanonicalize(des.VmSize, nw.VmSize) {
		nw.VmSize = des.VmSize
	}
	nw.SshConfig = canonicalizeNewClusterControlPlaneSshConfig(c, des.SshConfig, nw.SshConfig)
	nw.RootVolume = canonicalizeNewClusterControlPlaneRootVolume(c, des.RootVolume, nw.RootVolume)
	nw.MainVolume = canonicalizeNewClusterControlPlaneMainVolume(c, des.MainVolume, nw.MainVolume)
	nw.DatabaseEncryption = canonicalizeNewClusterControlPlaneDatabaseEncryption(c, des.DatabaseEncryption, nw.DatabaseEncryption)

	return nw
}

func canonicalizeNewClusterControlPlaneSet(c *Client, des, nw []ClusterControlPlane) []ClusterControlPlane {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterControlPlane
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterControlPlaneNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterControlPlaneSlice(c *Client, des, nw []ClusterControlPlane) []ClusterControlPlane {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterControlPlane
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterControlPlane(c, &d, &n))
	}

	return items
}

func canonicalizeClusterControlPlaneSshConfig(des, initial *ClusterControlPlaneSshConfig, opts ...dcl.ApplyOption) *ClusterControlPlaneSshConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterControlPlaneSshConfig{}

	if dcl.StringCanonicalize(des.AuthorizedKey, initial.AuthorizedKey) || dcl.IsZeroValue(des.AuthorizedKey) {
		cDes.AuthorizedKey = initial.AuthorizedKey
	} else {
		cDes.AuthorizedKey = des.AuthorizedKey
	}

	return cDes
}

func canonicalizeClusterControlPlaneSshConfigSlice(des, initial []ClusterControlPlaneSshConfig, opts ...dcl.ApplyOption) []ClusterControlPlaneSshConfig {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterControlPlaneSshConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterControlPlaneSshConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterControlPlaneSshConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterControlPlaneSshConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterControlPlaneSshConfig(c *Client, des, nw *ClusterControlPlaneSshConfig) *ClusterControlPlaneSshConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterControlPlaneSshConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.AuthorizedKey, nw.AuthorizedKey) {
		nw.AuthorizedKey = des.AuthorizedKey
	}

	return nw
}

func canonicalizeNewClusterControlPlaneSshConfigSet(c *Client, des, nw []ClusterControlPlaneSshConfig) []ClusterControlPlaneSshConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterControlPlaneSshConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterControlPlaneSshConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterControlPlaneSshConfigSlice(c *Client, des, nw []ClusterControlPlaneSshConfig) []ClusterControlPlaneSshConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterControlPlaneSshConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterControlPlaneSshConfig(c, &d, &n))
	}

	return items
}

func canonicalizeClusterControlPlaneRootVolume(des, initial *ClusterControlPlaneRootVolume, opts ...dcl.ApplyOption) *ClusterControlPlaneRootVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterControlPlaneRootVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}

	return cDes
}

func canonicalizeClusterControlPlaneRootVolumeSlice(des, initial []ClusterControlPlaneRootVolume, opts ...dcl.ApplyOption) []ClusterControlPlaneRootVolume {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterControlPlaneRootVolume, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterControlPlaneRootVolume(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterControlPlaneRootVolume, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterControlPlaneRootVolume(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterControlPlaneRootVolume(c *Client, des, nw *ClusterControlPlaneRootVolume) *ClusterControlPlaneRootVolume {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterControlPlaneRootVolume while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewClusterControlPlaneRootVolumeSet(c *Client, des, nw []ClusterControlPlaneRootVolume) []ClusterControlPlaneRootVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterControlPlaneRootVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterControlPlaneRootVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterControlPlaneRootVolumeSlice(c *Client, des, nw []ClusterControlPlaneRootVolume) []ClusterControlPlaneRootVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterControlPlaneRootVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterControlPlaneRootVolume(c, &d, &n))
	}

	return items
}

func canonicalizeClusterControlPlaneMainVolume(des, initial *ClusterControlPlaneMainVolume, opts ...dcl.ApplyOption) *ClusterControlPlaneMainVolume {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterControlPlaneMainVolume{}

	if dcl.IsZeroValue(des.SizeGib) {
		des.SizeGib = initial.SizeGib
	} else {
		cDes.SizeGib = des.SizeGib
	}

	return cDes
}

func canonicalizeClusterControlPlaneMainVolumeSlice(des, initial []ClusterControlPlaneMainVolume, opts ...dcl.ApplyOption) []ClusterControlPlaneMainVolume {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterControlPlaneMainVolume, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterControlPlaneMainVolume(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterControlPlaneMainVolume, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterControlPlaneMainVolume(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterControlPlaneMainVolume(c *Client, des, nw *ClusterControlPlaneMainVolume) *ClusterControlPlaneMainVolume {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterControlPlaneMainVolume while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewClusterControlPlaneMainVolumeSet(c *Client, des, nw []ClusterControlPlaneMainVolume) []ClusterControlPlaneMainVolume {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterControlPlaneMainVolume
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterControlPlaneMainVolumeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterControlPlaneMainVolumeSlice(c *Client, des, nw []ClusterControlPlaneMainVolume) []ClusterControlPlaneMainVolume {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterControlPlaneMainVolume
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterControlPlaneMainVolume(c, &d, &n))
	}

	return items
}

func canonicalizeClusterControlPlaneDatabaseEncryption(des, initial *ClusterControlPlaneDatabaseEncryption, opts ...dcl.ApplyOption) *ClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterControlPlaneDatabaseEncryption{}

	if dcl.StringCanonicalize(des.ResourceGroupId, initial.ResourceGroupId) || dcl.IsZeroValue(des.ResourceGroupId) {
		cDes.ResourceGroupId = initial.ResourceGroupId
	} else {
		cDes.ResourceGroupId = des.ResourceGroupId
	}
	if dcl.StringCanonicalize(des.KmsKeyIdentifier, initial.KmsKeyIdentifier) || dcl.IsZeroValue(des.KmsKeyIdentifier) {
		cDes.KmsKeyIdentifier = initial.KmsKeyIdentifier
	} else {
		cDes.KmsKeyIdentifier = des.KmsKeyIdentifier
	}

	return cDes
}

func canonicalizeClusterControlPlaneDatabaseEncryptionSlice(des, initial []ClusterControlPlaneDatabaseEncryption, opts ...dcl.ApplyOption) []ClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterControlPlaneDatabaseEncryption, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterControlPlaneDatabaseEncryption(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterControlPlaneDatabaseEncryption, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterControlPlaneDatabaseEncryption(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterControlPlaneDatabaseEncryption(c *Client, des, nw *ClusterControlPlaneDatabaseEncryption) *ClusterControlPlaneDatabaseEncryption {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterControlPlaneDatabaseEncryption while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceGroupId, nw.ResourceGroupId) {
		nw.ResourceGroupId = des.ResourceGroupId
	}
	if dcl.StringCanonicalize(des.KmsKeyIdentifier, nw.KmsKeyIdentifier) {
		nw.KmsKeyIdentifier = des.KmsKeyIdentifier
	}

	return nw
}

func canonicalizeNewClusterControlPlaneDatabaseEncryptionSet(c *Client, des, nw []ClusterControlPlaneDatabaseEncryption) []ClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterControlPlaneDatabaseEncryption
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterControlPlaneDatabaseEncryptionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterControlPlaneDatabaseEncryptionSlice(c *Client, des, nw []ClusterControlPlaneDatabaseEncryption) []ClusterControlPlaneDatabaseEncryption {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterControlPlaneDatabaseEncryption
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterControlPlaneDatabaseEncryption(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAuthorization(des, initial *ClusterAuthorization, opts ...dcl.ApplyOption) *ClusterAuthorization {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterAuthorization{}

	cDes.AdminUsers = canonicalizeClusterAuthorizationAdminUsersSlice(des.AdminUsers, initial.AdminUsers, opts...)

	return cDes
}

func canonicalizeClusterAuthorizationSlice(des, initial []ClusterAuthorization, opts ...dcl.ApplyOption) []ClusterAuthorization {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterAuthorization, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterAuthorization(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterAuthorization, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterAuthorization(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterAuthorization(c *Client, des, nw *ClusterAuthorization) *ClusterAuthorization {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterAuthorization while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.AdminUsers = canonicalizeNewClusterAuthorizationAdminUsersSlice(c, des.AdminUsers, nw.AdminUsers)

	return nw
}

func canonicalizeNewClusterAuthorizationSet(c *Client, des, nw []ClusterAuthorization) []ClusterAuthorization {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAuthorization
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterAuthorizationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterAuthorizationSlice(c *Client, des, nw []ClusterAuthorization) []ClusterAuthorization {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterAuthorization
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAuthorization(c, &d, &n))
	}

	return items
}

func canonicalizeClusterAuthorizationAdminUsers(des, initial *ClusterAuthorizationAdminUsers, opts ...dcl.ApplyOption) *ClusterAuthorizationAdminUsers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterAuthorizationAdminUsers{}

	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		cDes.Username = initial.Username
	} else {
		cDes.Username = des.Username
	}

	return cDes
}

func canonicalizeClusterAuthorizationAdminUsersSlice(des, initial []ClusterAuthorizationAdminUsers, opts ...dcl.ApplyOption) []ClusterAuthorizationAdminUsers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterAuthorizationAdminUsers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterAuthorizationAdminUsers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterAuthorizationAdminUsers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterAuthorizationAdminUsers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterAuthorizationAdminUsers(c *Client, des, nw *ClusterAuthorizationAdminUsers) *ClusterAuthorizationAdminUsers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterAuthorizationAdminUsers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}

	return nw
}

func canonicalizeNewClusterAuthorizationAdminUsersSet(c *Client, des, nw []ClusterAuthorizationAdminUsers) []ClusterAuthorizationAdminUsers {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterAuthorizationAdminUsers
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterAuthorizationAdminUsersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterAuthorizationAdminUsersSlice(c *Client, des, nw []ClusterAuthorizationAdminUsers) []ClusterAuthorizationAdminUsers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterAuthorizationAdminUsers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterAuthorizationAdminUsers(c, &d, &n))
	}

	return items
}

func canonicalizeClusterWorkloadIdentityConfig(des, initial *ClusterWorkloadIdentityConfig, opts ...dcl.ApplyOption) *ClusterWorkloadIdentityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ClusterWorkloadIdentityConfig{}

	if dcl.StringCanonicalize(des.IssuerUri, initial.IssuerUri) || dcl.IsZeroValue(des.IssuerUri) {
		cDes.IssuerUri = initial.IssuerUri
	} else {
		cDes.IssuerUri = des.IssuerUri
	}
	if dcl.StringCanonicalize(des.WorkloadPool, initial.WorkloadPool) || dcl.IsZeroValue(des.WorkloadPool) {
		cDes.WorkloadPool = initial.WorkloadPool
	} else {
		cDes.WorkloadPool = des.WorkloadPool
	}
	if dcl.StringCanonicalize(des.IdentityProvider, initial.IdentityProvider) || dcl.IsZeroValue(des.IdentityProvider) {
		cDes.IdentityProvider = initial.IdentityProvider
	} else {
		cDes.IdentityProvider = des.IdentityProvider
	}

	return cDes
}

func canonicalizeClusterWorkloadIdentityConfigSlice(des, initial []ClusterWorkloadIdentityConfig, opts ...dcl.ApplyOption) []ClusterWorkloadIdentityConfig {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ClusterWorkloadIdentityConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeClusterWorkloadIdentityConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ClusterWorkloadIdentityConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeClusterWorkloadIdentityConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewClusterWorkloadIdentityConfig(c *Client, des, nw *ClusterWorkloadIdentityConfig) *ClusterWorkloadIdentityConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for ClusterWorkloadIdentityConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.IssuerUri, nw.IssuerUri) {
		nw.IssuerUri = des.IssuerUri
	}
	if dcl.StringCanonicalize(des.WorkloadPool, nw.WorkloadPool) {
		nw.WorkloadPool = des.WorkloadPool
	}
	if dcl.StringCanonicalize(des.IdentityProvider, nw.IdentityProvider) {
		nw.IdentityProvider = des.IdentityProvider
	}

	return nw
}

func canonicalizeNewClusterWorkloadIdentityConfigSet(c *Client, des, nw []ClusterWorkloadIdentityConfig) []ClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []ClusterWorkloadIdentityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareClusterWorkloadIdentityConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewClusterWorkloadIdentityConfigSlice(c *Client, des, nw []ClusterWorkloadIdentityConfig) []ClusterWorkloadIdentityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ClusterWorkloadIdentityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewClusterWorkloadIdentityConfig(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffCluster(c *Client, desired, actual *Cluster, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AzureRegion, actual.AzureRegion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AzureRegion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceGroupId, actual.ResourceGroupId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceGroupId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Client, actual.Client, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AzureClient")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Networking, actual.Networking, dcl.Info{ObjectFunction: compareClusterNetworkingNewStyle, EmptyObject: EmptyClusterNetworking, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Networking")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ControlPlane, actual.ControlPlane, dcl.Info{ObjectFunction: compareClusterControlPlaneNewStyle, EmptyObject: EmptyClusterControlPlane, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ControlPlane")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Authorization, actual.Authorization, dcl.Info{ObjectFunction: compareClusterAuthorizationNewStyle, EmptyObject: EmptyClusterAuthorization, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Authorization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Reconciling, actual.Reconciling, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Reconciling")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadIdentityConfig, actual.WorkloadIdentityConfig, dcl.Info{OutputOnly: true, ObjectFunction: compareClusterWorkloadIdentityConfigNewStyle, EmptyObject: EmptyClusterWorkloadIdentityConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadIdentityConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareClusterNetworkingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterNetworking)
	if !ok {
		desiredNotPointer, ok := d.(ClusterNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworking or *ClusterNetworking", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterNetworking)
	if !ok {
		actualNotPointer, ok := a.(ClusterNetworking)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterNetworking", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.VirtualNetworkId, actual.VirtualNetworkId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VirtualNetworkId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PodAddressCidrBlocks, actual.PodAddressCidrBlocks, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PodAddressCidrBlocks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAddressCidrBlocks, actual.ServiceAddressCidrBlocks, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAddressCidrBlocks")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterControlPlaneNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterControlPlane)
	if !ok {
		desiredNotPointer, ok := d.(ClusterControlPlane)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlane or *ClusterControlPlane", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterControlPlane)
	if !ok {
		actualNotPointer, ok := a.(ClusterControlPlane)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlane", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubnetId, actual.SubnetId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubnetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VmSize, actual.VmSize, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VmSize")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SshConfig, actual.SshConfig, dcl.Info{ObjectFunction: compareClusterControlPlaneSshConfigNewStyle, EmptyObject: EmptyClusterControlPlaneSshConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SshConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RootVolume, actual.RootVolume, dcl.Info{ObjectFunction: compareClusterControlPlaneRootVolumeNewStyle, EmptyObject: EmptyClusterControlPlaneRootVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RootVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MainVolume, actual.MainVolume, dcl.Info{ObjectFunction: compareClusterControlPlaneMainVolumeNewStyle, EmptyObject: EmptyClusterControlPlaneMainVolume, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MainVolume")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatabaseEncryption, actual.DatabaseEncryption, dcl.Info{ObjectFunction: compareClusterControlPlaneDatabaseEncryptionNewStyle, EmptyObject: EmptyClusterControlPlaneDatabaseEncryption, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatabaseEncryption")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tags, actual.Tags, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tags")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterControlPlaneSshConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterControlPlaneSshConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterControlPlaneSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneSshConfig or *ClusterControlPlaneSshConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterControlPlaneSshConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterControlPlaneSshConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneSshConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AuthorizedKey, actual.AuthorizedKey, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthorizedKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterControlPlaneRootVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterControlPlaneRootVolume)
	if !ok {
		desiredNotPointer, ok := d.(ClusterControlPlaneRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneRootVolume or *ClusterControlPlaneRootVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterControlPlaneRootVolume)
	if !ok {
		actualNotPointer, ok := a.(ClusterControlPlaneRootVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneRootVolume", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SizeGib, actual.SizeGib, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGib")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterControlPlaneMainVolumeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterControlPlaneMainVolume)
	if !ok {
		desiredNotPointer, ok := d.(ClusterControlPlaneMainVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneMainVolume or *ClusterControlPlaneMainVolume", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterControlPlaneMainVolume)
	if !ok {
		actualNotPointer, ok := a.(ClusterControlPlaneMainVolume)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneMainVolume", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SizeGib, actual.SizeGib, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SizeGib")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterControlPlaneDatabaseEncryptionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterControlPlaneDatabaseEncryption)
	if !ok {
		desiredNotPointer, ok := d.(ClusterControlPlaneDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneDatabaseEncryption or *ClusterControlPlaneDatabaseEncryption", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterControlPlaneDatabaseEncryption)
	if !ok {
		actualNotPointer, ok := a.(ClusterControlPlaneDatabaseEncryption)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterControlPlaneDatabaseEncryption", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceGroupId, actual.ResourceGroupId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceGroupId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsKeyIdentifier, actual.KmsKeyIdentifier, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyIdentifier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAuthorizationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAuthorization)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAuthorization or *ClusterAuthorization", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAuthorization)
	if !ok {
		actualNotPointer, ok := a.(ClusterAuthorization)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAuthorization", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AdminUsers, actual.AdminUsers, dcl.Info{ObjectFunction: compareClusterAuthorizationAdminUsersNewStyle, EmptyObject: EmptyClusterAuthorizationAdminUsers, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdminUsers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterAuthorizationAdminUsersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterAuthorizationAdminUsers)
	if !ok {
		desiredNotPointer, ok := d.(ClusterAuthorizationAdminUsers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAuthorizationAdminUsers or *ClusterAuthorizationAdminUsers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterAuthorizationAdminUsers)
	if !ok {
		actualNotPointer, ok := a.(ClusterAuthorizationAdminUsers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterAuthorizationAdminUsers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Username, actual.Username, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Username")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareClusterWorkloadIdentityConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ClusterWorkloadIdentityConfig)
	if !ok {
		desiredNotPointer, ok := d.(ClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterWorkloadIdentityConfig or *ClusterWorkloadIdentityConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ClusterWorkloadIdentityConfig)
	if !ok {
		actualNotPointer, ok := a.(ClusterWorkloadIdentityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ClusterWorkloadIdentityConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IssuerUri, actual.IssuerUri, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IssuerUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadPool, actual.WorkloadPool, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadPool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IdentityProvider, actual.IdentityProvider, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdentityProvider")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Cluster) urlNormalized() *Cluster {
	normalized := dcl.Copy(*r).(Cluster)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.AzureRegion = dcl.SelfLinkToName(r.AzureRegion)
	normalized.ResourceGroupId = dcl.SelfLinkToName(r.ResourceGroupId)
	normalized.Client = dcl.SelfLinkToName(r.Client)
	normalized.Endpoint = dcl.SelfLinkToName(r.Endpoint)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Cluster) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Cluster resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Cluster) marshal(c *Client) ([]byte, error) {
	m, err := expandCluster(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Cluster: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalCluster decodes JSON responses into the Cluster resource schema.
func unmarshalCluster(b []byte, c *Client) (*Cluster, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCluster(m, c)
}

func unmarshalMapCluster(m map[string]interface{}, c *Client) (*Cluster, error) {

	return flattenCluster(c, m), nil
}

// expandCluster expands Cluster into a JSON request object.
func expandCluster(c *Client, f *Cluster) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/azureClusters/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.AzureRegion; !dcl.IsEmptyValueIndirect(v) {
		m["azureRegion"] = v
	}
	if v := f.ResourceGroupId; !dcl.IsEmptyValueIndirect(v) {
		m["resourceGroupId"] = v
	}
	if v := f.Client; !dcl.IsEmptyValueIndirect(v) {
		m["azureClient"] = v
	}
	if v, err := expandClusterNetworking(c, f.Networking); err != nil {
		return nil, fmt.Errorf("error expanding Networking into networking: %w", err)
	} else if v != nil {
		m["networking"] = v
	}
	if v, err := expandClusterControlPlane(c, f.ControlPlane); err != nil {
		return nil, fmt.Errorf("error expanding ControlPlane into controlPlane: %w", err)
	} else if v != nil {
		m["controlPlane"] = v
	}
	if v, err := expandClusterAuthorization(c, f.Authorization); err != nil {
		return nil, fmt.Errorf("error expanding Authorization into authorization: %w", err)
	} else if v != nil {
		m["authorization"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		m["annotations"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenCluster flattens Cluster from a JSON request object into the
// Cluster type.
func flattenCluster(c *Client, i interface{}) *Cluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Cluster{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.AzureRegion = dcl.FlattenString(m["azureRegion"])
	res.ResourceGroupId = dcl.FlattenString(m["resourceGroupId"])
	res.Client = dcl.FlattenString(m["azureClient"])
	res.Networking = flattenClusterNetworking(c, m["networking"])
	res.ControlPlane = flattenClusterControlPlane(c, m["controlPlane"])
	res.Authorization = flattenClusterAuthorization(c, m["authorization"])
	res.State = flattenClusterStateEnum(m["state"])
	res.Endpoint = dcl.FlattenString(m["endpoint"])
	res.Uid = dcl.FlattenString(m["uid"])
	res.Reconciling = dcl.FlattenBool(m["reconciling"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Etag = dcl.FlattenString(m["etag"])
	res.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	res.WorkloadIdentityConfig = flattenClusterWorkloadIdentityConfig(c, m["workloadIdentityConfig"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandClusterNetworkingMap expands the contents of ClusterNetworking into a JSON
// request object.
func expandClusterNetworkingMap(c *Client, f map[string]ClusterNetworking) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterNetworking(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterNetworkingSlice expands the contents of ClusterNetworking into a JSON
// request object.
func expandClusterNetworkingSlice(c *Client, f []ClusterNetworking) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterNetworking(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterNetworkingMap flattens the contents of ClusterNetworking from a JSON
// response object.
func flattenClusterNetworkingMap(c *Client, i interface{}) map[string]ClusterNetworking {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterNetworking{}
	}

	if len(a) == 0 {
		return map[string]ClusterNetworking{}
	}

	items := make(map[string]ClusterNetworking)
	for k, item := range a {
		items[k] = *flattenClusterNetworking(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterNetworkingSlice flattens the contents of ClusterNetworking from a JSON
// response object.
func flattenClusterNetworkingSlice(c *Client, i interface{}) []ClusterNetworking {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterNetworking{}
	}

	if len(a) == 0 {
		return []ClusterNetworking{}
	}

	items := make([]ClusterNetworking, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterNetworking(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterNetworking expands an instance of ClusterNetworking into a JSON
// request object.
func expandClusterNetworking(c *Client, f *ClusterNetworking) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.VirtualNetworkId; !dcl.IsEmptyValueIndirect(v) {
		m["virtualNetworkId"] = v
	}
	if v := f.PodAddressCidrBlocks; v != nil {
		m["podAddressCidrBlocks"] = v
	}
	if v := f.ServiceAddressCidrBlocks; v != nil {
		m["serviceAddressCidrBlocks"] = v
	}

	return m, nil
}

// flattenClusterNetworking flattens an instance of ClusterNetworking from a JSON
// response object.
func flattenClusterNetworking(c *Client, i interface{}) *ClusterNetworking {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterNetworking{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterNetworking
	}
	r.VirtualNetworkId = dcl.FlattenString(m["virtualNetworkId"])
	r.PodAddressCidrBlocks = dcl.FlattenStringSlice(m["podAddressCidrBlocks"])
	r.ServiceAddressCidrBlocks = dcl.FlattenStringSlice(m["serviceAddressCidrBlocks"])

	return r
}

// expandClusterControlPlaneMap expands the contents of ClusterControlPlane into a JSON
// request object.
func expandClusterControlPlaneMap(c *Client, f map[string]ClusterControlPlane) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterControlPlane(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterControlPlaneSlice expands the contents of ClusterControlPlane into a JSON
// request object.
func expandClusterControlPlaneSlice(c *Client, f []ClusterControlPlane) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterControlPlane(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterControlPlaneMap flattens the contents of ClusterControlPlane from a JSON
// response object.
func flattenClusterControlPlaneMap(c *Client, i interface{}) map[string]ClusterControlPlane {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterControlPlane{}
	}

	if len(a) == 0 {
		return map[string]ClusterControlPlane{}
	}

	items := make(map[string]ClusterControlPlane)
	for k, item := range a {
		items[k] = *flattenClusterControlPlane(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterControlPlaneSlice flattens the contents of ClusterControlPlane from a JSON
// response object.
func flattenClusterControlPlaneSlice(c *Client, i interface{}) []ClusterControlPlane {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterControlPlane{}
	}

	if len(a) == 0 {
		return []ClusterControlPlane{}
	}

	items := make([]ClusterControlPlane, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterControlPlane(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterControlPlane expands an instance of ClusterControlPlane into a JSON
// request object.
func expandClusterControlPlane(c *Client, f *ClusterControlPlane) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v := f.SubnetId; !dcl.IsEmptyValueIndirect(v) {
		m["subnetId"] = v
	}
	if v := f.VmSize; !dcl.IsEmptyValueIndirect(v) {
		m["vmSize"] = v
	}
	if v, err := expandClusterControlPlaneSshConfig(c, f.SshConfig); err != nil {
		return nil, fmt.Errorf("error expanding SshConfig into sshConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sshConfig"] = v
	}
	if v, err := expandClusterControlPlaneRootVolume(c, f.RootVolume); err != nil {
		return nil, fmt.Errorf("error expanding RootVolume into rootVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["rootVolume"] = v
	}
	if v, err := expandClusterControlPlaneMainVolume(c, f.MainVolume); err != nil {
		return nil, fmt.Errorf("error expanding MainVolume into mainVolume: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mainVolume"] = v
	}
	if v, err := expandClusterControlPlaneDatabaseEncryption(c, f.DatabaseEncryption); err != nil {
		return nil, fmt.Errorf("error expanding DatabaseEncryption into databaseEncryption: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["databaseEncryption"] = v
	}
	if v := f.Tags; !dcl.IsEmptyValueIndirect(v) {
		m["tags"] = v
	}

	return m, nil
}

// flattenClusterControlPlane flattens an instance of ClusterControlPlane from a JSON
// response object.
func flattenClusterControlPlane(c *Client, i interface{}) *ClusterControlPlane {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterControlPlane{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterControlPlane
	}
	r.Version = dcl.FlattenString(m["version"])
	r.SubnetId = dcl.FlattenString(m["subnetId"])
	r.VmSize = dcl.FlattenString(m["vmSize"])
	r.SshConfig = flattenClusterControlPlaneSshConfig(c, m["sshConfig"])
	r.RootVolume = flattenClusterControlPlaneRootVolume(c, m["rootVolume"])
	r.MainVolume = flattenClusterControlPlaneMainVolume(c, m["mainVolume"])
	r.DatabaseEncryption = flattenClusterControlPlaneDatabaseEncryption(c, m["databaseEncryption"])
	r.Tags = dcl.FlattenKeyValuePairs(m["tags"])

	return r
}

// expandClusterControlPlaneSshConfigMap expands the contents of ClusterControlPlaneSshConfig into a JSON
// request object.
func expandClusterControlPlaneSshConfigMap(c *Client, f map[string]ClusterControlPlaneSshConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterControlPlaneSshConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterControlPlaneSshConfigSlice expands the contents of ClusterControlPlaneSshConfig into a JSON
// request object.
func expandClusterControlPlaneSshConfigSlice(c *Client, f []ClusterControlPlaneSshConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterControlPlaneSshConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterControlPlaneSshConfigMap flattens the contents of ClusterControlPlaneSshConfig from a JSON
// response object.
func flattenClusterControlPlaneSshConfigMap(c *Client, i interface{}) map[string]ClusterControlPlaneSshConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterControlPlaneSshConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterControlPlaneSshConfig{}
	}

	items := make(map[string]ClusterControlPlaneSshConfig)
	for k, item := range a {
		items[k] = *flattenClusterControlPlaneSshConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterControlPlaneSshConfigSlice flattens the contents of ClusterControlPlaneSshConfig from a JSON
// response object.
func flattenClusterControlPlaneSshConfigSlice(c *Client, i interface{}) []ClusterControlPlaneSshConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterControlPlaneSshConfig{}
	}

	if len(a) == 0 {
		return []ClusterControlPlaneSshConfig{}
	}

	items := make([]ClusterControlPlaneSshConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterControlPlaneSshConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterControlPlaneSshConfig expands an instance of ClusterControlPlaneSshConfig into a JSON
// request object.
func expandClusterControlPlaneSshConfig(c *Client, f *ClusterControlPlaneSshConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AuthorizedKey; !dcl.IsEmptyValueIndirect(v) {
		m["authorizedKey"] = v
	}

	return m, nil
}

// flattenClusterControlPlaneSshConfig flattens an instance of ClusterControlPlaneSshConfig from a JSON
// response object.
func flattenClusterControlPlaneSshConfig(c *Client, i interface{}) *ClusterControlPlaneSshConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterControlPlaneSshConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterControlPlaneSshConfig
	}
	r.AuthorizedKey = dcl.FlattenString(m["authorizedKey"])

	return r
}

// expandClusterControlPlaneRootVolumeMap expands the contents of ClusterControlPlaneRootVolume into a JSON
// request object.
func expandClusterControlPlaneRootVolumeMap(c *Client, f map[string]ClusterControlPlaneRootVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterControlPlaneRootVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterControlPlaneRootVolumeSlice expands the contents of ClusterControlPlaneRootVolume into a JSON
// request object.
func expandClusterControlPlaneRootVolumeSlice(c *Client, f []ClusterControlPlaneRootVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterControlPlaneRootVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterControlPlaneRootVolumeMap flattens the contents of ClusterControlPlaneRootVolume from a JSON
// response object.
func flattenClusterControlPlaneRootVolumeMap(c *Client, i interface{}) map[string]ClusterControlPlaneRootVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterControlPlaneRootVolume{}
	}

	if len(a) == 0 {
		return map[string]ClusterControlPlaneRootVolume{}
	}

	items := make(map[string]ClusterControlPlaneRootVolume)
	for k, item := range a {
		items[k] = *flattenClusterControlPlaneRootVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterControlPlaneRootVolumeSlice flattens the contents of ClusterControlPlaneRootVolume from a JSON
// response object.
func flattenClusterControlPlaneRootVolumeSlice(c *Client, i interface{}) []ClusterControlPlaneRootVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterControlPlaneRootVolume{}
	}

	if len(a) == 0 {
		return []ClusterControlPlaneRootVolume{}
	}

	items := make([]ClusterControlPlaneRootVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterControlPlaneRootVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterControlPlaneRootVolume expands an instance of ClusterControlPlaneRootVolume into a JSON
// request object.
func expandClusterControlPlaneRootVolume(c *Client, f *ClusterControlPlaneRootVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}

	return m, nil
}

// flattenClusterControlPlaneRootVolume flattens an instance of ClusterControlPlaneRootVolume from a JSON
// response object.
func flattenClusterControlPlaneRootVolume(c *Client, i interface{}) *ClusterControlPlaneRootVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterControlPlaneRootVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterControlPlaneRootVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])

	return r
}

// expandClusterControlPlaneMainVolumeMap expands the contents of ClusterControlPlaneMainVolume into a JSON
// request object.
func expandClusterControlPlaneMainVolumeMap(c *Client, f map[string]ClusterControlPlaneMainVolume) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterControlPlaneMainVolume(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterControlPlaneMainVolumeSlice expands the contents of ClusterControlPlaneMainVolume into a JSON
// request object.
func expandClusterControlPlaneMainVolumeSlice(c *Client, f []ClusterControlPlaneMainVolume) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterControlPlaneMainVolume(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterControlPlaneMainVolumeMap flattens the contents of ClusterControlPlaneMainVolume from a JSON
// response object.
func flattenClusterControlPlaneMainVolumeMap(c *Client, i interface{}) map[string]ClusterControlPlaneMainVolume {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterControlPlaneMainVolume{}
	}

	if len(a) == 0 {
		return map[string]ClusterControlPlaneMainVolume{}
	}

	items := make(map[string]ClusterControlPlaneMainVolume)
	for k, item := range a {
		items[k] = *flattenClusterControlPlaneMainVolume(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterControlPlaneMainVolumeSlice flattens the contents of ClusterControlPlaneMainVolume from a JSON
// response object.
func flattenClusterControlPlaneMainVolumeSlice(c *Client, i interface{}) []ClusterControlPlaneMainVolume {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterControlPlaneMainVolume{}
	}

	if len(a) == 0 {
		return []ClusterControlPlaneMainVolume{}
	}

	items := make([]ClusterControlPlaneMainVolume, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterControlPlaneMainVolume(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterControlPlaneMainVolume expands an instance of ClusterControlPlaneMainVolume into a JSON
// request object.
func expandClusterControlPlaneMainVolume(c *Client, f *ClusterControlPlaneMainVolume) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SizeGib; !dcl.IsEmptyValueIndirect(v) {
		m["sizeGib"] = v
	}

	return m, nil
}

// flattenClusterControlPlaneMainVolume flattens an instance of ClusterControlPlaneMainVolume from a JSON
// response object.
func flattenClusterControlPlaneMainVolume(c *Client, i interface{}) *ClusterControlPlaneMainVolume {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterControlPlaneMainVolume{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterControlPlaneMainVolume
	}
	r.SizeGib = dcl.FlattenInteger(m["sizeGib"])

	return r
}

// expandClusterControlPlaneDatabaseEncryptionMap expands the contents of ClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandClusterControlPlaneDatabaseEncryptionMap(c *Client, f map[string]ClusterControlPlaneDatabaseEncryption) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterControlPlaneDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterControlPlaneDatabaseEncryptionSlice expands the contents of ClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandClusterControlPlaneDatabaseEncryptionSlice(c *Client, f []ClusterControlPlaneDatabaseEncryption) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterControlPlaneDatabaseEncryption(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterControlPlaneDatabaseEncryptionMap flattens the contents of ClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenClusterControlPlaneDatabaseEncryptionMap(c *Client, i interface{}) map[string]ClusterControlPlaneDatabaseEncryption {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterControlPlaneDatabaseEncryption{}
	}

	if len(a) == 0 {
		return map[string]ClusterControlPlaneDatabaseEncryption{}
	}

	items := make(map[string]ClusterControlPlaneDatabaseEncryption)
	for k, item := range a {
		items[k] = *flattenClusterControlPlaneDatabaseEncryption(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterControlPlaneDatabaseEncryptionSlice flattens the contents of ClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenClusterControlPlaneDatabaseEncryptionSlice(c *Client, i interface{}) []ClusterControlPlaneDatabaseEncryption {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterControlPlaneDatabaseEncryption{}
	}

	if len(a) == 0 {
		return []ClusterControlPlaneDatabaseEncryption{}
	}

	items := make([]ClusterControlPlaneDatabaseEncryption, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterControlPlaneDatabaseEncryption(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterControlPlaneDatabaseEncryption expands an instance of ClusterControlPlaneDatabaseEncryption into a JSON
// request object.
func expandClusterControlPlaneDatabaseEncryption(c *Client, f *ClusterControlPlaneDatabaseEncryption) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceGroupId; !dcl.IsEmptyValueIndirect(v) {
		m["resourceGroupId"] = v
	}
	if v := f.KmsKeyIdentifier; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyIdentifier"] = v
	}

	return m, nil
}

// flattenClusterControlPlaneDatabaseEncryption flattens an instance of ClusterControlPlaneDatabaseEncryption from a JSON
// response object.
func flattenClusterControlPlaneDatabaseEncryption(c *Client, i interface{}) *ClusterControlPlaneDatabaseEncryption {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterControlPlaneDatabaseEncryption{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterControlPlaneDatabaseEncryption
	}
	r.ResourceGroupId = dcl.FlattenString(m["resourceGroupId"])
	r.KmsKeyIdentifier = dcl.FlattenString(m["kmsKeyIdentifier"])

	return r
}

// expandClusterAuthorizationMap expands the contents of ClusterAuthorization into a JSON
// request object.
func expandClusterAuthorizationMap(c *Client, f map[string]ClusterAuthorization) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAuthorization(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAuthorizationSlice expands the contents of ClusterAuthorization into a JSON
// request object.
func expandClusterAuthorizationSlice(c *Client, f []ClusterAuthorization) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAuthorization(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAuthorizationMap flattens the contents of ClusterAuthorization from a JSON
// response object.
func flattenClusterAuthorizationMap(c *Client, i interface{}) map[string]ClusterAuthorization {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAuthorization{}
	}

	if len(a) == 0 {
		return map[string]ClusterAuthorization{}
	}

	items := make(map[string]ClusterAuthorization)
	for k, item := range a {
		items[k] = *flattenClusterAuthorization(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAuthorizationSlice flattens the contents of ClusterAuthorization from a JSON
// response object.
func flattenClusterAuthorizationSlice(c *Client, i interface{}) []ClusterAuthorization {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAuthorization{}
	}

	if len(a) == 0 {
		return []ClusterAuthorization{}
	}

	items := make([]ClusterAuthorization, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAuthorization(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAuthorization expands an instance of ClusterAuthorization into a JSON
// request object.
func expandClusterAuthorization(c *Client, f *ClusterAuthorization) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandClusterAuthorizationAdminUsersSlice(c, f.AdminUsers); err != nil {
		return nil, fmt.Errorf("error expanding AdminUsers into adminUsers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["adminUsers"] = v
	}

	return m, nil
}

// flattenClusterAuthorization flattens an instance of ClusterAuthorization from a JSON
// response object.
func flattenClusterAuthorization(c *Client, i interface{}) *ClusterAuthorization {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAuthorization{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterAuthorization
	}
	r.AdminUsers = flattenClusterAuthorizationAdminUsersSlice(c, m["adminUsers"])

	return r
}

// expandClusterAuthorizationAdminUsersMap expands the contents of ClusterAuthorizationAdminUsers into a JSON
// request object.
func expandClusterAuthorizationAdminUsersMap(c *Client, f map[string]ClusterAuthorizationAdminUsers) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterAuthorizationAdminUsers(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterAuthorizationAdminUsersSlice expands the contents of ClusterAuthorizationAdminUsers into a JSON
// request object.
func expandClusterAuthorizationAdminUsersSlice(c *Client, f []ClusterAuthorizationAdminUsers) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterAuthorizationAdminUsers(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterAuthorizationAdminUsersMap flattens the contents of ClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenClusterAuthorizationAdminUsersMap(c *Client, i interface{}) map[string]ClusterAuthorizationAdminUsers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterAuthorizationAdminUsers{}
	}

	if len(a) == 0 {
		return map[string]ClusterAuthorizationAdminUsers{}
	}

	items := make(map[string]ClusterAuthorizationAdminUsers)
	for k, item := range a {
		items[k] = *flattenClusterAuthorizationAdminUsers(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterAuthorizationAdminUsersSlice flattens the contents of ClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenClusterAuthorizationAdminUsersSlice(c *Client, i interface{}) []ClusterAuthorizationAdminUsers {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterAuthorizationAdminUsers{}
	}

	if len(a) == 0 {
		return []ClusterAuthorizationAdminUsers{}
	}

	items := make([]ClusterAuthorizationAdminUsers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterAuthorizationAdminUsers(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterAuthorizationAdminUsers expands an instance of ClusterAuthorizationAdminUsers into a JSON
// request object.
func expandClusterAuthorizationAdminUsers(c *Client, f *ClusterAuthorizationAdminUsers) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}

	return m, nil
}

// flattenClusterAuthorizationAdminUsers flattens an instance of ClusterAuthorizationAdminUsers from a JSON
// response object.
func flattenClusterAuthorizationAdminUsers(c *Client, i interface{}) *ClusterAuthorizationAdminUsers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterAuthorizationAdminUsers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterAuthorizationAdminUsers
	}
	r.Username = dcl.FlattenString(m["username"])

	return r
}

// expandClusterWorkloadIdentityConfigMap expands the contents of ClusterWorkloadIdentityConfig into a JSON
// request object.
func expandClusterWorkloadIdentityConfigMap(c *Client, f map[string]ClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandClusterWorkloadIdentityConfigSlice expands the contents of ClusterWorkloadIdentityConfig into a JSON
// request object.
func expandClusterWorkloadIdentityConfigSlice(c *Client, f []ClusterWorkloadIdentityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandClusterWorkloadIdentityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenClusterWorkloadIdentityConfigMap flattens the contents of ClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenClusterWorkloadIdentityConfigMap(c *Client, i interface{}) map[string]ClusterWorkloadIdentityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return map[string]ClusterWorkloadIdentityConfig{}
	}

	items := make(map[string]ClusterWorkloadIdentityConfig)
	for k, item := range a {
		items[k] = *flattenClusterWorkloadIdentityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenClusterWorkloadIdentityConfigSlice flattens the contents of ClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenClusterWorkloadIdentityConfigSlice(c *Client, i interface{}) []ClusterWorkloadIdentityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterWorkloadIdentityConfig{}
	}

	if len(a) == 0 {
		return []ClusterWorkloadIdentityConfig{}
	}

	items := make([]ClusterWorkloadIdentityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterWorkloadIdentityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandClusterWorkloadIdentityConfig expands an instance of ClusterWorkloadIdentityConfig into a JSON
// request object.
func expandClusterWorkloadIdentityConfig(c *Client, f *ClusterWorkloadIdentityConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IssuerUri; !dcl.IsEmptyValueIndirect(v) {
		m["issuerUri"] = v
	}
	if v := f.WorkloadPool; !dcl.IsEmptyValueIndirect(v) {
		m["workloadPool"] = v
	}
	if v := f.IdentityProvider; !dcl.IsEmptyValueIndirect(v) {
		m["identityProvider"] = v
	}

	return m, nil
}

// flattenClusterWorkloadIdentityConfig flattens an instance of ClusterWorkloadIdentityConfig from a JSON
// response object.
func flattenClusterWorkloadIdentityConfig(c *Client, i interface{}) *ClusterWorkloadIdentityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ClusterWorkloadIdentityConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyClusterWorkloadIdentityConfig
	}
	r.IssuerUri = dcl.FlattenString(m["issuerUri"])
	r.WorkloadPool = dcl.FlattenString(m["workloadPool"])
	r.IdentityProvider = dcl.FlattenString(m["identityProvider"])

	return r
}

// flattenClusterStateEnumMap flattens the contents of ClusterStateEnum from a JSON
// response object.
func flattenClusterStateEnumMap(c *Client, i interface{}) map[string]ClusterStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ClusterStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ClusterStateEnum{}
	}

	items := make(map[string]ClusterStateEnum)
	for k, item := range a {
		items[k] = *flattenClusterStateEnum(item.(interface{}))
	}

	return items
}

// flattenClusterStateEnumSlice flattens the contents of ClusterStateEnum from a JSON
// response object.
func flattenClusterStateEnumSlice(c *Client, i interface{}) []ClusterStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ClusterStateEnum{}
	}

	if len(a) == 0 {
		return []ClusterStateEnum{}
	}

	items := make([]ClusterStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenClusterStateEnum(item.(interface{})))
	}

	return items
}

// flattenClusterStateEnum asserts that an interface is a string, and returns a
// pointer to a *ClusterStateEnum with the same value as that string.
func flattenClusterStateEnum(i interface{}) *ClusterStateEnum {
	s, ok := i.(string)
	if !ok {
		return ClusterStateEnumRef("")
	}

	return ClusterStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Cluster) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCluster(b, c)
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
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

type clusterDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         clusterApiOperation
}

func convertFieldDiffsToClusterDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]clusterDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []clusterDiff
	// For each operation name, create a clusterDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := clusterDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToClusterApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToClusterApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (clusterApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractClusterFields(r *Cluster) error {
	return nil
}