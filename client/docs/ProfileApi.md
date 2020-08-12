# \ProfileApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicProfileResource**](ProfileApi.md#GetBasicProfileResource) | **Get** /profile/basic | Get Basic Profile of current user or a contact with id
[**GetProfileResource**](ProfileApi.md#GetProfileResource) | **Get** /profile/ | Get Profile of current user or a contact with id
[**PutProfileResource**](ProfileApi.md#PutProfileResource) | **Put** /profile/ | Update Profile of current user or a contact with id



## GetBasicProfileResource

> GetBasicProfileResource(ctx, optional)

Get Basic Profile of current user or a contact with id

The id is optional and if it is necessary must be sent as URL param.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBasicProfileResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBasicProfileResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **optional.String**| \&quot;ID of Contact to manage Basic Profile\&quot; | 

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


## GetProfileResource

> GetProfileResource(ctx, optional)

Get Profile of current user or a contact with id

The id is optional and if it is necessary must be sent as URL param.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProfileResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProfileResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **optional.String**| \&quot;ID of Contact to manage Profile\&quot; | 

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


## PutProfileResource

> PutProfileResource(ctx, payload, optional)

Update Profile of current user or a contact with id

The id is optional and if it is necessary must be sent as URL param.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**ProfileUpdate**](ProfileUpdate.md)|  | 
 **optional** | ***PutProfileResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutProfileResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactId** | **optional.String**| \&quot;ID of Contact to manage Profile\&quot; | 

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

