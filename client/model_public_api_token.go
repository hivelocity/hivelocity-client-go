/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type PublicApiToken struct {
	Name string `json:"name,omitempty"`
	// Token expiration datetime in Unix time format.
	TimeExpiration time.Time         `json:"timeExpiration,omitempty"`
	IpAddresses    *PublicApiTokenIp `json:"ipAddresses,omitempty"`
}
