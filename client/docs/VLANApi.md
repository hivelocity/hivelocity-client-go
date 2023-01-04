# \VLANApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVlanIdResource**](VLANApi.md#DeleteVlanIdResource) | **Delete** /vlan/{vlanId} | Delete VLAN
[**GetVlanIdResource**](VLANApi.md#GetVlanIdResource) | **Get** /vlan/{vlanId} | Get VLAN
[**GetVlanResource**](VLANApi.md#GetVlanResource) | **Get** /vlan/ | Get all VLANs
[**PostVlanIdClearResource**](VLANApi.md#PostVlanIdClearResource) | **Post** /vlan/{vlanId}/clear | Clear all configurations on this VLAN, including Ports, IPs and associated Static Routes
[**PostVlanResource**](VLANApi.md#PostVlanResource) | **Post** /vlan/ | Create VLAN
[**PutVlanIdResource**](VLANApi.md#PutVlanIdResource) | **Put** /vlan/{vlanId} | Modify VLAN


# **DeleteVlanIdResource**
> DeleteVlanIdResource(ctx, vlanId)
Delete VLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vlanId** | **int32**| Id of the VLAN to interact with | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVlanIdResource**
> Vlan GetVlanIdResource(ctx, vlanId, optional)
Get VLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vlanId** | **int32**| Id of the VLAN to interact with | 
 **optional** | ***VLANApiGetVlanIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VLANApiGetVlanIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vlan**](Vlan.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVlanResource**
> []Vlan GetVlanResource(ctx, optional)
Get all VLANs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VLANApiGetVlanResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VLANApiGetVlanResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Vlan**](Vlan.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVlanIdClearResource**
> NetworkTaskDump PostVlanIdClearResource(ctx, vlanId, optional)
Clear all configurations on this VLAN, including Ports, IPs and associated Static Routes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vlanId** | **int32**| Id of the VLAN to interact with | 
 **optional** | ***VLANApiPostVlanIdClearResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VLANApiPostVlanIdClearResourceOpts struct

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

# **PostVlanResource**
> Vlan PostVlanResource(ctx, payload, optional)
Create VLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**VlanCreate**](VlanCreate.md)|  | 
 **optional** | ***VLANApiPostVlanResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VLANApiPostVlanResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vlan**](Vlan.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutVlanIdResource**
> NetworkTaskDump PutVlanIdResource(ctx, vlanId, payload, optional)
Modify VLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vlanId** | **int32**| Id of the VLAN to interact with | 
  **payload** | [**VlanUpdate**](VlanUpdate.md)|  | 
 **optional** | ***VLANApiPutVlanIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VLANApiPutVlanIdResourceOpts struct

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

