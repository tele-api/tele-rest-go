# CreateInvoiceLinkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the link will be created. For payments in [Telegram Stars](https://t.me/BotNews/90) only. | [optional] 
**Title** | **string** | Product name, 1-32 characters | 
**Description** | **string** | Product description, 1-255 characters | 
**Payload** | **string** | Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes. | 
**ProviderToken** | Pointer to **string** | Payment provider token, obtained via [@BotFather](https://t.me/botfather). Pass an empty string for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**Currency** | **string** | Three-letter ISO 4217 currency code, see [more on currencies](https://core.telegram.org/bots/payments#supported-currencies). Pass “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90). | 
**Prices** | [**[]LabeledPrice**](LabeledPrice.md) | Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in [Telegram Stars](https://t.me/BotNews/90). | 
**SubscriptionPeriod** | Pointer to **int32** | The number of seconds the subscription will be active for before the next payment. The currency must be set to “XTR” (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 10000 Telegram Stars. | [optional] 
**MaxTipAmount** | Pointer to **int32** | The maximum accepted amount for tips in the *smallest units* of the currency (integer, **not** float/double). For example, for a maximum tip of &#x60;US$ 1.45&#x60; pass &#x60;max_tip_amount &#x3D; 145&#x60;. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] [default to 0]
**SuggestedTipAmounts** | Pointer to **[]int32** | A JSON-serialized array of suggested amounts of tips in the *smallest units* of the currency (integer, **not** float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed *max\\_tip\\_amount*. | [optional] 
**ProviderData** | Pointer to **string** | JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider. | [optional] 
**PhotoUrl** | Pointer to **string** | URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. | [optional] 
**PhotoSize** | Pointer to **int32** | Photo size in bytes | [optional] 
**PhotoWidth** | Pointer to **int32** | Photo width | [optional] 
**PhotoHeight** | Pointer to **int32** | Photo height | [optional] 
**NeedName** | Pointer to **bool** | Pass *True* if you require the user&#39;s full name to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**NeedPhoneNumber** | Pointer to **bool** | Pass *True* if you require the user&#39;s phone number to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**NeedEmail** | Pointer to **bool** | Pass *True* if you require the user&#39;s email address to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**NeedShippingAddress** | Pointer to **bool** | Pass *True* if you require the user&#39;s shipping address to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**SendPhoneNumberToProvider** | Pointer to **bool** | Pass *True* if the user&#39;s phone number should be sent to the provider. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**SendEmailToProvider** | Pointer to **bool** | Pass *True* if the user&#39;s email address should be sent to the provider. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**IsFlexible** | Pointer to **bool** | Pass *True* if the final price depends on the shipping method. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 

## Methods

### NewCreateInvoiceLinkPostRequest

`func NewCreateInvoiceLinkPostRequest(title string, description string, payload string, currency string, prices []LabeledPrice, ) *CreateInvoiceLinkPostRequest`

NewCreateInvoiceLinkPostRequest instantiates a new CreateInvoiceLinkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceLinkPostRequestWithDefaults

`func NewCreateInvoiceLinkPostRequestWithDefaults() *CreateInvoiceLinkPostRequest`

NewCreateInvoiceLinkPostRequestWithDefaults instantiates a new CreateInvoiceLinkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *CreateInvoiceLinkPostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *CreateInvoiceLinkPostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *CreateInvoiceLinkPostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *CreateInvoiceLinkPostRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateInvoiceLinkPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateInvoiceLinkPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateInvoiceLinkPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateInvoiceLinkPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInvoiceLinkPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInvoiceLinkPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPayload

`func (o *CreateInvoiceLinkPostRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateInvoiceLinkPostRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateInvoiceLinkPostRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetProviderToken

`func (o *CreateInvoiceLinkPostRequest) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *CreateInvoiceLinkPostRequest) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *CreateInvoiceLinkPostRequest) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.

### HasProviderToken

`func (o *CreateInvoiceLinkPostRequest) HasProviderToken() bool`

HasProviderToken returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateInvoiceLinkPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateInvoiceLinkPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateInvoiceLinkPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPrices

`func (o *CreateInvoiceLinkPostRequest) GetPrices() []LabeledPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *CreateInvoiceLinkPostRequest) GetPricesOk() (*[]LabeledPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *CreateInvoiceLinkPostRequest) SetPrices(v []LabeledPrice)`

SetPrices sets Prices field to given value.


### GetSubscriptionPeriod

`func (o *CreateInvoiceLinkPostRequest) GetSubscriptionPeriod() int32`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *CreateInvoiceLinkPostRequest) GetSubscriptionPeriodOk() (*int32, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *CreateInvoiceLinkPostRequest) SetSubscriptionPeriod(v int32)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *CreateInvoiceLinkPostRequest) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetMaxTipAmount

`func (o *CreateInvoiceLinkPostRequest) GetMaxTipAmount() int32`

