# \InvoiceApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceIdDetails**](InvoiceApi.md#GetInvoiceIdDetails) | **Get** /invoice/{invoiceId}/details | Return detailed information for an invoice
[**GetInvoiceIdResource**](InvoiceApi.md#GetInvoiceIdResource) | **Get** /invoice/{invoiceId} | Return serialized data on a single invoice
[**GetInvoicePdfResource**](InvoiceApi.md#GetInvoicePdfResource) | **Get** /invoice/{invoiceId}/pdf-download | Return an Invoice PDF file in Base64 Encoded Format
[**GetInvoiceResource**](InvoiceApi.md#GetInvoiceResource) | **Get** /invoice/ | Return serialized data on all invoices



## GetInvoiceIdDetails

> GetInvoiceIdDetails(ctx, invoiceId)

Return detailed information for an invoice

Same data as the PDF file but serialized

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32**| Invoice database ID | 

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


## GetInvoiceIdResource

> GetInvoiceIdResource(ctx, invoiceId)

Return serialized data on a single invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32**| Invoice database ID | 

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


## GetInvoicePdfResource

> GetInvoicePdfResource(ctx, invoiceId)

Return an Invoice PDF file in Base64 Encoded Format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32**| Invoice database ID | 

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


## GetInvoiceResource

> GetInvoiceResource(ctx, )

Return serialized data on all invoices

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

