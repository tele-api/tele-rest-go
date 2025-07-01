# PassportElementError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *unspecified* | [default to "unspecified"]
**Type** | **string** | Type of element of the user&#39;s Telegram Passport which has the issue | 
**FieldName** | **string** | Name of the data field which has the error | 
**DataHash** | **string** | Base64-encoded data hash | 
**Message** | **string** | Error message | 
**FileHash** | **string** | Base64-encoded file hash | 
**FileHashes** | **[]string** | List of base64-encoded file hashes | 
**ElementHash** | **string** | Base64-encoded element hash | 

## Methods

### NewPassportElementError

`func NewPassportElementError(source string, type_ string, fieldName string, dataHash string, message string, fileHash string, fileHashes []string, elementHash string, ) *PassportElementError`

NewPassportElementError instantiates a new PassportElementError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorWithDefaults

`func NewPassportElementErrorWithDefaults() *PassportElementError`

NewPassportElementErrorWithDefaults instantiates a new PassportElementError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementError) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementError) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementError) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementError) SetType(v string)`

SetType sets Type field to given value.


### GetFieldName

`func (o *PassportElementError) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *PassportElementError) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *PassportElementError) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetDataHash

`func (o *PassportElementError) GetDataHash() string`

GetDataHash returns the DataHash field if non-nil, zero value otherwise.

### GetDataHashOk

`func (o *PassportElementError) GetDataHashOk() (*string, bool)`

GetDataHashOk returns a tuple with the DataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHash

`func (o *PassportElementError) SetDataHash(v string)`

SetDataHash sets DataHash field to given value.


### GetMessage

`func (o *PassportElementError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFileHash

`func (o *PassportElementError) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *PassportElementError) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *PassportElementError) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.


### GetFileHashes

`func (o *PassportElementError) GetFileHashes() []string`

GetFileHashes returns the FileHashes field if non-nil, zero value otherwise.

### GetFileHashesOk

`func (o *PassportElementError) GetFileHashesOk() (*[]string, bool)`

GetFileHashesOk returns a tuple with the FileHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHashes

`func (o *PassportElementError) SetFileHashes(v []string)`

SetFileHashes sets FileHashes field to given value.


### GetElementHash

`func (o *PassportElementError) GetElementHash() string`

GetElementHash returns the ElementHash field if non-nil, zero value otherwise.

### GetElementHashOk

`func (o *PassportElementError) GetElementHashOk() (*string, bool)`

GetElementHashOk returns a tuple with the ElementHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementHash

`func (o *PassportElementError) SetElementHash(v string)`

SetElementHash sets ElementHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


