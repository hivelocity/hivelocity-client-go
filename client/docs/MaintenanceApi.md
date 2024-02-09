# \MaintenanceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMaintenanceEventClientResource**](MaintenanceApi.md#GetMaintenanceEventClientResource) | **Get** /maintenance/ | Get all Maintenance Events for a client
[**GetMaintenanceEventIdClientResource**](MaintenanceApi.md#GetMaintenanceEventIdClientResource) | **Get** /maintenance/{maintenanceEventId} | Get Maintenance Events


# **GetMaintenanceEventClientResource**
> []MaintenanceEvent GetMaintenanceEventClientResource(ctx, optional)
Get all Maintenance Events for a client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MaintenanceApiGetMaintenanceEventClientResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaintenanceApiGetMaintenanceEventClientResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]MaintenanceEvent**](MaintenanceEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMaintenanceEventIdClientResource**
> MaintenanceEvent GetMaintenanceEventIdClientResource(ctx, maintenanceEventId, optional)
Get Maintenance Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenanceEventId** | **int32**|  | 
 **optional** | ***MaintenanceApiGetMaintenanceEventIdClientResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaintenanceApiGetMaintenanceEventIdClientResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**MaintenanceEvent**](MaintenanceEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

