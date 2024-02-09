# \InventoryApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocationResource**](InventoryApi.md#GetLocationResource) | **Get** /inventory/locations | Get all facilities
[**GetStockByProductResource**](InventoryApi.md#GetStockByProductResource) | **Get** /inventory/product/{productId} | Get product
[**GetStockResource**](InventoryApi.md#GetStockResource) | **Get** /inventory/product | Get all products


# **GetLocationResource**
> []Location GetLocationResource(ctx, optional)
Get all facilities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoryApiGetLocationResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiGetLocationResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpsLocations** | **optional.String**| Filter by the VPS locations | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Location**](Location.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStockByProductResource**
> Stock GetStockByProductResource(ctx, productId, optional)
Get product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int32**| Product database ID | 
 **optional** | ***InventoryApiGetStockByProductResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiGetStockByProductResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Stock**](Stock.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStockResource**
> Inventory GetStockResource(ctx, optional)
Get all products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoryApiGetStockResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiGetStockResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **optional.String**| Retrieve products by their location in the Hivelocity store:   MAIN: The main list of instant servers. CUSTOM_DISPLAY: The main list of custom servers. MASSIVESTORAGE: Servers with up to 90 drives. GPU: Servers with GPUs. OUTLET: Discounted older hardware. LANDING: Periodic specials for events. | [default to MAIN]
 **includeCustom** | **optional.Bool**| Include custom server products. | [default to false]
 **bondingSupport** | **optional.String**| Filter by bonding support. Values are true/false. None will return a mix of both. *DEPRECATED:* Soon, all servers sold will have bonding support and this will be removed. | 
 **groupBy** | **optional.String**| Get results grouped by &#39;city&#39;, &#39;facility&#39;, or &#39;flat&#39; | [default to facility]

### Return type

[**Inventory**](Inventory.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

