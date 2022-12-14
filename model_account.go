/*
 * Barclays 2022 CBDC Hackathon API
 *
 * API for the Barclays 2022 CBDC Hackathon.  Please ensure you include your API key in the `x-api-key` header for your requests. 
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Account struct {
	AccountId int64 `json:"accountId,omitempty"`
	Status string `json:"status,omitempty"`
}
