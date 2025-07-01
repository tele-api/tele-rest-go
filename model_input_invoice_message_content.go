/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the InputInvoiceMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputInvoiceMessageContent{}

// InputInvoiceMessageContent Represents the [content](https://core.telegram.org/bots/api/#inputmessagecontent) of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`
	// *Optional*. Payment provider token, obtained via [@BotFather](https://t.me/botfather). Pass an empty string for payments in [Telegram Stars](https://t.me/BotNews/90).
	ProviderToken *string `json:"provider_token,omitempty"`
	// Three-letter ISO 4217 currency code, see [more on currencies](https://core.telegram.org/bots/payments#supported-currencies). Pass “XTR” for payments in [Telegram Stars](https://t.me/BotNews/90).
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in [Telegram Stars](https://t.me/BotNews/90).
	Prices []LabeledPrice `json:"prices"`
	// *Optional*. The maximum accepted amount for tips in the *smallest units* of the currency (integer, **not** float/double). For example, for a maximum tip of `US$ 1.45` pass `max_tip_amount = 145`. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in [Telegram Stars](https://t.me/BotNews/90).
	MaxTipAmount *int32 `json:"max_tip_amount,omitempty"`
	// *Optional*. A JSON-serialized array of suggested amounts of tip in the *smallest units* of the currency (integer, **not** float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed *max\\_tip\\_amount*.
	SuggestedTipAmounts []int32 `json:"suggested_tip_amounts,omitempty"`
	// *Optional*. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty"`
	// *Optional*. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoUrl *string `json:"photo_url,omitempty"`
	// *Optional*. Photo size in bytes
	PhotoSize *int32 `json:"photo_size,omitempty"`
	// *Optional*. Photo width
	PhotoWidth *int32 `json:"photo_width,omitempty"`
	// *Optional*. Photo height
	PhotoHeight *int32 `json:"photo_height,omitempty"`
	// *Optional*. Pass *True* if you require the user's full name to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedName *bool `json:"need_name,omitempty"`
	// *Optional*. Pass *True* if you require the user's phone number to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty"`
	// *Optional*. Pass *True* if you require the user's email address to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedEmail *bool `json:"need_email,omitempty"`
	// *Optional*. Pass *True* if you require the user's shipping address to complete the order. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty"`
	// *Optional*. Pass *True* if the user's phone number should be sent to the provider. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty"`
	// *Optional*. Pass *True* if the user's email address should be sent to the provider. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty"`
	// *Optional*. Pass *True* if the final price depends on the shipping method. Ignored for payments in [Telegram Stars](https://t.me/BotNews/90).
	IsFlexible *bool `json:"is_flexible,omitempty"`
}

type _InputInvoiceMessageContent InputInvoiceMessageContent

// NewInputInvoiceMessageContent instantiates a new InputInvoiceMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputInvoiceMessageContent(title string, description string, payload string, currency string, prices []LabeledPrice) *InputInvoiceMessageContent {
	this := InputInvoiceMessageContent{}
	this.Title = title
	this.Description = description
	this.Payload = payload
	this.Currency = currency
	this.Prices = prices
	var maxTipAmount int32 = 0
	this.MaxTipAmount = &maxTipAmount
	return &this
}

// NewInputInvoiceMessageContentWithDefaults instantiates a new InputInvoiceMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputInvoiceMessageContentWithDefaults() *InputInvoiceMessageContent {
	this := InputInvoiceMessageContent{}
	var maxTipAmount int32 = 0
	this.MaxTipAmount = &maxTipAmount
	return &this
}

// GetTitle returns the Title field value
func (o *InputInvoiceMessageContent) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InputInvoiceMessageContent) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *InputInvoiceMessageContent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InputInvoiceMessageContent) SetDescription(v string) {
	o.Description = v
}

// GetPayload returns the Payload field value
func (o *InputInvoiceMessageContent) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *InputInvoiceMessageContent) SetPayload(v string) {
	o.Payload = v
}

// GetProviderToken returns the ProviderToken field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetProviderToken() string {
	if o == nil || IsNil(o.ProviderToken) {
		var ret string
		return ret
	}
	return *o.ProviderToken
}

// GetProviderTokenOk returns a tuple with the ProviderToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetProviderTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderToken) {
		return nil, false
	}
	return o.ProviderToken, true
}

