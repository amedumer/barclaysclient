# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObCreatePaymentConsent**](OBPISPApi.md#ObCreatePaymentConsent) | **Post** /open-banking/pisp/domestic-payment-consents | Grant consent to a banking entity to initiate a payment.
[**ObGetPayment**](OBPISPApi.md#ObGetPayment) | **Get** /open-banking/pisp/domestic-payments/{domestic_payment_id} | Get payment details .
[**ObGetPaymentConsent**](OBPISPApi.md#ObGetPaymentConsent) | **Get** /open-banking/pisp/domestic-payment-consents/{consent_id} | Get the state of a consent resource.
[**ObMakePayment**](OBPISPApi.md#ObMakePayment) | **Post** /open-banking/pisp/domestic-payments | Make a payment.

# **ObCreatePaymentConsent**
> CreatePaymentConsentResponseBody ObCreatePaymentConsent(ctx, body, xEnvId, xCurrencyId)
Grant consent to a banking entity to initiate a payment.

Simulates the action of a party granting consent for a payment to be made from their account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenBankingPaymentConsentCreationRequestBody**](OpenBankingPaymentConsentCreationRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 

### Return type

[**CreatePaymentConsentResponseBody**](CreatePaymentConsentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObGetPayment**
> GetDomesticPaymentDetailsResponseBody ObGetPayment(ctx, xEnvId, xCurrencyId, xRequestingBankingEntityId, xConsentId, domesticPaymentId)
Get payment details .

Retrieve the latest status of a payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xRequestingBankingEntityId** | **int64**| Identifier for the banking entity (commercial bank/pip) which you are simulating is making this API call. This is the banking entity which needs to have been granted consent from the party. | 
  **xConsentId** | **int64**| Identifier for the consent resource which grants access to perform this request. | 
  **domesticPaymentId** | **int64**|  | 

### Return type

[**GetDomesticPaymentDetailsResponseBody**](GetDomesticPaymentDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObGetPaymentConsent**
> GetPaymentConsentResponseBody ObGetPaymentConsent(ctx, xEnvId, xCurrencyId, xRequestingBankingEntityId, consentId)
Get the state of a consent resource.

Get the state of a consent resource which has been previously created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xRequestingBankingEntityId** | **int64**| Identifier for the banking entity (commercial bank/pip) which you are simulating is making this API call. This is the banking entity which needs to have been granted consent from the party. | 
  **consentId** | **int64**|  | 

### Return type

[**GetPaymentConsentResponseBody**](GetPaymentConsentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObMakePayment**
> MakeDomesticPaymentResponseBody ObMakePayment(ctx, body, xEnvId, xCurrencyId, xRequestingBankingEntityId, xConsentId)
Make a payment.

Make a payment for which consent has previously been given.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeDomesticPaymentRequestBody**](MakeDomesticPaymentRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xRequestingBankingEntityId** | **int64**| Identifier for the banking entity (commercial bank/pip) which you are simulating is making this API call. This is the banking entity which needs to have been granted consent from the party. | 
  **xConsentId** | **int64**| Identifier for the consent resource which grants access to perform this request. | 

### Return type

[**MakeDomesticPaymentResponseBody**](MakeDomesticPaymentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

