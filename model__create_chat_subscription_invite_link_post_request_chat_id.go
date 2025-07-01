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
	"fmt"
)


// CreateChatSubscriptionInviteLinkPostRequestChatId Unique identifier for the target channel chat or username of the target channel (in the format `@channelusername`)
type CreateChatSubscriptionInviteLinkPostRequestChatId struct {
	Int32 *int32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateChatSubscriptionInviteLinkPostRequestChatId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreateChatSubscriptionInviteLinkPostRequestChatId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateChatSubscriptionInviteLinkPostRequestChatId) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCreateChatSubscriptionInviteLinkPostRequestChatId struct {
	value *CreateChatSubscriptionInviteLinkPostRequestChatId
	isSet bool
}

func (v NullableCreateChatSubscriptionInviteLinkPostRequestChatId) Get() *CreateChatSubscriptionInviteLinkPostRequestChatId {
	return v.value
}

func (v *NullableCreateChatSubscriptionInviteLinkPostRequestChatId) Set(val *CreateChatSubscriptionInviteLinkPostRequestChatId) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatSubscriptionInviteLinkPostRequestChatId) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatSubscriptionInviteLinkPostRequestChatId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatSubscriptionInviteLinkPostRequestChatId(val *CreateChatSubscriptionInviteLinkPostRequestChatId) *NullableCreateChatSubscriptionInviteLinkPostRequestChatId {
	return &NullableCreateChatSubscriptionInviteLinkPostRequestChatId{value: val, isSet: true}
}

func (v NullableCreateChatSubscriptionInviteLinkPostRequestChatId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatSubscriptionInviteLinkPostRequestChatId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


