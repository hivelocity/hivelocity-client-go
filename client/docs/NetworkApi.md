# \NetworkApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNullRoutesIdResource**](NetworkApi.md#DeleteNullRoutesIdResource) | **Delete** /network/null-route/{ipAddress} | Remove Null Route for an IP Address
[**GetDeviceNetworkPortResource**](NetworkApi.md#GetDeviceNetworkPortResource) | **Get** /network/device/ports | Get all ports by device
[**GetNetworkTaskClientResource**](NetworkApi.md#GetNetworkTaskClientResource) | **Get** /network/status/ | Get network task status
[**GetNetworkTaskIdResource**](NetworkApi.md#GetNetworkTaskIdResource) | **Get** /network/status/{id} | Get network task status by Device ID or Task ID
[**GetNullRoutesIdResource**](NetworkApi.md#GetNullRoutesIdResource) | **Get** /network/null-route/{ipAddress} | Return record if IP Address is Null Routed
[**GetNullRoutesResource**](NetworkApi.md#GetNullRoutesResource) | **Get** /network/null-route | Return Null Routed IPs
[**PostNullRoutesResource**](NetworkApi.md#PostNullRoutesResource) | **Post** /network/null-route | Add a Null Route for an IP Address


# **DeleteNullRoutesIdResource**
> NetworkTaskDump DeleteNullRoutesIdResource(ctx, ipAddress, optional)
Remove Null Route for an IP Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**|  | 
 **optional** | ***NetworkApiDeleteNullRoutesIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiDeleteNullRoutesIdResourceOpts struct

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

# **GetDeviceNetworkPortResource**
> []DevicePort GetDeviceNetworkPortResource(ctx, optional)
Get all ports by device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkApiGetDeviceNetworkPortResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiGetDeviceNetworkPortResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]DevicePort**](DevicePort.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkTaskClientResource**
> []NetworkTaskDump GetNetworkTaskClientResource(ctx, optional)
Get network task status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkApiGetNetworkTaskClientResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiGetNetworkTaskClientResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.Int32**| Filter to retrieve tasks for a specific client. | 
 **createdInLast** | **optional.Int32**| Filter to retrieve tasks created in the last timedelta hours. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkTaskIdResource**
> NetworkTaskDump GetNetworkTaskIdResource(ctx, id, optional)
Get network task status by Device ID or Task ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NetworkApiGetNetworkTaskIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiGetNetworkTaskIdResourceOpts struct

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

# **GetNullRoutesIdResource**
> NullRoute GetNullRoutesIdResource(ctx, ipAddress, optional)
Return record if IP Address is Null Routed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**|  | 
 **optional** | ***NetworkApiGetNullRoutesIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiGetNullRoutesIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NullRoute**](NullRoute.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNullRoutesResource**
> []NullRoute GetNullRoutesResource(ctx, optional)
Return Null Routed IPs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkApiGetNullRoutesResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiGetNullRoutesResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.Int32**| Filter to retrieve null routes for a specific client. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]NullRoute**](NullRoute.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostNullRoutesResource**
> NetworkTaskDump PostNullRoutesResource(ctx, payload, optional)
Add a Null Route for an IP Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**NullRouteRequest**](NullRouteRequest.md)|  | 
 **optional** | ***NetworkApiPostNullRoutesResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkApiPostNullRoutesResourceOpts struct

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

