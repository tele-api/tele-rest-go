# PassportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EncryptedPassportElement**](EncryptedPassportElement.md) | Array with information about documents and other Telegram Passport elements that was shared with the bot | 
**Credentials** | [**EncryptedCredentials**](EncryptedCredentials.md) |  | 

## Methods

### NewPassportData

`func NewPassportData(data []EncryptedPassportElement, credentials EncryptedCredentials, ) *PassportData`

NewPassportData instantiates a new PassportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassportDataWithDefaults

`func NewPassportDataWithDefaults() *PassportData`

NewPassportDataWithDefaults instantiates a new PassportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PassportData) GetData() []EncryptedPassportElement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PassportData) GetDataOk() (*[]EncryptedPassportElement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PassportData) SetData(v []EncryptedPassportElement)`

SetData sets Data field to given value.


### GetCredentials

`func (o *PassportData) GetCredentials() EncryptedCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PassportData) GetCredentialsOk() (*EncryptedCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PassportData) SetCredentials(v EncryptedCredentials)`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