// HasProviderToken returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasProviderToken() bool {
	if o != nil && !IsNil(o.ProviderToken) {
		return true
	}

	return false
}

// SetProviderToken gets a reference to the given string and assigns it to the ProviderToken field.
func (o *InputInvoiceMessageContent) SetProviderToken(v string) {
	o.ProviderToken = &v
}


// GetCurrency returns the Currency field value
func (o *InputInvoiceMessageContent) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InputInvoiceMessageContent) SetCurrency(v string) {
	o.Currency = v
}

// GetPrices returns the Prices field value
func (o *InputInvoiceMessageContent) GetPrices() []LabeledPrice {
	if o == nil {
		var ret []LabeledPrice
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetPricesOk() ([]LabeledPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *InputInvoiceMessageContent) SetPrices(v []LabeledPrice) {
	o.Prices = v
}

// GetMaxTipAmount returns the MaxTipAmount field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetMaxTipAmount() int32 {
	if o == nil || IsNil(o.MaxTipAmount) {
		var ret int32
		return ret
	}
	return *o.MaxTipAmount
}

// GetMaxTipAmountOk returns a tuple with the MaxTipAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetMaxTipAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTipAmount) {
		return nil, false
	}
	return o.MaxTipAmount, true
}

// HasMaxTipAmount returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasMaxTipAmount() bool {
	if o != nil && !IsNil(o.MaxTipAmount) {
		return true
	}

	return false
}

// SetMaxTipAmount gets a reference to the given int32 and assigns it to the MaxTipAmount field.
func (o *InputInvoiceMessageContent) SetMaxTipAmount(v int32) {
	o.MaxTipAmount = &v
}


// GetSuggestedTipAmounts returns the SuggestedTipAmounts field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetSuggestedTipAmounts() []int32 {
	if o == nil || IsNil(o.SuggestedTipAmounts) {
		var ret []int32
		return ret
	}
	return o.SuggestedTipAmounts
}

// GetSuggestedTipAmountsOk returns a tuple with the SuggestedTipAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetSuggestedTipAmountsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SuggestedTipAmounts) {
		return nil, false
	}
	return o.SuggestedTipAmounts, true
}

// HasSuggestedTipAmounts returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasSuggestedTipAmounts() bool {
	if o != nil && !IsNil(o.SuggestedTipAmounts) {
		return true
	}

	return false
}

// SetSuggestedTipAmounts gets a reference to the given []int32 and assigns it to the SuggestedTipAmounts field.
func (o *InputInvoiceMessageContent) SetSuggestedTipAmounts(v []int32) {
	o.SuggestedTipAmounts = v
}


// GetProviderData returns the ProviderData field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetProviderData() string {
	if o == nil || IsNil(o.ProviderData) {
		var ret string
		return ret
	}
	return *o.ProviderData
}

// GetProviderDataOk returns a tuple with the ProviderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetProviderDataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderData) {
		return nil, false
	}
	return o.ProviderData, true
}

// HasProviderData returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasProviderData() bool {
	if o != nil && !IsNil(o.ProviderData) {
		return true
	}

	return false
}

// SetProviderData gets a reference to the given string and assigns it to the ProviderData field.
func (o *InputInvoiceMessageContent) SetProviderData(v string) {
	o.ProviderData = &v
}


// GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetPhotoUrl() string {
	if o == nil || IsNil(o.PhotoUrl) {
		var ret string
		return ret
	}
	return *o.PhotoUrl
}

// GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetPhotoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PhotoUrl) {
		return nil, false
	}
	return o.PhotoUrl, true
}

// HasPhotoUrl returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasPhotoUrl() bool {
	if o != nil && !IsNil(o.PhotoUrl) {
		return true
	}

	return false
}

// SetPhotoUrl gets a reference to the given string and assigns it to the PhotoUrl field.
func (o *InputInvoiceMessageContent) SetPhotoUrl(v string) {
	o.PhotoUrl = &v
}


// GetPhotoSize returns the PhotoSize field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetPhotoSize() int32 {
	if o == nil || IsNil(o.PhotoSize) {
		var ret int32
		return ret
	}
	return *o.PhotoSize
}

// GetPhotoSizeOk returns a tuple with the PhotoSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetPhotoSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoSize) {
		return nil, false
	}
	return o.PhotoSize, true
}

