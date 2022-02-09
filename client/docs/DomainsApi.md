# \DomainsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteARecordIdResource**](DomainsApi.md#DeleteARecordIdResource) | **Delete** /domains/{domain}/a-record/{record} | Delete a single A Record value for given zone
[**DeleteAaaaRecordIdResource**](DomainsApi.md#DeleteAaaaRecordIdResource) | **Delete** /domains/{domainID}/aaaa-record/{recordId} | Delete an AAAA Record for given zone
[**DeleteDomainIdResource**](DomainsApi.md#DeleteDomainIdResource) | **Delete** /domains/{domainId} | Delete domain for given id
[**DeleteMxRecordIdResource**](DomainsApi.md#DeleteMxRecordIdResource) | **Delete** /domains/{domainID}/mx-record/{recordId} | Delete an MX Record for given zone
[**GetARecordIdResource**](DomainsApi.md#GetARecordIdResource) | **Get** /domains/{domain}/a-record/{record} | Return Single A Records found for given domain name
[**GetARecordResource**](DomainsApi.md#GetARecordResource) | **Get** /domains/{domain}/a-record | Return all A Records found for given zone name
[**GetAaaaRecordIdResource**](DomainsApi.md#GetAaaaRecordIdResource) | **Get** /domains/{domainID}/aaaa-record/{recordId} | Return Single AAAA Records found for given id
[**GetAaaaRecordResource**](DomainsApi.md#GetAaaaRecordResource) | **Get** /domains/{domainID}/aaaa-record | Return all AAAA Records found for given zone
[**GetDomainIdResource**](DomainsApi.md#GetDomainIdResource) | **Get** /domains/{domainId} | Return a domain for given id
[**GetDomainResource**](DomainsApi.md#GetDomainResource) | **Get** /domains/ | Return the domains of a client
[**GetMxRecordIdResource**](DomainsApi.md#GetMxRecordIdResource) | **Get** /domains/{domainID}/mx-record/{recordId} | Return Single MX Records found for given id
[**GetMxRecordResource**](DomainsApi.md#GetMxRecordResource) | **Get** /domains/{domainID}/mx-record | Return all MX Records found for given zone
[**GetPtrRecordIdResource**](DomainsApi.md#GetPtrRecordIdResource) | **Get** /domains/ptr/{recordId} | 
[**GetPtrRecordResource**](DomainsApi.md#GetPtrRecordResource) | **Get** /domains/ptr | Return the PTR records of a client
[**PostARecordResource**](DomainsApi.md#PostARecordResource) | **Post** /domains/{domain}/a-record | Create a new A Record for given zone
[**PostAaaaRecordResource**](DomainsApi.md#PostAaaaRecordResource) | **Post** /domains/{domainID}/aaaa-record | Create a new AAAA Record for given zone
[**PostDomainResource**](DomainsApi.md#PostDomainResource) | **Post** /domains/ | Create a new domain
[**PostMxRecordResource**](DomainsApi.md#PostMxRecordResource) | **Post** /domains/{domainID}/mx-record | Create a new MX Record for given zone
[**PutARecordIdResource**](DomainsApi.md#PutARecordIdResource) | **Put** /domains/{domain}/a-record/{record} | Update an A Record for given zone
[**PutAaaaRecordIdResource**](DomainsApi.md#PutAaaaRecordIdResource) | **Put** /domains/{domainID}/aaaa-record/{recordId} | Update an AAAA Record for given id
[**PutMxRecordIdResource**](DomainsApi.md#PutMxRecordIdResource) | **Put** /domains/{domainID}/mx-record/{recordId} | Update an MX Record for given zone
[**PutPtrRecordIdResource**](DomainsApi.md#PutPtrRecordIdResource) | **Put** /domains/ptr/{recordId} | 



## DeleteARecordIdResource

> DeleteARecordIdResource(ctx, domain, record)

Delete a single A Record value for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
**record** | **string**|  | 

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


## DeleteAaaaRecordIdResource

> DeleteAaaaRecordIdResource(ctx, domainID, recordId)

Delete an AAAA Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**recordId** | **int32**|  | 

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


## DeleteDomainIdResource

> DeleteDomainIdResource(ctx, domainId)

Delete domain for given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**|  | 

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


## DeleteMxRecordIdResource

> DeleteMxRecordIdResource(ctx, domainID, recordId)

Delete an MX Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**recordId** | **int32**|  | 

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


## GetARecordIdResource

> ARecord GetARecordIdResource(ctx, domain, record, optional)

Return Single A Records found for given domain name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
**record** | **string**|  | 
 **optional** | ***GetARecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetARecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ARecord**](ARecord.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetARecordResource

> []ARecord GetARecordResource(ctx, domain, optional)

Return all A Records found for given zone name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
 **optional** | ***GetARecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetARecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]ARecord**](ARecord.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAaaaRecordIdResource

> AaaaRecordReturn GetAaaaRecordIdResource(ctx, domainID, recordId, optional)

Return Single AAAA Records found for given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**recordId** | **int32**|  | 
 **optional** | ***GetAaaaRecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAaaaRecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**AaaaRecordReturn**](AAAARecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAaaaRecordResource

> []AaaaRecordReturn GetAaaaRecordResource(ctx, domainID, optional)

Return all AAAA Records found for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
 **optional** | ***GetAaaaRecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAaaaRecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]AaaaRecordReturn**](AAAARecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainIdResource

> DomainReturn GetDomainIdResource(ctx, domainId, optional)

Return a domain for given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**|  | 
 **optional** | ***GetDomainIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DomainReturn**](DomainReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainResource

> []DomainReturn GetDomainResource(ctx, optional)

Return the domains of a client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDomainResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]DomainReturn**](DomainReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMxRecordIdResource

> MxRecordReturn GetMxRecordIdResource(ctx, domainID, recordId, optional)

Return Single MX Records found for given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**recordId** | **int32**|  | 
 **optional** | ***GetMxRecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMxRecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**MxRecordReturn**](MXRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMxRecordResource

> []MxRecordReturn GetMxRecordResource(ctx, domainID, optional)

Return all MX Records found for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
 **optional** | ***GetMxRecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMxRecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]MxRecordReturn**](MXRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPtrRecordIdResource

> PtrRecordReturn GetPtrRecordIdResource(ctx, recordId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **int32**|  | 
 **optional** | ***GetPtrRecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPtrRecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**PtrRecordReturn**](PTRRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPtrRecordResource

> []PtrRecordReturn GetPtrRecordResource(ctx, optional)

Return the PTR records of a client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPtrRecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPtrRecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]PtrRecordReturn**](PTRRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostARecordResource

> ARecord PostARecordResource(ctx, domain, payload, optional)

Create a new A Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
**payload** | [**ARecord**](ARecord.md)|  | 
 **optional** | ***PostARecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostARecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ARecord**](ARecord.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAaaaRecordResource

> AaaaRecordReturn PostAaaaRecordResource(ctx, domainID, payload, optional)

Create a new AAAA Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**payload** | [**AaaaRecordCreate**](AaaaRecordCreate.md)|  | 
 **optional** | ***PostAaaaRecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAaaaRecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**AaaaRecordReturn**](AAAARecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDomainResource

> DomainReturn PostDomainResource(ctx, payload, optional)

Create a new domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**DomainCreate**](DomainCreate.md)|  | 
 **optional** | ***PostDomainResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostDomainResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**DomainReturn**](DomainReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMxRecordResource

> MxRecordReturn PostMxRecordResource(ctx, domainID, payload, optional)

Create a new MX Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**payload** | [**MxRecordCreate**](MxRecordCreate.md)|  | 
 **optional** | ***PostMxRecordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostMxRecordResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**MxRecordReturn**](MXRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutARecordIdResource

> ARecord PutARecordIdResource(ctx, domain, record, payload, optional)

Update an A Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
**record** | **string**|  | 
**payload** | [**ARecord**](ARecord.md)|  | 
 **optional** | ***PutARecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutARecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**ARecord**](ARecord.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAaaaRecordIdResource

> AaaaRecordReturn PutAaaaRecordIdResource(ctx, domainID, recordId, payload, optional)

Update an AAAA Record for given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**recordId** | **int32**|  | 
**payload** | [**AaaaRecordUpdate**](AaaaRecordUpdate.md)|  | 
 **optional** | ***PutAaaaRecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutAaaaRecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**AaaaRecordReturn**](AAAARecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMxRecordIdResource

> MxRecordReturn PutMxRecordIdResource(ctx, domainID, recordId, payload, optional)

Update an MX Record for given zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainID** | **int32**|  | 
**recordId** | **int32**|  | 
**payload** | [**MxRecordUpdate**](MxRecordUpdate.md)|  | 
 **optional** | ***PutMxRecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutMxRecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**MxRecordReturn**](MXRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPtrRecordIdResource

> PtrRecordReturn PutPtrRecordIdResource(ctx, recordId, payload, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **int32**|  | 
**payload** | [**PtrRecordUpdate**](PtrRecordUpdate.md)|  | 
 **optional** | ***PutPtrRecordIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutPtrRecordIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**PtrRecordReturn**](PTRRecordReturn.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

