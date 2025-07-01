# BackgroundFillSolid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background fill, always “solid” | [default to "solid"]
**Color** | **int32** | The color of the background fill in the RGB24 format | 

## Methods

### NewBackgroundFillSolid

`func NewBackgroundFillSolid(type_ string, color int32, ) *BackgroundFillSolid`

NewBackgroundFillSolid instantiates a new BackgroundFillSolid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundFillSolidWithDefaults

`func NewBackgroundFillSolidWithDefaults() *BackgroundFillSolid`

NewBackgroundFillSolidWithDefaults instantiates a new BackgroundFillSolid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundFillSolid) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundFillSolid) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundFillSolid) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *BackgroundFillSolid) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BackgroundFillSolid) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BackgroundFillSolid) SetColor(v int32)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


