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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Service struct {
	Name                  *string                       `json:"name"`
	DisplayName           *string                       `json:"displayName"`
	Custom                *ServiceCustom                `json:"custom"`
	AppEngine             *ServiceAppEngine             `json:"appEngine"`
	CloudEndpoints        *ServiceCloudEndpoints        `json:"cloudEndpoints"`
	ClusterIstio          *ServiceClusterIstio          `json:"clusterIstio"`
	MeshIstio             *ServiceMeshIstio             `json:"meshIstio"`
	IstioCanonicalService *ServiceIstioCanonicalService `json:"istioCanonicalService"`
	CloudRun              *ServiceCloudRun              `json:"cloudRun"`
	GkeNamespace          *ServiceGkeNamespace          `json:"gkeNamespace"`
	GkeWorkload           *ServiceGkeWorkload           `json:"gkeWorkload"`
	GkeService            *ServiceGkeService            `json:"gkeService"`
	Telemetry             *ServiceTelemetry             `json:"telemetry"`
	UserLabels            map[string]string             `json:"userLabels"`
	Deleted               *bool                         `json:"deleted"`
	Project               *string                       `json:"project"`
}

func (r *Service) String() string {
	return dcl.SprintResource(r)
}

type ServiceCustom struct {
	empty bool `json:"-"`
}

type jsonServiceCustom ServiceCustom

func (r *ServiceCustom) UnmarshalJSON(data []byte) error {
	var res jsonServiceCustom
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceCustom
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this ServiceCustom is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceCustom *ServiceCustom = &ServiceCustom{empty: true}

func (r *ServiceCustom) Empty() bool {
	return r.empty
}

func (r *ServiceCustom) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceCustom) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceAppEngine struct {
	empty    bool    `json:"-"`
	ModuleId *string `json:"moduleId"`
}

type jsonServiceAppEngine ServiceAppEngine

func (r *ServiceAppEngine) UnmarshalJSON(data []byte) error {
	var res jsonServiceAppEngine
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceAppEngine
	} else {

		r.ModuleId = res.ModuleId

	}
	return nil
}

// This object is used to assert a desired state where this ServiceAppEngine is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceAppEngine *ServiceAppEngine = &ServiceAppEngine{empty: true}

func (r *ServiceAppEngine) Empty() bool {
	return r.empty
}

func (r *ServiceAppEngine) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceAppEngine) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceCloudEndpoints struct {
	empty   bool    `json:"-"`
	Service *string `json:"service"`
}

type jsonServiceCloudEndpoints ServiceCloudEndpoints

func (r *ServiceCloudEndpoints) UnmarshalJSON(data []byte) error {
	var res jsonServiceCloudEndpoints
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceCloudEndpoints
	} else {

		r.Service = res.Service

	}
	return nil
}

// This object is used to assert a desired state where this ServiceCloudEndpoints is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceCloudEndpoints *ServiceCloudEndpoints = &ServiceCloudEndpoints{empty: true}

func (r *ServiceCloudEndpoints) Empty() bool {
	return r.empty
}

func (r *ServiceCloudEndpoints) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceCloudEndpoints) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceClusterIstio struct {
	empty            bool    `json:"-"`
	Location         *string `json:"location"`
	ClusterName      *string `json:"clusterName"`
	ServiceNamespace *string `json:"serviceNamespace"`
	ServiceName      *string `json:"serviceName"`
}

type jsonServiceClusterIstio ServiceClusterIstio

func (r *ServiceClusterIstio) UnmarshalJSON(data []byte) error {
	var res jsonServiceClusterIstio
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceClusterIstio
	} else {

		r.Location = res.Location

		r.ClusterName = res.ClusterName

		r.ServiceNamespace = res.ServiceNamespace

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceClusterIstio is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceClusterIstio *ServiceClusterIstio = &ServiceClusterIstio{empty: true}

func (r *ServiceClusterIstio) Empty() bool {
	return r.empty
}

func (r *ServiceClusterIstio) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceClusterIstio) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceMeshIstio struct {
	empty            bool    `json:"-"`
	MeshUid          *string `json:"meshUid"`
	ServiceNamespace *string `json:"serviceNamespace"`
	ServiceName      *string `json:"serviceName"`
}

type jsonServiceMeshIstio ServiceMeshIstio

