# BusinessIntro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | *Optional*. Title text of the business intro | [optional] 
**Message** | Pointer to **string** | *Optional*. Message text of the business intro | [optional] 
**Sticker** | Pointer to [**Sticker**](Sticker.md) |  | [optional] 

## Methods

### NewBusinessIntro

`func NewBusinessIntro() *BusinessIntro`

NewBusinessIntro instantiates a new BusinessIntro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessIntroWithDefaults

`func NewBusinessIntroWithDefaults() *BusinessIntro`

NewBusinessIntroWithDefaults instantiates a new BusinessIntro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BusinessIntro) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BusinessIntro) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BusinessIntro) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BusinessIntro) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessage

`func (o *BusinessIntro) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BusinessIntro) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BusinessIntro) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BusinessIntro) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSticker

`func (o *BusinessIntro) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *BusinessIntro) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *BusinessIntro) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *BusinessIntro) HasSticker() bool`

HasSticker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


