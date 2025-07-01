# PassportElementErrorFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *files* | [default to "files"]
**Type** | **string** | The section of the user&#39;s Telegram Passport which has the issue, one of “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration”, “temporary\\_registration” | 
**FileHashes** | **[]string** | List of base64-encoded file hashes | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorFiles

`func NewPassportElementErrorFiles(source string, type_ string, fileHashes []string, message string, ) *PassportElementErrorFiles`

NewPassportElementErrorFiles instantiates a new PassportElementErrorFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorFilesWithDefaults

`func NewPassportElementErrorFilesWithDefaults() *PassportElementErrorFiles`

NewPassportElementErrorFilesWithDefaults instantiates a new PassportElementErrorFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorFiles) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorFiles) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorFiles) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorFiles) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorFiles) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorFiles) SetType(v string)`

SetType sets Type field to given value.


### GetFileHashes

`func (o *PassportElementErrorFiles) GetFileHashes() []string`

GetFileHashes returns the FileHashes field if non-nil, zero value otherwise.

### GetFileHashesOk

`func (o *PassportElementErrorFiles) GetFileHashesOk() (*[]string, bool)`

GetFileHashesOk returns a tuple with the FileHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHashes

`func (o *PassportElementErrorFiles) SetFileHashes(v []string)`

SetFileHashes sets FileHashes field to given value.


### GetMessage

`func (o *PassportElementErrorFiles) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorFiles) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorFiles) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


