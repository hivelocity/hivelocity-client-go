/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProductOption struct {
	// Unique product id.
	ProductId int32 `json:"productId,omitempty"`
	// List of available options.
	Options interface{} `json:"options,omitempty"`
	// *DEPRECATED*
	Pricing interface{} `json:"pricing,omitempty"`
}
