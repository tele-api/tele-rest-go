# PassportElementErrorFrontSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *front\\_side* | [default to "front_side"]
**Type** | **string** | The section of the user&#39;s Telegram Passport which has the issue, one of “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport” | 
**FileHash** | **string** | Base64-encoded hash of the file with the front side of the document | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorFrontSide

`func NewPassportElementErrorFrontSide(source string, type_ string, fileHash string, message string, ) *PassportElementErrorFrontSide`

NewPassportElementErrorFrontSide instantiates a new PassportElementErrorFrontSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorFrontSideWithDefaults

`func NewPassportElementErrorFrontSideWithDefaults() *PassportElementErrorFrontSide`

NewPassportElementErrorFrontSideWithDefaults instantiates a new PassportElementErrorFrontSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorFrontSide) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorFrontSide) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorFrontSide) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorFrontSide) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorFrontSide) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorFrontSide) SetType(v string)`

SetType sets Type field to given value.


### GetFileHash

`func (o *PassportElementErrorFrontSide) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *PassportElementErrorFrontSide) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *PassportElementErrorFrontSide) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.


### GetMessage

`func (o *PassportElementErrorFrontSide) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorFrontSide) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorFrontSide) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


