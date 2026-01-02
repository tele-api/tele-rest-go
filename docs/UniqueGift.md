# UniqueGift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftId** | **string** | Identifier of the regular gift from which the gift was upgraded | 
**BaseName** | **string** | Human-readable name of the regular gift from which this unique gift was upgraded | 
**Name** | **string** | Unique name of the gift. This name can be used in &#x60;https://t.me/nft/...&#x60; links and story areas | 
**Number** | **int32** | Unique number of the upgraded gift among gifts upgraded from the same regular gift | 
**Model** | [**UniqueGiftModel**](UniqueGiftModel.md) |  | 
**Symbol** | [**UniqueGiftSymbol**](UniqueGiftSymbol.md) |  | 
**Backdrop** | [**UniqueGiftBackdrop**](UniqueGiftBackdrop.md) |  | 
**IsPremium** | Pointer to **bool** | *Optional*. *True*, if the original regular gift was exclusively purchaseable by Telegram Premium subscribers | [optional] [default to true]
**IsFromBlockchain** | Pointer to **bool** | *Optional*. *True*, if the gift is assigned from the TON blockchain and can&#39;t be resold or transferred in Telegram | [optional] [default to true]
**Colors** | Pointer to [**UniqueGiftColors**](UniqueGiftColors.md) |  | [optional] 
**PublisherChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 

## Methods

### NewUniqueGift

`func NewUniqueGift(giftId string, baseName string, name string, number int32, model UniqueGiftModel, symbol UniqueGiftSymbol, backdrop UniqueGiftBackdrop, ) *UniqueGift`

NewUniqueGift instantiates a new UniqueGift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueGiftWithDefaults

`func NewUniqueGiftWithDefaults() *UniqueGift`

NewUniqueGiftWithDefaults instantiates a new UniqueGift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftId

`func (o *UniqueGift) GetGiftId() string`

GetGiftId returns the GiftId field if non-nil, zero value otherwise.

### GetGiftIdOk

`func (o *UniqueGift) GetGiftIdOk() (*string, bool)`

GetGiftIdOk returns a tuple with the GiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftId

`func (o *UniqueGift) SetGiftId(v string)`

SetGiftId sets GiftId field to given value.


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


### GetIsPremium

`func (o *UniqueGift) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *UniqueGift) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *UniqueGift) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *UniqueGift) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetIsFromBlockchain

`func (o *UniqueGift) GetIsFromBlockchain() bool`

GetIsFromBlockchain returns the IsFromBlockchain field if non-nil, zero value otherwise.

### GetIsFromBlockchainOk

`func (o *UniqueGift) GetIsFromBlockchainOk() (*bool, bool)`

GetIsFromBlockchainOk returns a tuple with the IsFromBlockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromBlockchain

`func (o *UniqueGift) SetIsFromBlockchain(v bool)`

SetIsFromBlockchain sets IsFromBlockchain field to given value.

### HasIsFromBlockchain

`func (o *UniqueGift) HasIsFromBlockchain() bool`

HasIsFromBlockchain returns a boolean if a field has been set.

### GetColors

`func (o *UniqueGift) GetColors() UniqueGiftColors`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *UniqueGift) GetColorsOk() (*UniqueGiftColors, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *UniqueGift) SetColors(v UniqueGiftColors)`

SetColors sets Colors field to given value.

### HasColors

`func (o *UniqueGift) HasColors() bool`

HasColors returns a boolean if a field has been set.

### GetPublisherChat

`func (o *UniqueGift) GetPublisherChat() Chat`

GetPublisherChat returns the PublisherChat field if non-nil, zero value otherwise.

### GetPublisherChatOk

`func (o *UniqueGift) GetPublisherChatOk() (*Chat, bool)`

GetPublisherChatOk returns a tuple with the PublisherChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherChat

`func (o *UniqueGift) SetPublisherChat(v Chat)`

SetPublisherChat sets PublisherChat field to given value.

### HasPublisherChat

`func (o *UniqueGift) HasPublisherChat() bool`

HasPublisherChat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


