# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObCreateAccountAccessConsent**](OBAISPApi.md#ObCreateAccountAccessConsent) | **Post** /open-banking/aisp/account-access-consents | Grant consent to a banking entity to access account data.
[**ObGetAccountAccessConsent**](OBAISPApi.md#ObGetAccountAccessConsent) | **Get** /open-banking/aisp/account-access-consents/{consent_id} | Get the state of a consent.
[**ObGetAccountDetails**](OBAISPApi.md#ObGetAccountDetails) | **Get** /open-banking/aisp/accounts/{account_id} | Get account details.
[**ObGetPartyDetails**](OBAISPApi.md#ObGetPartyDetails) | **Get** /open-banking/aisp/accounts/{account_id}/party | Get party details.

# **ObCreateAccountAccessConsent**
> CreateAccountAccessConsentResponseBody ObCreateAccountAccessConsent(ctx, body, xEnvId, xCurrencyId)
Grant consent to a banking entity to access account data.

Simulates the action of a party granting consent to a banking entity to access its account data.  See the API documentation for a full description of the different actors involved with the open banking APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenBankingAccountAccessConsentCreationRequestBody**](OpenBankingAccountAccessConsentCreationRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 

### Return type

[**CreateAccountAccessConsentResponseBody**](CreateAccountAccessConsentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObGetAccountAccessConsent**
> GetAccountAccessConsentResponseBody ObGetAccountAccessConsent(ctx, xEnvId, xCurrencyId, xRequestingBankingEntityId, consentId)
Get the state of a consent.

Retrieve the latest state of a specific consent resource.  See the API documentation for a full description of the different actors involved with the open banking APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xRequestingBankingEntityId** | **int64**| Identifier for the banking entity (commercial bank/pip) which you are simulating is making this API call. This is the banking entity which needs to have been granted consent from the party. | 
  **consentId** | **int64**|  | 

### Return type

[**GetAccountAccessConsentResponseBody**](GetAccountAccessConsentResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObGetAccountDetails**
> GetBankingEntityAccountResponseBody ObGetAccountDetails(ctx, xEnvId, xCurrencyId, xRequestingBankingEntityId, xConsentId, accountId)
Get account details.

Retrieve details for the account for which consent was granted.  See the API documentation for a full description of the different actors involved with the open banking APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xRequestingBankingEntityId** | **int64**| Identifier for the banking entity (commercial bank/pip) which you are simulating is making this API call. This is the banking entity which needs to have been granted consent from the party. | 
  **xConsentId** | **int64**| Identifier for the consent resource which grants access to perform this request. | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObGetPartyDetails**
> GetPartyResponseBody ObGetPartyDetails(ctx, xEnvId, xCurrencyId, xRequestingBankingEntityId, xConsentId, accountId)
Get party details.

Retrieve details about the party which owns the account and has granted consent to access their data.  See the API documentation for a full description of the different actors involved with the open banking APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xRequestingBankingEntityId** | **int64**| Identifier for the banking entity (commercial bank/pip) which you are simulating is making this API call. This is the banking entity which needs to have been granted consent from the party. | 
  **xConsentId** | **int64**| Identifier for the consent resource which grants access to perform this request. | 
  **accountId** | **int64**|  | 

### Return type

[**GetPartyResponseBody**](GetPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

