# SetStickerEmojiListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**EmojiList** | **[]string** | A JSON-serialized list of 1-20 emoji associated with the sticker | 

## Methods

### NewSetStickerEmojiListRequest

`func NewSetStickerEmojiListRequest(sticker string, emojiList []string, ) *SetStickerEmojiListRequest`

NewSetStickerEmojiListRequest instantiates a new SetStickerEmojiListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerEmojiListRequestWithDefaults

`func NewSetStickerEmojiListRequestWithDefaults() *SetStickerEmojiListRequest`

NewSetStickerEmojiListRequestWithDefaults instantiates a new SetStickerEmojiListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *SetStickerEmojiListRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *SetStickerEmojiListRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *SetStickerEmojiListRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetEmojiList

`func (o *SetStickerEmojiListRequest) GetEmojiList() []string`

GetEmojiList returns the EmojiList field if non-nil, zero value otherwise.

### GetEmojiListOk

`func (o *SetStickerEmojiListRequest) GetEmojiListOk() (*[]string, bool)`

GetEmojiListOk returns a tuple with the EmojiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiList

`func (o *SetStickerEmojiListRequest) SetEmojiList(v []string)`

SetEmojiList sets EmojiList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


