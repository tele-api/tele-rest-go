# ChatBoostSourceGiveaway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Source of the boost, always “giveaway” | [default to "giveaway"]
**GiveawayMessageId** | **int32** | Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn&#39;t sent yet. | 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**PrizeStarCount** | Pointer to **int32** | *Optional*. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only | [optional] 
**IsUnclaimed** | Pointer to **bool** | *Optional*. *True*, if the giveaway was completed, but there was no user to win the prize | [optional] [default to true]

## Methods

### NewChatBoostSourceGiveaway

`func NewChatBoostSourceGiveaway(source string, giveawayMessageId int32, ) *ChatBoostSourceGiveaway`

NewChatBoostSourceGiveaway instantiates a new ChatBoostSourceGiveaway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatBoostSourceGiveawayWithDefaults

`func NewChatBoostSourceGiveawayWithDefaults() *ChatBoostSourceGiveaway`

NewChatBoostSourceGiveawayWithDefaults instantiates a new ChatBoostSourceGiveaway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ChatBoostSourceGiveaway) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChatBoostSourceGiveaway) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChatBoostSourceGiveaway) SetSource(v string)`

SetSource sets Source field to given value.


### GetGiveawayMessageId

`func (o *ChatBoostSourceGiveaway) GetGiveawayMessageId() int32`

GetGiveawayMessageId returns the GiveawayMessageId field if non-nil, zero value otherwise.

### GetGiveawayMessageIdOk

`func (o *ChatBoostSourceGiveaway) GetGiveawayMessageIdOk() (*int32, bool)`

GetGiveawayMessageIdOk returns a tuple with the GiveawayMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayMessageId

`func (o *ChatBoostSourceGiveaway) SetGiveawayMessageId(v int32)`

SetGiveawayMessageId sets GiveawayMessageId field to given value.


### GetUser

`func (o *ChatBoostSourceGiveaway) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatBoostSourceGiveaway) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatBoostSourceGiveaway) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatBoostSourceGiveaway) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPrizeStarCount

`func (o *ChatBoostSourceGiveaway) GetPrizeStarCount() int32`

GetPrizeStarCount returns the PrizeStarCount field if non-nil, zero value otherwise.

### GetPrizeStarCountOk

`func (o *ChatBoostSourceGiveaway) GetPrizeStarCountOk() (*int32, bool)`

GetPrizeStarCountOk returns a tuple with the PrizeStarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeStarCount

`func (o *ChatBoostSourceGiveaway) SetPrizeStarCount(v int32)`

SetPrizeStarCount sets PrizeStarCount field to given value.

### HasPrizeStarCount

`func (o *ChatBoostSourceGiveaway) HasPrizeStarCount() bool`

HasPrizeStarCount returns a boolean if a field has been set.

### GetIsUnclaimed

`func (o *ChatBoostSourceGiveaway) GetIsUnclaimed() bool`

GetIsUnclaimed returns the IsUnclaimed field if non-nil, zero value otherwise.

### GetIsUnclaimedOk

`func (o *ChatBoostSourceGiveaway) GetIsUnclaimedOk() (*bool, bool)`

GetIsUnclaimedOk returns a tuple with the IsUnclaimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnclaimed

`func (o *ChatBoostSourceGiveaway) SetIsUnclaimed(v bool)`

SetIsUnclaimed sets IsUnclaimed field to given value.

### HasIsUnclaimed

`func (o *ChatBoostSourceGiveaway) HasIsUnclaimed() bool`

HasIsUnclaimed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


