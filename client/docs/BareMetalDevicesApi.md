# \BareMetalDevicesApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBareMetalDeviceIdResource**](BareMetalDevicesApi.md#DeleteBareMetalDeviceIdResource) | **Delete** /bare-metal-devices/{deviceId} | Cancel/delete device
[**GetBareMetalDeviceIdResource**](BareMetalDevicesApi.md#GetBareMetalDeviceIdResource) | **Get** /bare-metal-devices/{deviceId} | Get device
[**GetBareMetalDeviceResource**](BareMetalDevicesApi.md#GetBareMetalDeviceResource) | **Get** /bare-metal-devices/ | Get all devices
[**PostBareMetalDeviceBatchResource**](BareMetalDevicesApi.md#PostBareMetalDeviceBatchResource) | **Post** /bare-metal-devices/batch | Batch provision instant devices
[**PostBareMetalDeviceResource**](BareMetalDevicesApi.md#PostBareMetalDeviceResource) | **Post** /bare-metal-devices/ | Provision an instant device
[**PutBareMetalDeviceIdResource**](BareMetalDevicesApi.md#PutBareMetalDeviceIdResource) | **Put** /bare-metal-devices/{deviceId} | Update/reload instant device


# **DeleteBareMetalDeviceIdResource**
> DeleteBareMetalDeviceIdResource(ctx, deviceId, optional)
Cancel/delete device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***BareMetalDevicesApiDeleteBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BareMetalDevicesApiDeleteBareMetalDeviceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | **optional.String**| Id of the deployment to interact with | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBareMetalDeviceIdResource**
> BareMetalDevice GetBareMetalDeviceIdResource(ctx, deviceId, optional)
Get device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***BareMetalDevicesApiGetBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BareMetalDevicesApiGetBareMetalDeviceIdResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBareMetalDeviceResource**
> []BareMetalDevice GetBareMetalDeviceResource(ctx, optional)
Get all devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BareMetalDevicesApiGetBareMetalDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BareMetalDevicesApiGetBareMetalDeviceResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]BareMetalDevice**](BareMetalDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBareMetalDeviceBatchResource**
> BareMetalDeviceBatch PostBareMetalDeviceBatchResource(ctx, payload, optional)
Batch provision instant devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**BareMetalDeviceBatchCreate**](BareMetalDeviceBatchCreate.md)|  | 
 **optional** | ***BareMetalDevicesApiPostBareMetalDeviceBatchResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BareMetalDevicesApiPostBareMetalDeviceBatchResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBareMetalDeviceResource**
> BareMetalDevice PostBareMetalDeviceResource(ctx, payload, optional)
Provision an instant device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**BareMetalDeviceCreate**](BareMetalDeviceCreate.md)|  | 
 **optional** | ***BareMetalDevicesApiPostBareMetalDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BareMetalDevicesApiPostBareMetalDeviceResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBareMetalDeviceIdResource**
> BareMetalDevice PutBareMetalDeviceIdResource(ctx, deviceId, payload, optional)
Update/reload instant device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **payload** | [**BareMetalDeviceUpdate**](BareMetalDeviceUpdate.md)|  | 
 **optional** | ***BareMetalDevicesApiPutBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BareMetalDevicesApiPutBareMetalDeviceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentId** | **optional.String**| Id of the deployment to interact with | 
 **skipPowerCheck** | **optional.Bool**| If true, bypass the powered off check. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**BareMetalDevice**](BareMetalDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

