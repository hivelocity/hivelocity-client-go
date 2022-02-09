# \VLANApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVlanIdResource**](VLANApi.md#DeleteVlanIdResource) | **Delete** /vlan/{vlanId} | Remove an existing Private VLAN
[**GetVlanIdResource**](VLANApi.md#GetVlanIdResource) | **Get** /vlan/{vlanId} | Fetch information from an existing Private VLAN
[**GetVlanResource**](VLANApi.md#GetVlanResource) | **Get** /vlan/ | Return a list with all Private VLANs
[**PostVlanResource**](VLANApi.md#PostVlanResource) | **Post** /vlan/ | Create a new Private VLAN
[**PutVlanIdResource**](VLANApi.md#PutVlanIdResource) | **Put** /vlan/{vlanId} | Update an existing Private VLAN, including adding and/or removing devices from it



## DeleteVlanIdResource

> Vlan DeleteVlanIdResource(ctx, vlanId, optional)

Remove an existing Private VLAN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vlanId** | **int32**| Id of the VLAN to interact with | 
 **optional** | ***DeleteVlanIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteVlanIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vlan**](Vlan.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanIdResource

> Vlan GetVlanIdResource(ctx, vlanId, optional)

Fetch information from an existing Private VLAN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vlanId** | **int32**| Id of the VLAN to interact with | 
 **optional** | ***GetVlanIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVlanIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vlan**](Vlan.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanResource

> []Vlan GetVlanResource(ctx, optional)

Return a list with all Private VLANs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetVlanResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVlanResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Vlan**](Vlan.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVlanResource

> Vlan PostVlanResource(ctx, payload, optional)

Create a new Private VLAN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**VlanCreate**](VlanCreate.md)|  | 
 **optional** | ***PostVlanResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostVlanResourceOpts struct


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVlanIdResource

> Vlan PutVlanIdResource(ctx, vlanId, payload, optional)

Update an existing Private VLAN, including adding and/or removing devices from it

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vlanId** | **int32**| Id of the VLAN to interact with | 
**payload** | [**VlanCreate**](VlanCreate.md)|  | 
 **optional** | ***PutVlanIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutVlanIdResourceOpts struct


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

