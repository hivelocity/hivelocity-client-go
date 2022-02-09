# \BillingInfoApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingInfoResource**](BillingInfoApi.md#GetBillingInfoResource) | **Get** /billing-info/ | Return a list with all Billing Info
[**PostBillingInfoResource**](BillingInfoApi.md#PostBillingInfoResource) | **Post** /billing-info/ | Create verification for credit card with all Billing Info
[**PutBillingInfoResource**](BillingInfoApi.md#PutBillingInfoResource) | **Put** /billing-info/ | Verify credit card with all Billing Info



## GetBillingInfoResource

> []BillingInfo GetBillingInfoResource(ctx, optional)

Return a list with all Billing Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBillingInfoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBillingInfoResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]BillingInfo**](BillingInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBillingInfoResource

> BillingInfo PostBillingInfoResource(ctx, payload, optional)

Create verification for credit card with all Billing Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**BillingInfoVerification**](BillingInfoVerification.md)|  | 
 **optional** | ***PostBillingInfoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostBillingInfoResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**BillingInfo**](BillingInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBillingInfoResource

> Credit PutBillingInfoResource(ctx, payload, optional)

Verify credit card with all Billing Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**BillingInfoVerification**](BillingInfoVerification.md)|  | 
 **optional** | ***PutBillingInfoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutBillingInfoResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Credit**](Credit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

