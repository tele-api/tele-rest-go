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


// PassportElementError This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:  * [PassportElementErrorDataField](https://core.telegram.org/bots/api/#passportelementerrordatafield) * [PassportElementErrorFrontSide](https://core.telegram.org/bots/api/#passportelementerrorfrontside) * [PassportElementErrorReverseSide](https://core.telegram.org/bots/api/#passportelementerrorreverseside) * [PassportElementErrorSelfie](https://core.telegram.org/bots/api/#passportelementerrorselfie) * [PassportElementErrorFile](https://core.telegram.org/bots/api/#passportelementerrorfile) * [PassportElementErrorFiles](https://core.telegram.org/bots/api/#passportelementerrorfiles) * [PassportElementErrorTranslationFile](https://core.telegram.org/bots/api/#passportelementerrortranslationfile) * [PassportElementErrorTranslationFiles](https://core.telegram.org/bots/api/#passportelementerrortranslationfiles) * [PassportElementErrorUnspecified](https://core.telegram.org/bots/api/#passportelementerrorunspecified)
type PassportElementError struct {
	PassportElementErrorDataField *PassportElementErrorDataField
	PassportElementErrorFile *PassportElementErrorFile
	PassportElementErrorFiles *PassportElementErrorFiles
	PassportElementErrorFrontSide *PassportElementErrorFrontSide
	PassportElementErrorReverseSide *PassportElementErrorReverseSide
	PassportElementErrorSelfie *PassportElementErrorSelfie
	PassportElementErrorTranslationFile *PassportElementErrorTranslationFile
	PassportElementErrorTranslationFiles *PassportElementErrorTranslationFiles
	PassportElementErrorUnspecified *PassportElementErrorUnspecified
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PassportElementError) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PassportElementErrorDataField
	err = json.Unmarshal(data, &dst.PassportElementErrorDataField);
	if err == nil {
		jsonPassportElementErrorDataField, _ := json.Marshal(dst.PassportElementErrorDataField)
		if string(jsonPassportElementErrorDataField) == "{}" { // empty struct
			dst.PassportElementErrorDataField = nil
		} else {
			return nil // data stored in dst.PassportElementErrorDataField, return on the first match
		}
	} else {
		dst.PassportElementErrorDataField = nil
	}

	// try to unmarshal JSON data into PassportElementErrorFile
	err = json.Unmarshal(data, &dst.PassportElementErrorFile);
	if err == nil {
		jsonPassportElementErrorFile, _ := json.Marshal(dst.PassportElementErrorFile)
		if string(jsonPassportElementErrorFile) == "{}" { // empty struct
			dst.PassportElementErrorFile = nil
		} else {
			return nil // data stored in dst.PassportElementErrorFile, return on the first match
		}
	} else {
		dst.PassportElementErrorFile = nil
	}

	// try to unmarshal JSON data into PassportElementErrorFiles
	err = json.Unmarshal(data, &dst.PassportElementErrorFiles);
	if err == nil {
		jsonPassportElementErrorFiles, _ := json.Marshal(dst.PassportElementErrorFiles)
		if string(jsonPassportElementErrorFiles) == "{}" { // empty struct
			dst.PassportElementErrorFiles = nil
		} else {
			return nil // data stored in dst.PassportElementErrorFiles, return on the first match
		}
	} else {
		dst.PassportElementErrorFiles = nil
	}

	// try to unmarshal JSON data into PassportElementErrorFrontSide
	err = json.Unmarshal(data, &dst.PassportElementErrorFrontSide);
	if err == nil {
		jsonPassportElementErrorFrontSide, _ := json.Marshal(dst.PassportElementErrorFrontSide)
		if string(jsonPassportElementErrorFrontSide) == "{}" { // empty struct
			dst.PassportElementErrorFrontSide = nil
		} else {
			return nil // data stored in dst.PassportElementErrorFrontSide, return on the first match
		}
	} else {
		dst.PassportElementErrorFrontSide = nil
	}

	// try to unmarshal JSON data into PassportElementErrorReverseSide
	err = json.Unmarshal(data, &dst.PassportElementErrorReverseSide);
	if err == nil {
		jsonPassportElementErrorReverseSide, _ := json.Marshal(dst.PassportElementErrorReverseSide)
		if string(jsonPassportElementErrorReverseSide) == "{}" { // empty struct
			dst.PassportElementErrorReverseSide = nil
		} else {
			return nil // data stored in dst.PassportElementErrorReverseSide, return on the first match
		}
	} else {
		dst.PassportElementErrorReverseSide = nil
	}

	// try to unmarshal JSON data into PassportElementErrorSelfie
	err = json.Unmarshal(data, &dst.PassportElementErrorSelfie);
	if err == nil {
		jsonPassportElementErrorSelfie, _ := json.Marshal(dst.PassportElementErrorSelfie)
		if string(jsonPassportElementErrorSelfie) == "{}" { // empty struct
			dst.PassportElementErrorSelfie = nil
		} else {
			return nil // data stored in dst.PassportElementErrorSelfie, return on the first match
		}
	} else {
		dst.PassportElementErrorSelfie = nil
	}

	// try to unmarshal JSON data into PassportElementErrorTranslationFile
	err = json.Unmarshal(data, &dst.PassportElementErrorTranslationFile);
	if err == nil {
		jsonPassportElementErrorTranslationFile, _ := json.Marshal(dst.PassportElementErrorTranslationFile)
		if string(jsonPassportElementErrorTranslationFile) == "{}" { // empty struct
			dst.PassportElementErrorTranslationFile = nil
		} else {
			return nil // data stored in dst.PassportElementErrorTranslationFile, return on the first match
		}
	} else {
		dst.PassportElementErrorTranslationFile = nil
	}

	// try to unmarshal JSON data into PassportElementErrorTranslationFiles
	err = json.Unmarshal(data, &dst.PassportElementErrorTranslationFiles);
	if err == nil {
		jsonPassportElementErrorTranslationFiles, _ := json.Marshal(dst.PassportElementErrorTranslationFiles)
		if string(jsonPassportElementErrorTranslationFiles) == "{}" { // empty struct
			dst.PassportElementErrorTranslationFiles = nil
		} else {
			return nil // data stored in dst.PassportElementErrorTranslationFiles, return on the first match
		}
	} else {
		dst.PassportElementErrorTranslationFiles = nil
	}

	// try to unmarshal JSON data into PassportElementErrorUnspecified
	err = json.Unmarshal(data, &dst.PassportElementErrorUnspecified);
	if err == nil {
		jsonPassportElementErrorUnspecified, _ := json.Marshal(dst.PassportElementErrorUnspecified)
		if string(jsonPassportElementErrorUnspecified) == "{}" { // empty struct
			dst.PassportElementErrorUnspecified = nil
		} else {
			return nil // data stored in dst.PassportElementErrorUnspecified, return on the first match
		}
	} else {
		dst.PassportElementErrorUnspecified = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PassportElementError)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PassportElementError) MarshalJSON() ([]byte, error) {
	if src.PassportElementErrorDataField != nil {
		return json.Marshal(&src.PassportElementErrorDataField)
	}

	if src.PassportElementErrorFile != nil {
		return json.Marshal(&src.PassportElementErrorFile)
	}

	if src.PassportElementErrorFiles != nil {
		return json.Marshal(&src.PassportElementErrorFiles)
	}

	if src.PassportElementErrorFrontSide != nil {
		return json.Marshal(&src.PassportElementErrorFrontSide)
	}

	if src.PassportElementErrorReverseSide != nil {
		return json.Marshal(&src.PassportElementErrorReverseSide)
	}

	if src.PassportElementErrorSelfie != nil {
		return json.Marshal(&src.PassportElementErrorSelfie)
	}

	if src.PassportElementErrorTranslationFile != nil {
		return json.Marshal(&src.PassportElementErrorTranslationFile)
	}

	if src.PassportElementErrorTranslationFiles != nil {
		return json.Marshal(&src.PassportElementErrorTranslationFiles)
	}

	if src.PassportElementErrorUnspecified != nil {
		return json.Marshal(&src.PassportElementErrorUnspecified)
	}

	return nil, nil // no data in anyOf schemas
}


type NullablePassportElementError struct {
	value *PassportElementError
	isSet bool
}

func (v NullablePassportElementError) Get() *PassportElementError {
	return v.value
}

func (v *NullablePassportElementError) Set(val *PassportElementError) {
	v.value = val
	v.isSet = true
}

func (v NullablePassportElementError) IsSet() bool {
	return v.isSet
}

func (v *NullablePassportElementError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassportElementError(val *PassportElementError) *NullablePassportElementError {
	return &NullablePassportElementError{value: val, isSet: true}
}

func (v NullablePassportElementError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassportElementError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


