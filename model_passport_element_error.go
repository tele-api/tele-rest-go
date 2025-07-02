/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// PassportElementError - This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:  * [PassportElementErrorDataField](https://core.telegram.org/bots/api/#passportelementerrordatafield) * [PassportElementErrorFrontSide](https://core.telegram.org/bots/api/#passportelementerrorfrontside) * [PassportElementErrorReverseSide](https://core.telegram.org/bots/api/#passportelementerrorreverseside) * [PassportElementErrorSelfie](https://core.telegram.org/bots/api/#passportelementerrorselfie) * [PassportElementErrorFile](https://core.telegram.org/bots/api/#passportelementerrorfile) * [PassportElementErrorFiles](https://core.telegram.org/bots/api/#passportelementerrorfiles) * [PassportElementErrorTranslationFile](https://core.telegram.org/bots/api/#passportelementerrortranslationfile) * [PassportElementErrorTranslationFiles](https://core.telegram.org/bots/api/#passportelementerrortranslationfiles) * [PassportElementErrorUnspecified](https://core.telegram.org/bots/api/#passportelementerrorunspecified)
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

// PassportElementErrorDataFieldAsPassportElementError is a convenience function that returns PassportElementErrorDataField wrapped in PassportElementError
func PassportElementErrorDataFieldAsPassportElementError(v *PassportElementErrorDataField) PassportElementError {
	return PassportElementError{
		PassportElementErrorDataField: v,
	}
}

// PassportElementErrorFileAsPassportElementError is a convenience function that returns PassportElementErrorFile wrapped in PassportElementError
func PassportElementErrorFileAsPassportElementError(v *PassportElementErrorFile) PassportElementError {
	return PassportElementError{
		PassportElementErrorFile: v,
	}
}

// PassportElementErrorFilesAsPassportElementError is a convenience function that returns PassportElementErrorFiles wrapped in PassportElementError
func PassportElementErrorFilesAsPassportElementError(v *PassportElementErrorFiles) PassportElementError {
	return PassportElementError{
		PassportElementErrorFiles: v,
	}
}

// PassportElementErrorFrontSideAsPassportElementError is a convenience function that returns PassportElementErrorFrontSide wrapped in PassportElementError
func PassportElementErrorFrontSideAsPassportElementError(v *PassportElementErrorFrontSide) PassportElementError {
	return PassportElementError{
		PassportElementErrorFrontSide: v,
	}
}

// PassportElementErrorReverseSideAsPassportElementError is a convenience function that returns PassportElementErrorReverseSide wrapped in PassportElementError
func PassportElementErrorReverseSideAsPassportElementError(v *PassportElementErrorReverseSide) PassportElementError {
	return PassportElementError{
		PassportElementErrorReverseSide: v,
	}
}

// PassportElementErrorSelfieAsPassportElementError is a convenience function that returns PassportElementErrorSelfie wrapped in PassportElementError
func PassportElementErrorSelfieAsPassportElementError(v *PassportElementErrorSelfie) PassportElementError {
	return PassportElementError{
		PassportElementErrorSelfie: v,
	}
}

// PassportElementErrorTranslationFileAsPassportElementError is a convenience function that returns PassportElementErrorTranslationFile wrapped in PassportElementError
func PassportElementErrorTranslationFileAsPassportElementError(v *PassportElementErrorTranslationFile) PassportElementError {
	return PassportElementError{
		PassportElementErrorTranslationFile: v,
	}
}

// PassportElementErrorTranslationFilesAsPassportElementError is a convenience function that returns PassportElementErrorTranslationFiles wrapped in PassportElementError
func PassportElementErrorTranslationFilesAsPassportElementError(v *PassportElementErrorTranslationFiles) PassportElementError {
	return PassportElementError{
		PassportElementErrorTranslationFiles: v,
	}
}

