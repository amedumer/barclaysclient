# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCurrency**](CurrencyApi.md#CreateCurrency) | **Post** /currencies | Create a new currency within an environment.
[**GetCurrenciesPage**](CurrencyApi.md#GetCurrenciesPage) | **Get** /currencies | Get a page listing the currencies you have created.
[**GetCurrency**](CurrencyApi.md#GetCurrency) | **Get** /currencies/{currency_id} | Get details about a specific currency.
[**TerminateCurrency**](CurrencyApi.md#TerminateCurrency) | **Delete** /currencies/{currency_id} | Terminate a currency (and by extension it&#x27;s CBDC Ledger and other child resources).

# **CreateCurrency**
> CreateCurrencyResponseBody CreateCurrency(ctx, body, xEnvId)
Create a new currency within an environment.

Create a new currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCurrencyRequestBody**](CreateCurrencyRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 

### Return type

[**CreateCurrencyResponseBody**](CreateCurrencyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrenciesPage**
> GetCurrencyDetailsPageResponseBody GetCurrenciesPage(ctx, xEnvId, optional)
Get a page listing the currencies you have created.

Get a page listing the currencies you have created.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
 **optional** | ***CurrencyApiGetCurrenciesPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrencyApiGetCurrenciesPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) of the page you want to return. | [default to 0]
 **pageSize** | **optional.Int32**| The max. number of items (size) a page will contain. | [default to 20]

### Return type

[**GetCurrencyDetailsPageResponseBody**](GetCurrencyDetailsPageResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrency**
> GetCurrencyDetailsResponseBody GetCurrency(ctx, xEnvId, currencyId)
Get details about a specific currency.

Retrieve details about a specific currency including the CBDC accounts on the currency's core ledger, as well as the commercial banks and pips registered to operate within that currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **currencyId** | **int64**|  | 

### Return type

[**GetCurrencyDetailsResponseBody**](GetCurrencyDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateCurrency**
> GetCurrencyDetailsResponseBody TerminateCurrency(ctx, xEnvId, currencyId)
Terminate a currency (and by extension it's CBDC Ledger and other child resources).

Terminates a currency. This will mean the CBDC ledger for that currency will no longer be available for experimentation. All child resources of the currency (the CBDC ledger accounts, the PIPs, the Commercial Banks) must be terminated prior to issuing this request, otherwise the request will fail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **currencyId** | **int64**|  | 

### Return type

[**GetCurrencyDetailsResponseBody**](GetCurrencyDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

