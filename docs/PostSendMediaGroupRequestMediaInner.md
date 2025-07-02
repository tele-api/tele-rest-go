# PostSendMediaGroupRequestMediaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *video* | [default to "video"]
**Media** | **string** | File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Thumbnail** | Pointer to **string** | *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail&#39;s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can&#39;t be reused and can be only uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the thumbnail was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the video to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the video caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**Duration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**Performer** | Pointer to **string** | *Optional*. Performer of the audio | [optional] 
**Title** | Pointer to **string** | *Optional*. Title of the audio | [optional] 
**DisableContentTypeDetection** | Pointer to **bool** | *Optional*. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always *True*, if the document is sent as part of an album. | [optional] 
**ShowCaptionAboveMedia** | Pointer to **bool** | *Optional*. Pass *True*, if the caption must be shown above the message media | [optional] 
**HasSpoiler** | Pointer to **bool** | *Optional*. Pass *True* if the video needs to be covered with a spoiler animation | [optional] 
**Cover** | Pointer to **string** | *Optional*. Cover for the video in the message. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**StartTimestamp** | Pointer to **int32** | *Optional*. Start timestamp for the video in the message | [optional] 
**Width** | Pointer to **int32** | *Optional*. Video width | [optional] 
**Height** | Pointer to **int32** | *Optional*. Video height | [optional] 
**SupportsStreaming** | Pointer to **bool** | *Optional*. Pass *True* if the uploaded video is suitable for streaming | [optional] 

## Methods

### NewPostSendMediaGroupRequestMediaInner

`func NewPostSendMediaGroupRequestMediaInner(type_ string, media string, ) *PostSendMediaGroupRequestMediaInner`

NewPostSendMediaGroupRequestMediaInner instantiates a new PostSendMediaGroupRequestMediaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSendMediaGroupRequestMediaInnerWithDefaults

`func NewPostSendMediaGroupRequestMediaInnerWithDefaults() *PostSendMediaGroupRequestMediaInner`

NewPostSendMediaGroupRequestMediaInnerWithDefaults instantiates a new PostSendMediaGroupRequestMediaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostSendMediaGroupRequestMediaInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostSendMediaGroupRequestMediaInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostSendMediaGroupRequestMediaInner) SetType(v string)`

SetType sets Type field to given value.


### GetMedia

`func (o *PostSendMediaGroupRequestMediaInner) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *PostSendMediaGroupRequestMediaInner) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *PostSendMediaGroupRequestMediaInner) SetMedia(v string)`

SetMedia sets Media field to given value.


### GetThumbnail

`func (o *PostSendMediaGroupRequestMediaInner) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *PostSendMediaGroupRequestMediaInner) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *PostSendMediaGroupRequestMediaInner) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *PostSendMediaGroupRequestMediaInner) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetCaption

`func (o *PostSendMediaGroupRequestMediaInner) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *PostSendMediaGroupRequestMediaInner) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *PostSendMediaGroupRequestMediaInner) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *PostSendMediaGroupRequestMediaInner) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *PostSendMediaGroupRequestMediaInner) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *PostSendMediaGroupRequestMediaInner) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *PostSendMediaGroupRequestMediaInner) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *PostSendMediaGroupRequestMediaInner) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *PostSendMediaGroupRequestMediaInner) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *PostSendMediaGroupRequestMediaInner) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *PostSendMediaGroupRequestMediaInner) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *PostSendMediaGroupRequestMediaInner) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetDuration

`func (o *PostSendMediaGroupRequestMediaInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PostSendMediaGroupRequestMediaInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PostSendMediaGroupRequestMediaInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PostSendMediaGroupRequestMediaInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPerformer

`func (o *PostSendMediaGroupRequestMediaInner) GetPerformer() string`

GetPerformer returns the Performer field if non-nil, zero value otherwise.

### GetPerformerOk

`func (o *PostSendMediaGroupRequestMediaInner) GetPerformerOk() (*string, bool)`

GetPerformerOk returns a tuple with the Performer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformer

`func (o *PostSendMediaGroupRequestMediaInner) SetPerformer(v string)`

SetPerformer sets Performer field to given value.

### HasPerformer

`func (o *PostSendMediaGroupRequestMediaInner) HasPerformer() bool`

HasPerformer returns a boolean if a field has been set.

### GetTitle

`func (o *PostSendMediaGroupRequestMediaInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSendMediaGroupRequestMediaInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSendMediaGroupRequestMediaInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostSendMediaGroupRequestMediaInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDisableContentTypeDetection

`func (o *PostSendMediaGroupRequestMediaInner) GetDisableContentTypeDetection() bool`

GetDisableContentTypeDetection returns the DisableContentTypeDetection field if non-nil, zero value otherwise.

### GetDisableContentTypeDetectionOk

`func (o *PostSendMediaGroupRequestMediaInner) GetDisableContentTypeDetectionOk() (*bool, bool)`

GetDisableContentTypeDetectionOk returns a tuple with the DisableContentTypeDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableContentTypeDetection

`func (o *PostSendMediaGroupRequestMediaInner) SetDisableContentTypeDetection(v bool)`

SetDisableContentTypeDetection sets DisableContentTypeDetection field to given value.

### HasDisableContentTypeDetection

`func (o *PostSendMediaGroupRequestMediaInner) HasDisableContentTypeDetection() bool`

HasDisableContentTypeDetection returns a boolean if a field has been set.

### GetShowCaptionAboveMedia

`func (o *PostSendMediaGroupRequestMediaInner) GetShowCaptionAboveMedia() bool`

GetShowCaptionAboveMedia returns the ShowCaptionAboveMedia field if non-nil, zero value otherwise.

### GetShowCaptionAboveMediaOk

`func (o *PostSendMediaGroupRequestMediaInner) GetShowCaptionAboveMediaOk() (*bool, bool)`

GetShowCaptionAboveMediaOk returns a tuple with the ShowCaptionAboveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCaptionAboveMedia

`func (o *PostSendMediaGroupRequestMediaInner) SetShowCaptionAboveMedia(v bool)`

SetShowCaptionAboveMedia sets ShowCaptionAboveMedia field to given value.

### HasShowCaptionAboveMedia

`func (o *PostSendMediaGroupRequestMediaInner) HasShowCaptionAboveMedia() bool`

HasShowCaptionAboveMedia returns a boolean if a field has been set.

### GetHasSpoiler

`func (o *PostSendMediaGroupRequestMediaInner) GetHasSpoiler() bool`

GetHasSpoiler returns the HasSpoiler field if non-nil, zero value otherwise.

### GetHasSpoilerOk

`func (o *PostSendMediaGroupRequestMediaInner) GetHasSpoilerOk() (*bool, bool)`

GetHasSpoilerOk returns a tuple with the HasSpoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpoiler

`func (o *PostSendMediaGroupRequestMediaInner) SetHasSpoiler(v bool)`

SetHasSpoiler sets HasSpoiler field to given value.

### HasHasSpoiler

`func (o *PostSendMediaGroupRequestMediaInner) HasHasSpoiler() bool`

HasHasSpoiler returns a boolean if a field has been set.

### GetCover

`func (o *PostSendMediaGroupRequestMediaInner) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *PostSendMediaGroupRequestMediaInner) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *PostSendMediaGroupRequestMediaInner) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *PostSendMediaGroupRequestMediaInner) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *PostSendMediaGroupRequestMediaInner) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *PostSendMediaGroupRequestMediaInner) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *PostSendMediaGroupRequestMediaInner) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *PostSendMediaGroupRequestMediaInner) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetWidth

`func (o *PostSendMediaGroupRequestMediaInner) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PostSendMediaGroupRequestMediaInner) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PostSendMediaGroupRequestMediaInner) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PostSendMediaGroupRequestMediaInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PostSendMediaGroupRequestMediaInner) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PostSendMediaGroupRequestMediaInner) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PostSendMediaGroupRequestMediaInner) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PostSendMediaGroupRequestMediaInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetSupportsStreaming

`func (o *PostSendMediaGroupRequestMediaInner) GetSupportsStreaming() bool`

GetSupportsStreaming returns the SupportsStreaming field if non-nil, zero value otherwise.

### GetSupportsStreamingOk

`func (o *PostSendMediaGroupRequestMediaInner) GetSupportsStreamingOk() (*bool, bool)`

GetSupportsStreamingOk returns a tuple with the SupportsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStreaming

`func (o *PostSendMediaGroupRequestMediaInner) SetSupportsStreaming(v bool)`

SetSupportsStreaming sets SupportsStreaming field to given value.

### HasSupportsStreaming

`func (o *PostSendMediaGroupRequestMediaInner) HasSupportsStreaming() bool`

HasSupportsStreaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


