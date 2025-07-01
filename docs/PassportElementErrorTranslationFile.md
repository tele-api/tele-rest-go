# PassportElementErrorTranslationFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *translation\\_file* | [default to "translation_file"]
**Type** | **string** | Type of element of the user&#39;s Telegram Passport which has the issue, one of “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration”, “temporary\\_registration” | 
**FileHash** | **string** | Base64-encoded file hash | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorTranslationFile

`func NewPassportElementErrorTranslationFile(source string, type_ string, fileHash string, message string, ) *PassportElementErrorTranslationFile`

NewPassportElementErrorTranslationFile instantiates a new PassportElementErrorTranslationFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorTranslationFileWithDefaults

`func NewPassportElementErrorTranslationFileWithDefaults() *PassportElementErrorTranslationFile`

NewPassportElementErrorTranslationFileWithDefaults instantiates a new PassportElementErrorTranslationFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorTranslationFile) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorTranslationFile) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorTranslationFile) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorTranslationFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorTranslationFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorTranslationFile) SetType(v string)`

SetType sets Type field to given value.


### GetFileHash

`func (o *PassportElementErrorTranslationFile) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *PassportElementErrorTranslationFile) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *PassportElementErrorTranslationFile) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.


### GetMessage

`func (o *PassportElementErrorTranslationFile) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorTranslationFile) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorTranslationFile) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


