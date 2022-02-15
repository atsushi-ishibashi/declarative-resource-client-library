// Copyright 2022 Google LLC. All Rights Reserved.
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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type ServerTlsPolicy struct {
	Name              *string                           `json:"name"`
	Description       *string                           `json:"description"`
	CreateTime        *string                           `json:"createTime"`
	UpdateTime        *string                           `json:"updateTime"`
	Labels            map[string]string                 `json:"labels"`
	AllowOpen         *bool                             `json:"allowOpen"`
	ServerCertificate *ServerTlsPolicyServerCertificate `json:"serverCertificate"`
	MtlsPolicy        *ServerTlsPolicyMtlsPolicy        `json:"mtlsPolicy"`
	Project           *string                           `json:"project"`
	Location          *string                           `json:"location"`
}

func (r *ServerTlsPolicy) String() string {
	return dcl.SprintResource(r)
}

type ServerTlsPolicyServerCertificate struct {
	empty                       bool                                                         `json:"-"`
	LocalFilepath               *ServerTlsPolicyServerCertificateLocalFilepath               `json:"localFilepath"`
	GrpcEndpoint                *ServerTlsPolicyServerCertificateGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ServerTlsPolicyServerCertificateCertificateProviderInstance `json:"certificateProviderInstance"`
}

type jsonServerTlsPolicyServerCertificate ServerTlsPolicyServerCertificate

func (r *ServerTlsPolicyServerCertificate) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyServerCertificate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyServerCertificate
	} else {

		r.LocalFilepath = res.LocalFilepath

		r.GrpcEndpoint = res.GrpcEndpoint

		r.CertificateProviderInstance = res.CertificateProviderInstance

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificate *ServerTlsPolicyServerCertificate = &ServerTlsPolicyServerCertificate{empty: true}

func (r *ServerTlsPolicyServerCertificate) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyServerCertificate) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyServerCertificateLocalFilepath struct {
	empty           bool    `json:"-"`
	CertificatePath *string `json:"certificatePath"`
	PrivateKeyPath  *string `json:"privateKeyPath"`
}

type jsonServerTlsPolicyServerCertificateLocalFilepath ServerTlsPolicyServerCertificateLocalFilepath

func (r *ServerTlsPolicyServerCertificateLocalFilepath) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyServerCertificateLocalFilepath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyServerCertificateLocalFilepath
	} else {

		r.CertificatePath = res.CertificatePath

		r.PrivateKeyPath = res.PrivateKeyPath

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificateLocalFilepath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificateLocalFilepath *ServerTlsPolicyServerCertificateLocalFilepath = &ServerTlsPolicyServerCertificateLocalFilepath{empty: true}

func (r *ServerTlsPolicyServerCertificateLocalFilepath) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyServerCertificateLocalFilepath) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificateLocalFilepath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyServerCertificateGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

type jsonServerTlsPolicyServerCertificateGrpcEndpoint ServerTlsPolicyServerCertificateGrpcEndpoint

func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyServerCertificateGrpcEndpoint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyServerCertificateGrpcEndpoint
	} else {

		r.TargetUri = res.TargetUri

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificateGrpcEndpoint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificateGrpcEndpoint *ServerTlsPolicyServerCertificateGrpcEndpoint = &ServerTlsPolicyServerCertificateGrpcEndpoint{empty: true}

func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificateGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyServerCertificateCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

type jsonServerTlsPolicyServerCertificateCertificateProviderInstance ServerTlsPolicyServerCertificateCertificateProviderInstance

func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyServerCertificateCertificateProviderInstance
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyServerCertificateCertificateProviderInstance
	} else {

		r.PluginInstance = res.PluginInstance

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyServerCertificateCertificateProviderInstance is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyServerCertificateCertificateProviderInstance *ServerTlsPolicyServerCertificateCertificateProviderInstance = &ServerTlsPolicyServerCertificateCertificateProviderInstance{empty: true}

func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyServerCertificateCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicy struct {
	empty              bool                                          `json:"-"`
	ClientValidationCa []ServerTlsPolicyMtlsPolicyClientValidationCa `json:"clientValidationCa"`
}

type jsonServerTlsPolicyMtlsPolicy ServerTlsPolicyMtlsPolicy

func (r *ServerTlsPolicyMtlsPolicy) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyMtlsPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyMtlsPolicy
	} else {

		r.ClientValidationCa = res.ClientValidationCa

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicy *ServerTlsPolicyMtlsPolicy = &ServerTlsPolicyMtlsPolicy{empty: true}

func (r *ServerTlsPolicyMtlsPolicy) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyMtlsPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicyClientValidationCa struct {
	empty                       bool                                                                    `json:"-"`
	CaCertPath                  *string                                                                 `json:"caCertPath"`
	GrpcEndpoint                *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance `json:"certificateProviderInstance"`
}

type jsonServerTlsPolicyMtlsPolicyClientValidationCa ServerTlsPolicyMtlsPolicyClientValidationCa

func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyMtlsPolicyClientValidationCa
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyMtlsPolicyClientValidationCa
	} else {

		r.CaCertPath = res.CaCertPath

		r.GrpcEndpoint = res.GrpcEndpoint

		r.CertificateProviderInstance = res.CertificateProviderInstance

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicyClientValidationCa is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicyClientValidationCa *ServerTlsPolicyMtlsPolicyClientValidationCa = &ServerTlsPolicyMtlsPolicyClientValidationCa{empty: true}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCa) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

type jsonServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint
	} else {

		r.TargetUri = res.TargetUri

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint = &ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{empty: true}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

type jsonServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) UnmarshalJSON(data []byte) error {
	var res jsonServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance
	} else {

		r.PluginInstance = res.PluginInstance

	}
	return nil
}

// This object is used to assert a desired state where this ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance = &ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{empty: true}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) Empty() bool {
	return r.empty
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ServerTlsPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_security",
		Type:    "ServerTlsPolicy",
		Version: "alpha",
	}
}

func (r *ServerTlsPolicy) ID() (string, error) {
	if err := extractServerTlsPolicyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":              dcl.ValueOrEmptyString(nr.Name),
		"description":       dcl.ValueOrEmptyString(nr.Description),
		"createTime":        dcl.ValueOrEmptyString(nr.CreateTime),
		"updateTime":        dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":            dcl.ValueOrEmptyString(nr.Labels),
		"allowOpen":         dcl.ValueOrEmptyString(nr.AllowOpen),
		"serverCertificate": dcl.ValueOrEmptyString(nr.ServerCertificate),
		"mtlsPolicy":        dcl.ValueOrEmptyString(nr.MtlsPolicy),
		"project":           dcl.ValueOrEmptyString(nr.Project),
		"location":          dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}", params), nil
}

const ServerTlsPolicyMaxPage = -1

type ServerTlsPolicyList struct {
	Items []*ServerTlsPolicy

	nextToken string

	pageSize int32

	resource *ServerTlsPolicy
}

func (l *ServerTlsPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServerTlsPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listServerTlsPolicy(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListServerTlsPolicy(ctx context.Context, project, location string) (*ServerTlsPolicyList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListServerTlsPolicyWithMaxResults(ctx, project, location, ServerTlsPolicyMaxPage)

}

func (c *Client) ListServerTlsPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ServerTlsPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &ServerTlsPolicy{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listServerTlsPolicy(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServerTlsPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetServerTlsPolicy(ctx context.Context, r *ServerTlsPolicy) (*ServerTlsPolicy, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractServerTlsPolicyFields(r)

	b, err := c.getServerTlsPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalServerTlsPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServerTlsPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractServerTlsPolicyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteServerTlsPolicy(ctx context.Context, r *ServerTlsPolicy) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ServerTlsPolicy resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting ServerTlsPolicy...")
	deleteOp := deleteServerTlsPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllServerTlsPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllServerTlsPolicy(ctx context.Context, project, location string, filter func(*ServerTlsPolicy) bool) error {
	listObj, err := c.ListServerTlsPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllServerTlsPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllServerTlsPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyServerTlsPolicy(ctx context.Context, rawDesired *ServerTlsPolicy, opts ...dcl.ApplyOption) (*ServerTlsPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *ServerTlsPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyServerTlsPolicyHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyServerTlsPolicyHelper(c *Client, ctx context.Context, rawDesired *ServerTlsPolicy, opts ...dcl.ApplyOption) (*ServerTlsPolicy, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyServerTlsPolicy...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractServerTlsPolicyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.serverTlsPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToServerTlsPolicyDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []serverTlsPolicyApiOperation
	if create {
		ops = append(ops, &createServerTlsPolicyOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyServerTlsPolicyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyServerTlsPolicyDiff(c *Client, ctx context.Context, desired *ServerTlsPolicy, rawDesired *ServerTlsPolicy, ops []serverTlsPolicyApiOperation, opts ...dcl.ApplyOption) (*ServerTlsPolicy, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetServerTlsPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServerTlsPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapServerTlsPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServerTlsPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServerTlsPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServerTlsPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractServerTlsPolicyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractServerTlsPolicyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffServerTlsPolicy(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}

func (r *ServerTlsPolicy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "POST", body, nil
}
