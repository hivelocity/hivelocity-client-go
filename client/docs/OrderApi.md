# \OrderApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderIdResource**](OrderApi.md#GetOrderIdResource) | **Get** /order/{orderId} | Return details of a specific Order
[**GetOrderResource**](OrderApi.md#GetOrderResource) | **Get** /order/ | Return a list with all Orders



## GetOrderIdResource

> GetOrderIdResource(ctx, orderId)

Return details of a specific Order

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32**| \&quot;ID of Order to View\&quot; | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderResource

> GetOrderResource(ctx, )

Return a list with all Orders

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

