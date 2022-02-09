# \BareMetalDevicesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBareMetalDeviceIdResource**](BareMetalDevicesApi.md#DeleteBareMetalDeviceIdResource) | **Delete** /bare-metal-devices/{deviceId} | Cancel the specified bare metal device
[**GetBareMetalDeviceIdResource**](BareMetalDevicesApi.md#GetBareMetalDeviceIdResource) | **Get** /bare-metal-devices/{deviceId} | Return bare metal device
[**GetBareMetalDeviceResource**](BareMetalDevicesApi.md#GetBareMetalDeviceResource) | **Get** /bare-metal-devices/ | Return a list with all servers as bare metal
[**PostBareMetalDeviceBatchResource**](BareMetalDevicesApi.md#PostBareMetalDeviceBatchResource) | **Post** /bare-metal-devices/batch | Deploy a list of new bare metal servers
[**PostBareMetalDeviceResource**](BareMetalDevicesApi.md#PostBareMetalDeviceResource) | **Post** /bare-metal-devices/ | Deploy a new bare metal server
[**PutBareMetalDeviceIdResource**](BareMetalDevicesApi.md#PutBareMetalDeviceIdResource) | **Put** /bare-metal-devices/{deviceId} | Update a bare metal device



## DeleteBareMetalDeviceIdResource

> DeleteBareMetalDeviceIdResource(ctx, deviceId, optional)

Cancel the specified bare metal device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32**|  | 
 **optional** | ***DeleteBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteBareMetalDeviceIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | **optional.String**| Id of the deployment to interact with | 

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


## GetBareMetalDeviceIdResource

> BareMetalDevice GetBareMetalDeviceIdResource(ctx, deviceId, optional)

Return bare metal device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32**|  | 
 **optional** | ***GetBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBareMetalDeviceIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | **optional.String**| Id of the deployment to interact with | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**BareMetalDevice**](BareMetalDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBareMetalDeviceResource

> []BareMetalDevice GetBareMetalDeviceResource(ctx, optional)

Return a list with all servers as bare metal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBareMetalDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBareMetalDeviceResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]BareMetalDevice**](BareMetalDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBareMetalDeviceBatchResource

> BareMetalDeviceBatch PostBareMetalDeviceBatchResource(ctx, payload, optional)

Deploy a list of new bare metal servers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**BareMetalDeviceBatchCreate**](BareMetalDeviceBatchCreate.md)|  | 
 **optional** | ***PostBareMetalDeviceBatchResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostBareMetalDeviceBatchResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**BareMetalDeviceBatch**](BareMetalDeviceBatch.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBareMetalDeviceResource

> BareMetalDevice PostBareMetalDeviceResource(ctx, payload, optional)

Deploy a new bare metal server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**BareMetalDeviceCreate**](BareMetalDeviceCreate.md)|  | 
 **optional** | ***PostBareMetalDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostBareMetalDeviceResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**BareMetalDevice**](BareMetalDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBareMetalDeviceIdResource

> BareMetalDevice PutBareMetalDeviceIdResource(ctx, deviceId, payload, optional)

Update a bare metal device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **int32**|  | 
**payload** | [**BareMetalDeviceUpdate**](BareMetalDeviceUpdate.md)|  | 
 **optional** | ***PutBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutBareMetalDeviceIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentId** | **optional.String**| Id of the deployment to interact with | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**BareMetalDevice**](BareMetalDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

