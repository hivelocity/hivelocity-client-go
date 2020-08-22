/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Stock struct for Stock
type Stock struct {
	Stock                  string  `json:"stock,omitempty"`
	ProductName            string  `json:"product_name,omitempty"`
	ProductDrive           string  `json:"product_drive,omitempty"`
	ProductOriginalPrice   float32 `json:"product_original_price,omitempty"`
	ProductId              int32   `json:"product_id,omitempty"`
	Edge                   bool    `json:"edge,omitempty"`
	Core                   bool    `json:"core,omitempty"`
	ProductCpuCores        string  `json:"product_cpu_cores,omitempty"`
	ProductBandwidth       string  `json:"product_bandwidth,omitempty"`
	ProductOnSale          bool    `json:"product_on_sale,omitempty"`
	ProductCpu             string  `json:"product_cpu,omitempty"`
	ProductDisplayPrice    float32 `json:"product_display_price,omitempty"`
	MonthlyLocationPremium float32 `json:"monthly_location_premium,omitempty"`
	ProductMonthlyPrice    float32 `json:"product_monthly_price,omitempty"`
	ProductMemory          string  `json:"product_memory,omitempty"`
}