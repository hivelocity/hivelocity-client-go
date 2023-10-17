# \IgnitionApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIgnitionResourceId**](IgnitionApi.md#DeleteIgnitionResourceId) | **Delete** /ignition/{ignitionId} | Delete an Ignition configuration by ID
[**GetIgnitionResource**](IgnitionApi.md#GetIgnitionResource) | **Get** /ignition/ | Return all Ignition configurations
[**GetIgnitionResourceId**](IgnitionApi.md#GetIgnitionResourceId) | **Get** /ignition/{ignitionId} | Return an Ignition configuration by ID
[**PostIgnitionResource**](IgnitionApi.md#PostIgnitionResource) | **Post** /ignition/ | Creates a new Ignition configuration
[**PutIgnitionResourceId**](IgnitionApi.md#PutIgnitionResourceId) | **Put** /ignition/{ignitionId} | Update an Ignition configuration by ID


# **DeleteIgnitionResourceId**
> DeleteIgnitionResourceId(ctx, ignitionId, optional)
Delete an Ignition configuration by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ignitionId** | **int32**|  | 
 **optional** | ***IgnitionApiDeleteIgnitionResourceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgnitionApiDeleteIgnitionResourceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIgnitionResource**
> []IgnitionResponse GetIgnitionResource(ctx, optional)
Return all Ignition configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IgnitionApiGetIgnitionResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgnitionApiGetIgnitionResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]IgnitionResponse**](IgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIgnitionResourceId**
> IgnitionResponse GetIgnitionResourceId(ctx, ignitionId, optional)
Return an Ignition configuration by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ignitionId** | **int32**|  | 
 **optional** | ***IgnitionApiGetIgnitionResourceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgnitionApiGetIgnitionResourceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IgnitionResponse**](IgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIgnitionResource**
> IgnitionResponse PostIgnitionResource(ctx, payload, optional)
Creates a new Ignition configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**CreateIgnition**](CreateIgnition.md)|  | 
 **optional** | ***IgnitionApiPostIgnitionResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgnitionApiPostIgnitionResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IgnitionResponse**](IgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIgnitionResourceId**
> IgnitionResponse PutIgnitionResourceId(ctx, ignitionId, payload, optional)
Update an Ignition configuration by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ignitionId** | **int32**|  | 
  **payload** | [**UpdateIgnition**](UpdateIgnition.md)|  | 
 **optional** | ***IgnitionApiPutIgnitionResourceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IgnitionApiPutIgnitionResourceIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IgnitionResponse**](IgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

