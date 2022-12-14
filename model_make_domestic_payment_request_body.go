/*
 * Barclays 2022 CBDC Hackathon API
 *
 * API for the Barclays 2022 CBDC Hackathon.  Please ensure you include your API key in the `x-api-key` header for your requests. 
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MakeDomesticPaymentRequestBody struct {
	SourceAccountId int64 `json:"sourceAccountId,omitempty"`
	DestinationAccountId int64 `json:"destinationAccountId,omitempty"`
	PaymentAmountInCurrencyUnits int64 `json:"paymentAmountInCurrencyUnits,omitempty"`
}
