# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewEnv**](SandboxEnvironmentApi.md#CreateNewEnv) | **Post** /envs | Create new environment.
[**GetEnv**](SandboxEnvironmentApi.md#GetEnv) | **Get** /envs/{env_id} | Get a description of a specific environment.
[**GetEnvs**](SandboxEnvironmentApi.md#GetEnvs) | **Get** /envs | Get a page listing the environments you have created.
[**TerminateEnv**](SandboxEnvironmentApi.md#TerminateEnv) | **Delete** /envs/{env_id} | Terminate a environment.

# **CreateNewEnv**
> CreateNewEnvironmentResponseBody CreateNewEnv(ctx, )
Create new environment.

Create a fresh, new environment which you can use to experiment with.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CreateNewEnvironmentResponseBody**](CreateNewEnvironmentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnv**
> GetEnvironmentDetailsResponseBody GetEnv(ctx, envId)
Get a description of a specific environment.

Returns a description of a specific environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **int64**|  | 

### Return type

[**GetEnvironmentDetailsResponseBody**](GetEnvironmentDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvs**
> GetEnvironmentDetailsPageResponseBody GetEnvs(ctx, optional)
Get a page listing the environments you have created.

Get a page listing the environments you have created.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SandboxEnvironmentApiGetEnvsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SandboxEnvironmentApiGetEnvsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) of the page you want to return. | [default to 0]
 **pageSize** | **optional.Int32**| The max. number of items (size) a page will contain. | [default to 20]

### Return type

[**GetEnvironmentDetailsPageResponseBody**](GetEnvironmentDetailsPageResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateEnv**
> GetEnvironmentDetailsResponseBody TerminateEnv(ctx, envId)
Terminate a environment.

Terminates a environment so it will no longer be available for experimentation. All child resources of the environment (ledgers/accounts/etc.) must be terminated prior to issuing this request otherwise the request will fail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **int64**|  | 

### Return type

[**GetEnvironmentDetailsResponseBody**](GetEnvironmentDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

