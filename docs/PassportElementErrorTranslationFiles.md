# PassportElementErrorTranslationFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *translation\\_files* | [default to "translation_files"]
**Type** | **string** | Type of element of the user&#39;s Telegram Passport which has the issue, one of “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration”, “temporary\\_registration” | 
**FileHashes** | **[]string** | List of base64-encoded file hashes | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorTranslationFiles

`func NewPassportElementErrorTranslationFiles(source string, type_ string, fileHashes []string, message string, ) *PassportElementErrorTranslationFiles`

NewPassportElementErrorTranslationFiles instantiates a new PassportElementErrorTranslationFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorTranslationFilesWithDefaults

`func NewPassportElementErrorTranslationFilesWithDefaults() *PassportElementErrorTranslationFiles`

NewPassportElementErrorTranslationFilesWithDefaults instantiates a new PassportElementErrorTranslationFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorTranslationFiles) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorTranslationFiles) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorTranslationFiles) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorTranslationFiles) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorTranslationFiles) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorTranslationFiles) SetType(v string)`

SetType sets Type field to given value.


### GetFileHashes

`func (o *PassportElementErrorTranslationFiles) GetFileHashes() []string`

GetFileHashes returns the FileHashes field if non-nil, zero value otherwise.

### GetFileHashesOk

`func (o *PassportElementErrorTranslationFiles) GetFileHashesOk() (*[]string, bool)`

GetFileHashesOk returns a tuple with the FileHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHashes

`func (o *PassportElementErrorTranslationFiles) SetFileHashes(v []string)`

SetFileHashes sets FileHashes field to given value.


### GetMessage

`func (o *PassportElementErrorTranslationFiles) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorTranslationFiles) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorTranslationFiles) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


