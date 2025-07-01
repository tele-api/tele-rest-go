# BackgroundTypeWallpaper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the background, always “wallpaper” | [default to "wallpaper"]
**Document** | [**Document**](Document.md) |  | 
**DarkThemeDimming** | **int32** | Dimming of the background in dark themes, as a percentage; 0-100 | 
**IsBlurred** | Pointer to **bool** | *Optional*. *True*, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12 | [optional] [default to true]
**IsMoving** | Pointer to **bool** | *Optional*. *True*, if the background moves slightly when the device is tilted | [optional] [default to true]

## Methods

### NewBackgroundTypeWallpaper

`func NewBackgroundTypeWallpaper(type_ string, document Document, darkThemeDimming int32, ) *BackgroundTypeWallpaper`

NewBackgroundTypeWallpaper instantiates a new BackgroundTypeWallpaper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTypeWallpaperWithDefaults

`func NewBackgroundTypeWallpaperWithDefaults() *BackgroundTypeWallpaper`

NewBackgroundTypeWallpaperWithDefaults instantiates a new BackgroundTypeWallpaper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackgroundTypeWallpaper) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackgroundTypeWallpaper) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackgroundTypeWallpaper) SetType(v string)`

SetType sets Type field to given value.


### GetDocument

`func (o *BackgroundTypeWallpaper) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *BackgroundTypeWallpaper) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *BackgroundTypeWallpaper) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetDarkThemeDimming

`func (o *BackgroundTypeWallpaper) GetDarkThemeDimming() int32`

GetDarkThemeDimming returns the DarkThemeDimming field if non-nil, zero value otherwise.

### GetDarkThemeDimmingOk

`func (o *BackgroundTypeWallpaper) GetDarkThemeDimmingOk() (*int32, bool)`

GetDarkThemeDimmingOk returns a tuple with the DarkThemeDimming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkThemeDimming

`func (o *BackgroundTypeWallpaper) SetDarkThemeDimming(v int32)`

SetDarkThemeDimming sets DarkThemeDimming field to given value.


### GetIsBlurred

`func (o *BackgroundTypeWallpaper) GetIsBlurred() bool`

GetIsBlurred returns the IsBlurred field if non-nil, zero value otherwise.

### GetIsBlurredOk

`func (o *BackgroundTypeWallpaper) GetIsBlurredOk() (*bool, bool)`

GetIsBlurredOk returns a tuple with the IsBlurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlurred

`func (o *BackgroundTypeWallpaper) SetIsBlurred(v bool)`

SetIsBlurred sets IsBlurred field to given value.

### HasIsBlurred

`func (o *BackgroundTypeWallpaper) HasIsBlurred() bool`

HasIsBlurred returns a boolean if a field has been set.

### GetIsMoving

`func (o *BackgroundTypeWallpaper) GetIsMoving() bool`

GetIsMoving returns the IsMoving field if non-nil, zero value otherwise.

### GetIsMovingOk

`func (o *BackgroundTypeWallpaper) GetIsMovingOk() (*bool, bool)`

GetIsMovingOk returns a tuple with the IsMoving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMoving

`func (o *BackgroundTypeWallpaper) SetIsMoving(v bool)`

SetIsMoving sets IsMoving field to given value.

### HasIsMoving

`func (o *BackgroundTypeWallpaper) HasIsMoving() bool`

HasIsMoving returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


