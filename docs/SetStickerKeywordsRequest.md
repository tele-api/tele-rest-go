# SetStickerKeywordsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**Keywords** | Pointer to **[]string** | A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters | [optional] 

## Methods

### NewSetStickerKeywordsRequest

`func NewSetStickerKeywordsRequest(sticker string, ) *SetStickerKeywordsRequest`

NewSetStickerKeywordsRequest instantiates a new SetStickerKeywordsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerKeywordsRequestWithDefaults

`func NewSetStickerKeywordsRequestWithDefaults() *SetStickerKeywordsRequest`

NewSetStickerKeywordsRequestWithDefaults instantiates a new SetStickerKeywordsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *SetStickerKeywordsRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *SetStickerKeywordsRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *SetStickerKeywordsRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetKeywords

`func (o *SetStickerKeywordsRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SetStickerKeywordsRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SetStickerKeywordsRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SetStickerKeywordsRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


