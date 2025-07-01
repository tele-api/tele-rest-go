/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Telegram Bot API service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package tele_rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SendInvoicePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendInvoicePostRequest{}

// SendInvoicePostRequest struct for SendInvoicePostRequest
type SendInvoicePostRequest struct {
	ChatId SendMessagePostRequestChatId `json:"chat_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int32 `json:"message_thread_id,omitempty"`
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`
	// Payment provider token, obtained via [@BotFather](https://t.me/botfather). Pass an empty string for payments in [Telegram Stars](https://t.me/BotNews/90).
	ProviderToken *string `json:"provider_token,omitempty"`
	// Three-letter ISO 4217 currency code, see [more on currencies](https://core.telegram.org/bots/payments#supported-currencies). Pass “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90).
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in [Telegram Stars](https://t.me/BotNews/90).
	Prices []LabeledPrice `json:"prices"`
	// The maximum accepted amount for tips in the *smallest units* of the currency (integer, **not** float/double). For example, for a maximum tip of `US$ 1.45` pass `max_tip_amount = 145`. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in [Telegram Stars](https://t.me/BotNews/90).
	MaxTipAmount *int32 `json:"max_tip_amount,omitempty"`
	// A JSON-serialized array of suggested amounts of tips in the *smallest units* of the currency (integer, **not** float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed *max\\_tip\\_amount*.
	SuggestedTipAmounts []int32 `json:"suggested_tip_amounts,omitempty"`
	// Unique deep-linking parameter. If left empty, **forwarded copies** of the sent message will have a *Pay* button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a *URL* button with a deep link to the bot (instead of a *Pay* button), with the value used as the start parameter
	StartParameter *string `json:"start_parameter,omitempty"`
	// JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoUrl *string `json:"photo_url,omitempty"`
	// Photo size in bytes
	PhotoSize *int32 `json:"photo_size,omitempty"`
	// Photo width
	PhotoWidth *int32 `json:"photo_width,omitempty"`
	// Photo height
	PhotoHeight *int32 `json:"photo_height,omitempty"`
	// Pass *True* if you require the user's full name to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedName *bool `json:"need_name,omitempty"`
	// Pass *True* if you require the user's phone number to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty"`
	// Pass *True* if you require the user's email address to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedEmail *bool `json:"need_email,omitempty"`
	// Pass *True* if you require the user's shipping address to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty"`
	// Pass *True* if the user's phone number should be sent to the provider. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty"`
	// Pass *True* if the user's email address should be sent to the provider. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty"`
	// Pass *True* if the final price depends on the shipping method. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	IsFlexible *bool `json:"is_flexible,omitempty"`
	// Sends the message [silently](https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass *True* to allow up to 1000 messages per second, ignoring [broadcasting limits](https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once) for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type _SendInvoicePostRequest SendInvoicePostRequest

// NewSendInvoicePostRequest instantiates a new SendInvoicePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendInvoicePostRequest(chatId SendMessagePostRequestChatId, title string, description string, payload string, currency string, prices []LabeledPrice) *SendInvoicePostRequest {
	this := SendInvoicePostRequest{}
	this.ChatId = chatId
	this.Title = title
	this.Description = description
	this.Payload = payload
	this.Currency = currency
	this.Prices = prices
	var maxTipAmount int32 = 0
	this.MaxTipAmount = &maxTipAmount
	return &this
}

// NewSendInvoicePostRequestWithDefaults instantiates a new SendInvoicePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendInvoicePostRequestWithDefaults() *SendInvoicePostRequest {
	this := SendInvoicePostRequest{}
	var maxTipAmount int32 = 0
	this.MaxTipAmount = &maxTipAmount
	return &this
}

// GetChatId returns the ChatId field value
func (o *SendInvoicePostRequest) GetChatId() SendMessagePostRequestChatId {
	if o == nil {
		var ret SendMessagePostRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SendInvoicePostRequest) SetChatId(v SendMessagePostRequestChatId) {
	o.ChatId = v
}

// GetMessageThreadId returns the MessageThreadId field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetMessageThreadId() int32 {
	if o == nil || IsNil(o.MessageThreadId) {
		var ret int32
		return ret
	}
	return *o.MessageThreadId
}

// GetMessageThreadIdOk returns a tuple with the MessageThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetMessageThreadIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageThreadId) {
		return nil, false
	}
	return o.MessageThreadId, true
}

