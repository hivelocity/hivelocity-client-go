/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// PublicApiTokenLoad struct for PublicApiTokenLoad
type PublicApiTokenLoad struct {
	Name        string           `json:"name,omitempty"`
	IpAddresses PublicApiTokenIp `json:"ipAddresses,omitempty"`
}
