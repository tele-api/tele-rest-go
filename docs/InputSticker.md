# InputSticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | The added sticker. Pass a *file\\_id* as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new file using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. Animated and video stickers can&#39;t be uploaded via HTTP URL. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Format** | **string** | Format of the added sticker, must be one of “static” for a **.WEBP** or **.PNG** image, “animated” for a **.TGS** animation, “video” for a **.WEBM** video | 
**EmojiList** | **[]string** | List of 1-20 emoji associated with the sticker | 
**MaskPosition** | Pointer to [**MaskPosition**](MaskPosition.md) |  | [optional] 
**Keywords** | Pointer to **[]string** | *Optional*. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom\\_emoji” stickers only. | [optional] 

## Methods

### NewInputSticker

`func NewInputSticker(sticker string, format string, emojiList []string, ) *InputSticker`

NewInputSticker instantiates a new InputSticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputStickerWithDefaults

`func NewInputStickerWithDefaults() *InputSticker`

NewInputStickerWithDefaults instantiates a new InputSticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *InputSticker) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *InputSticker) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *InputSticker) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetFormat

`func (o *InputSticker) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *InputSticker) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *InputSticker) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetEmojiList

`func (o *InputSticker) GetEmojiList() []string`

GetEmojiList returns the EmojiList field if non-nil, zero value otherwise.

### GetEmojiListOk

`func (o *InputSticker) GetEmojiListOk() (*[]string, bool)`

GetEmojiListOk returns a tuple with the EmojiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiList

`func (o *InputSticker) SetEmojiList(v []string)`

SetEmojiList sets EmojiList field to given value.


### GetMaskPosition

`func (o *InputSticker) GetMaskPosition() MaskPosition`

GetMaskPosition returns the MaskPosition field if non-nil, zero value otherwise.

### GetMaskPositionOk

`func (o *InputSticker) GetMaskPositionOk() (*MaskPosition, bool)`

GetMaskPositionOk returns a tuple with the MaskPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPosition

`func (o *InputSticker) SetMaskPosition(v MaskPosition)`

SetMaskPosition sets MaskPosition field to given value.

### HasMaskPosition

`func (o *InputSticker) HasMaskPosition() bool`

HasMaskPosition returns a boolean if a field has been set.

### GetKeywords

`func (o *InputSticker) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *InputSticker) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *InputSticker) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *InputSticker) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


