# GiftInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gift** | [**Gift**](Gift.md) |  | 
**OwnedGiftId** | Pointer to **string** | *Optional*. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts | [optional] 
**ConvertStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that can be claimed by the receiver by converting the gift; omitted if conversion to Telegram Stars is impossible | [optional] 
**PrepaidUpgradeStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that were prepaid by the sender for the ability to upgrade the gift | [optional] 
**CanBeUpgraded** | Pointer to **bool** | *Optional*. True, if the gift can be upgraded to a unique gift | [optional] [default to true]
**Text** | Pointer to **string** | *Optional*. Text of the message that was added to the gift | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the text | [optional] 
**IsPrivate** | Pointer to **bool** | *Optional*. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them | [optional] [default to true]

## Methods

### NewGiftInfo

`func NewGiftInfo(gift Gift, ) *GiftInfo`

NewGiftInfo instantiates a new GiftInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftInfoWithDefaults

`func NewGiftInfoWithDefaults() *GiftInfo`

NewGiftInfoWithDefaults instantiates a new GiftInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGift

`func (o *GiftInfo) GetGift() Gift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *GiftInfo) GetGiftOk() (*Gift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *GiftInfo) SetGift(v Gift)`

SetGift sets Gift field to given value.


### GetOwnedGiftId

`func (o *GiftInfo) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *GiftInfo) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *GiftInfo) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.

### HasOwnedGiftId

`func (o *GiftInfo) HasOwnedGiftId() bool`

HasOwnedGiftId returns a boolean if a field has been set.

### GetConvertStarCount

`func (o *GiftInfo) GetConvertStarCount() int32`

GetConvertStarCount returns the ConvertStarCount field if non-nil, zero value otherwise.

### GetConvertStarCountOk

`func (o *GiftInfo) GetConvertStarCountOk() (*int32, bool)`

GetConvertStarCountOk returns a tuple with the ConvertStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertStarCount

`func (o *GiftInfo) SetConvertStarCount(v int32)`

SetConvertStarCount sets ConvertStarCount field to given value.

### HasConvertStarCount

`func (o *GiftInfo) HasConvertStarCount() bool`

HasConvertStarCount returns a boolean if a field has been set.

### GetPrepaidUpgradeStarCount

`func (o *GiftInfo) GetPrepaidUpgradeStarCount() int32`

GetPrepaidUpgradeStarCount returns the PrepaidUpgradeStarCount field if non-nil, zero value otherwise.

### GetPrepaidUpgradeStarCountOk

`func (o *GiftInfo) GetPrepaidUpgradeStarCountOk() (*int32, bool)`

GetPrepaidUpgradeStarCountOk returns a tuple with the PrepaidUpgradeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidUpgradeStarCount

`func (o *GiftInfo) SetPrepaidUpgradeStarCount(v int32)`

SetPrepaidUpgradeStarCount sets PrepaidUpgradeStarCount field to given value.

### HasPrepaidUpgradeStarCount

`func (o *GiftInfo) HasPrepaidUpgradeStarCount() bool`

HasPrepaidUpgradeStarCount returns a boolean if a field has been set.

### GetCanBeUpgraded

`func (o *GiftInfo) GetCanBeUpgraded() bool`

GetCanBeUpgraded returns the CanBeUpgraded field if non-nil, zero value otherwise.

### GetCanBeUpgradedOk

`func (o *GiftInfo) GetCanBeUpgradedOk() (*bool, bool)`

GetCanBeUpgradedOk returns a tuple with the CanBeUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeUpgraded

`func (o *GiftInfo) SetCanBeUpgraded(v bool)`

SetCanBeUpgraded sets CanBeUpgraded field to given value.

### HasCanBeUpgraded

`func (o *GiftInfo) HasCanBeUpgraded() bool`

HasCanBeUpgraded returns a boolean if a field has been set.

### GetText

`func (o *GiftInfo) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GiftInfo) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GiftInfo) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GiftInfo) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEntities

`func (o *GiftInfo) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *GiftInfo) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *GiftInfo) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *GiftInfo) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetIsPrivate

`func (o *GiftInfo) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *GiftInfo) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *GiftInfo) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *GiftInfo) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


