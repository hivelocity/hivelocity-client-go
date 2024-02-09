/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TicketSeachResult struct {
	PerPage  int32       `json:"perPage,omitempty"`
	Items    *TicketPost `json:"items,omitempty"`
	Total    int32       `json:"total,omitempty"`
	Page     int32       `json:"page,omitempty"`
	HasNext  bool        `json:"hasNext,omitempty"`
	PrevPage int32       `json:"prevPage,omitempty"`
	HasPrev  bool        `json:"hasPrev,omitempty"`
	NextPage int32       `json:"nextPage,omitempty"`
	Pages    int32       `json:"pages,omitempty"`
}
