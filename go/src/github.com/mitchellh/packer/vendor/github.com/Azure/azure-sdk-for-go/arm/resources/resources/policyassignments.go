package resources

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// PolicyAssignmentsClient is the client for the PolicyAssignments methods of
// the Resources service.
type PolicyAssignmentsClient struct {
	ManagementClient
}

// NewPolicyAssignmentsClient creates an instance of the
// PolicyAssignmentsClient client.
func NewPolicyAssignmentsClient(subscriptionID string) PolicyAssignmentsClient {
	return NewPolicyAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPolicyAssignmentsClientWithBaseURI creates an instance of the
// PolicyAssignmentsClient client.
func NewPolicyAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) PolicyAssignmentsClient {
	return PolicyAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create policy assignment.
//
// scope is scope. policyAssignmentName is policy assignment name. parameters
// is policy assignment.
func (client PolicyAssignmentsClient) Create(scope string, policyAssignmentName string, parameters PolicyAssignment) (result PolicyAssignment, ae error) {
	req, err := client.CreatePreparer(scope, policyAssignmentName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PolicyAssignmentsClient) CreatePreparer(scope string, policyAssignmentName string, parameters PolicyAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": url.QueryEscape(policyAssignmentName),
		"scope":                scope,
		"subscriptionId":       url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) CreateResponder(resp *http.Response) (result PolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateByID create policy assignment by Id.
//
// policyAssignmentID is policy assignment Id parameters is policy assignment.
func (client PolicyAssignmentsClient) CreateByID(policyAssignmentID string, parameters PolicyAssignment) (result PolicyAssignment, ae error) {
	req, err := client.CreateByIDPreparer(policyAssignmentID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "CreateByID", nil, "Failure preparing request")
	}

	resp, err := client.CreateByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "CreateByID", resp, "Failure sending request")
	}

	result, err = client.CreateByIDResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "CreateByID", resp, "Failure responding to request")
	}

	return
}

// CreateByIDPreparer prepares the CreateByID request.
func (client PolicyAssignmentsClient) CreateByIDPreparer(policyAssignmentID string, parameters PolicyAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{policyAssignmentId}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateByIDSender sends the CreateByID request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) CreateByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// CreateByIDResponder handles the response to the CreateByID request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) CreateByIDResponder(resp *http.Response) (result PolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete policy assignment.
//
// scope is scope. policyAssignmentName is policy assignment name.
func (client PolicyAssignmentsClient) Delete(scope string, policyAssignmentName string) (result PolicyAssignment, ae error) {
	req, err := client.DeletePreparer(scope, policyAssignmentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PolicyAssignmentsClient) DeletePreparer(scope string, policyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": url.QueryEscape(policyAssignmentName),
		"scope":                scope,
		"subscriptionId":       url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) DeleteResponder(resp *http.Response) (result PolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteByID delete policy assignment.
//
// policyAssignmentID is policy assignment Id
func (client PolicyAssignmentsClient) DeleteByID(policyAssignmentID string) (result PolicyAssignment, ae error) {
	req, err := client.DeleteByIDPreparer(policyAssignmentID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "DeleteByID", nil, "Failure preparing request")
	}

	resp, err := client.DeleteByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "DeleteByID", resp, "Failure sending request")
	}

	result, err = client.DeleteByIDResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "DeleteByID", resp, "Failure responding to request")
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client PolicyAssignmentsClient) DeleteByIDPreparer(policyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{policyAssignmentId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) DeleteByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) DeleteByIDResponder(resp *http.Response) (result PolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get single policy assignment.
//
// scope is scope. policyAssignmentName is policy assignment name.
func (client PolicyAssignmentsClient) Get(scope string, policyAssignmentName string) (result PolicyAssignment, ae error) {
	req, err := client.GetPreparer(scope, policyAssignmentName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PolicyAssignmentsClient) GetPreparer(scope string, policyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": url.QueryEscape(policyAssignmentName),
		"scope":                scope,
		"subscriptionId":       url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) GetResponder(resp *http.Response) (result PolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID get single policy assignment.
//
// policyAssignmentID is policy assignment Id
func (client PolicyAssignmentsClient) GetByID(policyAssignmentID string) (result PolicyAssignment, ae error) {
	req, err := client.GetByIDPreparer(policyAssignmentID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "GetByID", nil, "Failure preparing request")
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "GetByID", resp, "Failure sending request")
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client PolicyAssignmentsClient) GetByIDPreparer(policyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{policyAssignmentId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) GetByIDResponder(resp *http.Response) (result PolicyAssignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets policy assignments of the subscription.
//
// filter is the filter to apply on the operation.
func (client PolicyAssignmentsClient) List(filter string) (result PolicyAssignmentListResult, ae error) {
	req, err := client.ListPreparer(filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PolicyAssignmentsClient) ListPreparer(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) ListResponder(resp *http.Response) (result PolicyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client PolicyAssignmentsClient) ListNextResults(lastResults PolicyAssignmentListResult) (result PolicyAssignmentListResult, ae error) {
	req, err := lastResults.PolicyAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListForResource gets policy assignments of the resource.
//
// resourceGroupName is the name of the resource group.
// resourceProviderNamespace is the name of the resource provider.
// parentResourcePath is the parent resource path. resourceType is the
// resource type. resourceName is the resource name. filter is the filter to
// apply on the operation.
func (client PolicyAssignmentsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result PolicyAssignmentListResult, ae error) {
	req, err := client.ListForResourcePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResource", nil, "Failure preparing request")
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResource", resp, "Failure sending request")
	}

	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResource", resp, "Failure responding to request")
	}

	return
}

// ListForResourcePreparer prepares the ListForResource request.
func (client PolicyAssignmentsClient) ListForResourcePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        url.QueryEscape(parentResourcePath),
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              url.QueryEscape(resourceType),
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}providers/Microsoft.Authorization/policyAssignments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListForResourceSender sends the ListForResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) ListForResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListForResourceResponder handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) ListForResourceResponder(resp *http.Response) (result PolicyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceNextResults retrieves the next set of results, if any.
func (client PolicyAssignmentsClient) ListForResourceNextResults(lastResults PolicyAssignmentListResult) (result PolicyAssignmentListResult, ae error) {
	req, err := lastResults.PolicyAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResource", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResource", resp, "Failure sending next results request request")
	}

	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResource", resp, "Failure responding to next results request request")
	}

	return
}

