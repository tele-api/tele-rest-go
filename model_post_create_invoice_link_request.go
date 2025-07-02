/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the PostCreateInvoiceLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateInvoiceLinkRequest{}

// PostCreateInvoiceLinkRequest struct for PostCreateInvoiceLinkRequest
type PostCreateInvoiceLinkRequest struct {
	// Unique identifier of the business connection on behalf of which the link will be created. For payments in [Telegram Stars](https://t.me/BotNews/90) only.
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
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
	// The number of seconds the subscription will be active for before the next payment. The currency must be set to “XTR” (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 10000 Telegram Stars.
	SubscriptionPeriod *int32 `json:"subscription_period,omitempty"`
	// The maximum accepted amount for tips in the *smallest units* of the currency (integer, **not** float/double). For example, for a maximum tip of `US$ 1.45` pass `max_tip_amount = 145`. See the *exp* parameter in [currencies.json](https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in [Telegram Stars](https://t.me/BotNews/90).
	MaxTipAmount *int32 `json:"max_tip_amount,omitempty"`
	// A JSON-serialized array of suggested amounts of tips in the *smallest units* of the currency (integer, **not** float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed *max\\_tip\\_amount*.
	SuggestedTipAmounts []int32 `json:"suggested_tip_amounts,omitempty"`
	// JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
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
}

type _PostCreateInvoiceLinkRequest PostCreateInvoiceLinkRequest

// NewPostCreateInvoiceLinkRequest instantiates a new PostCreateInvoiceLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateInvoiceLinkRequest(title string, description string, payload string, currency string, prices []LabeledPrice) *PostCreateInvoiceLinkRequest {
	this := PostCreateInvoiceLinkRequest{}
	this.Title = title
	this.Description = description
	this.Payload = payload
	this.Currency = currency
	this.Prices = prices
	var maxTipAmount int32 = 0
	this.MaxTipAmount = &maxTipAmount
	return &this
}

// NewPostCreateInvoiceLinkRequestWithDefaults instantiates a new PostCreateInvoiceLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateInvoiceLinkRequestWithDefaults() *PostCreateInvoiceLinkRequest {
	this := PostCreateInvoiceLinkRequest{}
	var maxTipAmount int32 = 0
	this.MaxTipAmount = &maxTipAmount
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetBusinessConnectionId() string {
	if o == nil || IsNil(o.BusinessConnectionId) {
		var ret string
		return ret
	}
	return *o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessConnectionId) {
		return nil, false
	}
	return o.BusinessConnectionId, true
}

// HasBusinessConnectionId returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasBusinessConnectionId() bool {
	if o != nil && !IsNil(o.BusinessConnectionId) {
		return true
	}

	return false
}

// SetBusinessConnectionId gets a reference to the given string and assigns it to the BusinessConnectionId field.
func (o *PostCreateInvoiceLinkRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = &v
}


// GetTitle returns the Title field value
func (o *PostCreateInvoiceLinkRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PostCreateInvoiceLinkRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *PostCreateInvoiceLinkRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostCreateInvoiceLinkRequest) SetDescription(v string) {
	o.Description = v
}

// GetPayload returns the Payload field value
func (o *PostCreateInvoiceLinkRequest) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *PostCreateInvoiceLinkRequest) SetPayload(v string) {
	o.Payload = v
}

// GetProviderToken returns the ProviderToken field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetProviderToken() string {
	if o == nil || IsNil(o.ProviderToken) {
		var ret string
		return ret
	}
	return *o.ProviderToken
}

// GetProviderTokenOk returns a tuple with the ProviderToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetProviderTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderToken) {
		return nil, false
	}
	return o.ProviderToken, true
}

// HasProviderToken returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasProviderToken() bool {
	if o != nil && !IsNil(o.ProviderToken) {
		return true
	}

	return false
}

// SetProviderToken gets a reference to the given string and assigns it to the ProviderToken field.
func (o *PostCreateInvoiceLinkRequest) SetProviderToken(v string) {
	o.ProviderToken = &v
}


// GetCurrency returns the Currency field value
func (o *PostCreateInvoiceLinkRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PostCreateInvoiceLinkRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetPrices returns the Prices field value
func (o *PostCreateInvoiceLinkRequest) GetPrices() []LabeledPrice {
	if o == nil {
		var ret []LabeledPrice
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetPricesOk() ([]LabeledPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *PostCreateInvoiceLinkRequest) SetPrices(v []LabeledPrice) {
	o.Prices = v
}

// GetSubscriptionPeriod returns the SubscriptionPeriod field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetSubscriptionPeriod() int32 {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		var ret int32
		return ret
	}
	return *o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetSubscriptionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		return nil, false
	}
	return o.SubscriptionPeriod, true
}

