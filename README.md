# Go API client for tele_rest

Auto-generated OpenAPI schema

## Overview

- API version: 9.0.0
- Package version: 9.0.0
- Build date: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
- Generator version: 7.14.0

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import tele_rest "github.com/tele-api/tele-rest-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `tele_rest.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), tele_rest.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `tele_rest.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), tele_rest.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `tele_rest.ContextOperationServerIndices` and `tele_rest.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), tele_rest.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), tele_rest.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.telegram.org/bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**AddStickerToSetPost**](docs/DefaultAPI.md#addstickertosetpost) | **Post** /addStickerToSet | 
*DefaultAPI* | [**AnswerCallbackQueryPost**](docs/DefaultAPI.md#answercallbackquerypost) | **Post** /answerCallbackQuery | 
*DefaultAPI* | [**AnswerInlineQueryPost**](docs/DefaultAPI.md#answerinlinequerypost) | **Post** /answerInlineQuery | 
*DefaultAPI* | [**AnswerPreCheckoutQueryPost**](docs/DefaultAPI.md#answerprecheckoutquerypost) | **Post** /answerPreCheckoutQuery | 
*DefaultAPI* | [**AnswerShippingQueryPost**](docs/DefaultAPI.md#answershippingquerypost) | **Post** /answerShippingQuery | 
*DefaultAPI* | [**AnswerWebAppQueryPost**](docs/DefaultAPI.md#answerwebappquerypost) | **Post** /answerWebAppQuery | 
*DefaultAPI* | [**ApproveChatJoinRequestPost**](docs/DefaultAPI.md#approvechatjoinrequestpost) | **Post** /approveChatJoinRequest | 
*DefaultAPI* | [**BanChatMemberPost**](docs/DefaultAPI.md#banchatmemberpost) | **Post** /banChatMember | 
*DefaultAPI* | [**BanChatSenderChatPost**](docs/DefaultAPI.md#banchatsenderchatpost) | **Post** /banChatSenderChat | 
*DefaultAPI* | [**CloseForumTopicPost**](docs/DefaultAPI.md#closeforumtopicpost) | **Post** /closeForumTopic | 
*DefaultAPI* | [**CloseGeneralForumTopicPost**](docs/DefaultAPI.md#closegeneralforumtopicpost) | **Post** /closeGeneralForumTopic | 
*DefaultAPI* | [**ClosePost**](docs/DefaultAPI.md#closepost) | **Post** /close | 
*DefaultAPI* | [**ConvertGiftToStarsPost**](docs/DefaultAPI.md#convertgifttostarspost) | **Post** /convertGiftToStars | 
*DefaultAPI* | [**CopyMessagePost**](docs/DefaultAPI.md#copymessagepost) | **Post** /copyMessage | 
*DefaultAPI* | [**CopyMessagesPost**](docs/DefaultAPI.md#copymessagespost) | **Post** /copyMessages | 
*DefaultAPI* | [**CreateChatInviteLinkPost**](docs/DefaultAPI.md#createchatinvitelinkpost) | **Post** /createChatInviteLink | 
*DefaultAPI* | [**CreateChatSubscriptionInviteLinkPost**](docs/DefaultAPI.md#createchatsubscriptioninvitelinkpost) | **Post** /createChatSubscriptionInviteLink | 
*DefaultAPI* | [**CreateForumTopicPost**](docs/DefaultAPI.md#createforumtopicpost) | **Post** /createForumTopic | 
*DefaultAPI* | [**CreateInvoiceLinkPost**](docs/DefaultAPI.md#createinvoicelinkpost) | **Post** /createInvoiceLink | 
*DefaultAPI* | [**CreateNewStickerSetPost**](docs/DefaultAPI.md#createnewstickersetpost) | **Post** /createNewStickerSet | 
*DefaultAPI* | [**DeclineChatJoinRequestPost**](docs/DefaultAPI.md#declinechatjoinrequestpost) | **Post** /declineChatJoinRequest | 
*DefaultAPI* | [**DeleteBusinessMessagesPost**](docs/DefaultAPI.md#deletebusinessmessagespost) | **Post** /deleteBusinessMessages | 
*DefaultAPI* | [**DeleteChatPhotoPost**](docs/DefaultAPI.md#deletechatphotopost) | **Post** /deleteChatPhoto | 
*DefaultAPI* | [**DeleteChatStickerSetPost**](docs/DefaultAPI.md#deletechatstickersetpost) | **Post** /deleteChatStickerSet | 
*DefaultAPI* | [**DeleteForumTopicPost**](docs/DefaultAPI.md#deleteforumtopicpost) | **Post** /deleteForumTopic | 
*DefaultAPI* | [**DeleteMessagePost**](docs/DefaultAPI.md#deletemessagepost) | **Post** /deleteMessage | 
*DefaultAPI* | [**DeleteMessagesPost**](docs/DefaultAPI.md#deletemessagespost) | **Post** /deleteMessages | 
*DefaultAPI* | [**DeleteMyCommandsPost**](docs/DefaultAPI.md#deletemycommandspost) | **Post** /deleteMyCommands | 
*DefaultAPI* | [**DeleteStickerFromSetPost**](docs/DefaultAPI.md#deletestickerfromsetpost) | **Post** /deleteStickerFromSet | 
*DefaultAPI* | [**DeleteStickerSetPost**](docs/DefaultAPI.md#deletestickersetpost) | **Post** /deleteStickerSet | 
*DefaultAPI* | [**DeleteStoryPost**](docs/DefaultAPI.md#deletestorypost) | **Post** /deleteStory | 
*DefaultAPI* | [**DeleteWebhookPost**](docs/DefaultAPI.md#deletewebhookpost) | **Post** /deleteWebhook | 
*DefaultAPI* | [**EditChatInviteLinkPost**](docs/DefaultAPI.md#editchatinvitelinkpost) | **Post** /editChatInviteLink | 
*DefaultAPI* | [**EditChatSubscriptionInviteLinkPost**](docs/DefaultAPI.md#editchatsubscriptioninvitelinkpost) | **Post** /editChatSubscriptionInviteLink | 
*DefaultAPI* | [**EditForumTopicPost**](docs/DefaultAPI.md#editforumtopicpost) | **Post** /editForumTopic | 
*DefaultAPI* | [**EditGeneralForumTopicPost**](docs/DefaultAPI.md#editgeneralforumtopicpost) | **Post** /editGeneralForumTopic | 
*DefaultAPI* | [**EditMessageCaptionPost**](docs/DefaultAPI.md#editmessagecaptionpost) | **Post** /editMessageCaption | 
*DefaultAPI* | [**EditMessageLiveLocationPost**](docs/DefaultAPI.md#editmessagelivelocationpost) | **Post** /editMessageLiveLocation | 
*DefaultAPI* | [**EditMessageMediaPost**](docs/DefaultAPI.md#editmessagemediapost) | **Post** /editMessageMedia | 
*DefaultAPI* | [**EditMessageReplyMarkupPost**](docs/DefaultAPI.md#editmessagereplymarkuppost) | **Post** /editMessageReplyMarkup | 
*DefaultAPI* | [**EditMessageTextPost**](docs/DefaultAPI.md#editmessagetextpost) | **Post** /editMessageText | 
*DefaultAPI* | [**EditStoryPost**](docs/DefaultAPI.md#editstorypost) | **Post** /editStory | 
*DefaultAPI* | [**EditUserStarSubscriptionPost**](docs/DefaultAPI.md#edituserstarsubscriptionpost) | **Post** /editUserStarSubscription | 
*DefaultAPI* | [**ExportChatInviteLinkPost**](docs/DefaultAPI.md#exportchatinvitelinkpost) | **Post** /exportChatInviteLink | 
*DefaultAPI* | [**ForwardMessagePost**](docs/DefaultAPI.md#forwardmessagepost) | **Post** /forwardMessage | 
*DefaultAPI* | [**ForwardMessagesPost**](docs/DefaultAPI.md#forwardmessagespost) | **Post** /forwardMessages | 
*DefaultAPI* | [**GetAvailableGiftsPost**](docs/DefaultAPI.md#getavailablegiftspost) | **Post** /getAvailableGifts | 
*DefaultAPI* | [**GetBusinessAccountGiftsPost**](docs/DefaultAPI.md#getbusinessaccountgiftspost) | **Post** /getBusinessAccountGifts | 
*DefaultAPI* | [**GetBusinessAccountStarBalancePost**](docs/DefaultAPI.md#getbusinessaccountstarbalancepost) | **Post** /getBusinessAccountStarBalance | 
*DefaultAPI* | [**GetBusinessConnectionPost**](docs/DefaultAPI.md#getbusinessconnectionpost) | **Post** /getBusinessConnection | 
*DefaultAPI* | [**GetChatAdministratorsPost**](docs/DefaultAPI.md#getchatadministratorspost) | **Post** /getChatAdministrators | 
*DefaultAPI* | [**GetChatMemberCountPost**](docs/DefaultAPI.md#getchatmembercountpost) | **Post** /getChatMemberCount | 
*DefaultAPI* | [**GetChatMemberPost**](docs/DefaultAPI.md#getchatmemberpost) | **Post** /getChatMember | 
*DefaultAPI* | [**GetChatMenuButtonPost**](docs/DefaultAPI.md#getchatmenubuttonpost) | **Post** /getChatMenuButton | 
*DefaultAPI* | [**GetChatPost**](docs/DefaultAPI.md#getchatpost) | **Post** /getChat | 
*DefaultAPI* | [**GetCustomEmojiStickersPost**](docs/DefaultAPI.md#getcustomemojistickerspost) | **Post** /getCustomEmojiStickers | 
*DefaultAPI* | [**GetFilePost**](docs/DefaultAPI.md#getfilepost) | **Post** /getFile | 
*DefaultAPI* | [**GetForumTopicIconStickersPost**](docs/DefaultAPI.md#getforumtopiciconstickerspost) | **Post** /getForumTopicIconStickers | 
*DefaultAPI* | [**GetGameHighScoresPost**](docs/DefaultAPI.md#getgamehighscorespost) | **Post** /getGameHighScores | 
*DefaultAPI* | [**GetMePost**](docs/DefaultAPI.md#getmepost) | **Post** /getMe | 
*DefaultAPI* | [**GetMyCommandsPost**](docs/DefaultAPI.md#getmycommandspost) | **Post** /getMyCommands | 
*DefaultAPI* | [**GetMyDefaultAdministratorRightsPost**](docs/DefaultAPI.md#getmydefaultadministratorrightspost) | **Post** /getMyDefaultAdministratorRights | 
*DefaultAPI* | [**GetMyDescriptionPost**](docs/DefaultAPI.md#getmydescriptionpost) | **Post** /getMyDescription | 
*DefaultAPI* | [**GetMyNamePost**](docs/DefaultAPI.md#getmynamepost) | **Post** /getMyName | 
*DefaultAPI* | [**GetMyShortDescriptionPost**](docs/DefaultAPI.md#getmyshortdescriptionpost) | **Post** /getMyShortDescription | 
*DefaultAPI* | [**GetStarTransactionsPost**](docs/DefaultAPI.md#getstartransactionspost) | **Post** /getStarTransactions | 
*DefaultAPI* | [**GetStickerSetPost**](docs/DefaultAPI.md#getstickersetpost) | **Post** /getStickerSet | 
*DefaultAPI* | [**GetUpdatesPost**](docs/DefaultAPI.md#getupdatespost) | **Post** /getUpdates | 
*DefaultAPI* | [**GetUserChatBoostsPost**](docs/DefaultAPI.md#getuserchatboostspost) | **Post** /getUserChatBoosts | 
*DefaultAPI* | [**GetUserProfilePhotosPost**](docs/DefaultAPI.md#getuserprofilephotospost) | **Post** /getUserProfilePhotos | 
*DefaultAPI* | [**GetWebhookInfoPost**](docs/DefaultAPI.md#getwebhookinfopost) | **Post** /getWebhookInfo | 
*DefaultAPI* | [**GiftPremiumSubscriptionPost**](docs/DefaultAPI.md#giftpremiumsubscriptionpost) | **Post** /giftPremiumSubscription | 
*DefaultAPI* | [**HideGeneralForumTopicPost**](docs/DefaultAPI.md#hidegeneralforumtopicpost) | **Post** /hideGeneralForumTopic | 
*DefaultAPI* | [**LeaveChatPost**](docs/DefaultAPI.md#leavechatpost) | **Post** /leaveChat | 
*DefaultAPI* | [**LogOutPost**](docs/DefaultAPI.md#logoutpost) | **Post** /logOut | 
*DefaultAPI* | [**PinChatMessagePost**](docs/DefaultAPI.md#pinchatmessagepost) | **Post** /pinChatMessage | 
*DefaultAPI* | [**PostStoryPost**](docs/DefaultAPI.md#poststorypost) | **Post** /postStory | 
*DefaultAPI* | [**PromoteChatMemberPost**](docs/DefaultAPI.md#promotechatmemberpost) | **Post** /promoteChatMember | 
*DefaultAPI* | [**ReadBusinessMessagePost**](docs/DefaultAPI.md#readbusinessmessagepost) | **Post** /readBusinessMessage | 
*DefaultAPI* | [**RefundStarPaymentPost**](docs/DefaultAPI.md#refundstarpaymentpost) | **Post** /refundStarPayment | 
*DefaultAPI* | [**RemoveBusinessAccountProfilePhotoPost**](docs/DefaultAPI.md#removebusinessaccountprofilephotopost) | **Post** /removeBusinessAccountProfilePhoto | 
*DefaultAPI* | [**RemoveChatVerificationPost**](docs/DefaultAPI.md#removechatverificationpost) | **Post** /removeChatVerification | 
*DefaultAPI* | [**RemoveUserVerificationPost**](docs/DefaultAPI.md#removeuserverificationpost) | **Post** /removeUserVerification | 
*DefaultAPI* | [**ReopenForumTopicPost**](docs/DefaultAPI.md#reopenforumtopicpost) | **Post** /reopenForumTopic | 
*DefaultAPI* | [**ReopenGeneralForumTopicPost**](docs/DefaultAPI.md#reopengeneralforumtopicpost) | **Post** /reopenGeneralForumTopic | 
*DefaultAPI* | [**ReplaceStickerInSetPost**](docs/DefaultAPI.md#replacestickerinsetpost) | **Post** /replaceStickerInSet | 
*DefaultAPI* | [**RestrictChatMemberPost**](docs/DefaultAPI.md#restrictchatmemberpost) | **Post** /restrictChatMember | 
*DefaultAPI* | [**RevokeChatInviteLinkPost**](docs/DefaultAPI.md#revokechatinvitelinkpost) | **Post** /revokeChatInviteLink | 
*DefaultAPI* | [**SavePreparedInlineMessagePost**](docs/DefaultAPI.md#savepreparedinlinemessagepost) | **Post** /savePreparedInlineMessage | 
*DefaultAPI* | [**SendAnimationPost**](docs/DefaultAPI.md#sendanimationpost) | **Post** /sendAnimation | 
*DefaultAPI* | [**SendAudioPost**](docs/DefaultAPI.md#sendaudiopost) | **Post** /sendAudio | 
*DefaultAPI* | [**SendChatActionPost**](docs/DefaultAPI.md#sendchatactionpost) | **Post** /sendChatAction | 
*DefaultAPI* | [**SendContactPost**](docs/DefaultAPI.md#sendcontactpost) | **Post** /sendContact | 
*DefaultAPI* | [**SendDicePost**](docs/DefaultAPI.md#senddicepost) | **Post** /sendDice | 
*DefaultAPI* | [**SendDocumentPost**](docs/DefaultAPI.md#senddocumentpost) | **Post** /sendDocument | 
*DefaultAPI* | [**SendGamePost**](docs/DefaultAPI.md#sendgamepost) | **Post** /sendGame | 
*DefaultAPI* | [**SendGiftPost**](docs/DefaultAPI.md#sendgiftpost) | **Post** /sendGift | 
*DefaultAPI* | [**SendInvoicePost**](docs/DefaultAPI.md#sendinvoicepost) | **Post** /sendInvoice | 
*DefaultAPI* | [**SendLocationPost**](docs/DefaultAPI.md#sendlocationpost) | **Post** /sendLocation | 
*DefaultAPI* | [**SendMediaGroupPost**](docs/DefaultAPI.md#sendmediagrouppost) | **Post** /sendMediaGroup | 
*DefaultAPI* | [**SendMessagePost**](docs/DefaultAPI.md#sendmessagepost) | **Post** /sendMessage | 
*DefaultAPI* | [**SendPaidMediaPost**](docs/DefaultAPI.md#sendpaidmediapost) | **Post** /sendPaidMedia | 
*DefaultAPI* | [**SendPhotoPost**](docs/DefaultAPI.md#sendphotopost) | **Post** /sendPhoto | 
*DefaultAPI* | [**SendPollPost**](docs/DefaultAPI.md#sendpollpost) | **Post** /sendPoll | 
*DefaultAPI* | [**SendStickerPost**](docs/DefaultAPI.md#sendstickerpost) | **Post** /sendSticker | 
*DefaultAPI* | [**SendVenuePost**](docs/DefaultAPI.md#sendvenuepost) | **Post** /sendVenue | 
*DefaultAPI* | [**SendVideoNotePost**](docs/DefaultAPI.md#sendvideonotepost) | **Post** /sendVideoNote | 
*DefaultAPI* | [**SendVideoPost**](docs/DefaultAPI.md#sendvideopost) | **Post** /sendVideo | 
*DefaultAPI* | [**SendVoicePost**](docs/DefaultAPI.md#sendvoicepost) | **Post** /sendVoice | 
*DefaultAPI* | [**SetBusinessAccountBioPost**](docs/DefaultAPI.md#setbusinessaccountbiopost) | **Post** /setBusinessAccountBio | 
*DefaultAPI* | [**SetBusinessAccountGiftSettingsPost**](docs/DefaultAPI.md#setbusinessaccountgiftsettingspost) | **Post** /setBusinessAccountGiftSettings | 
*DefaultAPI* | [**SetBusinessAccountNamePost**](docs/DefaultAPI.md#setbusinessaccountnamepost) | **Post** /setBusinessAccountName | 
*DefaultAPI* | [**SetBusinessAccountProfilePhotoPost**](docs/DefaultAPI.md#setbusinessaccountprofilephotopost) | **Post** /setBusinessAccountProfilePhoto | 
*DefaultAPI* | [**SetBusinessAccountUsernamePost**](docs/DefaultAPI.md#setbusinessaccountusernamepost) | **Post** /setBusinessAccountUsername | 
*DefaultAPI* | [**SetChatAdministratorCustomTitlePost**](docs/DefaultAPI.md#setchatadministratorcustomtitlepost) | **Post** /setChatAdministratorCustomTitle | 
*DefaultAPI* | [**SetChatDescriptionPost**](docs/DefaultAPI.md#setchatdescriptionpost) | **Post** /setChatDescription | 
*DefaultAPI* | [**SetChatMenuButtonPost**](docs/DefaultAPI.md#setchatmenubuttonpost) | **Post** /setChatMenuButton | 
*DefaultAPI* | [**SetChatPermissionsPost**](docs/DefaultAPI.md#setchatpermissionspost) | **Post** /setChatPermissions | 
*DefaultAPI* | [**SetChatPhotoPost**](docs/DefaultAPI.md#setchatphotopost) | **Post** /setChatPhoto | 
*DefaultAPI* | [**SetChatStickerSetPost**](docs/DefaultAPI.md#setchatstickersetpost) | **Post** /setChatStickerSet | 
*DefaultAPI* | [**SetChatTitlePost**](docs/DefaultAPI.md#setchattitlepost) | **Post** /setChatTitle | 
*DefaultAPI* | [**SetCustomEmojiStickerSetThumbnailPost**](docs/DefaultAPI.md#setcustomemojistickersetthumbnailpost) | **Post** /setCustomEmojiStickerSetThumbnail | 
*DefaultAPI* | [**SetGameScorePost**](docs/DefaultAPI.md#setgamescorepost) | **Post** /setGameScore | 
*DefaultAPI* | [**SetMessageReactionPost**](docs/DefaultAPI.md#setmessagereactionpost) | **Post** /setMessageReaction | 
*DefaultAPI* | [**SetMyCommandsPost**](docs/DefaultAPI.md#setmycommandspost) | **Post** /setMyCommands | 
*DefaultAPI* | [**SetMyDefaultAdministratorRightsPost**](docs/DefaultAPI.md#setmydefaultadministratorrightspost) | **Post** /setMyDefaultAdministratorRights | 
*DefaultAPI* | [**SetMyDescriptionPost**](docs/DefaultAPI.md#setmydescriptionpost) | **Post** /setMyDescription | 
*DefaultAPI* | [**SetMyNamePost**](docs/DefaultAPI.md#setmynamepost) | **Post** /setMyName | 
*DefaultAPI* | [**SetMyShortDescriptionPost**](docs/DefaultAPI.md#setmyshortdescriptionpost) | **Post** /setMyShortDescription | 
*DefaultAPI* | [**SetPassportDataErrorsPost**](docs/DefaultAPI.md#setpassportdataerrorspost) | **Post** /setPassportDataErrors | 
*DefaultAPI* | [**SetStickerEmojiListPost**](docs/DefaultAPI.md#setstickeremojilistpost) | **Post** /setStickerEmojiList | 
*DefaultAPI* | [**SetStickerKeywordsPost**](docs/DefaultAPI.md#setstickerkeywordspost) | **Post** /setStickerKeywords | 
*DefaultAPI* | [**SetStickerMaskPositionPost**](docs/DefaultAPI.md#setstickermaskpositionpost) | **Post** /setStickerMaskPosition | 
*DefaultAPI* | [**SetStickerPositionInSetPost**](docs/DefaultAPI.md#setstickerpositioninsetpost) | **Post** /setStickerPositionInSet | 
*DefaultAPI* | [**SetStickerSetThumbnailPost**](docs/DefaultAPI.md#setstickersetthumbnailpost) | **Post** /setStickerSetThumbnail | 
*DefaultAPI* | [**SetStickerSetTitlePost**](docs/DefaultAPI.md#setstickersettitlepost) | **Post** /setStickerSetTitle | 
*DefaultAPI* | [**SetUserEmojiStatusPost**](docs/DefaultAPI.md#setuseremojistatuspost) | **Post** /setUserEmojiStatus | 
*DefaultAPI* | [**SetWebhookPost**](docs/DefaultAPI.md#setwebhookpost) | **Post** /setWebhook | 
*DefaultAPI* | [**StopMessageLiveLocationPost**](docs/DefaultAPI.md#stopmessagelivelocationpost) | **Post** /stopMessageLiveLocation | 
*DefaultAPI* | [**StopPollPost**](docs/DefaultAPI.md#stoppollpost) | **Post** /stopPoll | 
*DefaultAPI* | [**TransferBusinessAccountStarsPost**](docs/DefaultAPI.md#transferbusinessaccountstarspost) | **Post** /transferBusinessAccountStars | 
*DefaultAPI* | [**TransferGiftPost**](docs/DefaultAPI.md#transfergiftpost) | **Post** /transferGift | 
*DefaultAPI* | [**UnbanChatMemberPost**](docs/DefaultAPI.md#unbanchatmemberpost) | **Post** /unbanChatMember | 
*DefaultAPI* | [**UnbanChatSenderChatPost**](docs/DefaultAPI.md#unbanchatsenderchatpost) | **Post** /unbanChatSenderChat | 
*DefaultAPI* | [**UnhideGeneralForumTopicPost**](docs/DefaultAPI.md#unhidegeneralforumtopicpost) | **Post** /unhideGeneralForumTopic | 
*DefaultAPI* | [**UnpinAllChatMessagesPost**](docs/DefaultAPI.md#unpinallchatmessagespost) | **Post** /unpinAllChatMessages | 
*DefaultAPI* | [**UnpinAllForumTopicMessagesPost**](docs/DefaultAPI.md#unpinallforumtopicmessagespost) | **Post** /unpinAllForumTopicMessages | 
*DefaultAPI* | [**UnpinAllGeneralForumTopicMessagesPost**](docs/DefaultAPI.md#unpinallgeneralforumtopicmessagespost) | **Post** /unpinAllGeneralForumTopicMessages | 
*DefaultAPI* | [**UnpinChatMessagePost**](docs/DefaultAPI.md#unpinchatmessagepost) | **Post** /unpinChatMessage | 
*DefaultAPI* | [**UpgradeGiftPost**](docs/DefaultAPI.md#upgradegiftpost) | **Post** /upgradeGift | 
*DefaultAPI* | [**UploadStickerFilePost**](docs/DefaultAPI.md#uploadstickerfilepost) | **Post** /uploadStickerFile | 
*DefaultAPI* | [**VerifyChatPost**](docs/DefaultAPI.md#verifychatpost) | **Post** /verifyChat | 
*DefaultAPI* | [**VerifyUserPost**](docs/DefaultAPI.md#verifyuserpost) | **Post** /verifyUser | 


## Documentation For Models

 - [AcceptedGiftTypes](docs/AcceptedGiftTypes.md)
 - [AffiliateInfo](docs/AffiliateInfo.md)
 - [Animation](docs/Animation.md)
 - [AnswerCallbackQueryPostRequest](docs/AnswerCallbackQueryPostRequest.md)
 - [AnswerInlineQueryPostRequest](docs/AnswerInlineQueryPostRequest.md)
 - [AnswerPreCheckoutQueryPostRequest](docs/AnswerPreCheckoutQueryPostRequest.md)
 - [AnswerShippingQueryPostRequest](docs/AnswerShippingQueryPostRequest.md)
 - [AnswerWebAppQueryPost200Response](docs/AnswerWebAppQueryPost200Response.md)
 - [AnswerWebAppQueryPostRequest](docs/AnswerWebAppQueryPostRequest.md)
 - [ApproveChatJoinRequestPostRequest](docs/ApproveChatJoinRequestPostRequest.md)
 - [Audio](docs/Audio.md)
 - [BackgroundFill](docs/BackgroundFill.md)
 - [BackgroundFillFreeformGradient](docs/BackgroundFillFreeformGradient.md)
 - [BackgroundFillGradient](docs/BackgroundFillGradient.md)
 - [BackgroundFillSolid](docs/BackgroundFillSolid.md)
 - [BackgroundType](docs/BackgroundType.md)
 - [BackgroundTypeChatTheme](docs/BackgroundTypeChatTheme.md)
 - [BackgroundTypeFill](docs/BackgroundTypeFill.md)
 - [BackgroundTypePattern](docs/BackgroundTypePattern.md)
 - [BackgroundTypeWallpaper](docs/BackgroundTypeWallpaper.md)
 - [BanChatMemberPostRequest](docs/BanChatMemberPostRequest.md)
 - [BanChatMemberPostRequestChatId](docs/BanChatMemberPostRequestChatId.md)
 - [BanChatSenderChatPostRequest](docs/BanChatSenderChatPostRequest.md)
 - [Birthdate](docs/Birthdate.md)
 - [BotCommand](docs/BotCommand.md)
 - [BotCommandScope](docs/BotCommandScope.md)
 - [BotCommandScopeAllChatAdministrators](docs/BotCommandScopeAllChatAdministrators.md)
 - [BotCommandScopeAllGroupChats](docs/BotCommandScopeAllGroupChats.md)
 - [BotCommandScopeAllPrivateChats](docs/BotCommandScopeAllPrivateChats.md)
 - [BotCommandScopeChat](docs/BotCommandScopeChat.md)
 - [BotCommandScopeChatAdministrators](docs/BotCommandScopeChatAdministrators.md)
 - [BotCommandScopeChatMember](docs/BotCommandScopeChatMember.md)
 - [BotCommandScopeDefault](docs/BotCommandScopeDefault.md)
 - [BotDescription](docs/BotDescription.md)
 - [BotName](docs/BotName.md)
 - [BotShortDescription](docs/BotShortDescription.md)
 - [BusinessBotRights](docs/BusinessBotRights.md)
 - [BusinessConnection](docs/BusinessConnection.md)
 - [BusinessIntro](docs/BusinessIntro.md)
 - [BusinessLocation](docs/BusinessLocation.md)
 - [BusinessMessagesDeleted](docs/BusinessMessagesDeleted.md)
 - [BusinessOpeningHours](docs/BusinessOpeningHours.md)
 - [BusinessOpeningHoursInterval](docs/BusinessOpeningHoursInterval.md)
 - [CallbackQuery](docs/CallbackQuery.md)
 - [Chat](docs/Chat.md)
 - [ChatAdministratorRights](docs/ChatAdministratorRights.md)
 - [ChatBackground](docs/ChatBackground.md)
 - [ChatBoost](docs/ChatBoost.md)
 - [ChatBoostAdded](docs/ChatBoostAdded.md)
 - [ChatBoostRemoved](docs/ChatBoostRemoved.md)
 - [ChatBoostSource](docs/ChatBoostSource.md)
 - [ChatBoostSourceGiftCode](docs/ChatBoostSourceGiftCode.md)
 - [ChatBoostSourceGiveaway](docs/ChatBoostSourceGiveaway.md)
 - [ChatBoostSourcePremium](docs/ChatBoostSourcePremium.md)
 - [ChatBoostUpdated](docs/ChatBoostUpdated.md)
 - [ChatFullInfo](docs/ChatFullInfo.md)
 - [ChatInviteLink](docs/ChatInviteLink.md)
 - [ChatJoinRequest](docs/ChatJoinRequest.md)
 - [ChatLocation](docs/ChatLocation.md)
 - [ChatMember](docs/ChatMember.md)
 - [ChatMemberAdministrator](docs/ChatMemberAdministrator.md)
 - [ChatMemberBanned](docs/ChatMemberBanned.md)
 - [ChatMemberLeft](docs/ChatMemberLeft.md)
 - [ChatMemberMember](docs/ChatMemberMember.md)
 - [ChatMemberOwner](docs/ChatMemberOwner.md)
 - [ChatMemberRestricted](docs/ChatMemberRestricted.md)
 - [ChatMemberUpdated](docs/ChatMemberUpdated.md)
 - [ChatPermissions](docs/ChatPermissions.md)
 - [ChatPhoto](docs/ChatPhoto.md)
 - [ChatShared](docs/ChatShared.md)
 - [ChosenInlineResult](docs/ChosenInlineResult.md)
 - [CloseForumTopicPostRequest](docs/CloseForumTopicPostRequest.md)
 - [Contact](docs/Contact.md)
 - [ConvertGiftToStarsPostRequest](docs/ConvertGiftToStarsPostRequest.md)
 - [CopyMessagePost200Response](docs/CopyMessagePost200Response.md)
 - [CopyMessagePostRequest](docs/CopyMessagePostRequest.md)
 - [CopyMessagesPostRequest](docs/CopyMessagesPostRequest.md)
 - [CopyTextButton](docs/CopyTextButton.md)
 - [CreateChatInviteLinkPost200Response](docs/CreateChatInviteLinkPost200Response.md)
 - [CreateChatInviteLinkPostRequest](docs/CreateChatInviteLinkPostRequest.md)
 - [CreateChatSubscriptionInviteLinkPostRequest](docs/CreateChatSubscriptionInviteLinkPostRequest.md)
 - [CreateChatSubscriptionInviteLinkPostRequestChatId](docs/CreateChatSubscriptionInviteLinkPostRequestChatId.md)
 - [CreateForumTopicPost200Response](docs/CreateForumTopicPost200Response.md)
 - [CreateForumTopicPostRequest](docs/CreateForumTopicPostRequest.md)
 - [CreateInvoiceLinkPostRequest](docs/CreateInvoiceLinkPostRequest.md)
 - [DeleteBusinessMessagesPostRequest](docs/DeleteBusinessMessagesPostRequest.md)
 - [DeleteChatStickerSetPostRequest](docs/DeleteChatStickerSetPostRequest.md)
 - [DeleteMessagePostRequest](docs/DeleteMessagePostRequest.md)
 - [DeleteMessagesPostRequest](docs/DeleteMessagesPostRequest.md)
 - [DeleteMyCommandsPostRequest](docs/DeleteMyCommandsPostRequest.md)
 - [DeleteStickerFromSetPostRequest](docs/DeleteStickerFromSetPostRequest.md)
 - [DeleteStickerSetPostRequest](docs/DeleteStickerSetPostRequest.md)
 - [DeleteStoryPostRequest](docs/DeleteStoryPostRequest.md)
 - [DeleteWebhookPostRequest](docs/DeleteWebhookPostRequest.md)
 - [Dice](docs/Dice.md)
 - [Document](docs/Document.md)
 - [EditChatInviteLinkPostRequest](docs/EditChatInviteLinkPostRequest.md)
 - [EditChatSubscriptionInviteLinkPostRequest](docs/EditChatSubscriptionInviteLinkPostRequest.md)
 - [EditForumTopicPostRequest](docs/EditForumTopicPostRequest.md)
 - [EditGeneralForumTopicPostRequest](docs/EditGeneralForumTopicPostRequest.md)
 - [EditMessageCaptionPostRequest](docs/EditMessageCaptionPostRequest.md)
 - [EditMessageLiveLocationPostRequest](docs/EditMessageLiveLocationPostRequest.md)
 - [EditMessageReplyMarkupPostRequest](docs/EditMessageReplyMarkupPostRequest.md)
 - [EditMessageTextPost200Response](docs/EditMessageTextPost200Response.md)
 - [EditMessageTextPost200ResponseResult](docs/EditMessageTextPost200ResponseResult.md)
 - [EditMessageTextPostRequest](docs/EditMessageTextPostRequest.md)
 - [EditMessageTextPostRequestChatId](docs/EditMessageTextPostRequestChatId.md)
 - [EditUserStarSubscriptionPostRequest](docs/EditUserStarSubscriptionPostRequest.md)
 - [EncryptedCredentials](docs/EncryptedCredentials.md)
 - [EncryptedPassportElement](docs/EncryptedPassportElement.md)
 - [Error](docs/Error.md)
 - [ExportChatInviteLinkPost200Response](docs/ExportChatInviteLinkPost200Response.md)
 - [ExportChatInviteLinkPostRequest](docs/ExportChatInviteLinkPostRequest.md)
 - [ExternalReplyInfo](docs/ExternalReplyInfo.md)
 - [File](docs/File.md)
 - [ForceReply](docs/ForceReply.md)
 - [ForumTopic](docs/ForumTopic.md)
 - [ForumTopicCreated](docs/ForumTopicCreated.md)
 - [ForumTopicEdited](docs/ForumTopicEdited.md)
 - [ForwardMessagePostRequest](docs/ForwardMessagePostRequest.md)
 - [ForwardMessagePostRequestFromChatId](docs/ForwardMessagePostRequestFromChatId.md)
 - [ForwardMessagesPost200Response](docs/ForwardMessagesPost200Response.md)
 - [ForwardMessagesPostRequest](docs/ForwardMessagesPostRequest.md)
 - [ForwardMessagesPostRequestFromChatId](docs/ForwardMessagesPostRequestFromChatId.md)
 - [Game](docs/Game.md)
 - [GameHighScore](docs/GameHighScore.md)
 - [GetAvailableGiftsPost200Response](docs/GetAvailableGiftsPost200Response.md)
 - [GetBusinessAccountGiftsPost200Response](docs/GetBusinessAccountGiftsPost200Response.md)
 - [GetBusinessAccountGiftsPostRequest](docs/GetBusinessAccountGiftsPostRequest.md)
 - [GetBusinessAccountStarBalancePost200Response](docs/GetBusinessAccountStarBalancePost200Response.md)
 - [GetBusinessConnectionPost200Response](docs/GetBusinessConnectionPost200Response.md)
 - [GetBusinessConnectionPostRequest](docs/GetBusinessConnectionPostRequest.md)
 - [GetChatAdministratorsPost200Response](docs/GetChatAdministratorsPost200Response.md)
 - [GetChatMemberCountPost200Response](docs/GetChatMemberCountPost200Response.md)
 - [GetChatMemberPost200Response](docs/GetChatMemberPost200Response.md)
 - [GetChatMemberPostRequest](docs/GetChatMemberPostRequest.md)
 - [GetChatMenuButtonPost200Response](docs/GetChatMenuButtonPost200Response.md)
 - [GetChatMenuButtonPostRequest](docs/GetChatMenuButtonPostRequest.md)
 - [GetChatPost200Response](docs/GetChatPost200Response.md)
 - [GetCustomEmojiStickersPostRequest](docs/GetCustomEmojiStickersPostRequest.md)
 - [GetFilePost200Response](docs/GetFilePost200Response.md)
 - [GetFilePostRequest](docs/GetFilePostRequest.md)
 - [GetForumTopicIconStickersPost200Response](docs/GetForumTopicIconStickersPost200Response.md)
 - [GetGameHighScoresPost200Response](docs/GetGameHighScoresPost200Response.md)
 - [GetGameHighScoresPostRequest](docs/GetGameHighScoresPostRequest.md)
 - [GetMePost200Response](docs/GetMePost200Response.md)
 - [GetMyCommandsPost200Response](docs/GetMyCommandsPost200Response.md)
 - [GetMyCommandsPostRequest](docs/GetMyCommandsPostRequest.md)
 - [GetMyDefaultAdministratorRightsPost200Response](docs/GetMyDefaultAdministratorRightsPost200Response.md)
 - [GetMyDefaultAdministratorRightsPostRequest](docs/GetMyDefaultAdministratorRightsPostRequest.md)
 - [GetMyDescriptionPost200Response](docs/GetMyDescriptionPost200Response.md)
 - [GetMyNamePost200Response](docs/GetMyNamePost200Response.md)
 - [GetMyNamePostRequest](docs/GetMyNamePostRequest.md)
 - [GetMyShortDescriptionPost200Response](docs/GetMyShortDescriptionPost200Response.md)
 - [GetStarTransactionsPost200Response](docs/GetStarTransactionsPost200Response.md)
 - [GetStarTransactionsPostRequest](docs/GetStarTransactionsPostRequest.md)
 - [GetStickerSetPost200Response](docs/GetStickerSetPost200Response.md)
 - [GetStickerSetPostRequest](docs/GetStickerSetPostRequest.md)
 - [GetUpdatesPost200Response](docs/GetUpdatesPost200Response.md)
 - [GetUpdatesPostRequest](docs/GetUpdatesPostRequest.md)
 - [GetUserChatBoostsPost200Response](docs/GetUserChatBoostsPost200Response.md)
 - [GetUserChatBoostsPostRequest](docs/GetUserChatBoostsPostRequest.md)
 - [GetUserChatBoostsPostRequestChatId](docs/GetUserChatBoostsPostRequestChatId.md)
 - [GetUserProfilePhotosPost200Response](docs/GetUserProfilePhotosPost200Response.md)
 - [GetUserProfilePhotosPostRequest](docs/GetUserProfilePhotosPostRequest.md)
 - [GetWebhookInfoPost200Response](docs/GetWebhookInfoPost200Response.md)
 - [Gift](docs/Gift.md)
 - [GiftInfo](docs/GiftInfo.md)
 - [GiftPremiumSubscriptionPostRequest](docs/GiftPremiumSubscriptionPostRequest.md)
 - [Gifts](docs/Gifts.md)
 - [Giveaway](docs/Giveaway.md)
 - [GiveawayCompleted](docs/GiveawayCompleted.md)
 - [GiveawayCreated](docs/GiveawayCreated.md)
 - [GiveawayWinners](docs/GiveawayWinners.md)
 - [InaccessibleMessage](docs/InaccessibleMessage.md)
 - [InlineKeyboardButton](docs/InlineKeyboardButton.md)
 - [InlineKeyboardMarkup](docs/InlineKeyboardMarkup.md)
 - [InlineQuery](docs/InlineQuery.md)
 - [InlineQueryResult](docs/InlineQueryResult.md)
 - [InlineQueryResultArticle](docs/InlineQueryResultArticle.md)
 - [InlineQueryResultAudio](docs/InlineQueryResultAudio.md)
 - [InlineQueryResultCachedAudio](docs/InlineQueryResultCachedAudio.md)
 - [InlineQueryResultCachedDocument](docs/InlineQueryResultCachedDocument.md)
 - [InlineQueryResultCachedGif](docs/InlineQueryResultCachedGif.md)
 - [InlineQueryResultCachedMpeg4Gif](docs/InlineQueryResultCachedMpeg4Gif.md)
 - [InlineQueryResultCachedPhoto](docs/InlineQueryResultCachedPhoto.md)
 - [InlineQueryResultCachedSticker](docs/InlineQueryResultCachedSticker.md)
 - [InlineQueryResultCachedVideo](docs/InlineQueryResultCachedVideo.md)
 - [InlineQueryResultCachedVoice](docs/InlineQueryResultCachedVoice.md)
 - [InlineQueryResultContact](docs/InlineQueryResultContact.md)
 - [InlineQueryResultDocument](docs/InlineQueryResultDocument.md)
 - [InlineQueryResultGame](docs/InlineQueryResultGame.md)
 - [InlineQueryResultGif](docs/InlineQueryResultGif.md)
 - [InlineQueryResultLocation](docs/InlineQueryResultLocation.md)
 - [InlineQueryResultMpeg4Gif](docs/InlineQueryResultMpeg4Gif.md)
 - [InlineQueryResultPhoto](docs/InlineQueryResultPhoto.md)
 - [InlineQueryResultVenue](docs/InlineQueryResultVenue.md)
 - [InlineQueryResultVideo](docs/InlineQueryResultVideo.md)
 - [InlineQueryResultVoice](docs/InlineQueryResultVoice.md)
 - [InlineQueryResultsButton](docs/InlineQueryResultsButton.md)
 - [InputContactMessageContent](docs/InputContactMessageContent.md)
 - [InputInvoiceMessageContent](docs/InputInvoiceMessageContent.md)
 - [InputLocationMessageContent](docs/InputLocationMessageContent.md)
 - [InputMedia](docs/InputMedia.md)
 - [InputMediaAnimation](docs/InputMediaAnimation.md)
 - [InputMediaAudio](docs/InputMediaAudio.md)
 - [InputMediaDocument](docs/InputMediaDocument.md)
 - [InputMediaPhoto](docs/InputMediaPhoto.md)
 - [InputMediaVideo](docs/InputMediaVideo.md)
 - [InputMessageContent](docs/InputMessageContent.md)
 - [InputPaidMedia](docs/InputPaidMedia.md)
 - [InputPaidMediaPhoto](docs/InputPaidMediaPhoto.md)
 - [InputPaidMediaVideo](docs/InputPaidMediaVideo.md)
 - [InputPollOption](docs/InputPollOption.md)
 - [InputProfilePhoto](docs/InputProfilePhoto.md)
 - [InputProfilePhotoAnimated](docs/InputProfilePhotoAnimated.md)
 - [InputProfilePhotoStatic](docs/InputProfilePhotoStatic.md)
 - [InputSticker](docs/InputSticker.md)
 - [InputStoryContent](docs/InputStoryContent.md)
 - [InputStoryContentPhoto](docs/InputStoryContentPhoto.md)
 - [InputStoryContentVideo](docs/InputStoryContentVideo.md)
 - [InputTextMessageContent](docs/InputTextMessageContent.md)
 - [InputVenueMessageContent](docs/InputVenueMessageContent.md)
 - [Invoice](docs/Invoice.md)
 - [KeyboardButton](docs/KeyboardButton.md)
 - [KeyboardButtonPollType](docs/KeyboardButtonPollType.md)
 - [KeyboardButtonRequestChat](docs/KeyboardButtonRequestChat.md)
 - [KeyboardButtonRequestUsers](docs/KeyboardButtonRequestUsers.md)
 - [LabeledPrice](docs/LabeledPrice.md)
 - [LeaveChatPostRequest](docs/LeaveChatPostRequest.md)
 - [LeaveChatPostRequestChatId](docs/LeaveChatPostRequestChatId.md)
 - [LinkPreviewOptions](docs/LinkPreviewOptions.md)
 - [Location](docs/Location.md)
 - [LocationAddress](docs/LocationAddress.md)
 - [LoginUrl](docs/LoginUrl.md)
 - [MaskPosition](docs/MaskPosition.md)
 - [MaybeInaccessibleMessage](docs/MaybeInaccessibleMessage.md)
 - [MenuButton](docs/MenuButton.md)
 - [MenuButtonCommands](docs/MenuButtonCommands.md)
 - [MenuButtonDefault](docs/MenuButtonDefault.md)
 - [MenuButtonWebApp](docs/MenuButtonWebApp.md)
 - [Message](docs/Message.md)
 - [MessageAutoDeleteTimerChanged](docs/MessageAutoDeleteTimerChanged.md)
 - [MessageEntity](docs/MessageEntity.md)
 - [MessageId](docs/MessageId.md)
 - [MessageOrigin](docs/MessageOrigin.md)
 - [MessageOriginChannel](docs/MessageOriginChannel.md)
 - [MessageOriginChat](docs/MessageOriginChat.md)
 - [MessageOriginHiddenUser](docs/MessageOriginHiddenUser.md)
 - [MessageOriginUser](docs/MessageOriginUser.md)
 - [MessageReactionCountUpdated](docs/MessageReactionCountUpdated.md)
 - [MessageReactionUpdated](docs/MessageReactionUpdated.md)
 - [OrderInfo](docs/OrderInfo.md)
 - [OwnedGift](docs/OwnedGift.md)
 - [OwnedGiftRegular](docs/OwnedGiftRegular.md)
 - [OwnedGiftUnique](docs/OwnedGiftUnique.md)
 - [OwnedGifts](docs/OwnedGifts.md)
 - [PaidMedia](docs/PaidMedia.md)
 - [PaidMediaInfo](docs/PaidMediaInfo.md)
 - [PaidMediaPhoto](docs/PaidMediaPhoto.md)
 - [PaidMediaPreview](docs/PaidMediaPreview.md)
 - [PaidMediaPurchased](docs/PaidMediaPurchased.md)
 - [PaidMediaVideo](docs/PaidMediaVideo.md)
 - [PaidMessagePriceChanged](docs/PaidMessagePriceChanged.md)
 - [PassportData](docs/PassportData.md)
 - [PassportElementError](docs/PassportElementError.md)
 - [PassportElementErrorDataField](docs/PassportElementErrorDataField.md)
 - [PassportElementErrorFile](docs/PassportElementErrorFile.md)
 - [PassportElementErrorFiles](docs/PassportElementErrorFiles.md)
 - [PassportElementErrorFrontSide](docs/PassportElementErrorFrontSide.md)
 - [PassportElementErrorReverseSide](docs/PassportElementErrorReverseSide.md)
 - [PassportElementErrorSelfie](docs/PassportElementErrorSelfie.md)
 - [PassportElementErrorTranslationFile](docs/PassportElementErrorTranslationFile.md)
 - [PassportElementErrorTranslationFiles](docs/PassportElementErrorTranslationFiles.md)
 - [PassportElementErrorUnspecified](docs/PassportElementErrorUnspecified.md)
 - [PassportFile](docs/PassportFile.md)
 - [PhotoSize](docs/PhotoSize.md)
 - [PinChatMessagePostRequest](docs/PinChatMessagePostRequest.md)
 - [Poll](docs/Poll.md)
 - [PollAnswer](docs/PollAnswer.md)
 - [PollOption](docs/PollOption.md)
 - [PostStoryPost200Response](docs/PostStoryPost200Response.md)
 - [PreCheckoutQuery](docs/PreCheckoutQuery.md)
 - [PreparedInlineMessage](docs/PreparedInlineMessage.md)
 - [PromoteChatMemberPostRequest](docs/PromoteChatMemberPostRequest.md)
 - [ProximityAlertTriggered](docs/ProximityAlertTriggered.md)
 - [ReactionCount](docs/ReactionCount.md)
 - [ReactionType](docs/ReactionType.md)
 - [ReactionTypeCustomEmoji](docs/ReactionTypeCustomEmoji.md)
 - [ReactionTypeEmoji](docs/ReactionTypeEmoji.md)
 - [ReactionTypePaid](docs/ReactionTypePaid.md)
 - [ReadBusinessMessagePostRequest](docs/ReadBusinessMessagePostRequest.md)
 - [RefundStarPaymentPostRequest](docs/RefundStarPaymentPostRequest.md)
 - [RefundedPayment](docs/RefundedPayment.md)
 - [RemoveBusinessAccountProfilePhotoPostRequest](docs/RemoveBusinessAccountProfilePhotoPostRequest.md)
 - [RemoveUserVerificationPostRequest](docs/RemoveUserVerificationPostRequest.md)
 - [ReplyKeyboardMarkup](docs/ReplyKeyboardMarkup.md)
 - [ReplyKeyboardRemove](docs/ReplyKeyboardRemove.md)
 - [ReplyParameters](docs/ReplyParameters.md)
 - [ReplyParametersChatId](docs/ReplyParametersChatId.md)
 - [ResponseParameters](docs/ResponseParameters.md)
 - [RestrictChatMemberPostRequest](docs/RestrictChatMemberPostRequest.md)
 - [RestrictChatMemberPostRequestChatId](docs/RestrictChatMemberPostRequestChatId.md)
 - [RevenueWithdrawalState](docs/RevenueWithdrawalState.md)
 - [RevenueWithdrawalStateFailed](docs/RevenueWithdrawalStateFailed.md)
 - [RevenueWithdrawalStatePending](docs/RevenueWithdrawalStatePending.md)
 - [RevenueWithdrawalStateSucceeded](docs/RevenueWithdrawalStateSucceeded.md)
 - [RevokeChatInviteLinkPostRequest](docs/RevokeChatInviteLinkPostRequest.md)
 - [RevokeChatInviteLinkPostRequestChatId](docs/RevokeChatInviteLinkPostRequestChatId.md)
 - [SavePreparedInlineMessagePost200Response](docs/SavePreparedInlineMessagePost200Response.md)
 - [SavePreparedInlineMessagePostRequest](docs/SavePreparedInlineMessagePostRequest.md)
 - [SendAnimationPostRequestAnimation](docs/SendAnimationPostRequestAnimation.md)
 - [SendAudioPostRequestAudio](docs/SendAudioPostRequestAudio.md)
 - [SendAudioPostRequestThumbnail](docs/SendAudioPostRequestThumbnail.md)
 - [SendChatActionPostRequest](docs/SendChatActionPostRequest.md)
 - [SendContactPostRequest](docs/SendContactPostRequest.md)
 - [SendDicePostRequest](docs/SendDicePostRequest.md)
 - [SendDocumentPostRequestDocument](docs/SendDocumentPostRequestDocument.md)
 - [SendGamePostRequest](docs/SendGamePostRequest.md)
 - [SendGiftPostRequest](docs/SendGiftPostRequest.md)
 - [SendGiftPostRequestChatId](docs/SendGiftPostRequestChatId.md)
 - [SendInvoicePostRequest](docs/SendInvoicePostRequest.md)
 - [SendLocationPostRequest](docs/SendLocationPostRequest.md)
 - [SendMediaGroupPost200Response](docs/SendMediaGroupPost200Response.md)
 - [SendMediaGroupPostRequestMediaInner](docs/SendMediaGroupPostRequestMediaInner.md)
 - [SendMessagePost200Response](docs/SendMessagePost200Response.md)
 - [SendMessagePostRequest](docs/SendMessagePostRequest.md)
 - [SendMessagePostRequestChatId](docs/SendMessagePostRequestChatId.md)
 - [SendMessagePostRequestReplyMarkup](docs/SendMessagePostRequestReplyMarkup.md)
 - [SendPaidMediaPostRequestChatId](docs/SendPaidMediaPostRequestChatId.md)
 - [SendPhotoPostRequestPhoto](docs/SendPhotoPostRequestPhoto.md)
 - [SendPollPostRequest](docs/SendPollPostRequest.md)
 - [SendStickerPostRequestSticker](docs/SendStickerPostRequestSticker.md)
 - [SendVenuePostRequest](docs/SendVenuePostRequest.md)
 - [SendVideoNotePostRequestVideoNote](docs/SendVideoNotePostRequestVideoNote.md)
 - [SendVideoPostRequestCover](docs/SendVideoPostRequestCover.md)
 - [SendVideoPostRequestVideo](docs/SendVideoPostRequestVideo.md)
 - [SendVoicePostRequestVoice](docs/SendVoicePostRequestVoice.md)
 - [SentWebAppMessage](docs/SentWebAppMessage.md)
 - [SetBusinessAccountBioPostRequest](docs/SetBusinessAccountBioPostRequest.md)
 - [SetBusinessAccountGiftSettingsPostRequest](docs/SetBusinessAccountGiftSettingsPostRequest.md)
 - [SetBusinessAccountNamePostRequest](docs/SetBusinessAccountNamePostRequest.md)
 - [SetBusinessAccountUsernamePostRequest](docs/SetBusinessAccountUsernamePostRequest.md)
 - [SetChatAdministratorCustomTitlePostRequest](docs/SetChatAdministratorCustomTitlePostRequest.md)
 - [SetChatDescriptionPostRequest](docs/SetChatDescriptionPostRequest.md)
 - [SetChatMenuButtonPostRequest](docs/SetChatMenuButtonPostRequest.md)
 - [SetChatPermissionsPostRequest](docs/SetChatPermissionsPostRequest.md)
 - [SetChatStickerSetPostRequest](docs/SetChatStickerSetPostRequest.md)
 - [SetChatTitlePostRequest](docs/SetChatTitlePostRequest.md)
 - [SetCustomEmojiStickerSetThumbnailPostRequest](docs/SetCustomEmojiStickerSetThumbnailPostRequest.md)
 - [SetGameScorePostRequest](docs/SetGameScorePostRequest.md)
 - [SetMessageReactionPostRequest](docs/SetMessageReactionPostRequest.md)
 - [SetMyCommandsPostRequest](docs/SetMyCommandsPostRequest.md)
 - [SetMyDefaultAdministratorRightsPostRequest](docs/SetMyDefaultAdministratorRightsPostRequest.md)
 - [SetMyDescriptionPostRequest](docs/SetMyDescriptionPostRequest.md)
 - [SetMyNamePostRequest](docs/SetMyNamePostRequest.md)
 - [SetMyShortDescriptionPostRequest](docs/SetMyShortDescriptionPostRequest.md)
 - [SetPassportDataErrorsPostRequest](docs/SetPassportDataErrorsPostRequest.md)
 - [SetStickerEmojiListPostRequest](docs/SetStickerEmojiListPostRequest.md)
 - [SetStickerKeywordsPostRequest](docs/SetStickerKeywordsPostRequest.md)
 - [SetStickerMaskPositionPostRequest](docs/SetStickerMaskPositionPostRequest.md)
 - [SetStickerPositionInSetPostRequest](docs/SetStickerPositionInSetPostRequest.md)
 - [SetStickerSetThumbnailPostRequestThumbnail](docs/SetStickerSetThumbnailPostRequestThumbnail.md)
 - [SetStickerSetTitlePostRequest](docs/SetStickerSetTitlePostRequest.md)
 - [SetUserEmojiStatusPostRequest](docs/SetUserEmojiStatusPostRequest.md)
 - [SetWebhookPost200Response](docs/SetWebhookPost200Response.md)
 - [SharedUser](docs/SharedUser.md)
 - [ShippingAddress](docs/ShippingAddress.md)
 - [ShippingOption](docs/ShippingOption.md)
 - [ShippingQuery](docs/ShippingQuery.md)
 - [StarAmount](docs/StarAmount.md)
 - [StarTransaction](docs/StarTransaction.md)
 - [StarTransactions](docs/StarTransactions.md)
 - [Sticker](docs/Sticker.md)
 - [StickerSet](docs/StickerSet.md)
 - [StopMessageLiveLocationPostRequest](docs/StopMessageLiveLocationPostRequest.md)
 - [StopPollPost200Response](docs/StopPollPost200Response.md)
 - [StopPollPostRequest](docs/StopPollPostRequest.md)
 - [Story](docs/Story.md)
 - [StoryArea](docs/StoryArea.md)
 - [StoryAreaPosition](docs/StoryAreaPosition.md)
 - [StoryAreaType](docs/StoryAreaType.md)
 - [StoryAreaTypeLink](docs/StoryAreaTypeLink.md)
 - [StoryAreaTypeLocation](docs/StoryAreaTypeLocation.md)
 - [StoryAreaTypeSuggestedReaction](docs/StoryAreaTypeSuggestedReaction.md)
 - [StoryAreaTypeUniqueGift](docs/StoryAreaTypeUniqueGift.md)
 - [StoryAreaTypeWeather](docs/StoryAreaTypeWeather.md)
 - [SuccessfulPayment](docs/SuccessfulPayment.md)
 - [SwitchInlineQueryChosenChat](docs/SwitchInlineQueryChosenChat.md)
 - [TextQuote](docs/TextQuote.md)
 - [TransactionPartner](docs/TransactionPartner.md)
 - [TransactionPartnerAffiliateProgram](docs/TransactionPartnerAffiliateProgram.md)
 - [TransactionPartnerChat](docs/TransactionPartnerChat.md)
 - [TransactionPartnerFragment](docs/TransactionPartnerFragment.md)
 - [TransactionPartnerOther](docs/TransactionPartnerOther.md)
 - [TransactionPartnerTelegramAds](docs/TransactionPartnerTelegramAds.md)
 - [TransactionPartnerTelegramApi](docs/TransactionPartnerTelegramApi.md)
 - [TransactionPartnerUser](docs/TransactionPartnerUser.md)
 - [TransferBusinessAccountStarsPostRequest](docs/TransferBusinessAccountStarsPostRequest.md)
 - [TransferGiftPostRequest](docs/TransferGiftPostRequest.md)
 - [UnbanChatMemberPostRequest](docs/UnbanChatMemberPostRequest.md)
 - [UniqueGift](docs/UniqueGift.md)
 - [UniqueGiftBackdrop](docs/UniqueGiftBackdrop.md)
 - [UniqueGiftBackdropColors](docs/UniqueGiftBackdropColors.md)
 - [UniqueGiftInfo](docs/UniqueGiftInfo.md)
 - [UniqueGiftModel](docs/UniqueGiftModel.md)
 - [UniqueGiftSymbol](docs/UniqueGiftSymbol.md)
 - [UnpinChatMessagePostRequest](docs/UnpinChatMessagePostRequest.md)
 - [Update](docs/Update.md)
 - [UpgradeGiftPostRequest](docs/UpgradeGiftPostRequest.md)
 - [User](docs/User.md)
 - [UserChatBoosts](docs/UserChatBoosts.md)
 - [UserProfilePhotos](docs/UserProfilePhotos.md)
 - [UsersShared](docs/UsersShared.md)
 - [Venue](docs/Venue.md)
 - [VerifyChatPostRequest](docs/VerifyChatPostRequest.md)
 - [VerifyUserPostRequest](docs/VerifyUserPostRequest.md)
 - [Video](docs/Video.md)
 - [VideoChatEnded](docs/VideoChatEnded.md)
 - [VideoChatParticipantsInvited](docs/VideoChatParticipantsInvited.md)
 - [VideoChatScheduled](docs/VideoChatScheduled.md)
 - [VideoNote](docs/VideoNote.md)
 - [Voice](docs/Voice.md)
 - [WebAppData](docs/WebAppData.md)
 - [WebAppInfo](docs/WebAppInfo.md)
 - [WebhookInfo](docs/WebhookInfo.md)
 - [WriteAccessAllowed](docs/WriteAccessAllowed.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



