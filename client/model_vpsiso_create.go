/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VpsisoCreate struct {
	// Determines whether the VPS instance ISO is bootable or not.
	Bootable bool `json:"bootable"`
	// The name of the VPS instance ISO.
	Name string `json:"name"`
	// The unique client account ID.
	ClientId int32 `json:"clientId,omitempty"`
	// The URL to be directed to when downloading the VPS instance ISO.
	Url string `json:"url"`
	// The description of the VPS instance ISO.
	Description string `json:"description"`
	// The facility code associated wth the VPS instance.  For example:TPA1.
	FacilityCode string `json:"facilityCode"`
}
