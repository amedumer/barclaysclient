# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseCommercialBankAccount**](CommercialBanksApi.md#CloseCommercialBankAccount) | **Delete** /commercial-banks/{bank_id}/accounts/{account_id} | Close the commercial bank account.
[**CreateNewCommercialBank**](CommercialBanksApi.md#CreateNewCommercialBank) | **Post** /commercial-banks | Create a new commercial bank entity.
[**DepositIntoCommercialBankAccount**](CommercialBanksApi.md#DepositIntoCommercialBankAccount) | **Post** /commercial-banks/{bank_id}/accounts/{account_id}/deposit | Deposit a cash amount into the account.
[**DeregisterCommercialBankParty**](CommercialBanksApi.md#DeregisterCommercialBankParty) | **Delete** /commercial-banks/{bank_id}/parties/{party_id} | De-register a specific party from a commercial bank.
[**GetCommercialBankAccount**](CommercialBanksApi.md#GetCommercialBankAccount) | **Get** /commercial-banks/{bank_id}/accounts/{account_id} | Get details of a specific account at the commercial bank.
[**GetCommercialBankAccountsPage**](CommercialBanksApi.md#GetCommercialBankAccountsPage) | **Get** /commercial-banks/{bank_id}/accounts | Get details of all of the accounts at the specific commercial bank.
[**GetCommercialBankDetails**](CommercialBanksApi.md#GetCommercialBankDetails) | **Get** /commercial-banks/{bank_id} | Get commercial bank details.
[**GetCommercialBankDetailsPage**](CommercialBanksApi.md#GetCommercialBankDetailsPage) | **Get** /commercial-banks | Get a page listing the commercial banks you have created.
[**GetCommercialBankPartyDetails**](CommercialBanksApi.md#GetCommercialBankPartyDetails) | **Get** /commercial-banks/{bank_id}/parties/{party_id} | Get details about a specific party registered at a commercial bank.
[**GetCommercialBankPartyPageDetails**](CommercialBanksApi.md#GetCommercialBankPartyPageDetails) | **Get** /commercial-banks/{bank_id}/parties | Get details of all parties registered at the commercial bank.
[**OpenCommercialBankAccount**](CommercialBanksApi.md#OpenCommercialBankAccount) | **Post** /commercial-banks/{bank_id}/accounts | Open an account at the commercial bank.
[**RegisterCommercialBankParty**](CommercialBanksApi.md#RegisterCommercialBankParty) | **Post** /commercial-banks/{bank_id}/parties | Register a party with the commercial bank.
[**TerminateCommercialBank**](CommercialBanksApi.md#TerminateCommercialBank) | **Delete** /commercial-banks/{bank_id} | Terminate a commercial bank.
[**WithdrawFromCommercialBankAccount**](CommercialBanksApi.md#WithdrawFromCommercialBankAccount) | **Post** /commercial-banks/{bank_id}/accounts/{account_id}/withdrawal | Withdraw a cash amount from the account.

# **CloseCommercialBankAccount**
> GetBankingEntityAccountResponseBody CloseCommercialBankAccount(ctx, xEnvId, xCurrencyId, bankId, accountId)
Close the commercial bank account.

Close the commercial bank account. Note: the account must have a 0 balance for this API call to succeed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNewCommercialBank**
> CreateCommercialBankResponseBody CreateNewCommercialBank(ctx, body, xEnvId, xCurrencyId)
Create a new commercial bank entity.

Create a new commercial bank which will manage its own ledger of accounts in the given currency. Once created, you can register parties with the commercial bank and also open up accounts owned by a registered party.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCommercialBankRequestBody**](CreateCommercialBankRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 

### Return type

[**CreateCommercialBankResponseBody**](CreateCommercialBankResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DepositIntoCommercialBankAccount**
> GetBankingEntityAccountResponseBody DepositIntoCommercialBankAccount(ctx, body, xEnvId, xCurrencyId, bankId, accountId)
Deposit a cash amount into the account.

Simulation of depositing cash into the account.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeDepositRequestBody**](MakeDepositRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeregisterCommercialBankParty**
> GetPartyResponseBody DeregisterCommercialBankParty(ctx, xEnvId, xCurrencyId, bankId, partyId)
De-register a specific party from a commercial bank.

De-register a specific party from a commercial bank.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
  **partyId** | **int64**|  | 

### Return type

[**GetPartyResponseBody**](GetPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommercialBankAccount**
> GetBankingEntityAccountResponseBody GetCommercialBankAccount(ctx, xEnvId, xCurrencyId, bankId, accountId)
Get details of a specific account at the commercial bank.

Get details of a specific account at the commercial bank.  <h2> Balance details </h2>  The `balance` is given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For your convenience there is also a field, `formattedBalance` which displays the balance as a decimal number.  For example if you had £1.03 in the account you would receive a `balance` of 103 and a `formattedBalance` of £1.03.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommercialBankAccountsPage**
> GetBankingEntityAccountsPageResponseBody GetCommercialBankAccountsPage(ctx, xEnvId, xCurrencyId, bankId, optional)
Get details of all of the accounts at the specific commercial bank.

Get details of all of the accounts at the specific commercial bank.    <h2> Balance details </h2>  The `balance` is given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For your convenience there is also a field, `formattedBalance` which displays the balance as a decimal number.  For example if you had £1.03 in the account you would receive a `balance` of 103 and a `formattedBalance` of £1.03.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
 **optional** | ***CommercialBanksApiGetCommercialBankAccountsPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommercialBanksApiGetCommercialBankAccountsPageOpts struct
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

# **GetCommercialBankDetails**
> GetCommercialBankDetailsResponseBody GetCommercialBankDetails(ctx, xEnvId, xCurrencyId, bankId)
Get commercial bank details.

Retrieve the details of a specific commercial bank.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 

### Return type

[**GetCommercialBankDetailsResponseBody**](GetCommercialBankDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommercialBankDetailsPage**
> GetCommercialBankDetailsPageResponseBody GetCommercialBankDetailsPage(ctx, xEnvId, xCurrencyId, optional)
Get a page listing the commercial banks you have created.

Get a page listing the commercial banks you have created.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
 **optional** | ***CommercialBanksApiGetCommercialBankDetailsPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommercialBanksApiGetCommercialBankDetailsPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) of the page you want to return. | [default to 0]
 **pageSize** | **optional.Int32**| The max. number of items (size) a page will contain. | [default to 20]

### Return type

[**GetCommercialBankDetailsPageResponseBody**](GetCommercialBankDetailsPageResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommercialBankPartyDetails**
> GetPartyResponseBody GetCommercialBankPartyDetails(ctx, xEnvId, xCurrencyId, bankId, partyId)
Get details about a specific party registered at a commercial bank.

Returns the details about the specific party registered at the commercial bank.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
  **partyId** | **int64**|  | 

### Return type

[**GetPartyResponseBody**](GetPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommercialBankPartyPageDetails**
> GetPartyViewsPageResponseBody GetCommercialBankPartyPageDetails(ctx, xEnvId, xCurrencyId, bankId, optional)
Get details of all parties registered at the commercial bank.

Get details of all parties registered at the commercial bank.  <h2> Paging </h2>  This endpoint can potentially return quite a large result and so you must use paging to control the size of the response.  Pagination is handled by using the `pageIndex` and `pageSize` parameters.  - The `pageIndex` refers to the index of the page of data you want returned. It is 0-indexed meaning the first page will always have the `pageIndex` value of 0.  - The `pageSize` refers to the maximum number of entries you would like to see within a single page (hence page size).    - You cannot request a `pageSize` greater than  1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
 **optional** | ***CommercialBanksApiGetCommercialBankPartyPageDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommercialBanksApiGetCommercialBankPartyPageDetailsOpts struct
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

# **OpenCommercialBankAccount**
> OpenAccountResponseBody OpenCommercialBankAccount(ctx, body, xEnvId, xCurrencyId, bankId)
Open an account at the commercial bank.

Open an account at the commercial bank. You need to provide a reference to the ID of the party which is opening the account in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenAccountRequestBody**](OpenAccountRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 

### Return type

[**OpenAccountResponseBody**](OpenAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterCommercialBankParty**
> RegisterPartyResponseBody RegisterCommercialBankParty(ctx, body, xEnvId, xCurrencyId, bankId)
Register a party with the commercial bank.

Register a party with the commercial bank.  Once registered you can then open accounts belonging to that party. In the real world the registration would also include performing KYC and AML checks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterPartyRequestBody**](RegisterPartyRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 

### Return type

[**RegisterPartyResponseBody**](RegisterPartyResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateCommercialBank**
> GetCommercialBankDetailsResponseBody TerminateCommercialBank(ctx, xEnvId, xCurrencyId, bankId)
Terminate a commercial bank.

Terminate a commercial bank entity. Note: all child resources (parties/accounts) must be terminated as well otherwise this API call will fail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 

### Return type

[**GetCommercialBankDetailsResponseBody**](GetCommercialBankDetailsResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawFromCommercialBankAccount**
> GetBankingEntityAccountResponseBody WithdrawFromCommercialBankAccount(ctx, body, xEnvId, xCurrencyId, bankId, accountId)
Withdraw a cash amount from the account.

Simulation of withdrawing money from the account.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeWithdrawalRequestBody**](MakeWithdrawalRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **bankId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**GetBankingEntityAccountResponseBody**](GetBankingEntityAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

