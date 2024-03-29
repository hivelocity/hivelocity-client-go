/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Vpsiso struct {
	// Determines whether the VPS instance ISO is bootable or not.
	Bootable bool `json:"bootable,omitempty"`
	// The status of the VPS instance ISO.
	Status string `json:"status,omitempty"`
	// These are the device IDs with the ISO attached.
	DeviceIds []int32 `json:"deviceIds,omitempty"`
	// The size of the VPS instance ISO.
	Size int32 `json:"size,omitempty"`
	// The name of the VPS instance ISO.
	Name string `json:"name,omitempty"`
	// The ready state of the VPS instance ISO.
	Ready bool `json:"ready,omitempty"`
	// The unique ID of the VPS instance ISO.
	IsoId string `json:"isoId,omitempty"`
	// The URL to be directed to when downloading the VPS instance ISO.
	Url string `json:"url,omitempty"`
	// The description of the VPS instance ISO.
	Description string `json:"description,omitempty"`
	// The facility code associated with the ISO instance. For example:TPA1.
	FacilityCode string `json:"facilityCode,omitempty"`
}
