/*
 * Barclays 2022 CBDC Hackathon API
 *
 * API for the Barclays 2022 CBDC Hackathon.  Please ensure you include your API key in the `x-api-key` header for your requests. 
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CbdcAccountChild struct {
	Id int64 `json:"id,omitempty"`
	AccountOwningPIPId int64 `json:"accountOwningPIPId,omitempty"`
	Status string `json:"status,omitempty"`
}