// HasMessageThreadId returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasMessageThreadId() bool {
	if o != nil && !IsNil(o.MessageThreadId) {
		return true
	}

	return false
}

// SetMessageThreadId gets a reference to the given int32 and assigns it to the MessageThreadId field.
func (o *SendInvoicePostRequest) SetMessageThreadId(v int32) {
	o.MessageThreadId = &v
}


// GetTitle returns the Title field value
func (o *SendInvoicePostRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SendInvoicePostRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *SendInvoicePostRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SendInvoicePostRequest) SetDescription(v string) {
	o.Description = v
}

// GetPayload returns the Payload field value
func (o *SendInvoicePostRequest) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *SendInvoicePostRequest) SetPayload(v string) {
	o.Payload = v
}

// GetProviderToken returns the ProviderToken field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetProviderToken() string {
	if o == nil || IsNil(o.ProviderToken) {
		var ret string
		return ret
	}
	return *o.ProviderToken
}

// GetProviderTokenOk returns a tuple with the ProviderToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetProviderTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderToken) {
		return nil, false
	}
	return o.ProviderToken, true
}

// HasProviderToken returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasProviderToken() bool {
	if o != nil && !IsNil(o.ProviderToken) {
		return true
	}

	return false
}

// SetProviderToken gets a reference to the given string and assigns it to the ProviderToken field.
func (o *SendInvoicePostRequest) SetProviderToken(v string) {
	o.ProviderToken = &v
}


// GetCurrency returns the Currency field value
func (o *SendInvoicePostRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SendInvoicePostRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetPrices returns the Prices field value
func (o *SendInvoicePostRequest) GetPrices() []LabeledPrice {
	if o == nil {
		var ret []LabeledPrice
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetPricesOk() ([]LabeledPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *SendInvoicePostRequest) SetPrices(v []LabeledPrice) {
	o.Prices = v
}

// GetMaxTipAmount returns the MaxTipAmount field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetMaxTipAmount() int32 {
	if o == nil || IsNil(o.MaxTipAmount) {
		var ret int32
		return ret
	}
	return *o.MaxTipAmount
}

// GetMaxTipAmountOk returns a tuple with the MaxTipAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetMaxTipAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTipAmount) {
		return nil, false
	}
	return o.MaxTipAmount, true
}

// HasMaxTipAmount returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasMaxTipAmount() bool {
	if o != nil && !IsNil(o.MaxTipAmount) {
		return true
	}

	return false
}

// SetMaxTipAmount gets a reference to the given int32 and assigns it to the MaxTipAmount field.
func (o *SendInvoicePostRequest) SetMaxTipAmount(v int32) {
	o.MaxTipAmount = &v
}


// GetSuggestedTipAmounts returns the SuggestedTipAmounts field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetSuggestedTipAmounts() []int32 {
	if o == nil || IsNil(o.SuggestedTipAmounts) {
		var ret []int32
		return ret
	}
	return o.SuggestedTipAmounts
}

// GetSuggestedTipAmountsOk returns a tuple with the SuggestedTipAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetSuggestedTipAmountsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SuggestedTipAmounts) {
		return nil, false
	}
	return o.SuggestedTipAmounts, true
}

// HasSuggestedTipAmounts returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasSuggestedTipAmounts() bool {
	if o != nil && !IsNil(o.SuggestedTipAmounts) {
		return true
	}

	return false
}

// SetSuggestedTipAmounts gets a reference to the given []int32 and assigns it to the SuggestedTipAmounts field.
func (o *SendInvoicePostRequest) SetSuggestedTipAmounts(v []int32) {
	o.SuggestedTipAmounts = v
}


// GetStartParameter returns the StartParameter field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetStartParameter() string {
	if o == nil || IsNil(o.StartParameter) {
		var ret string
		return ret
	}
	return *o.StartParameter
}

// GetStartParameterOk returns a tuple with the StartParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetStartParameterOk() (*string, bool) {
	if o == nil || IsNil(o.StartParameter) {
		return nil, false
	}
	return o.StartParameter, true
}

// HasStartParameter returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasStartParameter() bool {
	if o != nil && !IsNil(o.StartParameter) {
		return true
	}

	return false
}

// SetStartParameter gets a reference to the given string and assigns it to the StartParameter field.
func (o *SendInvoicePostRequest) SetStartParameter(v string) {
	o.StartParameter = &v
}