// HasPhotoSize returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasPhotoSize() bool {
	if o != nil && !IsNil(o.PhotoSize) {
		return true
	}

	return false
}

// SetPhotoSize gets a reference to the given int32 and assigns it to the PhotoSize field.
func (o *InputInvoiceMessageContent) SetPhotoSize(v int32) {
	o.PhotoSize = &v
}


// GetPhotoWidth returns the PhotoWidth field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetPhotoWidth() int32 {
	if o == nil || IsNil(o.PhotoWidth) {
		var ret int32
		return ret
	}
	return *o.PhotoWidth
}

// GetPhotoWidthOk returns a tuple with the PhotoWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetPhotoWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoWidth) {
		return nil, false
	}
	return o.PhotoWidth, true
}

// HasPhotoWidth returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasPhotoWidth() bool {
	if o != nil && !IsNil(o.PhotoWidth) {
		return true
	}

	return false
}

// SetPhotoWidth gets a reference to the given int32 and assigns it to the PhotoWidth field.
func (o *InputInvoiceMessageContent) SetPhotoWidth(v int32) {
	o.PhotoWidth = &v
}


// GetPhotoHeight returns the PhotoHeight field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetPhotoHeight() int32 {
	if o == nil || IsNil(o.PhotoHeight) {
		var ret int32
		return ret
	}
	return *o.PhotoHeight
}

// GetPhotoHeightOk returns a tuple with the PhotoHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetPhotoHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoHeight) {
		return nil, false
	}
	return o.PhotoHeight, true
}

// HasPhotoHeight returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasPhotoHeight() bool {
	if o != nil && !IsNil(o.PhotoHeight) {
		return true
	}

	return false
}

// SetPhotoHeight gets a reference to the given int32 and assigns it to the PhotoHeight field.
func (o *InputInvoiceMessageContent) SetPhotoHeight(v int32) {
	o.PhotoHeight = &v
}


// GetNeedName returns the NeedName field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetNeedName() bool {
	if o == nil || IsNil(o.NeedName) {
		var ret bool
		return ret
	}
	return *o.NeedName
}

// GetNeedNameOk returns a tuple with the NeedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetNeedNameOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedName) {
		return nil, false
	}
	return o.NeedName, true
}

// HasNeedName returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasNeedName() bool {
	if o != nil && !IsNil(o.NeedName) {
		return true
	}

	return false
}

// SetNeedName gets a reference to the given bool and assigns it to the NeedName field.
func (o *InputInvoiceMessageContent) SetNeedName(v bool) {
	o.NeedName = &v
}


// GetNeedPhoneNumber returns the NeedPhoneNumber field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetNeedPhoneNumber() bool {
	if o == nil || IsNil(o.NeedPhoneNumber) {
		var ret bool
		return ret
	}
	return *o.NeedPhoneNumber
}

// GetNeedPhoneNumberOk returns a tuple with the NeedPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetNeedPhoneNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedPhoneNumber) {
		return nil, false
	}
	return o.NeedPhoneNumber, true
}

// HasNeedPhoneNumber returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasNeedPhoneNumber() bool {
	if o != nil && !IsNil(o.NeedPhoneNumber) {
		return true
	}

	return false
}

// SetNeedPhoneNumber gets a reference to the given bool and assigns it to the NeedPhoneNumber field.
func (o *InputInvoiceMessageContent) SetNeedPhoneNumber(v bool) {
	o.NeedPhoneNumber = &v
}


// GetNeedEmail returns the NeedEmail field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetNeedEmail() bool {
	if o == nil || IsNil(o.NeedEmail) {
		var ret bool
		return ret
	}
	return *o.NeedEmail
}

// GetNeedEmailOk returns a tuple with the NeedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetNeedEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedEmail) {
		return nil, false
	}
	return o.NeedEmail, true
}

// HasNeedEmail returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasNeedEmail() bool {
	if o != nil && !IsNil(o.NeedEmail) {
		return true
	}

	return false
}

// SetNeedEmail gets a reference to the given bool and assigns it to the NeedEmail field.
func (o *InputInvoiceMessageContent) SetNeedEmail(v bool) {
	o.NeedEmail = &v
}


// GetNeedShippingAddress returns the NeedShippingAddress field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetNeedShippingAddress() bool {
	if o == nil || IsNil(o.NeedShippingAddress) {
		var ret bool
		return ret
	}
	return *o.NeedShippingAddress
}

