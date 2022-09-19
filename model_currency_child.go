/*
 * Barclays 2022 CBDC Hackathon API
 *
 * API for the Barclays 2022 CBDC Hackathon.  Please ensure you include your API key in the `x-api-key` header for your requests. 
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CurrencyChild struct {
	CurrencyId int64 `json:"currencyId,omitempty"`
	Status string `json:"status,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
}
