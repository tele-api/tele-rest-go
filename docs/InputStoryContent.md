# InputStoryContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the content, must be *video* | [default to "video"]
**Photo** | **string** | The photo to post as a story. The photo must be of the size 1080x1920 and must not exceed 10 MB. The photo can&#39;t be reused and can only be uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the photo was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Video** | **string** | The video to post as a story. The video must be of the size 720x1280, streamable, encoded with H.265 codec, with key frames added each second in the MPEG4 format, and must not exceed 30 MB. The video can&#39;t be reused and can only be uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the video was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Duration** | Pointer to **float32** | *Optional*. Precise duration of the video in seconds; 0-60 | [optional] 
**CoverFrameTimestamp** | Pointer to **float32** | *Optional*. Timestamp in seconds of the frame that will be used as the static cover for the story. Defaults to 0.0. | [optional] 
**IsAnimation** | Pointer to **bool** | *Optional*. Pass *True* if the video has no sound | [optional] 

## Methods

### NewInputStoryContent

`func NewInputStoryContent(type_ string, photo string, video string, ) *InputStoryContent`

NewInputStoryContent instantiates a new InputStoryContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputStoryContentWithDefaults

`func NewInputStoryContentWithDefaults() *InputStoryContent`

NewInputStoryContentWithDefaults instantiates a new InputStoryContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputStoryContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputStoryContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputStoryContent) SetType(v string)`

SetType sets Type field to given value.


### GetPhoto

`func (o *InputStoryContent) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *InputStoryContent) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *InputStoryContent) SetPhoto(v string)`

SetPhoto sets Photo field to given value.


### GetVideo

`func (o *InputStoryContent) GetVideo() string`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *InputStoryContent) GetVideoOk() (*string, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *InputStoryContent) SetVideo(v string)`

SetVideo sets Video field to given value.


### GetDuration

`func (o *InputStoryContent) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InputStoryContent) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InputStoryContent) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InputStoryContent) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCoverFrameTimestamp

`func (o *InputStoryContent) GetCoverFrameTimestamp() float32`

GetCoverFrameTimestamp returns the CoverFrameTimestamp field if non-nil, zero value otherwise.

### GetCoverFrameTimestampOk

`func (o *InputStoryContent) GetCoverFrameTimestampOk() (*float32, bool)`

GetCoverFrameTimestampOk returns a tuple with the CoverFrameTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverFrameTimestamp

`func (o *InputStoryContent) SetCoverFrameTimestamp(v float32)`

SetCoverFrameTimestamp sets CoverFrameTimestamp field to given value.

### HasCoverFrameTimestamp

`func (o *InputStoryContent) HasCoverFrameTimestamp() bool`

HasCoverFrameTimestamp returns a boolean if a field has been set.

### GetIsAnimation

`func (o *InputStoryContent) GetIsAnimation() bool`

GetIsAnimation returns the IsAnimation field if non-nil, zero value otherwise.

### GetIsAnimationOk

`func (o *InputStoryContent) GetIsAnimationOk() (*bool, bool)`

GetIsAnimationOk returns a tuple with the IsAnimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnimation

`func (o *InputStoryContent) SetIsAnimation(v bool)`

SetIsAnimation sets IsAnimation field to given value.

### HasIsAnimation

`func (o *InputStoryContent) HasIsAnimation() bool`

HasIsAnimation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


