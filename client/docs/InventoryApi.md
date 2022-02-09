# \InventoryApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocationResource**](InventoryApi.md#GetLocationResource) | **Get** /inventory/locations | Return sps facilities and locations
[**GetStockByProductResource**](InventoryApi.md#GetStockByProductResource) | **Get** /inventory/product/{productId} | Return a structured sps stock data, grouped by city or facility code for a single product
[**GetStockResource**](InventoryApi.md#GetStockResource) | **Get** /inventory/product | Return structured sps stock data, grouped by city or facility code for all products



## GetLocationResource

> []Location GetLocationResource(ctx, optional)

Return sps facilities and locations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLocationResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLocationResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Location**](Location.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockByProductResource

> Stock GetStockByProductResource(ctx, productId, optional)

Return a structured sps stock data, grouped by city or facility code for a single product

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32**| Product database ID | 
 **optional** | ***GetStockByProductResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStockByProductResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Stock**](Stock.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockResource

> map[string][]Stock GetStockResource(ctx, optional)

Return structured sps stock data, grouped by city or facility code for all products

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetStockResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStockResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **optional.String**| Filter products by location | [default to MAIN]
 **groupBy** | **optional.String**| Get results grouped by &#39;city&#39;, &#39;facility&#39;, or &#39;flat&#39; | [default to facility]

### Return type

[**map[string][]Stock**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

