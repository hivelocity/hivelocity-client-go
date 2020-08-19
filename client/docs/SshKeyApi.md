# \SshKeyApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSshKeyIdResource**](SshKeyApi.md#DeleteSshKeyIdResource) | **Delete** /ssh_key/{sshKeyId} | Removes public ssh key
[**GetSshKeyResource**](SshKeyApi.md#GetSshKeyResource) | **Get** /ssh_key/ | Gets all public ssh key
[**PostSshKeyResource**](SshKeyApi.md#PostSshKeyResource) | **Post** /ssh_key/ | Adds public ssh key
[**PutSshKeyIdResource**](SshKeyApi.md#PutSshKeyIdResource) | **Put** /ssh_key/{sshKeyId} | Updates public ssh key



## DeleteSshKeyIdResource

> DeleteSshKeyIdResource(ctx, sshKeyId)

Removes public ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32**|  | 

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


## GetSshKeyResource

> GetSshKeyResource(ctx, )

Gets all public ssh key

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


## PostSshKeyResource

> PostSshKeyResource(ctx, payload)

Adds public ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**SshKey**](SshKey.md)|  | 

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


## PutSshKeyIdResource

> PutSshKeyIdResource(ctx, sshKeyId, payload)

Updates public ssh key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32**|  | 
**payload** | [**SshKeyUpdate**](SshKeyUpdate.md)|  | 

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

