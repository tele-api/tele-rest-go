# InlineQueryResultVoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *voice* | [default to "voice"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**VoiceUrl** | **string** | A valid URL for the voice recording | 
**Title** | **string** | Recording title | 
**Caption** | Pointer to **string** | *Optional*. Caption, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the voice message caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**VoiceDuration** | Pointer to **int32** | *Optional*. Recording duration in seconds | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultVoice

`func NewInlineQueryResultVoice(type_ string, id string, voiceUrl string, title string, ) *InlineQueryResultVoice`

NewInlineQueryResultVoice instantiates a new InlineQueryResultVoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultVoiceWithDefaults

`func NewInlineQueryResultVoiceWithDefaults() *InlineQueryResultVoice`

NewInlineQueryResultVoiceWithDefaults instantiates a new InlineQueryResultVoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultVoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultVoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultVoice) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultVoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultVoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultVoice) SetId(v string)`

SetId sets Id field to given value.


### GetVoiceUrl

`func (o *InlineQueryResultVoice) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *InlineQueryResultVoice) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *InlineQueryResultVoice) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.


### GetTitle

`func (o *InlineQueryResultVoice) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultVoice) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultVoice) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCaption

`func (o *InlineQueryResultVoice) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultVoice) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultVoice) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultVoice) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultVoice) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultVoice) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultVoice) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultVoice) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultVoice) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultVoice) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultVoice) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultVoice) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetVoiceDuration

`func (o *InlineQueryResultVoice) GetVoiceDuration() int32`

GetVoiceDuration returns the VoiceDuration field if non-nil, zero value otherwise.

### GetVoiceDurationOk

`func (o *InlineQueryResultVoice) GetVoiceDurationOk() (*int32, bool)`

GetVoiceDurationOk returns a tuple with the VoiceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDuration

`func (o *InlineQueryResultVoice) SetVoiceDuration(v int32)`

SetVoiceDuration sets VoiceDuration field to given value.

### HasVoiceDuration

`func (o *InlineQueryResultVoice) HasVoiceDuration() bool`

HasVoiceDuration returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultVoice) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultVoice) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultVoice) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultVoice) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultVoice) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultVoice) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultVoice) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultVoice) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


