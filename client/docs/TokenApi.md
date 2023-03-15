# \TokenApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTokenResource**](TokenApi.md#DeleteTokenResource) | **Delete** /token/ | Deletes the Public API Token with apiTokenId
[**GetTokenResource**](TokenApi.md#GetTokenResource) | **Get** /token/ | Returns a list of Public API Tokens for the current user (last 4 characters only)
[**PostTokenResource**](TokenApi.md#PostTokenResource) | **Post** /token/ | Create a new Public API Token for the current user
[**PutTokenResource**](TokenApi.md#PutTokenResource) | **Put** /token/ | Updates the Public API Token with apiTokenId


# **DeleteTokenResource**
> DeleteTokenResource(ctx, payload)
Deletes the Public API Token with apiTokenId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**DeleteApiTokenLoad**](DeleteApiTokenLoad.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenResource**
> []PublicApiHideTokenDump GetTokenResource(ctx, optional)
Returns a list of Public API Tokens for the current user (last 4 characters only)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenApiGetTokenResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetTokenResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]PublicApiHideTokenDump**](PublicApiHideToken-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTokenResource**
> PublicApiTokenDump PostTokenResource(ctx, payload, optional)
Create a new Public API Token for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**PublicApiToken**](PublicApiToken.md)|  | 
 **optional** | ***TokenApiPostTokenResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiPostTokenResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**PublicApiTokenDump**](PublicApiToken-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTokenResource**
> PublicApiHideTokenDump PutTokenResource(ctx, payload, optional)
Updates the Public API Token with apiTokenId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**PublicApiTokenLoad**](PublicApiTokenLoad.md)|  | 
 **optional** | ***TokenApiPutTokenResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiPutTokenResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**PublicApiHideTokenDump**](PublicApiHideToken-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

