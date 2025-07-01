# MessageReactionCountUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**MessageId** | **int32** | Unique message identifier inside the chat | 
**Date** | **int32** | Date of the change in Unix time | 
**Reactions** | [**[]ReactionCount**](ReactionCount.md) | List of reactions that are present on the message | 

## Methods

### NewMessageReactionCountUpdated

`func NewMessageReactionCountUpdated(chat Chat, messageId int32, date int32, reactions []ReactionCount, ) *MessageReactionCountUpdated`

NewMessageReactionCountUpdated instantiates a new MessageReactionCountUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageReactionCountUpdatedWithDefaults

`func NewMessageReactionCountUpdatedWithDefaults() *MessageReactionCountUpdated`

NewMessageReactionCountUpdatedWithDefaults instantiates a new MessageReactionCountUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *MessageReactionCountUpdated) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MessageReactionCountUpdated) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MessageReactionCountUpdated) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetMessageId

`func (o *MessageReactionCountUpdated) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageReactionCountUpdated) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageReactionCountUpdated) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetDate

`func (o *MessageReactionCountUpdated) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageReactionCountUpdated) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageReactionCountUpdated) SetDate(v int32)`

SetDate sets Date field to given value.


### GetReactions

`func (o *MessageReactionCountUpdated) GetReactions() []ReactionCount`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *MessageReactionCountUpdated) GetReactionsOk() (*[]ReactionCount, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *MessageReactionCountUpdated) SetReactions(v []ReactionCount)`

SetReactions sets Reactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


