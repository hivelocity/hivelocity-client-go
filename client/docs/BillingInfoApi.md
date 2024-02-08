# \BillingInfoApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBillingInfoIdResource**](BillingInfoApi.md#DeleteBillingInfoIdResource) | **Delete** /billing-info/{billingInfoId} | Delete a Payment Method by Billing Info ID
[**GetBillingInfoIdResource**](BillingInfoApi.md#GetBillingInfoIdResource) | **Get** /billing-info/{billingInfoId} | Return a Payment Method by Billing Info ID
[**GetBillingInfoResource**](BillingInfoApi.md#GetBillingInfoResource) | **Get** /billing-info/ | Return a list with all Payment Methods (Billing Info)
[**GetPaymentMethodsEnableReadResource**](BillingInfoApi.md#GetPaymentMethodsEnableReadResource) | **Get** /billing-info/enable | Return Payment Methods enabled for editing
[**PostBillingInfoBankAccountCreateResource**](BillingInfoApi.md#PostBillingInfoBankAccountCreateResource) | **Post** /billing-info/bank-account | Add a new Bank Account as a Payment Method
[**PostBillingInfoCreditCardCreateResource**](BillingInfoApi.md#PostBillingInfoCreditCardCreateResource) | **Post** /billing-info/credit-card | Add a new Credit Card as a Payment Method
[**PostBillingInfoResource**](BillingInfoApi.md#PostBillingInfoResource) | **Post** /billing-info/ | Create verification for credit card with all Billing Info
[**PutBillingInfoBankAccountUpdateResource**](BillingInfoApi.md#PutBillingInfoBankAccountUpdateResource) | **Put** /billing-info/bank-account/{billingInfoId} | Update a Bank Account
[**PutBillingInfoCreditCardUpdateResource**](BillingInfoApi.md#PutBillingInfoCreditCardUpdateResource) | **Put** /billing-info/credit-card/{billingInfoId} | Update a Credit Card
[**PutBillingInfoResource**](BillingInfoApi.md#PutBillingInfoResource) | **Put** /billing-info/ | Verify credit card with all Billing Info


# **DeleteBillingInfoIdResource**
> DeleteBillingInfoIdResource(ctx, billingInfoId, optional)
Delete a Payment Method by Billing Info ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **billingInfoId** | **int32**|  | 
 **optional** | ***BillingInfoApiDeleteBillingInfoIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiDeleteBillingInfoIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingInfoIdResource**
> BillingInfo GetBillingInfoIdResource(ctx, billingInfoId, optional)
Return a Payment Method by Billing Info ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **billingInfoId** | **int32**|  | 
 **optional** | ***BillingInfoApiGetBillingInfoIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiGetBillingInfoIdResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingInfoResource**
> []BillingInfo GetBillingInfoResource(ctx, optional)
Return a list with all Payment Methods (Billing Info)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingInfoApiGetBillingInfoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiGetBillingInfoResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.Int32**| Client ID to fetch billing info for | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]BillingInfo**](BillingInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentMethodsEnableReadResource**
> PaymentMethodsEnabled GetPaymentMethodsEnableReadResource(ctx, optional)
Return Payment Methods enabled for editing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingInfoApiGetPaymentMethodsEnableReadResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiGetPaymentMethodsEnableReadResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**PaymentMethodsEnabled**](PaymentMethodsEnabled.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBillingInfoBankAccountCreateResource**
> BillingInfo PostBillingInfoBankAccountCreateResource(ctx, payload, optional)
Add a new Bank Account as a Payment Method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**BillingInfoCreateBankAccount**](BillingInfoCreateBankAccount.md)|  | 
 **optional** | ***BillingInfoApiPostBillingInfoBankAccountCreateResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiPostBillingInfoBankAccountCreateResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBillingInfoCreditCardCreateResource**
> BillingInfo PostBillingInfoCreditCardCreateResource(ctx, payload, optional)
Add a new Credit Card as a Payment Method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**BillingInfoCreateCreditCard**](BillingInfoCreateCreditCard.md)|  | 
 **optional** | ***BillingInfoApiPostBillingInfoCreditCardCreateResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiPostBillingInfoCreditCardCreateResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBillingInfoResource**
> BillingInfo PostBillingInfoResource(ctx, payload, optional)
Create verification for credit card with all Billing Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**BillingInfoVerification**](BillingInfoVerification.md)|  | 
 **optional** | ***BillingInfoApiPostBillingInfoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiPostBillingInfoResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBillingInfoBankAccountUpdateResource**
> BillingInfo PutBillingInfoBankAccountUpdateResource(ctx, billingInfoId, payload, optional)
Update a Bank Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **billingInfoId** | **int32**|  | 
  **payload** | [**BillingInfoCreateBankAccount**](BillingInfoCreateBankAccount.md)|  | 
 **optional** | ***BillingInfoApiPutBillingInfoBankAccountUpdateResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiPutBillingInfoBankAccountUpdateResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBillingInfoCreditCardUpdateResource**
> BillingInfo PutBillingInfoCreditCardUpdateResource(ctx, billingInfoId, payload, optional)
Update a Credit Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **billingInfoId** | **int32**|  | 
  **payload** | [**BillingInfoCreateCreditCard**](BillingInfoCreateCreditCard.md)|  | 
 **optional** | ***BillingInfoApiPutBillingInfoCreditCardUpdateResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiPutBillingInfoCreditCardUpdateResourceOpts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutBillingInfoResource**
> Credit PutBillingInfoResource(ctx, payload, optional)
Verify credit card with all Billing Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**BillingInfoVerification**](BillingInfoVerification.md)|  | 
 **optional** | ***BillingInfoApiPutBillingInfoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingInfoApiPutBillingInfoResourceOpts struct

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

