# \IPAssignmentApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpAssignmentIdResource**](IPAssignmentApi.md#DeleteIpAssignmentIdResource) | **Delete** /ip/{ipAssignmentId} | Remove an existing IP Assignment
[**GetElasticPoolResource**](IPAssignmentApi.md#GetElasticPoolResource) | **Get** /ip/elastic | Return existing Elastic IP Assignment
[**GetIpAssignmentIdResource**](IPAssignmentApi.md#GetIpAssignmentIdResource) | **Get** /ip/{ipAssignmentId} | Return details about a Single IP Assignment
[**GetIpAssignmentResource**](IPAssignmentApi.md#GetIpAssignmentResource) | **Get** /ip/ | Returns all IP assignments a client has
[**PostElasticPoolResource**](IPAssignmentApi.md#PostElasticPoolResource) | **Post** /ip/elastic | Create a new Elastic IP Asignment
[**PostIpAssignmentResource**](IPAssignmentApi.md#PostIpAssignmentResource) | **Post** /ip/ | Request a new Secondary IP Assignment
[**PutElasticPoolIdResource**](IPAssignmentApi.md#PutElasticPoolIdResource) | **Put** /ip/elastic/{ipAssignmentId} | Update an existing Elastic Assignment, including adding and/or removing devices from it



## DeleteIpAssignmentIdResource

> IpAssignment DeleteIpAssignmentIdResource(ctx, ipAssignmentId, optional)

Remove an existing IP Assignment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAssignmentId** | **int32**|  | 
 **optional** | ***DeleteIpAssignmentIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteIpAssignmentIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElasticPoolResource

> []IpAssignment GetElasticPoolResource(ctx, optional)

Return existing Elastic IP Assignment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetElasticPoolResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetElasticPoolResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpAssignmentIdResource

> IpAssignment GetIpAssignmentIdResource(ctx, ipAssignmentId, optional)

Return details about a Single IP Assignment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAssignmentId** | **int32**|  | 
 **optional** | ***GetIpAssignmentIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetIpAssignmentIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpAssignmentResource

> []IpAssignment GetIpAssignmentResource(ctx, optional)

Returns all IP assignments a client has

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetIpAssignmentResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetIpAssignmentResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostElasticPoolResource

> IpAssignment PostElasticPoolResource(ctx, payload, optional)

Create a new Elastic IP Asignment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**ElasticPoolCreate**](ElasticPoolCreate.md)|  | 
 **optional** | ***PostElasticPoolResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostElasticPoolResourceOpts struct


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIpAssignmentResource

> PostIpAssignmentResource(ctx, )

Request a new Secondary IP Assignment

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutElasticPoolIdResource

> IpAssignment PutElasticPoolIdResource(ctx, ipAssignmentId, payload, optional)

Update an existing Elastic Assignment, including adding and/or removing devices from it

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAssignmentId** | **int32**| ID of the Elastic Assignment to interact with. | 
**payload** | [**ElasticPoolModify**](ElasticPoolModify.md)|  | 
 **optional** | ***PutElasticPoolIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutElasticPoolIdResourceOpts struct


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

