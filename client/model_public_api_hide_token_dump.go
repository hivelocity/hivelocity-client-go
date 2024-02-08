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

type PublicApiHideTokenDump struct {
	IpAddresses *PublicApiTokenIp `json:"ipAddresses,omitempty"`
	ApiTokenId  string            `json:"apiTokenId,omitempty"`
	TimeCreated time.Time         `json:"timeCreated,omitempty"`
	Name        string            `json:"name,omitempty"`
	Token       string            `json:"token"`
	TimeUpdated time.Time         `json:"timeUpdated,omitempty"`
	// Token expiration datetime in UTC format.
	TimeExpiration time.Time `json:"timeExpiration"`
}
