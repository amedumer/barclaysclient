# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClosePipAccount**](PaymentInterfaceProvidersPIPsApi.md#ClosePipAccount) | **Delete** /pips/{pip_id}/accounts/{account_id} | Close the account at the PIP.
[**CreateNewPip**](PaymentInterfaceProvidersPIPsApi.md#CreateNewPip) | **Post** /pips | Create a new PIP.
[**CreateNewPipParty**](PaymentInterfaceProvidersPIPsApi.md#CreateNewPipParty) | **Post** /pips/{pip_id}/parties | Register a party with the PIP.
[**DepositIntoPipAccount**](PaymentInterfaceProvidersPIPsApi.md#DepositIntoPipAccount) | **Post** /pips/{pip_id}/accounts/{account_id}/deposit | Deposit a cash amount into the account.
[**DeregisterPipParty**](PaymentInterfaceProvidersPIPsApi.md#DeregisterPipParty) | **Delete** /pips/{pip_id}/parties/{party_id} | De-register a specific party from a PIP.
[**GetPipAccount**](PaymentInterfaceProvidersPIPsApi.md#GetPipAccount) | **Get** /pips/{pip_id}/accounts/{account_id} | Get details of a specific account at the PIP.
[**GetPipAccountPage**](PaymentInterfaceProvidersPIPsApi.md#GetPipAccountPage) | **Get** /pips/{pip_id}/accounts | Get details of all of the accounts at the PIP.
[**GetPipDetails**](PaymentInterfaceProvidersPIPsApi.md#GetPipDetails) | **Get** /pips/{pip_id} | Get PIP details.
[**GetPipDetailsPage**](PaymentInterfaceProvidersPIPsApi.md#GetPipDetailsPage) | **Get** /pips | Get a page listing the PIPs you have created.
[**GetPipParty**](PaymentInterfaceProvidersPIPsApi.md#GetPipParty) | **Get** /pips/{pip_id}/parties/{party_id} | Get details about a specific party registered at a PIP.
[**GetPipPartyPageDetails**](PaymentInterfaceProvidersPIPsApi.md#GetPipPartyPageDetails) | **Get** /pips/{pip_id}/parties | Get details of all parties registered at the PIP.
[**OpenPipAccount**](PaymentInterfaceProvidersPIPsApi.md#OpenPipAccount) | **Post** /pips/{pip_id}/accounts | Open a CBDC account at the PIP.
[**TerminatePip**](PaymentInterfaceProvidersPIPsApi.md#TerminatePip) | **Delete** /pips/{pip_id} | Terminate a PIP.
[**WithdrawFromPipAccount**](PaymentInterfaceProvidersPIPsApi.md#WithdrawFromPipAccount) | **Post** /pips/{pip_id}/accounts/{account_id}/withdrawal | Withdraw a cash amount from the account.

# **ClosePipAccount**
> GetBankingEntityAccountResponseBody ClosePipAccount(ctx, xEnvId, xCurrencyId, pipId, accountId)
Close the account at the PIP.

Close the account. Note: the account must have a 0 balance for this API call to succeed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNewPip**
> CreatePaymentInterfaceProviderResponseBody CreateNewPip(ctx, body, xEnvId, xCurrencyId)
Create a new PIP.

Create a new PIP which will have access to the CBDC ledger for the given currency. Once created, you can register parties with the PIP and also open up accounts. An account at a PIP is a proxy for an account on the CBDC ledger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePaymentInterfaceProviderRequestBody**](CreatePaymentInterfaceProviderRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 

### Return type

[**CreatePaymentInterfaceProviderResponseBody**](CreatePaymentInterfaceProviderResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNewPipParty**
> RegisterPartyResponseBody CreateNewPipParty(ctx, body, xEnvId, xCurrencyId, pipId)
Register a party with the PIP.

Register a party with the PIP.  It is the PIP which holds the personally identifiable data of the party and which interacts with the CBDC core ledger on behalf of the party.   Once registered you can then open accounts belonging to that party. In the real world the registration would also include performing KYC and AML checks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterPartyRequestBody**](RegisterPartyRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 

### Return type

[**RegisterPartyResponseBody**](RegisterPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DepositIntoPipAccount**
> GetBankingEntityAccountResponseBody DepositIntoPipAccount(ctx, body, xEnvId, xCurrencyId, pipId, accountId)
Deposit a cash amount into the account.

Simulation of depositing cash into the account.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeDepositRequestBody**](MakeDepositRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeregisterPipParty**
> GetPartyResponseBody DeregisterPipParty(ctx, xEnvId, xCurrencyId, pipId, partyId)
De-register a specific party from a PIP.

De-register a specific party from a PIP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
  **partyId** | **int64**|  | 

### Return type

[**GetPartyResponseBody**](GetPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipAccount**
> GetBankingEntityAccountResponseBody GetPipAccount(ctx, xEnvId, xCurrencyId, pipId, accountId)
Get details of a specific account at the PIP.

Get details of a specific account at the PIP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipAccountPage**
> GetBankingEntityAccountsPageResponseBody GetPipAccountPage(ctx, xEnvId, xCurrencyId, pipId, optional)
Get details of all of the accounts at the PIP.

Get details of all of the accounts at the PIP.    <h2> Balance details </h2>  The `balance` is given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For your convenience there is also a field, `formattedBalance` which displays the balance as a decimal number.  For example if you had £1.03 in the account you would receive a `balance` of 103 and a `formattedBalance` of £1.03.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
 **optional** | ***PaymentInterfaceProvidersPIPsApiGetPipAccountPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentInterfaceProvidersPIPsApiGetPipAccountPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) of the page you want to return. | [default to 0]
 **pageSize** | **optional.Int32**| The max. number of items (size) a page will contain. | [default to 20]

### Return type

[**GetBankingEntityAccountsPageResponseBody**](GetBankingEntityAccountsPageResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipDetails**
> GetPipDetailsResponseBody GetPipDetails(ctx, xEnvId, xCurrencyId, pipId)
Get PIP details.

Retrieve the details of a specific PIP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 

### Return type

[**GetPipDetailsResponseBody**](GetPIPDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipDetailsPage**
> GetPipDetailsPageResponseBody GetPipDetailsPage(ctx, xEnvId, xCurrencyId, optional)
Get a page listing the PIPs you have created.

Get a page listing the PIPs you have created.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
 **optional** | ***PaymentInterfaceProvidersPIPsApiGetPipDetailsPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentInterfaceProvidersPIPsApiGetPipDetailsPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) of the page you want to return. | [default to 0]
 **pageSize** | **optional.Int32**| The max. number of items (size) a page will contain. | [default to 20]

### Return type

[**GetPipDetailsPageResponseBody**](GetPIPDetailsPageResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipParty**
> GetPartyResponseBody GetPipParty(ctx, xEnvId, xCurrencyId, pipId, partyId)
Get details about a specific party registered at a PIP.

Returns the details about the specific party registered at the PIP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
  **partyId** | **int64**|  | 

### Return type

[**GetPartyResponseBody**](GetPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipPartyPageDetails**
> GetPartyViewsPageResponseBody GetPipPartyPageDetails(ctx, xEnvId, xCurrencyId, pipId, optional)
Get details of all parties registered at the PIP.

Get details of all parties registered at the PIP.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
 **optional** | ***PaymentInterfaceProvidersPIPsApiGetPipPartyPageDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentInterfaceProvidersPIPsApiGetPipPartyPageDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) of the page you want to return. | [default to 0]
 **pageSize** | **optional.Int32**| The max. number of items (size) a page will contain. | [default to 20]

### Return type

[**GetPartyViewsPageResponseBody**](GetPartyViewsPageResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenPipAccount**
> OpenAccountResponseBody OpenPipAccount(ctx, body, xEnvId, xCurrencyId, pipId)
Open a CBDC account at the PIP.

Opens a CBDC account in the given currency at the payment interface provider. This account is essentially a proxy for the CBDC account stored on the central bank's core ledger. You do not need to interact directly with the core ledger - the PIP will do that behind the scenes. You need to provide a reference to the ID of the party which is opening the account in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenAccountRequestBody**](OpenAccountRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 

### Return type

[**OpenAccountResponseBody**](OpenAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminatePip**
> GetPipDetailsResponseBody TerminatePip(ctx, xEnvId, xCurrencyId, pipId)
Terminate a PIP.

Terminate a PIP entity. Note: all child resources (parties/accounts) must be terminated as well otherwise this API call will fail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 

### Return type

[**GetPipDetailsResponseBody**](GetPIPDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawFromPipAccount**
> GetBankingEntityAccountResponseBody WithdrawFromPipAccount(ctx, body, xEnvId, xCurrencyId, pipId, accountId)
Withdraw a cash amount from the account.

Simulation of withdrawing money from the account.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeWithdrawalRequestBody**](MakeWithdrawalRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **pipId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

