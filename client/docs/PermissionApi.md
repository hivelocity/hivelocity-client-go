# \PermissionApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionAllResource**](PermissionApi.md#GetPermissionAllResource) | **Get** /permission/ | Endpoint to get All Permissions
[**GetPermissionContactResource**](PermissionApi.md#GetPermissionContactResource) | **Get** /permission/contact/{contactId} | Endpoint to get Contact Permissions
[**GetPermissionUserResource**](PermissionApi.md#GetPermissionUserResource) | **Get** /permission/user | Endpoint to get User Permissions
[**PostPermissionAssignContactResource**](PermissionApi.md#PostPermissionAssignContactResource) | **Post** /permission/contact | Endpoint to assign a new Permission to a Contact


# **GetPermissionAllResource**
> GetPermissionAllResource(ctx, )
Endpoint to get All Permissions

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionContactResource**
> GetPermissionContactResource(ctx, contactId)
Endpoint to get Contact Permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contactId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionUserResource**
> GetPermissionUserResource(ctx, )
Endpoint to get User Permissions

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPermissionAssignContactResource**
> PostPermissionAssignContactResource(ctx, payload)
Endpoint to assign a new Permission to a Contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**Permission**](Permission.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

