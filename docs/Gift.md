# Gift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the gift | 
**Sticker** | [**Sticker**](Sticker.md) |  | 
**StarCount** | **int32** | The number of Telegram Stars that must be paid to send the sticker | 
**UpgradeStarCount** | Pointer to **int32** | *Optional*. The number of Telegram Stars that must be paid to upgrade the gift to a unique one | [optional] 
**TotalCount** | Pointer to **int32** | *Optional*. The total number of the gifts of this type that can be sent; for limited gifts only | [optional] 
**RemainingCount** | Pointer to **int32** | *Optional*. The number of remaining gifts of this type that can be sent; for limited gifts only | [optional] 

## Methods

### NewGift

`func NewGift(id string, sticker Sticker, starCount int32, ) *Gift`

NewGift instantiates a new Gift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftWithDefaults

`func NewGiftWithDefaults() *Gift`

NewGiftWithDefaults instantiates a new Gift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gift) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gift) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gift) SetId(v string)`

SetId sets Id field to given value.


### GetSticker

`func (o *Gift) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *Gift) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *Gift) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.


### GetStarCount

`func (o *Gift) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *Gift) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *Gift) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.


### GetUpgradeStarCount

`func (o *Gift) GetUpgradeStarCount() int32`

GetUpgradeStarCount returns the UpgradeStarCount field if non-nil, zero value otherwise.

### GetUpgradeStarCountOk

`func (o *Gift) GetUpgradeStarCountOk() (*int32, bool)`

GetUpgradeStarCountOk returns a tuple with the UpgradeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStarCount

`func (o *Gift) SetUpgradeStarCount(v int32)`

SetUpgradeStarCount sets UpgradeStarCount field to given value.

### HasUpgradeStarCount

`func (o *Gift) HasUpgradeStarCount() bool`

HasUpgradeStarCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *Gift) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Gift) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Gift) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Gift) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetRemainingCount

`func (o *Gift) GetRemainingCount() int32`

GetRemainingCount returns the RemainingCount field if non-nil, zero value otherwise.

### GetRemainingCountOk

`func (o *Gift) GetRemainingCountOk() (*int32, bool)`

GetRemainingCountOk returns a tuple with the RemainingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCount

`func (o *Gift) SetRemainingCount(v int32)`

SetRemainingCount sets RemainingCount field to given value.

### HasRemainingCount

`func (o *Gift) HasRemainingCount() bool`

HasRemainingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


