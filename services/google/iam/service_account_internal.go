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
package iam

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

func (r *ServiceAccount) validate() error {

	if !dcl.IsEmptyValueIndirect(r.ActasResources) {
		if err := r.ActasResources.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ServiceAccountActasResources) validate() error {
	return nil
}
func (r *ServiceAccountActasResourcesResources) validate() error {
	return nil
}

func serviceAccountGetURL(userBasePath string, r *ServiceAccount) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com", "https://iam.googleapis.com/v1/", userBasePath, params), nil
}

func serviceAccountListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/serviceAccounts", "https://iam.googleapis.com/v1/", userBasePath, params), nil

}

func serviceAccountCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/serviceAccounts", "https://iam.googleapis.com/v1/", userBasePath, params), nil

}

func serviceAccountDeleteURL(userBasePath string, r *ServiceAccount) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com", "https://iam.googleapis.com/v1/", userBasePath, params), nil
}

func (r *ServiceAccount) SetPolicyURL(userBasePath string) string {
	n := r.URLNormalized()
	fields := map[string]interface{}{
		"project": *n.Project,
		"name":    *n.Name,
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com:setIamPolicy", "https://iam.googleapis.com/v1/", userBasePath, fields)
}

func (r *ServiceAccount) SetPolicyVerb() string {
	return "POST"
}

func (r *ServiceAccount) getPolicyURL(userBasePath string) string {
	n := r.URLNormalized()
	fields := map[string]interface{}{
		"project": *n.Project,
		"name":    *n.Name,
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com:getIamPolicy", "https://iam.googleapis.com/v1/", userBasePath, fields)
}

func (r *ServiceAccount) IAMPolicyVersion() int {
	return 3
}

// serviceAccountApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type serviceAccountApiOperation interface {
	do(context.Context, *ServiceAccount, *Client) error
}

// newUpdateServiceAccountPatchServiceAccountRequest creates a request for an
// ServiceAccount resource's PatchServiceAccount update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateServiceAccountPatchServiceAccountRequest(ctx context.Context, f *ServiceAccount, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	return req, nil
}

// marshalUpdateServiceAccountPatchServiceAccountRequest converts the update into
// the final JSON request body.
func marshalUpdateServiceAccountPatchServiceAccountRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"email"},
		[]string{},
	)
	return json.Marshal(m)
}

type updateServiceAccountPatchServiceAccountOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	Diffs        []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateServiceAccountPatchServiceAccountOperation) do(ctx context.Context, r *ServiceAccount, c *Client) error {
	_, err := c.GetServiceAccount(ctx, r.URLNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchServiceAccount")
	if err != nil {
		return err
	}

	req, err := newUpdateServiceAccountPatchServiceAccountRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateServiceAccountPatchServiceAccountRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://iam.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listServiceAccountRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := serviceAccountListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ServiceAccountMaxPage {
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

type listServiceAccountOperation struct {
	Accounts []map[string]interface{} `json:"accounts"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listServiceAccount(ctx context.Context, project, pageToken string, pageSize int32) ([]*ServiceAccount, string, error) {
	b, err := c.listServiceAccountRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listServiceAccountOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*ServiceAccount
	for _, v := range m.Accounts {
		res, err := unmarshalMapServiceAccount(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllServiceAccount(ctx context.Context, f func(*ServiceAccount) bool, resources []*ServiceAccount) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteServiceAccount(ctx, res)
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

type deleteServiceAccountOperation struct{}

func (op *deleteServiceAccountOperation) do(ctx context.Context, r *ServiceAccount, c *Client) error {
	r, err := c.GetServiceAccount(ctx, r.URLNormalized())
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.Infof("ServiceAccount not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.Warningf("GetServiceAccount checking for existence. error: %v", err)
		return err
	}

	u, err := serviceAccountDeleteURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete ServiceAccount: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetServiceAccount(ctx, r.URLNormalized())
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
type createServiceAccountOperation struct {
	response map[string]interface{}
}

func (op *createServiceAccountOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createServiceAccountOperation) do(ctx context.Context, r *ServiceAccount, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := serviceAccountCreateURL(c.Config.BasePath, project)

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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetServiceAccount(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getServiceAccountRaw(ctx context.Context, r *ServiceAccount) ([]byte, error) {

	u, err := serviceAccountGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) serviceAccountDiffsForRawDesired(ctx context.Context, rawDesired *ServiceAccount, opts ...dcl.ApplyOption) (initial, desired *ServiceAccount, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *ServiceAccount
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*ServiceAccount); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected ServiceAccount, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetServiceAccount(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a ServiceAccount resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve ServiceAccount resource: %v", err)
		}
		c.Config.Logger.Info("Found that ServiceAccount resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeServiceAccountDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for ServiceAccount: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for ServiceAccount: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeServiceAccountInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for ServiceAccount: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeServiceAccountDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for ServiceAccount: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffServiceAccount(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeServiceAccountInitialState(rawInitial, rawDesired *ServiceAccount) (*ServiceAccount, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeServiceAccountDesiredState(rawDesired, rawInitial *ServiceAccount, opts ...dcl.ApplyOption) (*ServiceAccount, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.ActasResources = canonicalizeServiceAccountActasResources(rawDesired.ActasResources, nil, opts...)

		return rawDesired, nil
	}

	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		rawDesired.DisplayName = rawInitial.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		rawDesired.Description = rawInitial.Description
	}
	rawDesired.ActasResources = canonicalizeServiceAccountActasResources(rawDesired.ActasResources, rawInitial.ActasResources, opts...)

	return rawDesired, nil
}

func canonicalizeServiceAccountNewState(c *Client, rawNew, rawDesired *ServiceAccount) (*ServiceAccount, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
		if dcl.NameToSelfLink(rawDesired.Project, rawNew.Project) {
			rawNew.Project = rawDesired.Project
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.UniqueId) && dcl.IsEmptyValueIndirect(rawDesired.UniqueId) {
		rawNew.UniqueId = rawDesired.UniqueId
	} else {
		if dcl.StringCanonicalize(rawDesired.UniqueId, rawNew.UniqueId) {
			rawNew.UniqueId = rawDesired.UniqueId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Email) && dcl.IsEmptyValueIndirect(rawDesired.Email) {
		rawNew.Email = rawDesired.Email
	} else {
		if dcl.StringCanonicalize(rawDesired.Email, rawNew.Email) {
			rawNew.Email = rawDesired.Email
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.OAuth2ClientId) && dcl.IsEmptyValueIndirect(rawDesired.OAuth2ClientId) {
		rawNew.OAuth2ClientId = rawDesired.OAuth2ClientId
	} else {
		if dcl.StringCanonicalize(rawDesired.OAuth2ClientId, rawNew.OAuth2ClientId) {
			rawNew.OAuth2ClientId = rawDesired.OAuth2ClientId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ActasResources) && dcl.IsEmptyValueIndirect(rawDesired.ActasResources) {
		rawNew.ActasResources = rawDesired.ActasResources
	} else {
		rawNew.ActasResources = canonicalizeNewServiceAccountActasResources(c, rawDesired.ActasResources, rawNew.ActasResources)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Disabled) && dcl.IsEmptyValueIndirect(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	return rawNew, nil
}

func canonicalizeServiceAccountActasResources(des, initial *ServiceAccountActasResources, opts ...dcl.ApplyOption) *ServiceAccountActasResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.Resources) {
		des.Resources = initial.Resources
	}

	return des
}

func canonicalizeNewServiceAccountActasResources(c *Client, des, nw *ServiceAccountActasResources) *ServiceAccountActasResources {
	if des == nil || nw == nil {
		return nw
	}

	nw.Resources = canonicalizeNewServiceAccountActasResourcesResourcesSlice(c, des.Resources, nw.Resources)

	return nw
}

func canonicalizeNewServiceAccountActasResourcesSet(c *Client, des, nw []ServiceAccountActasResources) []ServiceAccountActasResources {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceAccountActasResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceAccountActasResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceAccountActasResourcesSlice(c *Client, des, nw []ServiceAccountActasResources) []ServiceAccountActasResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceAccountActasResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceAccountActasResources(c, &d, &n))
	}

	return items
}

func canonicalizeServiceAccountActasResourcesResources(des, initial *ServiceAccountActasResourcesResources, opts ...dcl.ApplyOption) *ServiceAccountActasResourcesResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.FullResourceName, initial.FullResourceName) || dcl.IsZeroValue(des.FullResourceName) {
		des.FullResourceName = initial.FullResourceName
	}

	return des
}

func canonicalizeNewServiceAccountActasResourcesResources(c *Client, des, nw *ServiceAccountActasResourcesResources) *ServiceAccountActasResourcesResources {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.FullResourceName, nw.FullResourceName) {
		nw.FullResourceName = des.FullResourceName
	}

	return nw
}

func canonicalizeNewServiceAccountActasResourcesResourcesSet(c *Client, des, nw []ServiceAccountActasResourcesResources) []ServiceAccountActasResourcesResources {
	if des == nil {
		return nw
	}
	var reorderedNew []ServiceAccountActasResourcesResources
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareServiceAccountActasResourcesResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewServiceAccountActasResourcesResourcesSlice(c *Client, des, nw []ServiceAccountActasResourcesResources) []ServiceAccountActasResourcesResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ServiceAccountActasResourcesResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewServiceAccountActasResourcesResources(c, &d, &n))
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
func diffServiceAccount(c *Client, desired, actual *ServiceAccount, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{OutputOnly: true, Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UniqueId, actual.UniqueId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UniqueId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Email, actual.Email, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Email")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceAccountPatchServiceAccountOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateServiceAccountPatchServiceAccountOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OAuth2ClientId, actual.OAuth2ClientId, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oauth2ClientId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ActasResources, actual.ActasResources, dcl.Info{ObjectFunction: compareServiceAccountActasResourcesNewStyle, EmptyObject: EmptyServiceAccountActasResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ActasResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareServiceAccountActasResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceAccountActasResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceAccountActasResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResources or *ServiceAccountActasResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceAccountActasResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceAccountActasResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.Info{ObjectFunction: compareServiceAccountActasResourcesResourcesNewStyle, EmptyObject: EmptyServiceAccountActasResourcesResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareServiceAccountActasResourcesResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ServiceAccountActasResourcesResources)
	if !ok {
		desiredNotPointer, ok := d.(ServiceAccountActasResourcesResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResourcesResources or *ServiceAccountActasResourcesResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ServiceAccountActasResourcesResources)
	if !ok {
		actualNotPointer, ok := a.(ServiceAccountActasResourcesResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ServiceAccountActasResourcesResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FullResourceName, actual.FullResourceName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullResourceName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *ServiceAccount) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServiceAccount) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *ServiceAccount) deleteFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *ServiceAccount) updateURL(userBasePath, updateName string) (string, error) {
	n := r.URLNormalized()
	if updateName == "PatchServiceAccount" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(n.Project),
			"name":    dcl.ValueOrEmptyString(n.Name),
		}
		return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com", "https://iam.googleapis.com/v1/", userBasePath, fields), nil

	}
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the ServiceAccount resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *ServiceAccount) marshal(c *Client) ([]byte, error) {
	m, err := expandServiceAccount(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ServiceAccount: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"email"},
		[]string{},
	)
	m = EncodeServiceAccountCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalServiceAccount decodes JSON responses into the ServiceAccount resource schema.
func unmarshalServiceAccount(b []byte, c *Client) (*ServiceAccount, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapServiceAccount(m, c)
}

func unmarshalMapServiceAccount(m map[string]interface{}, c *Client) (*ServiceAccount, error) {

	return flattenServiceAccount(c, m), nil
}

// expandServiceAccount expands ServiceAccount into a JSON request object.
func expandServiceAccount(c *Client, f *ServiceAccount) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccounts.com", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Project; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.UniqueId; !dcl.IsEmptyValueIndirect(v) {
		m["uniqueId"] = v
	}
	if v := f.Email; !dcl.IsEmptyValueIndirect(v) {
		m["email"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.OAuth2ClientId; !dcl.IsEmptyValueIndirect(v) {
		m["oauth2ClientId"] = v
	}
	if v, err := expandServiceAccountActasResources(c, f.ActasResources); err != nil {
		return nil, fmt.Errorf("error expanding ActasResources into actasResources: %w", err)
	} else if v != nil {
		m["actasResources"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		m["disabled"] = v
	}

	return m, nil
}

// flattenServiceAccount flattens ServiceAccount from a JSON request object into the
// ServiceAccount type.
func flattenServiceAccount(c *Client, i interface{}) *ServiceAccount {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &ServiceAccount{}
	res.Name = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["projectId"])
	res.UniqueId = dcl.FlattenString(m["uniqueId"])
	res.Email = dcl.FlattenString(m["email"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Description = dcl.FlattenString(m["description"])
	res.OAuth2ClientId = dcl.FlattenString(m["oauth2ClientId"])
	res.ActasResources = flattenServiceAccountActasResources(c, m["actasResources"])
	res.Disabled = dcl.FlattenBool(m["disabled"])

	return res
}

// expandServiceAccountActasResourcesMap expands the contents of ServiceAccountActasResources into a JSON
// request object.
func expandServiceAccountActasResourcesMap(c *Client, f map[string]ServiceAccountActasResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceAccountActasResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceAccountActasResourcesSlice expands the contents of ServiceAccountActasResources into a JSON
// request object.
func expandServiceAccountActasResourcesSlice(c *Client, f []ServiceAccountActasResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceAccountActasResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceAccountActasResourcesMap flattens the contents of ServiceAccountActasResources from a JSON
// response object.
func flattenServiceAccountActasResourcesMap(c *Client, i interface{}) map[string]ServiceAccountActasResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceAccountActasResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceAccountActasResources{}
	}

	items := make(map[string]ServiceAccountActasResources)
	for k, item := range a {
		items[k] = *flattenServiceAccountActasResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceAccountActasResourcesSlice flattens the contents of ServiceAccountActasResources from a JSON
// response object.
func flattenServiceAccountActasResourcesSlice(c *Client, i interface{}) []ServiceAccountActasResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceAccountActasResources{}
	}

	if len(a) == 0 {
		return []ServiceAccountActasResources{}
	}

	items := make([]ServiceAccountActasResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceAccountActasResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceAccountActasResources expands an instance of ServiceAccountActasResources into a JSON
// request object.
func expandServiceAccountActasResources(c *Client, f *ServiceAccountActasResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandServiceAccountActasResourcesResourcesSlice(c, f.Resources); err != nil {
		return nil, fmt.Errorf("error expanding Resources into resources: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resources"] = v
	}

	return m, nil
}

// flattenServiceAccountActasResources flattens an instance of ServiceAccountActasResources from a JSON
// response object.
func flattenServiceAccountActasResources(c *Client, i interface{}) *ServiceAccountActasResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceAccountActasResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceAccountActasResources
	}
	r.Resources = flattenServiceAccountActasResourcesResourcesSlice(c, m["resources"])

	return r
}

// expandServiceAccountActasResourcesResourcesMap expands the contents of ServiceAccountActasResourcesResources into a JSON
// request object.
func expandServiceAccountActasResourcesResourcesMap(c *Client, f map[string]ServiceAccountActasResourcesResources) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandServiceAccountActasResourcesResources(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandServiceAccountActasResourcesResourcesSlice expands the contents of ServiceAccountActasResourcesResources into a JSON
// request object.
func expandServiceAccountActasResourcesResourcesSlice(c *Client, f []ServiceAccountActasResourcesResources) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandServiceAccountActasResourcesResources(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenServiceAccountActasResourcesResourcesMap flattens the contents of ServiceAccountActasResourcesResources from a JSON
// response object.
func flattenServiceAccountActasResourcesResourcesMap(c *Client, i interface{}) map[string]ServiceAccountActasResourcesResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ServiceAccountActasResourcesResources{}
	}

	if len(a) == 0 {
		return map[string]ServiceAccountActasResourcesResources{}
	}

	items := make(map[string]ServiceAccountActasResourcesResources)
	for k, item := range a {
		items[k] = *flattenServiceAccountActasResourcesResources(c, item.(map[string]interface{}))
	}

	return items
}

// flattenServiceAccountActasResourcesResourcesSlice flattens the contents of ServiceAccountActasResourcesResources from a JSON
// response object.
func flattenServiceAccountActasResourcesResourcesSlice(c *Client, i interface{}) []ServiceAccountActasResourcesResources {
	a, ok := i.([]interface{})
	if !ok {
		return []ServiceAccountActasResourcesResources{}
	}

	if len(a) == 0 {
		return []ServiceAccountActasResourcesResources{}
	}

	items := make([]ServiceAccountActasResourcesResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenServiceAccountActasResourcesResources(c, item.(map[string]interface{})))
	}

	return items
}

// expandServiceAccountActasResourcesResources expands an instance of ServiceAccountActasResourcesResources into a JSON
// request object.
func expandServiceAccountActasResourcesResources(c *Client, f *ServiceAccountActasResourcesResources) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FullResourceName; !dcl.IsEmptyValueIndirect(v) {
		m["fullResourceName"] = v
	}

	return m, nil
}

// flattenServiceAccountActasResourcesResources flattens an instance of ServiceAccountActasResourcesResources from a JSON
// response object.
func flattenServiceAccountActasResourcesResources(c *Client, i interface{}) *ServiceAccountActasResourcesResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ServiceAccountActasResourcesResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyServiceAccountActasResourcesResources
	}
	r.FullResourceName = dcl.FlattenString(m["fullResourceName"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *ServiceAccount) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalServiceAccount(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
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

type serviceAccountDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         serviceAccountApiOperation
}

func convertFieldDiffToServiceAccountOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]serviceAccountDiff, error) {
	var diffs []serviceAccountDiff
	for _, op := range ops {
		diff := serviceAccountDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameToserviceAccountApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToserviceAccountApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (serviceAccountApiOperation, error) {
	switch op {

	case "updateServiceAccountPatchServiceAccountOperation":
		return &updateServiceAccountPatchServiceAccountOperation{Diffs: diffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
