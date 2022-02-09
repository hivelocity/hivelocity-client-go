# \ProductApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductListResource**](ProductApi.md#GetProductListResource) | **Get** /product/list | Return structured sps stock data in a list
[**GetProductOperatingSystemsResource**](ProductApi.md#GetProductOperatingSystemsResource) | **Get** /product/{productId}/operating-systems | Return List of operating systems found for a Product
[**GetProductOptionResource**](ProductApi.md#GetProductOptionResource) | **Get** /product/{productId}/options | Return List of Options found for a Product
[**GetProductsAndOptionsResource**](ProductApi.md#GetProductsAndOptionsResource) | **Get** /product/options | Return a mapping of Products and Options with pricing per-period



## GetProductListResource

> []ProductInfo GetProductListResource(ctx, optional)

Return structured sps stock data in a list

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProductListResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProductListResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]ProductInfo**](ProductInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductOperatingSystemsResource

> []Option GetProductOperatingSystemsResource(ctx, productId, optional)

Return List of operating systems found for a Product

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32**| ID of the Product | 
 **optional** | ***GetProductOperatingSystemsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProductOperatingSystemsResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Option**](Option.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductOptionResource

> map[string][]Option GetProductOptionResource(ctx, productId, optional)

Return List of Options found for a Product

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32**| ID of the Product | 
 **optional** | ***GetProductOptionResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProductOptionResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **optional.String**| Get results grouped by &#39;upgrade&#39; or without grouping. | [default to upgrade]

### Return type

[**map[string][]Option**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductsAndOptionsResource

> []ProductOption GetProductsAndOptionsResource(ctx, optional)

Return a mapping of Products and Options with pricing per-period

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProductsAndOptionsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProductsAndOptionsResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]ProductOption**](ProductOption.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

