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
)

// checks if the DeleteWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteWebhookRequest{}

// DeleteWebhookRequest Request parameters for deleteWebhook
type DeleteWebhookRequest struct {
	// Pass *True* to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
}

// NewDeleteWebhookRequest instantiates a new DeleteWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteWebhookRequest() *DeleteWebhookRequest {
	this := DeleteWebhookRequest{}
	return &this
}

// NewDeleteWebhookRequestWithDefaults instantiates a new DeleteWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteWebhookRequestWithDefaults() *DeleteWebhookRequest {
	this := DeleteWebhookRequest{}
	return &this
}

// GetDropPendingUpdates returns the DropPendingUpdates field value if set, zero value otherwise.
func (o *DeleteWebhookRequest) GetDropPendingUpdates() bool {
	if o == nil || IsNil(o.DropPendingUpdates) {
		var ret bool
		return ret
	}
	return *o.DropPendingUpdates
}

// GetDropPendingUpdatesOk returns a tuple with the DropPendingUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteWebhookRequest) GetDropPendingUpdatesOk() (*bool, bool) {
	if o == nil || IsNil(o.DropPendingUpdates) {
		return nil, false
	}
	return o.DropPendingUpdates, true
}

// HasDropPendingUpdates returns a boolean if a field has been set.
func (o *DeleteWebhookRequest) HasDropPendingUpdates() bool {
	if o != nil && !IsNil(o.DropPendingUpdates) {
		return true
	}

	return false
}

// SetDropPendingUpdates gets a reference to the given bool and assigns it to the DropPendingUpdates field.
func (o *DeleteWebhookRequest) SetDropPendingUpdates(v bool) {
	o.DropPendingUpdates = &v
}


func (o DeleteWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DropPendingUpdates) {
		toSerialize["drop_pending_updates"] = o.DropPendingUpdates
	}
	return toSerialize, nil
}

type NullableDeleteWebhookRequest struct {
	value *DeleteWebhookRequest
	isSet bool
}

func (v NullableDeleteWebhookRequest) Get() *DeleteWebhookRequest {
	return v.value
}

func (v *NullableDeleteWebhookRequest) Set(val *DeleteWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteWebhookRequest(val *DeleteWebhookRequest) *NullableDeleteWebhookRequest {
	return &NullableDeleteWebhookRequest{value: val, isSet: true}
}

func (v NullableDeleteWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


