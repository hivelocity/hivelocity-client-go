/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrderDump struct {
	OrderId int32       `json:"orderId,omitempty"`
	Info    interface{} `json:"info,omitempty"`
	Total   float32     `json:"total,omitempty"`
	Status  string      `json:"status,omitempty"`
	Owner   string      `json:"owner,omitempty"`
}
