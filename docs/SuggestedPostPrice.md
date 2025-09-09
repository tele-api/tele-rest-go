# SuggestedPostPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency in which the post will be paid. Currently, must be one of “XTR” for Telegram Stars or “TON” for toncoins | 
**Amount** | **int32** | The amount of the currency that will be paid for the post in the *smallest units* of the currency, i.e. Telegram Stars or nanotoncoins. Currently, price in Telegram Stars must be between 5 and 100000, and price in nanotoncoins must be between 10000000 and 10000000000000. | 

## Methods

### NewSuggestedPostPrice

`func NewSuggestedPostPrice(currency string, amount int32, ) *SuggestedPostPrice`

NewSuggestedPostPrice instantiates a new SuggestedPostPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostPriceWithDefaults

`func NewSuggestedPostPriceWithDefaults() *SuggestedPostPrice`

NewSuggestedPostPriceWithDefaults instantiates a new SuggestedPostPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SuggestedPostPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SuggestedPostPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SuggestedPostPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *SuggestedPostPrice) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SuggestedPostPrice) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SuggestedPostPrice) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


