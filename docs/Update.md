# Update

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateId** | **int32** | The update&#39;s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This identifier becomes especially handy if you&#39;re using [webhooks](https://core.telegram.org/bots/api/#setwebhook), since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially. | 
**Message** | Pointer to [**Message**](Message.md) |  | [optional] 
**EditedMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**ChannelPost** | Pointer to [**Message**](Message.md) |  | [optional] 
**EditedChannelPost** | Pointer to [**Message**](Message.md) |  | [optional] 
**BusinessConnection** | Pointer to [**BusinessConnection**](BusinessConnection.md) |  | [optional] 
**BusinessMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**EditedBusinessMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**DeletedBusinessMessages** | Pointer to [**BusinessMessagesDeleted**](BusinessMessagesDeleted.md) |  | [optional] 
**MessageReaction** | Pointer to [**MessageReactionUpdated**](MessageReactionUpdated.md) |  | [optional] 
**MessageReactionCount** | Pointer to [**MessageReactionCountUpdated**](MessageReactionCountUpdated.md) |  | [optional] 
**InlineQuery** | Pointer to [**InlineQuery**](InlineQuery.md) |  | [optional] 
**ChosenInlineResult** | Pointer to [**ChosenInlineResult**](ChosenInlineResult.md) |  | [optional] 
**CallbackQuery** | Pointer to [**CallbackQuery**](CallbackQuery.md) |  | [optional] 
**ShippingQuery** | Pointer to [**ShippingQuery**](ShippingQuery.md) |  | [optional] 
**PreCheckoutQuery** | Pointer to [**PreCheckoutQuery**](PreCheckoutQuery.md) |  | [optional] 
**PurchasedPaidMedia** | Pointer to [**PaidMediaPurchased**](PaidMediaPurchased.md) |  | [optional] 
**Poll** | Pointer to [**Poll**](Poll.md) |  | [optional] 
**PollAnswer** | Pointer to [**PollAnswer**](PollAnswer.md) |  | [optional] 
**MyChatMember** | Pointer to [**ChatMemberUpdated**](ChatMemberUpdated.md) |  | [optional] 
**ChatMember** | Pointer to [**ChatMemberUpdated**](ChatMemberUpdated.md) |  | [optional] 
**ChatJoinRequest** | Pointer to [**ChatJoinRequest**](ChatJoinRequest.md) |  | [optional] 
**ChatBoost** | Pointer to [**ChatBoostUpdated**](ChatBoostUpdated.md) |  | [optional] 
**RemovedChatBoost** | Pointer to [**ChatBoostRemoved**](ChatBoostRemoved.md) |  | [optional] 

## Methods

### NewUpdate

`func NewUpdate(updateId int32, ) *Update`

NewUpdate instantiates a new Update object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWithDefaults

`func NewUpdateWithDefaults() *Update`

NewUpdateWithDefaults instantiates a new Update object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateId

`func (o *Update) GetUpdateId() int32`

GetUpdateId returns the UpdateId field if non-nil, zero value otherwise.

### GetUpdateIdOk

`func (o *Update) GetUpdateIdOk() (*int32, bool)`

GetUpdateIdOk returns a tuple with the UpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateId

`func (o *Update) SetUpdateId(v int32)`

SetUpdateId sets UpdateId field to given value.


### GetMessage

`func (o *Update) GetMessage() Message`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Update) GetMessageOk() (*Message, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Update) SetMessage(v Message)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Update) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEditedMessage

`func (o *Update) GetEditedMessage() Message`

GetEditedMessage returns the EditedMessage field if non-nil, zero value otherwise.

### GetEditedMessageOk

`func (o *Update) GetEditedMessageOk() (*Message, bool)`

GetEditedMessageOk returns a tuple with the EditedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedMessage

`func (o *Update) SetEditedMessage(v Message)`

SetEditedMessage sets EditedMessage field to given value.

### HasEditedMessage

`func (o *Update) HasEditedMessage() bool`

HasEditedMessage returns a boolean if a field has been set.

### GetChannelPost

`func (o *Update) GetChannelPost() Message`

GetChannelPost returns the ChannelPost field if non-nil, zero value otherwise.

### GetChannelPostOk

`func (o *Update) GetChannelPostOk() (*Message, bool)`

GetChannelPostOk returns a tuple with the ChannelPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPost

`func (o *Update) SetChannelPost(v Message)`

SetChannelPost sets ChannelPost field to given value.

