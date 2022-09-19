# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthPing**](StartHereApi.md#AuthPing) | **Get** /start-here/auth-ping | Send a HTTP ping request and expect a pong back if everything is OK. Requires you to be authenticated for the request to succeed.
[**PublicPing**](StartHereApi.md#PublicPing) | **Get** /start-here/public-ping | Send a HTTP ping request and expect a pong back if everything is OK.

# **AuthPing**
> PingResponse AuthPing(ctx, )
Send a HTTP ping request and expect a pong back if everything is OK. Requires you to be authenticated for the request to succeed.

Useful for testing whether you can reach the server or not. This request does not require you to provide your team's API key.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicPing**
> PingResponse PublicPing(ctx, )
Send a HTTP ping request and expect a pong back if everything is OK.

Useful for testing whether you can reach the server or not. This request does not require you to provide the API key.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

