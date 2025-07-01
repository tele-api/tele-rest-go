# PassportElementErrorUnspecified

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *unspecified* | [default to "unspecified"]
**Type** | **string** | Type of element of the user&#39;s Telegram Passport which has the issue | 
**ElementHash** | **string** | Base64-encoded element hash | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorUnspecified

`func NewPassportElementErrorUnspecified(source string, type_ string, elementHash string, message string, ) *PassportElementErrorUnspecified`

NewPassportElementErrorUnspecified instantiates a new PassportElementErrorUnspecified object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorUnspecifiedWithDefaults

`func NewPassportElementErrorUnspecifiedWithDefaults() *PassportElementErrorUnspecified`

NewPassportElementErrorUnspecifiedWithDefaults instantiates a new PassportElementErrorUnspecified object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorUnspecified) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorUnspecified) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorUnspecified) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorUnspecified) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorUnspecified) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorUnspecified) SetType(v string)`

SetType sets Type field to given value.


### GetElementHash

`func (o *PassportElementErrorUnspecified) GetElementHash() string`

GetElementHash returns the ElementHash field if non-nil, zero value otherwise.

### GetElementHashOk

`func (o *PassportElementErrorUnspecified) GetElementHashOk() (*string, bool)`

GetElementHashOk returns a tuple with the ElementHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementHash

`func (o *PassportElementErrorUnspecified) SetElementHash(v string)`

SetElementHash sets ElementHash field to given value.


### GetMessage

`func (o *PassportElementErrorUnspecified) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorUnspecified) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorUnspecified) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


