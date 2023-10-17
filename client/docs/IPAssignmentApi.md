# \IPAssignmentApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpAssignmentIdResource**](IPAssignmentApi.md#DeleteIpAssignmentIdResource) | **Delete** /ip/{ipAssignmentId} | Delete assignment
[**GetIpAssignmentIdResource**](IPAssignmentApi.md#GetIpAssignmentIdResource) | **Get** /ip/{ipAssignmentId} | Get assignment
[**GetIpAssignmentResource**](IPAssignmentApi.md#GetIpAssignmentResource) | **Get** /ip/ | Get all assignments
[**GetIpAssignmentSubnetResource**](IPAssignmentApi.md#GetIpAssignmentSubnetResource) | **Get** /ip/address/{ipAddress} | Get assignment by IP
[**PostIpAssignmentCombineResource**](IPAssignmentApi.md#PostIpAssignmentCombineResource) | **Post** /ip/combine | Combine assignments
[**PostIpAssignmentIdClearResource**](IPAssignmentApi.md#PostIpAssignmentIdClearResource) | **Post** /ip/{ipAssignmentId}/clear | Clear all configuration on an IP Assignment
[**PostIpAssignmentResource**](IPAssignmentApi.md#PostIpAssignmentResource) | **Post** /ip/ | Request new assignment
[**PostIpAssignmentSplitResource**](IPAssignmentApi.md#PostIpAssignmentSplitResource) | **Post** /ip/{ipAssignmentId}/split | Split assignment
[**PutIpAssignmentIdResource**](IPAssignmentApi.md#PutIpAssignmentIdResource) | **Put** /ip/{ipAssignmentId} | Route assignment to IP


# **DeleteIpAssignmentIdResource**
> DeleteIpAssignmentIdResource(ctx, ipAssignmentId)
Delete assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAssignmentId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpAssignmentIdResource**
> IpAssignment GetIpAssignmentIdResource(ctx, ipAssignmentId, optional)
Get assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAssignmentId** | **int32**|  | 
 **optional** | ***IPAssignmentApiGetIpAssignmentIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiGetIpAssignmentIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpAssignmentResource**
> []IpAssignment GetIpAssignmentResource(ctx, optional)
Get all assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAssignmentApiGetIpAssignmentResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiGetIpAssignmentResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayPrivateNetworks** | **optional.Bool**| Return Private (RFC 1918) Addresses in Results | [default to false]
 **displayIPv6** | **optional.Bool**| Return IPv6 Addresses in Results | [default to true]
 **clientId** | **optional.Int32**| Return IPs from this Client ID | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpAssignmentSubnetResource**
> IpAssignment GetIpAssignmentSubnetResource(ctx, ipAddress, optional)
Get assignment by IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**|  | 
 **optional** | ***IPAssignmentApiGetIpAssignmentSubnetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiGetIpAssignmentSubnetResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpAssignmentCombineResource**
> IpAssignment PostIpAssignmentCombineResource(ctx, payload, optional)
Combine assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**IpAssignmentCombine**](IpAssignmentCombine.md)|  | 
 **optional** | ***IPAssignmentApiPostIpAssignmentCombineResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiPostIpAssignmentCombineResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpAssignmentIdClearResource**
> NetworkTaskDump PostIpAssignmentIdClearResource(ctx, ipAssignmentId, optional)
Clear all configuration on an IP Assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAssignmentId** | **int32**|  | 
 **optional** | ***IPAssignmentApiPostIpAssignmentIdClearResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiPostIpAssignmentIdClearResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpAssignmentResource**
> PostIpAssignmentResource(ctx, payload)
Request new assignment

This endpoint is used to request a new IP Assignment from MyV and the API itself.  It will create a request on Zendesk for Networking support team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**IpAssignmentRequest**](IpAssignmentRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIpAssignmentSplitResource**
> []IpAssignment PostIpAssignmentSplitResource(ctx, ipAssignmentId, optional)
Split assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAssignmentId** | **int32**|  | 
 **optional** | ***IPAssignmentApiPostIpAssignmentSplitResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiPostIpAssignmentSplitResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIpAssignmentIdResource**
> NetworkTaskDump PutIpAssignmentIdResource(ctx, ipAssignmentId, payload, optional)
Route assignment to IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAssignmentId** | **int32**|  | 
  **payload** | [**IpAssignmentPut**](IpAssignmentPut.md)|  | 
 **optional** | ***IPAssignmentApiPutIpAssignmentIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAssignmentApiPutIpAssignmentIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