func (r *ServiceMeshIstio) UnmarshalJSON(data []byte) error {
	var res jsonServiceMeshIstio
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceMeshIstio
	} else {

		r.MeshUid = res.MeshUid

		r.ServiceNamespace = res.ServiceNamespace

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceMeshIstio is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceMeshIstio *ServiceMeshIstio = &ServiceMeshIstio{empty: true}

func (r *ServiceMeshIstio) Empty() bool {
	return r.empty
}

func (r *ServiceMeshIstio) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceMeshIstio) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceIstioCanonicalService struct {
	empty                     bool    `json:"-"`
	MeshUid                   *string `json:"meshUid"`
	CanonicalServiceNamespace *string `json:"canonicalServiceNamespace"`
	CanonicalService          *string `json:"canonicalService"`
}

type jsonServiceIstioCanonicalService ServiceIstioCanonicalService

func (r *ServiceIstioCanonicalService) UnmarshalJSON(data []byte) error {
	var res jsonServiceIstioCanonicalService
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceIstioCanonicalService
	} else {

		r.MeshUid = res.MeshUid

		r.CanonicalServiceNamespace = res.CanonicalServiceNamespace

		r.CanonicalService = res.CanonicalService

	}
	return nil
}

// This object is used to assert a desired state where this ServiceIstioCanonicalService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceIstioCanonicalService *ServiceIstioCanonicalService = &ServiceIstioCanonicalService{empty: true}

func (r *ServiceIstioCanonicalService) Empty() bool {
	return r.empty
}

func (r *ServiceIstioCanonicalService) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceIstioCanonicalService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceCloudRun struct {
	empty       bool    `json:"-"`
	ServiceName *string `json:"serviceName"`
	Location    *string `json:"location"`
}

type jsonServiceCloudRun ServiceCloudRun

func (r *ServiceCloudRun) UnmarshalJSON(data []byte) error {
	var res jsonServiceCloudRun
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceCloudRun
	} else {

		r.ServiceName = res.ServiceName

		r.Location = res.Location

	}
	return nil
}

// This object is used to assert a desired state where this ServiceCloudRun is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceCloudRun *ServiceCloudRun = &ServiceCloudRun{empty: true}

func (r *ServiceCloudRun) Empty() bool {
	return r.empty
}

func (r *ServiceCloudRun) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceCloudRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceGkeNamespace struct {
	empty         bool    `json:"-"`
	ProjectId     *string `json:"projectId"`
	Location      *string `json:"location"`
	ClusterName   *string `json:"clusterName"`
	NamespaceName *string `json:"namespaceName"`
}

type jsonServiceGkeNamespace ServiceGkeNamespace

func (r *ServiceGkeNamespace) UnmarshalJSON(data []byte) error {
	var res jsonServiceGkeNamespace
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceGkeNamespace
	} else {

		r.ProjectId = res.ProjectId

		r.Location = res.Location

		r.ClusterName = res.ClusterName

		r.NamespaceName = res.NamespaceName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceGkeNamespace is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceGkeNamespace *ServiceGkeNamespace = &ServiceGkeNamespace{empty: true}

func (r *ServiceGkeNamespace) Empty() bool {
	return r.empty
}

func (r *ServiceGkeNamespace) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceGkeNamespace) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceGkeWorkload struct {
	empty                  bool    `json:"-"`
	ProjectId              *string `json:"projectId"`
	Location               *string `json:"location"`
	ClusterName            *string `json:"clusterName"`
	NamespaceName          *string `json:"namespaceName"`
	TopLevelControllerType *string `json:"topLevelControllerType"`
	TopLevelControllerName *string `json:"topLevelControllerName"`
}

type jsonServiceGkeWorkload ServiceGkeWorkload

func (r *ServiceGkeWorkload) UnmarshalJSON(data []byte) error {
	var res jsonServiceGkeWorkload
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceGkeWorkload
	} else {

		r.ProjectId = res.ProjectId

		r.Location = res.Location

		r.ClusterName = res.ClusterName

		r.NamespaceName = res.NamespaceName

		r.TopLevelControllerType = res.TopLevelControllerType

		r.TopLevelControllerName = res.TopLevelControllerName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceGkeWorkload is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceGkeWorkload *ServiceGkeWorkload = &ServiceGkeWorkload{empty: true}

func (r *ServiceGkeWorkload) Empty() bool {
	return r.empty
}

func (r *ServiceGkeWorkload) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceGkeWorkload) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceGkeService struct {
	empty         bool    `json:"-"`
	ProjectId     *string `json:"projectId"`
	Location      *string `json:"location"`
	ClusterName   *string `json:"clusterName"`
	NamespaceName *string `json:"namespaceName"`
	ServiceName   *string `json:"serviceName"`
}

