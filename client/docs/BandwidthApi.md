# \BandwidthApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDeviceIdBandwidthImageResource**](BandwidthApi.md#PostDeviceIdBandwidthImageResource) | **Post** /bandwidth/device/{deviceId}/image | Get PNG by device
[**PostDeviceIdBandwidthResource**](BandwidthApi.md#PostDeviceIdBandwidthResource) | **Post** /bandwidth/device/{deviceId} | Get data by device
[**PostServiceIdBandwidthImageResource**](BandwidthApi.md#PostServiceIdBandwidthImageResource) | **Post** /bandwidth/service/{serviceId}/image | Get PNG by service
[**PostServiceIdBandwidthResource**](BandwidthApi.md#PostServiceIdBandwidthResource) | **Post** /bandwidth/service/{serviceId} | Get data by service


# **PostDeviceIdBandwidthImageResource**
> []BandwidthImage PostDeviceIdBandwidthImageResource(ctx, deviceId, period, interface_, optional)
Get PNG by device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View | 
  **period** | **string**| Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to day]
  **interface_** | **string**| Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to eth0]
 **optional** | ***BandwidthApiPostDeviceIdBandwidthImageResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BandwidthApiPostDeviceIdBandwidthImageResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.Int32**| Start Time of Custom Time Period. (Unix Epoch Time) | [default to 0]
 **end** | **optional.Int32**| End Time of Custom Time Period (Unix Epoch Time) | [default to 1672840541]
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]BandwidthImage**](BandwidthImage.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceIdBandwidthResource**
> []Bandwidth PostDeviceIdBandwidthResource(ctx, deviceId, period, interface_, step, optional)
Get data by device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View | 
  **period** | **string**| Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to day]
  **interface_** | **string**| Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to eth0]
  **step** | **int32**| Interval of data in seconds. Historical data is condensed, if provided value is smaller than existing steps for the date range the endpoint will return data with the smallest available step. | [default to 300]
 **optional** | ***BandwidthApiPostDeviceIdBandwidthResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BandwidthApiPostDeviceIdBandwidthResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **historical** | **optional.Bool**| If you are a reseller, this will include historical interface data for device regardless of the current device owner. | [default to false]
 **start** | **optional.Int32**| Start time of custom time period. (Unix Epoch Time) | [default to 0]
 **end** | **optional.Int32**| End time of custom time period (Unix Epoch Time) | [default to 1672840541]
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Bandwidth**](Bandwidth.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceIdBandwidthImageResource**
> []BandwidthImage PostServiceIdBandwidthImageResource(ctx, serviceId, period, interface_, optional)
Get PNG by service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **int32**| ID of Service to View | 
  **period** | **string**| Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to day]
  **interface_** | **string**| Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to eth0]
 **optional** | ***BandwidthApiPostServiceIdBandwidthImageResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BandwidthApiPostServiceIdBandwidthImageResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.Int32**| Start Time of Custom Time Period. (Unix Epoch Time) | [default to 0]
 **end** | **optional.Int32**| End Time of Custom Time Period (Unix Epoch Time) | [default to 1672840541]
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]BandwidthImage**](BandwidthImage.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostServiceIdBandwidthResource**
> []Bandwidth PostServiceIdBandwidthResource(ctx, serviceId, period, interface_, step, optional)
Get data by service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **int32**| ID of Service to View | 
  **period** | **string**| Return data in the given period. Day, week, month will return the previous day, week, month from now. | [default to day]
  **interface_** | **string**| Network interface to get bandwidth usage from. eth0 and eth1 are your first and second nic respectively. If unsure use the public, private, and all values. Overages are billed on public traffic. | [default to eth0]
  **step** | **int32**| Interval of data in seconds. Historical data is condensed, if provided value is smaller than existing steps for the date range the endpoint will return data with the smallest available step. | [default to 300]
 **optional** | ***BandwidthApiPostServiceIdBandwidthResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BandwidthApiPostServiceIdBandwidthResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| Start Time of Custom Time Period. (Unix Epoch Time) | [default to 0]
 **end** | **optional.Int32**| End Time of Custom Time Period (Unix Epoch Time) | [default to 1672840541]
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Bandwidth**](Bandwidth.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

