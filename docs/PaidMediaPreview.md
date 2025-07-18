# PaidMediaPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the paid media, always “preview” | [default to "preview"]
**Width** | Pointer to **int32** | *Optional*. Media width as defined by the sender | [optional] 
**Height** | Pointer to **int32** | *Optional*. Media height as defined by the sender | [optional] 
**Duration** | Pointer to **int32** | *Optional*. Duration of the media in seconds as defined by the sender | [optional] 

## Methods

### NewPaidMediaPreview

`func NewPaidMediaPreview(type_ string, ) *PaidMediaPreview`

NewPaidMediaPreview instantiates a new PaidMediaPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaidMediaPreviewWithDefaults

`func NewPaidMediaPreviewWithDefaults() *PaidMediaPreview`

NewPaidMediaPreviewWithDefaults instantiates a new PaidMediaPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaidMediaPreview) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaidMediaPreview) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaidMediaPreview) SetType(v string)`

SetType sets Type field to given value.


### GetWidth

`func (o *PaidMediaPreview) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PaidMediaPreview) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PaidMediaPreview) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PaidMediaPreview) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PaidMediaPreview) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PaidMediaPreview) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PaidMediaPreview) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PaidMediaPreview) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDuration

`func (o *PaidMediaPreview) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PaidMediaPreview) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PaidMediaPreview) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PaidMediaPreview) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


