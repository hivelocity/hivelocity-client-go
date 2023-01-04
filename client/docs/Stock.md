# Stock

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductQuarterlyPrice** | **float32** | Price per quarter (3 months). | [optional] [default to null]
**ProductDisplayPrice** | **float32** | *DEPRECATED*. | [optional] [default to null]
**ProductSemiAnnuallyPrice** | **float32** | Price per half year (6 months). | [optional] [default to null]
**MonthlyLocationPremium** | **float32** | Additional monthly fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**ProductMonthlyPrice** | **float32** | Price per month. | [optional] [default to null]
**ProductName** | **string** | The unique name of this product. | [optional] [default to null]
**TriennialLocationPremium** | **float32** | Additional triennial fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**ProductBiennialPrice** | **float32** | Price per 2 year period (24 months). | [optional] [default to null]
**DataCenter** | **string** | A facility code. For example &#x60;NYC1&#x60;. | [optional] [default to null]
**ProductOnSale** | **bool** | true|false. Use with &#x60;product_original_price&#x60;. | [optional] [default to null]
**ProductDisabledBillingPeriods** | **[]string** | Orders for the periods in the array will fail. Possible values: hourly|monthly|quarterly|semi-annually|biennial|triennial | [optional] [default to null]
**Core** | **bool** | true|false if core site. | [optional] [default to null]
**ProcessorInfo** | **interface{}** | JSON CPU info for cores, threads, sockets, and vCPUs. | [optional] [default to null]
**ProductGpu** | **string** | Human readable GPU specs. | [optional] [default to null]
**ProductHourlyPrice** | **float32** | Price per hour. | [optional] [default to null]
**ProductAnnuallyPrice** | **float32** | Price per year (12 months). | [optional] [default to null]
**AnnuallyLocationPremium** | **float32** | Additional annual fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**BiennialLocationPremium** | **float32** | Additional biennial fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**SemiAnnuallyLocationPremium** | **float32** | Additional semi-annual fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**ProductTriennialPrice** | **float32** | Price per 3 year period (36 months). | [optional] [default to null]
**ProductCpuCores** | **string** | Human readable CPU core and thread info in HTML format. | [optional] [default to null]
**QuarterlyLocationPremium** | **float32** | Additional quarterly fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**Edge** | **bool** | true|false if edge site. | [optional] [default to null]
**ProductOriginalPrice** | **float32** | Retail price of product. Use with &#x60;product_on_sale&#x60;. | [optional] [default to null]
**ProductMemory** | **string** | Human readable memory specs. | [optional] [default to null]
**HourlyLocationPremium** | **float32** | Additional hourly fee for this &#x60;product_id&#x60; in this &#x60;data_center&#x60;. | [optional] [default to null]
**ProductDrive** | **string** | Human readable drive specs. Can include multiple drives. | [optional] [default to null]
**Stock** | **string** | available|limited|unavailable. Limited denotes minimal stock. | [optional] [default to null]
**ProductCpu** | **string** | Human readable CPU specs. | [optional] [default to null]
**ProductId** | **int32** | The unique ID of this product. | [optional] [default to null]
**ProductBandwidth** | **string** | Human readable networking specs in the format: Free Outbound Transfer / NIC Size | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


