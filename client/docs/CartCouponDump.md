# CartCouponDump

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageTypeLimit** | **[]string** | Coupon usage type limit max_total|one_per_user|one_per_user_in_24_hours | [optional] [default to null]
**MaxUseSameCart** | **int32** | Max use same cart | [optional] [default to null]
**ExpireType** | **string** | Coupon expire type first_billing_period|number_of_days|never_expire | [default to null]
**ServiceEligibility** | **string** | Coupon service eligibility none|all | [optional] [default to null]
**MinimumPurchaseAmount** | **float32** | Minimum purchase requirement | [optional] [default to null]
**Name** | **string** | Coupon name | [default to null]
**ExpireValue** | **int32** | Coupon expire value | [optional] [default to null]
**AccountApplicableCredit** | **bool** | Account applicable credit | [optional] [default to null]
**End** | **string** | Coupon end date | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Active** | **bool** | Coupon active | [optional] [default to null]
**MaximumPurchaseAmount** | **float32** | Maximum purchase requirement max | [optional] [default to null]
**Code** | **string** | Coupon code | [default to null]
**Value** | **float32** | Coupon value | [default to null]
**CustomerEligibilityRules** | **[]string** | Eligible client rules to use this coupon block_enterprise_client|block_reseller|block_partner|has_no_purchase|has_one_purchase|has_multiple_purchase|account_created_recently | [optional] [default to null]
**AllowOtherDiscounts** | **bool** | Allow other discounts | [default to null]
**MaxUseSameCartUseType** | **string** | Coupon max use same cart use type first_item|highest_price | [optional] [default to null]
**CustomerEligibility** | **string** | Coupon customer eligibility all|rules|specific | [default to null]
**UserLoginRequired** | **interface{}** |  | [optional] [default to null]
**Id** | **int32** |  | [optional] [default to null]
**ProductIds** | **[]int32** | Products ids the coupon will be available | [optional] [default to null]
**Type_** | **string** | Coupon type %|$|credit | [default to null]
**MaxUsageLimit** | **int32** | Coupon max uses | [optional] [default to null]
**Start** | **string** | Coupon start date | [default to null]
**ClientIds** | **[]int32** | Clients ids the coupon will be available | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


