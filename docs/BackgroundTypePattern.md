# BackgroundTypePattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background, always “pattern” | [default to "pattern"]
**Document** | [**Document**](Document.md) |  | 
**Fill** | [**BackgroundFill**](BackgroundFill.md) |  | 
**Intensity** | **int32** | Intensity of the pattern when it is shown above the filled background; 0-100 | 
**IsInverted** | Pointer to **bool** | *Optional*. *True*, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only | [optional] [default to true]
**IsMoving** | Pointer to **bool** | *Optional*. *True*, if the background moves slightly when the device is tilted | [optional] [default to true]

## Methods

### NewBackgroundTypePattern

`func NewBackgroundTypePattern(type_ string, document Document, fill BackgroundFill, intensity int32, ) *BackgroundTypePattern`

NewBackgroundTypePattern instantiates a new BackgroundTypePattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTypePatternWithDefaults

`func NewBackgroundTypePatternWithDefaults() *BackgroundTypePattern`

NewBackgroundTypePatternWithDefaults instantiates a new BackgroundTypePattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundTypePattern) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundTypePattern) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundTypePattern) SetType(v string)`

SetType sets Type field to given value.


### GetDocument

`func (o *BackgroundTypePattern) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *BackgroundTypePattern) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *BackgroundTypePattern) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetFill

`func (o *BackgroundTypePattern) GetFill() BackgroundFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *BackgroundTypePattern) GetFillOk() (*BackgroundFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *BackgroundTypePattern) SetFill(v BackgroundFill)`

SetFill sets Fill field to given value.


### GetIntensity

`func (o *BackgroundTypePattern) GetIntensity() int32`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *BackgroundTypePattern) GetIntensityOk() (*int32, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *BackgroundTypePattern) SetIntensity(v int32)`

SetIntensity sets Intensity field to given value.


### GetIsInverted

`func (o *BackgroundTypePattern) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *BackgroundTypePattern) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *BackgroundTypePattern) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.

### HasIsInverted

`func (o *BackgroundTypePattern) HasIsInverted() bool`

HasIsInverted returns a boolean if a field has been set.

### GetIsMoving

`func (o *BackgroundTypePattern) GetIsMoving() bool`

GetIsMoving returns the IsMoving field if non-nil, zero value otherwise.

### GetIsMovingOk

`func (o *BackgroundTypePattern) GetIsMovingOk() (*bool, bool)`

GetIsMovingOk returns a tuple with the IsMoving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMoving

`func (o *BackgroundTypePattern) SetIsMoving(v bool)`

SetIsMoving sets IsMoving field to given value.

### HasIsMoving

`func (o *BackgroundTypePattern) HasIsMoving() bool`

HasIsMoving returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


