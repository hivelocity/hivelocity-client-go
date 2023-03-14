# \DeviceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeviceBondResource**](DeviceApi.md#DeleteDeviceBondResource) | **Delete** /device/{deviceId}/ports/bond | Unbond ports
[**GetAllDeviceTagOrderResource**](DeviceApi.md#GetAllDeviceTagOrderResource) | **Get** /device/tags-order/all | Get all tag orders
[**GetClientDeviceTagOrderResource**](DeviceApi.md#GetClientDeviceTagOrderResource) | **Get** /device/tags-order | Get tag order
[**GetClientDeviceTagResource**](DeviceApi.md#GetClientDeviceTagResource) | **Get** /device/tags | Get all tags
[**GetDeviceIdEventResource**](DeviceApi.md#GetDeviceIdEventResource) | **Get** /device/{deviceId}/events | Get all events
[**GetDeviceIdResource**](DeviceApi.md#GetDeviceIdResource) | **Get** /device/{deviceId} | Get device
[**GetDeviceIdServiceResource**](DeviceApi.md#GetDeviceIdServiceResource) | **Get** /device/{deviceId}/services | Get all device services
[**GetDeviceIpAssignmentsResource**](DeviceApi.md#GetDeviceIpAssignmentsResource) | **Get** /device/{deviceId}/ips | Get IPs
[**GetDeviceIpminatRuleResource**](DeviceApi.md#GetDeviceIpminatRuleResource) | **Get** /device/{deviceId}/ipmi/nat | Clear IPMI whitelist
[**GetDevicePortResource**](DeviceApi.md#GetDevicePortResource) | **Get** /device/{deviceId}/ports | Get ports
[**GetDeviceResource**](DeviceApi.md#GetDeviceResource) | **Get** /device/ | Get all devices
[**GetDeviceTagIdResource**](DeviceApi.md#GetDeviceTagIdResource) | **Get** /device/{deviceId}/tags | Get device tags
[**GetEffectiveIgnitionIdResource**](DeviceApi.md#GetEffectiveIgnitionIdResource) | **Get** /device/{deviceId}/ignition | Get Ignition injections
[**GetInitialCredsIdResource**](DeviceApi.md#GetInitialCredsIdResource) | **Get** /device/{deviceId}/initial-creds | Get initial creds
[**GetInitialPasswordIdResource**](DeviceApi.md#GetInitialPasswordIdResource) | **Get** /device/{deviceId}/initial-password | Get initial password
[**GetIpmiInfoIdResource**](DeviceApi.md#GetIpmiInfoIdResource) | **Get** /device/{deviceId}/ipmi | Get IPMI data
[**GetIpmiInfoLoginDataResource**](DeviceApi.md#GetIpmiInfoLoginDataResource) | **Get** /device/{deviceId}/ipmi/login-data | Get IPMI creds
[**GetIpmiThresholdsIdResource**](DeviceApi.md#GetIpmiThresholdsIdResource) | **Get** /device/{deviceId}/ipmi/thresholds | Get IPMI thresholds
[**GetIpmiValidLoginIdResource**](DeviceApi.md#GetIpmiValidLoginIdResource) | **Get** /device/{deviceId}/ipmi/valid-login | Get IPMI access status
[**GetPowerResource**](DeviceApi.md#GetPowerResource) | **Get** /device/{deviceId}/power | Get power status
[**PostDeviceBondResource**](DeviceApi.md#PostDeviceBondResource) | **Post** /device/{deviceId}/ports/bond | Bond ports
[**PostDeviceIpmiWhitelistResource**](DeviceApi.md#PostDeviceIpmiWhitelistResource) | **Post** /device/{deviceId}/ipmi/whitelist/ | Whitelist IP for IPMI
[**PostDevicePortIdClearResource**](DeviceApi.md#PostDevicePortIdClearResource) | **Post** /device/{deviceId}/port/{portId}/clear | Clear all Port configurations
[**PostDeviceReloadResource**](DeviceApi.md#PostDeviceReloadResource) | **Post** /device/{deviceId}/reload | Reload device
[**PostPowerResource**](DeviceApi.md#PostPowerResource) | **Post** /device/{deviceId}/power | Update power status
[**PostPreviewEffectiveIgnitionResource**](DeviceApi.md#PostPreviewEffectiveIgnitionResource) | **Post** /device/preview-ignition | Preview Ignition injections
[**PutClientDeviceTagOrderResource**](DeviceApi.md#PutClientDeviceTagOrderResource) | **Put** /device/tags-order | Update tag order
[**PutDeviceIdResource**](DeviceApi.md#PutDeviceIdResource) | **Put** /device/{deviceId} | Update device
[**PutDevicePortResource**](DeviceApi.md#PutDevicePortResource) | **Put** /device/{deviceId}/ports | Update port network
[**PutDeviceTagIdResource**](DeviceApi.md#PutDeviceTagIdResource) | **Put** /device/{deviceId}/tags | Update device tags
[**PutIpmiDevicesThresholdsIdResource**](DeviceApi.md#PutIpmiDevicesThresholdsIdResource) | **Put** /device/ipmi/thresholds | Bulk update IPMI thresholds
[**PutIpmiThresholdsIdResource**](DeviceApi.md#PutIpmiThresholdsIdResource) | **Put** /device/{deviceId}/ipmi/thresholds | Update IPMI thresholds


# **DeleteDeviceBondResource**
> NetworkTaskDump DeleteDeviceBondResource(ctx, deviceId, optional)
Unbond ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***DeviceApiDeleteDeviceBondResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiDeleteDeviceBondResourceOpts struct

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

# **GetAllDeviceTagOrderResource**
> DeviceTag GetAllDeviceTagOrderResource(ctx, optional)
Get all tag orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceApiGetAllDeviceTagOrderResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetAllDeviceTagOrderResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientDeviceTagOrderResource**
> DeviceTag GetClientDeviceTagOrderResource(ctx, optional)
Get tag order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceApiGetClientDeviceTagOrderResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetClientDeviceTagOrderResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientDeviceTagResource**
> DeviceTag GetClientDeviceTagResource(ctx, optional)
Get all tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceApiGetClientDeviceTagResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetClientDeviceTagResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceIdEventResource**
> []DeviceEvent GetDeviceIdEventResource(ctx, deviceId, optional)
Get all events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**|  | 
 **optional** | ***DeviceApiGetDeviceIdEventResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDeviceIdEventResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]DeviceEvent**](DeviceEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceIdResource**
> DeviceDump GetDeviceIdResource(ctx, deviceId, optional)
Get device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View / Update | 
 **optional** | ***DeviceApiGetDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDeviceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceDump**](Device-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceIdServiceResource**
> Services GetDeviceIdServiceResource(ctx, deviceId, optional)
Get all device services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to get all services by type | 
 **optional** | ***DeviceApiGetDeviceIdServiceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDeviceIdServiceResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **optional.String**| Get results grouped by service type code &#39;type_code&#39;, or &#39;flat&#39; | [default to flat]
 **typeCode** | **optional.String**| Return service having the same service type code, the default value is all  The list of service types can be accessed on https://core.hivelocity.net/api/v2/service/types | [default to all]

### Return type

[**Services**](Services.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceIpAssignmentsResource**
> []IpAssignment GetDeviceIpAssignmentsResource(ctx, deviceId, optional)
Get IPs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***DeviceApiGetDeviceIpAssignmentsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDeviceIpAssignmentsResourceOpts struct

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

# **GetDeviceIpminatRuleResource**
> GetDeviceIpminatRuleResource(ctx, deviceId)
Clear IPMI whitelist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of a client Device | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicePortResource**
> []DevicePort GetDevicePortResource(ctx, deviceId, optional)
Get ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***DeviceApiGetDevicePortResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDevicePortResourceOpts struct

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

# **GetDeviceResource**
> []DeviceDump GetDeviceResource(ctx, optional)
Get all devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceApiGetDeviceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDeviceResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rackId** | **optional.Int32**| Filter Devices only in this Rack ID | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]DeviceDump**](Device-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceTagIdResource**
> DeviceTag GetDeviceTagIdResource(ctx, deviceId, optional)
Get device tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View / Update | 
 **optional** | ***DeviceApiGetDeviceTagIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetDeviceTagIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveIgnitionIdResource**
> EffectiveIgnitionResponse GetEffectiveIgnitionIdResource(ctx, deviceId, optional)
Get Ignition injections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***DeviceApiGetEffectiveIgnitionIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetEffectiveIgnitionIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**EffectiveIgnitionResponse**](EffectiveIgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInitialCredsIdResource**
> DeviceInitialCreds GetInitialCredsIdResource(ctx, deviceId, optional)
Get initial creds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to retrieve initial authentication credentials for | 
 **optional** | ***DeviceApiGetInitialCredsIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetInitialCredsIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceInitialCreds**](DeviceInitialCreds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInitialPasswordIdResource**
> DevicePassword GetInitialPasswordIdResource(ctx, deviceId, optional)
Get initial password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to retrieve initial password | 
 **optional** | ***DeviceApiGetInitialPasswordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetInitialPasswordIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DevicePassword**](DevicePassword.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpmiInfoIdResource**
> DeviceIpmiInfo GetIpmiInfoIdResource(ctx, deviceId, optional)
Get IPMI data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to retrieve IPMI info. | 
 **optional** | ***DeviceApiGetIpmiInfoIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetIpmiInfoIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceIpmiInfo**](DeviceIPMIInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpmiInfoLoginDataResource**
> IpmiLoginData GetIpmiInfoLoginDataResource(ctx, deviceId, optional)
Get IPMI creds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to retrieve IPMI Login data. | 
 **optional** | ***DeviceApiGetIpmiInfoLoginDataResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetIpmiInfoLoginDataResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpmiLoginData**](IPMILoginData.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpmiThresholdsIdResource**
> DeviceIpmiThresholds GetIpmiThresholdsIdResource(ctx, deviceId, optional)
Get IPMI thresholds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View | 
 **optional** | ***DeviceApiGetIpmiThresholdsIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetIpmiThresholdsIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceIpmiThresholds**](DeviceIPMIThresholds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpmiValidLoginIdResource**
> IpmiValidLogin GetIpmiValidLoginIdResource(ctx, deviceId, optional)
Get IPMI access status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to check IPMI credentials | 
 **optional** | ***DeviceApiGetIpmiValidLoginIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetIpmiValidLoginIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**IpmiValidLogin**](IPMIValidLogin.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPowerResource**
> DevicePower GetPowerResource(ctx, deviceId, optional)
Get power status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View / Update | 
 **optional** | ***DeviceApiGetPowerResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiGetPowerResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DevicePower**](DevicePower.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDeviceBondResource**
> NetworkTaskDump PostDeviceBondResource(ctx, deviceId, optional)
Bond ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***DeviceApiPostDeviceBondResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPostDeviceBondResourceOpts struct

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

# **PostDeviceIpmiWhitelistResource**
> PostDeviceIpmiWhitelistResource(ctx, deviceId, payload)
Whitelist IP for IPMI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of the Device to put IP in Whitelist | 
  **payload** | [**DeviceIpmiWhitelistIp**](DeviceIpmiWhitelistIp.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDevicePortIdClearResource**
> NetworkTaskDump PostDevicePortIdClearResource(ctx, deviceId, portId, optional)
Clear all Port configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **portId** | **int32**|  | 
 **optional** | ***DeviceApiPostDevicePortIdClearResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPostDevicePortIdClearResourceOpts struct

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

# **PostDeviceReloadResource**
> ServiceOption PostDeviceReloadResource(ctx, deviceId, payload, optional)
Reload device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **payload** | [**DeviceReload**](DeviceReload.md)|  | 
 **optional** | ***DeviceApiPostDeviceReloadResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPostDeviceReloadResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ServiceOption**](ServiceOption.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPowerResource**
> DevicePower PostPowerResource(ctx, deviceId, action, optional)
Update power status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View / Update | 
  **action** | **string**| Must be one of boot|reboot|shutdown | 
 **optional** | ***DeviceApiPostPowerResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPostPowerResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DevicePower**](DevicePower.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPreviewEffectiveIgnitionResource**
> PreviewEffectiveIgnitionResponse PostPreviewEffectiveIgnitionResource(ctx, payload, optional)
Preview Ignition injections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**PreviewEffectiveIgnition**](PreviewEffectiveIgnition.md)|  | 
 **optional** | ***DeviceApiPostPreviewEffectiveIgnitionResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPostPreviewEffectiveIgnitionResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**PreviewEffectiveIgnitionResponse**](PreviewEffectiveIgnitionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutClientDeviceTagOrderResource**
> DeviceTag PutClientDeviceTagOrderResource(ctx, payload, optional)
Update tag order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**DeviceTag**](DeviceTag.md)|  | 
 **optional** | ***DeviceApiPutClientDeviceTagOrderResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPutClientDeviceTagOrderResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDeviceIdResource**
> DeviceDump PutDeviceIdResource(ctx, deviceId, payload, optional)
Update device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View / Update | 
  **payload** | [**DeviceUpdate**](DeviceUpdate.md)|  | 
 **optional** | ***DeviceApiPutDeviceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPutDeviceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceDump**](Device-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDevicePortResource**
> []NetworkTaskDump PutDevicePortResource(ctx, deviceId, payload, optional)
Update port network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **payload** | [**DevicePortsUpdate**](DevicePortsUpdate.md)|  | 
 **optional** | ***DeviceApiPutDevicePortResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPutDevicePortResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDeviceTagIdResource**
> DeviceTag PutDeviceTagIdResource(ctx, deviceId, payload, optional)
Update device tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View / Update | 
  **payload** | [**DeviceTag**](DeviceTag.md)|  | 
 **optional** | ***DeviceApiPutDeviceTagIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPutDeviceTagIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceTag**](DeviceTag.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIpmiDevicesThresholdsIdResource**
> DevicesIpmiThresholds PutIpmiDevicesThresholdsIdResource(ctx, payload, optional)
Bulk update IPMI thresholds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**UpdateDevicesIpmiThresholds**](UpdateDevicesIpmiThresholds.md)|  | 
 **optional** | ***DeviceApiPutIpmiDevicesThresholdsIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPutIpmiDevicesThresholdsIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DevicesIpmiThresholds**](DevicesIPMIThresholds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIpmiThresholdsIdResource**
> DeviceIpmiThresholds PutIpmiThresholdsIdResource(ctx, deviceId, payload, optional)
Update IPMI thresholds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| ID of Device to View | 
  **payload** | [**DeviceIpmiThresholds**](DeviceIpmiThresholds.md)|  | 
 **optional** | ***DeviceApiPutIpmiThresholdsIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeviceApiPutIpmiThresholdsIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DeviceIpmiThresholds**](DeviceIPMIThresholds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