// GetProviderData returns the ProviderData field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetProviderData() string {
	if o == nil || IsNil(o.ProviderData) {
		var ret string
		return ret
	}
	return *o.ProviderData
}

// GetProviderDataOk returns a tuple with the ProviderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetProviderDataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderData) {
		return nil, false
	}
	return o.ProviderData, true
}

// HasProviderData returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasProviderData() bool {
	if o != nil && !IsNil(o.ProviderData) {
		return true
	}

	return false
}

// SetProviderData gets a reference to the given string and assigns it to the ProviderData field.
func (o *SendInvoicePostRequest) SetProviderData(v string) {
	o.ProviderData = &v
}


// GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetPhotoUrl() string {
	if o == nil || IsNil(o.PhotoUrl) {
		var ret string
		return ret
	}
	return *o.PhotoUrl
}

// GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetPhotoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PhotoUrl) {
		return nil, false
	}
	return o.PhotoUrl, true
}

// HasPhotoUrl returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasPhotoUrl() bool {
	if o != nil && !IsNil(o.PhotoUrl) {
		return true
	}

	return false
}

// SetPhotoUrl gets a reference to the given string and assigns it to the PhotoUrl field.
func (o *SendInvoicePostRequest) SetPhotoUrl(v string) {
	o.PhotoUrl = &v
}


// GetPhotoSize returns the PhotoSize field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetPhotoSize() int32 {
	if o == nil || IsNil(o.PhotoSize) {
		var ret int32
		return ret
	}
	return *o.PhotoSize
}

// GetPhotoSizeOk returns a tuple with the PhotoSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetPhotoSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoSize) {
		return nil, false
	}
	return o.PhotoSize, true
}

// HasPhotoSize returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasPhotoSize() bool {
	if o != nil && !IsNil(o.PhotoSize) {
		return true
	}

	return false
}

// SetPhotoSize gets a reference to the given int32 and assigns it to the PhotoSize field.
func (o *SendInvoicePostRequest) SetPhotoSize(v int32) {
	o.PhotoSize = &v
}


// GetPhotoWidth returns the PhotoWidth field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetPhotoWidth() int32 {
	if o == nil || IsNil(o.PhotoWidth) {
		var ret int32
		return ret
	}
	return *o.PhotoWidth
}

// GetPhotoWidthOk returns a tuple with the PhotoWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetPhotoWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoWidth) {
		return nil, false
	}
	return o.PhotoWidth, true
}

// HasPhotoWidth returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasPhotoWidth() bool {
	if o != nil && !IsNil(o.PhotoWidth) {
		return true
	}

	return false
}

// SetPhotoWidth gets a reference to the given int32 and assigns it to the PhotoWidth field.
func (o *SendInvoicePostRequest) SetPhotoWidth(v int32) {
	o.PhotoWidth = &v
}


// GetPhotoHeight returns the PhotoHeight field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetPhotoHeight() int32 {
	if o == nil || IsNil(o.PhotoHeight) {
		var ret int32
		return ret
	}
	return *o.PhotoHeight
}

// GetPhotoHeightOk returns a tuple with the PhotoHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetPhotoHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoHeight) {
		return nil, false
	}
	return o.PhotoHeight, true
}

// HasPhotoHeight returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasPhotoHeight() bool {
	if o != nil && !IsNil(o.PhotoHeight) {
		return true
	}

	return false
}

// SetPhotoHeight gets a reference to the given int32 and assigns it to the PhotoHeight field.
func (o *SendInvoicePostRequest) SetPhotoHeight(v int32) {
	o.PhotoHeight = &v
}


// GetNeedName returns the NeedName field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetNeedName() bool {
	if o == nil || IsNil(o.NeedName) {
		var ret bool
		return ret
	}
	return *o.NeedName
}

// GetNeedNameOk returns a tuple with the NeedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetNeedNameOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedName) {
		return nil, false
	}
	return o.NeedName, true
}

// HasNeedName returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasNeedName() bool {
	if o != nil && !IsNil(o.NeedName) {
		return true
	}

	return false
}

// SetNeedName gets a reference to the given bool and assigns it to the NeedName field.
func (o *SendInvoicePostRequest) SetNeedName(v bool) {
	o.NeedName = &v
}


