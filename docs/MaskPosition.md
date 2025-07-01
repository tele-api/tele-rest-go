# MaskPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | **string** | The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”. | 
**XShift** | **float32** | Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position. | 
**YShift** | **float32** | Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position. | 
**Scale** | **float32** | Mask scaling coefficient. For example, 2.0 means double size. | 

## Methods

### NewMaskPosition

`func NewMaskPosition(point string, xShift float32, yShift float32, scale float32, ) *MaskPosition`

NewMaskPosition instantiates a new MaskPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskPositionWithDefaults

`func NewMaskPositionWithDefaults() *MaskPosition`

NewMaskPositionWithDefaults instantiates a new MaskPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *MaskPosition) GetPoint() string`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *MaskPosition) GetPointOk() (*string, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *MaskPosition) SetPoint(v string)`

SetPoint sets Point field to given value.


### GetXShift

`func (o *MaskPosition) GetXShift() float32`

GetXShift returns the XShift field if non-nil, zero value otherwise.

### GetXShiftOk

`func (o *MaskPosition) GetXShiftOk() (*float32, bool)`

GetXShiftOk returns a tuple with the XShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXShift

`func (o *MaskPosition) SetXShift(v float32)`

SetXShift sets XShift field to given value.


### GetYShift

`func (o *MaskPosition) GetYShift() float32`

GetYShift returns the YShift field if non-nil, zero value otherwise.

### GetYShiftOk

`func (o *MaskPosition) GetYShiftOk() (*float32, bool)`

GetYShiftOk returns a tuple with the YShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYShift

`func (o *MaskPosition) SetYShift(v float32)`

SetYShift sets YShift field to given value.


### GetScale

`func (o *MaskPosition) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *MaskPosition) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *MaskPosition) SetScale(v float32)`

SetScale sets Scale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