// HasSubscriptionPeriod returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasSubscriptionPeriod() bool {
	if o != nil && !IsNil(o.SubscriptionPeriod) {
		return true
	}

	return false
}

// SetSubscriptionPeriod gets a reference to the given int32 and assigns it to the SubscriptionPeriod field.
func (o *PostCreateInvoiceLinkRequest) SetSubscriptionPeriod(v int32) {
	o.SubscriptionPeriod = &v
}


// GetMaxTipAmount returns the MaxTipAmount field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetMaxTipAmount() int32 {
	if o == nil || IsNil(o.MaxTipAmount) {
		var ret int32
		return ret
	}
	return *o.MaxTipAmount
}

// GetMaxTipAmountOk returns a tuple with the MaxTipAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetMaxTipAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTipAmount) {
		return nil, false
	}
	return o.MaxTipAmount, true
}

// HasMaxTipAmount returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasMaxTipAmount() bool {
	if o != nil && !IsNil(o.MaxTipAmount) {
		return true
	}

	return false
}

// SetMaxTipAmount gets a reference to the given int32 and assigns it to the MaxTipAmount field.
func (o *PostCreateInvoiceLinkRequest) SetMaxTipAmount(v int32) {
	o.MaxTipAmount = &v
}


// GetSuggestedTipAmounts returns the SuggestedTipAmounts field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetSuggestedTipAmounts() []int32 {
	if o == nil || IsNil(o.SuggestedTipAmounts) {
		var ret []int32
		return ret
	}
	return o.SuggestedTipAmounts
}

// GetSuggestedTipAmountsOk returns a tuple with the SuggestedTipAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetSuggestedTipAmountsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SuggestedTipAmounts) {
		return nil, false
	}
	return o.SuggestedTipAmounts, true
}

// HasSuggestedTipAmounts returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasSuggestedTipAmounts() bool {
	if o != nil && !IsNil(o.SuggestedTipAmounts) {
		return true
	}

	return false
}

// SetSuggestedTipAmounts gets a reference to the given []int32 and assigns it to the SuggestedTipAmounts field.
func (o *PostCreateInvoiceLinkRequest) SetSuggestedTipAmounts(v []int32) {
	o.SuggestedTipAmounts = v
}


// GetProviderData returns the ProviderData field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetProviderData() string {
	if o == nil || IsNil(o.ProviderData) {
		var ret string
		return ret
	}
	return *o.ProviderData
}

// GetProviderDataOk returns a tuple with the ProviderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetProviderDataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderData) {
		return nil, false
	}
	return o.ProviderData, true
}

// HasProviderData returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasProviderData() bool {
	if o != nil && !IsNil(o.ProviderData) {
		return true
	}

	return false
}

// SetProviderData gets a reference to the given string and assigns it to the ProviderData field.
func (o *PostCreateInvoiceLinkRequest) SetProviderData(v string) {
	o.ProviderData = &v
}


// GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetPhotoUrl() string {
	if o == nil || IsNil(o.PhotoUrl) {
		var ret string
		return ret
	}
	return *o.PhotoUrl
}

// GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetPhotoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PhotoUrl) {
		return nil, false
	}
	return o.PhotoUrl, true
}

// HasPhotoUrl returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasPhotoUrl() bool {
	if o != nil && !IsNil(o.PhotoUrl) {
		return true
	}

	return false
}

// SetPhotoUrl gets a reference to the given string and assigns it to the PhotoUrl field.
func (o *PostCreateInvoiceLinkRequest) SetPhotoUrl(v string) {
	o.PhotoUrl = &v
}


// GetPhotoSize returns the PhotoSize field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetPhotoSize() int32 {
	if o == nil || IsNil(o.PhotoSize) {
		var ret int32
		return ret
	}
	return *o.PhotoSize
}

// GetPhotoSizeOk returns a tuple with the PhotoSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetPhotoSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoSize) {
		return nil, false
	}
	return o.PhotoSize, true
}

// HasPhotoSize returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasPhotoSize() bool {
	if o != nil && !IsNil(o.PhotoSize) {
		return true
	}

	return false
}

// SetPhotoSize gets a reference to the given int32 and assigns it to the PhotoSize field.
func (o *PostCreateInvoiceLinkRequest) SetPhotoSize(v int32) {
	o.PhotoSize = &v
}


