# PassportFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**FileSize** | **int32** | File size in bytes | 
**FileDate** | **int32** | Unix time when the file was uploaded | 

## Methods

### NewPassportFile

`func NewPassportFile(fileId string, fileUniqueId string, fileSize int32, fileDate int32, ) *PassportFile`

NewPassportFile instantiates a new PassportFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportFileWithDefaults

`func NewPassportFileWithDefaults() *PassportFile`

NewPassportFileWithDefaults instantiates a new PassportFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *PassportFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PassportFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PassportFile) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *PassportFile) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *PassportFile) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *PassportFile) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetFileSize

`func (o *PassportFile) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *PassportFile) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *PassportFile) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetFileDate

`func (o *PassportFile) GetFileDate() int32`

GetFileDate returns the FileDate field if non-nil, zero value otherwise.

### GetFileDateOk

`func (o *PassportFile) GetFileDateOk() (*int32, bool)`

GetFileDateOk returns a tuple with the FileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDate

`func (o *PassportFile) SetFileDate(v int32)`

SetFileDate sets FileDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


