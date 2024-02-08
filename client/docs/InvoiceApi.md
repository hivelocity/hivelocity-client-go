# \InvoiceApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceIdDetails**](InvoiceApi.md#GetInvoiceIdDetails) | **Get** /invoice/{invoiceId}/details | Return detailed information for an invoice
[**GetInvoiceIdResource**](InvoiceApi.md#GetInvoiceIdResource) | **Get** /invoice/{invoiceId} | Return serialized data on a single invoice
[**GetInvoicePdfResource**](InvoiceApi.md#GetInvoicePdfResource) | **Get** /invoice/{invoiceId}/pdf-download | Return an Invoice PDF file in Base64 Encoded Format
[**GetInvoiceResource**](InvoiceApi.md#GetInvoiceResource) | **Get** /invoice/ | Return serialized data on all invoices
[**GetInvoiceSearchResource**](InvoiceApi.md#GetInvoiceSearchResource) | **Get** /invoice/search | Return results of invoice search
[**GetInvoiceUnpaidResource**](InvoiceApi.md#GetInvoiceUnpaidResource) | **Get** /invoice/unpaid | Return total balance of all unpaid invoices
[**PostInvoiceIdApplyAccountCredit**](InvoiceApi.md#PostInvoiceIdApplyAccountCredit) | **Post** /invoice/{invoiceId}/apply-account-credit | Apply account credit to an invoice


# **GetInvoiceIdDetails**
> InvoiceDetails GetInvoiceIdDetails(ctx, invoiceId, optional)
Return detailed information for an invoice

Same data as the PDF file but serialized

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int32**| Invoice database ID | 
 **optional** | ***InvoiceApiGetInvoiceIdDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetInvoiceIdDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**InvoiceDetails**](InvoiceDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceIdResource**
> Invoice GetInvoiceIdResource(ctx, invoiceId, optional)
Return serialized data on a single invoice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int32**| Invoice database ID | 
 **optional** | ***InvoiceApiGetInvoiceIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetInvoiceIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoicePdfResource**
> InvoicePdf GetInvoicePdfResource(ctx, invoiceId, optional)
Return an Invoice PDF file in Base64 Encoded Format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int32**| Invoice database ID | 
 **optional** | ***InvoiceApiGetInvoicePdfResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetInvoicePdfResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**InvoicePdf**](InvoicePDF.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceResource**
> []Invoice GetInvoiceResource(ctx, optional)
Return serialized data on all invoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoiceApiGetInvoiceResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetInvoiceResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceSearchResource**
> InvoiceSeachResult GetInvoiceSearchResource(ctx, optional)
Return results of invoice search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoiceApiGetInvoiceSearchResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetInvoiceSearchResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endDate** | **optional.Int32**| The end date range | 
 **startDate** | **optional.Int32**| The start date range | 
 **perPage** | **optional.Int32**| Number of items per page. | [default to 10]
 **page** | **optional.Int32**| The page number of search. | [default to 1]
 **q** | **optional.String**| Content search. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**InvoiceSeachResult**](InvoiceSeachResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceUnpaidResource**
> InvoiceUnpaid GetInvoiceUnpaidResource(ctx, optional)
Return total balance of all unpaid invoices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoiceApiGetInvoiceUnpaidResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiGetInvoiceUnpaidResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**InvoiceUnpaid**](InvoiceUnpaid.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInvoiceIdApplyAccountCredit**
> Invoice PostInvoiceIdApplyAccountCredit(ctx, invoiceId, optional)
Apply account credit to an invoice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **int32**| Invoice database ID | 
 **optional** | ***InvoiceApiPostInvoiceIdApplyAccountCreditOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceApiPostInvoiceIdApplyAccountCreditOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