// GetPhotoWidth returns the PhotoWidth field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetPhotoWidth() int32 {
	if o == nil || IsNil(o.PhotoWidth) {
		var ret int32
		return ret
	}
	return *o.PhotoWidth
}

// GetPhotoWidthOk returns a tuple with the PhotoWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetPhotoWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoWidth) {
		return nil, false
	}
	return o.PhotoWidth, true
}

// HasPhotoWidth returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasPhotoWidth() bool {
	if o != nil && !IsNil(o.PhotoWidth) {
		return true
	}

	return false
}

// SetPhotoWidth gets a reference to the given int32 and assigns it to the PhotoWidth field.
func (o *PostCreateInvoiceLinkRequest) SetPhotoWidth(v int32) {
	o.PhotoWidth = &v
}


// GetPhotoHeight returns the PhotoHeight field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetPhotoHeight() int32 {
	if o == nil || IsNil(o.PhotoHeight) {
		var ret int32
		return ret
	}
	return *o.PhotoHeight
}

// GetPhotoHeightOk returns a tuple with the PhotoHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetPhotoHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.PhotoHeight) {
		return nil, false
	}
	return o.PhotoHeight, true
}

// HasPhotoHeight returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasPhotoHeight() bool {
	if o != nil && !IsNil(o.PhotoHeight) {
		return true
	}

	return false
}

// SetPhotoHeight gets a reference to the given int32 and assigns it to the PhotoHeight field.
func (o *PostCreateInvoiceLinkRequest) SetPhotoHeight(v int32) {
	o.PhotoHeight = &v
}


// GetNeedName returns the NeedName field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetNeedName() bool {
	if o == nil || IsNil(o.NeedName) {
		var ret bool
		return ret
	}
	return *o.NeedName
}

// GetNeedNameOk returns a tuple with the NeedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetNeedNameOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedName) {
		return nil, false
	}
	return o.NeedName, true
}

// HasNeedName returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasNeedName() bool {
	if o != nil && !IsNil(o.NeedName) {
		return true
	}

	return false
}

// SetNeedName gets a reference to the given bool and assigns it to the NeedName field.
func (o *PostCreateInvoiceLinkRequest) SetNeedName(v bool) {
	o.NeedName = &v
}


// GetNeedPhoneNumber returns the NeedPhoneNumber field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetNeedPhoneNumber() bool {
	if o == nil || IsNil(o.NeedPhoneNumber) {
		var ret bool
		return ret
	}
	return *o.NeedPhoneNumber
}

// GetNeedPhoneNumberOk returns a tuple with the NeedPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetNeedPhoneNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedPhoneNumber) {
		return nil, false
	}
	return o.NeedPhoneNumber, true
}

// HasNeedPhoneNumber returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasNeedPhoneNumber() bool {
	if o != nil && !IsNil(o.NeedPhoneNumber) {
		return true
	}

	return false
}

// SetNeedPhoneNumber gets a reference to the given bool and assigns it to the NeedPhoneNumber field.
func (o *PostCreateInvoiceLinkRequest) SetNeedPhoneNumber(v bool) {
	o.NeedPhoneNumber = &v
}


// GetNeedEmail returns the NeedEmail field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetNeedEmail() bool {
	if o == nil || IsNil(o.NeedEmail) {
		var ret bool
		return ret
	}
	return *o.NeedEmail
}

// GetNeedEmailOk returns a tuple with the NeedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetNeedEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedEmail) {
		return nil, false
	}
	return o.NeedEmail, true
}

// HasNeedEmail returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasNeedEmail() bool {
	if o != nil && !IsNil(o.NeedEmail) {
		return true
	}

	return false
}

// SetNeedEmail gets a reference to the given bool and assigns it to the NeedEmail field.
func (o *PostCreateInvoiceLinkRequest) SetNeedEmail(v bool) {
	o.NeedEmail = &v
}


// GetNeedShippingAddress returns the NeedShippingAddress field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetNeedShippingAddress() bool {
	if o == nil || IsNil(o.NeedShippingAddress) {
		var ret bool
		return ret
	}
	return *o.NeedShippingAddress
}

// GetNeedShippingAddressOk returns a tuple with the NeedShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetNeedShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedShippingAddress) {
		return nil, false
	}
	return o.NeedShippingAddress, true
}

// HasNeedShippingAddress returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasNeedShippingAddress() bool {
	if o != nil && !IsNil(o.NeedShippingAddress) {
		return true
	}

	return false
}

// SetNeedShippingAddress gets a reference to the given bool and assigns it to the NeedShippingAddress field.
func (o *PostCreateInvoiceLinkRequest) SetNeedShippingAddress(v bool) {
	o.NeedShippingAddress = &v
}


