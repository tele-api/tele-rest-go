# InputStoryContentVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the content, must be *video* | [default to "video"]
**Video** | **string** | The video to post as a story. The video must be of the size 720x1280, streamable, encoded with H.265 codec, with key frames added each second in the MPEG4 format, and must not exceed 30 MB. The video can&#39;t be reused and can only be uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the video was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Duration** | Pointer to **float32** | *Optional*. Precise duration of the video in seconds; 0-60 | [optional] 
**CoverFrameTimestamp** | Pointer to **float32** | *Optional*. Timestamp in seconds of the frame that will be used as the static cover for the story. Defaults to 0.0. | [optional] 
**IsAnimation** | Pointer to **bool** | *Optional*. Pass *True* if the video has no sound | [optional] 

## Methods

### NewInputStoryContentVideo

`func NewInputStoryContentVideo(type_ string, video string, ) *InputStoryContentVideo`

NewInputStoryContentVideo instantiates a new InputStoryContentVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputStoryContentVideoWithDefaults

`func NewInputStoryContentVideoWithDefaults() *InputStoryContentVideo`

NewInputStoryContentVideoWithDefaults instantiates a new InputStoryContentVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputStoryContentVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputStoryContentVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputStoryContentVideo) SetType(v string)`

SetType sets Type field to given value.


### GetVideo

`func (o *InputStoryContentVideo) GetVideo() string`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *InputStoryContentVideo) GetVideoOk() (*string, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *InputStoryContentVideo) SetVideo(v string)`

SetVideo sets Video field to given value.


### GetDuration

`func (o *InputStoryContentVideo) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InputStoryContentVideo) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InputStoryContentVideo) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InputStoryContentVideo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCoverFrameTimestamp

`func (o *InputStoryContentVideo) GetCoverFrameTimestamp() float32`

GetCoverFrameTimestamp returns the CoverFrameTimestamp field if non-nil, zero value otherwise.

### GetCoverFrameTimestampOk

`func (o *InputStoryContentVideo) GetCoverFrameTimestampOk() (*float32, bool)`

GetCoverFrameTimestampOk returns a tuple with the CoverFrameTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverFrameTimestamp

`func (o *InputStoryContentVideo) SetCoverFrameTimestamp(v float32)`

SetCoverFrameTimestamp sets CoverFrameTimestamp field to given value.

### HasCoverFrameTimestamp

`func (o *InputStoryContentVideo) HasCoverFrameTimestamp() bool`

HasCoverFrameTimestamp returns a boolean if a field has been set.

### GetIsAnimation

`func (o *InputStoryContentVideo) GetIsAnimation() bool`

GetIsAnimation returns the IsAnimation field if non-nil, zero value otherwise.

### GetIsAnimationOk

`func (o *InputStoryContentVideo) GetIsAnimationOk() (*bool, bool)`

GetIsAnimationOk returns a tuple with the IsAnimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnimation

`func (o *InputStoryContentVideo) SetIsAnimation(v bool)`

SetIsAnimation sets IsAnimation field to given value.

### HasIsAnimation

`func (o *InputStoryContentVideo) HasIsAnimation() bool`

HasIsAnimation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


