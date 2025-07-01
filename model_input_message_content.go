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


// InputMessageContent This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:  * [InputTextMessageContent](https://core.telegram.org/bots/api/#inputtextmessagecontent) * [InputLocationMessageContent](https://core.telegram.org/bots/api/#inputlocationmessagecontent) * [InputVenueMessageContent](https://core.telegram.org/bots/api/#inputvenuemessagecontent) * [InputContactMessageContent](https://core.telegram.org/bots/api/#inputcontactmessagecontent) * [InputInvoiceMessageContent](https://core.telegram.org/bots/api/#inputinvoicemessagecontent)
type InputMessageContent struct {
	InputContactMessageContent *InputContactMessageContent
	InputInvoiceMessageContent *InputInvoiceMessageContent
	InputLocationMessageContent *InputLocationMessageContent
	InputTextMessageContent *InputTextMessageContent
	InputVenueMessageContent *InputVenueMessageContent
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InputMessageContent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InputContactMessageContent
	err = json.Unmarshal(data, &dst.InputContactMessageContent);
	if err == nil {
		jsonInputContactMessageContent, _ := json.Marshal(dst.InputContactMessageContent)
		if string(jsonInputContactMessageContent) == "{}" { // empty struct
			dst.InputContactMessageContent = nil
		} else {
			return nil // data stored in dst.InputContactMessageContent, return on the first match
		}
	} else {
		dst.InputContactMessageContent = nil
	}

	// try to unmarshal JSON data into InputInvoiceMessageContent
	err = json.Unmarshal(data, &dst.InputInvoiceMessageContent);
	if err == nil {
		jsonInputInvoiceMessageContent, _ := json.Marshal(dst.InputInvoiceMessageContent)
		if string(jsonInputInvoiceMessageContent) == "{}" { // empty struct
			dst.InputInvoiceMessageContent = nil
		} else {
			return nil // data stored in dst.InputInvoiceMessageContent, return on the first match
		}
	} else {
		dst.InputInvoiceMessageContent = nil
	}

	// try to unmarshal JSON data into InputLocationMessageContent
	err = json.Unmarshal(data, &dst.InputLocationMessageContent);
	if err == nil {
		jsonInputLocationMessageContent, _ := json.Marshal(dst.InputLocationMessageContent)
		if string(jsonInputLocationMessageContent) == "{}" { // empty struct
			dst.InputLocationMessageContent = nil
		} else {
			return nil // data stored in dst.InputLocationMessageContent, return on the first match
		}
	} else {
		dst.InputLocationMessageContent = nil
	}

	// try to unmarshal JSON data into InputTextMessageContent
	err = json.Unmarshal(data, &dst.InputTextMessageContent);
	if err == nil {
		jsonInputTextMessageContent, _ := json.Marshal(dst.InputTextMessageContent)
		if string(jsonInputTextMessageContent) == "{}" { // empty struct
			dst.InputTextMessageContent = nil
		} else {
			return nil // data stored in dst.InputTextMessageContent, return on the first match
		}
	} else {
		dst.InputTextMessageContent = nil
	}

	// try to unmarshal JSON data into InputVenueMessageContent
	err = json.Unmarshal(data, &dst.InputVenueMessageContent);
	if err == nil {
		jsonInputVenueMessageContent, _ := json.Marshal(dst.InputVenueMessageContent)
		if string(jsonInputVenueMessageContent) == "{}" { // empty struct
			dst.InputVenueMessageContent = nil
		} else {
			return nil // data stored in dst.InputVenueMessageContent, return on the first match
		}
	} else {
		dst.InputVenueMessageContent = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(InputMessageContent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InputMessageContent) MarshalJSON() ([]byte, error) {
	if src.InputContactMessageContent != nil {
		return json.Marshal(&src.InputContactMessageContent)
	}

	if src.InputInvoiceMessageContent != nil {
		return json.Marshal(&src.InputInvoiceMessageContent)
	}

	if src.InputLocationMessageContent != nil {
		return json.Marshal(&src.InputLocationMessageContent)
	}

	if src.InputTextMessageContent != nil {
		return json.Marshal(&src.InputTextMessageContent)
	}

	if src.InputVenueMessageContent != nil {
		return json.Marshal(&src.InputVenueMessageContent)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableInputMessageContent struct {
	value *InputMessageContent
	isSet bool
}

func (v NullableInputMessageContent) Get() *InputMessageContent {
	return v.value
}

func (v *NullableInputMessageContent) Set(val *InputMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputMessageContent(val *InputMessageContent) *NullableInputMessageContent {
	return &NullableInputMessageContent{value: val, isSet: true}
}

func (v NullableInputMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


