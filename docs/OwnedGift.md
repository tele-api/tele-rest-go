# OwnedGift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the gift, always “unique” | [default to "unique"]
**Gift** | [**UniqueGift**](UniqueGift.md) |  | 
**OwnedGiftId** | Pointer to **string** | *Optional*. Unique identifier of the received gift for the bot; for gifts received on behalf of business accounts only | [optional] 
**SenderUser** | Pointer to [**User**](User.md) |  | [optional] 
**SendDate** | **int32** | Date the gift was sent in Unix time | 
**Text** | Pointer to **string** | *Optional*. Text of the message that was added to the gift | [optional] 
**Entities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the text | [optional] 
**IsPrivate** | Pointer to **bool** | *Optional*. *True*, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them | [optional] [default to true]
**IsSaved** | Pointer to **bool** | *Optional*. *True*, if the gift is displayed on the account&#39;s profile page; for gifts received on behalf of business accounts only | [optional] [default to true]
**CanBeUpgraded** | Pointer to **bool** | *Optional*. *True*, if the gift can be upgraded to a unique gift; for gifts received on behalf of business accounts only | [optional] [default to true]
**WasRefunded** | Pointer to **bool** | *Optional*. *True*, if the gift was refunded and isn&#39;t available anymore | [optional] [default to true]
**ConvertStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that can be claimed by the receiver instead of the gift; omitted if the gift cannot be converted to Telegram Stars | [optional] 
**PrepaidUpgradeStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that were paid by the sender for the ability to upgrade the gift | [optional] 
**CanBeTransferred** | Pointer to **bool** | *Optional*. *True*, if the gift can be transferred to another owner; for gifts received on behalf of business accounts only | [optional] [default to true]
**TransferStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift | [optional] 
**NextTransferDate** | Pointer to **int32** | *Optional*. Point in time (Unix timestamp) when the gift can be transferred. If it is in the past, then the gift can be transferred now | [optional] 

## Methods

### NewOwnedGift

`func NewOwnedGift(type_ string, gift UniqueGift, sendDate int32, ) *OwnedGift`

NewOwnedGift instantiates a new OwnedGift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnedGiftWithDefaults

`func NewOwnedGiftWithDefaults() *OwnedGift`

NewOwnedGiftWithDefaults instantiates a new OwnedGift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OwnedGift) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OwnedGift) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OwnedGift) SetType(v string)`

SetType sets Type field to given value.


### GetGift

`func (o *OwnedGift) GetGift() UniqueGift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *OwnedGift) GetGiftOk() (*UniqueGift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *OwnedGift) SetGift(v UniqueGift)`

SetGift sets Gift field to given value.


### GetOwnedGiftId

`func (o *OwnedGift) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *OwnedGift) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *OwnedGift) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.

### HasOwnedGiftId

`func (o *OwnedGift) HasOwnedGiftId() bool`

HasOwnedGiftId returns a boolean if a field has been set.

### GetSenderUser

`func (o *OwnedGift) GetSenderUser() User`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *OwnedGift) GetSenderUserOk() (*User, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *OwnedGift) SetSenderUser(v User)`

SetSenderUser sets SenderUser field to given value.

### HasSenderUser

`func (o *OwnedGift) HasSenderUser() bool`

HasSenderUser returns a boolean if a field has been set.

### GetSendDate

`func (o *OwnedGift) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *OwnedGift) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *OwnedGift) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.


### GetText

`func (o *OwnedGift) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OwnedGift) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OwnedGift) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OwnedGift) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEntities

`func (o *OwnedGift) GetEntities() []MessageEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *OwnedGift) GetEntitiesOk() (*[]MessageEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *OwnedGift) SetEntities(v []MessageEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *OwnedGift) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetIsPrivate

`func (o *OwnedGift) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *OwnedGift) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *OwnedGift) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *OwnedGift) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsSaved

`func (o *OwnedGift) GetIsSaved() bool`

GetIsSaved returns the IsSaved field if non-nil, zero value otherwise.

### GetIsSavedOk

`func (o *OwnedGift) GetIsSavedOk() (*bool, bool)`

GetIsSavedOk returns a tuple with the IsSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSaved

`func (o *OwnedGift) SetIsSaved(v bool)`

SetIsSaved sets IsSaved field to given value.

### HasIsSaved

`func (o *OwnedGift) HasIsSaved() bool`

HasIsSaved returns a boolean if a field has been set.

### GetCanBeUpgraded

`func (o *OwnedGift) GetCanBeUpgraded() bool`

GetCanBeUpgraded returns the CanBeUpgraded field if non-nil, zero value otherwise.

### GetCanBeUpgradedOk

`func (o *OwnedGift) GetCanBeUpgradedOk() (*bool, bool)`

GetCanBeUpgradedOk returns a tuple with the CanBeUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeUpgraded

`func (o *OwnedGift) SetCanBeUpgraded(v bool)`

SetCanBeUpgraded sets CanBeUpgraded field to given value.

### HasCanBeUpgraded

`func (o *OwnedGift) HasCanBeUpgraded() bool`

HasCanBeUpgraded returns a boolean if a field has been set.

### GetWasRefunded

`func (o *OwnedGift) GetWasRefunded() bool`

GetWasRefunded returns the WasRefunded field if non-nil, zero value otherwise.

### GetWasRefundedOk

`func (o *OwnedGift) GetWasRefundedOk() (*bool, bool)`

GetWasRefundedOk returns a tuple with the WasRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasRefunded

`func (o *OwnedGift) SetWasRefunded(v bool)`

SetWasRefunded sets WasRefunded field to given value.

### HasWasRefunded

`func (o *OwnedGift) HasWasRefunded() bool`

HasWasRefunded returns a boolean if a field has been set.

### GetConvertStarCount

`func (o *OwnedGift) GetConvertStarCount() int32`

GetConvertStarCount returns the ConvertStarCount field if non-nil, zero value otherwise.

### GetConvertStarCountOk

`func (o *OwnedGift) GetConvertStarCountOk() (*int32, bool)`

GetConvertStarCountOk returns a tuple with the ConvertStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertStarCount

`func (o *OwnedGift) SetConvertStarCount(v int32)`

SetConvertStarCount sets ConvertStarCount field to given value.

### HasConvertStarCount

`func (o *OwnedGift) HasConvertStarCount() bool`

HasConvertStarCount returns a boolean if a field has been set.

### GetPrepaidUpgradeStarCount

`func (o *OwnedGift) GetPrepaidUpgradeStarCount() int32`

GetPrepaidUpgradeStarCount returns the PrepaidUpgradeStarCount field if non-nil, zero value otherwise.

### GetPrepaidUpgradeStarCountOk

`func (o *OwnedGift) GetPrepaidUpgradeStarCountOk() (*int32, bool)`

GetPrepaidUpgradeStarCountOk returns a tuple with the PrepaidUpgradeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidUpgradeStarCount

`func (o *OwnedGift) SetPrepaidUpgradeStarCount(v int32)`

SetPrepaidUpgradeStarCount sets PrepaidUpgradeStarCount field to given value.

### HasPrepaidUpgradeStarCount

`func (o *OwnedGift) HasPrepaidUpgradeStarCount() bool`

HasPrepaidUpgradeStarCount returns a boolean if a field has been set.

### GetCanBeTransferred

`func (o *OwnedGift) GetCanBeTransferred() bool`

GetCanBeTransferred returns the CanBeTransferred field if non-nil, zero value otherwise.

### GetCanBeTransferredOk

`func (o *OwnedGift) GetCanBeTransferredOk() (*bool, bool)`

GetCanBeTransferredOk returns a tuple with the CanBeTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeTransferred

`func (o *OwnedGift) SetCanBeTransferred(v bool)`

SetCanBeTransferred sets CanBeTransferred field to given value.

### HasCanBeTransferred

`func (o *OwnedGift) HasCanBeTransferred() bool`

HasCanBeTransferred returns a boolean if a field has been set.

### GetTransferStarCount

`func (o *OwnedGift) GetTransferStarCount() int32`

GetTransferStarCount returns the TransferStarCount field if non-nil, zero value otherwise.

### GetTransferStarCountOk

`func (o *OwnedGift) GetTransferStarCountOk() (*int32, bool)`

GetTransferStarCountOk returns a tuple with the TransferStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferStarCount

`func (o *OwnedGift) SetTransferStarCount(v int32)`

SetTransferStarCount sets TransferStarCount field to given value.

### HasTransferStarCount

`func (o *OwnedGift) HasTransferStarCount() bool`

HasTransferStarCount returns a boolean if a field has been set.

### GetNextTransferDate

`func (o *OwnedGift) GetNextTransferDate() int32`

GetNextTransferDate returns the NextTransferDate field if non-nil, zero value otherwise.

### GetNextTransferDateOk

`func (o *OwnedGift) GetNextTransferDateOk() (*int32, bool)`

GetNextTransferDateOk returns a tuple with the NextTransferDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTransferDate

`func (o *OwnedGift) SetNextTransferDate(v int32)`

SetNextTransferDate sets NextTransferDate field to given value.

### HasNextTransferDate

`func (o *OwnedGift) HasNextTransferDate() bool`

HasNextTransferDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


