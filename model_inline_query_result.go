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

// InlineQueryResult - This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:  * [InlineQueryResultCachedAudio](https://core.telegram.org/bots/api/#inlinequeryresultcachedaudio) * [InlineQueryResultCachedDocument](https://core.telegram.org/bots/api/#inlinequeryresultcacheddocument) * [InlineQueryResultCachedGif](https://core.telegram.org/bots/api/#inlinequeryresultcachedgif) * [InlineQueryResultCachedMpeg4Gif](https://core.telegram.org/bots/api/#inlinequeryresultcachedmpeg4gif) * [InlineQueryResultCachedPhoto](https://core.telegram.org/bots/api/#inlinequeryresultcachedphoto) * [InlineQueryResultCachedSticker](https://core.telegram.org/bots/api/#inlinequeryresultcachedsticker) * [InlineQueryResultCachedVideo](https://core.telegram.org/bots/api/#inlinequeryresultcachedvideo) * [InlineQueryResultCachedVoice](https://core.telegram.org/bots/api/#inlinequeryresultcachedvoice) * [InlineQueryResultArticle](https://core.telegram.org/bots/api/#inlinequeryresultarticle) * [InlineQueryResultAudio](https://core.telegram.org/bots/api/#inlinequeryresultaudio) * [InlineQueryResultContact](https://core.telegram.org/bots/api/#inlinequeryresultcontact) * [InlineQueryResultGame](https://core.telegram.org/bots/api/#inlinequeryresultgame) * [InlineQueryResultDocument](https://core.telegram.org/bots/api/#inlinequeryresultdocument) * [InlineQueryResultGif](https://core.telegram.org/bots/api/#inlinequeryresultgif) * [InlineQueryResultLocation](https://core.telegram.org/bots/api/#inlinequeryresultlocation) * [InlineQueryResultMpeg4Gif](https://core.telegram.org/bots/api/#inlinequeryresultmpeg4gif) * [InlineQueryResultPhoto](https://core.telegram.org/bots/api/#inlinequeryresultphoto) * [InlineQueryResultVenue](https://core.telegram.org/bots/api/#inlinequeryresultvenue) * [InlineQueryResultVideo](https://core.telegram.org/bots/api/#inlinequeryresultvideo) * [InlineQueryResultVoice](https://core.telegram.org/bots/api/#inlinequeryresultvoice)
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

// InlineQueryResultArticleAsInlineQueryResult is a convenience function that returns InlineQueryResultArticle wrapped in InlineQueryResult
func InlineQueryResultArticleAsInlineQueryResult(v *InlineQueryResultArticle) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultArticle: v,
	}
}

// InlineQueryResultAudioAsInlineQueryResult is a convenience function that returns InlineQueryResultAudio wrapped in InlineQueryResult
func InlineQueryResultAudioAsInlineQueryResult(v *InlineQueryResultAudio) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultAudio: v,
	}
}

// InlineQueryResultCachedAudioAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedAudio wrapped in InlineQueryResult
func InlineQueryResultCachedAudioAsInlineQueryResult(v *InlineQueryResultCachedAudio) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedAudio: v,
	}
}

// InlineQueryResultCachedDocumentAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedDocument wrapped in InlineQueryResult
func InlineQueryResultCachedDocumentAsInlineQueryResult(v *InlineQueryResultCachedDocument) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedDocument: v,
	}
}

// InlineQueryResultCachedGifAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedGif wrapped in InlineQueryResult
func InlineQueryResultCachedGifAsInlineQueryResult(v *InlineQueryResultCachedGif) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedGif: v,
	}
}

// InlineQueryResultCachedMpeg4GifAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedMpeg4Gif wrapped in InlineQueryResult
func InlineQueryResultCachedMpeg4GifAsInlineQueryResult(v *InlineQueryResultCachedMpeg4Gif) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedMpeg4Gif: v,
	}
}

// InlineQueryResultCachedPhotoAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedPhoto wrapped in InlineQueryResult
func InlineQueryResultCachedPhotoAsInlineQueryResult(v *InlineQueryResultCachedPhoto) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedPhoto: v,
	}
}

// InlineQueryResultCachedStickerAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedSticker wrapped in InlineQueryResult
func InlineQueryResultCachedStickerAsInlineQueryResult(v *InlineQueryResultCachedSticker) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedSticker: v,
	}
}

// InlineQueryResultCachedVideoAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedVideo wrapped in InlineQueryResult
func InlineQueryResultCachedVideoAsInlineQueryResult(v *InlineQueryResultCachedVideo) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedVideo: v,
	}
}

