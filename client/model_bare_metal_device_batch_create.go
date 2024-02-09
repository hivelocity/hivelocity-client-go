/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BareMetalDeviceBatchCreate struct {
	// Unique ID to identify the group order.
	OrderGroupId int32 `json:"orderGroupId,omitempty"`
	// List of devices to provision.
	Devices []BareMetalDeviceCreate `json:"devices,omitempty"`
}
