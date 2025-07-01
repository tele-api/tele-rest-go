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


// InlineQueryResult This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:  * [InlineQueryResultCachedAudio](https://core.telegram.org/bots/api/#inlinequeryresultcachedaudio) * [InlineQueryResultCachedDocument](https://core.telegram.org/bots/api/#inlinequeryresultcacheddocument) * [InlineQueryResultCachedGif](https://core.telegram.org/bots/api/#inlinequeryresultcachedgif) * [InlineQueryResultCachedMpeg4Gif](https://core.telegram.org/bots/api/#inlinequeryresultcachedmpeg4gif) * [InlineQueryResultCachedPhoto](https://core.telegram.org/bots/api/#inlinequeryresultcachedphoto) * [InlineQueryResultCachedSticker](https://core.telegram.org/bots/api/#inlinequeryresultcachedsticker) * [InlineQueryResultCachedVideo](https://core.telegram.org/bots/api/#inlinequeryresultcachedvideo) * [InlineQueryResultCachedVoice](https://core.telegram.org/bots/api/#inlinequeryresultcachedvoice) * [InlineQueryResultArticle](https://core.telegram.org/bots/api/#inlinequeryresultarticle) * [InlineQueryResultAudio](https://core.telegram.org/bots/api/#inlinequeryresultaudio) * [InlineQueryResultContact](https://core.telegram.org/bots/api/#inlinequeryresultcontact) * [InlineQueryResultGame](https://core.telegram.org/bots/api/#inlinequeryresultgame) * [InlineQueryResultDocument](https://core.telegram.org/bots/api/#inlinequeryresultdocument) * [InlineQueryResultGif](https://core.telegram.org/bots/api/#inlinequeryresultgif) * [InlineQueryResultLocation](https://core.telegram.org/bots/api/#inlinequeryresultlocation) * [InlineQueryResultMpeg4Gif](https://core.telegram.org/bots/api/#inlinequeryresultmpeg4gif) * [InlineQueryResultPhoto](https://core.telegram.org/bots/api/#inlinequeryresultphoto) * [InlineQueryResultVenue](https://core.telegram.org/bots/api/#inlinequeryresultvenue) * [InlineQueryResultVideo](https://core.telegram.org/bots/api/#inlinequeryresultvideo) * [InlineQueryResultVoice](https://core.telegram.org/bots/api/#inlinequeryresultvoice)
type InlineQueryResult struct {
	InlineQueryResultArticle *InlineQueryResultArticle
	InlineQueryResultAudio *InlineQueryResultAudio
	InlineQueryResultCachedAudio *InlineQueryResultCachedAudio
	InlineQueryResultCachedDocument *InlineQueryResultCachedDocument
	InlineQueryResultCachedGif *InlineQueryResultCachedGif
	InlineQueryResultCachedMpeg4Gif *InlineQueryResultCachedMpeg4Gif
	InlineQueryResultCachedPhoto *InlineQueryResultCachedPhoto
	InlineQueryResultCachedSticker *InlineQueryResultCachedSticker
	InlineQueryResultCachedVideo *InlineQueryResultCachedVideo
	InlineQueryResultCachedVoice *InlineQueryResultCachedVoice
	InlineQueryResultContact *InlineQueryResultContact
	InlineQueryResultDocument *InlineQueryResultDocument
	InlineQueryResultGame *InlineQueryResultGame
	InlineQueryResultGif *InlineQueryResultGif
	InlineQueryResultLocation *InlineQueryResultLocation
	InlineQueryResultMpeg4Gif *InlineQueryResultMpeg4Gif
	InlineQueryResultPhoto *InlineQueryResultPhoto
	InlineQueryResultVenue *InlineQueryResultVenue
	InlineQueryResultVideo *InlineQueryResultVideo
	InlineQueryResultVoice *InlineQueryResultVoice
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InlineQueryResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InlineQueryResultArticle
	err = json.Unmarshal(data, &dst.InlineQueryResultArticle);
	if err == nil {
		jsonInlineQueryResultArticle, _ := json.Marshal(dst.InlineQueryResultArticle)
		if string(jsonInlineQueryResultArticle) == "{}" { // empty struct
			dst.InlineQueryResultArticle = nil
		} else {
			return nil // data stored in dst.InlineQueryResultArticle, return on the first match
		}
	} else {
		dst.InlineQueryResultArticle = nil
	}

	// try to unmarshal JSON data into InlineQueryResultAudio
	err = json.Unmarshal(data, &dst.InlineQueryResultAudio);
	if err == nil {
		jsonInlineQueryResultAudio, _ := json.Marshal(dst.InlineQueryResultAudio)
		if string(jsonInlineQueryResultAudio) == "{}" { // empty struct
			dst.InlineQueryResultAudio = nil
		} else {
			return nil // data stored in dst.InlineQueryResultAudio, return on the first match
		}
	} else {
		dst.InlineQueryResultAudio = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedAudio
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedAudio);
	if err == nil {
		jsonInlineQueryResultCachedAudio, _ := json.Marshal(dst.InlineQueryResultCachedAudio)
		if string(jsonInlineQueryResultCachedAudio) == "{}" { // empty struct
			dst.InlineQueryResultCachedAudio = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedAudio, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedAudio = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedDocument
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedDocument);
	if err == nil {
		jsonInlineQueryResultCachedDocument, _ := json.Marshal(dst.InlineQueryResultCachedDocument)
		if string(jsonInlineQueryResultCachedDocument) == "{}" { // empty struct
			dst.InlineQueryResultCachedDocument = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedDocument, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedDocument = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedGif
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedGif);
	if err == nil {
		jsonInlineQueryResultCachedGif, _ := json.Marshal(dst.InlineQueryResultCachedGif)
		if string(jsonInlineQueryResultCachedGif) == "{}" { // empty struct
			dst.InlineQueryResultCachedGif = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedGif, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedGif = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedMpeg4Gif
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedMpeg4Gif);
	if err == nil {
		jsonInlineQueryResultCachedMpeg4Gif, _ := json.Marshal(dst.InlineQueryResultCachedMpeg4Gif)
		if string(jsonInlineQueryResultCachedMpeg4Gif) == "{}" { // empty struct
			dst.InlineQueryResultCachedMpeg4Gif = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedMpeg4Gif, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedMpeg4Gif = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedPhoto
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedPhoto);
	if err == nil {
		jsonInlineQueryResultCachedPhoto, _ := json.Marshal(dst.InlineQueryResultCachedPhoto)
		if string(jsonInlineQueryResultCachedPhoto) == "{}" { // empty struct
			dst.InlineQueryResultCachedPhoto = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedPhoto, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedPhoto = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedSticker
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedSticker);
	if err == nil {
		jsonInlineQueryResultCachedSticker, _ := json.Marshal(dst.InlineQueryResultCachedSticker)
		if string(jsonInlineQueryResultCachedSticker) == "{}" { // empty struct
			dst.InlineQueryResultCachedSticker = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedSticker, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedSticker = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedVideo
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedVideo);
	if err == nil {
		jsonInlineQueryResultCachedVideo, _ := json.Marshal(dst.InlineQueryResultCachedVideo)
		if string(jsonInlineQueryResultCachedVideo) == "{}" { // empty struct
			dst.InlineQueryResultCachedVideo = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedVideo, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedVideo = nil
	}

	// try to unmarshal JSON data into InlineQueryResultCachedVoice
	err = json.Unmarshal(data, &dst.InlineQueryResultCachedVoice);
	if err == nil {
		jsonInlineQueryResultCachedVoice, _ := json.Marshal(dst.InlineQueryResultCachedVoice)
		if string(jsonInlineQueryResultCachedVoice) == "{}" { // empty struct
			dst.InlineQueryResultCachedVoice = nil
		} else {
			return nil // data stored in dst.InlineQueryResultCachedVoice, return on the first match
		}
	} else {
		dst.InlineQueryResultCachedVoice = nil
	}

	// try to unmarshal JSON data into InlineQueryResultContact
	err = json.Unmarshal(data, &dst.InlineQueryResultContact);
	if err == nil {
		jsonInlineQueryResultContact, _ := json.Marshal(dst.InlineQueryResultContact)
		if string(jsonInlineQueryResultContact) == "{}" { // empty struct
			dst.InlineQueryResultContact = nil
		} else {
			return nil // data stored in dst.InlineQueryResultContact, return on the first match
		}
	} else {
		dst.InlineQueryResultContact = nil
	}

	// try to unmarshal JSON data into InlineQueryResultDocument
	err = json.Unmarshal(data, &dst.InlineQueryResultDocument);
	if err == nil {
		jsonInlineQueryResultDocument, _ := json.Marshal(dst.InlineQueryResultDocument)
		if string(jsonInlineQueryResultDocument) == "{}" { // empty struct
			dst.InlineQueryResultDocument = nil
		} else {
			return nil // data stored in dst.InlineQueryResultDocument, return on the first match
		}
	} else {
		dst.InlineQueryResultDocument = nil
	}

	// try to unmarshal JSON data into InlineQueryResultGame
	err = json.Unmarshal(data, &dst.InlineQueryResultGame);
	if err == nil {
		jsonInlineQueryResultGame, _ := json.Marshal(dst.InlineQueryResultGame)
		if string(jsonInlineQueryResultGame) == "{}" { // empty struct
			dst.InlineQueryResultGame = nil
		} else {
			return nil // data stored in dst.InlineQueryResultGame, return on the first match
		}
	} else {
		dst.InlineQueryResultGame = nil
	}

	// try to unmarshal JSON data into InlineQueryResultGif
	err = json.Unmarshal(data, &dst.InlineQueryResultGif);
	if err == nil {
		jsonInlineQueryResultGif, _ := json.Marshal(dst.InlineQueryResultGif)
		if string(jsonInlineQueryResultGif) == "{}" { // empty struct
			dst.InlineQueryResultGif = nil
		} else {
			return nil // data stored in dst.InlineQueryResultGif, return on the first match
		}
	} else {
		dst.InlineQueryResultGif = nil
	}

	// try to unmarshal JSON data into InlineQueryResultLocation
	err = json.Unmarshal(data, &dst.InlineQueryResultLocation);
	if err == nil {
		jsonInlineQueryResultLocation, _ := json.Marshal(dst.InlineQueryResultLocation)
		if string(jsonInlineQueryResultLocation) == "{}" { // empty struct
			dst.InlineQueryResultLocation = nil
		} else {
			return nil // data stored in dst.InlineQueryResultLocation, return on the first match
		}
	} else {
		dst.InlineQueryResultLocation = nil
	}

	// try to unmarshal JSON data into InlineQueryResultMpeg4Gif
	err = json.Unmarshal(data, &dst.InlineQueryResultMpeg4Gif);
	if err == nil {
		jsonInlineQueryResultMpeg4Gif, _ := json.Marshal(dst.InlineQueryResultMpeg4Gif)
		if string(jsonInlineQueryResultMpeg4Gif) == "{}" { // empty struct
			dst.InlineQueryResultMpeg4Gif = nil
		} else {
			return nil // data stored in dst.InlineQueryResultMpeg4Gif, return on the first match
		}
	} else {
		dst.InlineQueryResultMpeg4Gif = nil
	}

	// try to unmarshal JSON data into InlineQueryResultPhoto
	err = json.Unmarshal(data, &dst.InlineQueryResultPhoto);
	if err == nil {
		jsonInlineQueryResultPhoto, _ := json.Marshal(dst.InlineQueryResultPhoto)
		if string(jsonInlineQueryResultPhoto) == "{}" { // empty struct
			dst.InlineQueryResultPhoto = nil
		} else {
			return nil // data stored in dst.InlineQueryResultPhoto, return on the first match
		}
	} else {
		dst.InlineQueryResultPhoto = nil
	}

	// try to unmarshal JSON data into InlineQueryResultVenue
	err = json.Unmarshal(data, &dst.InlineQueryResultVenue);
	if err == nil {
		jsonInlineQueryResultVenue, _ := json.Marshal(dst.InlineQueryResultVenue)
		if string(jsonInlineQueryResultVenue) == "{}" { // empty struct
			dst.InlineQueryResultVenue = nil
		} else {
			return nil // data stored in dst.InlineQueryResultVenue, return on the first match
		}
	} else {
		dst.InlineQueryResultVenue = nil
	}

	// try to unmarshal JSON data into InlineQueryResultVideo
	err = json.Unmarshal(data, &dst.InlineQueryResultVideo);
	if err == nil {
		jsonInlineQueryResultVideo, _ := json.Marshal(dst.InlineQueryResultVideo)
		if string(jsonInlineQueryResultVideo) == "{}" { // empty struct
			dst.InlineQueryResultVideo = nil
		} else {
			return nil // data stored in dst.InlineQueryResultVideo, return on the first match
		}
	} else {
		dst.InlineQueryResultVideo = nil
	}

	// try to unmarshal JSON data into InlineQueryResultVoice
	err = json.Unmarshal(data, &dst.InlineQueryResultVoice);
	if err == nil {
		jsonInlineQueryResultVoice, _ := json.Marshal(dst.InlineQueryResultVoice)
		if string(jsonInlineQueryResultVoice) == "{}" { // empty struct
			dst.InlineQueryResultVoice = nil
		} else {
			return nil // data stored in dst.InlineQueryResultVoice, return on the first match
		}
	} else {
		dst.InlineQueryResultVoice = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(InlineQueryResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InlineQueryResult) MarshalJSON() ([]byte, error) {
	if src.InlineQueryResultArticle != nil {
		return json.Marshal(&src.InlineQueryResultArticle)
	}

	if src.InlineQueryResultAudio != nil {
		return json.Marshal(&src.InlineQueryResultAudio)
	}

	if src.InlineQueryResultCachedAudio != nil {
		return json.Marshal(&src.InlineQueryResultCachedAudio)
	}

	if src.InlineQueryResultCachedDocument != nil {
		return json.Marshal(&src.InlineQueryResultCachedDocument)
	}

	if src.InlineQueryResultCachedGif != nil {
		return json.Marshal(&src.InlineQueryResultCachedGif)
	}

	if src.InlineQueryResultCachedMpeg4Gif != nil {
		return json.Marshal(&src.InlineQueryResultCachedMpeg4Gif)
	}

	if src.InlineQueryResultCachedPhoto != nil {
		return json.Marshal(&src.InlineQueryResultCachedPhoto)
	}

	if src.InlineQueryResultCachedSticker != nil {
		return json.Marshal(&src.InlineQueryResultCachedSticker)
	}

	if src.InlineQueryResultCachedVideo != nil {
		return json.Marshal(&src.InlineQueryResultCachedVideo)
	}

	if src.InlineQueryResultCachedVoice != nil {
		return json.Marshal(&src.InlineQueryResultCachedVoice)
	}

	if src.InlineQueryResultContact != nil {
		return json.Marshal(&src.InlineQueryResultContact)
	}

	if src.InlineQueryResultDocument != nil {
		return json.Marshal(&src.InlineQueryResultDocument)
	}

	if src.InlineQueryResultGame != nil {
		return json.Marshal(&src.InlineQueryResultGame)
	}

	if src.InlineQueryResultGif != nil {
		return json.Marshal(&src.InlineQueryResultGif)
	}

	if src.InlineQueryResultLocation != nil {
		return json.Marshal(&src.InlineQueryResultLocation)
	}

	if src.InlineQueryResultMpeg4Gif != nil {
		return json.Marshal(&src.InlineQueryResultMpeg4Gif)
	}

	if src.InlineQueryResultPhoto != nil {
		return json.Marshal(&src.InlineQueryResultPhoto)
	}

	if src.InlineQueryResultVenue != nil {
		return json.Marshal(&src.InlineQueryResultVenue)
	}

	if src.InlineQueryResultVideo != nil {
		return json.Marshal(&src.InlineQueryResultVideo)
	}

	if src.InlineQueryResultVoice != nil {
		return json.Marshal(&src.InlineQueryResultVoice)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableInlineQueryResult struct {
	value *InlineQueryResult
	isSet bool
}

func (v NullableInlineQueryResult) Get() *InlineQueryResult {
	return v.value
}

func (v *NullableInlineQueryResult) Set(val *InlineQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResult(val *InlineQueryResult) *NullableInlineQueryResult {
	return &NullableInlineQueryResult{value: val, isSet: true}
}

func (v NullableInlineQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