// InlineQueryResultCachedVoiceAsInlineQueryResult is a convenience function that returns InlineQueryResultCachedVoice wrapped in InlineQueryResult
func InlineQueryResultCachedVoiceAsInlineQueryResult(v *InlineQueryResultCachedVoice) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultCachedVoice: v,
	}
}

// InlineQueryResultContactAsInlineQueryResult is a convenience function that returns InlineQueryResultContact wrapped in InlineQueryResult
func InlineQueryResultContactAsInlineQueryResult(v *InlineQueryResultContact) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultContact: v,
	}
}

// InlineQueryResultDocumentAsInlineQueryResult is a convenience function that returns InlineQueryResultDocument wrapped in InlineQueryResult
func InlineQueryResultDocumentAsInlineQueryResult(v *InlineQueryResultDocument) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultDocument: v,
	}
}

// InlineQueryResultGameAsInlineQueryResult is a convenience function that returns InlineQueryResultGame wrapped in InlineQueryResult
func InlineQueryResultGameAsInlineQueryResult(v *InlineQueryResultGame) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultGame: v,
	}
}

// InlineQueryResultGifAsInlineQueryResult is a convenience function that returns InlineQueryResultGif wrapped in InlineQueryResult
func InlineQueryResultGifAsInlineQueryResult(v *InlineQueryResultGif) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultGif: v,
	}
}

// InlineQueryResultLocationAsInlineQueryResult is a convenience function that returns InlineQueryResultLocation wrapped in InlineQueryResult
func InlineQueryResultLocationAsInlineQueryResult(v *InlineQueryResultLocation) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultLocation: v,
	}
}

// InlineQueryResultMpeg4GifAsInlineQueryResult is a convenience function that returns InlineQueryResultMpeg4Gif wrapped in InlineQueryResult
func InlineQueryResultMpeg4GifAsInlineQueryResult(v *InlineQueryResultMpeg4Gif) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultMpeg4Gif: v,
	}
}

// InlineQueryResultPhotoAsInlineQueryResult is a convenience function that returns InlineQueryResultPhoto wrapped in InlineQueryResult
func InlineQueryResultPhotoAsInlineQueryResult(v *InlineQueryResultPhoto) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultPhoto: v,
	}
}

// InlineQueryResultVenueAsInlineQueryResult is a convenience function that returns InlineQueryResultVenue wrapped in InlineQueryResult
func InlineQueryResultVenueAsInlineQueryResult(v *InlineQueryResultVenue) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultVenue: v,
	}
}

// InlineQueryResultVideoAsInlineQueryResult is a convenience function that returns InlineQueryResultVideo wrapped in InlineQueryResult
func InlineQueryResultVideoAsInlineQueryResult(v *InlineQueryResultVideo) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultVideo: v,
	}
}

