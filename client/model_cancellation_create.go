/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CancellationCreate struct {
	// Unique ID of a service. Must be related to `device_id`.
	ServiceId int32 `json:"serviceId"`
	// Unique ID of a device. Must be related to `service_id`.
	DeviceId int32 `json:"deviceId"`
	// Details why the device is being cancelled
	Comments string `json:"comments,omitempty"`
	// My current service is being upgraded|Hardware/software/network or power issues|I think I have found a better deal|I have experienced tech support problems|I am consolidating my Hivelocity accounts|I am a reseller and my customer cancelled|I am moving to a different technology solution: Public Cloud|I am moving to a different technology solution: Managed Hosting Company|A product difference (Example: AWS)
	Reason string `json:"reason"`
}
