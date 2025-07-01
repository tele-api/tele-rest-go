# PassportElementErrorSelfie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *selfie* | [default to "selfie"]
**Type** | **string** | The section of the user&#39;s Telegram Passport which has the issue, one of “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport” | 
**FileHash** | **string** | Base64-encoded hash of the file with the selfie | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorSelfie

`func NewPassportElementErrorSelfie(source string, type_ string, fileHash string, message string, ) *PassportElementErrorSelfie`

NewPassportElementErrorSelfie instantiates a new PassportElementErrorSelfie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorSelfieWithDefaults

`func NewPassportElementErrorSelfieWithDefaults() *PassportElementErrorSelfie`

NewPassportElementErrorSelfieWithDefaults instantiates a new PassportElementErrorSelfie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorSelfie) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorSelfie) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorSelfie) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorSelfie) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorSelfie) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorSelfie) SetType(v string)`

SetType sets Type field to given value.


### GetFileHash

`func (o *PassportElementErrorSelfie) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *PassportElementErrorSelfie) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *PassportElementErrorSelfie) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.


### GetMessage

`func (o *PassportElementErrorSelfie) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorSelfie) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorSelfie) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


