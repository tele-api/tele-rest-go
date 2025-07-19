/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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
	"gopkg.in/validator.v2"
)

// InputMessageContent - This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:  * [InputTextMessageContent](https://core.telegram.org/bots/api/#inputtextmessagecontent) * [InputLocationMessageContent](https://core.telegram.org/bots/api/#inputlocationmessagecontent) * [InputVenueMessageContent](https://core.telegram.org/bots/api/#inputvenuemessagecontent) * [InputContactMessageContent](https://core.telegram.org/bots/api/#inputcontactmessagecontent) * [InputInvoiceMessageContent](https://core.telegram.org/bots/api/#inputinvoicemessagecontent)
type InputMessageContent struct {
	InputContactMessageContent *InputContactMessageContent
	InputInvoiceMessageContent *InputInvoiceMessageContent
	InputLocationMessageContent *InputLocationMessageContent
	InputTextMessageContent *InputTextMessageContent
	InputVenueMessageContent *InputVenueMessageContent
}

// InputContactMessageContentAsInputMessageContent is a convenience function that returns InputContactMessageContent wrapped in InputMessageContent
func InputContactMessageContentAsInputMessageContent(v *InputContactMessageContent) InputMessageContent {
	return InputMessageContent{
		InputContactMessageContent: v,
	}
}

// InputInvoiceMessageContentAsInputMessageContent is a convenience function that returns InputInvoiceMessageContent wrapped in InputMessageContent
func InputInvoiceMessageContentAsInputMessageContent(v *InputInvoiceMessageContent) InputMessageContent {
	return InputMessageContent{
		InputInvoiceMessageContent: v,
	}
}

// InputLocationMessageContentAsInputMessageContent is a convenience function that returns InputLocationMessageContent wrapped in InputMessageContent
func InputLocationMessageContentAsInputMessageContent(v *InputLocationMessageContent) InputMessageContent {
	return InputMessageContent{
		InputLocationMessageContent: v,
	}
}

// InputTextMessageContentAsInputMessageContent is a convenience function that returns InputTextMessageContent wrapped in InputMessageContent
func InputTextMessageContentAsInputMessageContent(v *InputTextMessageContent) InputMessageContent {
	return InputMessageContent{
		InputTextMessageContent: v,
	}
}

// InputVenueMessageContentAsInputMessageContent is a convenience function that returns InputVenueMessageContent wrapped in InputMessageContent
func InputVenueMessageContentAsInputMessageContent(v *InputVenueMessageContent) InputMessageContent {
	return InputMessageContent{
		InputVenueMessageContent: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InputMessageContent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InputContactMessageContent
	err = newStrictDecoder(data).Decode(&dst.InputContactMessageContent)
	if err == nil {
		jsonInputContactMessageContent, _ := json.Marshal(dst.InputContactMessageContent)
		if string(jsonInputContactMessageContent) == "{}" { // empty struct
			dst.InputContactMessageContent = nil
		} else {
			if err = validator.Validate(dst.InputContactMessageContent); err != nil {
				dst.InputContactMessageContent = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputContactMessageContent = nil
	}

	// try to unmarshal data into InputInvoiceMessageContent
	err = newStrictDecoder(data).Decode(&dst.InputInvoiceMessageContent)
	if err == nil {
		jsonInputInvoiceMessageContent, _ := json.Marshal(dst.InputInvoiceMessageContent)
		if string(jsonInputInvoiceMessageContent) == "{}" { // empty struct
			dst.InputInvoiceMessageContent = nil
		} else {
			if err = validator.Validate(dst.InputInvoiceMessageContent); err != nil {
				dst.InputInvoiceMessageContent = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputInvoiceMessageContent = nil
	}

	// try to unmarshal data into InputLocationMessageContent
	err = newStrictDecoder(data).Decode(&dst.InputLocationMessageContent)
	if err == nil {
		jsonInputLocationMessageContent, _ := json.Marshal(dst.InputLocationMessageContent)
		if string(jsonInputLocationMessageContent) == "{}" { // empty struct
			dst.InputLocationMessageContent = nil
		} else {
			if err = validator.Validate(dst.InputLocationMessageContent); err != nil {
				dst.InputLocationMessageContent = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputLocationMessageContent = nil
	}

	// try to unmarshal data into InputTextMessageContent
	err = newStrictDecoder(data).Decode(&dst.InputTextMessageContent)
	if err == nil {
		jsonInputTextMessageContent, _ := json.Marshal(dst.InputTextMessageContent)
		if string(jsonInputTextMessageContent) == "{}" { // empty struct
			dst.InputTextMessageContent = nil
		} else {
			if err = validator.Validate(dst.InputTextMessageContent); err != nil {
				dst.InputTextMessageContent = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputTextMessageContent = nil
	}

	// try to unmarshal data into InputVenueMessageContent
	err = newStrictDecoder(data).Decode(&dst.InputVenueMessageContent)
	if err == nil {
		jsonInputVenueMessageContent, _ := json.Marshal(dst.InputVenueMessageContent)
		if string(jsonInputVenueMessageContent) == "{}" { // empty struct
			dst.InputVenueMessageContent = nil
		} else {
			if err = validator.Validate(dst.InputVenueMessageContent); err != nil {
				dst.InputVenueMessageContent = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputVenueMessageContent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InputContactMessageContent = nil
		dst.InputInvoiceMessageContent = nil
		dst.InputLocationMessageContent = nil
		dst.InputTextMessageContent = nil
		dst.InputVenueMessageContent = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InputMessageContent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InputMessageContent)")
	}
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

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InputMessageContent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InputContactMessageContent != nil {
		return obj.InputContactMessageContent
	}

	if obj.InputInvoiceMessageContent != nil {
		return obj.InputInvoiceMessageContent
	}

	if obj.InputLocationMessageContent != nil {
		return obj.InputLocationMessageContent
	}

	if obj.InputTextMessageContent != nil {
		return obj.InputTextMessageContent
	}

	if obj.InputVenueMessageContent != nil {
		return obj.InputVenueMessageContent
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InputMessageContent) GetActualInstanceValue() (interface{}) {
	if obj.InputContactMessageContent != nil {
		return *obj.InputContactMessageContent
	}

	if obj.InputInvoiceMessageContent != nil {
		return *obj.InputInvoiceMessageContent
	}

	if obj.InputLocationMessageContent != nil {
		return *obj.InputLocationMessageContent
	}

	if obj.InputTextMessageContent != nil {
		return *obj.InputTextMessageContent
	}

	if obj.InputVenueMessageContent != nil {
		return *obj.InputVenueMessageContent
	}

	// all schemas are nil
	return nil
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


