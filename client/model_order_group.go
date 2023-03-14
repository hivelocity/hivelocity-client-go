/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrderGroup struct {
	Name      string  `json:"name,omitempty"`
	Id        int32   `json:"id,omitempty"`
	SameRack  bool    `json:"same_rack,omitempty"`
	DeviceIds []int32 `json:"device_ids,omitempty"`
}