GetMaxTipAmount returns the MaxTipAmount field if non-nil, zero value otherwise.

### GetMaxTipAmountOk

`func (o *CreateInvoiceLinkPostRequest) GetMaxTipAmountOk() (*int32, bool)`

GetMaxTipAmountOk returns a tuple with the MaxTipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTipAmount

`func (o *CreateInvoiceLinkPostRequest) SetMaxTipAmount(v int32)`

SetMaxTipAmount sets MaxTipAmount field to given value.

### HasMaxTipAmount

`func (o *CreateInvoiceLinkPostRequest) HasMaxTipAmount() bool`

HasMaxTipAmount returns a boolean if a field has been set.

### GetSuggestedTipAmounts

`func (o *CreateInvoiceLinkPostRequest) GetSuggestedTipAmounts() []int32`

GetSuggestedTipAmounts returns the SuggestedTipAmounts field if non-nil, zero value otherwise.

### GetSuggestedTipAmountsOk

`func (o *CreateInvoiceLinkPostRequest) GetSuggestedTipAmountsOk() (*[]int32, bool)`

GetSuggestedTipAmountsOk returns a tuple with the SuggestedTipAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTipAmounts

`func (o *CreateInvoiceLinkPostRequest) SetSuggestedTipAmounts(v []int32)`

SetSuggestedTipAmounts sets SuggestedTipAmounts field to given value.

### HasSuggestedTipAmounts

`func (o *CreateInvoiceLinkPostRequest) HasSuggestedTipAmounts() bool`

HasSuggestedTipAmounts returns a boolean if a field has been set.

### GetProviderData

