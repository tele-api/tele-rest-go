# PostSetStickerKeywordsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**Keywords** | Pointer to **[]string** | A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters | [optional] 

## Methods

### NewPostSetStickerKeywordsRequest

`func NewPostSetStickerKeywordsRequest(sticker string, ) *PostSetStickerKeywordsRequest`

NewPostSetStickerKeywordsRequest instantiates a new PostSetStickerKeywordsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetStickerKeywordsRequestWithDefaults

`func NewPostSetStickerKeywordsRequestWithDefaults() *PostSetStickerKeywordsRequest`

NewPostSetStickerKeywordsRequestWithDefaults instantiates a new PostSetStickerKeywordsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *PostSetStickerKeywordsRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *PostSetStickerKeywordsRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *PostSetStickerKeywordsRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetKeywords

`func (o *PostSetStickerKeywordsRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *PostSetStickerKeywordsRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *PostSetStickerKeywordsRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *PostSetStickerKeywordsRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


