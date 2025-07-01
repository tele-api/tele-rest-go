# MessageReactionUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**MessageId** | **int32** | Unique identifier of the message inside the chat | 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**ActorChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**Date** | **int32** | Date of the change in Unix time | 
**OldReaction** | [**[]ReactionType**](ReactionType.md) | Previous list of reaction types that were set by the user | 
**NewReaction** | [**[]ReactionType**](ReactionType.md) | New list of reaction types that have been set by the user | 

## Methods

### NewMessageReactionUpdated

`func NewMessageReactionUpdated(chat Chat, messageId int32, date int32, oldReaction []ReactionType, newReaction []ReactionType, ) *MessageReactionUpdated`

NewMessageReactionUpdated instantiates a new MessageReactionUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageReactionUpdatedWithDefaults

`func NewMessageReactionUpdatedWithDefaults() *MessageReactionUpdated`

NewMessageReactionUpdatedWithDefaults instantiates a new MessageReactionUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *MessageReactionUpdated) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MessageReactionUpdated) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MessageReactionUpdated) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetMessageId

`func (o *MessageReactionUpdated) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageReactionUpdated) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageReactionUpdated) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetUser

`func (o *MessageReactionUpdated) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MessageReactionUpdated) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MessageReactionUpdated) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *MessageReactionUpdated) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetActorChat

`func (o *MessageReactionUpdated) GetActorChat() Chat`

GetActorChat returns the ActorChat field if non-nil, zero value otherwise.

### GetActorChatOk

`func (o *MessageReactionUpdated) GetActorChatOk() (*Chat, bool)`

GetActorChatOk returns a tuple with the ActorChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorChat

`func (o *MessageReactionUpdated) SetActorChat(v Chat)`

SetActorChat sets ActorChat field to given value.

### HasActorChat

`func (o *MessageReactionUpdated) HasActorChat() bool`

HasActorChat returns a boolean if a field has been set.

### GetDate

`func (o *MessageReactionUpdated) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageReactionUpdated) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageReactionUpdated) SetDate(v int32)`

SetDate sets Date field to given value.


### GetOldReaction

`func (o *MessageReactionUpdated) GetOldReaction() []ReactionType`

GetOldReaction returns the OldReaction field if non-nil, zero value otherwise.

### GetOldReactionOk

`func (o *MessageReactionUpdated) GetOldReactionOk() (*[]ReactionType, bool)`

GetOldReactionOk returns a tuple with the OldReaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldReaction

`func (o *MessageReactionUpdated) SetOldReaction(v []ReactionType)`

SetOldReaction sets OldReaction field to given value.


### GetNewReaction

`func (o *MessageReactionUpdated) GetNewReaction() []ReactionType`

GetNewReaction returns the NewReaction field if non-nil, zero value otherwise.

### GetNewReactionOk

`func (o *MessageReactionUpdated) GetNewReactionOk() (*[]ReactionType, bool)`

GetNewReactionOk returns a tuple with the NewReaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewReaction

`func (o *MessageReactionUpdated) SetNewReaction(v []ReactionType)`

SetNewReaction sets NewReaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


