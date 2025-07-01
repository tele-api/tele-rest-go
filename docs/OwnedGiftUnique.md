# OwnedGiftUnique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the gift, always “unique” | [default to "unique"]
**Gift** | [**UniqueGift**](UniqueGift.md) |  | 
**OwnedGiftId** | Pointer to **string** | *Optional*. Unique identifier of the received gift for the bot; for gifts received on behalf of business accounts only | [optional] 
**SenderUser** | Pointer to [**User**](User.md) |  | [optional] 
**SendDate** | **int32** | Date the gift was sent in Unix time | 
**IsSaved** | Pointer to **bool** | *Optional*. True, if the gift is displayed on the account&#39;s profile page; for gifts received on behalf of business accounts only | [optional] [default to true]
**CanBeTransferred** | Pointer to **bool** | *Optional*. True, if the gift can be transferred to another owner; for gifts received on behalf of business accounts only | [optional] [default to true]
**TransferStarCount** | Pointer to **int32** | *Optional*. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift | [optional] 

## Methods

### NewOwnedGiftUnique

`func NewOwnedGiftUnique(type_ string, gift UniqueGift, sendDate int32, ) *OwnedGiftUnique`

NewOwnedGiftUnique instantiates a new OwnedGiftUnique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnedGiftUniqueWithDefaults

`func NewOwnedGiftUniqueWithDefaults() *OwnedGiftUnique`

NewOwnedGiftUniqueWithDefaults instantiates a new OwnedGiftUnique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OwnedGiftUnique) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OwnedGiftUnique) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OwnedGiftUnique) SetType(v string)`

SetType sets Type field to given value.


### GetGift

`func (o *OwnedGiftUnique) GetGift() UniqueGift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *OwnedGiftUnique) GetGiftOk() (*UniqueGift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *OwnedGiftUnique) SetGift(v UniqueGift)`

SetGift sets Gift field to given value.


### GetOwnedGiftId

`func (o *OwnedGiftUnique) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *OwnedGiftUnique) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *OwnedGiftUnique) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.

### HasOwnedGiftId

`func (o *OwnedGiftUnique) HasOwnedGiftId() bool`

HasOwnedGiftId returns a boolean if a field has been set.

### GetSenderUser

`func (o *OwnedGiftUnique) GetSenderUser() User`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *OwnedGiftUnique) GetSenderUserOk() (*User, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *OwnedGiftUnique) SetSenderUser(v User)`

SetSenderUser sets SenderUser field to given value.

### HasSenderUser

`func (o *OwnedGiftUnique) HasSenderUser() bool`

HasSenderUser returns a boolean if a field has been set.

### GetSendDate

`func (o *OwnedGiftUnique) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *OwnedGiftUnique) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *OwnedGiftUnique) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.


### GetIsSaved

`func (o *OwnedGiftUnique) GetIsSaved() bool`

GetIsSaved returns the IsSaved field if non-nil, zero value otherwise.

### GetIsSavedOk

`func (o *OwnedGiftUnique) GetIsSavedOk() (*bool, bool)`

GetIsSavedOk returns a tuple with the IsSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSaved

`func (o *OwnedGiftUnique) SetIsSaved(v bool)`

SetIsSaved sets IsSaved field to given value.

### HasIsSaved

`func (o *OwnedGiftUnique) HasIsSaved() bool`

HasIsSaved returns a boolean if a field has been set.

### GetCanBeTransferred

`func (o *OwnedGiftUnique) GetCanBeTransferred() bool`

GetCanBeTransferred returns the CanBeTransferred field if non-nil, zero value otherwise.

### GetCanBeTransferredOk

`func (o *OwnedGiftUnique) GetCanBeTransferredOk() (*bool, bool)`

GetCanBeTransferredOk returns a tuple with the CanBeTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeTransferred

`func (o *OwnedGiftUnique) SetCanBeTransferred(v bool)`

SetCanBeTransferred sets CanBeTransferred field to given value.

### HasCanBeTransferred

`func (o *OwnedGiftUnique) HasCanBeTransferred() bool`

HasCanBeTransferred returns a boolean if a field has been set.

### GetTransferStarCount

`func (o *OwnedGiftUnique) GetTransferStarCount() int32`

GetTransferStarCount returns the TransferStarCount field if non-nil, zero value otherwise.

### GetTransferStarCountOk

`func (o *OwnedGiftUnique) GetTransferStarCountOk() (*int32, bool)`

GetTransferStarCountOk returns a tuple with the TransferStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferStarCount

`func (o *OwnedGiftUnique) SetTransferStarCount(v int32)`

SetTransferStarCount sets TransferStarCount field to given value.

### HasTransferStarCount

`func (o *OwnedGiftUnique) HasTransferStarCount() bool`

HasTransferStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


