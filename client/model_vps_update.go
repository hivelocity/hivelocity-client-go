/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VpsUpdate struct {
	// ID of VPS product to scale to
	ProductId int32 `json:"productId,omitempty"`
	// List of SSH Key IDs
	SshKeyIds []int32 `json:"sshKeyIds,omitempty"`
	// Cloud-init script
	CloudInit string `json:"cloudInit,omitempty"`
}
