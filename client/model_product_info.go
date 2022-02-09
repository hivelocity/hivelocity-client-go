/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// ProductInfo struct for ProductInfo
type ProductInfo struct {
	ProductGpu                    string   `json:"product_gpu,omitempty"`
	ProductOnSale                 bool     `json:"product_on_sale,omitempty"`
	ProductMemory                 string   `json:"product_memory,omitempty"`
	DataCenter                    string   `json:"data_center,omitempty"`
	AnnuallyLocationPremium       float32  `json:"annually_location_premium,omitempty"`
	ProductBiennialPrice          float32  `json:"product_biennial_price,omitempty"`
	ProductMonthlyPrice           float32  `json:"product_monthly_price,omitempty"`
	QuarterlyLocationPremium      float32  `json:"quarterly_location_premium,omitempty"`
	Edge                          bool     `json:"edge,omitempty"`
	ProductSemiAnnuallyPrice      float32  `json:"product_semi_annually_price,omitempty"`
	HourlyLocationPremium         float32  `json:"hourly_location_premium,omitempty"`
	ProductAnnuallyPrice          float32  `json:"product_annually_price,omitempty"`
	ProductCpu                    string   `json:"product_cpu,omitempty"`
	ProductHourlyPrice            float32  `json:"product_hourly_price,omitempty"`
	ProductTriennialPrice         float32  `json:"product_triennial_price,omitempty"`
	ProductCpuCores               string   `json:"product_cpu_cores,omitempty"`
	ProductQuarterlyPrice         float32  `json:"product_quarterly_price,omitempty"`
	ProductDisplayPrice           float32  `json:"product_display_price,omitempty"`
	Core                          bool     `json:"core,omitempty"`
	SemiAnnuallyLocationPremium   float32  `json:"semi_annually_location_premium,omitempty"`
	ProductId                     int32    `json:"product_id,omitempty"`
	Stock                         string   `json:"stock,omitempty"`
	ProductDisabledBillingPeriods []string `json:"product_disabled_billing_periods,omitempty"`
	ProductName                   string   `json:"product_name,omitempty"`
	ProductOriginalPrice          float32  `json:"product_original_price,omitempty"`
	ProductBandwidth              string   `json:"product_bandwidth,omitempty"`
	TriennialLocationPremium      float32  `json:"triennial_location_premium,omitempty"`
	ProductDrive                  string   `json:"product_drive,omitempty"`
	BiennialLocationPremium       float32  `json:"biennial_location_premium,omitempty"`
	MonthlyLocationPremium        float32  `json:"monthly_location_premium,omitempty"`
}
