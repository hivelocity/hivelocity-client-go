# \CreditApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCreditResource**](CreditApi.md#GetCreditResource) | **Get** /credit/ | Return a list with all Credits
[**GetTotalActiveCreditResource**](CreditApi.md#GetTotalActiveCreditResource) | **Get** /credit/total | Return the client&#39;s total active credit amount
[**PostCreditResource**](CreditApi.md#PostCreditResource) | **Post** /credit/ | Receive billing method id and amount and return the created Credit



## GetCreditResource

> GetCreditResource(ctx, )

Return a list with all Credits

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


## GetTotalActiveCreditResource

> GetTotalActiveCreditResource(ctx, )

Return the client's total active credit amount

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


## PostCreditResource

> PostCreditResource(ctx, payload)

Receive billing method id and amount and return the created Credit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**CreateCredit**](CreateCredit.md)|  | 

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

