# SetMessageReactionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**MessageId** | **int32** | Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead. | 
**Reaction** | Pointer to [**[]ReactionType**](ReactionType.md) | A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can&#39;t be used by bots. | [optional] 
**IsBig** | Pointer to **bool** | Pass *True* to set the reaction with a big animation | [optional] 

## Methods

### NewSetMessageReactionPostRequest

`func NewSetMessageReactionPostRequest(chatId SendMessagePostRequestChatId, messageId int32, ) *SetMessageReactionPostRequest`

NewSetMessageReactionPostRequest instantiates a new SetMessageReactionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMessageReactionPostRequestWithDefaults

`func NewSetMessageReactionPostRequestWithDefaults() *SetMessageReactionPostRequest`

NewSetMessageReactionPostRequestWithDefaults instantiates a new SetMessageReactionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetMessageReactionPostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetMessageReactionPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetMessageReactionPostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageId

`func (o *SetMessageReactionPostRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SetMessageReactionPostRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SetMessageReactionPostRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetReaction

`func (o *SetMessageReactionPostRequest) GetReaction() []ReactionType`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *SetMessageReactionPostRequest) GetReactionOk() (*[]ReactionType, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *SetMessageReactionPostRequest) SetReaction(v []ReactionType)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *SetMessageReactionPostRequest) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetIsBig

`func (o *SetMessageReactionPostRequest) GetIsBig() bool`

GetIsBig returns the IsBig field if non-nil, zero value otherwise.

### GetIsBigOk

`func (o *SetMessageReactionPostRequest) GetIsBigOk() (*bool, bool)`

GetIsBigOk returns a tuple with the IsBig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBig

`func (o *SetMessageReactionPostRequest) SetIsBig(v bool)`

SetIsBig sets IsBig field to given value.

### HasIsBig

`func (o *SetMessageReactionPostRequest) HasIsBig() bool`

HasIsBig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


