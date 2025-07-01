# SuccessfulPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code, or “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90) | 
**TotalAmount** | **int32** | Total price in the *smallest units* of the currency (integer, **not** float/double). For example, for a price of &#x60;US$ 1.45&#x60; pass &#x60;amount &#x3D; 145&#x60;. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). | 
**InvoicePayload** | **string** | Bot-specified invoice payload | 
**SubscriptionExpirationDate** | Pointer to **int32** | *Optional*. Expiration date of the subscription, in Unix time; for recurring payments only | [optional] 
**IsRecurring** | Pointer to **bool** | *Optional*. True, if the payment is a recurring payment for a subscription | [optional] [default to true]
**IsFirstRecurring** | Pointer to **bool** | *Optional*. True, if the payment is the first payment for a subscription | [optional] [default to true]
**ShippingOptionId** | Pointer to **string** | *Optional*. Identifier of the shipping option chosen by the user | [optional] 
**OrderInfo** | Pointer to [**OrderInfo**](OrderInfo.md) |  | [optional] 
**TelegramPaymentChargeId** | **string** | Telegram payment identifier | 
**ProviderPaymentChargeId** | **string** | Provider payment identifier | 

## Methods

### NewSuccessfulPayment

`func NewSuccessfulPayment(currency string, totalAmount int32, invoicePayload string, telegramPaymentChargeId string, providerPaymentChargeId string, ) *SuccessfulPayment`

NewSuccessfulPayment instantiates a new SuccessfulPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessfulPaymentWithDefaults

`func NewSuccessfulPaymentWithDefaults() *SuccessfulPayment`

NewSuccessfulPaymentWithDefaults instantiates a new SuccessfulPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SuccessfulPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SuccessfulPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SuccessfulPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalAmount

`func (o *SuccessfulPayment) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *SuccessfulPayment) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *SuccessfulPayment) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetInvoicePayload

`func (o *SuccessfulPayment) GetInvoicePayload() string`

GetInvoicePayload returns the InvoicePayload field if non-nil, zero value otherwise.

### GetInvoicePayloadOk

`func (o *SuccessfulPayment) GetInvoicePayloadOk() (*string, bool)`

GetInvoicePayloadOk returns a tuple with the InvoicePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayload

`func (o *SuccessfulPayment) SetInvoicePayload(v string)`

SetInvoicePayload sets InvoicePayload field to given value.


### GetSubscriptionExpirationDate

`func (o *SuccessfulPayment) GetSubscriptionExpirationDate() int32`

GetSubscriptionExpirationDate returns the SubscriptionExpirationDate field if non-nil, zero value otherwise.

### GetSubscriptionExpirationDateOk

`func (o *SuccessfulPayment) GetSubscriptionExpirationDateOk() (*int32, bool)`

GetSubscriptionExpirationDateOk returns a tuple with the SubscriptionExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpirationDate

`func (o *SuccessfulPayment) SetSubscriptionExpirationDate(v int32)`

SetSubscriptionExpirationDate sets SubscriptionExpirationDate field to given value.

### HasSubscriptionExpirationDate

`func (o *SuccessfulPayment) HasSubscriptionExpirationDate() bool`

HasSubscriptionExpirationDate returns a boolean if a field has been set.

### GetIsRecurring

`func (o *SuccessfulPayment) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *SuccessfulPayment) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *SuccessfulPayment) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *SuccessfulPayment) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### GetIsFirstRecurring

`func (o *SuccessfulPayment) GetIsFirstRecurring() bool`

GetIsFirstRecurring returns the IsFirstRecurring field if non-nil, zero value otherwise.

### GetIsFirstRecurringOk

`func (o *SuccessfulPayment) GetIsFirstRecurringOk() (*bool, bool)`

GetIsFirstRecurringOk returns a tuple with the IsFirstRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstRecurring

`func (o *SuccessfulPayment) SetIsFirstRecurring(v bool)`

SetIsFirstRecurring sets IsFirstRecurring field to given value.

### HasIsFirstRecurring

`func (o *SuccessfulPayment) HasIsFirstRecurring() bool`

HasIsFirstRecurring returns a boolean if a field has been set.

### GetShippingOptionId

`func (o *SuccessfulPayment) GetShippingOptionId() string`

GetShippingOptionId returns the ShippingOptionId field if non-nil, zero value otherwise.

### GetShippingOptionIdOk

`func (o *SuccessfulPayment) GetShippingOptionIdOk() (*string, bool)`

GetShippingOptionIdOk returns a tuple with the ShippingOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptionId

`func (o *SuccessfulPayment) SetShippingOptionId(v string)`

SetShippingOptionId sets ShippingOptionId field to given value.

### HasShippingOptionId

`func (o *SuccessfulPayment) HasShippingOptionId() bool`

HasShippingOptionId returns a boolean if a field has been set.

### GetOrderInfo

`func (o *SuccessfulPayment) GetOrderInfo() OrderInfo`

GetOrderInfo returns the OrderInfo field if non-nil, zero value otherwise.

### GetOrderInfoOk

`func (o *SuccessfulPayment) GetOrderInfoOk() (*OrderInfo, bool)`

GetOrderInfoOk returns a tuple with the OrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderInfo

`func (o *SuccessfulPayment) SetOrderInfo(v OrderInfo)`

SetOrderInfo sets OrderInfo field to given value.

### HasOrderInfo

`func (o *SuccessfulPayment) HasOrderInfo() bool`

HasOrderInfo returns a boolean if a field has been set.

### GetTelegramPaymentChargeId

`func (o *SuccessfulPayment) GetTelegramPaymentChargeId() string`

GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field if non-nil, zero value otherwise.

### GetTelegramPaymentChargeIdOk

`func (o *SuccessfulPayment) GetTelegramPaymentChargeIdOk() (*string, bool)`

GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramPaymentChargeId

`func (o *SuccessfulPayment) SetTelegramPaymentChargeId(v string)`

SetTelegramPaymentChargeId sets TelegramPaymentChargeId field to given value.


### GetProviderPaymentChargeId

`func (o *SuccessfulPayment) GetProviderPaymentChargeId() string`

GetProviderPaymentChargeId returns the ProviderPaymentChargeId field if non-nil, zero value otherwise.

### GetProviderPaymentChargeIdOk

`func (o *SuccessfulPayment) GetProviderPaymentChargeIdOk() (*string, bool)`

GetProviderPaymentChargeIdOk returns a tuple with the ProviderPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderPaymentChargeId

`func (o *SuccessfulPayment) SetProviderPaymentChargeId(v string)`

SetProviderPaymentChargeId sets ProviderPaymentChargeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


