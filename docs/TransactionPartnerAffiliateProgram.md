# TransactionPartnerAffiliateProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the transaction partner, always “affiliate\\_program” | [default to "affiliate_program"]
**SponsorUser** | Pointer to [**User**](User.md) |  | [optional] 
**CommissionPerMille** | **int32** | The number of Telegram Stars received by the bot for each 1000 Telegram Stars received by the affiliate program sponsor from referred users | 

## Methods

### NewTransactionPartnerAffiliateProgram

`func NewTransactionPartnerAffiliateProgram(type_ string, commissionPerMille int32, ) *TransactionPartnerAffiliateProgram`

NewTransactionPartnerAffiliateProgram instantiates a new TransactionPartnerAffiliateProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPartnerAffiliateProgramWithDefaults

`func NewTransactionPartnerAffiliateProgramWithDefaults() *TransactionPartnerAffiliateProgram`

NewTransactionPartnerAffiliateProgramWithDefaults instantiates a new TransactionPartnerAffiliateProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPartnerAffiliateProgram) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPartnerAffiliateProgram) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPartnerAffiliateProgram) SetType(v string)`

SetType sets Type field to given value.


### GetSponsorUser

`func (o *TransactionPartnerAffiliateProgram) GetSponsorUser() User`

GetSponsorUser returns the SponsorUser field if non-nil, zero value otherwise.

### GetSponsorUserOk

`func (o *TransactionPartnerAffiliateProgram) GetSponsorUserOk() (*User, bool)`

GetSponsorUserOk returns a tuple with the SponsorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorUser

`func (o *TransactionPartnerAffiliateProgram) SetSponsorUser(v User)`

SetSponsorUser sets SponsorUser field to given value.

### HasSponsorUser

`func (o *TransactionPartnerAffiliateProgram) HasSponsorUser() bool`

HasSponsorUser returns a boolean if a field has been set.

### GetCommissionPerMille

`func (o *TransactionPartnerAffiliateProgram) GetCommissionPerMille() int32`

GetCommissionPerMille returns the CommissionPerMille field if non-nil, zero value otherwise.

### GetCommissionPerMilleOk

`func (o *TransactionPartnerAffiliateProgram) GetCommissionPerMilleOk() (*int32, bool)`

GetCommissionPerMilleOk returns a tuple with the CommissionPerMille field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPerMille

`func (o *TransactionPartnerAffiliateProgram) SetCommissionPerMille(v int32)`

SetCommissionPerMille sets CommissionPerMille field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


