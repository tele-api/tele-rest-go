# TransactionPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the transaction partner, always “other” | [default to "other"]
**TransactionType** | **string** | Type of the transaction, currently one of “invoice\\_payment” for payments via invoices, “paid\\_media\\_payment” for payments for paid media, “gift\\_purchase” for gifts sent by the bot, “premium\\_purchase” for Telegram Premium subscriptions gifted by the bot, “business\\_account\\_transfer” for direct transfers from managed business accounts | 
**User** | [**User**](User.md) |  | 
**Affiliate** | Pointer to [**AffiliateInfo**](AffiliateInfo.md) |  | [optional] 
**InvoicePayload** | Pointer to **string** | *Optional*. Bot-specified invoice payload. Can be available only for “invoice\\_payment” transactions. | [optional] 
**SubscriptionPeriod** | Pointer to **int32** | *Optional*. The duration of the paid subscription. Can be available only for “invoice\\_payment” transactions. | [optional] 
**PaidMedia** | Pointer to [**[]PaidMedia**](PaidMedia.md) | *Optional*. Information about the paid media bought by the user; for “paid\\_media\\_payment” transactions only | [optional] 
**PaidMediaPayload** | Pointer to **string** | *Optional*. Bot-specified paid media payload. Can be available only for “paid\\_media\\_payment” transactions. | [optional] 
**Gift** | Pointer to [**Gift**](Gift.md) |  | [optional] 
**PremiumSubscriptionDuration** | Pointer to **int32** | *Optional*. Number of months the gifted Telegram Premium subscription will be active for; for “premium\\_purchase” transactions only | [optional] 
**Chat** | [**Chat**](Chat.md) |  | 
**SponsorUser** | Pointer to [**User**](User.md) |  | [optional] 
**CommissionPerMille** | **int32** | The number of Telegram Stars received by the bot for each 1000 Telegram Stars received by the affiliate program sponsor from referred users | 
**WithdrawalState** | Pointer to [**RevenueWithdrawalState**](RevenueWithdrawalState.md) |  | [optional] 
**RequestCount** | **int32** | The number of successful requests that exceeded regular limits and were therefore billed | 

## Methods

### NewTransactionPartner

`func NewTransactionPartner(type_ string, transactionType string, user User, chat Chat, commissionPerMille int32, requestCount int32, ) *TransactionPartner`

NewTransactionPartner instantiates a new TransactionPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPartnerWithDefaults

`func NewTransactionPartnerWithDefaults() *TransactionPartner`

NewTransactionPartnerWithDefaults instantiates a new TransactionPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPartner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPartner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPartner) SetType(v string)`

SetType sets Type field to given value.


### GetTransactionType

`func (o *TransactionPartner) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionPartner) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionPartner) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetUser

`func (o *TransactionPartner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransactionPartner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransactionPartner) SetUser(v User)`

SetUser sets User field to given value.


### GetAffiliate

`func (o *TransactionPartner) GetAffiliate() AffiliateInfo`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *TransactionPartner) GetAffiliateOk() (*AffiliateInfo, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *TransactionPartner) SetAffiliate(v AffiliateInfo)`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *TransactionPartner) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.

### GetInvoicePayload

`func (o *TransactionPartner) GetInvoicePayload() string`

GetInvoicePayload returns the InvoicePayload field if non-nil, zero value otherwise.

### GetInvoicePayloadOk

`func (o *TransactionPartner) GetInvoicePayloadOk() (*string, bool)`

GetInvoicePayloadOk returns a tuple with the InvoicePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayload

`func (o *TransactionPartner) SetInvoicePayload(v string)`

SetInvoicePayload sets InvoicePayload field to given value.

### HasInvoicePayload

`func (o *TransactionPartner) HasInvoicePayload() bool`

HasInvoicePayload returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *TransactionPartner) GetSubscriptionPeriod() int32`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *TransactionPartner) GetSubscriptionPeriodOk() (*int32, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *TransactionPartner) SetSubscriptionPeriod(v int32)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *TransactionPartner) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetPaidMedia

`func (o *TransactionPartner) GetPaidMedia() []PaidMedia`

GetPaidMedia returns the PaidMedia field if non-nil, zero value otherwise.

### GetPaidMediaOk

`func (o *TransactionPartner) GetPaidMediaOk() (*[]PaidMedia, bool)`

GetPaidMediaOk returns a tuple with the PaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMedia

`func (o *TransactionPartner) SetPaidMedia(v []PaidMedia)`

SetPaidMedia sets PaidMedia field to given value.

### HasPaidMedia

`func (o *TransactionPartner) HasPaidMedia() bool`

HasPaidMedia returns a boolean if a field has been set.

### GetPaidMediaPayload

`func (o *TransactionPartner) GetPaidMediaPayload() string`

GetPaidMediaPayload returns the PaidMediaPayload field if non-nil, zero value otherwise.

### GetPaidMediaPayloadOk

`func (o *TransactionPartner) GetPaidMediaPayloadOk() (*string, bool)`

GetPaidMediaPayloadOk returns a tuple with the PaidMediaPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMediaPayload

`func (o *TransactionPartner) SetPaidMediaPayload(v string)`

SetPaidMediaPayload sets PaidMediaPayload field to given value.

### HasPaidMediaPayload

`func (o *TransactionPartner) HasPaidMediaPayload() bool`

HasPaidMediaPayload returns a boolean if a field has been set.

### GetGift

`func (o *TransactionPartner) GetGift() Gift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *TransactionPartner) GetGiftOk() (*Gift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *TransactionPartner) SetGift(v Gift)`

SetGift sets Gift field to given value.

### HasGift

`func (o *TransactionPartner) HasGift() bool`

HasGift returns a boolean if a field has been set.

### GetPremiumSubscriptionDuration

`func (o *TransactionPartner) GetPremiumSubscriptionDuration() int32`

GetPremiumSubscriptionDuration returns the PremiumSubscriptionDuration field if non-nil, zero value otherwise.

### GetPremiumSubscriptionDurationOk

`func (o *TransactionPartner) GetPremiumSubscriptionDurationOk() (*int32, bool)`

GetPremiumSubscriptionDurationOk returns a tuple with the PremiumSubscriptionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionDuration

`func (o *TransactionPartner) SetPremiumSubscriptionDuration(v int32)`

SetPremiumSubscriptionDuration sets PremiumSubscriptionDuration field to given value.

### HasPremiumSubscriptionDuration

`func (o *TransactionPartner) HasPremiumSubscriptionDuration() bool`

HasPremiumSubscriptionDuration returns a boolean if a field has been set.

### GetChat

`func (o *TransactionPartner) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *TransactionPartner) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *TransactionPartner) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetSponsorUser

`func (o *TransactionPartner) GetSponsorUser() User`

GetSponsorUser returns the SponsorUser field if non-nil, zero value otherwise.

### GetSponsorUserOk

`func (o *TransactionPartner) GetSponsorUserOk() (*User, bool)`

GetSponsorUserOk returns a tuple with the SponsorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorUser

`func (o *TransactionPartner) SetSponsorUser(v User)`

SetSponsorUser sets SponsorUser field to given value.

### HasSponsorUser

`func (o *TransactionPartner) HasSponsorUser() bool`

HasSponsorUser returns a boolean if a field has been set.

### GetCommissionPerMille

`func (o *TransactionPartner) GetCommissionPerMille() int32`

GetCommissionPerMille returns the CommissionPerMille field if non-nil, zero value otherwise.

### GetCommissionPerMilleOk

`func (o *TransactionPartner) GetCommissionPerMilleOk() (*int32, bool)`

GetCommissionPerMilleOk returns a tuple with the CommissionPerMille field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPerMille

`func (o *TransactionPartner) SetCommissionPerMille(v int32)`

SetCommissionPerMille sets CommissionPerMille field to given value.


### GetWithdrawalState

`func (o *TransactionPartner) GetWithdrawalState() RevenueWithdrawalState`

GetWithdrawalState returns the WithdrawalState field if non-nil, zero value otherwise.

### GetWithdrawalStateOk

`func (o *TransactionPartner) GetWithdrawalStateOk() (*RevenueWithdrawalState, bool)`

GetWithdrawalStateOk returns a tuple with the WithdrawalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalState

`func (o *TransactionPartner) SetWithdrawalState(v RevenueWithdrawalState)`

SetWithdrawalState sets WithdrawalState field to given value.

### HasWithdrawalState

`func (o *TransactionPartner) HasWithdrawalState() bool`

HasWithdrawalState returns a boolean if a field has been set.

### GetRequestCount

`func (o *TransactionPartner) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *TransactionPartner) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *TransactionPartner) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


