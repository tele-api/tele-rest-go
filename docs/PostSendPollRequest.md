# PostSendPollRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the message will be sent | [optional] 
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
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
**ReplyMarkup** | Pointer to [**PostSendMessageRequestReplyMarkup**](PostSendMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewPostSendPollRequest

`func NewPostSendPollRequest(chatId PostSendMessageRequestChatId, question string, options []InputPollOption, ) *PostSendPollRequest`

NewPostSendPollRequest instantiates a new PostSendPollRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSendPollRequestWithDefaults

`func NewPostSendPollRequestWithDefaults() *PostSendPollRequest`

NewPostSendPollRequestWithDefaults instantiates a new PostSendPollRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostSendPollRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostSendPollRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostSendPollRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *PostSendPollRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *PostSendPollRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostSendPollRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostSendPollRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *PostSendPollRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *PostSendPollRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *PostSendPollRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *PostSendPollRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetQuestion

`func (o *PostSendPollRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PostSendPollRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PostSendPollRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetQuestionParseMode

`func (o *PostSendPollRequest) GetQuestionParseMode() string`

GetQuestionParseMode returns the QuestionParseMode field if non-nil, zero value otherwise.

### GetQuestionParseModeOk

`func (o *PostSendPollRequest) GetQuestionParseModeOk() (*string, bool)`

GetQuestionParseModeOk returns a tuple with the QuestionParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionParseMode

`func (o *PostSendPollRequest) SetQuestionParseMode(v string)`

SetQuestionParseMode sets QuestionParseMode field to given value.

### HasQuestionParseMode

`func (o *PostSendPollRequest) HasQuestionParseMode() bool`

HasQuestionParseMode returns a boolean if a field has been set.

### GetQuestionEntities

`func (o *PostSendPollRequest) GetQuestionEntities() []MessageEntity`

GetQuestionEntities returns the QuestionEntities field if non-nil, zero value otherwise.

### GetQuestionEntitiesOk

`func (o *PostSendPollRequest) GetQuestionEntitiesOk() (*[]MessageEntity, bool)`

GetQuestionEntitiesOk returns a tuple with the QuestionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionEntities

`func (o *PostSendPollRequest) SetQuestionEntities(v []MessageEntity)`

SetQuestionEntities sets QuestionEntities field to given value.

### HasQuestionEntities

`func (o *PostSendPollRequest) HasQuestionEntities() bool`

HasQuestionEntities returns a boolean if a field has been set.

### GetOptions

`func (o *PostSendPollRequest) GetOptions() []InputPollOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PostSendPollRequest) GetOptionsOk() (*[]InputPollOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PostSendPollRequest) SetOptions(v []InputPollOption)`

SetOptions sets Options field to given value.


### GetIsAnonymous

`func (o *PostSendPollRequest) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *PostSendPollRequest) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *PostSendPollRequest) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *PostSendPollRequest) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetType

`func (o *PostSendPollRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostSendPollRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostSendPollRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostSendPollRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAllowsMultipleAnswers

`func (o *PostSendPollRequest) GetAllowsMultipleAnswers() bool`

GetAllowsMultipleAnswers returns the AllowsMultipleAnswers field if non-nil, zero value otherwise.

### GetAllowsMultipleAnswersOk

`func (o *PostSendPollRequest) GetAllowsMultipleAnswersOk() (*bool, bool)`

GetAllowsMultipleAnswersOk returns a tuple with the AllowsMultipleAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsMultipleAnswers

`func (o *PostSendPollRequest) SetAllowsMultipleAnswers(v bool)`

SetAllowsMultipleAnswers sets AllowsMultipleAnswers field to given value.

### HasAllowsMultipleAnswers

`func (o *PostSendPollRequest) HasAllowsMultipleAnswers() bool`

HasAllowsMultipleAnswers returns a boolean if a field has been set.

### GetCorrectOptionId

`func (o *PostSendPollRequest) GetCorrectOptionId() int32`

GetCorrectOptionId returns the CorrectOptionId field if non-nil, zero value otherwise.

### GetCorrectOptionIdOk

`func (o *PostSendPollRequest) GetCorrectOptionIdOk() (*int32, bool)`

GetCorrectOptionIdOk returns a tuple with the CorrectOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectOptionId

`func (o *PostSendPollRequest) SetCorrectOptionId(v int32)`

SetCorrectOptionId sets CorrectOptionId field to given value.

### HasCorrectOptionId

`func (o *PostSendPollRequest) HasCorrectOptionId() bool`

HasCorrectOptionId returns a boolean if a field has been set.

### GetExplanation

`func (o *PostSendPollRequest) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *PostSendPollRequest) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *PostSendPollRequest) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *PostSendPollRequest) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetExplanationParseMode

