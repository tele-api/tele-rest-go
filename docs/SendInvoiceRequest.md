# SendInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread (topic) of the forum; for forum supergroups only | [optional] 
**Title** | **string** | Product name, 1-32 characters | 
**Description** | **string** | Product description, 1-255 characters | 
**Payload** | **string** | Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes. | 
**ProviderToken** | Pointer to **string** | Payment provider token, obtained via [@BotFather](https://t.me/botfather). Pass an empty string for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] 
**Currency** | **string** | Three-letter ISO 4217 currency code, see [more on currencies](https://core.telegram.org/bots/payments#supported-currencies). Pass “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90). | 
**Prices** | [**[]LabeledPrice**](LabeledPrice.md) | Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in [Telegram Stars](https://t.me/BotNews/90). | 
**MaxTipAmount** | Pointer to **int32** | The maximum accepted amount for tips in the *smallest units* of the currency (integer, **not** float/double). For example, for a maximum tip of &#x60;US$ 1.45&#x60; pass &#x60;max_tip_amount &#x3D; 145&#x60;. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in [Telegram Stars](https://t.me/BotNews/90). | [optional] [default to 0]
**SuggestedTipAmounts** | Pointer to **[]int32** | A JSON-serialized array of suggested amounts of tips in the *smallest units* of the currency (integer, **not** float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed *max\\_tip\\_amount*. | [optional] 
**StartParameter** | Pointer to **string** | Unique deep-linking parameter. If left empty, **forwarded copies** of the sent message will have a *Pay* button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a *URL* button with a deep link to the bot (instead of a *Pay* button), with the value used as the start parameter | [optional] 
**ProviderData** | Pointer to **string** | JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider. | [optional] 
**PhotoUrl** | Pointer to **string** | URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for. | [optional] 
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
**DisableNotification** | Pointer to **bool** | Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound. | [optional] 
**ProtectContent** | Pointer to **bool** | Protects the contents of the sent message from forwarding and saving | [optional] 
**AllowPaidBroadcast** | Pointer to **bool** | Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot&#39;s balance | [optional] 
**MessageEffectId** | Pointer to **string** | Unique identifier of the message effect to be added to the message; for private chats only | [optional] 
**ReplyParameters** | Pointer to [**ReplyParameters**](ReplyParameters.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewSendInvoiceRequest

`func NewSendInvoiceRequest(chatId SendMessageRequestChatId, title string, description string, payload string, currency string, prices []LabeledPrice, ) *SendInvoiceRequest`

NewSendInvoiceRequest instantiates a new SendInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInvoiceRequestWithDefaults

`func NewSendInvoiceRequestWithDefaults() *SendInvoiceRequest`

NewSendInvoiceRequestWithDefaults instantiates a new SendInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SendInvoiceRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendInvoiceRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendInvoiceRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendInvoiceRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendInvoiceRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendInvoiceRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendInvoiceRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetTitle

`func (o *SendInvoiceRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendInvoiceRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendInvoiceRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SendInvoiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SendInvoiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SendInvoiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPayload

`func (o *SendInvoiceRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendInvoiceRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendInvoiceRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetProviderToken

`func (o *SendInvoiceRequest) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *SendInvoiceRequest) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *SendInvoiceRequest) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.

### HasProviderToken

`func (o *SendInvoiceRequest) HasProviderToken() bool`

HasProviderToken returns a boolean if a field has been set.

### GetCurrency

`func (o *SendInvoiceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SendInvoiceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SendInvoiceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPrices

`func (o *SendInvoiceRequest) GetPrices() []LabeledPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SendInvoiceRequest) GetPricesOk() (*[]LabeledPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SendInvoiceRequest) SetPrices(v []LabeledPrice)`

SetPrices sets Prices field to given value.


### GetMaxTipAmount

`func (o *SendInvoiceRequest) GetMaxTipAmount() int32`

GetMaxTipAmount returns the MaxTipAmount field if non-nil, zero value otherwise.

### GetMaxTipAmountOk

`func (o *SendInvoiceRequest) GetMaxTipAmountOk() (*int32, bool)`

GetMaxTipAmountOk returns a tuple with the MaxTipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTipAmount

`func (o *SendInvoiceRequest) SetMaxTipAmount(v int32)`

SetMaxTipAmount sets MaxTipAmount field to given value.

### HasMaxTipAmount

`func (o *SendInvoiceRequest) HasMaxTipAmount() bool`

HasMaxTipAmount returns a boolean if a field has been set.

### GetSuggestedTipAmounts

`func (o *SendInvoiceRequest) GetSuggestedTipAmounts() []int32`

GetSuggestedTipAmounts returns the SuggestedTipAmounts field if non-nil, zero value otherwise.

### GetSuggestedTipAmountsOk

`func (o *SendInvoiceRequest) GetSuggestedTipAmountsOk() (*[]int32, bool)`

GetSuggestedTipAmountsOk returns a tuple with the SuggestedTipAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTipAmounts

`func (o *SendInvoiceRequest) SetSuggestedTipAmounts(v []int32)`

SetSuggestedTipAmounts sets SuggestedTipAmounts field to given value.

### HasSuggestedTipAmounts

`func (o *SendInvoiceRequest) HasSuggestedTipAmounts() bool`

HasSuggestedTipAmounts returns a boolean if a field has been set.

### GetStartParameter

`func (o *SendInvoiceRequest) GetStartParameter() string`

GetStartParameter returns the StartParameter field if non-nil, zero value otherwise.

### GetStartParameterOk

`func (o *SendInvoiceRequest) GetStartParameterOk() (*string, bool)`

GetStartParameterOk returns a tuple with the StartParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartParameter

`func (o *SendInvoiceRequest) SetStartParameter(v string)`

SetStartParameter sets StartParameter field to given value.

### HasStartParameter

`func (o *SendInvoiceRequest) HasStartParameter() bool`

HasStartParameter returns a boolean if a field has been set.

### GetProviderData

`func (o *SendInvoiceRequest) GetProviderData() string`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *SendInvoiceRequest) GetProviderDataOk() (*string, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *SendInvoiceRequest) SetProviderData(v string)`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *SendInvoiceRequest) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *SendInvoiceRequest) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *SendInvoiceRequest) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *SendInvoiceRequest) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *SendInvoiceRequest) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetPhotoSize

`func (o *SendInvoiceRequest) GetPhotoSize() int32`

GetPhotoSize returns the PhotoSize field if non-nil, zero value otherwise.

### GetPhotoSizeOk

`func (o *SendInvoiceRequest) GetPhotoSizeOk() (*int32, bool)`

GetPhotoSizeOk returns a tuple with the PhotoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoSize

`func (o *SendInvoiceRequest) SetPhotoSize(v int32)`

SetPhotoSize sets PhotoSize field to given value.

### HasPhotoSize

`func (o *SendInvoiceRequest) HasPhotoSize() bool`

HasPhotoSize returns a boolean if a field has been set.

### GetPhotoWidth

`func (o *SendInvoiceRequest) GetPhotoWidth() int32`

GetPhotoWidth returns the PhotoWidth field if non-nil, zero value otherwise.

### GetPhotoWidthOk

`func (o *SendInvoiceRequest) GetPhotoWidthOk() (*int32, bool)`

GetPhotoWidthOk returns a tuple with the PhotoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoWidth

`func (o *SendInvoiceRequest) SetPhotoWidth(v int32)`

SetPhotoWidth sets PhotoWidth field to given value.

### HasPhotoWidth

`func (o *SendInvoiceRequest) HasPhotoWidth() bool`

HasPhotoWidth returns a boolean if a field has been set.

### GetPhotoHeight

`func (o *SendInvoiceRequest) GetPhotoHeight() int32`

GetPhotoHeight returns the PhotoHeight field if non-nil, zero value otherwise.

### GetPhotoHeightOk

`func (o *SendInvoiceRequest) GetPhotoHeightOk() (*int32, bool)`

GetPhotoHeightOk returns a tuple with the PhotoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoHeight

`func (o *SendInvoiceRequest) SetPhotoHeight(v int32)`

SetPhotoHeight sets PhotoHeight field to given value.

### HasPhotoHeight

`func (o *SendInvoiceRequest) HasPhotoHeight() bool`

HasPhotoHeight returns a boolean if a field has been set.

### GetNeedName

`func (o *SendInvoiceRequest) GetNeedName() bool`

GetNeedName returns the NeedName field if non-nil, zero value otherwise.

### GetNeedNameOk

`func (o *SendInvoiceRequest) GetNeedNameOk() (*bool, bool)`

GetNeedNameOk returns a tuple with the NeedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedName

`func (o *SendInvoiceRequest) SetNeedName(v bool)`

SetNeedName sets NeedName field to given value.

### HasNeedName

`func (o *SendInvoiceRequest) HasNeedName() bool`

HasNeedName returns a boolean if a field has been set.

### GetNeedPhoneNumber

`func (o *SendInvoiceRequest) GetNeedPhoneNumber() bool`

GetNeedPhoneNumber returns the NeedPhoneNumber field if non-nil, zero value otherwise.

### GetNeedPhoneNumberOk

`func (o *SendInvoiceRequest) GetNeedPhoneNumberOk() (*bool, bool)`

GetNeedPhoneNumberOk returns a tuple with the NeedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedPhoneNumber

`func (o *SendInvoiceRequest) SetNeedPhoneNumber(v bool)`

SetNeedPhoneNumber sets NeedPhoneNumber field to given value.

### HasNeedPhoneNumber

`func (o *SendInvoiceRequest) HasNeedPhoneNumber() bool`

HasNeedPhoneNumber returns a boolean if a field has been set.

### GetNeedEmail

`func (o *SendInvoiceRequest) GetNeedEmail() bool`

GetNeedEmail returns the NeedEmail field if non-nil, zero value otherwise.

### GetNeedEmailOk

`func (o *SendInvoiceRequest) GetNeedEmailOk() (*bool, bool)`

GetNeedEmailOk returns a tuple with the NeedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedEmail

`func (o *SendInvoiceRequest) SetNeedEmail(v bool)`

SetNeedEmail sets NeedEmail field to given value.

### HasNeedEmail

`func (o *SendInvoiceRequest) HasNeedEmail() bool`

HasNeedEmail returns a boolean if a field has been set.

### GetNeedShippingAddress

`func (o *SendInvoiceRequest) GetNeedShippingAddress() bool`

GetNeedShippingAddress returns the NeedShippingAddress field if non-nil, zero value otherwise.

### GetNeedShippingAddressOk

`func (o *SendInvoiceRequest) GetNeedShippingAddressOk() (*bool, bool)`

GetNeedShippingAddressOk returns a tuple with the NeedShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedShippingAddress

`func (o *SendInvoiceRequest) SetNeedShippingAddress(v bool)`

SetNeedShippingAddress sets NeedShippingAddress field to given value.

### HasNeedShippingAddress

`func (o *SendInvoiceRequest) HasNeedShippingAddress() bool`

HasNeedShippingAddress returns a boolean if a field has been set.

### GetSendPhoneNumberToProvider

`func (o *SendInvoiceRequest) GetSendPhoneNumberToProvider() bool`

GetSendPhoneNumberToProvider returns the SendPhoneNumberToProvider field if non-nil, zero value otherwise.

### GetSendPhoneNumberToProviderOk

`func (o *SendInvoiceRequest) GetSendPhoneNumberToProviderOk() (*bool, bool)`

GetSendPhoneNumberToProviderOk returns a tuple with the SendPhoneNumberToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPhoneNumberToProvider

`func (o *SendInvoiceRequest) SetSendPhoneNumberToProvider(v bool)`

SetSendPhoneNumberToProvider sets SendPhoneNumberToProvider field to given value.

### HasSendPhoneNumberToProvider

`func (o *SendInvoiceRequest) HasSendPhoneNumberToProvider() bool`

HasSendPhoneNumberToProvider returns a boolean if a field has been set.

### GetSendEmailToProvider

`func (o *SendInvoiceRequest) GetSendEmailToProvider() bool`

GetSendEmailToProvider returns the SendEmailToProvider field if non-nil, zero value otherwise.

### GetSendEmailToProviderOk

`func (o *SendInvoiceRequest) GetSendEmailToProviderOk() (*bool, bool)`

GetSendEmailToProviderOk returns a tuple with the SendEmailToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmailToProvider

`func (o *SendInvoiceRequest) SetSendEmailToProvider(v bool)`

SetSendEmailToProvider sets SendEmailToProvider field to given value.

### HasSendEmailToProvider

`func (o *SendInvoiceRequest) HasSendEmailToProvider() bool`

HasSendEmailToProvider returns a boolean if a field has been set.

### GetIsFlexible

`func (o *SendInvoiceRequest) GetIsFlexible() bool`

GetIsFlexible returns the IsFlexible field if non-nil, zero value otherwise.

### GetIsFlexibleOk

`func (o *SendInvoiceRequest) GetIsFlexibleOk() (*bool, bool)`

GetIsFlexibleOk returns a tuple with the IsFlexible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlexible

`func (o *SendInvoiceRequest) SetIsFlexible(v bool)`

SetIsFlexible sets IsFlexible field to given value.

### HasIsFlexible

`func (o *SendInvoiceRequest) HasIsFlexible() bool`

HasIsFlexible returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendInvoiceRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendInvoiceRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendInvoiceRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendInvoiceRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendInvoiceRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendInvoiceRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendInvoiceRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendInvoiceRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendInvoiceRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendInvoiceRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendInvoiceRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendInvoiceRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendInvoiceRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendInvoiceRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendInvoiceRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendInvoiceRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendInvoiceRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendInvoiceRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendInvoiceRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendInvoiceRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendInvoiceRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendInvoiceRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendInvoiceRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendInvoiceRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


