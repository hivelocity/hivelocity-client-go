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

type MaintenanceEvent struct {
	// Maintenance Event End (UTC)
	End time.Time `json:"end,omitempty"`
	// Device IDs affected by event
	DeviceIds []int32 `json:"device_ids,omitempty"`
	// Maintenance Event Title
	Title string `json:"title,omitempty"`
	// Maintenance Event ID
	MaintenanceEventId int32 `json:"maintenanceEventId,omitempty"`
	// Maintenance Event Start (UTC)
	Start time.Time `json:"start,omitempty"`
	// Maintenance Event Description
	Description string `json:"description,omitempty"`
}