### HasChannelPost

`func (o *Update) HasChannelPost() bool`

HasChannelPost returns a boolean if a field has been set.

### GetEditedChannelPost

`func (o *Update) GetEditedChannelPost() Message`

GetEditedChannelPost returns the EditedChannelPost field if non-nil, zero value otherwise.

### GetEditedChannelPostOk

`func (o *Update) GetEditedChannelPostOk() (*Message, bool)`

GetEditedChannelPostOk returns a tuple with the EditedChannelPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedChannelPost

`func (o *Update) SetEditedChannelPost(v Message)`

SetEditedChannelPost sets EditedChannelPost field to given value.

### HasEditedChannelPost

`func (o *Update) HasEditedChannelPost() bool`

HasEditedChannelPost returns a boolean if a field has been set.

### GetBusinessConnection

`func (o *Update) GetBusinessConnection() BusinessConnection`

GetBusinessConnection returns the BusinessConnection field if non-nil, zero value otherwise.

### GetBusinessConnectionOk

`func (o *Update) GetBusinessConnectionOk() (*BusinessConnection, bool)`

GetBusinessConnectionOk returns a tuple with the BusinessConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnection

`func (o *Update) SetBusinessConnection(v BusinessConnection)`

SetBusinessConnection sets BusinessConnection field to given value.

### HasBusinessConnection

`func (o *Update) HasBusinessConnection() bool`

HasBusinessConnection returns a boolean if a field has been set.

### GetBusinessMessage

`func (o *Update) GetBusinessMessage() Message`

GetBusinessMessage returns the BusinessMessage field if non-nil, zero value otherwise.

### GetBusinessMessageOk

`func (o *Update) GetBusinessMessageOk() (*Message, bool)`

GetBusinessMessageOk returns a tuple with the BusinessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessMessage

`func (o *Update) SetBusinessMessage(v Message)`

SetBusinessMessage sets BusinessMessage field to given value.

### HasBusinessMessage

`func (o *Update) HasBusinessMessage() bool`

HasBusinessMessage returns a boolean if a field has been set.

### GetEditedBusinessMessage

`func (o *Update) GetEditedBusinessMessage() Message`

GetEditedBusinessMessage returns the EditedBusinessMessage field if non-nil, zero value otherwise.

### GetEditedBusinessMessageOk

`func (o *Update) GetEditedBusinessMessageOk() (*Message, bool)`

GetEditedBusinessMessageOk returns a tuple with the EditedBusinessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedBusinessMessage

`func (o *Update) SetEditedBusinessMessage(v Message)`

SetEditedBusinessMessage sets EditedBusinessMessage field to given value.

### HasEditedBusinessMessage

`func (o *Update) HasEditedBusinessMessage() bool`

HasEditedBusinessMessage returns a boolean if a field has been set.

### GetDeletedBusinessMessages

`func (o *Update) GetDeletedBusinessMessages() BusinessMessagesDeleted`

GetDeletedBusinessMessages returns the DeletedBusinessMessages field if non-nil, zero value otherwise.

### GetDeletedBusinessMessagesOk

`func (o *Update) GetDeletedBusinessMessagesOk() (*BusinessMessagesDeleted, bool)`

GetDeletedBusinessMessagesOk returns a tuple with the DeletedBusinessMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBusinessMessages

`func (o *Update) SetDeletedBusinessMessages(v BusinessMessagesDeleted)`

SetDeletedBusinessMessages sets DeletedBusinessMessages field to given value.

### HasDeletedBusinessMessages

`func (o *Update) HasDeletedBusinessMessages() bool`

HasDeletedBusinessMessages returns a boolean if a field has been set.

### GetMessageReaction

`func (o *Update) GetMessageReaction() MessageReactionUpdated`

GetMessageReaction returns the MessageReaction field if non-nil, zero value otherwise.

### GetMessageReactionOk

`func (o *Update) GetMessageReactionOk() (*MessageReactionUpdated, bool)`

GetMessageReactionOk returns a tuple with the MessageReaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReaction

`func (o *Update) SetMessageReaction(v MessageReactionUpdated)`

SetMessageReaction sets MessageReaction field to given value.

### HasMessageReaction

`func (o *Update) HasMessageReaction() bool`

HasMessageReaction returns a boolean if a field has been set.

### GetMessageReactionCount

`func (o *Update) GetMessageReactionCount() MessageReactionCountUpdated`

GetMessageReactionCount returns the MessageReactionCount field if non-nil, zero value otherwise.

