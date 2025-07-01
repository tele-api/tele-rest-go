# InlineQueryResultAudio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *audio* | [default to "audio"]
**Id** | **string** | Unique identifier for this result, 1-64 bytes | 
**AudioUrl** | **string** | A valid URL for the audio file | 
**Title** | **string** | Title | 
**Caption** | Pointer to **string** | *Optional*. Caption, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the audio caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**Performer** | Pointer to **string** | *Optional*. Performer | [optional] 
**AudioDuration** | Pointer to **int32** | *Optional*. Audio duration in seconds | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 
**InputMessageContent** | Pointer to [**InputMessageContent**](InputMessageContent.md) |  | [optional] 

## Methods

### NewInlineQueryResultAudio

`func NewInlineQueryResultAudio(type_ string, id string, audioUrl string, title string, ) *InlineQueryResultAudio`

NewInlineQueryResultAudio instantiates a new InlineQueryResultAudio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultAudioWithDefaults

`func NewInlineQueryResultAudioWithDefaults() *InlineQueryResultAudio`

NewInlineQueryResultAudioWithDefaults instantiates a new InlineQueryResultAudio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineQueryResultAudio) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineQueryResultAudio) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineQueryResultAudio) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InlineQueryResultAudio) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineQueryResultAudio) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineQueryResultAudio) SetId(v string)`

SetId sets Id field to given value.


### GetAudioUrl

`func (o *InlineQueryResultAudio) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *InlineQueryResultAudio) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *InlineQueryResultAudio) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.


### GetTitle

`func (o *InlineQueryResultAudio) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineQueryResultAudio) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineQueryResultAudio) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCaption

`func (o *InlineQueryResultAudio) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InlineQueryResultAudio) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InlineQueryResultAudio) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InlineQueryResultAudio) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InlineQueryResultAudio) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InlineQueryResultAudio) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InlineQueryResultAudio) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InlineQueryResultAudio) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InlineQueryResultAudio) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InlineQueryResultAudio) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InlineQueryResultAudio) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InlineQueryResultAudio) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetPerformer

`func (o *InlineQueryResultAudio) GetPerformer() string`

GetPerformer returns the Performer field if non-nil, zero value otherwise.

### GetPerformerOk

`func (o *InlineQueryResultAudio) GetPerformerOk() (*string, bool)`

GetPerformerOk returns a tuple with the Performer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformer

`func (o *InlineQueryResultAudio) SetPerformer(v string)`

SetPerformer sets Performer field to given value.

### HasPerformer

`func (o *InlineQueryResultAudio) HasPerformer() bool`

HasPerformer returns a boolean if a field has been set.

### GetAudioDuration

`func (o *InlineQueryResultAudio) GetAudioDuration() int32`

GetAudioDuration returns the AudioDuration field if non-nil, zero value otherwise.

### GetAudioDurationOk

`func (o *InlineQueryResultAudio) GetAudioDurationOk() (*int32, bool)`

GetAudioDurationOk returns a tuple with the AudioDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioDuration

`func (o *InlineQueryResultAudio) SetAudioDuration(v int32)`

SetAudioDuration sets AudioDuration field to given value.

### HasAudioDuration

`func (o *InlineQueryResultAudio) HasAudioDuration() bool`

HasAudioDuration returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *InlineQueryResultAudio) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *InlineQueryResultAudio) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *InlineQueryResultAudio) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *InlineQueryResultAudio) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetInputMessageContent

`func (o *InlineQueryResultAudio) GetInputMessageContent() InputMessageContent`

GetInputMessageContent returns the InputMessageContent field if non-nil, zero value otherwise.

### GetInputMessageContentOk

`func (o *InlineQueryResultAudio) GetInputMessageContentOk() (*InputMessageContent, bool)`

GetInputMessageContentOk returns a tuple with the InputMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageContent

`func (o *InlineQueryResultAudio) SetInputMessageContent(v InputMessageContent)`

SetInputMessageContent sets InputMessageContent field to given value.

### HasInputMessageContent

`func (o *InlineQueryResultAudio) HasInputMessageContent() bool`

HasInputMessageContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


