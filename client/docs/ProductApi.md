# \ProductApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductOperatingSystemsResource**](ProductApi.md#GetProductOperatingSystemsResource) | **Get** /product/{productId}/operating-systems | Get product operating systems
[**GetProductOptionResource**](ProductApi.md#GetProductOptionResource) | **Get** /product/{productId}/options | Get product options
[**GetProductsAndOptionsResource**](ProductApi.md#GetProductsAndOptionsResource) | **Get** /product/options | Get all options


# **GetProductOperatingSystemsResource**
> []Option GetProductOperatingSystemsResource(ctx, productId, optional)
Get product operating systems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int32**| ID of the Product | 
 **optional** | ***ProductApiGetProductOperatingSystemsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiGetProductOperatingSystemsResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Option**](Option.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductOptionResource**
> Options GetProductOptionResource(ctx, productId, optional)
Get product options

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int32**| ID of the Product | 
 **optional** | ***ProductApiGetProductOptionResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiGetProductOptionResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **optional.String**| Get results grouped by &#39;upgrade&#39; or without grouping. | [default to upgrade]

### Return type

[**Options**](Options.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductsAndOptionsResource**
> []ProductOption GetProductsAndOptionsResource(ctx, optional)
Get all options

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiGetProductsAndOptionsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiGetProductsAndOptionsResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]ProductOption**](ProductOption.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

