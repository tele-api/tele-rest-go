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
	"fmt"
	"gopkg.in/validator.v2"
)

// RevenueWithdrawalState - This object describes the state of a revenue withdrawal operation. Currently, it can be one of  * [RevenueWithdrawalStatePending](https://core.telegram.org/bots/api/#revenuewithdrawalstatepending) * [RevenueWithdrawalStateSucceeded](https://core.telegram.org/bots/api/#revenuewithdrawalstatesucceeded) * [RevenueWithdrawalStateFailed](https://core.telegram.org/bots/api/#revenuewithdrawalstatefailed)
type RevenueWithdrawalState struct {
	RevenueWithdrawalStateFailed *RevenueWithdrawalStateFailed
	RevenueWithdrawalStatePending *RevenueWithdrawalStatePending
	RevenueWithdrawalStateSucceeded *RevenueWithdrawalStateSucceeded
}

// RevenueWithdrawalStateFailedAsRevenueWithdrawalState is a convenience function that returns RevenueWithdrawalStateFailed wrapped in RevenueWithdrawalState
func RevenueWithdrawalStateFailedAsRevenueWithdrawalState(v *RevenueWithdrawalStateFailed) RevenueWithdrawalState {
	return RevenueWithdrawalState{
		RevenueWithdrawalStateFailed: v,
	}
}

// RevenueWithdrawalStatePendingAsRevenueWithdrawalState is a convenience function that returns RevenueWithdrawalStatePending wrapped in RevenueWithdrawalState
func RevenueWithdrawalStatePendingAsRevenueWithdrawalState(v *RevenueWithdrawalStatePending) RevenueWithdrawalState {
	return RevenueWithdrawalState{
		RevenueWithdrawalStatePending: v,
	}
}

// RevenueWithdrawalStateSucceededAsRevenueWithdrawalState is a convenience function that returns RevenueWithdrawalStateSucceeded wrapped in RevenueWithdrawalState
func RevenueWithdrawalStateSucceededAsRevenueWithdrawalState(v *RevenueWithdrawalStateSucceeded) RevenueWithdrawalState {
	return RevenueWithdrawalState{
		RevenueWithdrawalStateSucceeded: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RevenueWithdrawalStateFailed
	err = newStrictDecoder(data).Decode(&dst.RevenueWithdrawalStateFailed)
	if err == nil {
		jsonRevenueWithdrawalStateFailed, _ := json.Marshal(dst.RevenueWithdrawalStateFailed)
		if string(jsonRevenueWithdrawalStateFailed) == "{}" { // empty struct
			dst.RevenueWithdrawalStateFailed = nil
		} else {
			if err = validator.Validate(dst.RevenueWithdrawalStateFailed); err != nil {
				dst.RevenueWithdrawalStateFailed = nil
			} else {
				match++
			}
		}
	} else {
		dst.RevenueWithdrawalStateFailed = nil
	}

	// try to unmarshal data into RevenueWithdrawalStatePending
	err = newStrictDecoder(data).Decode(&dst.RevenueWithdrawalStatePending)
	if err == nil {
		jsonRevenueWithdrawalStatePending, _ := json.Marshal(dst.RevenueWithdrawalStatePending)
		if string(jsonRevenueWithdrawalStatePending) == "{}" { // empty struct
			dst.RevenueWithdrawalStatePending = nil
		} else {
			if err = validator.Validate(dst.RevenueWithdrawalStatePending); err != nil {
				dst.RevenueWithdrawalStatePending = nil
			} else {
				match++
			}
		}
	} else {
		dst.RevenueWithdrawalStatePending = nil
	}

	// try to unmarshal data into RevenueWithdrawalStateSucceeded
	err = newStrictDecoder(data).Decode(&dst.RevenueWithdrawalStateSucceeded)
	if err == nil {
		jsonRevenueWithdrawalStateSucceeded, _ := json.Marshal(dst.RevenueWithdrawalStateSucceeded)
		if string(jsonRevenueWithdrawalStateSucceeded) == "{}" { // empty struct
			dst.RevenueWithdrawalStateSucceeded = nil
		} else {
			if err = validator.Validate(dst.RevenueWithdrawalStateSucceeded); err != nil {
				dst.RevenueWithdrawalStateSucceeded = nil
			} else {
				match++
			}
		}
	} else {
		dst.RevenueWithdrawalStateSucceeded = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RevenueWithdrawalStateFailed = nil
		dst.RevenueWithdrawalStatePending = nil
		dst.RevenueWithdrawalStateSucceeded = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RevenueWithdrawalState)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RevenueWithdrawalState)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RevenueWithdrawalState) MarshalJSON() ([]byte, error) {
	if src.RevenueWithdrawalStateFailed != nil {
		return json.Marshal(&src.RevenueWithdrawalStateFailed)
	}

	if src.RevenueWithdrawalStatePending != nil {
		return json.Marshal(&src.RevenueWithdrawalStatePending)
	}

	if src.RevenueWithdrawalStateSucceeded != nil {
		return json.Marshal(&src.RevenueWithdrawalStateSucceeded)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RevenueWithdrawalState) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RevenueWithdrawalStateFailed != nil {
		return obj.RevenueWithdrawalStateFailed
	}

	if obj.RevenueWithdrawalStatePending != nil {
		return obj.RevenueWithdrawalStatePending
	}

	if obj.RevenueWithdrawalStateSucceeded != nil {
		return obj.RevenueWithdrawalStateSucceeded
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RevenueWithdrawalState) GetActualInstanceValue() (interface{}) {
	if obj.RevenueWithdrawalStateFailed != nil {
		return *obj.RevenueWithdrawalStateFailed
	}

	if obj.RevenueWithdrawalStatePending != nil {
		return *obj.RevenueWithdrawalStatePending
	}

	if obj.RevenueWithdrawalStateSucceeded != nil {
		return *obj.RevenueWithdrawalStateSucceeded
	}

	// all schemas are nil
	return nil
}

type NullableRevenueWithdrawalState struct {
	value *RevenueWithdrawalState
	isSet bool
}

func (v NullableRevenueWithdrawalState) Get() *RevenueWithdrawalState {
	return v.value
}

func (v *NullableRevenueWithdrawalState) Set(val *RevenueWithdrawalState) {
	v.value = val
	v.isSet = true
}

func (v NullableRevenueWithdrawalState) IsSet() bool {
	return v.isSet
}

func (v *NullableRevenueWithdrawalState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevenueWithdrawalState(val *RevenueWithdrawalState) *NullableRevenueWithdrawalState {
	return &NullableRevenueWithdrawalState{value: val, isSet: true}
}

func (v NullableRevenueWithdrawalState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevenueWithdrawalState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


