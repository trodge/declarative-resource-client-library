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
package vertexai

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type MetadataStore struct {
	Name           *string                      `json:"name"`
	CreateTime     *string                      `json:"createTime"`
	UpdateTime     *string                      `json:"updateTime"`
	EncryptionSpec *MetadataStoreEncryptionSpec `json:"encryptionSpec"`
	Description    *string                      `json:"description"`
	State          *MetadataStoreState          `json:"state"`
	Project        *string                      `json:"project"`
	Location       *string                      `json:"location"`
}

func (r *MetadataStore) String() string {
	return dcl.SprintResource(r)
}

type MetadataStoreEncryptionSpec struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonMetadataStoreEncryptionSpec MetadataStoreEncryptionSpec

func (r *MetadataStoreEncryptionSpec) UnmarshalJSON(data []byte) error {
	var res jsonMetadataStoreEncryptionSpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMetadataStoreEncryptionSpec
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this MetadataStoreEncryptionSpec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMetadataStoreEncryptionSpec *MetadataStoreEncryptionSpec = &MetadataStoreEncryptionSpec{empty: true}

func (r *MetadataStoreEncryptionSpec) Empty() bool {
	return r.empty
}

func (r *MetadataStoreEncryptionSpec) String() string {
	return dcl.SprintResource(r)
}

func (r *MetadataStoreEncryptionSpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MetadataStoreState struct {
	empty                bool   `json:"-"`
	DiskUtilizationBytes *int64 `json:"diskUtilizationBytes"`
}

type jsonMetadataStoreState MetadataStoreState

func (r *MetadataStoreState) UnmarshalJSON(data []byte) error {
	var res jsonMetadataStoreState
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMetadataStoreState
	} else {

		r.DiskUtilizationBytes = res.DiskUtilizationBytes

	}
	return nil
}

// This object is used to assert a desired state where this MetadataStoreState is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMetadataStoreState *MetadataStoreState = &MetadataStoreState{empty: true}

func (r *MetadataStoreState) Empty() bool {
	return r.empty
}

func (r *MetadataStoreState) String() string {
	return dcl.SprintResource(r)
}

func (r *MetadataStoreState) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *MetadataStore) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vertex_ai",
		Type:    "MetadataStore",
		Version: "vertexai",
	}
}

func (r *MetadataStore) ID() (string, error) {
	if err := extractMetadataStoreFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":            dcl.ValueOrEmptyString(nr.Name),
		"create_time":     dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":     dcl.ValueOrEmptyString(nr.UpdateTime),
		"encryption_spec": dcl.ValueOrEmptyString(nr.EncryptionSpec),
		"description":     dcl.ValueOrEmptyString(nr.Description),
		"state":           dcl.ValueOrEmptyString(nr.State),
		"project":         dcl.ValueOrEmptyString(nr.Project),
		"location":        dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/metadataStores/{{name}}", params), nil
}

const MetadataStoreMaxPage = -1

type MetadataStoreList struct {
	Items []*MetadataStore

	nextToken string

	pageSize int32

	resource *MetadataStore
}

func (l *MetadataStoreList) HasNext() bool {
	return l.nextToken != ""
}

func (l *MetadataStoreList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listMetadataStore(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListMetadataStore(ctx context.Context, project, location string) (*MetadataStoreList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListMetadataStoreWithMaxResults(ctx, project, location, MetadataStoreMaxPage)

}

func (c *Client) ListMetadataStoreWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*MetadataStoreList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &MetadataStore{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listMetadataStore(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &MetadataStoreList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetMetadataStore(ctx context.Context, r *MetadataStore) (*MetadataStore, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractMetadataStoreFields(r)

	b, err := c.getMetadataStoreRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalMetadataStore(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeMetadataStoreNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractMetadataStoreFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteMetadataStore(ctx context.Context, r *MetadataStore) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("MetadataStore resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting MetadataStore...")
	deleteOp := deleteMetadataStoreOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllMetadataStore deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllMetadataStore(ctx context.Context, project, location string, filter func(*MetadataStore) bool) error {
	listObj, err := c.ListMetadataStore(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllMetadataStore(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllMetadataStore(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyMetadataStore(ctx context.Context, rawDesired *MetadataStore, opts ...dcl.ApplyOption) (*MetadataStore, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *MetadataStore
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyMetadataStoreHelper(c, ctx, rawDesired, opts...)
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

func applyMetadataStoreHelper(c *Client, ctx context.Context, rawDesired *MetadataStore, opts ...dcl.ApplyOption) (*MetadataStore, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyMetadataStore...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractMetadataStoreFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.metadataStoreDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToMetadataStoreDiffs(c.Config, fieldDiffs, opts)
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
	var ops []metadataStoreApiOperation
	if create {
		ops = append(ops, &createMetadataStoreOperation{})
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
	return applyMetadataStoreDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyMetadataStoreDiff(c *Client, ctx context.Context, desired *MetadataStore, rawDesired *MetadataStore, ops []metadataStoreApiOperation, opts ...dcl.ApplyOption) (*MetadataStore, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetMetadataStore(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createMetadataStoreOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapMetadataStore(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeMetadataStoreNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeMetadataStoreNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeMetadataStoreDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractMetadataStoreFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractMetadataStoreFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffMetadataStore(c, newDesired, newState)
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