// PassportElementErrorUnspecifiedAsPassportElementError is a convenience function that returns PassportElementErrorUnspecified wrapped in PassportElementError
func PassportElementErrorUnspecifiedAsPassportElementError(v *PassportElementErrorUnspecified) PassportElementError {
	return PassportElementError{
		PassportElementErrorUnspecified: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PassportElementError) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PassportElementErrorDataField
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorDataField)
	if err == nil {
		jsonPassportElementErrorDataField, _ := json.Marshal(dst.PassportElementErrorDataField)
		if string(jsonPassportElementErrorDataField) == "{}" { // empty struct
			dst.PassportElementErrorDataField = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorDataField); err != nil {
				dst.PassportElementErrorDataField = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorDataField = nil
	}

	// try to unmarshal data into PassportElementErrorFile
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorFile)
	if err == nil {
		jsonPassportElementErrorFile, _ := json.Marshal(dst.PassportElementErrorFile)
		if string(jsonPassportElementErrorFile) == "{}" { // empty struct
			dst.PassportElementErrorFile = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorFile); err != nil {
				dst.PassportElementErrorFile = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorFile = nil
	}

	// try to unmarshal data into PassportElementErrorFiles
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorFiles)
	if err == nil {
		jsonPassportElementErrorFiles, _ := json.Marshal(dst.PassportElementErrorFiles)
		if string(jsonPassportElementErrorFiles) == "{}" { // empty struct
			dst.PassportElementErrorFiles = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorFiles); err != nil {
				dst.PassportElementErrorFiles = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorFiles = nil
	}

	// try to unmarshal data into PassportElementErrorFrontSide
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorFrontSide)
	if err == nil {
		jsonPassportElementErrorFrontSide, _ := json.Marshal(dst.PassportElementErrorFrontSide)
		if string(jsonPassportElementErrorFrontSide) == "{}" { // empty struct
			dst.PassportElementErrorFrontSide = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorFrontSide); err != nil {
				dst.PassportElementErrorFrontSide = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorFrontSide = nil
	}

	// try to unmarshal data into PassportElementErrorReverseSide
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorReverseSide)
	if err == nil {
		jsonPassportElementErrorReverseSide, _ := json.Marshal(dst.PassportElementErrorReverseSide)
		if string(jsonPassportElementErrorReverseSide) == "{}" { // empty struct
			dst.PassportElementErrorReverseSide = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorReverseSide); err != nil {
				dst.PassportElementErrorReverseSide = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorReverseSide = nil
	}

	// try to unmarshal data into PassportElementErrorSelfie
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorSelfie)
	if err == nil {
		jsonPassportElementErrorSelfie, _ := json.Marshal(dst.PassportElementErrorSelfie)
		if string(jsonPassportElementErrorSelfie) == "{}" { // empty struct
			dst.PassportElementErrorSelfie = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorSelfie); err != nil {
				dst.PassportElementErrorSelfie = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorSelfie = nil
	}

	// try to unmarshal data into PassportElementErrorTranslationFile
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorTranslationFile)
	if err == nil {
		jsonPassportElementErrorTranslationFile, _ := json.Marshal(dst.PassportElementErrorTranslationFile)
		if string(jsonPassportElementErrorTranslationFile) == "{}" { // empty struct
			dst.PassportElementErrorTranslationFile = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorTranslationFile); err != nil {
				dst.PassportElementErrorTranslationFile = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorTranslationFile = nil
	}

	// try to unmarshal data into PassportElementErrorTranslationFiles
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorTranslationFiles)
	if err == nil {
		jsonPassportElementErrorTranslationFiles, _ := json.Marshal(dst.PassportElementErrorTranslationFiles)
		if string(jsonPassportElementErrorTranslationFiles) == "{}" { // empty struct
			dst.PassportElementErrorTranslationFiles = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorTranslationFiles); err != nil {
				dst.PassportElementErrorTranslationFiles = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorTranslationFiles = nil
	}

	// try to unmarshal data into PassportElementErrorUnspecified
	err = newStrictDecoder(data).Decode(&dst.PassportElementErrorUnspecified)
	if err == nil {
		jsonPassportElementErrorUnspecified, _ := json.Marshal(dst.PassportElementErrorUnspecified)
		if string(jsonPassportElementErrorUnspecified) == "{}" { // empty struct
			dst.PassportElementErrorUnspecified = nil
		} else {
			if err = validator.Validate(dst.PassportElementErrorUnspecified); err != nil {
				dst.PassportElementErrorUnspecified = nil
			} else {
				match++
			}
		}
	} else {
		dst.PassportElementErrorUnspecified = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PassportElementErrorDataField = nil
		dst.PassportElementErrorFile = nil
		dst.PassportElementErrorFiles = nil
		dst.PassportElementErrorFrontSide = nil
		dst.PassportElementErrorReverseSide = nil
		dst.PassportElementErrorSelfie = nil
		dst.PassportElementErrorTranslationFile = nil
		dst.PassportElementErrorTranslationFiles = nil
		dst.PassportElementErrorUnspecified = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PassportElementError)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PassportElementError)")
	}
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

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PassportElementError) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PassportElementErrorDataField != nil {
		return obj.PassportElementErrorDataField
	}

	if obj.PassportElementErrorFile != nil {
		return obj.PassportElementErrorFile
	}

	if obj.PassportElementErrorFiles != nil {
		return obj.PassportElementErrorFiles
	}

	if obj.PassportElementErrorFrontSide != nil {
		return obj.PassportElementErrorFrontSide
	}

	if obj.PassportElementErrorReverseSide != nil {
		return obj.PassportElementErrorReverseSide
	}

	if obj.PassportElementErrorSelfie != nil {
		return obj.PassportElementErrorSelfie
	}

	if obj.PassportElementErrorTranslationFile != nil {
		return obj.PassportElementErrorTranslationFile
	}

	if obj.PassportElementErrorTranslationFiles != nil {
		return obj.PassportElementErrorTranslationFiles
	}

	if obj.PassportElementErrorUnspecified != nil {
		return obj.PassportElementErrorUnspecified
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj PassportElementError) GetActualInstanceValue() (interface{}) {
	if obj.PassportElementErrorDataField != nil {
		return *obj.PassportElementErrorDataField
	}

	if obj.PassportElementErrorFile != nil {
		return *obj.PassportElementErrorFile
	}

	if obj.PassportElementErrorFiles != nil {
		return *obj.PassportElementErrorFiles
	}

	if obj.PassportElementErrorFrontSide != nil {
		return *obj.PassportElementErrorFrontSide
	}

	if obj.PassportElementErrorReverseSide != nil {
		return *obj.PassportElementErrorReverseSide
	}

	if obj.PassportElementErrorSelfie != nil {
		return *obj.PassportElementErrorSelfie
	}

	if obj.PassportElementErrorTranslationFile != nil {
		return *obj.PassportElementErrorTranslationFile
	}

	if obj.PassportElementErrorTranslationFiles != nil {
		return *obj.PassportElementErrorTranslationFiles
	}

	if obj.PassportElementErrorUnspecified != nil {
		return *obj.PassportElementErrorUnspecified
	}

	// all schemas are nil
	return nil
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


