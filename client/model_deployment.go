/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Deployment struct {
	DeploymentId            int32         `json:"deploymentId,omitempty"`
	StartedProvisioning     bool          `json:"startedProvisioning,omitempty"`
	DeploymentName          string        `json:"deploymentName,omitempty"`
	Empty                   bool          `json:"empty,omitempty"`
	Price                   float32       `json:"price,omitempty"`
	OrderNumber             string        `json:"orderNumber,omitempty"`
	DeploymentConfiguration []interface{} `json:"deploymentConfiguration,omitempty"`
}
