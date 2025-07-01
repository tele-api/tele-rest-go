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


// MenuButton This object describes the bot's menu button in a private chat. It should be one of  * [MenuButtonCommands](https://core.telegram.org/bots/api/#menubuttoncommands) * [MenuButtonWebApp](https://core.telegram.org/bots/api/#menubuttonwebapp) * [MenuButtonDefault](https://core.telegram.org/bots/api/#menubuttondefault)
type MenuButton struct {
	MenuButtonCommands *MenuButtonCommands
	MenuButtonDefault *MenuButtonDefault
	MenuButtonWebApp *MenuButtonWebApp
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MenuButton) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MenuButtonCommands
	err = json.Unmarshal(data, &dst.MenuButtonCommands);
	if err == nil {
		jsonMenuButtonCommands, _ := json.Marshal(dst.MenuButtonCommands)
		if string(jsonMenuButtonCommands) == "{}" { // empty struct
			dst.MenuButtonCommands = nil
		} else {
			return nil // data stored in dst.MenuButtonCommands, return on the first match
		}
	} else {
		dst.MenuButtonCommands = nil
	}

	// try to unmarshal JSON data into MenuButtonDefault
	err = json.Unmarshal(data, &dst.MenuButtonDefault);
	if err == nil {
		jsonMenuButtonDefault, _ := json.Marshal(dst.MenuButtonDefault)
		if string(jsonMenuButtonDefault) == "{}" { // empty struct
			dst.MenuButtonDefault = nil
		} else {
			return nil // data stored in dst.MenuButtonDefault, return on the first match
		}
	} else {
		dst.MenuButtonDefault = nil
	}

	// try to unmarshal JSON data into MenuButtonWebApp
	err = json.Unmarshal(data, &dst.MenuButtonWebApp);
	if err == nil {
		jsonMenuButtonWebApp, _ := json.Marshal(dst.MenuButtonWebApp)
		if string(jsonMenuButtonWebApp) == "{}" { // empty struct
			dst.MenuButtonWebApp = nil
		} else {
			return nil // data stored in dst.MenuButtonWebApp, return on the first match
		}
	} else {
		dst.MenuButtonWebApp = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MenuButton)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MenuButton) MarshalJSON() ([]byte, error) {
	if src.MenuButtonCommands != nil {
		return json.Marshal(&src.MenuButtonCommands)
	}

	if src.MenuButtonDefault != nil {
		return json.Marshal(&src.MenuButtonDefault)
	}

	if src.MenuButtonWebApp != nil {
		return json.Marshal(&src.MenuButtonWebApp)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableMenuButton struct {
	value *MenuButton
	isSet bool
}

func (v NullableMenuButton) Get() *MenuButton {
	return v.value
}

func (v *NullableMenuButton) Set(val *MenuButton) {
	v.value = val
	v.isSet = true
}

func (v NullableMenuButton) IsSet() bool {
	return v.isSet
}

func (v *NullableMenuButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMenuButton(val *MenuButton) *NullableMenuButton {
	return &NullableMenuButton{value: val, isSet: true}
}

func (v NullableMenuButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMenuButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


