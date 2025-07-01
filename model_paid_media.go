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
	"fmt"
	"gopkg.in/validator.v2"
)

// PaidMedia - This object describes paid media. Currently, it can be one of  * [PaidMediaPreview](https://core.telegram.org/bots/api/#paidmediapreview) * [PaidMediaPhoto](https://core.telegram.org/bots/api/#paidmediaphoto) * [PaidMediaVideo](https://core.telegram.org/bots/api/#paidmediavideo)
type PaidMedia struct {
	PaidMediaPhoto *PaidMediaPhoto
	PaidMediaPreview *PaidMediaPreview
	PaidMediaVideo *PaidMediaVideo
}

// PaidMediaPhotoAsPaidMedia is a convenience function that returns PaidMediaPhoto wrapped in PaidMedia
func PaidMediaPhotoAsPaidMedia(v *PaidMediaPhoto) PaidMedia {
	return PaidMedia{
		PaidMediaPhoto: v,
	}
}

// PaidMediaPreviewAsPaidMedia is a convenience function that returns PaidMediaPreview wrapped in PaidMedia
func PaidMediaPreviewAsPaidMedia(v *PaidMediaPreview) PaidMedia {
	return PaidMedia{
		PaidMediaPreview: v,
	}
}

// PaidMediaVideoAsPaidMedia is a convenience function that returns PaidMediaVideo wrapped in PaidMedia
func PaidMediaVideoAsPaidMedia(v *PaidMediaVideo) PaidMedia {
	return PaidMedia{
		PaidMediaVideo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaidMedia) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PaidMediaPhoto
	err = newStrictDecoder(data).Decode(&dst.PaidMediaPhoto)
	if err == nil {
		jsonPaidMediaPhoto, _ := json.Marshal(dst.PaidMediaPhoto)
		if string(jsonPaidMediaPhoto) == "{}" { // empty struct
			dst.PaidMediaPhoto = nil
		} else {
			if err = validator.Validate(dst.PaidMediaPhoto); err != nil {
				dst.PaidMediaPhoto = nil
			} else {
				match++
			}
		}
	} else {
		dst.PaidMediaPhoto = nil
	}

	// try to unmarshal data into PaidMediaPreview
	err = newStrictDecoder(data).Decode(&dst.PaidMediaPreview)
	if err == nil {
		jsonPaidMediaPreview, _ := json.Marshal(dst.PaidMediaPreview)
		if string(jsonPaidMediaPreview) == "{}" { // empty struct
			dst.PaidMediaPreview = nil
		} else {
			if err = validator.Validate(dst.PaidMediaPreview); err != nil {
				dst.PaidMediaPreview = nil
			} else {
				match++
			}
		}
	} else {
		dst.PaidMediaPreview = nil
	}

	// try to unmarshal data into PaidMediaVideo
	err = newStrictDecoder(data).Decode(&dst.PaidMediaVideo)
	if err == nil {
		jsonPaidMediaVideo, _ := json.Marshal(dst.PaidMediaVideo)
		if string(jsonPaidMediaVideo) == "{}" { // empty struct
			dst.PaidMediaVideo = nil
		} else {
			if err = validator.Validate(dst.PaidMediaVideo); err != nil {
				dst.PaidMediaVideo = nil
			} else {
				match++
			}
		}
	} else {
		dst.PaidMediaVideo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PaidMediaPhoto = nil
		dst.PaidMediaPreview = nil
		dst.PaidMediaVideo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PaidMedia)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PaidMedia)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaidMedia) MarshalJSON() ([]byte, error) {
	if src.PaidMediaPhoto != nil {
		return json.Marshal(&src.PaidMediaPhoto)
	}

	if src.PaidMediaPreview != nil {
		return json.Marshal(&src.PaidMediaPreview)
	}

	if src.PaidMediaVideo != nil {
		return json.Marshal(&src.PaidMediaVideo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaidMedia) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaidMediaPhoto != nil {
		return obj.PaidMediaPhoto
	}

	if obj.PaidMediaPreview != nil {
		return obj.PaidMediaPreview
	}

	if obj.PaidMediaVideo != nil {
		return obj.PaidMediaVideo
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj PaidMedia) GetActualInstanceValue() (interface{}) {
	if obj.PaidMediaPhoto != nil {
		return *obj.PaidMediaPhoto
	}

	if obj.PaidMediaPreview != nil {
		return *obj.PaidMediaPreview
	}

	if obj.PaidMediaVideo != nil {
		return *obj.PaidMediaVideo
	}

	// all schemas are nil
	return nil
}

type NullablePaidMedia struct {
	value *PaidMedia
	isSet bool
}

func (v NullablePaidMedia) Get() *PaidMedia {
	return v.value
}

func (v *NullablePaidMedia) Set(val *PaidMedia) {
	v.value = val
	v.isSet = true
}

func (v NullablePaidMedia) IsSet() bool {
	return v.isSet
}

func (v *NullablePaidMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaidMedia(val *PaidMedia) *NullablePaidMedia {
	return &NullablePaidMedia{value: val, isSet: true}
}

func (v NullablePaidMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaidMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


