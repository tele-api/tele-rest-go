# PreCheckoutQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique query identifier | 
**From** | [**User**](User.md) |  | 
**Currency** | **string** | Three-letter ISO 4217 [currency](https://core.telegram.org/bots/payments#supported-currencies) code, or “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90) | 
**TotalAmount** | **int32** | Total price in the *smallest units* of the currency (integer, **not** float/double). For example, for a price of &#x60;US$ 1.45&#x60; pass &#x60;amount &#x3D; 145&#x60;. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). | 
**InvoicePayload** | **string** | Bot-specified invoice payload | 
**ShippingOptionId** | Pointer to **string** | *Optional*. Identifier of the shipping option chosen by the user | [optional] 
**OrderInfo** | Pointer to [**OrderInfo**](OrderInfo.md) |  | [optional] 

## Methods

### NewPreCheckoutQuery

`func NewPreCheckoutQuery(id string, from User, currency string, totalAmount int32, invoicePayload string, ) *PreCheckoutQuery`

NewPreCheckoutQuery instantiates a new PreCheckoutQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreCheckoutQueryWithDefaults

`func NewPreCheckoutQueryWithDefaults() *PreCheckoutQuery`

NewPreCheckoutQueryWithDefaults instantiates a new PreCheckoutQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PreCheckoutQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreCheckoutQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreCheckoutQuery) SetId(v string)`

SetId sets Id field to given value.


### GetFrom

`func (o *PreCheckoutQuery) GetFrom() User`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PreCheckoutQuery) GetFromOk() (*User, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PreCheckoutQuery) SetFrom(v User)`

SetFrom sets From field to given value.


### GetCurrency

`func (o *PreCheckoutQuery) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PreCheckoutQuery) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PreCheckoutQuery) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalAmount

`func (o *PreCheckoutQuery) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PreCheckoutQuery) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PreCheckoutQuery) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetInvoicePayload

`func (o *PreCheckoutQuery) GetInvoicePayload() string`

GetInvoicePayload returns the InvoicePayload field if non-nil, zero value otherwise.

### GetInvoicePayloadOk

`func (o *PreCheckoutQuery) GetInvoicePayloadOk() (*string, bool)`

GetInvoicePayloadOk returns a tuple with the InvoicePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayload

`func (o *PreCheckoutQuery) SetInvoicePayload(v string)`

SetInvoicePayload sets InvoicePayload field to given value.


### GetShippingOptionId

`func (o *PreCheckoutQuery) GetShippingOptionId() string`

GetShippingOptionId returns the ShippingOptionId field if non-nil, zero value otherwise.

### GetShippingOptionIdOk

`func (o *PreCheckoutQuery) GetShippingOptionIdOk() (*string, bool)`

GetShippingOptionIdOk returns a tuple with the ShippingOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptionId

`func (o *PreCheckoutQuery) SetShippingOptionId(v string)`

SetShippingOptionId sets ShippingOptionId field to given value.

### HasShippingOptionId

`func (o *PreCheckoutQuery) HasShippingOptionId() bool`

HasShippingOptionId returns a boolean if a field has been set.

### GetOrderInfo

`func (o *PreCheckoutQuery) GetOrderInfo() OrderInfo`

GetOrderInfo returns the OrderInfo field if non-nil, zero value otherwise.

### GetOrderInfoOk

`func (o *PreCheckoutQuery) GetOrderInfoOk() (*OrderInfo, bool)`

GetOrderInfoOk returns a tuple with the OrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderInfo

`func (o *PreCheckoutQuery) SetOrderInfo(v OrderInfo)`

SetOrderInfo sets OrderInfo field to given value.

### HasOrderInfo

`func (o *PreCheckoutQuery) HasOrderInfo() bool`

HasOrderInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


