# PaidMediaPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the paid media, always “photo” | [default to "photo"]
**Photo** | [**[]PhotoSize**](PhotoSize.md) | The photo | 

## Methods

### NewPaidMediaPhoto

`func NewPaidMediaPhoto(type_ string, photo []PhotoSize, ) *PaidMediaPhoto`

NewPaidMediaPhoto instantiates a new PaidMediaPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidMediaPhotoWithDefaults

`func NewPaidMediaPhotoWithDefaults() *PaidMediaPhoto`

NewPaidMediaPhotoWithDefaults instantiates a new PaidMediaPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaidMediaPhoto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaidMediaPhoto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaidMediaPhoto) SetType(v string)`

SetType sets Type field to given value.


### GetPhoto

`func (o *PaidMediaPhoto) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *PaidMediaPhoto) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *PaidMediaPhoto) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