// ListForResourceGroup gets policy assignments of the resource group.
//
// resourceGroupName is resource group name. filter is the filter to apply on
// the operation.
func (client PolicyAssignmentsClient) ListForResourceGroup(resourceGroupName string, filter string) (result PolicyAssignmentListResult, ae error) {
	req, err := client.ListForResourceGroupPreparer(resourceGroupName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client PolicyAssignmentsClient) ListForResourceGroupPreparer(resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyAssignments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) ListForResourceGroupResponder(resp *http.Response) (result PolicyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForResourceGroupNextResults retrieves the next set of results, if any.
func (client PolicyAssignmentsClient) ListForResourceGroupNextResults(lastResults PolicyAssignmentListResult) (result PolicyAssignmentListResult, ae error) {
	req, err := lastResults.PolicyAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResourceGroup", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResourceGroup", resp, "Failure sending next results request request")
	}

	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForResourceGroup", resp, "Failure responding to next results request request")
	}

	return
}

// ListForScope gets policy assignments of the scope.
//
// scope is scope. filter is the filter to apply on the operation.
func (client PolicyAssignmentsClient) ListForScope(scope string, filter string) (result PolicyAssignmentListResult, ae error) {
	req, err := client.ListForScopePreparer(scope, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForScope", nil, "Failure preparing request")
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForScope", resp, "Failure sending request")
	}

	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForScope", resp, "Failure responding to request")
	}

	return
}

// ListForScopePreparer prepares the ListForScope request.
func (client PolicyAssignmentsClient) ListForScopePreparer(scope string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope":          scope,
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{scope}/providers/Microsoft.Authorization/policyAssignments"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListForScopeSender sends the ListForScope request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyAssignmentsClient) ListForScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListForScopeResponder handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (client PolicyAssignmentsClient) ListForScopeResponder(resp *http.Response) (result PolicyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForScopeNextResults retrieves the next set of results, if any.
func (client PolicyAssignmentsClient) ListForScopeNextResults(lastResults PolicyAssignmentListResult) (result PolicyAssignmentListResult, ae error) {
	req, err := lastResults.PolicyAssignmentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForScope", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForScope", resp, "Failure sending next results request request")
	}

	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/PolicyAssignmentsClient", "ListForScope", resp, "Failure responding to next results request request")
	}

	return
}