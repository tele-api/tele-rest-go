# Audio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**Duration** | **int32** | Duration of the audio in seconds as defined by the sender | 
**Performer** | Pointer to **string** | *Optional*. Performer of the audio as defined by the sender or by audio tags | [optional] 
**Title** | Pointer to **string** | *Optional*. Title of the audio as defined by the sender or by audio tags | [optional] 
**FileName** | Pointer to **string** | *Optional*. Original filename as defined by the sender | [optional] 
**MimeType** | Pointer to **string** | *Optional*. MIME type of the file as defined by the sender | [optional] 
**FileSize** | Pointer to **int32** | *Optional*. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value. | [optional] 
**Thumbnail** | Pointer to [**PhotoSize**](PhotoSize.md) |  | [optional] 

## Methods

### NewAudio

`func NewAudio(fileId string, fileUniqueId string, duration int32, ) *Audio`

NewAudio instantiates a new Audio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioWithDefaults

`func NewAudioWithDefaults() *Audio`

NewAudioWithDefaults instantiates a new Audio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *Audio) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Audio) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Audio) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *Audio) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *Audio) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *Audio) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetDuration

`func (o *Audio) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Audio) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Audio) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetPerformer

`func (o *Audio) GetPerformer() string`

GetPerformer returns the Performer field if non-nil, zero value otherwise.

### GetPerformerOk

`func (o *Audio) GetPerformerOk() (*string, bool)`

GetPerformerOk returns a tuple with the Performer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformer

`func (o *Audio) SetPerformer(v string)`

SetPerformer sets Performer field to given value.

### HasPerformer

`func (o *Audio) HasPerformer() bool`

HasPerformer returns a boolean if a field has been set.

### GetTitle

`func (o *Audio) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Audio) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Audio) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Audio) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFileName

`func (o *Audio) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Audio) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Audio) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Audio) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetMimeType

`func (o *Audio) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Audio) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Audio) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Audio) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileSize

`func (o *Audio) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Audio) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Audio) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *Audio) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetThumbnail

`func (o *Audio) GetThumbnail() PhotoSize`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Audio) GetThumbnailOk() (*PhotoSize, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Audio) SetThumbnail(v PhotoSize)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Audio) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


