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


// TransactionPartner This object describes the source of a transaction, or its recipient for outgoing transactions. Currently, it can be one of  * [TransactionPartnerUser](https://core.telegram.org/bots/api/#transactionpartneruser) * [TransactionPartnerChat](https://core.telegram.org/bots/api/#transactionpartnerchat) * [TransactionPartnerAffiliateProgram](https://core.telegram.org/bots/api/#transactionpartneraffiliateprogram) * [TransactionPartnerFragment](https://core.telegram.org/bots/api/#transactionpartnerfragment) * [TransactionPartnerTelegramAds](https://core.telegram.org/bots/api/#transactionpartnertelegramads) * [TransactionPartnerTelegramApi](https://core.telegram.org/bots/api/#transactionpartnertelegramapi) * [TransactionPartnerOther](https://core.telegram.org/bots/api/#transactionpartnerother)
type TransactionPartner struct {
	TransactionPartnerAffiliateProgram *TransactionPartnerAffiliateProgram
	TransactionPartnerChat *TransactionPartnerChat
	TransactionPartnerFragment *TransactionPartnerFragment
	TransactionPartnerOther *TransactionPartnerOther
	TransactionPartnerTelegramAds *TransactionPartnerTelegramAds
	TransactionPartnerTelegramApi *TransactionPartnerTelegramApi
	TransactionPartnerUser *TransactionPartnerUser
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionPartner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransactionPartnerAffiliateProgram
	err = json.Unmarshal(data, &dst.TransactionPartnerAffiliateProgram);
	if err == nil {
		jsonTransactionPartnerAffiliateProgram, _ := json.Marshal(dst.TransactionPartnerAffiliateProgram)
		if string(jsonTransactionPartnerAffiliateProgram) == "{}" { // empty struct
			dst.TransactionPartnerAffiliateProgram = nil
		} else {
			return nil // data stored in dst.TransactionPartnerAffiliateProgram, return on the first match
		}
	} else {
		dst.TransactionPartnerAffiliateProgram = nil
	}

	// try to unmarshal JSON data into TransactionPartnerChat
	err = json.Unmarshal(data, &dst.TransactionPartnerChat);
	if err == nil {
		jsonTransactionPartnerChat, _ := json.Marshal(dst.TransactionPartnerChat)
		if string(jsonTransactionPartnerChat) == "{}" { // empty struct
			dst.TransactionPartnerChat = nil
		} else {
			return nil // data stored in dst.TransactionPartnerChat, return on the first match
		}
	} else {
		dst.TransactionPartnerChat = nil
	}

	// try to unmarshal JSON data into TransactionPartnerFragment
	err = json.Unmarshal(data, &dst.TransactionPartnerFragment);
	if err == nil {
		jsonTransactionPartnerFragment, _ := json.Marshal(dst.TransactionPartnerFragment)
		if string(jsonTransactionPartnerFragment) == "{}" { // empty struct
			dst.TransactionPartnerFragment = nil
		} else {
			return nil // data stored in dst.TransactionPartnerFragment, return on the first match
		}
	} else {
		dst.TransactionPartnerFragment = nil
	}

	// try to unmarshal JSON data into TransactionPartnerOther
	err = json.Unmarshal(data, &dst.TransactionPartnerOther);
	if err == nil {
		jsonTransactionPartnerOther, _ := json.Marshal(dst.TransactionPartnerOther)
		if string(jsonTransactionPartnerOther) == "{}" { // empty struct
			dst.TransactionPartnerOther = nil
		} else {
			return nil // data stored in dst.TransactionPartnerOther, return on the first match
		}
	} else {
		dst.TransactionPartnerOther = nil
	}

	// try to unmarshal JSON data into TransactionPartnerTelegramAds
	err = json.Unmarshal(data, &dst.TransactionPartnerTelegramAds);
	if err == nil {
		jsonTransactionPartnerTelegramAds, _ := json.Marshal(dst.TransactionPartnerTelegramAds)
		if string(jsonTransactionPartnerTelegramAds) == "{}" { // empty struct
			dst.TransactionPartnerTelegramAds = nil
		} else {
			return nil // data stored in dst.TransactionPartnerTelegramAds, return on the first match
		}
	} else {
		dst.TransactionPartnerTelegramAds = nil
	}

	// try to unmarshal JSON data into TransactionPartnerTelegramApi
	err = json.Unmarshal(data, &dst.TransactionPartnerTelegramApi);
	if err == nil {
		jsonTransactionPartnerTelegramApi, _ := json.Marshal(dst.TransactionPartnerTelegramApi)
		if string(jsonTransactionPartnerTelegramApi) == "{}" { // empty struct
			dst.TransactionPartnerTelegramApi = nil
		} else {
			return nil // data stored in dst.TransactionPartnerTelegramApi, return on the first match
		}
	} else {
		dst.TransactionPartnerTelegramApi = nil
	}

	// try to unmarshal JSON data into TransactionPartnerUser
	err = json.Unmarshal(data, &dst.TransactionPartnerUser);
	if err == nil {
		jsonTransactionPartnerUser, _ := json.Marshal(dst.TransactionPartnerUser)
		if string(jsonTransactionPartnerUser) == "{}" { // empty struct
			dst.TransactionPartnerUser = nil
		} else {
			return nil // data stored in dst.TransactionPartnerUser, return on the first match
		}
	} else {
		dst.TransactionPartnerUser = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionPartner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionPartner) MarshalJSON() ([]byte, error) {
	if src.TransactionPartnerAffiliateProgram != nil {
		return json.Marshal(&src.TransactionPartnerAffiliateProgram)
	}

	if src.TransactionPartnerChat != nil {
		return json.Marshal(&src.TransactionPartnerChat)
	}

	if src.TransactionPartnerFragment != nil {
		return json.Marshal(&src.TransactionPartnerFragment)
	}

	if src.TransactionPartnerOther != nil {
		return json.Marshal(&src.TransactionPartnerOther)
	}

	if src.TransactionPartnerTelegramAds != nil {
		return json.Marshal(&src.TransactionPartnerTelegramAds)
	}

	if src.TransactionPartnerTelegramApi != nil {
		return json.Marshal(&src.TransactionPartnerTelegramApi)
	}

	if src.TransactionPartnerUser != nil {
		return json.Marshal(&src.TransactionPartnerUser)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableTransactionPartner struct {
	value *TransactionPartner
	isSet bool
}

func (v NullableTransactionPartner) Get() *TransactionPartner {
	return v.value
}

func (v *NullableTransactionPartner) Set(val *TransactionPartner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPartner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPartner(val *TransactionPartner) *NullableTransactionPartner {
	return &NullableTransactionPartner{value: val, isSet: true}
}

func (v NullableTransactionPartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


