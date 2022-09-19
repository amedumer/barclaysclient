# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomesticPaymentEcosystemServiceGetPayment**](EcosystemServiceDomesticPaymentProcessorApi.md#DomesticPaymentEcosystemServiceGetPayment) | **Get** /domestic-payments/{domestic_payment_id} | Get details about a payment.
[**DomesticPaymentEcosystemServiceMakePayment**](EcosystemServiceDomesticPaymentProcessorApi.md#DomesticPaymentEcosystemServiceMakePayment) | **Post** /domestic-payments | Make a payment between two accounts.

# **DomesticPaymentEcosystemServiceGetPayment**
> GetDomesticPaymentDetailsResponseBody DomesticPaymentEcosystemServiceGetPayment(ctx, xEnvId, xCurrencyId, domesticPaymentId)
Get details about a payment.

Retrieve the latest state of a payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **domesticPaymentId** | **int64**|  | 

### Return type

[**GetDomesticPaymentDetailsResponseBody**](GetDomesticPaymentDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomesticPaymentEcosystemServiceMakePayment**
> MakeDomesticPaymentResponseBody DomesticPaymentEcosystemServiceMakePayment(ctx, body, xEnvId, xCurrencyId)
Make a payment between two accounts.

Make a payment between two accounts.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeDomesticPaymentRequestBody**](MakeDomesticPaymentRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 

### Return type

[**MakeDomesticPaymentResponseBody**](MakeDomesticPaymentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

