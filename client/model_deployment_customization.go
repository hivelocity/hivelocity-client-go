/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// DeploymentCustomization struct for DeploymentCustomization
type DeploymentCustomization struct {
	ProductId       int32  `json:"productId"`
	OperatingSystem string `json:"operatingSystem"`
	// must be one of ['monthly', 'quarterly', 'semi-annually', 'annually', 'biennial', 'triennial']
	BillingPeriod   string   `json:"billingPeriod,omitempty"`
	Options         []int32  `json:"options,omitempty"`
	AdditionalNotes []string `json:"additionalNotes,omitempty"`
	Hostnames       []string `json:"hostnames"`
	LocationCode    string   `json:"locationCode,omitempty"`
	Quantity        int32    `json:"quantity,omitempty"`
}
