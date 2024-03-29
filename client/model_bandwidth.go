/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Bandwidth struct {
	// Returns metadata about the query including the 95th percentile values of the returned data, legend for returned data arrays, and query args passed to the query.
	Metadata interface{} `json:"metadata,omitempty"`
	// Returns an array where each item contains a dictionary of data and an identifier for each requested interface. The data itself is an array of arrays where each item is the measurementat the requested `step` and the index of each integer value corresponds with the legend returned in the `metadata`.
	BandwidthData [][]float32 `json:"bandwidthData,omitempty"`
	// The interface(s) displayed in the image.
	Interface_ string `json:"interface,omitempty"`
	// The unique ID of the switch where bandwidth data is measured.
	SwitchId string `json:"switchId,omitempty"`
}