// GetSendPhoneNumberToProvider returns the SendPhoneNumberToProvider field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetSendPhoneNumberToProvider() bool {
	if o == nil || IsNil(o.SendPhoneNumberToProvider) {
		var ret bool
		return ret
	}
	return *o.SendPhoneNumberToProvider
}

// GetSendPhoneNumberToProviderOk returns a tuple with the SendPhoneNumberToProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetSendPhoneNumberToProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.SendPhoneNumberToProvider) {
		return nil, false
	}
	return o.SendPhoneNumberToProvider, true
}

// HasSendPhoneNumberToProvider returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasSendPhoneNumberToProvider() bool {
	if o != nil && !IsNil(o.SendPhoneNumberToProvider) {
		return true
	}

	return false
}

// SetSendPhoneNumberToProvider gets a reference to the given bool and assigns it to the SendPhoneNumberToProvider field.
func (o *PostCreateInvoiceLinkRequest) SetSendPhoneNumberToProvider(v bool) {
	o.SendPhoneNumberToProvider = &v
}


// GetSendEmailToProvider returns the SendEmailToProvider field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetSendEmailToProvider() bool {
	if o == nil || IsNil(o.SendEmailToProvider) {
		var ret bool
		return ret
	}
	return *o.SendEmailToProvider
}

// GetSendEmailToProviderOk returns a tuple with the SendEmailToProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetSendEmailToProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmailToProvider) {
		return nil, false
	}
	return o.SendEmailToProvider, true
}

// HasSendEmailToProvider returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasSendEmailToProvider() bool {
	if o != nil && !IsNil(o.SendEmailToProvider) {
		return true
	}

	return false
}

// SetSendEmailToProvider gets a reference to the given bool and assigns it to the SendEmailToProvider field.
func (o *PostCreateInvoiceLinkRequest) SetSendEmailToProvider(v bool) {
	o.SendEmailToProvider = &v
}


// GetIsFlexible returns the IsFlexible field value if set, zero value otherwise.
func (o *PostCreateInvoiceLinkRequest) GetIsFlexible() bool {
	if o == nil || IsNil(o.IsFlexible) {
		var ret bool
		return ret
	}
	return *o.IsFlexible
}

// GetIsFlexibleOk returns a tuple with the IsFlexible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateInvoiceLinkRequest) GetIsFlexibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFlexible) {
		return nil, false
	}
	return o.IsFlexible, true
}

// HasIsFlexible returns a boolean if a field has been set.
func (o *PostCreateInvoiceLinkRequest) HasIsFlexible() bool {
	if o != nil && !IsNil(o.IsFlexible) {
		return true
	}

	return false
}

// SetIsFlexible gets a reference to the given bool and assigns it to the IsFlexible field.
func (o *PostCreateInvoiceLinkRequest) SetIsFlexible(v bool) {
	o.IsFlexible = &v
}


func (o PostCreateInvoiceLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateInvoiceLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessConnectionId) {
		toSerialize["business_connection_id"] = o.BusinessConnectionId
	}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["payload"] = o.Payload
	if !IsNil(o.ProviderToken) {
		toSerialize["provider_token"] = o.ProviderToken
	}
	toSerialize["currency"] = o.Currency
	toSerialize["prices"] = o.Prices
	if !IsNil(o.SubscriptionPeriod) {
		toSerialize["subscription_period"] = o.SubscriptionPeriod
	}
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

func (o *PostCreateInvoiceLinkRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPostCreateInvoiceLinkRequest := _PostCreateInvoiceLinkRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateInvoiceLinkRequest)

	if err != nil {
		return err
	}

	*o = PostCreateInvoiceLinkRequest(varPostCreateInvoiceLinkRequest)

	return err
}

type NullablePostCreateInvoiceLinkRequest struct {
	value *PostCreateInvoiceLinkRequest
	isSet bool
}

func (v NullablePostCreateInvoiceLinkRequest) Get() *PostCreateInvoiceLinkRequest {
	return v.value
}

func (v *NullablePostCreateInvoiceLinkRequest) Set(val *PostCreateInvoiceLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateInvoiceLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateInvoiceLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateInvoiceLinkRequest(val *PostCreateInvoiceLinkRequest) *NullablePostCreateInvoiceLinkRequest {
	return &NullablePostCreateInvoiceLinkRequest{value: val, isSet: true}
}

func (v NullablePostCreateInvoiceLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateInvoiceLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


