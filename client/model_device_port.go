/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DevicePort struct {
	Private bool `json:"private"`
	// The unique ID of the port.
	PortId int32 `json:"portId,omitempty"`
	// The vlan tag of the port's native vlan, if applicable.
	NativeVlanTag int32 `json:"nativeVlanTag,omitempty"`
	// ENABLED|DISABLED|UNKOWN
	Status string `json:"status,omitempty"`
	// Indicates if is a bond interface. If not, indicates the Mbps rate of the port.
	Type_ string `json:"type,omitempty"`
	// The unique ID of the native VLAN, if applicable.
	NativeVlanId int32  `json:"nativeVlanId,omitempty"`
	Name         string `json:"name,omitempty"`
	// Your client account's unique ID.
	ClientId int32 `json:"clientId,omitempty"`
	// The unique ID of the port's device.
	DeviceId int32 `json:"deviceId,omitempty"`
	// IPs applied to this port.
	Ips []IpAssignment `json:"ips,omitempty"`
	Mtu int32          `json:"mtu,omitempty"`
}