### GetMessageReactionCountOk

`func (o *Update) GetMessageReactionCountOk() (*MessageReactionCountUpdated, bool)`

GetMessageReactionCountOk returns a tuple with the MessageReactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReactionCount

`func (o *Update) SetMessageReactionCount(v MessageReactionCountUpdated)`

SetMessageReactionCount sets MessageReactionCount field to given value.

### HasMessageReactionCount

`func (o *Update) HasMessageReactionCount() bool`

HasMessageReactionCount returns a boolean if a field has been set.

### GetInlineQuery

`func (o *Update) GetInlineQuery() InlineQuery`

GetInlineQuery returns the InlineQuery field if non-nil, zero value otherwise.

### GetInlineQueryOk

`func (o *Update) GetInlineQueryOk() (*InlineQuery, bool)`

GetInlineQueryOk returns a tuple with the InlineQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineQuery

`func (o *Update) SetInlineQuery(v InlineQuery)`

SetInlineQuery sets InlineQuery field to given value.

### HasInlineQuery

`func (o *Update) HasInlineQuery() bool`

HasInlineQuery returns a boolean if a field has been set.

### GetChosenInlineResult

`func (o *Update) GetChosenInlineResult() ChosenInlineResult`

GetChosenInlineResult returns the ChosenInlineResult field if non-nil, zero value otherwise.

### GetChosenInlineResultOk

`func (o *Update) GetChosenInlineResultOk() (*ChosenInlineResult, bool)`

GetChosenInlineResultOk returns a tuple with the ChosenInlineResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenInlineResult

`func (o *Update) SetChosenInlineResult(v ChosenInlineResult)`

SetChosenInlineResult sets ChosenInlineResult field to given value.

### HasChosenInlineResult

`func (o *Update) HasChosenInlineResult() bool`

HasChosenInlineResult returns a boolean if a field has been set.

### GetCallbackQuery

`func (o *Update) GetCallbackQuery() CallbackQuery`

GetCallbackQuery returns the CallbackQuery field if non-nil, zero value otherwise.

### GetCallbackQueryOk

`func (o *Update) GetCallbackQueryOk() (*CallbackQuery, bool)`

GetCallbackQueryOk returns a tuple with the CallbackQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackQuery

`func (o *Update) SetCallbackQuery(v CallbackQuery)`

SetCallbackQuery sets CallbackQuery field to given value.

### HasCallbackQuery

`func (o *Update) HasCallbackQuery() bool`

HasCallbackQuery returns a boolean if a field has been set.

### GetShippingQuery

`func (o *Update) GetShippingQuery() ShippingQuery`

GetShippingQuery returns the ShippingQuery field if non-nil, zero value otherwise.

### GetShippingQueryOk

`func (o *Update) GetShippingQueryOk() (*ShippingQuery, bool)`

GetShippingQueryOk returns a tuple with the ShippingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingQuery

`func (o *Update) SetShippingQuery(v ShippingQuery)`

SetShippingQuery sets ShippingQuery field to given value.

### HasShippingQuery

`func (o *Update) HasShippingQuery() bool`

HasShippingQuery returns a boolean if a field has been set.

### GetPreCheckoutQuery

`func (o *Update) GetPreCheckoutQuery() PreCheckoutQuery`

GetPreCheckoutQuery returns the PreCheckoutQuery field if non-nil, zero value otherwise.

### GetPreCheckoutQueryOk

`func (o *Update) GetPreCheckoutQueryOk() (*PreCheckoutQuery, bool)`

GetPreCheckoutQueryOk returns a tuple with the PreCheckoutQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCheckoutQuery

`func (o *Update) SetPreCheckoutQuery(v PreCheckoutQuery)`

SetPreCheckoutQuery sets PreCheckoutQuery field to given value.

### HasPreCheckoutQuery

`func (o *Update) HasPreCheckoutQuery() bool`

HasPreCheckoutQuery returns a boolean if a field has been set.

### GetPurchasedPaidMedia

`func (o *Update) GetPurchasedPaidMedia() PaidMediaPurchased`

GetPurchasedPaidMedia returns the PurchasedPaidMedia field if non-nil, zero value otherwise.

### GetPurchasedPaidMediaOk

`func (o *Update) GetPurchasedPaidMediaOk() (*PaidMediaPurchased, bool)`

GetPurchasedPaidMediaOk returns a tuple with the PurchasedPaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedPaidMedia

`func (o *Update) SetPurchasedPaidMedia(v PaidMediaPurchased)`

SetPurchasedPaidMedia sets PurchasedPaidMedia field to given value.

### HasPurchasedPaidMedia

`func (o *Update) HasPurchasedPaidMedia() bool`

HasPurchasedPaidMedia returns a boolean if a field has been set.

### GetPoll

`func (o *Update) GetPoll() Poll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *Update) GetPollOk() (*Poll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *Update) SetPoll(v Poll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *Update) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetPollAnswer

`func (o *Update) GetPollAnswer() PollAnswer`

GetPollAnswer returns the PollAnswer field if non-nil, zero value otherwise.

### GetPollAnswerOk

`func (o *Update) GetPollAnswerOk() (*PollAnswer, bool)`

GetPollAnswerOk returns a tuple with the PollAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollAnswer

`func (o *Update) SetPollAnswer(v PollAnswer)`

SetPollAnswer sets PollAnswer field to given value.

### HasPollAnswer

`func (o *Update) HasPollAnswer() bool`

HasPollAnswer returns a boolean if a field has been set.

### GetMyChatMember

`func (o *Update) GetMyChatMember() ChatMemberUpdated`

GetMyChatMember returns the MyChatMember field if non-nil, zero value otherwise.

### GetMyChatMemberOk

`func (o *Update) GetMyChatMemberOk() (*ChatMemberUpdated, bool)`

GetMyChatMemberOk returns a tuple with the MyChatMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyChatMember

`func (o *Update) SetMyChatMember(v ChatMemberUpdated)`

SetMyChatMember sets MyChatMember field to given value.

### HasMyChatMember

`func (o *Update) HasMyChatMember() bool`

HasMyChatMember returns a boolean if a field has been set.

### GetChatMember

`func (o *Update) GetChatMember() ChatMemberUpdated`

GetChatMember returns the ChatMember field if non-nil, zero value otherwise.

### GetChatMemberOk

`func (o *Update) GetChatMemberOk() (*ChatMemberUpdated, bool)`

GetChatMemberOk returns a tuple with the ChatMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatMember

`func (o *Update) SetChatMember(v ChatMemberUpdated)`

SetChatMember sets ChatMember field to given value.

### HasChatMember

`func (o *Update) HasChatMember() bool`

HasChatMember returns a boolean if a field has been set.

### GetChatJoinRequest

`func (o *Update) GetChatJoinRequest() ChatJoinRequest`

GetChatJoinRequest returns the ChatJoinRequest field if non-nil, zero value otherwise.

### GetChatJoinRequestOk

`func (o *Update) GetChatJoinRequestOk() (*ChatJoinRequest, bool)`

GetChatJoinRequestOk returns a tuple with the ChatJoinRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatJoinRequest

`func (o *Update) SetChatJoinRequest(v ChatJoinRequest)`

SetChatJoinRequest sets ChatJoinRequest field to given value.

### HasChatJoinRequest

`func (o *Update) HasChatJoinRequest() bool`

HasChatJoinRequest returns a boolean if a field has been set.

### GetChatBoost

`func (o *Update) GetChatBoost() ChatBoostUpdated`

GetChatBoost returns the ChatBoost field if non-nil, zero value otherwise.

### GetChatBoostOk

`func (o *Update) GetChatBoostOk() (*ChatBoostUpdated, bool)`

GetChatBoostOk returns a tuple with the ChatBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatBoost

`func (o *Update) SetChatBoost(v ChatBoostUpdated)`

SetChatBoost sets ChatBoost field to given value.

### HasChatBoost

`func (o *Update) HasChatBoost() bool`

HasChatBoost returns a boolean if a field has been set.

### GetRemovedChatBoost

`func (o *Update) GetRemovedChatBoost() ChatBoostRemoved`

GetRemovedChatBoost returns the RemovedChatBoost field if non-nil, zero value otherwise.

### GetRemovedChatBoostOk

`func (o *Update) GetRemovedChatBoostOk() (*ChatBoostRemoved, bool)`

GetRemovedChatBoostOk returns a tuple with the RemovedChatBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedChatBoost

`func (o *Update) SetRemovedChatBoost(v ChatBoostRemoved)`

SetRemovedChatBoost sets RemovedChatBoost field to given value.

### HasRemovedChatBoost

`func (o *Update) HasRemovedChatBoost() bool`

HasRemovedChatBoost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


