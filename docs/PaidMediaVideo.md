# PaidMediaVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the paid media, always “video” | [default to "video"]
**Video** | [**Video**](Video.md) |  | 

## Methods

### NewPaidMediaVideo

`func NewPaidMediaVideo(type_ string, video Video, ) *PaidMediaVideo`

NewPaidMediaVideo instantiates a new PaidMediaVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidMediaVideoWithDefaults

`func NewPaidMediaVideoWithDefaults() *PaidMediaVideo`

NewPaidMediaVideoWithDefaults instantiates a new PaidMediaVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaidMediaVideo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaidMediaVideo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaidMediaVideo) SetType(v string)`

SetType sets Type field to given value.


### GetVideo

`func (o *PaidMediaVideo) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PaidMediaVideo) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PaidMediaVideo) SetVideo(v Video)`

SetVideo sets Video field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


