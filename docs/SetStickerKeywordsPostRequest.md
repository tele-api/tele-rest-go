# SetStickerKeywordsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**Keywords** | Pointer to **[]string** | A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters | [optional] 

## Methods

### NewSetStickerKeywordsPostRequest

`func NewSetStickerKeywordsPostRequest(sticker string, ) *SetStickerKeywordsPostRequest`

NewSetStickerKeywordsPostRequest instantiates a new SetStickerKeywordsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerKeywordsPostRequestWithDefaults

`func NewSetStickerKeywordsPostRequestWithDefaults() *SetStickerKeywordsPostRequest`

NewSetStickerKeywordsPostRequestWithDefaults instantiates a new SetStickerKeywordsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *SetStickerKeywordsPostRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *SetStickerKeywordsPostRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *SetStickerKeywordsPostRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetKeywords

`func (o *SetStickerKeywordsPostRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SetStickerKeywordsPostRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SetStickerKeywordsPostRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SetStickerKeywordsPostRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


