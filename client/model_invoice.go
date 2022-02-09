/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Invoice struct for Invoice
type Invoice struct {
	Id           int32   `json:"id,omitempty"`
	Amount       float32 `json:"amount,omitempty"`
	AmountUnpaid float32 `json:"amountUnpaid,omitempty"`
	ClientId     int32   `json:"clientId,omitempty"`
	Created      int32   `json:"created,omitempty"`
	DatePaid     int32   `json:"datePaid,omitempty"`
	Due          int32   `json:"due,omitempty"`
	Sent         int32   `json:"sent,omitempty"`
	Status       string  `json:"status,omitempty"`
}
