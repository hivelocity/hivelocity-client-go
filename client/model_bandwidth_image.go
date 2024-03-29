/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BandwidthImage struct {
	// A PNG image of bandwidth usage.
	GraphImage string `json:"graphImage,omitempty"`
	// The interface(s) displayed in the image.
	Interface_ string `json:"interface,omitempty"`
	// The unique ID of the switch where bandwidth data is measured.
	SwitchId string `json:"switchId,omitempty"`
}
