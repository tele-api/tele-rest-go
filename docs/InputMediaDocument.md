# InputMediaDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the result, must be *document* | [default to "document"]
**Media** | **string** | File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 
**Thumbnail** | Pointer to **string** | *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail&#39;s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can&#39;t be reused and can be only uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the thumbnail was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | [optional] 
**Caption** | Pointer to **string** | *Optional*. Caption of the document to be sent, 0-1024 characters after entities parsing | [optional] 
**ParseMode** | Pointer to **string** | *Optional*. Mode for parsing entities in the document caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details. | [optional] 
**CaptionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. List of special entities that appear in the caption, which can be specified instead of *parse\\_mode* | [optional] 
**DisableContentTypeDetection** | Pointer to **bool** | *Optional*. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always *True*, if the document is sent as part of an album. | [optional] 

## Methods

### NewInputMediaDocument

`func NewInputMediaDocument(type_ string, media string, ) *InputMediaDocument`

NewInputMediaDocument instantiates a new InputMediaDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMediaDocumentWithDefaults

`func NewInputMediaDocumentWithDefaults() *InputMediaDocument`

NewInputMediaDocumentWithDefaults instantiates a new InputMediaDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMediaDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMediaDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMediaDocument) SetType(v string)`

SetType sets Type field to given value.


### GetMedia

`func (o *InputMediaDocument) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *InputMediaDocument) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *InputMediaDocument) SetMedia(v string)`

SetMedia sets Media field to given value.


### GetThumbnail

`func (o *InputMediaDocument) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *InputMediaDocument) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *InputMediaDocument) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *InputMediaDocument) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetCaption

`func (o *InputMediaDocument) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *InputMediaDocument) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *InputMediaDocument) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *InputMediaDocument) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetParseMode

`func (o *InputMediaDocument) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *InputMediaDocument) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *InputMediaDocument) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *InputMediaDocument) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetCaptionEntities

`func (o *InputMediaDocument) GetCaptionEntities() []MessageEntity`

GetCaptionEntities returns the CaptionEntities field if non-nil, zero value otherwise.

### GetCaptionEntitiesOk

`func (o *InputMediaDocument) GetCaptionEntitiesOk() (*[]MessageEntity, bool)`

GetCaptionEntitiesOk returns a tuple with the CaptionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionEntities

`func (o *InputMediaDocument) SetCaptionEntities(v []MessageEntity)`

SetCaptionEntities sets CaptionEntities field to given value.

### HasCaptionEntities

`func (o *InputMediaDocument) HasCaptionEntities() bool`

HasCaptionEntities returns a boolean if a field has been set.

### GetDisableContentTypeDetection

`func (o *InputMediaDocument) GetDisableContentTypeDetection() bool`

GetDisableContentTypeDetection returns the DisableContentTypeDetection field if non-nil, zero value otherwise.

### GetDisableContentTypeDetectionOk

`func (o *InputMediaDocument) GetDisableContentTypeDetectionOk() (*bool, bool)`

GetDisableContentTypeDetectionOk returns a tuple with the DisableContentTypeDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableContentTypeDetection

`func (o *InputMediaDocument) SetDisableContentTypeDetection(v bool)`

SetDisableContentTypeDetection sets DisableContentTypeDetection field to given value.

### HasDisableContentTypeDetection

`func (o *InputMediaDocument) HasDisableContentTypeDetection() bool`

HasDisableContentTypeDetection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


