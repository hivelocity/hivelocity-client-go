/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// TicketSeachResult struct for TicketSeachResult
type TicketSeachResult struct {
	Pages    int32      `json:"pages,omitempty"`
	HasPrev  bool       `json:"hasPrev,omitempty"`
	Items    TicketPost `json:"items,omitempty"`
	HasNext  bool       `json:"hasNext,omitempty"`
	PerPage  int32      `json:"perPage,omitempty"`
	NextPage int32      `json:"nextPage,omitempty"`
	Page     int32      `json:"page,omitempty"`
	Total    int32      `json:"total,omitempty"`
	PrevPage int32      `json:"prevPage,omitempty"`
}