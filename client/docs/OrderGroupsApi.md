# \OrderGroupsApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrderGroupIdResource**](OrderGroupsApi.md#DeleteOrderGroupIdResource) | **Delete** /order-groups/{orderGroupId} | Delete the specified Order Group
[**GetOrderGroupIdResource**](OrderGroupsApi.md#GetOrderGroupIdResource) | **Get** /order-groups/{orderGroupId} | Return an Order Group
[**GetOrderGroupResource**](OrderGroupsApi.md#GetOrderGroupResource) | **Get** /order-groups/ | Return a list with all Order Groups
[**PostOrderGroupResource**](OrderGroupsApi.md#PostOrderGroupResource) | **Post** /order-groups/ | Create a new Order Group
[**PutOrderGroupIdResource**](OrderGroupsApi.md#PutOrderGroupIdResource) | **Put** /order-groups/{orderGroupId} | Update an Order Group


# **DeleteOrderGroupIdResource**
> DeleteOrderGroupIdResource(ctx, orderGroupId)
Delete the specified Order Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderGroupId** | **int32**| Id of the order group to interact with | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderGroupIdResource**
> OrderGroup GetOrderGroupIdResource(ctx, orderGroupId, optional)
Return an Order Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderGroupId** | **int32**| Id of the order group to interact with | 
 **optional** | ***OrderGroupsApiGetOrderGroupIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderGroupsApiGetOrderGroupIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**OrderGroup**](OrderGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderGroupResource**
> []OrderGroup GetOrderGroupResource(ctx, optional)
Return a list with all Order Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderGroupsApiGetOrderGroupResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderGroupsApiGetOrderGroupResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]OrderGroup**](OrderGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOrderGroupResource**
> OrderGroup PostOrderGroupResource(ctx, payload, optional)
Create a new Order Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**OrderGroupCreate**](OrderGroupCreate.md)|  | 
 **optional** | ***OrderGroupsApiPostOrderGroupResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderGroupsApiPostOrderGroupResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**OrderGroup**](OrderGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutOrderGroupIdResource**
> OrderGroup PutOrderGroupIdResource(ctx, orderGroupId, payload, optional)
Update an Order Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderGroupId** | **int32**| Id of the order group to interact with | 
  **payload** | [**OrderGroupUpdate**](OrderGroupUpdate.md)|  | 
 **optional** | ***OrderGroupsApiPutOrderGroupIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderGroupsApiPutOrderGroupIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**OrderGroup**](OrderGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

