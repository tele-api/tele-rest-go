# SendPollPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Question** | **string** | Poll question, 1-300 characters | 
**QuestionParseMode** | Pointer to **string** | Mode for parsing entities in the question. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. Currently, only custom emoji entities are allowed | [optional] 
**QuestionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of *question\\_parse\\_mode* | [optional] 
**Options** | [**[]InputPollOption**](InputPollOption.md) | A JSON-serialized list of 2-10 answer options | 
**IsAnonymous** | Pointer to **bool** | *True*, if the poll needs to be anonymous, defaults to *True* | [optional] 
**Type** | Pointer to **string** | Poll type, “quiz” or “regular”, defaults to “regular” | [optional] 
**AllowsMultipleAnswers** | Pointer to **bool** | *True*, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to *False* | [optional] 
**CorrectOptionId** | Pointer to **int32** | 0-based identifier of the correct answer option, required for polls in quiz mode | [optional] 
**Explanation** | Pointer to **string** | Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing | [optional] 
**ExplanationParseMode** | Pointer to **string** | Mode for parsing entities in the explanation. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**ExplanationEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of *explanation\\_parse\\_mode* | [optional] 
**OpenPeriod** | Pointer to **int32** | Amount of time in seconds the poll will be active after creation, 5-600. Can&#39;t be used together with *close\\_date*. | [optional] 
**CloseDate** | Pointer to **int32** | Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can&#39;t be used together with *open\\_period*. | [optional] 
**IsClosed** | Pointer to **bool** | Pass *True* if the poll needs to be immediately closed. This can be useful for poll preview. | [optional] 
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendMessagePostRequestReplyMarkup**](SendMessagePostRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewSendPollPostRequest

`func NewSendPollPostRequest(chatId SendMessagePostRequestChatId, question string, options []InputPollOption, ) *SendPollPostRequest`

NewSendPollPostRequest instantiates a new SendPollPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPollPostRequestWithDefaults

`func NewSendPollPostRequestWithDefaults() *SendPollPostRequest`

NewSendPollPostRequestWithDefaults instantiates a new SendPollPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SendPollPostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SendPollPostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SendPollPostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *SendPollPostRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *SendPollPostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendPollPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendPollPostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendPollPostRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendPollPostRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendPollPostRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendPollPostRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetQuestion

`func (o *SendPollPostRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SendPollPostRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SendPollPostRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetQuestionParseMode

`func (o *SendPollPostRequest) GetQuestionParseMode() string`

GetQuestionParseMode returns the QuestionParseMode field if non-nil, zero value otherwise.

### GetQuestionParseModeOk

`func (o *SendPollPostRequest) GetQuestionParseModeOk() (*string, bool)`

GetQuestionParseModeOk returns a tuple with the QuestionParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionParseMode

`func (o *SendPollPostRequest) SetQuestionParseMode(v string)`

SetQuestionParseMode sets QuestionParseMode field to given value.

### HasQuestionParseMode

`func (o *SendPollPostRequest) HasQuestionParseMode() bool`

HasQuestionParseMode returns a boolean if a field has been set.

### GetQuestionEntities

`func (o *SendPollPostRequest) GetQuestionEntities() []MessageEntity`

GetQuestionEntities returns the QuestionEntities field if non-nil, zero value otherwise.

### GetQuestionEntitiesOk

`func (o *SendPollPostRequest) GetQuestionEntitiesOk() (*[]MessageEntity, bool)`

GetQuestionEntitiesOk returns a tuple with the QuestionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionEntities

`func (o *SendPollPostRequest) SetQuestionEntities(v []MessageEntity)`

SetQuestionEntities sets QuestionEntities field to given value.

### HasQuestionEntities

`func (o *SendPollPostRequest) HasQuestionEntities() bool`

HasQuestionEntities returns a boolean if a field has been set.

### GetOptions

`func (o *SendPollPostRequest) GetOptions() []InputPollOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SendPollPostRequest) GetOptionsOk() (*[]InputPollOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SendPollPostRequest) SetOptions(v []InputPollOption)`

SetOptions sets Options field to given value.


### GetIsAnonymous

`func (o *SendPollPostRequest) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *SendPollPostRequest) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *SendPollPostRequest) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *SendPollPostRequest) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetType

`func (o *SendPollPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendPollPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendPollPostRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SendPollPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAllowsMultipleAnswers

`func (o *SendPollPostRequest) GetAllowsMultipleAnswers() bool`

GetAllowsMultipleAnswers returns the AllowsMultipleAnswers field if non-nil, zero value otherwise.

### GetAllowsMultipleAnswersOk

`func (o *SendPollPostRequest) GetAllowsMultipleAnswersOk() (*bool, bool)`

GetAllowsMultipleAnswersOk returns a tuple with the AllowsMultipleAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsMultipleAnswers

`func (o *SendPollPostRequest) SetAllowsMultipleAnswers(v bool)`

SetAllowsMultipleAnswers sets AllowsMultipleAnswers field to given value.

### HasAllowsMultipleAnswers

`func (o *SendPollPostRequest) HasAllowsMultipleAnswers() bool`

HasAllowsMultipleAnswers returns a boolean if a field has been set.

### GetCorrectOptionId

`func (o *SendPollPostRequest) GetCorrectOptionId() int32`

GetCorrectOptionId returns the CorrectOptionId field if non-nil, zero value otherwise.

### GetCorrectOptionIdOk

`func (o *SendPollPostRequest) GetCorrectOptionIdOk() (*int32, bool)`

GetCorrectOptionIdOk returns a tuple with the CorrectOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectOptionId

`func (o *SendPollPostRequest) SetCorrectOptionId(v int32)`

SetCorrectOptionId sets CorrectOptionId field to given value.

### HasCorrectOptionId

`func (o *SendPollPostRequest) HasCorrectOptionId() bool`

HasCorrectOptionId returns a boolean if a field has been set.

### GetExplanation

`func (o *SendPollPostRequest) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SendPollPostRequest) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SendPollPostRequest) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SendPollPostRequest) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetExplanationParseMode

`func (o *SendPollPostRequest) GetExplanationParseMode() string`

GetExplanationParseMode returns the ExplanationParseMode field if non-nil, zero value otherwise.

### GetExplanationParseModeOk

`func (o *SendPollPostRequest) GetExplanationParseModeOk() (*string, bool)`

GetExplanationParseModeOk returns a tuple with the ExplanationParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationParseMode

`func (o *SendPollPostRequest) SetExplanationParseMode(v string)`

SetExplanationParseMode sets ExplanationParseMode field to given value.

### HasExplanationParseMode

`func (o *SendPollPostRequest) HasExplanationParseMode() bool`

HasExplanationParseMode returns a boolean if a field has been set.

### GetExplanationEntities

`func (o *SendPollPostRequest) GetExplanationEntities() []MessageEntity`

GetExplanationEntities returns the ExplanationEntities field if non-nil, zero value otherwise.

### GetExplanationEntitiesOk

`func (o *SendPollPostRequest) GetExplanationEntitiesOk() (*[]MessageEntity, bool)`

GetExplanationEntitiesOk returns a tuple with the ExplanationEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationEntities

`func (o *SendPollPostRequest) SetExplanationEntities(v []MessageEntity)`

SetExplanationEntities sets ExplanationEntities field to given value.

### HasExplanationEntities

`func (o *SendPollPostRequest) HasExplanationEntities() bool`

HasExplanationEntities returns a boolean if a field has been set.

### GetOpenPeriod

`func (o *SendPollPostRequest) GetOpenPeriod() int32`

GetOpenPeriod returns the OpenPeriod field if non-nil, zero value otherwise.

### GetOpenPeriodOk

`func (o *SendPollPostRequest) GetOpenPeriodOk() (*int32, bool)`

GetOpenPeriodOk returns a tuple with the OpenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPeriod

`func (o *SendPollPostRequest) SetOpenPeriod(v int32)`

SetOpenPeriod sets OpenPeriod field to given value.

### HasOpenPeriod

`func (o *SendPollPostRequest) HasOpenPeriod() bool`

HasOpenPeriod returns a boolean if a field has been set.

### GetCloseDate

`func (o *SendPollPostRequest) GetCloseDate() int32`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *SendPollPostRequest) GetCloseDateOk() (*int32, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *SendPollPostRequest) SetCloseDate(v int32)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *SendPollPostRequest) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetIsClosed

`func (o *SendPollPostRequest) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *SendPollPostRequest) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *SendPollPostRequest) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *SendPollPostRequest) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendPollPostRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendPollPostRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendPollPostRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendPollPostRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendPollPostRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendPollPostRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendPollPostRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendPollPostRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendPollPostRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendPollPostRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendPollPostRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendPollPostRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendPollPostRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendPollPostRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendPollPostRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendPollPostRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendPollPostRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendPollPostRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendPollPostRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendPollPostRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendPollPostRequest) GetReplyMarkup() SendMessagePostRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendPollPostRequest) GetReplyMarkupOk() (*SendMessagePostRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendPollPostRequest) SetReplyMarkup(v SendMessagePostRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendPollPostRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


