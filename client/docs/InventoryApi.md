# \InventoryApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStockByProductResource**](InventoryApi.md#GetStockByProductResource) | **Get** /inventory/product/{productId} | Return a structured sps stock data, grouped by city or facility code for a single product
[**GetStockResource**](InventoryApi.md#GetStockResource) | **Get** /inventory/product | Return a structured sps stock data, grouped by city or facility code for all products


# **GetStockByProductResource**
> GetStockByProductResource(ctx, productId)
Return a structured sps stock data, grouped by city or facility code for a single product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int32**| Product database ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStockResource**
> GetStockResource(ctx, optional)
Return a structured sps stock data, grouped by city or facility code for all products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoryApiGetStockResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiGetStockResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **optional.String**| Filter products by location | [default to MAIN]
 **groupBy** | **optional.String**| Get results grouped by &#39;city&#39; or &#39;facility&#39; | [default to facility]

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