// GetNeedShippingAddressOk returns a tuple with the NeedShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetNeedShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedShippingAddress) {
		return nil, false
	}
	return o.NeedShippingAddress, true
}

// HasNeedShippingAddress returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasNeedShippingAddress() bool {
	if o != nil && !IsNil(o.NeedShippingAddress) {
		return true
	}

	return false
}

// SetNeedShippingAddress gets a reference to the given bool and assigns it to the NeedShippingAddress field.
func (o *InputInvoiceMessageContent) SetNeedShippingAddress(v bool) {
	o.NeedShippingAddress = &v
}


// GetSendPhoneNumberToProvider returns the SendPhoneNumberToProvider field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetSendPhoneNumberToProvider() bool {
	if o == nil || IsNil(o.SendPhoneNumberToProvider) {
		var ret bool
		return ret
	}
	return *o.SendPhoneNumberToProvider
}

// GetSendPhoneNumberToProviderOk returns a tuple with the SendPhoneNumberToProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetSendPhoneNumberToProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.SendPhoneNumberToProvider) {
		return nil, false
	}
	return o.SendPhoneNumberToProvider, true
}

// HasSendPhoneNumberToProvider returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasSendPhoneNumberToProvider() bool {
	if o != nil && !IsNil(o.SendPhoneNumberToProvider) {
		return true
	}

	return false
}

// SetSendPhoneNumberToProvider gets a reference to the given bool and assigns it to the SendPhoneNumberToProvider field.
func (o *InputInvoiceMessageContent) SetSendPhoneNumberToProvider(v bool) {
	o.SendPhoneNumberToProvider = &v
}


// GetSendEmailToProvider returns the SendEmailToProvider field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetSendEmailToProvider() bool {
	if o == nil || IsNil(o.SendEmailToProvider) {
		var ret bool
		return ret
	}
	return *o.SendEmailToProvider
}

// GetSendEmailToProviderOk returns a tuple with the SendEmailToProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetSendEmailToProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmailToProvider) {
		return nil, false
	}
	return o.SendEmailToProvider, true
}

// HasSendEmailToProvider returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasSendEmailToProvider() bool {
	if o != nil && !IsNil(o.SendEmailToProvider) {
		return true
	}

	return false
}

// SetSendEmailToProvider gets a reference to the given bool and assigns it to the SendEmailToProvider field.
func (o *InputInvoiceMessageContent) SetSendEmailToProvider(v bool) {
	o.SendEmailToProvider = &v
}


// GetIsFlexible returns the IsFlexible field value if set, zero value otherwise.
func (o *InputInvoiceMessageContent) GetIsFlexible() bool {
	if o == nil || IsNil(o.IsFlexible) {
		var ret bool
		return ret
	}
	return *o.IsFlexible
}

// GetIsFlexibleOk returns a tuple with the IsFlexible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputInvoiceMessageContent) GetIsFlexibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFlexible) {
		return nil, false
	}
	return o.IsFlexible, true
}

// HasIsFlexible returns a boolean if a field has been set.
func (o *InputInvoiceMessageContent) HasIsFlexible() bool {
	if o != nil && !IsNil(o.IsFlexible) {
		return true
	}

	return false
}

// SetIsFlexible gets a reference to the given bool and assigns it to the IsFlexible field.
func (o *InputInvoiceMessageContent) SetIsFlexible(v bool) {
	o.IsFlexible = &v
}


func (o InputInvoiceMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputInvoiceMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

func (o *InputInvoiceMessageContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varInputInvoiceMessageContent := _InputInvoiceMessageContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputInvoiceMessageContent)

	if err != nil {
		return err
	}

	*o = InputInvoiceMessageContent(varInputInvoiceMessageContent)

	return err
}

type NullableInputInvoiceMessageContent struct {
	value *InputInvoiceMessageContent
	isSet bool
}

func (v NullableInputInvoiceMessageContent) Get() *InputInvoiceMessageContent {
	return v.value
}

func (v *NullableInputInvoiceMessageContent) Set(val *InputInvoiceMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputInvoiceMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputInvoiceMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputInvoiceMessageContent(val *InputInvoiceMessageContent) *NullableInputInvoiceMessageContent {
	return &NullableInputInvoiceMessageContent{value: val, isSet: true}
}

func (v NullableInputInvoiceMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputInvoiceMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


