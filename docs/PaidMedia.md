# PaidMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the paid media, always “video” | [default to "video"]
**Width** | Pointer to **int32** | *Optional*. Media width as defined by the sender | [optional] 
**Height** | Pointer to **int32** | *Optional*. Media height as defined by the sender | [optional] 
**Duration** | Pointer to **int32** | *Optional*. Duration of the media in seconds as defined by the sender | [optional] 
**Photo** | [**[]PhotoSize**](PhotoSize.md) | The photo | 
**Video** | [**Video**](Video.md) |  | 

## Methods

### NewPaidMedia

`func NewPaidMedia(type_ string, photo []PhotoSize, video Video, ) *PaidMedia`

NewPaidMedia instantiates a new PaidMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidMediaWithDefaults

`func NewPaidMediaWithDefaults() *PaidMedia`

NewPaidMediaWithDefaults instantiates a new PaidMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaidMedia) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaidMedia) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaidMedia) SetType(v string)`

SetType sets Type field to given value.


### GetWidth

`func (o *PaidMedia) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PaidMedia) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PaidMedia) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PaidMedia) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PaidMedia) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PaidMedia) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PaidMedia) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PaidMedia) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDuration

`func (o *PaidMedia) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PaidMedia) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PaidMedia) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PaidMedia) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPhoto

`func (o *PaidMedia) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *PaidMedia) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *PaidMedia) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.


### GetVideo

`func (o *PaidMedia) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PaidMedia) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PaidMedia) SetVideo(v Video)`

SetVideo sets Video field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