`func (o *PostSendPollRequest) GetExplanationParseMode() string`

GetExplanationParseMode returns the ExplanationParseMode field if non-nil, zero value otherwise.

### GetExplanationParseModeOk

`func (o *PostSendPollRequest) GetExplanationParseModeOk() (*string, bool)`

GetExplanationParseModeOk returns a tuple with the ExplanationParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationParseMode

`func (o *PostSendPollRequest) SetExplanationParseMode(v string)`

SetExplanationParseMode sets ExplanationParseMode field to given value.

### HasExplanationParseMode

`func (o *PostSendPollRequest) HasExplanationParseMode() bool`

HasExplanationParseMode returns a boolean if a field has been set.

### GetExplanationEntities

`func (o *PostSendPollRequest) GetExplanationEntities() []MessageEntity`

GetExplanationEntities returns the ExplanationEntities field if non-nil, zero value otherwise.

### GetExplanationEntitiesOk

`func (o *PostSendPollRequest) GetExplanationEntitiesOk() (*[]MessageEntity, bool)`

GetExplanationEntitiesOk returns a tuple with the ExplanationEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationEntities

`func (o *PostSendPollRequest) SetExplanationEntities(v []MessageEntity)`

SetExplanationEntities sets ExplanationEntities field to given value.

### HasExplanationEntities

`func (o *PostSendPollRequest) HasExplanationEntities() bool`

HasExplanationEntities returns a boolean if a field has been set.

### GetOpenPeriod

`func (o *PostSendPollRequest) GetOpenPeriod() int32`

GetOpenPeriod returns the OpenPeriod field if non-nil, zero value otherwise.

### GetOpenPeriodOk

`func (o *PostSendPollRequest) GetOpenPeriodOk() (*int32, bool)`

GetOpenPeriodOk returns a tuple with the OpenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPeriod

`func (o *PostSendPollRequest) SetOpenPeriod(v int32)`

SetOpenPeriod sets OpenPeriod field to given value.

### HasOpenPeriod

`func (o *PostSendPollRequest) HasOpenPeriod() bool`

HasOpenPeriod returns a boolean if a field has been set.

### GetCloseDate

`func (o *PostSendPollRequest) GetCloseDate() int32`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *PostSendPollRequest) GetCloseDateOk() (*int32, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *PostSendPollRequest) SetCloseDate(v int32)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *PostSendPollRequest) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetIsClosed

`func (o *PostSendPollRequest) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *PostSendPollRequest) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *PostSendPollRequest) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *PostSendPollRequest) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetDisableNotification

`func (o *PostSendPollRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *PostSendPollRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *PostSendPollRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *PostSendPollRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *PostSendPollRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *PostSendPollRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *PostSendPollRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *PostSendPollRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *PostSendPollRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *PostSendPollRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *PostSendPollRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *PostSendPollRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *PostSendPollRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *PostSendPollRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *PostSendPollRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *PostSendPollRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *PostSendPollRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *PostSendPollRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *PostSendPollRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *PostSendPollRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *PostSendPollRequest) GetReplyMarkup() PostSendMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *PostSendPollRequest) GetReplyMarkupOk() (*PostSendMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *PostSendPollRequest) SetReplyMarkup(v PostSendMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *PostSendPollRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


