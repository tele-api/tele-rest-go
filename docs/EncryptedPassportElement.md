# EncryptedPassportElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Element type. One of “personal\\_details”, “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “address”, “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration”, “temporary\\_registration”, “phone\\_number”, “email”. | 
**Data** | Pointer to **string** | *Optional*. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for “personal\\_details”, “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport” and “address” types. Can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api/#encryptedcredentials). | [optional] 
**PhoneNumber** | Pointer to **string** | *Optional*. User&#39;s verified phone number; available only for “phone\\_number” type | [optional] 
**Email** | Pointer to **string** | *Optional*. User&#39;s verified email address; available only for “email” type | [optional] 
**Files** | Pointer to [**[]PassportFile**](PassportFile.md) | *Optional*. Array of encrypted files with documents provided by the user; available only for “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration” and “temporary\\_registration” types. Files can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api/#encryptedcredentials). | [optional] 
**FrontSide** | Pointer to [**PassportFile**](PassportFile.md) |  | [optional] 
**ReverseSide** | Pointer to [**PassportFile**](PassportFile.md) |  | [optional] 
**Selfie** | Pointer to [**PassportFile**](PassportFile.md) |  | [optional] 
**Translation** | Pointer to [**[]PassportFile**](PassportFile.md) | *Optional*. Array of encrypted files with translated versions of documents provided by the user; available if requested for “passport”, “driver\\_license”, “identity\\_card”, “internal\\_passport”, “utility\\_bill”, “bank\\_statement”, “rental\\_agreement”, “passport\\_registration” and “temporary\\_registration” types. Files can be decrypted and verified using the accompanying [EncryptedCredentials](https://core.telegram.org/bots/api/#encryptedcredentials). | [optional] 
**Hash** | **string** | Base64-encoded element hash for using in [PassportElementErrorUnspecified](https://core.telegram.org/bots/api/#passportelementerrorunspecified) | 

## Methods

### NewEncryptedPassportElement

`func NewEncryptedPassportElement(type_ string, hash string, ) *EncryptedPassportElement`

NewEncryptedPassportElement instantiates a new EncryptedPassportElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptedPassportElementWithDefaults

`func NewEncryptedPassportElementWithDefaults() *EncryptedPassportElement`

NewEncryptedPassportElementWithDefaults instantiates a new EncryptedPassportElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EncryptedPassportElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EncryptedPassportElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EncryptedPassportElement) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *EncryptedPassportElement) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EncryptedPassportElement) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EncryptedPassportElement) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EncryptedPassportElement) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *EncryptedPassportElement) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *EncryptedPassportElement) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *EncryptedPassportElement) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *EncryptedPassportElement) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *EncryptedPassportElement) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EncryptedPassportElement) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EncryptedPassportElement) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EncryptedPassportElement) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFiles

`func (o *EncryptedPassportElement) GetFiles() []PassportFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *EncryptedPassportElement) GetFilesOk() (*[]PassportFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *EncryptedPassportElement) SetFiles(v []PassportFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *EncryptedPassportElement) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFrontSide

`func (o *EncryptedPassportElement) GetFrontSide() PassportFile`

GetFrontSide returns the FrontSide field if non-nil, zero value otherwise.

### GetFrontSideOk

`func (o *EncryptedPassportElement) GetFrontSideOk() (*PassportFile, bool)`

GetFrontSideOk returns a tuple with the FrontSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontSide

`func (o *EncryptedPassportElement) SetFrontSide(v PassportFile)`

SetFrontSide sets FrontSide field to given value.

### HasFrontSide

`func (o *EncryptedPassportElement) HasFrontSide() bool`

HasFrontSide returns a boolean if a field has been set.

### GetReverseSide

`func (o *EncryptedPassportElement) GetReverseSide() PassportFile`

GetReverseSide returns the ReverseSide field if non-nil, zero value otherwise.

### GetReverseSideOk

`func (o *EncryptedPassportElement) GetReverseSideOk() (*PassportFile, bool)`

GetReverseSideOk returns a tuple with the ReverseSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseSide

`func (o *EncryptedPassportElement) SetReverseSide(v PassportFile)`

SetReverseSide sets ReverseSide field to given value.

### HasReverseSide

`func (o *EncryptedPassportElement) HasReverseSide() bool`

HasReverseSide returns a boolean if a field has been set.

### GetSelfie

`func (o *EncryptedPassportElement) GetSelfie() PassportFile`

GetSelfie returns the Selfie field if non-nil, zero value otherwise.

### GetSelfieOk

`func (o *EncryptedPassportElement) GetSelfieOk() (*PassportFile, bool)`

GetSelfieOk returns a tuple with the Selfie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfie

`func (o *EncryptedPassportElement) SetSelfie(v PassportFile)`

SetSelfie sets Selfie field to given value.

### HasSelfie

`func (o *EncryptedPassportElement) HasSelfie() bool`

HasSelfie returns a boolean if a field has been set.

### GetTranslation

`func (o *EncryptedPassportElement) GetTranslation() []PassportFile`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *EncryptedPassportElement) GetTranslationOk() (*[]PassportFile, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *EncryptedPassportElement) SetTranslation(v []PassportFile)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *EncryptedPassportElement) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### GetHash

`func (o *EncryptedPassportElement) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *EncryptedPassportElement) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *EncryptedPassportElement) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


