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


// RevenueWithdrawalState This object describes the state of a revenue withdrawal operation. Currently, it can be one of  * [RevenueWithdrawalStatePending](https://core.telegram.org/bots/api/#revenuewithdrawalstatepending) * [RevenueWithdrawalStateSucceeded](https://core.telegram.org/bots/api/#revenuewithdrawalstatesucceeded) * [RevenueWithdrawalStateFailed](https://core.telegram.org/bots/api/#revenuewithdrawalstatefailed)
type RevenueWithdrawalState struct {
	RevenueWithdrawalStateFailed *RevenueWithdrawalStateFailed
	RevenueWithdrawalStatePending *RevenueWithdrawalStatePending
	RevenueWithdrawalStateSucceeded *RevenueWithdrawalStateSucceeded
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RevenueWithdrawalStateFailed
	err = json.Unmarshal(data, &dst.RevenueWithdrawalStateFailed);
	if err == nil {
		jsonRevenueWithdrawalStateFailed, _ := json.Marshal(dst.RevenueWithdrawalStateFailed)
		if string(jsonRevenueWithdrawalStateFailed) == "{}" { // empty struct
			dst.RevenueWithdrawalStateFailed = nil
		} else {
			return nil // data stored in dst.RevenueWithdrawalStateFailed, return on the first match
		}
	} else {
		dst.RevenueWithdrawalStateFailed = nil
	}

	// try to unmarshal JSON data into RevenueWithdrawalStatePending
	err = json.Unmarshal(data, &dst.RevenueWithdrawalStatePending);
	if err == nil {
		jsonRevenueWithdrawalStatePending, _ := json.Marshal(dst.RevenueWithdrawalStatePending)
		if string(jsonRevenueWithdrawalStatePending) == "{}" { // empty struct
			dst.RevenueWithdrawalStatePending = nil
		} else {
			return nil // data stored in dst.RevenueWithdrawalStatePending, return on the first match
		}
	} else {
		dst.RevenueWithdrawalStatePending = nil
	}

	// try to unmarshal JSON data into RevenueWithdrawalStateSucceeded
	err = json.Unmarshal(data, &dst.RevenueWithdrawalStateSucceeded);
	if err == nil {
		jsonRevenueWithdrawalStateSucceeded, _ := json.Marshal(dst.RevenueWithdrawalStateSucceeded)
		if string(jsonRevenueWithdrawalStateSucceeded) == "{}" { // empty struct
			dst.RevenueWithdrawalStateSucceeded = nil
		} else {
			return nil // data stored in dst.RevenueWithdrawalStateSucceeded, return on the first match
		}
	} else {
		dst.RevenueWithdrawalStateSucceeded = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RevenueWithdrawalState)")
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

	return nil, nil // no data in anyOf schemas
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