`func (o *CreateInvoiceLinkPostRequest) GetProviderData() string`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *CreateInvoiceLinkPostRequest) GetProviderDataOk() (*string, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *CreateInvoiceLinkPostRequest) SetProviderData(v string)`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *CreateInvoiceLinkPostRequest) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *CreateInvoiceLinkPostRequest) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *CreateInvoiceLinkPostRequest) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *CreateInvoiceLinkPostRequest) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *CreateInvoiceLinkPostRequest) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetPhotoSize

`func (o *CreateInvoiceLinkPostRequest) GetPhotoSize() int32`

GetPhotoSize returns the PhotoSize field if non-nil, zero value otherwise.

### GetPhotoSizeOk

`func (o *CreateInvoiceLinkPostRequest) GetPhotoSizeOk() (*int32, bool)`

GetPhotoSizeOk returns a tuple with the PhotoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoSize

`func (o *CreateInvoiceLinkPostRequest) SetPhotoSize(v int32)`

SetPhotoSize sets PhotoSize field to given value.

### HasPhotoSize

`func (o *CreateInvoiceLinkPostRequest) HasPhotoSize() bool`

HasPhotoSize returns a boolean if a field has been set.

### GetPhotoWidth

`func (o *CreateInvoiceLinkPostRequest) GetPhotoWidth() int32`

GetPhotoWidth returns the PhotoWidth field if non-nil, zero value otherwise.

### GetPhotoWidthOk

`func (o *CreateInvoiceLinkPostRequest) GetPhotoWidthOk() (*int32, bool)`

GetPhotoWidthOk returns a tuple with the PhotoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoWidth

`func (o *CreateInvoiceLinkPostRequest) SetPhotoWidth(v int32)`

SetPhotoWidth sets PhotoWidth field to given value.

### HasPhotoWidth

`func (o *CreateInvoiceLinkPostRequest) HasPhotoWidth() bool`

HasPhotoWidth returns a boolean if a field has been set.

### GetPhotoHeight

`func (o *CreateInvoiceLinkPostRequest) GetPhotoHeight() int32`

GetPhotoHeight returns the PhotoHeight field if non-nil, zero value otherwise.

### GetPhotoHeightOk

`func (o *CreateInvoiceLinkPostRequest) GetPhotoHeightOk() (*int32, bool)`

GetPhotoHeightOk returns a tuple with the PhotoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoHeight

`func (o *CreateInvoiceLinkPostRequest) SetPhotoHeight(v int32)`

SetPhotoHeight sets PhotoHeight field to given value.

### HasPhotoHeight

`func (o *CreateInvoiceLinkPostRequest) HasPhotoHeight() bool`

HasPhotoHeight returns a boolean if a field has been set.

### GetNeedName

`func (o *CreateInvoiceLinkPostRequest) GetNeedName() bool`

GetNeedName returns the NeedName field if non-nil, zero value otherwise.

### GetNeedNameOk

`func (o *CreateInvoiceLinkPostRequest) GetNeedNameOk() (*bool, bool)`

GetNeedNameOk returns a tuple with the NeedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedName

`func (o *CreateInvoiceLinkPostRequest) SetNeedName(v bool)`

SetNeedName sets NeedName field to given value.

### HasNeedName

`func (o *CreateInvoiceLinkPostRequest) HasNeedName() bool`

HasNeedName returns a boolean if a field has been set.

### GetNeedPhoneNumber

`func (o *CreateInvoiceLinkPostRequest) GetNeedPhoneNumber() bool`

GetNeedPhoneNumber returns the NeedPhoneNumber field if non-nil, zero value otherwise.

### GetNeedPhoneNumberOk

`func (o *CreateInvoiceLinkPostRequest) GetNeedPhoneNumberOk() (*bool, bool)`

GetNeedPhoneNumberOk returns a tuple with the NeedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedPhoneNumber

`func (o *CreateInvoiceLinkPostRequest) SetNeedPhoneNumber(v bool)`

SetNeedPhoneNumber sets NeedPhoneNumber field to given value.

### HasNeedPhoneNumber

`func (o *CreateInvoiceLinkPostRequest) HasNeedPhoneNumber() bool`

HasNeedPhoneNumber returns a boolean if a field has been set.

### GetNeedEmail

`func (o *CreateInvoiceLinkPostRequest) GetNeedEmail() bool`

GetNeedEmail returns the NeedEmail field if non-nil, zero value otherwise.

### GetNeedEmailOk

`func (o *CreateInvoiceLinkPostRequest) GetNeedEmailOk() (*bool, bool)`

GetNeedEmailOk returns a tuple with the NeedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedEmail

`func (o *CreateInvoiceLinkPostRequest) SetNeedEmail(v bool)`

SetNeedEmail sets NeedEmail field to given value.

### HasNeedEmail

`func (o *CreateInvoiceLinkPostRequest) HasNeedEmail() bool`

HasNeedEmail returns a boolean if a field has been set.

### GetNeedShippingAddress

`func (o *CreateInvoiceLinkPostRequest) GetNeedShippingAddress() bool`

GetNeedShippingAddress returns the NeedShippingAddress field if non-nil, zero value otherwise.

### GetNeedShippingAddressOk

`func (o *CreateInvoiceLinkPostRequest) GetNeedShippingAddressOk() (*bool, bool)`

GetNeedShippingAddressOk returns a tuple with the NeedShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedShippingAddress

`func (o *CreateInvoiceLinkPostRequest) SetNeedShippingAddress(v bool)`

SetNeedShippingAddress sets NeedShippingAddress field to given value.

### HasNeedShippingAddress

`func (o *CreateInvoiceLinkPostRequest) HasNeedShippingAddress() bool`

HasNeedShippingAddress returns a boolean if a field has been set.

### GetSendPhoneNumberToProvider

`func (o *CreateInvoiceLinkPostRequest) GetSendPhoneNumberToProvider() bool`

GetSendPhoneNumberToProvider returns the SendPhoneNumberToProvider field if non-nil, zero value otherwise.

### GetSendPhoneNumberToProviderOk

`func (o *CreateInvoiceLinkPostRequest) GetSendPhoneNumberToProviderOk() (*bool, bool)`

GetSendPhoneNumberToProviderOk returns a tuple with the SendPhoneNumberToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPhoneNumberToProvider

`func (o *CreateInvoiceLinkPostRequest) SetSendPhoneNumberToProvider(v bool)`

SetSendPhoneNumberToProvider sets SendPhoneNumberToProvider field to given value.

### HasSendPhoneNumberToProvider

`func (o *CreateInvoiceLinkPostRequest) HasSendPhoneNumberToProvider() bool`

HasSendPhoneNumberToProvider returns a boolean if a field has been set.

### GetSendEmailToProvider

`func (o *CreateInvoiceLinkPostRequest) GetSendEmailToProvider() bool`

GetSendEmailToProvider returns the SendEmailToProvider field if non-nil, zero value otherwise.

### GetSendEmailToProviderOk

`func (o *CreateInvoiceLinkPostRequest) GetSendEmailToProviderOk() (*bool, bool)`

GetSendEmailToProviderOk returns a tuple with the SendEmailToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmailToProvider

`func (o *CreateInvoiceLinkPostRequest) SetSendEmailToProvider(v bool)`

SetSendEmailToProvider sets SendEmailToProvider field to given value.

### HasSendEmailToProvider

`func (o *CreateInvoiceLinkPostRequest) HasSendEmailToProvider() bool`

HasSendEmailToProvider returns a boolean if a field has been set.

### GetIsFlexible

`func (o *CreateInvoiceLinkPostRequest) GetIsFlexible() bool`

GetIsFlexible returns the IsFlexible field if non-nil, zero value otherwise.

### GetIsFlexibleOk

`func (o *CreateInvoiceLinkPostRequest) GetIsFlexibleOk() (*bool, bool)`

GetIsFlexibleOk returns a tuple with the IsFlexible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlexible

`func (o *CreateInvoiceLinkPostRequest) SetIsFlexible(v bool)`

SetIsFlexible sets IsFlexible field to given value.

### HasIsFlexible

`func (o *CreateInvoiceLinkPostRequest) HasIsFlexible() bool`

HasIsFlexible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


