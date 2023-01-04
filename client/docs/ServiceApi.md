# \ServiceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceIdResource**](ServiceApi.md#GetServiceIdResource) | **Get** /service/{serviceId} | Return a dictionary with data from a specific account service
[**GetServiceManageableResource**](ServiceApi.md#GetServiceManageableResource) | **Get** /service/{serviceId}/manageable | Return a dictionary with data verifying managed services eligibility
[**GetServiceManagedReqResource**](ServiceApi.md#GetServiceManagedReqResource) | **Get** /service/managed-requirements | Return a dictionary with managed services operating system and panel requirements
[**GetServiceResource**](ServiceApi.md#GetServiceResource) | **Get** /service/ | Return a list of all account services
[**GetServiceTypeResource**](ServiceApi.md#GetServiceTypeResource) | **Get** /service/types | Return all available service types


# **GetServiceIdResource**
> Service GetServiceIdResource(ctx, serviceId, optional)
Return a dictionary with data from a specific account service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **int32**|  | 
 **optional** | ***ServiceApiGetServiceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeZeroPriceOptions** | **optional.Bool**| Includes on the field serviceOptions the options having price equal to zero | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Service**](Service.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceManageableResource**
> ServiceManageable GetServiceManageableResource(ctx, serviceId, optional)
Return a dictionary with data verifying managed services eligibility

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **int32**|  | 
 **optional** | ***ServiceApiGetServiceManageableResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceManageableResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ServiceManageable**](ServiceManageable.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceManagedReqResource**
> ServiceManagedReq GetServiceManagedReqResource(ctx, optional)
Return a dictionary with managed services operating system and panel requirements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceApiGetServiceManagedReqResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceManagedReqResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ServiceManagedReq**](ServiceManagedReq.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceResource**
> []FastServiceDump GetServiceResource(ctx, status, optional)
Return a list of all account services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **status** | **string**| The current status of the service | [default to all]
 **optional** | ***ServiceApiGetServiceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typeCode** | **optional.String**| Return service having the same service type code, the default value is all  The list of service types can be accessed on https://core.hivelocity.net/api/v2/service/types | [default to null]
 **orderId** | **optional.Int32**| Order id of the service | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]FastServiceDump**](FastService-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceTypeResource**
> []ServiceType GetServiceTypeResource(ctx, optional)
Return all available service types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceApiGetServiceTypeResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceTypeResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]ServiceType**](ServiceType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

