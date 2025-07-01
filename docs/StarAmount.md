# StarAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Integer amount of Telegram Stars, rounded to 0; can be negative | 
**NanostarAmount** | Pointer to **int32** | *Optional*. The number of 1/1000000000 shares of Telegram Stars; from -999999999 to 999999999; can be negative if and only if *amount* is non-positive | [optional] 

## Methods

### NewStarAmount

`func NewStarAmount(amount int32, ) *StarAmount`

NewStarAmount instantiates a new StarAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarAmountWithDefaults

`func NewStarAmountWithDefaults() *StarAmount`

NewStarAmountWithDefaults instantiates a new StarAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StarAmount) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StarAmount) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StarAmount) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetNanostarAmount

`func (o *StarAmount) GetNanostarAmount() int32`

GetNanostarAmount returns the NanostarAmount field if non-nil, zero value otherwise.

### GetNanostarAmountOk

`func (o *StarAmount) GetNanostarAmountOk() (*int32, bool)`

GetNanostarAmountOk returns a tuple with the NanostarAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanostarAmount

`func (o *StarAmount) SetNanostarAmount(v int32)`

SetNanostarAmount sets NanostarAmount field to given value.

### HasNanostarAmount

`func (o *StarAmount) HasNanostarAmount() bool`

HasNanostarAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


