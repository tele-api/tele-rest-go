/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// checks if the SetWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetWebhookRequest{}

// SetWebhookRequest Request parameters for setWebhook
type SetWebhookRequest struct {
	// HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Url string `json:"url"`
	Certificate interface{} `json:"certificate,omitempty"`
	// The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	IpAddress *string `json:"ip_address,omitempty"`
	// The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to *40*. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// A JSON-serialized list of the update types you want your bot to receive. For example, specify `[\"message\", \"edited_channel_post\", \"callback_query\"]` to only receive updates of these types. See [Update](https://core.telegram.org/bots/api/#update) for a complete list of available update types. Specify an empty list to receive all update types except *chat\\_member*, *message\\_reaction*, and *message\\_reaction\\_count* (default). If not specified, the previous setting will be used.   Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
	// Pass *True* to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters `A-Z`, `a-z`, `0-9`, `_` and `-` are allowed. The header is useful to ensure that the request comes from a webhook set by you.
	SecretToken *string `json:"secret_token,omitempty"`
}

type _SetWebhookRequest SetWebhookRequest

// NewSetWebhookRequest instantiates a new SetWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetWebhookRequest(url string) *SetWebhookRequest {
	this := SetWebhookRequest{}
	this.Url = url
	var maxConnections int32 = 40
	this.MaxConnections = &maxConnections
	return &this
}

// NewSetWebhookRequestWithDefaults instantiates a new SetWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetWebhookRequestWithDefaults() *SetWebhookRequest {
	this := SetWebhookRequest{}
	var maxConnections int32 = 40
	this.MaxConnections = &maxConnections
	return &this
}

// GetUrl returns the Url field value
func (o *SetWebhookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SetWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SetWebhookRequest) SetUrl(v string) {
	o.Url = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetWebhookRequest) GetCertificate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetWebhookRequest) GetCertificateOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SetWebhookRequest) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given interface{} and assigns it to the Certificate field.
func (o *SetWebhookRequest) SetCertificate(v interface{}) {
	o.Certificate = v
}


// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SetWebhookRequest) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebhookRequest) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SetWebhookRequest) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SetWebhookRequest) SetIpAddress(v string) {
	o.IpAddress = &v
}


// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *SetWebhookRequest) GetMaxConnections() int32 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebhookRequest) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *SetWebhookRequest) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int32 and assigns it to the MaxConnections field.
func (o *SetWebhookRequest) SetMaxConnections(v int32) {
	o.MaxConnections = &v
}


// GetAllowedUpdates returns the AllowedUpdates field value if set, zero value otherwise.
func (o *SetWebhookRequest) GetAllowedUpdates() []string {
	if o == nil || IsNil(o.AllowedUpdates) {
		var ret []string
		return ret
	}
	return o.AllowedUpdates
}

// GetAllowedUpdatesOk returns a tuple with the AllowedUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebhookRequest) GetAllowedUpdatesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedUpdates) {
		return nil, false
	}
	return o.AllowedUpdates, true
}

// HasAllowedUpdates returns a boolean if a field has been set.
func (o *SetWebhookRequest) HasAllowedUpdates() bool {
	if o != nil && !IsNil(o.AllowedUpdates) {
		return true
	}

	return false
}

// SetAllowedUpdates gets a reference to the given []string and assigns it to the AllowedUpdates field.
func (o *SetWebhookRequest) SetAllowedUpdates(v []string) {
	o.AllowedUpdates = v
}


// GetDropPendingUpdates returns the DropPendingUpdates field value if set, zero value otherwise.
func (o *SetWebhookRequest) GetDropPendingUpdates() bool {
	if o == nil || IsNil(o.DropPendingUpdates) {
		var ret bool
		return ret
	}
	return *o.DropPendingUpdates
}

// GetDropPendingUpdatesOk returns a tuple with the DropPendingUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebhookRequest) GetDropPendingUpdatesOk() (*bool, bool) {
	if o == nil || IsNil(o.DropPendingUpdates) {
		return nil, false
	}
	return o.DropPendingUpdates, true
}

// HasDropPendingUpdates returns a boolean if a field has been set.
func (o *SetWebhookRequest) HasDropPendingUpdates() bool {
	if o != nil && !IsNil(o.DropPendingUpdates) {
		return true
	}

	return false
}

// SetDropPendingUpdates gets a reference to the given bool and assigns it to the DropPendingUpdates field.
func (o *SetWebhookRequest) SetDropPendingUpdates(v bool) {
	o.DropPendingUpdates = &v
}


// GetSecretToken returns the SecretToken field value if set, zero value otherwise.
func (o *SetWebhookRequest) GetSecretToken() string {
	if o == nil || IsNil(o.SecretToken) {
		var ret string
		return ret
	}
	return *o.SecretToken
}

// GetSecretTokenOk returns a tuple with the SecretToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebhookRequest) GetSecretTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SecretToken) {
		return nil, false
	}
	return o.SecretToken, true
}

// HasSecretToken returns a boolean if a field has been set.
func (o *SetWebhookRequest) HasSecretToken() bool {
	if o != nil && !IsNil(o.SecretToken) {
		return true
	}

	return false
}

// SetSecretToken gets a reference to the given string and assigns it to the SecretToken field.
func (o *SetWebhookRequest) SetSecretToken(v string) {
	o.SecretToken = &v
}


func (o SetWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["max_connections"] = o.MaxConnections
	}
	if !IsNil(o.AllowedUpdates) {
		toSerialize["allowed_updates"] = o.AllowedUpdates
	}
	if !IsNil(o.DropPendingUpdates) {
		toSerialize["drop_pending_updates"] = o.DropPendingUpdates
	}
	if !IsNil(o.SecretToken) {
		toSerialize["secret_token"] = o.SecretToken
	}
	return toSerialize, nil
}

func (o *SetWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varSetWebhookRequest := _SetWebhookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetWebhookRequest)

	if err != nil {
		return err
	}

	*o = SetWebhookRequest(varSetWebhookRequest)

	return err
}

type NullableSetWebhookRequest struct {
	value *SetWebhookRequest
	isSet bool
}

func (v NullableSetWebhookRequest) Get() *SetWebhookRequest {
	return v.value
}

func (v *NullableSetWebhookRequest) Set(val *SetWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetWebhookRequest(val *SetWebhookRequest) *NullableSetWebhookRequest {
	return &NullableSetWebhookRequest{value: val, isSet: true}
}

func (v NullableSetWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


