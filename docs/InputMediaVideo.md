# InputMediaVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *video* | [default to "video"]
**Media** | **string** | File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Thumbnail** | Pointer to **string** | *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail&#39;s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can&#39;t be reused and can be only uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the thumbnail was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**Cover** | Pointer to **string** | *Optional*. Cover for the video in the message. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**StartTimestamp** | Pointer to **int32** | *Optional*. Start timestamp for the video in the message | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the video to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**Width** | Pointer to **int32** | *Optional*. Video width | [optional] 
**Height** | Pointer to **int32** | *Optional*. Video height | [optional] 
**Duration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**SupportsStreaming** | Pointer to **bool** | *Optional*. Pass *True* if the uploaded video is suitable for streaming | [optional] 
**HasSpoiler** | Pointer to **bool** | *Optional*. Pass *True* if the video needs to be covered with a spoiler animation | [optional] 

## Methods

### NewInputMediaVideo

`func NewInputMediaVideo(type_ string, media string, ) *InputMediaVideo`

NewInputMediaVideo instantiates a new InputMediaVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMediaVideoWithDefaults

`func NewInputMediaVideoWithDefaults() *InputMediaVideo`

NewInputMediaVideoWithDefaults instantiates a new InputMediaVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMediaVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMediaVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMediaVideo) SetType(v string)`

SetType sets Type field to given value.


### GetMedia

`func (o *InputMediaVideo) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *InputMediaVideo) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *InputMediaVideo) SetMedia(v string)`

SetMedia sets Media field to given value.


### GetThumbnail

`func (o *InputMediaVideo) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *InputMediaVideo) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *InputMediaVideo) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *InputMediaVideo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetCover

`func (o *InputMediaVideo) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *InputMediaVideo) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *InputMediaVideo) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *InputMediaVideo) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *InputMediaVideo) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *InputMediaVideo) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *InputMediaVideo) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *InputMediaVideo) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetCaption

`func (o *InputMediaVideo) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InputMediaVideo) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InputMediaVideo) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InputMediaVideo) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InputMediaVideo) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InputMediaVideo) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InputMediaVideo) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InputMediaVideo) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InputMediaVideo) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InputMediaVideo) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InputMediaVideo) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InputMediaVideo) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *InputMediaVideo) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *InputMediaVideo) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *InputMediaVideo) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *InputMediaVideo) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetWidth

`func (o *InputMediaVideo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InputMediaVideo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InputMediaVideo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InputMediaVideo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *InputMediaVideo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InputMediaVideo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InputMediaVideo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InputMediaVideo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDuration

`func (o *InputMediaVideo) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InputMediaVideo) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InputMediaVideo) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InputMediaVideo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSupportsStreaming

`func (o *InputMediaVideo) GetSupportsStreaming() bool`

GetSupportsStreaming returns the SupportsStreaming field if non-nil, zero value otherwise.

### GetSupportsStreamingOk

`func (o *InputMediaVideo) GetSupportsStreamingOk() (*bool, bool)`

GetSupportsStreamingOk returns a tuple with the SupportsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStreaming

`func (o *InputMediaVideo) SetSupportsStreaming(v bool)`

SetSupportsStreaming sets SupportsStreaming field to given value.

### HasSupportsStreaming

`func (o *InputMediaVideo) HasSupportsStreaming() bool`

HasSupportsStreaming returns a boolean if a field has been set.

### GetHasSpoiler

`func (o *InputMediaVideo) GetHasSpoiler() bool`

GetHasSpoiler returns the HasSpoiler field if non-nil, zero value otherwise.

### GetHasSpoilerOk

`func (o *InputMediaVideo) GetHasSpoilerOk() (*bool, bool)`

GetHasSpoilerOk returns a tuple with the HasSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpoiler

`func (o *InputMediaVideo) SetHasSpoiler(v bool)`

SetHasSpoiler sets HasSpoiler field to given value.

### HasHasSpoiler

`func (o *InputMediaVideo) HasHasSpoiler() bool`

HasHasSpoiler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


