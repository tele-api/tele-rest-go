# UniqueGift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseName** | **string** | Human-readable name of the regular gift from which this unique gift was upgraded | 
**Name** | **string** | Unique name of the gift. This name can be used in &#x60;https://t.me/nft/...&#x60; links and story areas | 
**Number** | **int32** | Unique number of the upgraded gift among gifts upgraded from the same regular gift | 
**Model** | [**UniqueGiftModel**](UniqueGiftModel.md) |  | 
**Symbol** | [**UniqueGiftSymbol**](UniqueGiftSymbol.md) |  | 
**Backdrop** | [**UniqueGiftBackdrop**](UniqueGiftBackdrop.md) |  | 

## Methods

### NewUniqueGift

`func NewUniqueGift(baseName string, name string, number int32, model UniqueGiftModel, symbol UniqueGiftSymbol, backdrop UniqueGiftBackdrop, ) *UniqueGift`

NewUniqueGift instantiates a new UniqueGift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueGiftWithDefaults

`func NewUniqueGiftWithDefaults() *UniqueGift`

NewUniqueGiftWithDefaults instantiates a new UniqueGift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseName

`func (o *UniqueGift) GetBaseName() string`

GetBaseName returns the BaseName field if non-nil, zero value otherwise.

### GetBaseNameOk

`func (o *UniqueGift) GetBaseNameOk() (*string, bool)`

GetBaseNameOk returns a tuple with the BaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseName

`func (o *UniqueGift) SetBaseName(v string)`

SetBaseName sets BaseName field to given value.


### GetName

`func (o *UniqueGift) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniqueGift) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniqueGift) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *UniqueGift) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UniqueGift) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UniqueGift) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetModel

`func (o *UniqueGift) GetModel() UniqueGiftModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UniqueGift) GetModelOk() (*UniqueGiftModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UniqueGift) SetModel(v UniqueGiftModel)`

SetModel sets Model field to given value.


### GetSymbol

`func (o *UniqueGift) GetSymbol() UniqueGiftSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UniqueGift) GetSymbolOk() (*UniqueGiftSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UniqueGift) SetSymbol(v UniqueGiftSymbol)`

SetSymbol sets Symbol field to given value.


### GetBackdrop

`func (o *UniqueGift) GetBackdrop() UniqueGiftBackdrop`

GetBackdrop returns the Backdrop field if non-nil, zero value otherwise.

### GetBackdropOk

`func (o *UniqueGift) GetBackdropOk() (*UniqueGiftBackdrop, bool)`

GetBackdropOk returns a tuple with the Backdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdrop

`func (o *UniqueGift) SetBackdrop(v UniqueGiftBackdrop)`

SetBackdrop sets Backdrop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


