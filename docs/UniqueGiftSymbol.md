# UniqueGiftSymbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the symbol | 
**Sticker** | [**Sticker**](Sticker.md) |  | 
**RarityPerMille** | **int32** | The number of unique gifts that receive this model for every 1000 gifts upgraded | 

## Methods

### NewUniqueGiftSymbol

`func NewUniqueGiftSymbol(name string, sticker Sticker, rarityPerMille int32, ) *UniqueGiftSymbol`

NewUniqueGiftSymbol instantiates a new UniqueGiftSymbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueGiftSymbolWithDefaults

`func NewUniqueGiftSymbolWithDefaults() *UniqueGiftSymbol`

NewUniqueGiftSymbolWithDefaults instantiates a new UniqueGiftSymbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UniqueGiftSymbol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniqueGiftSymbol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniqueGiftSymbol) SetName(v string)`

SetName sets Name field to given value.


### GetSticker

`func (o *UniqueGiftSymbol) GetSticker() Sticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *UniqueGiftSymbol) GetStickerOk() (*Sticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *UniqueGiftSymbol) SetSticker(v Sticker)`

SetSticker sets Sticker field to given value.


### GetRarityPerMille

`func (o *UniqueGiftSymbol) GetRarityPerMille() int32`

GetRarityPerMille returns the RarityPerMille field if non-nil, zero value otherwise.

### GetRarityPerMilleOk

`func (o *UniqueGiftSymbol) GetRarityPerMilleOk() (*int32, bool)`

GetRarityPerMilleOk returns a tuple with the RarityPerMille field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityPerMille

`func (o *UniqueGiftSymbol) SetRarityPerMille(v int32)`

SetRarityPerMille sets RarityPerMille field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


