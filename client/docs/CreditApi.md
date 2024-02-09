# \CreditApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCreditResource**](CreditApi.md#GetCreditResource) | **Get** /credit/ | Return a list with all Credits
[**GetTotalActiveCreditResource**](CreditApi.md#GetTotalActiveCreditResource) | **Get** /credit/total | Return the client&#39;s total active credit amount
[**PostCreditApplyCouponResource**](CreditApi.md#PostCreditApplyCouponResource) | **Post** /credit/apply-coupon | Apply the coupon code for a service
[**PostCreditResource**](CreditApi.md#PostCreditResource) | **Post** /credit/ | Receive billing method id and amount and return the created Credit
[**PostCreditValidateCouponResource**](CreditApi.md#PostCreditValidateCouponResource) | **Post** /credit/validate-coupon | Validate the coupon code for a service


# **GetCreditResource**
> []Credit GetCreditResource(ctx, optional)
Return a list with all Credits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreditApiGetCreditResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiGetCreditResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| The current status of the service (all, active, inactive) | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Credit**](Credit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTotalActiveCreditResource**
> TotalActiveCredit GetTotalActiveCreditResource(ctx, optional)
Return the client's total active credit amount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreditApiGetTotalActiveCreditResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiGetTotalActiveCreditResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**TotalActiveCredit**](TotalActiveCredit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreditApplyCouponResource**
> Credit PostCreditApplyCouponResource(ctx, payload, optional)
Apply the coupon code for a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**CartCoupon**](CartCoupon.md)|  | 
 **optional** | ***CreditApiPostCreditApplyCouponResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiPostCreditApplyCouponResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreditResource**
> Credit PostCreditResource(ctx, payload, optional)
Receive billing method id and amount and return the created Credit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**CreateCredit**](CreateCredit.md)|  | 
 **optional** | ***CreditApiPostCreditResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiPostCreditResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreditValidateCouponResource**
> ValidateCartCouponDump PostCreditValidateCouponResource(ctx, payload, optional)
Validate the coupon code for a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**CartCoupon**](CartCoupon.md)|  | 
 **optional** | ***CreditApiPostCreditValidateCouponResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreditApiPostCreditValidateCouponResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ValidateCartCouponDump**](ValidateCartCoupon-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

