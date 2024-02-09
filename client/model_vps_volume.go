/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type VpsVolume struct {
	// The unique VPS instance volume ID.
	VolumeId string `json:"volumeId,omitempty"`
	// Disk type. Options are `ROOT` for root volumes and `DATADISK` for additional disks.
	Type_ string `json:"type,omitempty"`
	// Disk size in bytes
	Size int32 `json:"size,omitempty"`
	// The created date/time of the VPS instance volume.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Device ID the disk is attached to
	DeviceId int32 `json:"deviceId,omitempty"`
	// The facility code associated with the VPS Volume. For example: `TPA1`.
	FacilityCode string `json:"facilityCode,omitempty"`
}