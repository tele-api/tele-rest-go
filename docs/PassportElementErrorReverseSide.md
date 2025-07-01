# PassportElementErrorReverseSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *reverse\\_side* | [default to "reverse_side"]
**Type** | **string** | The section of the user&#39;s Telegram Passport which has the issue, one of “driver\\_license”, “identity\\_card” | 
**FileHash** | **string** | Base64-encoded hash of the file with the reverse side of the document | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorReverseSide

`func NewPassportElementErrorReverseSide(source string, type_ string, fileHash string, message string, ) *PassportElementErrorReverseSide`

NewPassportElementErrorReverseSide instantiates a new PassportElementErrorReverseSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorReverseSideWithDefaults

`func NewPassportElementErrorReverseSideWithDefaults() *PassportElementErrorReverseSide`

NewPassportElementErrorReverseSideWithDefaults instantiates a new PassportElementErrorReverseSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorReverseSide) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorReverseSide) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorReverseSide) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorReverseSide) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorReverseSide) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorReverseSide) SetType(v string)`

SetType sets Type field to given value.


### GetFileHash

`func (o *PassportElementErrorReverseSide) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *PassportElementErrorReverseSide) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *PassportElementErrorReverseSide) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.


### GetMessage

`func (o *PassportElementErrorReverseSide) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorReverseSide) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorReverseSide) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