// GetNeedPhoneNumber returns the NeedPhoneNumber field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetNeedPhoneNumber() bool {
	if o == nil || IsNil(o.NeedPhoneNumber) {
		var ret bool
		return ret
	}
	return *o.NeedPhoneNumber
}

// GetNeedPhoneNumberOk returns a tuple with the NeedPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetNeedPhoneNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedPhoneNumber) {
		return nil, false
	}
	return o.NeedPhoneNumber, true
}

// HasNeedPhoneNumber returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasNeedPhoneNumber() bool {
	if o != nil && !IsNil(o.NeedPhoneNumber) {
		return true
	}

	return false
}

// SetNeedPhoneNumber gets a reference to the given bool and assigns it to the NeedPhoneNumber field.
func (o *SendInvoicePostRequest) SetNeedPhoneNumber(v bool) {
	o.NeedPhoneNumber = &v
}


// GetNeedEmail returns the NeedEmail field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetNeedEmail() bool {
	if o == nil || IsNil(o.NeedEmail) {
		var ret bool
		return ret
	}
	return *o.NeedEmail
}

// GetNeedEmailOk returns a tuple with the NeedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetNeedEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedEmail) {
		return nil, false
	}
	return o.NeedEmail, true
}

// HasNeedEmail returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasNeedEmail() bool {
	if o != nil && !IsNil(o.NeedEmail) {
		return true
	}

	return false
}

// SetNeedEmail gets a reference to the given bool and assigns it to the NeedEmail field.
func (o *SendInvoicePostRequest) SetNeedEmail(v bool) {
	o.NeedEmail = &v
}


// GetNeedShippingAddress returns the NeedShippingAddress field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetNeedShippingAddress() bool {
	if o == nil || IsNil(o.NeedShippingAddress) {
		var ret bool
		return ret
	}
	return *o.NeedShippingAddress
}

// GetNeedShippingAddressOk returns a tuple with the NeedShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetNeedShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedShippingAddress) {
		return nil, false
	}
	return o.NeedShippingAddress, true
}

// HasNeedShippingAddress returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasNeedShippingAddress() bool {
	if o != nil && !IsNil(o.NeedShippingAddress) {
		return true
	}

	return false
}

// SetNeedShippingAddress gets a reference to the given bool and assigns it to the NeedShippingAddress field.
func (o *SendInvoicePostRequest) SetNeedShippingAddress(v bool) {
	o.NeedShippingAddress = &v
}


// GetSendPhoneNumberToProvider returns the SendPhoneNumberToProvider field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetSendPhoneNumberToProvider() bool {
	if o == nil || IsNil(o.SendPhoneNumberToProvider) {
		var ret bool
		return ret
	}
	return *o.SendPhoneNumberToProvider
}

// GetSendPhoneNumberToProviderOk returns a tuple with the SendPhoneNumberToProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetSendPhoneNumberToProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.SendPhoneNumberToProvider) {
		return nil, false
	}
	return o.SendPhoneNumberToProvider, true
}

// HasSendPhoneNumberToProvider returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasSendPhoneNumberToProvider() bool {
	if o != nil && !IsNil(o.SendPhoneNumberToProvider) {
		return true
	}

	return false
}

// SetSendPhoneNumberToProvider gets a reference to the given bool and assigns it to the SendPhoneNumberToProvider field.
func (o *SendInvoicePostRequest) SetSendPhoneNumberToProvider(v bool) {
	o.SendPhoneNumberToProvider = &v
}


// GetSendEmailToProvider returns the SendEmailToProvider field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetSendEmailToProvider() bool {
	if o == nil || IsNil(o.SendEmailToProvider) {
		var ret bool
		return ret
	}
	return *o.SendEmailToProvider
}

// GetSendEmailToProviderOk returns a tuple with the SendEmailToProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetSendEmailToProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmailToProvider) {
		return nil, false
	}
	return o.SendEmailToProvider, true
}

// HasSendEmailToProvider returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasSendEmailToProvider() bool {
	if o != nil && !IsNil(o.SendEmailToProvider) {
		return true
	}

	return false
}

// SetSendEmailToProvider gets a reference to the given bool and assigns it to the SendEmailToProvider field.
func (o *SendInvoicePostRequest) SetSendEmailToProvider(v bool) {
	o.SendEmailToProvider = &v
}


// GetIsFlexible returns the IsFlexible field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetIsFlexible() bool {
	if o == nil || IsNil(o.IsFlexible) {
		var ret bool
		return ret
	}
	return *o.IsFlexible
}

