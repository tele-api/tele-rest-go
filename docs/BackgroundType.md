# BackgroundType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background, always “chat\\_theme” | [default to "chat_theme"]
**Fill** | [**BackgroundFill**](BackgroundFill.md) |  | 
**DarkThemeDimming** | **int32** | Dimming of the background in dark themes, as a percentage; 0-100 | 
**Document** | [**Document**](Document.md) |  | 
**IsBlurred** | Pointer to **bool** | *Optional*. *True*, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12 | [optional] [default to true]
**IsMoving** | Pointer to **bool** | *Optional*. *True*, if the background moves slightly when the device is tilted | [optional] [default to true]
**Intensity** | **int32** | Intensity of the pattern when it is shown above the filled background; 0-100 | 
**IsInverted** | Pointer to **bool** | *Optional*. *True*, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only | [optional] [default to true]
**ThemeName** | **string** | Name of the chat theme, which is usually an emoji | 

## Methods

### NewBackgroundType

`func NewBackgroundType(type_ string, fill BackgroundFill, darkThemeDimming int32, document Document, intensity int32, themeName string, ) *BackgroundType`

NewBackgroundType instantiates a new BackgroundType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTypeWithDefaults

`func NewBackgroundTypeWithDefaults() *BackgroundType`

NewBackgroundTypeWithDefaults instantiates a new BackgroundType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundType) SetType(v string)`

SetType sets Type field to given value.


### GetFill

`func (o *BackgroundType) GetFill() BackgroundFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *BackgroundType) GetFillOk() (*BackgroundFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *BackgroundType) SetFill(v BackgroundFill)`

SetFill sets Fill field to given value.


### GetDarkThemeDimming

`func (o *BackgroundType) GetDarkThemeDimming() int32`

GetDarkThemeDimming returns the DarkThemeDimming field if non-nil, zero value otherwise.

### GetDarkThemeDimmingOk

`func (o *BackgroundType) GetDarkThemeDimmingOk() (*int32, bool)`

GetDarkThemeDimmingOk returns a tuple with the DarkThemeDimming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkThemeDimming

`func (o *BackgroundType) SetDarkThemeDimming(v int32)`

SetDarkThemeDimming sets DarkThemeDimming field to given value.


### GetDocument

`func (o *BackgroundType) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *BackgroundType) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *BackgroundType) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetIsBlurred

`func (o *BackgroundType) GetIsBlurred() bool`

GetIsBlurred returns the IsBlurred field if non-nil, zero value otherwise.

### GetIsBlurredOk

`func (o *BackgroundType) GetIsBlurredOk() (*bool, bool)`

GetIsBlurredOk returns a tuple with the IsBlurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlurred

`func (o *BackgroundType) SetIsBlurred(v bool)`

SetIsBlurred sets IsBlurred field to given value.

### HasIsBlurred

`func (o *BackgroundType) HasIsBlurred() bool`

HasIsBlurred returns a boolean if a field has been set.

### GetIsMoving

`func (o *BackgroundType) GetIsMoving() bool`

GetIsMoving returns the IsMoving field if non-nil, zero value otherwise.

### GetIsMovingOk

`func (o *BackgroundType) GetIsMovingOk() (*bool, bool)`

GetIsMovingOk returns a tuple with the IsMoving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMoving

`func (o *BackgroundType) SetIsMoving(v bool)`

SetIsMoving sets IsMoving field to given value.

### HasIsMoving

`func (o *BackgroundType) HasIsMoving() bool`

HasIsMoving returns a boolean if a field has been set.

### GetIntensity

`func (o *BackgroundType) GetIntensity() int32`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *BackgroundType) GetIntensityOk() (*int32, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *BackgroundType) SetIntensity(v int32)`

SetIntensity sets Intensity field to given value.


### GetIsInverted

`func (o *BackgroundType) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *BackgroundType) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *BackgroundType) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.

### HasIsInverted

`func (o *BackgroundType) HasIsInverted() bool`

HasIsInverted returns a boolean if a field has been set.

### GetThemeName

`func (o *BackgroundType) GetThemeName() string`

GetThemeName returns the ThemeName field if non-nil, zero value otherwise.

### GetThemeNameOk

`func (o *BackgroundType) GetThemeNameOk() (*string, bool)`

GetThemeNameOk returns a tuple with the ThemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeName

`func (o *BackgroundType) SetThemeName(v string)`

SetThemeName sets ThemeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


