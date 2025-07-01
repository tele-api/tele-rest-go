# InputPaidMediaVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the media, must be *video* | [default to "video"]
**Media** | **string** | File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Thumbnail** | Pointer to **string** | *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail&#39;s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can&#39;t be reused and can be only uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the thumbnail was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**Cover** | Pointer to **string** | *Optional*. Cover for the video in the message. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**StartTimestamp** | Pointer to **int32** | *Optional*. Start timestamp for the video in the message | [optional] 
**Width** | Pointer to **int32** | *Optional*. Video width | [optional] 
**Height** | Pointer to **int32** | *Optional*. Video height | [optional] 
**Duration** | Pointer to **int32** | *Optional*. Video duration in seconds | [optional] 
**SupportsStreaming** | Pointer to **bool** | *Optional*. Pass *True* if the uploaded video is suitable for streaming | [optional] 

## Methods

### NewInputPaidMediaVideo

`func NewInputPaidMediaVideo(type_ string, media string, ) *InputPaidMediaVideo`

NewInputPaidMediaVideo instantiates a new InputPaidMediaVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputPaidMediaVideoWithDefaults

`func NewInputPaidMediaVideoWithDefaults() *InputPaidMediaVideo`

NewInputPaidMediaVideoWithDefaults instantiates a new InputPaidMediaVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputPaidMediaVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputPaidMediaVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputPaidMediaVideo) SetType(v string)`

SetType sets Type field to given value.


### GetMedia

`func (o *InputPaidMediaVideo) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *InputPaidMediaVideo) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *InputPaidMediaVideo) SetMedia(v string)`

SetMedia sets Media field to given value.


### GetThumbnail

`func (o *InputPaidMediaVideo) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *InputPaidMediaVideo) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *InputPaidMediaVideo) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *InputPaidMediaVideo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetCover

`func (o *InputPaidMediaVideo) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *InputPaidMediaVideo) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *InputPaidMediaVideo) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *InputPaidMediaVideo) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *InputPaidMediaVideo) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *InputPaidMediaVideo) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *InputPaidMediaVideo) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *InputPaidMediaVideo) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetWidth

`func (o *InputPaidMediaVideo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InputPaidMediaVideo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InputPaidMediaVideo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InputPaidMediaVideo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *InputPaidMediaVideo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InputPaidMediaVideo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InputPaidMediaVideo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InputPaidMediaVideo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDuration

`func (o *InputPaidMediaVideo) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InputPaidMediaVideo) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InputPaidMediaVideo) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InputPaidMediaVideo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSupportsStreaming

`func (o *InputPaidMediaVideo) GetSupportsStreaming() bool`

GetSupportsStreaming returns the SupportsStreaming field if non-nil, zero value otherwise.

### GetSupportsStreamingOk

`func (o *InputPaidMediaVideo) GetSupportsStreamingOk() (*bool, bool)`

GetSupportsStreamingOk returns a tuple with the SupportsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStreaming

`func (o *InputPaidMediaVideo) SetSupportsStreaming(v bool)`

SetSupportsStreaming sets SupportsStreaming field to given value.

### HasSupportsStreaming

`func (o *InputPaidMediaVideo) HasSupportsStreaming() bool`

HasSupportsStreaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


