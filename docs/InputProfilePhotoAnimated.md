# InputProfilePhotoAnimated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the profile photo, must be *animated* | [default to "animated"]
**Animation** | **string** | The animated profile photo. Profile photos can&#39;t be reused and can only be uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the photo was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**MainFrameTimestamp** | Pointer to **float32** | *Optional*. Timestamp in seconds of the frame that will be used as the static profile photo. Defaults to 0.0. | [optional] 

## Methods

### NewInputProfilePhotoAnimated

`func NewInputProfilePhotoAnimated(type_ string, animation string, ) *InputProfilePhotoAnimated`

NewInputProfilePhotoAnimated instantiates a new InputProfilePhotoAnimated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProfilePhotoAnimatedWithDefaults

`func NewInputProfilePhotoAnimatedWithDefaults() *InputProfilePhotoAnimated`

NewInputProfilePhotoAnimatedWithDefaults instantiates a new InputProfilePhotoAnimated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputProfilePhotoAnimated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputProfilePhotoAnimated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputProfilePhotoAnimated) SetType(v string)`

SetType sets Type field to given value.


### GetAnimation

`func (o *InputProfilePhotoAnimated) GetAnimation() string`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *InputProfilePhotoAnimated) GetAnimationOk() (*string, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *InputProfilePhotoAnimated) SetAnimation(v string)`

SetAnimation sets Animation field to given value.


### GetMainFrameTimestamp

`func (o *InputProfilePhotoAnimated) GetMainFrameTimestamp() float32`

GetMainFrameTimestamp returns the MainFrameTimestamp field if non-nil, zero value otherwise.

### GetMainFrameTimestampOk

`func (o *InputProfilePhotoAnimated) GetMainFrameTimestampOk() (*float32, bool)`

GetMainFrameTimestampOk returns a tuple with the MainFrameTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainFrameTimestamp

`func (o *InputProfilePhotoAnimated) SetMainFrameTimestamp(v float32)`

SetMainFrameTimestamp sets MainFrameTimestamp field to given value.

### HasMainFrameTimestamp

`func (o *InputProfilePhotoAnimated) HasMainFrameTimestamp() bool`

HasMainFrameTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


