/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// DeviceInitialCreds struct for DeviceInitialCreds
type DeviceInitialCreds struct {
	Password  string `json:"password,omitempty"`
	User      string `json:"user,omitempty"`
	Port      int32  `json:"port,omitempty"`
	LockerUrl string `json:"lockerUrl,omitempty"`
}