// GetIsFlexibleOk returns a tuple with the IsFlexible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetIsFlexibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFlexible) {
		return nil, false
	}
	return o.IsFlexible, true
}

// HasIsFlexible returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasIsFlexible() bool {
	if o != nil && !IsNil(o.IsFlexible) {
		return true
	}

	return false
}

// SetIsFlexible gets a reference to the given bool and assigns it to the IsFlexible field.
func (o *SendInvoicePostRequest) SetIsFlexible(v bool) {
	o.IsFlexible = &v
}


// GetDisableNotification returns the DisableNotification field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetDisableNotification() bool {
	if o == nil || IsNil(o.DisableNotification) {
		var ret bool
		return ret
	}
	return *o.DisableNotification
}

// GetDisableNotificationOk returns a tuple with the DisableNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetDisableNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableNotification) {
		return nil, false
	}
	return o.DisableNotification, true
}

// HasDisableNotification returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasDisableNotification() bool {
	if o != nil && !IsNil(o.DisableNotification) {
		return true
	}

	return false
}

// SetDisableNotification gets a reference to the given bool and assigns it to the DisableNotification field.
func (o *SendInvoicePostRequest) SetDisableNotification(v bool) {
	o.DisableNotification = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *SendInvoicePostRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


// GetAllowPaidBroadcast returns the AllowPaidBroadcast field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetAllowPaidBroadcast() bool {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		var ret bool
		return ret
	}
	return *o.AllowPaidBroadcast
}

// GetAllowPaidBroadcastOk returns a tuple with the AllowPaidBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetAllowPaidBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPaidBroadcast) {
		return nil, false
	}
	return o.AllowPaidBroadcast, true
}

// HasAllowPaidBroadcast returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasAllowPaidBroadcast() bool {
	if o != nil && !IsNil(o.AllowPaidBroadcast) {
		return true
	}

	return false
}

// SetAllowPaidBroadcast gets a reference to the given bool and assigns it to the AllowPaidBroadcast field.
func (o *SendInvoicePostRequest) SetAllowPaidBroadcast(v bool) {
	o.AllowPaidBroadcast = &v
}


// GetMessageEffectId returns the MessageEffectId field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetMessageEffectId() string {
	if o == nil || IsNil(o.MessageEffectId) {
		var ret string
		return ret
	}
	return *o.MessageEffectId
}

// GetMessageEffectIdOk returns a tuple with the MessageEffectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetMessageEffectIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageEffectId) {
		return nil, false
	}
	return o.MessageEffectId, true
}

// HasMessageEffectId returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasMessageEffectId() bool {
	if o != nil && !IsNil(o.MessageEffectId) {
		return true
	}

	return false
}

// SetMessageEffectId gets a reference to the given string and assigns it to the MessageEffectId field.
func (o *SendInvoicePostRequest) SetMessageEffectId(v string) {
	o.MessageEffectId = &v
}


// GetReplyParameters returns the ReplyParameters field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetReplyParameters() ReplyParameters {
	if o == nil || IsNil(o.ReplyParameters) {
		var ret ReplyParameters
		return ret
	}
	return *o.ReplyParameters
}

// GetReplyParametersOk returns a tuple with the ReplyParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetReplyParametersOk() (*ReplyParameters, bool) {
	if o == nil || IsNil(o.ReplyParameters) {
		return nil, false
	}
	return o.ReplyParameters, true
}

// HasReplyParameters returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasReplyParameters() bool {
	if o != nil && !IsNil(o.ReplyParameters) {
		return true
	}

	return false
}

// SetReplyParameters gets a reference to the given ReplyParameters and assigns it to the ReplyParameters field.
func (o *SendInvoicePostRequest) SetReplyParameters(v ReplyParameters) {
	o.ReplyParameters = &v
}


// GetReplyMarkup returns the ReplyMarkup field value if set, zero value otherwise.
func (o *SendInvoicePostRequest) GetReplyMarkup() InlineKeyboardMarkup {
	if o == nil || IsNil(o.ReplyMarkup) {
		var ret InlineKeyboardMarkup
		return ret
	}
	return *o.ReplyMarkup
}

// GetReplyMarkupOk returns a tuple with the ReplyMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendInvoicePostRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool) {
	if o == nil || IsNil(o.ReplyMarkup) {
		return nil, false
	}
	return o.ReplyMarkup, true
}

