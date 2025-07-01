# PassportElementErrorDataField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Error source, must be *data* | [default to "data"]
**Type** | **string** | The section of the user&#39;s Telegram Passport which has the error, one of “personal\\_details”, “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “address” | 
**FieldName** | **string** | Name of the data field which has the error | 
**DataHash** | **string** | Base64-encoded data hash | 
**Message** | **string** | Error message | 

## Methods

### NewPassportElementErrorDataField

`func NewPassportElementErrorDataField(source string, type_ string, fieldName string, dataHash string, message string, ) *PassportElementErrorDataField`

NewPassportElementErrorDataField instantiates a new PassportElementErrorDataField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportElementErrorDataFieldWithDefaults

`func NewPassportElementErrorDataFieldWithDefaults() *PassportElementErrorDataField`

NewPassportElementErrorDataFieldWithDefaults instantiates a new PassportElementErrorDataField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PassportElementErrorDataField) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PassportElementErrorDataField) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PassportElementErrorDataField) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *PassportElementErrorDataField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassportElementErrorDataField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassportElementErrorDataField) SetType(v string)`

SetType sets Type field to given value.


### GetFieldName

`func (o *PassportElementErrorDataField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *PassportElementErrorDataField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *PassportElementErrorDataField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetDataHash

`func (o *PassportElementErrorDataField) GetDataHash() string`

GetDataHash returns the DataHash field if non-nil, zero value otherwise.

### GetDataHashOk

`func (o *PassportElementErrorDataField) GetDataHashOk() (*string, bool)`

GetDataHashOk returns a tuple with the DataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHash

`func (o *PassportElementErrorDataField) SetDataHash(v string)`

SetDataHash sets DataHash field to given value.


### GetMessage

`func (o *PassportElementErrorDataField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PassportElementErrorDataField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PassportElementErrorDataField) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