// InlineQueryResultVoiceAsInlineQueryResult is a convenience function that returns InlineQueryResultVoice wrapped in InlineQueryResult
func InlineQueryResultVoiceAsInlineQueryResult(v *InlineQueryResultVoice) InlineQueryResult {
	return InlineQueryResult{
		InlineQueryResultVoice: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InlineQueryResult) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InlineQueryResultArticle
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultArticle)
	if err == nil {
		jsonInlineQueryResultArticle, _ := json.Marshal(dst.InlineQueryResultArticle)
		if string(jsonInlineQueryResultArticle) == "{}" { // empty struct
			dst.InlineQueryResultArticle = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultArticle); err != nil {
				dst.InlineQueryResultArticle = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultArticle = nil
	}

	// try to unmarshal data into InlineQueryResultAudio
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultAudio)
	if err == nil {
		jsonInlineQueryResultAudio, _ := json.Marshal(dst.InlineQueryResultAudio)
		if string(jsonInlineQueryResultAudio) == "{}" { // empty struct
			dst.InlineQueryResultAudio = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultAudio); err != nil {
				dst.InlineQueryResultAudio = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultAudio = nil
	}

	// try to unmarshal data into InlineQueryResultCachedAudio
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedAudio)
	if err == nil {
		jsonInlineQueryResultCachedAudio, _ := json.Marshal(dst.InlineQueryResultCachedAudio)
		if string(jsonInlineQueryResultCachedAudio) == "{}" { // empty struct
			dst.InlineQueryResultCachedAudio = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedAudio); err != nil {
				dst.InlineQueryResultCachedAudio = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedAudio = nil
	}

	// try to unmarshal data into InlineQueryResultCachedDocument
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedDocument)
	if err == nil {
		jsonInlineQueryResultCachedDocument, _ := json.Marshal(dst.InlineQueryResultCachedDocument)
		if string(jsonInlineQueryResultCachedDocument) == "{}" { // empty struct
			dst.InlineQueryResultCachedDocument = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedDocument); err != nil {
				dst.InlineQueryResultCachedDocument = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedDocument = nil
	}

	// try to unmarshal data into InlineQueryResultCachedGif
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedGif)
	if err == nil {
		jsonInlineQueryResultCachedGif, _ := json.Marshal(dst.InlineQueryResultCachedGif)
		if string(jsonInlineQueryResultCachedGif) == "{}" { // empty struct
			dst.InlineQueryResultCachedGif = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedGif); err != nil {
				dst.InlineQueryResultCachedGif = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedGif = nil
	}

	// try to unmarshal data into InlineQueryResultCachedMpeg4Gif
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedMpeg4Gif)
	if err == nil {
		jsonInlineQueryResultCachedMpeg4Gif, _ := json.Marshal(dst.InlineQueryResultCachedMpeg4Gif)
		if string(jsonInlineQueryResultCachedMpeg4Gif) == "{}" { // empty struct
			dst.InlineQueryResultCachedMpeg4Gif = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedMpeg4Gif); err != nil {
				dst.InlineQueryResultCachedMpeg4Gif = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedMpeg4Gif = nil
	}

	// try to unmarshal data into InlineQueryResultCachedPhoto
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedPhoto)
	if err == nil {
		jsonInlineQueryResultCachedPhoto, _ := json.Marshal(dst.InlineQueryResultCachedPhoto)
		if string(jsonInlineQueryResultCachedPhoto) == "{}" { // empty struct
			dst.InlineQueryResultCachedPhoto = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedPhoto); err != nil {
				dst.InlineQueryResultCachedPhoto = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedPhoto = nil
	}

	// try to unmarshal data into InlineQueryResultCachedSticker
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedSticker)
	if err == nil {
		jsonInlineQueryResultCachedSticker, _ := json.Marshal(dst.InlineQueryResultCachedSticker)
		if string(jsonInlineQueryResultCachedSticker) == "{}" { // empty struct
			dst.InlineQueryResultCachedSticker = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedSticker); err != nil {
				dst.InlineQueryResultCachedSticker = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedSticker = nil
	}

	// try to unmarshal data into InlineQueryResultCachedVideo
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedVideo)
	if err == nil {
		jsonInlineQueryResultCachedVideo, _ := json.Marshal(dst.InlineQueryResultCachedVideo)
		if string(jsonInlineQueryResultCachedVideo) == "{}" { // empty struct
			dst.InlineQueryResultCachedVideo = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedVideo); err != nil {
				dst.InlineQueryResultCachedVideo = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedVideo = nil
	}

	// try to unmarshal data into InlineQueryResultCachedVoice
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultCachedVoice)
	if err == nil {
		jsonInlineQueryResultCachedVoice, _ := json.Marshal(dst.InlineQueryResultCachedVoice)
		if string(jsonInlineQueryResultCachedVoice) == "{}" { // empty struct
			dst.InlineQueryResultCachedVoice = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultCachedVoice); err != nil {
				dst.InlineQueryResultCachedVoice = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultCachedVoice = nil
	}

	// try to unmarshal data into InlineQueryResultContact
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultContact)
	if err == nil {
		jsonInlineQueryResultContact, _ := json.Marshal(dst.InlineQueryResultContact)
		if string(jsonInlineQueryResultContact) == "{}" { // empty struct
			dst.InlineQueryResultContact = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultContact); err != nil {
				dst.InlineQueryResultContact = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultContact = nil
	}

	// try to unmarshal data into InlineQueryResultDocument
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultDocument)
	if err == nil {
		jsonInlineQueryResultDocument, _ := json.Marshal(dst.InlineQueryResultDocument)
		if string(jsonInlineQueryResultDocument) == "{}" { // empty struct
			dst.InlineQueryResultDocument = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultDocument); err != nil {
				dst.InlineQueryResultDocument = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultDocument = nil
	}

	// try to unmarshal data into InlineQueryResultGame
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultGame)
	if err == nil {
		jsonInlineQueryResultGame, _ := json.Marshal(dst.InlineQueryResultGame)
		if string(jsonInlineQueryResultGame) == "{}" { // empty struct
			dst.InlineQueryResultGame = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultGame); err != nil {
				dst.InlineQueryResultGame = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultGame = nil
	}

	// try to unmarshal data into InlineQueryResultGif
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultGif)
	if err == nil {
		jsonInlineQueryResultGif, _ := json.Marshal(dst.InlineQueryResultGif)
		if string(jsonInlineQueryResultGif) == "{}" { // empty struct
			dst.InlineQueryResultGif = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultGif); err != nil {
				dst.InlineQueryResultGif = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultGif = nil
	}

	// try to unmarshal data into InlineQueryResultLocation
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultLocation)
	if err == nil {
		jsonInlineQueryResultLocation, _ := json.Marshal(dst.InlineQueryResultLocation)
		if string(jsonInlineQueryResultLocation) == "{}" { // empty struct
			dst.InlineQueryResultLocation = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultLocation); err != nil {
				dst.InlineQueryResultLocation = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultLocation = nil
	}

	// try to unmarshal data into InlineQueryResultMpeg4Gif
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultMpeg4Gif)
	if err == nil {
		jsonInlineQueryResultMpeg4Gif, _ := json.Marshal(dst.InlineQueryResultMpeg4Gif)
		if string(jsonInlineQueryResultMpeg4Gif) == "{}" { // empty struct
			dst.InlineQueryResultMpeg4Gif = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultMpeg4Gif); err != nil {
				dst.InlineQueryResultMpeg4Gif = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultMpeg4Gif = nil
	}

	// try to unmarshal data into InlineQueryResultPhoto
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultPhoto)
	if err == nil {
		jsonInlineQueryResultPhoto, _ := json.Marshal(dst.InlineQueryResultPhoto)
		if string(jsonInlineQueryResultPhoto) == "{}" { // empty struct
			dst.InlineQueryResultPhoto = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultPhoto); err != nil {
				dst.InlineQueryResultPhoto = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultPhoto = nil
	}

	// try to unmarshal data into InlineQueryResultVenue
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultVenue)
	if err == nil {
		jsonInlineQueryResultVenue, _ := json.Marshal(dst.InlineQueryResultVenue)
		if string(jsonInlineQueryResultVenue) == "{}" { // empty struct
			dst.InlineQueryResultVenue = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultVenue); err != nil {
				dst.InlineQueryResultVenue = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultVenue = nil
	}

	// try to unmarshal data into InlineQueryResultVideo
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultVideo)
	if err == nil {
		jsonInlineQueryResultVideo, _ := json.Marshal(dst.InlineQueryResultVideo)
		if string(jsonInlineQueryResultVideo) == "{}" { // empty struct
			dst.InlineQueryResultVideo = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultVideo); err != nil {
				dst.InlineQueryResultVideo = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultVideo = nil
	}

	// try to unmarshal data into InlineQueryResultVoice
	err = newStrictDecoder(data).Decode(&dst.InlineQueryResultVoice)
	if err == nil {
		jsonInlineQueryResultVoice, _ := json.Marshal(dst.InlineQueryResultVoice)
		if string(jsonInlineQueryResultVoice) == "{}" { // empty struct
			dst.InlineQueryResultVoice = nil
		} else {
			if err = validator.Validate(dst.InlineQueryResultVoice); err != nil {
				dst.InlineQueryResultVoice = nil
			} else {
				match++
			}
		}
	} else {
		dst.InlineQueryResultVoice = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InlineQueryResultArticle = nil
		dst.InlineQueryResultAudio = nil
		dst.InlineQueryResultCachedAudio = nil
		dst.InlineQueryResultCachedDocument = nil
		dst.InlineQueryResultCachedGif = nil
		dst.InlineQueryResultCachedMpeg4Gif = nil
		dst.InlineQueryResultCachedPhoto = nil
		dst.InlineQueryResultCachedSticker = nil
		dst.InlineQueryResultCachedVideo = nil
		dst.InlineQueryResultCachedVoice = nil
		dst.InlineQueryResultContact = nil
		dst.InlineQueryResultDocument = nil
		dst.InlineQueryResultGame = nil
		dst.InlineQueryResultGif = nil
		dst.InlineQueryResultLocation = nil
		dst.InlineQueryResultMpeg4Gif = nil
		dst.InlineQueryResultPhoto = nil
		dst.InlineQueryResultVenue = nil
		dst.InlineQueryResultVideo = nil
		dst.InlineQueryResultVoice = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InlineQueryResult)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InlineQueryResult)")
	}
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

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InlineQueryResult) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InlineQueryResultArticle != nil {
		return obj.InlineQueryResultArticle
	}

	if obj.InlineQueryResultAudio != nil {
		return obj.InlineQueryResultAudio
	}

	if obj.InlineQueryResultCachedAudio != nil {
		return obj.InlineQueryResultCachedAudio
	}

	if obj.InlineQueryResultCachedDocument != nil {
		return obj.InlineQueryResultCachedDocument
	}

	if obj.InlineQueryResultCachedGif != nil {
		return obj.InlineQueryResultCachedGif
	}

	if obj.InlineQueryResultCachedMpeg4Gif != nil {
		return obj.InlineQueryResultCachedMpeg4Gif
	}

	if obj.InlineQueryResultCachedPhoto != nil {
		return obj.InlineQueryResultCachedPhoto
	}

	if obj.InlineQueryResultCachedSticker != nil {
		return obj.InlineQueryResultCachedSticker
	}

	if obj.InlineQueryResultCachedVideo != nil {
		return obj.InlineQueryResultCachedVideo
	}

	if obj.InlineQueryResultCachedVoice != nil {
		return obj.InlineQueryResultCachedVoice
	}

	if obj.InlineQueryResultContact != nil {
		return obj.InlineQueryResultContact
	}

	if obj.InlineQueryResultDocument != nil {
		return obj.InlineQueryResultDocument
	}

	if obj.InlineQueryResultGame != nil {
		return obj.InlineQueryResultGame
	}

	if obj.InlineQueryResultGif != nil {
		return obj.InlineQueryResultGif
	}

	if obj.InlineQueryResultLocation != nil {
		return obj.InlineQueryResultLocation
	}

	if obj.InlineQueryResultMpeg4Gif != nil {
		return obj.InlineQueryResultMpeg4Gif
	}

	if obj.InlineQueryResultPhoto != nil {
		return obj.InlineQueryResultPhoto
	}

	if obj.InlineQueryResultVenue != nil {
		return obj.InlineQueryResultVenue
	}

	if obj.InlineQueryResultVideo != nil {
		return obj.InlineQueryResultVideo
	}

	if obj.InlineQueryResultVoice != nil {
		return obj.InlineQueryResultVoice
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj InlineQueryResult) GetActualInstanceValue() (interface{}) {
	if obj.InlineQueryResultArticle != nil {
		return *obj.InlineQueryResultArticle
	}

	if obj.InlineQueryResultAudio != nil {
		return *obj.InlineQueryResultAudio
	}

	if obj.InlineQueryResultCachedAudio != nil {
		return *obj.InlineQueryResultCachedAudio
	}

	if obj.InlineQueryResultCachedDocument != nil {
		return *obj.InlineQueryResultCachedDocument
	}

	if obj.InlineQueryResultCachedGif != nil {
		return *obj.InlineQueryResultCachedGif
	}

	if obj.InlineQueryResultCachedMpeg4Gif != nil {
		return *obj.InlineQueryResultCachedMpeg4Gif
	}

	if obj.InlineQueryResultCachedPhoto != nil {
		return *obj.InlineQueryResultCachedPhoto
	}

	if obj.InlineQueryResultCachedSticker != nil {
		return *obj.InlineQueryResultCachedSticker
	}

	if obj.InlineQueryResultCachedVideo != nil {
		return *obj.InlineQueryResultCachedVideo
	}

	if obj.InlineQueryResultCachedVoice != nil {
		return *obj.InlineQueryResultCachedVoice
	}

	if obj.InlineQueryResultContact != nil {
		return *obj.InlineQueryResultContact
	}

	if obj.InlineQueryResultDocument != nil {
		return *obj.InlineQueryResultDocument
	}

	if obj.InlineQueryResultGame != nil {
		return *obj.InlineQueryResultGame
	}

	if obj.InlineQueryResultGif != nil {
		return *obj.InlineQueryResultGif
	}

	if obj.InlineQueryResultLocation != nil {
		return *obj.InlineQueryResultLocation
	}

	if obj.InlineQueryResultMpeg4Gif != nil {
		return *obj.InlineQueryResultMpeg4Gif
	}

	if obj.InlineQueryResultPhoto != nil {
		return *obj.InlineQueryResultPhoto
	}

	if obj.InlineQueryResultVenue != nil {
		return *obj.InlineQueryResultVenue
	}

	if obj.InlineQueryResultVideo != nil {
		return *obj.InlineQueryResultVideo
	}

	if obj.InlineQueryResultVoice != nil {
		return *obj.InlineQueryResultVoice
	}

	// all schemas are nil
	return nil
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


