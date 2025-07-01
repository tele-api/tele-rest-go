# RefundedPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code, or “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90). Currently, always “XTR” | [default to "XTR"]
**TotalAmount** | **int32** | Total refunded price in the *smallest units* of the currency (integer, **not** float/double). For example, for a price of &#x60;US$ 1.45&#x60;, &#x60;total_amount &#x3D; 145&#x60;. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). | 
**InvoicePayload** | **string** | Bot-specified invoice payload | 
**TelegramPaymentChargeId** | **string** | Telegram payment identifier | 
**ProviderPaymentChargeId** | Pointer to **string** | *Optional*. Provider payment identifier | [optional] 

## Methods

### NewRefundedPayment

`func NewRefundedPayment(currency string, totalAmount int32, invoicePayload string, telegramPaymentChargeId string, ) *RefundedPayment`

NewRefundedPayment instantiates a new RefundedPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundedPaymentWithDefaults

`func NewRefundedPaymentWithDefaults() *RefundedPayment`

NewRefundedPaymentWithDefaults instantiates a new RefundedPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *RefundedPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RefundedPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RefundedPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalAmount

`func (o *RefundedPayment) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *RefundedPayment) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *RefundedPayment) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetInvoicePayload

`func (o *RefundedPayment) GetInvoicePayload() string`

GetInvoicePayload returns the InvoicePayload field if non-nil, zero value otherwise.

### GetInvoicePayloadOk

`func (o *RefundedPayment) GetInvoicePayloadOk() (*string, bool)`

GetInvoicePayloadOk returns a tuple with the InvoicePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayload

`func (o *RefundedPayment) SetInvoicePayload(v string)`

SetInvoicePayload sets InvoicePayload field to given value.


### GetTelegramPaymentChargeId

`func (o *RefundedPayment) GetTelegramPaymentChargeId() string`

GetTelegramPaymentChargeId returns the TelegramPaymentChargeId field if non-nil, zero value otherwise.

### GetTelegramPaymentChargeIdOk

`func (o *RefundedPayment) GetTelegramPaymentChargeIdOk() (*string, bool)`

GetTelegramPaymentChargeIdOk returns a tuple with the TelegramPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramPaymentChargeId

`func (o *RefundedPayment) SetTelegramPaymentChargeId(v string)`

SetTelegramPaymentChargeId sets TelegramPaymentChargeId field to given value.


### GetProviderPaymentChargeId

`func (o *RefundedPayment) GetProviderPaymentChargeId() string`

GetProviderPaymentChargeId returns the ProviderPaymentChargeId field if non-nil, zero value otherwise.

### GetProviderPaymentChargeIdOk

`func (o *RefundedPayment) GetProviderPaymentChargeIdOk() (*string, bool)`

GetProviderPaymentChargeIdOk returns a tuple with the ProviderPaymentChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderPaymentChargeId

`func (o *RefundedPayment) SetProviderPaymentChargeId(v string)`

SetProviderPaymentChargeId sets ProviderPaymentChargeId field to given value.

### HasProviderPaymentChargeId

`func (o *RefundedPayment) HasProviderPaymentChargeId() bool`

HasProviderPaymentChargeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