// HasReplyMarkup returns a boolean if a field has been set.
func (o *SendInvoicePostRequest) HasReplyMarkup() bool {
	if o != nil && !IsNil(o.ReplyMarkup) {
		return true
	}

	return false
}

// SetReplyMarkup gets a reference to the given InlineKeyboardMarkup and assigns it to the ReplyMarkup field.
func (o *SendInvoicePostRequest) SetReplyMarkup(v InlineKeyboardMarkup) {
	o.ReplyMarkup = &v
}


func (o SendInvoicePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendInvoicePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.MessageThreadId) {
		toSerialize["message_thread_id"] = o.MessageThreadId
	}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["payload"] = o.Payload
	if !IsNil(o.ProviderToken) {
		toSerialize["provider_token"] = o.ProviderToken
	}
	toSerialize["currency"] = o.Currency
	toSerialize["prices"] = o.Prices
	if !IsNil(o.MaxTipAmount) {
		toSerialize["max_tip_amount"] = o.MaxTipAmount
	}
	if !IsNil(o.SuggestedTipAmounts) {
		toSerialize["suggested_tip_amounts"] = o.SuggestedTipAmounts
	}
	if !IsNil(o.StartParameter) {
		toSerialize["start_parameter"] = o.StartParameter
	}
	if !IsNil(o.ProviderData) {
		toSerialize["provider_data"] = o.ProviderData
	}
	if !IsNil(o.PhotoUrl) {
		toSerialize["photo_url"] = o.PhotoUrl
	}
	if !IsNil(o.PhotoSize) {
		toSerialize["photo_size"] = o.PhotoSize
	}
	if !IsNil(o.PhotoWidth) {
		toSerialize["photo_width"] = o.PhotoWidth
	}
	if !IsNil(o.PhotoHeight) {
		toSerialize["photo_height"] = o.PhotoHeight
	}
	if !IsNil(o.NeedName) {
		toSerialize["need_name"] = o.NeedName
	}
	if !IsNil(o.NeedPhoneNumber) {
		toSerialize["need_phone_number"] = o.NeedPhoneNumber
	}
	if !IsNil(o.NeedEmail) {
		toSerialize["need_email"] = o.NeedEmail
	}
	if !IsNil(o.NeedShippingAddress) {
		toSerialize["need_shipping_address"] = o.NeedShippingAddress
	}
	if !IsNil(o.SendPhoneNumberToProvider) {
		toSerialize["send_phone_number_to_provider"] = o.SendPhoneNumberToProvider
	}
	if !IsNil(o.SendEmailToProvider) {
		toSerialize["send_email_to_provider"] = o.SendEmailToProvider
	}
	if !IsNil(o.IsFlexible) {
		toSerialize["is_flexible"] = o.IsFlexible
	}
	if !IsNil(o.DisableNotification) {
		toSerialize["disable_notification"] = o.DisableNotification
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	if !IsNil(o.AllowPaidBroadcast) {
		toSerialize["allow_paid_broadcast"] = o.AllowPaidBroadcast
	}
	if !IsNil(o.MessageEffectId) {
		toSerialize["message_effect_id"] = o.MessageEffectId
	}
	if !IsNil(o.ReplyParameters) {
		toSerialize["reply_parameters"] = o.ReplyParameters
	}
	if !IsNil(o.ReplyMarkup) {
		toSerialize["reply_markup"] = o.ReplyMarkup
	}
	return toSerialize, nil
}

func (o *SendInvoicePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"title",
		"description",
		"payload",
		"currency",
		"prices",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSendInvoicePostRequest := _SendInvoicePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendInvoicePostRequest)

	if err != nil {
		return err
	}

	*o = SendInvoicePostRequest(varSendInvoicePostRequest)

	return err
}

type NullableSendInvoicePostRequest struct {
	value *SendInvoicePostRequest
	isSet bool
}

func (v NullableSendInvoicePostRequest) Get() *SendInvoicePostRequest {
	return v.value
}

func (v *NullableSendInvoicePostRequest) Set(val *SendInvoicePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendInvoicePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendInvoicePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendInvoicePostRequest(val *SendInvoicePostRequest) *NullableSendInvoicePostRequest {
	return &NullableSendInvoicePostRequest{value: val, isSet: true}
}

func (v NullableSendInvoicePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendInvoicePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


