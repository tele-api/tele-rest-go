# EditStoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**StoryId** | **int32** | Unique identifier of the story to edit | 
**Content** | [**InputStoryContent**](InputStoryContent.md) |  | 
**Caption** | Pointer to **string** | Caption of the story, 0-2048 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | Mode for parsing entities in the story caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**Areas** | Pointer to [**[]StoryArea**](StoryArea.md) | A JSON-serialized list of clickable areas to be shown on the story | [optional] 

## Methods

### NewEditStoryRequest

`func NewEditStoryRequest(businessConnectionId string, storyId int32, content InputStoryContent, ) *EditStoryRequest`

NewEditStoryRequest instantiates a new EditStoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditStoryRequestWithDefaults

`func NewEditStoryRequestWithDefaults() *EditStoryRequest`

NewEditStoryRequestWithDefaults instantiates a new EditStoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *EditStoryRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *EditStoryRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *EditStoryRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetStoryId

`func (o *EditStoryRequest) GetStoryId() int32`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *EditStoryRequest) GetStoryIdOk() (*int32, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *EditStoryRequest) SetStoryId(v int32)`

SetStoryId sets StoryId field to given value.


### GetContent

`func (o *EditStoryRequest) GetContent() InputStoryContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EditStoryRequest) GetContentOk() (*InputStoryContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EditStoryRequest) SetContent(v InputStoryContent)`

SetContent sets Content field to given value.


### GetCaption

`func (o *EditStoryRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *EditStoryRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *EditStoryRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *EditStoryRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *EditStoryRequest) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *EditStoryRequest) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *EditStoryRequest) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *EditStoryRequest) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *EditStoryRequest) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *EditStoryRequest) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *EditStoryRequest) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *EditStoryRequest) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetAreas

`func (o *EditStoryRequest) GetAreas() []StoryArea`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *EditStoryRequest) GetAreasOk() (*[]StoryArea, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *EditStoryRequest) SetAreas(v []StoryArea)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *EditStoryRequest) HasAreas() bool`

HasAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


