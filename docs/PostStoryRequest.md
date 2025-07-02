# PostStoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**Content** | [**InputStoryContent**](InputStoryContent.md) |  | 
**ActivePeriod** | **int32** | Period after which the story is moved to the archive, in seconds; must be one of &#x60;6 * 3600&#x60;, &#x60;12 * 3600&#x60;, &#x60;86400&#x60;, or &#x60;2 * 86400&#x60; | 
**Caption** | Pointer to **string** | Caption of the story, 0-2048 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the story caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**Areas** | Pointer to [**[]StoryArea**](StoryArea.md) | A JSON-serialized list of clickable areas to be shown on the story | [optional] 
**PostToChatPage** | Pointer to **bool** | Pass *True* to keep the story accessible after it expires | [optional] 
**ProtectContent** | Pointer to **bool** | Pass *True* if the content of the story must be protected from forwarding and screenshotting | [optional] 

## Methods

### NewPostStoryRequest

`func NewPostStoryRequest(businessConnectionId string, content InputStoryContent, activePeriod int32, ) *PostStoryRequest`

NewPostStoryRequest instantiates a new PostStoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostStoryRequestWithDefaults

`func NewPostStoryRequestWithDefaults() *PostStoryRequest`

NewPostStoryRequestWithDefaults instantiates a new PostStoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostStoryRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostStoryRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostStoryRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetContent

`func (o *PostStoryRequest) GetContent() InputStoryContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostStoryRequest) GetContentOk() (*InputStoryContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostStoryRequest) SetContent(v InputStoryContent)`

SetContent sets Content field to given value.


### GetActivePeriod

`func (o *PostStoryRequest) GetActivePeriod() int32`

GetActivePeriod returns the ActivePeriod field if non-nil, zero value otherwise.

### GetActivePeriodOk

`func (o *PostStoryRequest) GetActivePeriodOk() (*int32, bool)`

GetActivePeriodOk returns a tuple with the ActivePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePeriod

`func (o *PostStoryRequest) SetActivePeriod(v int32)`

SetActivePeriod sets ActivePeriod field to given value.


### GetCaption

`func (o *PostStoryRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *PostStoryRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *PostStoryRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *PostStoryRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *PostStoryRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *PostStoryRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *PostStoryRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *PostStoryRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *PostStoryRequest) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *PostStoryRequest) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *PostStoryRequest) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *PostStoryRequest) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetAreas

`func (o *PostStoryRequest) GetAreas() []StoryArea`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *PostStoryRequest) GetAreasOk() (*[]StoryArea, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *PostStoryRequest) SetAreas(v []StoryArea)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *PostStoryRequest) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetPostToChatPage

`func (o *PostStoryRequest) GetPostToChatPage() bool`

GetPostToChatPage returns the PostToChatPage field if non-nil, zero value otherwise.

### GetPostToChatPageOk

`func (o *PostStoryRequest) GetPostToChatPageOk() (*bool, bool)`

GetPostToChatPageOk returns a tuple with the PostToChatPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostToChatPage

`func (o *PostStoryRequest) SetPostToChatPage(v bool)`

SetPostToChatPage sets PostToChatPage field to given value.

### HasPostToChatPage

`func (o *PostStoryRequest) HasPostToChatPage() bool`

HasPostToChatPage returns a boolean if a field has been set.

### GetProtectContent

`func (o *PostStoryRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *PostStoryRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *PostStoryRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *PostStoryRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


