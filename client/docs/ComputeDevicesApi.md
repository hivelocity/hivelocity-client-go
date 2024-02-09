# \ComputeDevicesApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBareMetalDeviceIdResource**](ComputeDevicesApi.md#DeleteBareMetalDeviceIdResource) | **Delete** /compute/{deviceId} | Cancel/delete device
[**GetBareMetalDeviceIdResource**](ComputeDevicesApi.md#GetBareMetalDeviceIdResource) | **Get** /compute/{deviceId} | Get device
[**GetBareMetalDeviceResource**](ComputeDevicesApi.md#GetBareMetalDeviceResource) | **Get** /compute/ | Get all devices
[**PostBareMetalDeviceBatchResource**](ComputeDevicesApi.md#PostBareMetalDeviceBatchResource) | **Post** /compute/batch | Batch provision instant devices
[**PostBareMetalDeviceResource**](ComputeDevicesApi.md#PostBareMetalDeviceResource) | **Post** /compute/ | Provision an instant device
[**PutBareMetalDeviceIdResource**](ComputeDevicesApi.md#PutBareMetalDeviceIdResource) | **Put** /compute/{deviceId} | Update/reload instant device


# **DeleteBareMetalDeviceIdResource**
> DeleteBareMetalDeviceIdResource(ctx, deviceId)
Cancel/delete device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 

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
 **optional** | ***ComputeDevicesApiGetBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComputeDevicesApiGetBareMetalDeviceIdResourceOpts struct

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

# **GetBareMetalDeviceResource**
> []BareMetalDevice GetBareMetalDeviceResource(ctx, optional)
Get all devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComputeDevicesApiGetBareMetalDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComputeDevicesApiGetBareMetalDeviceResourceOpts struct

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
 **optional** | ***ComputeDevicesApiPostBareMetalDeviceBatchResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComputeDevicesApiPostBareMetalDeviceBatchResourceOpts struct

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
 **optional** | ***ComputeDevicesApiPostBareMetalDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComputeDevicesApiPostBareMetalDeviceResourceOpts struct

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
 **optional** | ***ComputeDevicesApiPutBareMetalDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComputeDevicesApiPutBareMetalDeviceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

