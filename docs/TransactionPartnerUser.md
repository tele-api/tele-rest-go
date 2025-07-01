# TransactionPartnerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the transaction partner, always “user” | [default to "user"]
**TransactionType** | **string** | Type of the transaction, currently one of “invoice\\_payment” for payments via invoices, “paid\\_media\\_payment” for payments for paid media, “gift\\_purchase” for gifts sent by the bot, “premium\\_purchase” for Telegram Premium subscriptions gifted by the bot, “business\\_account\\_transfer” for direct transfers from managed business accounts | 
**User** | [**User**](User.md) |  | 
**Affiliate** | Pointer to [**AffiliateInfo**](AffiliateInfo.md) |  | [optional] 
**InvoicePayload** | Pointer to **string** | *Optional*. Bot-specified invoice payload. Can be available only for “invoice\\_payment” transactions. | [optional] 
**SubscriptionPeriod** | Pointer to **int32** | *Optional*. The duration of the paid subscription. Can be available only for “invoice\\_payment” transactions. | [optional] 
**PaidMedia** | Pointer to [**[]PaidMedia**](PaidMedia.md) | *Optional*. Information about the paid media bought by the user; for “paid\\_media\\_payment” transactions only | [optional] 
**PaidMediaPayload** | Pointer to **string** | *Optional*. Bot-specified paid media payload. Can be available only for “paid\\_media\\_payment” transactions. | [optional] 
**Gift** | Pointer to [**Gift**](Gift.md) |  | [optional] 
**PremiumSubscriptionDuration** | Pointer to **int32** | *Optional*. Number of months the gifted Telegram Premium subscription will be active for; for “premium\\_purchase” transactions only | [optional] 

## Methods

### NewTransactionPartnerUser

`func NewTransactionPartnerUser(type_ string, transactionType string, user User, ) *TransactionPartnerUser`

NewTransactionPartnerUser instantiates a new TransactionPartnerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPartnerUserWithDefaults

`func NewTransactionPartnerUserWithDefaults() *TransactionPartnerUser`

NewTransactionPartnerUserWithDefaults instantiates a new TransactionPartnerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPartnerUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPartnerUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPartnerUser) SetType(v string)`

SetType sets Type field to given value.


### GetTransactionType

`func (o *TransactionPartnerUser) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionPartnerUser) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionPartnerUser) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetUser

`func (o *TransactionPartnerUser) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransactionPartnerUser) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransactionPartnerUser) SetUser(v User)`

SetUser sets User field to given value.


### GetAffiliate

`func (o *TransactionPartnerUser) GetAffiliate() AffiliateInfo`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *TransactionPartnerUser) GetAffiliateOk() (*AffiliateInfo, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *TransactionPartnerUser) SetAffiliate(v AffiliateInfo)`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *TransactionPartnerUser) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.

### GetInvoicePayload

`func (o *TransactionPartnerUser) GetInvoicePayload() string`

GetInvoicePayload returns the InvoicePayload field if non-nil, zero value otherwise.

### GetInvoicePayloadOk

`func (o *TransactionPartnerUser) GetInvoicePayloadOk() (*string, bool)`

GetInvoicePayloadOk returns a tuple with the InvoicePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayload

`func (o *TransactionPartnerUser) SetInvoicePayload(v string)`

SetInvoicePayload sets InvoicePayload field to given value.

### HasInvoicePayload

`func (o *TransactionPartnerUser) HasInvoicePayload() bool`

HasInvoicePayload returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *TransactionPartnerUser) GetSubscriptionPeriod() int32`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *TransactionPartnerUser) GetSubscriptionPeriodOk() (*int32, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *TransactionPartnerUser) SetSubscriptionPeriod(v int32)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *TransactionPartnerUser) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetPaidMedia

`func (o *TransactionPartnerUser) GetPaidMedia() []PaidMedia`

GetPaidMedia returns the PaidMedia field if non-nil, zero value otherwise.

### GetPaidMediaOk

`func (o *TransactionPartnerUser) GetPaidMediaOk() (*[]PaidMedia, bool)`

GetPaidMediaOk returns a tuple with the PaidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMedia

`func (o *TransactionPartnerUser) SetPaidMedia(v []PaidMedia)`

SetPaidMedia sets PaidMedia field to given value.

### HasPaidMedia

`func (o *TransactionPartnerUser) HasPaidMedia() bool`

HasPaidMedia returns a boolean if a field has been set.

### GetPaidMediaPayload

`func (o *TransactionPartnerUser) GetPaidMediaPayload() string`

GetPaidMediaPayload returns the PaidMediaPayload field if non-nil, zero value otherwise.

### GetPaidMediaPayloadOk

`func (o *TransactionPartnerUser) GetPaidMediaPayloadOk() (*string, bool)`

GetPaidMediaPayloadOk returns a tuple with the PaidMediaPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMediaPayload

`func (o *TransactionPartnerUser) SetPaidMediaPayload(v string)`

SetPaidMediaPayload sets PaidMediaPayload field to given value.

### HasPaidMediaPayload

`func (o *TransactionPartnerUser) HasPaidMediaPayload() bool`

HasPaidMediaPayload returns a boolean if a field has been set.

### GetGift

`func (o *TransactionPartnerUser) GetGift() Gift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *TransactionPartnerUser) GetGiftOk() (*Gift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *TransactionPartnerUser) SetGift(v Gift)`

SetGift sets Gift field to given value.

### HasGift

`func (o *TransactionPartnerUser) HasGift() bool`

HasGift returns a boolean if a field has been set.

### GetPremiumSubscriptionDuration

`func (o *TransactionPartnerUser) GetPremiumSubscriptionDuration() int32`

GetPremiumSubscriptionDuration returns the PremiumSubscriptionDuration field if non-nil, zero value otherwise.

### GetPremiumSubscriptionDurationOk

`func (o *TransactionPartnerUser) GetPremiumSubscriptionDurationOk() (*int32, bool)`

GetPremiumSubscriptionDurationOk returns a tuple with the PremiumSubscriptionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSubscriptionDuration

`func (o *TransactionPartnerUser) SetPremiumSubscriptionDuration(v int32)`

SetPremiumSubscriptionDuration sets PremiumSubscriptionDuration field to given value.

### HasPremiumSubscriptionDuration

`func (o *TransactionPartnerUser) HasPremiumSubscriptionDuration() bool`

HasPremiumSubscriptionDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


