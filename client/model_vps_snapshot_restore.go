/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VpsSnapshotRestore struct {
	// The unique client acount ID.
	ClientId int32 `json:"clientId,omitempty"`
	// The facility code associated with the VPS instance. For example:TPA1.
	FacilityCode string `json:"facilityCode"`
}