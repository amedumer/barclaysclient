/*
 * Barclays 2022 CBDC Hackathon API
 *
 * API for the Barclays 2022 CBDC Hackathon.  Please ensure you include your API key in the `x-api-key` header for your requests. 
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ErrorResponseBody struct {
	StatusCode int32 `json:"statusCode,omitempty"`
	StatusDescription string `json:"statusDescription,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Message string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