type jsonServiceGkeService ServiceGkeService

func (r *ServiceGkeService) UnmarshalJSON(data []byte) error {
	var res jsonServiceGkeService
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceGkeService
	} else {

		r.ProjectId = res.ProjectId

		r.Location = res.Location

		r.ClusterName = res.ClusterName

		r.NamespaceName = res.NamespaceName

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceGkeService is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceGkeService *ServiceGkeService = &ServiceGkeService{empty: true}

func (r *ServiceGkeService) Empty() bool {
	return r.empty
}

func (r *ServiceGkeService) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceGkeService) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTelemetry struct {
	empty        bool    `json:"-"`
	ResourceName *string `json:"resourceName"`
}

type jsonServiceTelemetry ServiceTelemetry

func (r *ServiceTelemetry) UnmarshalJSON(data []byte) error {
	var res jsonServiceTelemetry
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTelemetry
	} else {

		r.ResourceName = res.ResourceName

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTelemetry is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyServiceTelemetry *ServiceTelemetry = &ServiceTelemetry{empty: true}

func (r *ServiceTelemetry) Empty() bool {
	return r.empty
}

func (r *ServiceTelemetry) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTelemetry) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Service) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "monitoring",
		Type:    "Service",
		Version: "alpha",
	}
}

func (r *Service) ID() (string, error) {
	if err := extractServiceFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                  dcl.ValueOrEmptyString(nr.Name),
		"displayName":           dcl.ValueOrEmptyString(nr.DisplayName),
		"custom":                dcl.ValueOrEmptyString(nr.Custom),
		"appEngine":             dcl.ValueOrEmptyString(nr.AppEngine),
		"cloudEndpoints":        dcl.ValueOrEmptyString(nr.CloudEndpoints),
		"clusterIstio":          dcl.ValueOrEmptyString(nr.ClusterIstio),
		"meshIstio":             dcl.ValueOrEmptyString(nr.MeshIstio),
		"istioCanonicalService": dcl.ValueOrEmptyString(nr.IstioCanonicalService),
		"cloudRun":              dcl.ValueOrEmptyString(nr.CloudRun),
		"gkeNamespace":          dcl.ValueOrEmptyString(nr.GkeNamespace),
		"gkeWorkload":           dcl.ValueOrEmptyString(nr.GkeWorkload),
		"gkeService":            dcl.ValueOrEmptyString(nr.GkeService),
		"telemetry":             dcl.ValueOrEmptyString(nr.Telemetry),
		"userLabels":            dcl.ValueOrEmptyString(nr.UserLabels),
		"deleted":               dcl.ValueOrEmptyString(nr.Deleted),
		"project":               dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/services/{{name}}", params), nil
}

const ServiceMaxPage = -1

type ServiceList struct {
	Items []*Service

	nextToken string

	pageSize int32

	resource *Service
}

func (l *ServiceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServiceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listService(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListService(ctx context.Context, project string) (*ServiceList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListServiceWithMaxResults(ctx, project, ServiceMaxPage)

}

func (c *Client) ListServiceWithMaxResults(ctx context.Context, project string, pageSize int32) (*ServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Service{
		Project: &project,
	}
	items, token, err := c.listService(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServiceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetService(ctx context.Context, r *Service) (*Service, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractServiceFields(r)

	b, err := c.getServiceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalService(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServiceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractServiceFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteService(ctx context.Context, r *Service) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Service resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Service...")
	deleteOp := deleteServiceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllService deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllService(ctx context.Context, project string, filter func(*Service) bool) error {
	listObj, err := c.ListService(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllService(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllService(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyService(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (*Service, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Service
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyServiceHelper(c, ctx, rawDesired, opts...)
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

func applyServiceHelper(c *Client, ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (*Service, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyService...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractServiceFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.serviceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToServiceDiffs(c.Config, fieldDiffs, opts)
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
	var ops []serviceApiOperation
	if create {
		ops = append(ops, &createServiceOperation{})
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

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetService(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServiceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapService(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServiceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServiceNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServiceDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractServiceFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractServiceFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffService(c, newDesired, newState)
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
