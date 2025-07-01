# RevenueWithdrawalStateSucceeded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the state, always “succeeded” | [default to "succeeded"]
**Date** | **int32** | Date the withdrawal was completed in Unix time | 
**Url** | **string** | An HTTPS URL that can be used to see transaction details | 

## Methods

### NewRevenueWithdrawalStateSucceeded

`func NewRevenueWithdrawalStateSucceeded(type_ string, date int32, url string, ) *RevenueWithdrawalStateSucceeded`

NewRevenueWithdrawalStateSucceeded instantiates a new RevenueWithdrawalStateSucceeded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueWithdrawalStateSucceededWithDefaults

`func NewRevenueWithdrawalStateSucceededWithDefaults() *RevenueWithdrawalStateSucceeded`

NewRevenueWithdrawalStateSucceededWithDefaults instantiates a new RevenueWithdrawalStateSucceeded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RevenueWithdrawalStateSucceeded) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevenueWithdrawalStateSucceeded) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevenueWithdrawalStateSucceeded) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *RevenueWithdrawalStateSucceeded) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RevenueWithdrawalStateSucceeded) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RevenueWithdrawalStateSucceeded) SetDate(v int32)`

SetDate sets Date field to given value.


### GetUrl

`func (o *RevenueWithdrawalStateSucceeded) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RevenueWithdrawalStateSucceeded) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RevenueWithdrawalStateSucceeded) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


