# SuggestedPostPaid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestedPostMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Currency** | **string** | Currency in which the payment was made. Currently, one of “XTR” for Telegram Stars or “TON” for toncoins | 
**Amount** | Pointer to **int32** | *Optional*. The amount of the currency that was received by the channel in nanotoncoins; for payments in toncoins only | [optional] 
**StarAmount** | Pointer to [**StarAmount**](StarAmount.md) |  | [optional] 

## Methods

### NewSuggestedPostPaid

`func NewSuggestedPostPaid(currency string, ) *SuggestedPostPaid`

NewSuggestedPostPaid instantiates a new SuggestedPostPaid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostPaidWithDefaults

`func NewSuggestedPostPaidWithDefaults() *SuggestedPostPaid`

NewSuggestedPostPaidWithDefaults instantiates a new SuggestedPostPaid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestedPostMessage

`func (o *SuggestedPostPaid) GetSuggestedPostMessage() Message`

GetSuggestedPostMessage returns the SuggestedPostMessage field if non-nil, zero value otherwise.

### GetSuggestedPostMessageOk

`func (o *SuggestedPostPaid) GetSuggestedPostMessageOk() (*Message, bool)`

GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPostMessage

`func (o *SuggestedPostPaid) SetSuggestedPostMessage(v Message)`

SetSuggestedPostMessage sets SuggestedPostMessage field to given value.

### HasSuggestedPostMessage

`func (o *SuggestedPostPaid) HasSuggestedPostMessage() bool`

HasSuggestedPostMessage returns a boolean if a field has been set.

### GetCurrency

`func (o *SuggestedPostPaid) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SuggestedPostPaid) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SuggestedPostPaid) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *SuggestedPostPaid) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SuggestedPostPaid) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SuggestedPostPaid) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SuggestedPostPaid) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStarAmount

`func (o *SuggestedPostPaid) GetStarAmount() StarAmount`

GetStarAmount returns the StarAmount field if non-nil, zero value otherwise.

### GetStarAmountOk

`func (o *SuggestedPostPaid) GetStarAmountOk() (*StarAmount, bool)`

GetStarAmountOk returns a tuple with the StarAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarAmount

`func (o *SuggestedPostPaid) SetStarAmount(v StarAmount)`

SetStarAmount sets StarAmount field to given value.

### HasStarAmount

`func (o *SuggestedPostPaid) HasStarAmount() bool`

HasStarAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


