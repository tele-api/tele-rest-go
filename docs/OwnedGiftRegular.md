# OwnedGiftRegular

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the gift, always “regular” | [default to "regular"]
**Gift** | [**Gift**](Gift.md) |  | 
**OwnedGiftId** | Pointer to **string** | *Optional*. Unique identifier of the gift for the bot; for gifts received on behalf of business accounts only | [optional] 
**SenderUser** | Pointer to [**User**](User.md) |  | [optional] 
**SendDate** | **int32** | Date the gift was sent in Unix time | 
**Text** | Pointer to **string** | *Optional*. Text of the message that was added to the gift | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the text | [optional] 
**IsPrivate** | Pointer to **bool** | *Optional*. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them | [optional] [default to true]
**IsSaved** | Pointer to **bool** | *Optional*. True, if the gift is displayed on the account&#39;s profile page; for gifts received on behalf of business accounts only | [optional] [default to true]
**CanBeUpgraded** | Pointer to **bool** | *Optional*. True, if the gift can be upgraded to a unique gift; for gifts received on behalf of business accounts only | [optional] [default to true]
**WasRefunded** | Pointer to **bool** | *Optional*. True, if the gift was refunded and isn&#39;t available anymore | [optional] [default to true]
**ConvertStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that can be claimed by the receiver instead of the gift; omitted if the gift cannot be converted to Telegram Stars | [optional] 
**PrepaidUpgradeStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that were paid by the sender for the ability to upgrade the gift | [optional] 

## Methods

### NewOwnedGiftRegular

`func NewOwnedGiftRegular(type_ string, gift Gift, sendDate int32, ) *OwnedGiftRegular`

NewOwnedGiftRegular instantiates a new OwnedGiftRegular object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnedGiftRegularWithDefaults

`func NewOwnedGiftRegularWithDefaults() *OwnedGiftRegular`

NewOwnedGiftRegularWithDefaults instantiates a new OwnedGiftRegular object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OwnedGiftRegular) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OwnedGiftRegular) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OwnedGiftRegular) SetType(v string)`

SetType sets Type field to given value.


### GetGift

`func (o *OwnedGiftRegular) GetGift() Gift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *OwnedGiftRegular) GetGiftOk() (*Gift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *OwnedGiftRegular) SetGift(v Gift)`

SetGift sets Gift field to given value.


### GetOwnedGiftId

`func (o *OwnedGiftRegular) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *OwnedGiftRegular) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *OwnedGiftRegular) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.

### HasOwnedGiftId

`func (o *OwnedGiftRegular) HasOwnedGiftId() bool`

HasOwnedGiftId returns a boolean if a field has been set.

### GetSenderUser

`func (o *OwnedGiftRegular) GetSenderUser() User`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *OwnedGiftRegular) GetSenderUserOk() (*User, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *OwnedGiftRegular) SetSenderUser(v User)`

SetSenderUser sets SenderUser field to given value.

### HasSenderUser

`func (o *OwnedGiftRegular) HasSenderUser() bool`

HasSenderUser returns a boolean if a field has been set.

### GetSendDate

`func (o *OwnedGiftRegular) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *OwnedGiftRegular) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *OwnedGiftRegular) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.


### GetText

`func (o *OwnedGiftRegular) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OwnedGiftRegular) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OwnedGiftRegular) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OwnedGiftRegular) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEntities

`func (o *OwnedGiftRegular) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *OwnedGiftRegular) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *OwnedGiftRegular) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *OwnedGiftRegular) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetIsPrivate

`func (o *OwnedGiftRegular) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *OwnedGiftRegular) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *OwnedGiftRegular) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *OwnedGiftRegular) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsSaved

`func (o *OwnedGiftRegular) GetIsSaved() bool`

GetIsSaved returns the IsSaved field if non-nil, zero value otherwise.

### GetIsSavedOk

`func (o *OwnedGiftRegular) GetIsSavedOk() (*bool, bool)`

GetIsSavedOk returns a tuple with the IsSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSaved

`func (o *OwnedGiftRegular) SetIsSaved(v bool)`

SetIsSaved sets IsSaved field to given value.

### HasIsSaved

`func (o *OwnedGiftRegular) HasIsSaved() bool`

HasIsSaved returns a boolean if a field has been set.

### GetCanBeUpgraded

`func (o *OwnedGiftRegular) GetCanBeUpgraded() bool`

GetCanBeUpgraded returns the CanBeUpgraded field if non-nil, zero value otherwise.

### GetCanBeUpgradedOk

`func (o *OwnedGiftRegular) GetCanBeUpgradedOk() (*bool, bool)`

GetCanBeUpgradedOk returns a tuple with the CanBeUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeUpgraded

`func (o *OwnedGiftRegular) SetCanBeUpgraded(v bool)`

SetCanBeUpgraded sets CanBeUpgraded field to given value.

### HasCanBeUpgraded

`func (o *OwnedGiftRegular) HasCanBeUpgraded() bool`

HasCanBeUpgraded returns a boolean if a field has been set.

### GetWasRefunded

`func (o *OwnedGiftRegular) GetWasRefunded() bool`

GetWasRefunded returns the WasRefunded field if non-nil, zero value otherwise.

### GetWasRefundedOk

`func (o *OwnedGiftRegular) GetWasRefundedOk() (*bool, bool)`

GetWasRefundedOk returns a tuple with the WasRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasRefunded

`func (o *OwnedGiftRegular) SetWasRefunded(v bool)`

SetWasRefunded sets WasRefunded field to given value.

### HasWasRefunded

`func (o *OwnedGiftRegular) HasWasRefunded() bool`

HasWasRefunded returns a boolean if a field has been set.

### GetConvertStarCount

`func (o *OwnedGiftRegular) GetConvertStarCount() int32`

GetConvertStarCount returns the ConvertStarCount field if non-nil, zero value otherwise.

### GetConvertStarCountOk

`func (o *OwnedGiftRegular) GetConvertStarCountOk() (*int32, bool)`

GetConvertStarCountOk returns a tuple with the ConvertStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertStarCount

`func (o *OwnedGiftRegular) SetConvertStarCount(v int32)`

SetConvertStarCount sets ConvertStarCount field to given value.

### HasConvertStarCount

`func (o *OwnedGiftRegular) HasConvertStarCount() bool`

HasConvertStarCount returns a boolean if a field has been set.

### GetPrepaidUpgradeStarCount

`func (o *OwnedGiftRegular) GetPrepaidUpgradeStarCount() int32`

GetPrepaidUpgradeStarCount returns the PrepaidUpgradeStarCount field if non-nil, zero value otherwise.

### GetPrepaidUpgradeStarCountOk

`func (o *OwnedGiftRegular) GetPrepaidUpgradeStarCountOk() (*int32, bool)`

GetPrepaidUpgradeStarCountOk returns a tuple with the PrepaidUpgradeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidUpgradeStarCount

`func (o *OwnedGiftRegular) SetPrepaidUpgradeStarCount(v int32)`

SetPrepaidUpgradeStarCount sets PrepaidUpgradeStarCount field to given value.

### HasPrepaidUpgradeStarCount

`func (o *OwnedGiftRegular) HasPrepaidUpgradeStarCount() bool`

HasPrepaidUpgradeStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


