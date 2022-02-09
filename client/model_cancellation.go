/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// Cancellation struct for Cancellation
type Cancellation struct {
	Id             int32     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	ClientId       int32     `json:"clientId,omitempty"`
	StartDate      time.Time `json:"startDate,omitempty"`
	ActiveMonths   int32     `json:"activeMonths,omitempty"`
	PlanId         int32     `json:"planId,omitempty"`
	RequestDate    time.Time `json:"requestDate,omitempty"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
	MonthlyPayment int32     `json:"monthlyPayment,omitempty"`
	DeviceId       int32     `json:"deviceId,omitempty"`
	Cancelled      bool      `json:"cancelled,omitempty"`
	ServiceId      int32     `json:"serviceId,omitempty"`
}
