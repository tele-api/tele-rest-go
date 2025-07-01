# EncryptedCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Base64-encoded encrypted JSON-serialized data with unique user&#39;s payload, data hashes and secrets required for [EncryptedPassportElement](https://core.telegram.org/bots/api/#encryptedpassportelement) decryption and authentication | 
**Hash** | **string** | Base64-encoded data hash for data authentication | 
**Secret** | **string** | Base64-encoded secret, encrypted with the bot&#39;s public RSA key, required for data decryption | 

## Methods

### NewEncryptedCredentials

`func NewEncryptedCredentials(data string, hash string, secret string, ) *EncryptedCredentials`

NewEncryptedCredentials instantiates a new EncryptedCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptedCredentialsWithDefaults

`func NewEncryptedCredentialsWithDefaults() *EncryptedCredentials`

NewEncryptedCredentialsWithDefaults instantiates a new EncryptedCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EncryptedCredentials) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EncryptedCredentials) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EncryptedCredentials) SetData(v string)`

SetData sets Data field to given value.


### GetHash

`func (o *EncryptedCredentials) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *EncryptedCredentials) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *EncryptedCredentials) SetHash(v string)`

SetHash sets Hash field to given value.


### GetSecret

`func (o *EncryptedCredentials) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EncryptedCredentials) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EncryptedCredentials) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


