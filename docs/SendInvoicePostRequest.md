# SendInvoicePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
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

### NewSendInvoicePostRequest

`func NewSendInvoicePostRequest(chatId SendMessagePostRequestChatId, title string, description string, payload string, currency string, prices []LabeledPrice, ) *SendInvoicePostRequest`

NewSendInvoicePostRequest instantiates a new SendInvoicePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInvoicePostRequestWithDefaults

`func NewSendInvoicePostRequestWithDefaults() *SendInvoicePostRequest`

NewSendInvoicePostRequestWithDefaults instantiates a new SendInvoicePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SendInvoicePostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendInvoicePostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendInvoicePostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *SendInvoicePostRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *SendInvoicePostRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *SendInvoicePostRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *SendInvoicePostRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetTitle

`func (o *SendInvoicePostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendInvoicePostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendInvoicePostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SendInvoicePostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SendInvoicePostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SendInvoicePostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPayload

`func (o *SendInvoicePostRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendInvoicePostRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendInvoicePostRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetProviderToken

`func (o *SendInvoicePostRequest) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *SendInvoicePostRequest) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *SendInvoicePostRequest) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.

### HasProviderToken

`func (o *SendInvoicePostRequest) HasProviderToken() bool`

HasProviderToken returns a boolean if a field has been set.

### GetCurrency

`func (o *SendInvoicePostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SendInvoicePostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SendInvoicePostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPrices

`func (o *SendInvoicePostRequest) GetPrices() []LabeledPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SendInvoicePostRequest) GetPricesOk() (*[]LabeledPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SendInvoicePostRequest) SetPrices(v []LabeledPrice)`

SetPrices sets Prices field to given value.


### GetMaxTipAmount

`func (o *SendInvoicePostRequest) GetMaxTipAmount() int32`

GetMaxTipAmount returns the MaxTipAmount field if non-nil, zero value otherwise.

### GetMaxTipAmountOk

`func (o *SendInvoicePostRequest) GetMaxTipAmountOk() (*int32, bool)`

GetMaxTipAmountOk returns a tuple with the MaxTipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTipAmount

`func (o *SendInvoicePostRequest) SetMaxTipAmount(v int32)`

SetMaxTipAmount sets MaxTipAmount field to given value.

### HasMaxTipAmount

`func (o *SendInvoicePostRequest) HasMaxTipAmount() bool`

HasMaxTipAmount returns a boolean if a field has been set.

### GetSuggestedTipAmounts

`func (o *SendInvoicePostRequest) GetSuggestedTipAmounts() []int32`

GetSuggestedTipAmounts returns the SuggestedTipAmounts field if non-nil, zero value otherwise.

### GetSuggestedTipAmountsOk

`func (o *SendInvoicePostRequest) GetSuggestedTipAmountsOk() (*[]int32, bool)`

GetSuggestedTipAmountsOk returns a tuple with the SuggestedTipAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTipAmounts

`func (o *SendInvoicePostRequest) SetSuggestedTipAmounts(v []int32)`

SetSuggestedTipAmounts sets SuggestedTipAmounts field to given value.

### HasSuggestedTipAmounts

`func (o *SendInvoicePostRequest) HasSuggestedTipAmounts() bool`

HasSuggestedTipAmounts returns a boolean if a field has been set.

### GetStartParameter

`func (o *SendInvoicePostRequest) GetStartParameter() string`

GetStartParameter returns the StartParameter field if non-nil, zero value otherwise.

### GetStartParameterOk

`func (o *SendInvoicePostRequest) GetStartParameterOk() (*string, bool)`

GetStartParameterOk returns a tuple with the StartParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartParameter

`func (o *SendInvoicePostRequest) SetStartParameter(v string)`

SetStartParameter sets StartParameter field to given value.

### HasStartParameter

`func (o *SendInvoicePostRequest) HasStartParameter() bool`

HasStartParameter returns a boolean if a field has been set.

### GetProviderData

`func (o *SendInvoicePostRequest) GetProviderData() string`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *SendInvoicePostRequest) GetProviderDataOk() (*string, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *SendInvoicePostRequest) SetProviderData(v string)`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *SendInvoicePostRequest) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *SendInvoicePostRequest) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *SendInvoicePostRequest) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *SendInvoicePostRequest) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *SendInvoicePostRequest) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetPhotoSize

`func (o *SendInvoicePostRequest) GetPhotoSize() int32`

GetPhotoSize returns the PhotoSize field if non-nil, zero value otherwise.

### GetPhotoSizeOk

`func (o *SendInvoicePostRequest) GetPhotoSizeOk() (*int32, bool)`

GetPhotoSizeOk returns a tuple with the PhotoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoSize

`func (o *SendInvoicePostRequest) SetPhotoSize(v int32)`

SetPhotoSize sets PhotoSize field to given value.

### HasPhotoSize

`func (o *SendInvoicePostRequest) HasPhotoSize() bool`

HasPhotoSize returns a boolean if a field has been set.

### GetPhotoWidth

`func (o *SendInvoicePostRequest) GetPhotoWidth() int32`

GetPhotoWidth returns the PhotoWidth field if non-nil, zero value otherwise.

### GetPhotoWidthOk

`func (o *SendInvoicePostRequest) GetPhotoWidthOk() (*int32, bool)`

GetPhotoWidthOk returns a tuple with the PhotoWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoWidth

`func (o *SendInvoicePostRequest) SetPhotoWidth(v int32)`

SetPhotoWidth sets PhotoWidth field to given value.

### HasPhotoWidth

`func (o *SendInvoicePostRequest) HasPhotoWidth() bool`

HasPhotoWidth returns a boolean if a field has been set.

### GetPhotoHeight

`func (o *SendInvoicePostRequest) GetPhotoHeight() int32`

GetPhotoHeight returns the PhotoHeight field if non-nil, zero value otherwise.

### GetPhotoHeightOk

`func (o *SendInvoicePostRequest) GetPhotoHeightOk() (*int32, bool)`

GetPhotoHeightOk returns a tuple with the PhotoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoHeight

`func (o *SendInvoicePostRequest) SetPhotoHeight(v int32)`

SetPhotoHeight sets PhotoHeight field to given value.

### HasPhotoHeight

`func (o *SendInvoicePostRequest) HasPhotoHeight() bool`

HasPhotoHeight returns a boolean if a field has been set.

### GetNeedName

`func (o *SendInvoicePostRequest) GetNeedName() bool`

GetNeedName returns the NeedName field if non-nil, zero value otherwise.

### GetNeedNameOk

`func (o *SendInvoicePostRequest) GetNeedNameOk() (*bool, bool)`

GetNeedNameOk returns a tuple with the NeedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedName

`func (o *SendInvoicePostRequest) SetNeedName(v bool)`

SetNeedName sets NeedName field to given value.

### HasNeedName

`func (o *SendInvoicePostRequest) HasNeedName() bool`

HasNeedName returns a boolean if a field has been set.

### GetNeedPhoneNumber

`func (o *SendInvoicePostRequest) GetNeedPhoneNumber() bool`

GetNeedPhoneNumber returns the NeedPhoneNumber field if non-nil, zero value otherwise.

### GetNeedPhoneNumberOk

`func (o *SendInvoicePostRequest) GetNeedPhoneNumberOk() (*bool, bool)`

GetNeedPhoneNumberOk returns a tuple with the NeedPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedPhoneNumber

`func (o *SendInvoicePostRequest) SetNeedPhoneNumber(v bool)`

SetNeedPhoneNumber sets NeedPhoneNumber field to given value.

### HasNeedPhoneNumber

`func (o *SendInvoicePostRequest) HasNeedPhoneNumber() bool`

HasNeedPhoneNumber returns a boolean if a field has been set.

### GetNeedEmail

`func (o *SendInvoicePostRequest) GetNeedEmail() bool`

GetNeedEmail returns the NeedEmail field if non-nil, zero value otherwise.

### GetNeedEmailOk

`func (o *SendInvoicePostRequest) GetNeedEmailOk() (*bool, bool)`

GetNeedEmailOk returns a tuple with the NeedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedEmail

`func (o *SendInvoicePostRequest) SetNeedEmail(v bool)`

SetNeedEmail sets NeedEmail field to given value.

### HasNeedEmail

`func (o *SendInvoicePostRequest) HasNeedEmail() bool`

HasNeedEmail returns a boolean if a field has been set.

### GetNeedShippingAddress

`func (o *SendInvoicePostRequest) GetNeedShippingAddress() bool`

GetNeedShippingAddress returns the NeedShippingAddress field if non-nil, zero value otherwise.

### GetNeedShippingAddressOk

`func (o *SendInvoicePostRequest) GetNeedShippingAddressOk() (*bool, bool)`

GetNeedShippingAddressOk returns a tuple with the NeedShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedShippingAddress

`func (o *SendInvoicePostRequest) SetNeedShippingAddress(v bool)`

SetNeedShippingAddress sets NeedShippingAddress field to given value.

### HasNeedShippingAddress

`func (o *SendInvoicePostRequest) HasNeedShippingAddress() bool`

HasNeedShippingAddress returns a boolean if a field has been set.

### GetSendPhoneNumberToProvider

`func (o *SendInvoicePostRequest) GetSendPhoneNumberToProvider() bool`

GetSendPhoneNumberToProvider returns the SendPhoneNumberToProvider field if non-nil, zero value otherwise.

### GetSendPhoneNumberToProviderOk

`func (o *SendInvoicePostRequest) GetSendPhoneNumberToProviderOk() (*bool, bool)`

GetSendPhoneNumberToProviderOk returns a tuple with the SendPhoneNumberToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPhoneNumberToProvider

`func (o *SendInvoicePostRequest) SetSendPhoneNumberToProvider(v bool)`

SetSendPhoneNumberToProvider sets SendPhoneNumberToProvider field to given value.

### HasSendPhoneNumberToProvider

`func (o *SendInvoicePostRequest) HasSendPhoneNumberToProvider() bool`

HasSendPhoneNumberToProvider returns a boolean if a field has been set.

### GetSendEmailToProvider

`func (o *SendInvoicePostRequest) GetSendEmailToProvider() bool`

GetSendEmailToProvider returns the SendEmailToProvider field if non-nil, zero value otherwise.

### GetSendEmailToProviderOk

`func (o *SendInvoicePostRequest) GetSendEmailToProviderOk() (*bool, bool)`

GetSendEmailToProviderOk returns a tuple with the SendEmailToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmailToProvider

`func (o *SendInvoicePostRequest) SetSendEmailToProvider(v bool)`

SetSendEmailToProvider sets SendEmailToProvider field to given value.

### HasSendEmailToProvider

`func (o *SendInvoicePostRequest) HasSendEmailToProvider() bool`

HasSendEmailToProvider returns a boolean if a field has been set.

### GetIsFlexible

`func (o *SendInvoicePostRequest) GetIsFlexible() bool`

GetIsFlexible returns the IsFlexible field if non-nil, zero value otherwise.

### GetIsFlexibleOk

`func (o *SendInvoicePostRequest) GetIsFlexibleOk() (*bool, bool)`

GetIsFlexibleOk returns a tuple with the IsFlexible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlexible

`func (o *SendInvoicePostRequest) SetIsFlexible(v bool)`

SetIsFlexible sets IsFlexible field to given value.

### HasIsFlexible

`func (o *SendInvoicePostRequest) HasIsFlexible() bool`

HasIsFlexible returns a boolean if a field has been set.

### GetDisableNotification

`func (o *SendInvoicePostRequest) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *SendInvoicePostRequest) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *SendInvoicePostRequest) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *SendInvoicePostRequest) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *SendInvoicePostRequest) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *SendInvoicePostRequest) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *SendInvoicePostRequest) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *SendInvoicePostRequest) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetAllowPaidBroadcast

`func (o *SendInvoicePostRequest) GetAllowPaidBroadcast() bool`

GetAllowPaidBroadcast returns the AllowPaidBroadcast field if non-nil, zero value otherwise.

### GetAllowPaidBroadcastOk

`func (o *SendInvoicePostRequest) GetAllowPaidBroadcastOk() (*bool, bool)`

GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaidBroadcast

`func (o *SendInvoicePostRequest) SetAllowPaidBroadcast(v bool)`

SetAllowPaidBroadcast sets AllowPaidBroadcast field to given value.

### HasAllowPaidBroadcast

`func (o *SendInvoicePostRequest) HasAllowPaidBroadcast() bool`

HasAllowPaidBroadcast returns a boolean if a field has been set.

### GetMessageEffectId

`func (o *SendInvoicePostRequest) GetMessageEffectId() string`

GetMessageEffectId returns the MessageEffectId field if non-nil, zero value otherwise.

### GetMessageEffectIdOk

`func (o *SendInvoicePostRequest) GetMessageEffectIdOk() (*string, bool)`

GetMessageEffectIdOk returns a tuple with the MessageEffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEffectId

`func (o *SendInvoicePostRequest) SetMessageEffectId(v string)`

SetMessageEffectId sets MessageEffectId field to given value.

### HasMessageEffectId

`func (o *SendInvoicePostRequest) HasMessageEffectId() bool`

HasMessageEffectId returns a boolean if a field has been set.

### GetReplyParameters

`func (o *SendInvoicePostRequest) GetReplyParameters() ReplyParameters`

GetReplyParameters returns the ReplyParameters field if non-nil, zero value otherwise.

### GetReplyParametersOk

`func (o *SendInvoicePostRequest) GetReplyParametersOk() (*ReplyParameters, bool)`

GetReplyParametersOk returns a tuple with the ReplyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyParameters

`func (o *SendInvoicePostRequest) SetReplyParameters(v ReplyParameters)`

SetReplyParameters sets ReplyParameters field to given value.

### HasReplyParameters

`func (o *SendInvoicePostRequest) HasReplyParameters() bool`

HasReplyParameters returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendInvoicePostRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendInvoicePostRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendInvoicePostRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendInvoicePostRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


