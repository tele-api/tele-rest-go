# AffiliateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffiliateUser** | Pointer to [**User**](User.md) |  | [optional] 
**AffiliateChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**CommissionPerMille** | **int32** | The number of Telegram Stars received by the affiliate for each 1000 Telegram Stars received by the bot from referred users | 
**Amount** | **int32** | Integer amount of Telegram Stars received by the affiliate from the transaction, rounded to 0; can be negative for refunds | 
**NanostarAmount** | Pointer to **int32** | *Optional*. The number of 1/1000000000 shares of Telegram Stars received by the affiliate; from -999999999 to 999999999; can be negative for refunds | [optional] 

## Methods

### NewAffiliateInfo

`func NewAffiliateInfo(commissionPerMille int32, amount int32, ) *AffiliateInfo`

NewAffiliateInfo instantiates a new AffiliateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateInfoWithDefaults

`func NewAffiliateInfoWithDefaults() *AffiliateInfo`

NewAffiliateInfoWithDefaults instantiates a new AffiliateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliateUser

`func (o *AffiliateInfo) GetAffiliateUser() User`

GetAffiliateUser returns the AffiliateUser field if non-nil, zero value otherwise.

### GetAffiliateUserOk

`func (o *AffiliateInfo) GetAffiliateUserOk() (*User, bool)`

GetAffiliateUserOk returns a tuple with the AffiliateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateUser

`func (o *AffiliateInfo) SetAffiliateUser(v User)`

SetAffiliateUser sets AffiliateUser field to given value.

### HasAffiliateUser

`func (o *AffiliateInfo) HasAffiliateUser() bool`

HasAffiliateUser returns a boolean if a field has been set.

### GetAffiliateChat

`func (o *AffiliateInfo) GetAffiliateChat() Chat`

GetAffiliateChat returns the AffiliateChat field if non-nil, zero value otherwise.

### GetAffiliateChatOk

`func (o *AffiliateInfo) GetAffiliateChatOk() (*Chat, bool)`

GetAffiliateChatOk returns a tuple with the AffiliateChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateChat

`func (o *AffiliateInfo) SetAffiliateChat(v Chat)`

SetAffiliateChat sets AffiliateChat field to given value.

### HasAffiliateChat

`func (o *AffiliateInfo) HasAffiliateChat() bool`

HasAffiliateChat returns a boolean if a field has been set.

### GetCommissionPerMille

`func (o *AffiliateInfo) GetCommissionPerMille() int32`

GetCommissionPerMille returns the CommissionPerMille field if non-nil, zero value otherwise.

### GetCommissionPerMilleOk

`func (o *AffiliateInfo) GetCommissionPerMilleOk() (*int32, bool)`

GetCommissionPerMilleOk returns a tuple with the CommissionPerMille field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPerMille

`func (o *AffiliateInfo) SetCommissionPerMille(v int32)`

SetCommissionPerMille sets CommissionPerMille field to given value.


### GetAmount

`func (o *AffiliateInfo) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AffiliateInfo) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AffiliateInfo) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetNanostarAmount

`func (o *AffiliateInfo) GetNanostarAmount() int32`

GetNanostarAmount returns the NanostarAmount field if non-nil, zero value otherwise.

### GetNanostarAmountOk

`func (o *AffiliateInfo) GetNanostarAmountOk() (*int32, bool)`

GetNanostarAmountOk returns a tuple with the NanostarAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanostarAmount

`func (o *AffiliateInfo) SetNanostarAmount(v int32)`

SetNanostarAmount sets NanostarAmount field to given value.

### HasNanostarAmount

`func (o *AffiliateInfo) HasNanostarAmount() bool`

HasNanostarAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


