# \ProductApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductOperatingSystemsResource**](ProductApi.md#GetProductOperatingSystemsResource) | **Get** /product/{productId}/operating-systems | Returns List of operating systems found for a Product
[**GetProductOptionResource**](ProductApi.md#GetProductOptionResource) | **Get** /product/{productId}/options | Returns List of Options found for a Product
[**PostProductMatchResource**](ProductApi.md#PostProductMatchResource) | **Post** /product/match | Return a list of Products matching the provided lshw output of a server



## GetProductOperatingSystemsResource

> GetProductOperatingSystemsResource(ctx, productId)

Returns List of operating systems found for a Product

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32**| ID of the Product | 

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


## GetProductOptionResource

> GetProductOptionResource(ctx, productId)

Returns List of Options found for a Product

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32**| ID of the Product | 

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


## PostProductMatchResource

> PostProductMatchResource(ctx, payload)

Return a list of Products matching the provided lshw output of a server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**ProductMatch**](ProductMatch.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

