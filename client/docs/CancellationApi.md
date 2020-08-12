# \CancellationApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCancellationDeviceResource**](CancellationApi.md#GetCancellationDeviceResource) | **Get** /cancellation/device/{deviceId} | Return the Cancellation found for a Device
[**GetCancellationIdResource**](CancellationApi.md#GetCancellationIdResource) | **Get** /cancellation/cancellation/{cancellationId} | Return any cancellation by ID
[**GetCancellationResource**](CancellationApi.md#GetCancellationResource) | **Get** /cancellation/cancellation | Returns the services cancellations of a client
[**GetCancellationServiceResource**](CancellationApi.md#GetCancellationServiceResource) | **Get** /cancellation/service/{serviceId} | Return the Cancellation found for a Service
[**PostCancellationResource**](CancellationApi.md#PostCancellationResource) | **Post** /cancellation/cancellation | Creates Cancellation for a device/service


# **GetCancellationDeviceResource**
> GetCancellationDeviceResource(ctx, deviceId)
Return the Cancellation found for a Device

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

# **GetCancellationIdResource**
> GetCancellationIdResource(ctx, cancellationId)
Return any cancellation by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cancellationId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCancellationResource**
> GetCancellationResource(ctx, )
Returns the services cancellations of a client

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

# **GetCancellationServiceResource**
> GetCancellationServiceResource(ctx, serviceId)
Return the Cancellation found for a Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCancellationResource**
> PostCancellationResource(ctx, payload)
Creates Cancellation for a device/service

reason options are:     'My current service is being upgraded',     'Hardware/software/network or power issues',     'I think I have found a better deal',     'I have experienced tech support problems',     'I am consolidating my Hivelocity accounts',     'I am a reseller and my customer cancelled',     'I am moving to a different technology solution: Public Cloud',     'I am moving to a different technology solution: Managed Hosting Company',     'A product difference (Example: AWS)'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**CancellationCreate**](CancellationCreate.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

