# {{classname}}

All URIs are relative to *https://cbdchackathon-dev.barclays.nayaone.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseCbdcAccount**](CBDCLedgerApi.md#CloseCbdcAccount) | **Delete** /cbdc-ledger/accounts/{account_id} | Close a CBDC account.
[**DepositIntoCbdcAccount**](CBDCLedgerApi.md#DepositIntoCbdcAccount) | **Post** /cbdc-ledger/accounts/{account_id}/deposit | Deposit a cash amount into the account.
[**GetCbdcAccount**](CBDCLedgerApi.md#GetCbdcAccount) | **Get** /cbdc-ledger/accounts/{account_id} | Get details about a specific CBDC account.
[**OpenCbdcAccount**](CBDCLedgerApi.md#OpenCbdcAccount) | **Post** /cbdc-ledger/accounts | Open a CBDC account on the central bank&#x27;s core ledger.
[**WithdrawFromCbdcAccount**](CBDCLedgerApi.md#WithdrawFromCbdcAccount) | **Post** /cbdc-ledger/accounts/{account_id}/withdrawal | Withdraw a cash amount from the account.

# **CloseCbdcAccount**
> GetCbdcAccountResponseBody CloseCbdcAccount(ctx, xEnvId, xCurrencyId, xPipId, accountId)
Close a CBDC account.

Close a CBDC account. The account must have a 0 balance for this operation to succeed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xPipId** | **int64**| Identifier for the pip which you are simulating is making this API call. | 
  **accountId** | **int64**|  | 

### Return type

[**GetCbdcAccountResponseBody**](GetCBDCAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DepositIntoCbdcAccount**
> GetCbdcAccountResponseBody DepositIntoCbdcAccount(ctx, body, xEnvId, xCurrencyId, xPipId, accountId)
Deposit a cash amount into the account.

Simulation of depositing cash into the account.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeDepositRequestBody**](MakeDepositRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xPipId** | **int64**| Identifier for the pip which you are simulating is making this API call. | 
  **accountId** | **int64**|  | 

### Return type

[**GetCbdcAccountResponseBody**](GetCBDCAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCbdcAccount**
> GetCbdcAccountResponseBody GetCbdcAccount(ctx, xEnvId, xCurrencyId, xPipId, accountId)
Get details about a specific CBDC account.

Retrieve details about a specific CBDC account including its balance.  As these details are for an account on the core ledger there is no information about the party which owns the account. This is because the core ledger, by design, does not store any information which can directly identify the party.  This API endpoint is intended to be used by PIPs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xPipId** | **int64**| Identifier for the pip which you are simulating is making this API call. | 
  **accountId** | **int64**|  | 

### Return type

[**GetCbdcAccountResponseBody**](GetCBDCAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenCbdcAccount**
> OpenAccountResponseBody OpenCbdcAccount(ctx, xEnvId, xCurrencyId, xPipId)
Open a CBDC account on the central bank's core ledger.

Open a CBDC account on the core ledger of the central bank.  This API endpoint is intended to be used by PIPs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xPipId** | **int64**| Identifier for the pip which you are simulating is making this API call. | 

### Return type

[**OpenAccountResponseBody**](OpenAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawFromCbdcAccount**
> GetCbdcAccountResponseBody WithdrawFromCbdcAccount(ctx, body, xEnvId, xCurrencyId, xPipId, accountId)
Withdraw a cash amount from the account.

Simulation of withdrawing money from the account.  <h2> Amount details </h2>  The amount must be given in the units of the currency (for GBP that means pence, for  USD/EUR: cents, etc.).  For example if you wanted £1/$1/€1 to be the amount you would give the value 100 (as the value refers to the amount of pence/cents).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MakeWithdrawalRequestBody**](MakeWithdrawalRequestBody.md)|  | 
  **xEnvId** | **int64**| Identifier for the environment. | 
  **xCurrencyId** | **int64**| Identifier for the currency. | 
  **xPipId** | **int64**| Identifier for the pip which you are simulating is making this API call. | 
  **accountId** | **int64**|  | 

### Return type

[**GetCbdcAccountResponseBody**](GetCBDCAccountResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

