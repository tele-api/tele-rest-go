# Go API client for tele_rest

The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram.
To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.

## Overview

- API version: 9.1.0
- Package version: 9.1.0
- Build date: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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
*DefaultAPI* | [**PostAddStickerToSet**](docs/DefaultAPI.md#postaddstickertoset) | **Post** /addStickerToSet | addStickerToSet
*DefaultAPI* | [**PostAnswerCallbackQuery**](docs/DefaultAPI.md#postanswercallbackquery) | **Post** /answerCallbackQuery | answerCallbackQuery
*DefaultAPI* | [**PostAnswerInlineQuery**](docs/DefaultAPI.md#postanswerinlinequery) | **Post** /answerInlineQuery | answerInlineQuery
*DefaultAPI* | [**PostAnswerPreCheckoutQuery**](docs/DefaultAPI.md#postanswerprecheckoutquery) | **Post** /answerPreCheckoutQuery | answerPreCheckoutQuery
*DefaultAPI* | [**PostAnswerShippingQuery**](docs/DefaultAPI.md#postanswershippingquery) | **Post** /answerShippingQuery | answerShippingQuery
*DefaultAPI* | [**PostAnswerWebAppQuery**](docs/DefaultAPI.md#postanswerwebappquery) | **Post** /answerWebAppQuery | answerWebAppQuery
*DefaultAPI* | [**PostApproveChatJoinRequest**](docs/DefaultAPI.md#postapprovechatjoinrequest) | **Post** /approveChatJoinRequest | approveChatJoinRequest
*DefaultAPI* | [**PostBanChatMember**](docs/DefaultAPI.md#postbanchatmember) | **Post** /banChatMember | banChatMember
*DefaultAPI* | [**PostBanChatSenderChat**](docs/DefaultAPI.md#postbanchatsenderchat) | **Post** /banChatSenderChat | banChatSenderChat
*DefaultAPI* | [**PostClose**](docs/DefaultAPI.md#postclose) | **Post** /close | close
*DefaultAPI* | [**PostCloseForumTopic**](docs/DefaultAPI.md#postcloseforumtopic) | **Post** /closeForumTopic | closeForumTopic
*DefaultAPI* | [**PostCloseGeneralForumTopic**](docs/DefaultAPI.md#postclosegeneralforumtopic) | **Post** /closeGeneralForumTopic | closeGeneralForumTopic
*DefaultAPI* | [**PostConvertGiftToStars**](docs/DefaultAPI.md#postconvertgifttostars) | **Post** /convertGiftToStars | convertGiftToStars
*DefaultAPI* | [**PostCopyMessage**](docs/DefaultAPI.md#postcopymessage) | **Post** /copyMessage | copyMessage
*DefaultAPI* | [**PostCopyMessages**](docs/DefaultAPI.md#postcopymessages) | **Post** /copyMessages | copyMessages
*DefaultAPI* | [**PostCreateChatInviteLink**](docs/DefaultAPI.md#postcreatechatinvitelink) | **Post** /createChatInviteLink | createChatInviteLink
*DefaultAPI* | [**PostCreateChatSubscriptionInviteLink**](docs/DefaultAPI.md#postcreatechatsubscriptioninvitelink) | **Post** /createChatSubscriptionInviteLink | createChatSubscriptionInviteLink
*DefaultAPI* | [**PostCreateForumTopic**](docs/DefaultAPI.md#postcreateforumtopic) | **Post** /createForumTopic | createForumTopic
*DefaultAPI* | [**PostCreateInvoiceLink**](docs/DefaultAPI.md#postcreateinvoicelink) | **Post** /createInvoiceLink | createInvoiceLink
*DefaultAPI* | [**PostCreateNewStickerSet**](docs/DefaultAPI.md#postcreatenewstickerset) | **Post** /createNewStickerSet | createNewStickerSet
*DefaultAPI* | [**PostDeclineChatJoinRequest**](docs/DefaultAPI.md#postdeclinechatjoinrequest) | **Post** /declineChatJoinRequest | declineChatJoinRequest
*DefaultAPI* | [**PostDeleteBusinessMessages**](docs/DefaultAPI.md#postdeletebusinessmessages) | **Post** /deleteBusinessMessages | deleteBusinessMessages
*DefaultAPI* | [**PostDeleteChatPhoto**](docs/DefaultAPI.md#postdeletechatphoto) | **Post** /deleteChatPhoto | deleteChatPhoto
*DefaultAPI* | [**PostDeleteChatStickerSet**](docs/DefaultAPI.md#postdeletechatstickerset) | **Post** /deleteChatStickerSet | deleteChatStickerSet
*DefaultAPI* | [**PostDeleteForumTopic**](docs/DefaultAPI.md#postdeleteforumtopic) | **Post** /deleteForumTopic | deleteForumTopic
*DefaultAPI* | [**PostDeleteMessage**](docs/DefaultAPI.md#postdeletemessage) | **Post** /deleteMessage | deleteMessage
*DefaultAPI* | [**PostDeleteMessages**](docs/DefaultAPI.md#postdeletemessages) | **Post** /deleteMessages | deleteMessages
*DefaultAPI* | [**PostDeleteMyCommands**](docs/DefaultAPI.md#postdeletemycommands) | **Post** /deleteMyCommands | deleteMyCommands
*DefaultAPI* | [**PostDeleteStickerFromSet**](docs/DefaultAPI.md#postdeletestickerfromset) | **Post** /deleteStickerFromSet | deleteStickerFromSet
*DefaultAPI* | [**PostDeleteStickerSet**](docs/DefaultAPI.md#postdeletestickerset) | **Post** /deleteStickerSet | deleteStickerSet
*DefaultAPI* | [**PostDeleteStory**](docs/DefaultAPI.md#postdeletestory) | **Post** /deleteStory | deleteStory
*DefaultAPI* | [**PostDeleteWebhook**](docs/DefaultAPI.md#postdeletewebhook) | **Post** /deleteWebhook | deleteWebhook
*DefaultAPI* | [**PostEditChatInviteLink**](docs/DefaultAPI.md#posteditchatinvitelink) | **Post** /editChatInviteLink | editChatInviteLink
*DefaultAPI* | [**PostEditChatSubscriptionInviteLink**](docs/DefaultAPI.md#posteditchatsubscriptioninvitelink) | **Post** /editChatSubscriptionInviteLink | editChatSubscriptionInviteLink
*DefaultAPI* | [**PostEditForumTopic**](docs/DefaultAPI.md#posteditforumtopic) | **Post** /editForumTopic | editForumTopic
*DefaultAPI* | [**PostEditGeneralForumTopic**](docs/DefaultAPI.md#posteditgeneralforumtopic) | **Post** /editGeneralForumTopic | editGeneralForumTopic
*DefaultAPI* | [**PostEditMessageCaption**](docs/DefaultAPI.md#posteditmessagecaption) | **Post** /editMessageCaption | editMessageCaption
*DefaultAPI* | [**PostEditMessageChecklist**](docs/DefaultAPI.md#posteditmessagechecklist) | **Post** /editMessageChecklist | editMessageChecklist
*DefaultAPI* | [**PostEditMessageLiveLocation**](docs/DefaultAPI.md#posteditmessagelivelocation) | **Post** /editMessageLiveLocation | editMessageLiveLocation
*DefaultAPI* | [**PostEditMessageMedia**](docs/DefaultAPI.md#posteditmessagemedia) | **Post** /editMessageMedia | editMessageMedia
*DefaultAPI* | [**PostEditMessageReplyMarkup**](docs/DefaultAPI.md#posteditmessagereplymarkup) | **Post** /editMessageReplyMarkup | editMessageReplyMarkup
*DefaultAPI* | [**PostEditMessageText**](docs/DefaultAPI.md#posteditmessagetext) | **Post** /editMessageText | editMessageText
*DefaultAPI* | [**PostEditStory**](docs/DefaultAPI.md#posteditstory) | **Post** /editStory | editStory
*DefaultAPI* | [**PostEditUserStarSubscription**](docs/DefaultAPI.md#postedituserstarsubscription) | **Post** /editUserStarSubscription | editUserStarSubscription
*DefaultAPI* | [**PostExportChatInviteLink**](docs/DefaultAPI.md#postexportchatinvitelink) | **Post** /exportChatInviteLink | exportChatInviteLink
*DefaultAPI* | [**PostForwardMessage**](docs/DefaultAPI.md#postforwardmessage) | **Post** /forwardMessage | forwardMessage
*DefaultAPI* | [**PostForwardMessages**](docs/DefaultAPI.md#postforwardmessages) | **Post** /forwardMessages | forwardMessages
*DefaultAPI* | [**PostGetAvailableGifts**](docs/DefaultAPI.md#postgetavailablegifts) | **Post** /getAvailableGifts | getAvailableGifts
*DefaultAPI* | [**PostGetBusinessAccountGifts**](docs/DefaultAPI.md#postgetbusinessaccountgifts) | **Post** /getBusinessAccountGifts | getBusinessAccountGifts
*DefaultAPI* | [**PostGetBusinessAccountStarBalance**](docs/DefaultAPI.md#postgetbusinessaccountstarbalance) | **Post** /getBusinessAccountStarBalance | getBusinessAccountStarBalance
*DefaultAPI* | [**PostGetBusinessConnection**](docs/DefaultAPI.md#postgetbusinessconnection) | **Post** /getBusinessConnection | getBusinessConnection
*DefaultAPI* | [**PostGetChat**](docs/DefaultAPI.md#postgetchat) | **Post** /getChat | getChat
*DefaultAPI* | [**PostGetChatAdministrators**](docs/DefaultAPI.md#postgetchatadministrators) | **Post** /getChatAdministrators | getChatAdministrators
*DefaultAPI* | [**PostGetChatMember**](docs/DefaultAPI.md#postgetchatmember) | **Post** /getChatMember | getChatMember
*DefaultAPI* | [**PostGetChatMemberCount**](docs/DefaultAPI.md#postgetchatmembercount) | **Post** /getChatMemberCount | getChatMemberCount
*DefaultAPI* | [**PostGetChatMenuButton**](docs/DefaultAPI.md#postgetchatmenubutton) | **Post** /getChatMenuButton | getChatMenuButton
*DefaultAPI* | [**PostGetCustomEmojiStickers**](docs/DefaultAPI.md#postgetcustomemojistickers) | **Post** /getCustomEmojiStickers | getCustomEmojiStickers
*DefaultAPI* | [**PostGetFile**](docs/DefaultAPI.md#postgetfile) | **Post** /getFile | getFile
*DefaultAPI* | [**PostGetForumTopicIconStickers**](docs/DefaultAPI.md#postgetforumtopiciconstickers) | **Post** /getForumTopicIconStickers | getForumTopicIconStickers
*DefaultAPI* | [**PostGetGameHighScores**](docs/DefaultAPI.md#postgetgamehighscores) | **Post** /getGameHighScores | getGameHighScores
*DefaultAPI* | [**PostGetMe**](docs/DefaultAPI.md#postgetme) | **Post** /getMe | getMe
*DefaultAPI* | [**PostGetMyCommands**](docs/DefaultAPI.md#postgetmycommands) | **Post** /getMyCommands | getMyCommands
*DefaultAPI* | [**PostGetMyDefaultAdministratorRights**](docs/DefaultAPI.md#postgetmydefaultadministratorrights) | **Post** /getMyDefaultAdministratorRights | getMyDefaultAdministratorRights
*DefaultAPI* | [**PostGetMyDescription**](docs/DefaultAPI.md#postgetmydescription) | **Post** /getMyDescription | getMyDescription
*DefaultAPI* | [**PostGetMyName**](docs/DefaultAPI.md#postgetmyname) | **Post** /getMyName | getMyName
*DefaultAPI* | [**PostGetMyShortDescription**](docs/DefaultAPI.md#postgetmyshortdescription) | **Post** /getMyShortDescription | getMyShortDescription
*DefaultAPI* | [**PostGetMyStarBalance**](docs/DefaultAPI.md#postgetmystarbalance) | **Post** /getMyStarBalance | getMyStarBalance
*DefaultAPI* | [**PostGetStarTransactions**](docs/DefaultAPI.md#postgetstartransactions) | **Post** /getStarTransactions | getStarTransactions
*DefaultAPI* | [**PostGetStickerSet**](docs/DefaultAPI.md#postgetstickerset) | **Post** /getStickerSet | getStickerSet
*DefaultAPI* | [**PostGetUpdates**](docs/DefaultAPI.md#postgetupdates) | **Post** /getUpdates | getUpdates
*DefaultAPI* | [**PostGetUserChatBoosts**](docs/DefaultAPI.md#postgetuserchatboosts) | **Post** /getUserChatBoosts | getUserChatBoosts
*DefaultAPI* | [**PostGetUserProfilePhotos**](docs/DefaultAPI.md#postgetuserprofilephotos) | **Post** /getUserProfilePhotos | getUserProfilePhotos
*DefaultAPI* | [**PostGetWebhookInfo**](docs/DefaultAPI.md#postgetwebhookinfo) | **Post** /getWebhookInfo | getWebhookInfo
*DefaultAPI* | [**PostGiftPremiumSubscription**](docs/DefaultAPI.md#postgiftpremiumsubscription) | **Post** /giftPremiumSubscription | giftPremiumSubscription
*DefaultAPI* | [**PostHideGeneralForumTopic**](docs/DefaultAPI.md#posthidegeneralforumtopic) | **Post** /hideGeneralForumTopic | hideGeneralForumTopic
*DefaultAPI* | [**PostLeaveChat**](docs/DefaultAPI.md#postleavechat) | **Post** /leaveChat | leaveChat
*DefaultAPI* | [**PostLogOut**](docs/DefaultAPI.md#postlogout) | **Post** /logOut | logOut
*DefaultAPI* | [**PostPinChatMessage**](docs/DefaultAPI.md#postpinchatmessage) | **Post** /pinChatMessage | pinChatMessage
*DefaultAPI* | [**PostPostStory**](docs/DefaultAPI.md#postpoststory) | **Post** /postStory | postStory
*DefaultAPI* | [**PostPromoteChatMember**](docs/DefaultAPI.md#postpromotechatmember) | **Post** /promoteChatMember | promoteChatMember
*DefaultAPI* | [**PostReadBusinessMessage**](docs/DefaultAPI.md#postreadbusinessmessage) | **Post** /readBusinessMessage | readBusinessMessage
*DefaultAPI* | [**PostRefundStarPayment**](docs/DefaultAPI.md#postrefundstarpayment) | **Post** /refundStarPayment | refundStarPayment
*DefaultAPI* | [**PostRemoveBusinessAccountProfilePhoto**](docs/DefaultAPI.md#postremovebusinessaccountprofilephoto) | **Post** /removeBusinessAccountProfilePhoto | removeBusinessAccountProfilePhoto
*DefaultAPI* | [**PostRemoveChatVerification**](docs/DefaultAPI.md#postremovechatverification) | **Post** /removeChatVerification | removeChatVerification
*DefaultAPI* | [**PostRemoveUserVerification**](docs/DefaultAPI.md#postremoveuserverification) | **Post** /removeUserVerification | removeUserVerification
*DefaultAPI* | [**PostReopenForumTopic**](docs/DefaultAPI.md#postreopenforumtopic) | **Post** /reopenForumTopic | reopenForumTopic
*DefaultAPI* | [**PostReopenGeneralForumTopic**](docs/DefaultAPI.md#postreopengeneralforumtopic) | **Post** /reopenGeneralForumTopic | reopenGeneralForumTopic
*DefaultAPI* | [**PostReplaceStickerInSet**](docs/DefaultAPI.md#postreplacestickerinset) | **Post** /replaceStickerInSet | replaceStickerInSet
*DefaultAPI* | [**PostRestrictChatMember**](docs/DefaultAPI.md#postrestrictchatmember) | **Post** /restrictChatMember | restrictChatMember
*DefaultAPI* | [**PostRevokeChatInviteLink**](docs/DefaultAPI.md#postrevokechatinvitelink) | **Post** /revokeChatInviteLink | revokeChatInviteLink
*DefaultAPI* | [**PostSavePreparedInlineMessage**](docs/DefaultAPI.md#postsavepreparedinlinemessage) | **Post** /savePreparedInlineMessage | savePreparedInlineMessage
*DefaultAPI* | [**PostSendAnimation**](docs/DefaultAPI.md#postsendanimation) | **Post** /sendAnimation | sendAnimation
*DefaultAPI* | [**PostSendAudio**](docs/DefaultAPI.md#postsendaudio) | **Post** /sendAudio | sendAudio
*DefaultAPI* | [**PostSendChatAction**](docs/DefaultAPI.md#postsendchataction) | **Post** /sendChatAction | sendChatAction
*DefaultAPI* | [**PostSendChecklist**](docs/DefaultAPI.md#postsendchecklist) | **Post** /sendChecklist | sendChecklist
*DefaultAPI* | [**PostSendContact**](docs/DefaultAPI.md#postsendcontact) | **Post** /sendContact | sendContact
*DefaultAPI* | [**PostSendDice**](docs/DefaultAPI.md#postsenddice) | **Post** /sendDice | sendDice
*DefaultAPI* | [**PostSendDocument**](docs/DefaultAPI.md#postsenddocument) | **Post** /sendDocument | sendDocument
*DefaultAPI* | [**PostSendGame**](docs/DefaultAPI.md#postsendgame) | **Post** /sendGame | sendGame
*DefaultAPI* | [**PostSendGift**](docs/DefaultAPI.md#postsendgift) | **Post** /sendGift | sendGift
*DefaultAPI* | [**PostSendInvoice**](docs/DefaultAPI.md#postsendinvoice) | **Post** /sendInvoice | sendInvoice
*DefaultAPI* | [**PostSendLocation**](docs/DefaultAPI.md#postsendlocation) | **Post** /sendLocation | sendLocation
*DefaultAPI* | [**PostSendMediaGroup**](docs/DefaultAPI.md#postsendmediagroup) | **Post** /sendMediaGroup | sendMediaGroup
*DefaultAPI* | [**PostSendMessage**](docs/DefaultAPI.md#postsendmessage) | **Post** /sendMessage | sendMessage
*DefaultAPI* | [**PostSendPaidMedia**](docs/DefaultAPI.md#postsendpaidmedia) | **Post** /sendPaidMedia | sendPaidMedia
*DefaultAPI* | [**PostSendPhoto**](docs/DefaultAPI.md#postsendphoto) | **Post** /sendPhoto | sendPhoto
*DefaultAPI* | [**PostSendPoll**](docs/DefaultAPI.md#postsendpoll) | **Post** /sendPoll | sendPoll
*DefaultAPI* | [**PostSendSticker**](docs/DefaultAPI.md#postsendsticker) | **Post** /sendSticker | sendSticker
*DefaultAPI* | [**PostSendVenue**](docs/DefaultAPI.md#postsendvenue) | **Post** /sendVenue | sendVenue
*DefaultAPI* | [**PostSendVideo**](docs/DefaultAPI.md#postsendvideo) | **Post** /sendVideo | sendVideo
*DefaultAPI* | [**PostSendVideoNote**](docs/DefaultAPI.md#postsendvideonote) | **Post** /sendVideoNote | sendVideoNote
*DefaultAPI* | [**PostSendVoice**](docs/DefaultAPI.md#postsendvoice) | **Post** /sendVoice | sendVoice
*DefaultAPI* | [**PostSetBusinessAccountBio**](docs/DefaultAPI.md#postsetbusinessaccountbio) | **Post** /setBusinessAccountBio | setBusinessAccountBio
*DefaultAPI* | [**PostSetBusinessAccountGiftSettings**](docs/DefaultAPI.md#postsetbusinessaccountgiftsettings) | **Post** /setBusinessAccountGiftSettings | setBusinessAccountGiftSettings
*DefaultAPI* | [**PostSetBusinessAccountName**](docs/DefaultAPI.md#postsetbusinessaccountname) | **Post** /setBusinessAccountName | setBusinessAccountName
*DefaultAPI* | [**PostSetBusinessAccountProfilePhoto**](docs/DefaultAPI.md#postsetbusinessaccountprofilephoto) | **Post** /setBusinessAccountProfilePhoto | setBusinessAccountProfilePhoto
*DefaultAPI* | [**PostSetBusinessAccountUsername**](docs/DefaultAPI.md#postsetbusinessaccountusername) | **Post** /setBusinessAccountUsername | setBusinessAccountUsername
*DefaultAPI* | [**PostSetChatAdministratorCustomTitle**](docs/DefaultAPI.md#postsetchatadministratorcustomtitle) | **Post** /setChatAdministratorCustomTitle | setChatAdministratorCustomTitle
*DefaultAPI* | [**PostSetChatDescription**](docs/DefaultAPI.md#postsetchatdescription) | **Post** /setChatDescription | setChatDescription
*DefaultAPI* | [**PostSetChatMenuButton**](docs/DefaultAPI.md#postsetchatmenubutton) | **Post** /setChatMenuButton | setChatMenuButton
*DefaultAPI* | [**PostSetChatPermissions**](docs/DefaultAPI.md#postsetchatpermissions) | **Post** /setChatPermissions | setChatPermissions
*DefaultAPI* | [**PostSetChatPhoto**](docs/DefaultAPI.md#postsetchatphoto) | **Post** /setChatPhoto | setChatPhoto
*DefaultAPI* | [**PostSetChatStickerSet**](docs/DefaultAPI.md#postsetchatstickerset) | **Post** /setChatStickerSet | setChatStickerSet
*DefaultAPI* | [**PostSetChatTitle**](docs/DefaultAPI.md#postsetchattitle) | **Post** /setChatTitle | setChatTitle
*DefaultAPI* | [**PostSetCustomEmojiStickerSetThumbnail**](docs/DefaultAPI.md#postsetcustomemojistickersetthumbnail) | **Post** /setCustomEmojiStickerSetThumbnail | setCustomEmojiStickerSetThumbnail
*DefaultAPI* | [**PostSetGameScore**](docs/DefaultAPI.md#postsetgamescore) | **Post** /setGameScore | setGameScore
*DefaultAPI* | [**PostSetMessageReaction**](docs/DefaultAPI.md#postsetmessagereaction) | **Post** /setMessageReaction | setMessageReaction
*DefaultAPI* | [**PostSetMyCommands**](docs/DefaultAPI.md#postsetmycommands) | **Post** /setMyCommands | setMyCommands
*DefaultAPI* | [**PostSetMyDefaultAdministratorRights**](docs/DefaultAPI.md#postsetmydefaultadministratorrights) | **Post** /setMyDefaultAdministratorRights | setMyDefaultAdministratorRights
*DefaultAPI* | [**PostSetMyDescription**](docs/DefaultAPI.md#postsetmydescription) | **Post** /setMyDescription | setMyDescription
*DefaultAPI* | [**PostSetMyName**](docs/DefaultAPI.md#postsetmyname) | **Post** /setMyName | setMyName
*DefaultAPI* | [**PostSetMyShortDescription**](docs/DefaultAPI.md#postsetmyshortdescription) | **Post** /setMyShortDescription | setMyShortDescription
*DefaultAPI* | [**PostSetPassportDataErrors**](docs/DefaultAPI.md#postsetpassportdataerrors) | **Post** /setPassportDataErrors | setPassportDataErrors
*DefaultAPI* | [**PostSetStickerEmojiList**](docs/DefaultAPI.md#postsetstickeremojilist) | **Post** /setStickerEmojiList | setStickerEmojiList
*DefaultAPI* | [**PostSetStickerKeywords**](docs/DefaultAPI.md#postsetstickerkeywords) | **Post** /setStickerKeywords | setStickerKeywords
*DefaultAPI* | [**PostSetStickerMaskPosition**](docs/DefaultAPI.md#postsetstickermaskposition) | **Post** /setStickerMaskPosition | setStickerMaskPosition
*DefaultAPI* | [**PostSetStickerPositionInSet**](docs/DefaultAPI.md#postsetstickerpositioninset) | **Post** /setStickerPositionInSet | setStickerPositionInSet
*DefaultAPI* | [**PostSetStickerSetThumbnail**](docs/DefaultAPI.md#postsetstickersetthumbnail) | **Post** /setStickerSetThumbnail | setStickerSetThumbnail
*DefaultAPI* | [**PostSetStickerSetTitle**](docs/DefaultAPI.md#postsetstickersettitle) | **Post** /setStickerSetTitle | setStickerSetTitle
*DefaultAPI* | [**PostSetUserEmojiStatus**](docs/DefaultAPI.md#postsetuseremojistatus) | **Post** /setUserEmojiStatus | setUserEmojiStatus
*DefaultAPI* | [**PostSetWebhook**](docs/DefaultAPI.md#postsetwebhook) | **Post** /setWebhook | setWebhook
*DefaultAPI* | [**PostStopMessageLiveLocation**](docs/DefaultAPI.md#poststopmessagelivelocation) | **Post** /stopMessageLiveLocation | stopMessageLiveLocation
*DefaultAPI* | [**PostStopPoll**](docs/DefaultAPI.md#poststoppoll) | **Post** /stopPoll | stopPoll
*DefaultAPI* | [**PostTransferBusinessAccountStars**](docs/DefaultAPI.md#posttransferbusinessaccountstars) | **Post** /transferBusinessAccountStars | transferBusinessAccountStars
*DefaultAPI* | [**PostTransferGift**](docs/DefaultAPI.md#posttransfergift) | **Post** /transferGift | transferGift
*DefaultAPI* | [**PostUnbanChatMember**](docs/DefaultAPI.md#postunbanchatmember) | **Post** /unbanChatMember | unbanChatMember
*DefaultAPI* | [**PostUnbanChatSenderChat**](docs/DefaultAPI.md#postunbanchatsenderchat) | **Post** /unbanChatSenderChat | unbanChatSenderChat
*DefaultAPI* | [**PostUnhideGeneralForumTopic**](docs/DefaultAPI.md#postunhidegeneralforumtopic) | **Post** /unhideGeneralForumTopic | unhideGeneralForumTopic
*DefaultAPI* | [**PostUnpinAllChatMessages**](docs/DefaultAPI.md#postunpinallchatmessages) | **Post** /unpinAllChatMessages | unpinAllChatMessages
*DefaultAPI* | [**PostUnpinAllForumTopicMessages**](docs/DefaultAPI.md#postunpinallforumtopicmessages) | **Post** /unpinAllForumTopicMessages | unpinAllForumTopicMessages
*DefaultAPI* | [**PostUnpinAllGeneralForumTopicMessages**](docs/DefaultAPI.md#postunpinallgeneralforumtopicmessages) | **Post** /unpinAllGeneralForumTopicMessages | unpinAllGeneralForumTopicMessages
*DefaultAPI* | [**PostUnpinChatMessage**](docs/DefaultAPI.md#postunpinchatmessage) | **Post** /unpinChatMessage | unpinChatMessage
*DefaultAPI* | [**PostUpgradeGift**](docs/DefaultAPI.md#postupgradegift) | **Post** /upgradeGift | upgradeGift
*DefaultAPI* | [**PostUploadStickerFile**](docs/DefaultAPI.md#postuploadstickerfile) | **Post** /uploadStickerFile | uploadStickerFile
*DefaultAPI* | [**PostVerifyChat**](docs/DefaultAPI.md#postverifychat) | **Post** /verifyChat | verifyChat
*DefaultAPI* | [**PostVerifyUser**](docs/DefaultAPI.md#postverifyuser) | **Post** /verifyUser | verifyUser


## Documentation For Models

 - [AcceptedGiftTypes](docs/AcceptedGiftTypes.md)
 - [AddStickerToSetRequest](docs/AddStickerToSetRequest.md)
 - [AddStickerToSetResponse](docs/AddStickerToSetResponse.md)
 - [AffiliateInfo](docs/AffiliateInfo.md)
 - [Animation](docs/Animation.md)
 - [AnswerCallbackQueryRequest](docs/AnswerCallbackQueryRequest.md)
 - [AnswerCallbackQueryResponse](docs/AnswerCallbackQueryResponse.md)
 - [AnswerInlineQueryRequest](docs/AnswerInlineQueryRequest.md)
 - [AnswerInlineQueryResponse](docs/AnswerInlineQueryResponse.md)
 - [AnswerPreCheckoutQueryRequest](docs/AnswerPreCheckoutQueryRequest.md)
 - [AnswerPreCheckoutQueryResponse](docs/AnswerPreCheckoutQueryResponse.md)
 - [AnswerShippingQueryRequest](docs/AnswerShippingQueryRequest.md)
 - [AnswerShippingQueryResponse](docs/AnswerShippingQueryResponse.md)
 - [AnswerWebAppQueryRequest](docs/AnswerWebAppQueryRequest.md)
 - [AnswerWebAppQueryResponse](docs/AnswerWebAppQueryResponse.md)
 - [ApproveChatJoinRequestRequest](docs/ApproveChatJoinRequestRequest.md)
 - [ApproveChatJoinRequestResponse](docs/ApproveChatJoinRequestResponse.md)
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
 - [BanChatMemberRequest](docs/BanChatMemberRequest.md)
 - [BanChatMemberRequestChatId](docs/BanChatMemberRequestChatId.md)
 - [BanChatMemberResponse](docs/BanChatMemberResponse.md)
 - [BanChatSenderChatRequest](docs/BanChatSenderChatRequest.md)
 - [BanChatSenderChatResponse](docs/BanChatSenderChatResponse.md)
 - [Birthdate](docs/Birthdate.md)
 - [BotCommand](docs/BotCommand.md)
 - [BotCommandScope](docs/BotCommandScope.md)
 - [BotCommandScopeAllChatAdministrators](docs/BotCommandScopeAllChatAdministrators.md)
 - [BotCommandScopeAllGroupChats](docs/BotCommandScopeAllGroupChats.md)
 - [BotCommandScopeAllPrivateChats](docs/BotCommandScopeAllPrivateChats.md)
 - [BotCommandScopeChat](docs/BotCommandScopeChat.md)
 - [BotCommandScopeChatAdministrators](docs/BotCommandScopeChatAdministrators.md)
 - [BotCommandScopeChatChatId](docs/BotCommandScopeChatChatId.md)
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
 - [Checklist](docs/Checklist.md)
 - [ChecklistTask](docs/ChecklistTask.md)
 - [ChecklistTasksAdded](docs/ChecklistTasksAdded.md)
 - [ChecklistTasksDone](docs/ChecklistTasksDone.md)
 - [ChosenInlineResult](docs/ChosenInlineResult.md)
 - [CloseForumTopicRequest](docs/CloseForumTopicRequest.md)
 - [CloseForumTopicResponse](docs/CloseForumTopicResponse.md)
 - [CloseGeneralForumTopicRequest](docs/CloseGeneralForumTopicRequest.md)
 - [CloseGeneralForumTopicResponse](docs/CloseGeneralForumTopicResponse.md)
 - [CloseResponse](docs/CloseResponse.md)
 - [Contact](docs/Contact.md)
 - [ConvertGiftToStarsRequest](docs/ConvertGiftToStarsRequest.md)
 - [ConvertGiftToStarsResponse](docs/ConvertGiftToStarsResponse.md)
 - [CopyMessageRequest](docs/CopyMessageRequest.md)
 - [CopyMessageResponse](docs/CopyMessageResponse.md)
 - [CopyMessagesRequest](docs/CopyMessagesRequest.md)
 - [CopyMessagesResponse](docs/CopyMessagesResponse.md)
 - [CopyTextButton](docs/CopyTextButton.md)
 - [CreateChatInviteLinkRequest](docs/CreateChatInviteLinkRequest.md)
 - [CreateChatInviteLinkResponse](docs/CreateChatInviteLinkResponse.md)
 - [CreateChatSubscriptionInviteLinkRequest](docs/CreateChatSubscriptionInviteLinkRequest.md)
 - [CreateChatSubscriptionInviteLinkRequestChatId](docs/CreateChatSubscriptionInviteLinkRequestChatId.md)
 - [CreateChatSubscriptionInviteLinkResponse](docs/CreateChatSubscriptionInviteLinkResponse.md)
 - [CreateForumTopicRequest](docs/CreateForumTopicRequest.md)
 - [CreateForumTopicResponse](docs/CreateForumTopicResponse.md)
 - [CreateInvoiceLinkRequest](docs/CreateInvoiceLinkRequest.md)
 - [CreateInvoiceLinkResponse](docs/CreateInvoiceLinkResponse.md)
 - [CreateNewStickerSetRequest](docs/CreateNewStickerSetRequest.md)
 - [CreateNewStickerSetResponse](docs/CreateNewStickerSetResponse.md)
 - [DeclineChatJoinRequestRequest](docs/DeclineChatJoinRequestRequest.md)
 - [DeclineChatJoinRequestResponse](docs/DeclineChatJoinRequestResponse.md)
 - [DeleteBusinessMessagesRequest](docs/DeleteBusinessMessagesRequest.md)
 - [DeleteBusinessMessagesResponse](docs/DeleteBusinessMessagesResponse.md)
 - [DeleteChatPhotoRequest](docs/DeleteChatPhotoRequest.md)
 - [DeleteChatPhotoResponse](docs/DeleteChatPhotoResponse.md)
 - [DeleteChatStickerSetRequest](docs/DeleteChatStickerSetRequest.md)
 - [DeleteChatStickerSetResponse](docs/DeleteChatStickerSetResponse.md)
 - [DeleteForumTopicRequest](docs/DeleteForumTopicRequest.md)
 - [DeleteForumTopicResponse](docs/DeleteForumTopicResponse.md)
 - [DeleteMessageRequest](docs/DeleteMessageRequest.md)
 - [DeleteMessageResponse](docs/DeleteMessageResponse.md)
 - [DeleteMessagesRequest](docs/DeleteMessagesRequest.md)
 - [DeleteMessagesResponse](docs/DeleteMessagesResponse.md)
 - [DeleteMyCommandsRequest](docs/DeleteMyCommandsRequest.md)
 - [DeleteMyCommandsResponse](docs/DeleteMyCommandsResponse.md)
 - [DeleteStickerFromSetRequest](docs/DeleteStickerFromSetRequest.md)
 - [DeleteStickerFromSetResponse](docs/DeleteStickerFromSetResponse.md)
 - [DeleteStickerSetRequest](docs/DeleteStickerSetRequest.md)
 - [DeleteStickerSetResponse](docs/DeleteStickerSetResponse.md)
 - [DeleteStoryRequest](docs/DeleteStoryRequest.md)
 - [DeleteStoryResponse](docs/DeleteStoryResponse.md)
 - [DeleteWebhookRequest](docs/DeleteWebhookRequest.md)
 - [DeleteWebhookResponse](docs/DeleteWebhookResponse.md)
 - [Dice](docs/Dice.md)
 - [DirectMessagePriceChanged](docs/DirectMessagePriceChanged.md)
 - [Document](docs/Document.md)
 - [EditChatInviteLinkRequest](docs/EditChatInviteLinkRequest.md)
 - [EditChatInviteLinkResponse](docs/EditChatInviteLinkResponse.md)
 - [EditChatSubscriptionInviteLinkRequest](docs/EditChatSubscriptionInviteLinkRequest.md)
 - [EditChatSubscriptionInviteLinkResponse](docs/EditChatSubscriptionInviteLinkResponse.md)
 - [EditForumTopicRequest](docs/EditForumTopicRequest.md)
 - [EditForumTopicResponse](docs/EditForumTopicResponse.md)
 - [EditGeneralForumTopicRequest](docs/EditGeneralForumTopicRequest.md)
 - [EditGeneralForumTopicResponse](docs/EditGeneralForumTopicResponse.md)
 - [EditMessageCaptionRequest](docs/EditMessageCaptionRequest.md)
 - [EditMessageCaptionResponse](docs/EditMessageCaptionResponse.md)
 - [EditMessageChecklistRequest](docs/EditMessageChecklistRequest.md)
 - [EditMessageChecklistResponse](docs/EditMessageChecklistResponse.md)
 - [EditMessageLiveLocationRequest](docs/EditMessageLiveLocationRequest.md)
 - [EditMessageLiveLocationResponse](docs/EditMessageLiveLocationResponse.md)
 - [EditMessageMediaRequest](docs/EditMessageMediaRequest.md)
 - [EditMessageMediaResponse](docs/EditMessageMediaResponse.md)
 - [EditMessageReplyMarkupRequest](docs/EditMessageReplyMarkupRequest.md)
 - [EditMessageReplyMarkupResponse](docs/EditMessageReplyMarkupResponse.md)
 - [EditMessageTextRequest](docs/EditMessageTextRequest.md)
 - [EditMessageTextRequestChatId](docs/EditMessageTextRequestChatId.md)
 - [EditMessageTextResponse](docs/EditMessageTextResponse.md)
 - [EditMessageTextResponseResult](docs/EditMessageTextResponseResult.md)
 - [EditStoryRequest](docs/EditStoryRequest.md)
 - [EditStoryResponse](docs/EditStoryResponse.md)
 - [EditUserStarSubscriptionRequest](docs/EditUserStarSubscriptionRequest.md)
 - [EditUserStarSubscriptionResponse](docs/EditUserStarSubscriptionResponse.md)
 - [EncryptedCredentials](docs/EncryptedCredentials.md)
 - [EncryptedPassportElement](docs/EncryptedPassportElement.md)
 - [Error](docs/Error.md)
 - [ExportChatInviteLinkRequest](docs/ExportChatInviteLinkRequest.md)
 - [ExportChatInviteLinkResponse](docs/ExportChatInviteLinkResponse.md)
 - [ExternalReplyInfo](docs/ExternalReplyInfo.md)
 - [File](docs/File.md)
 - [ForceReply](docs/ForceReply.md)
 - [ForumTopic](docs/ForumTopic.md)
 - [ForumTopicCreated](docs/ForumTopicCreated.md)
 - [ForumTopicEdited](docs/ForumTopicEdited.md)
 - [ForwardMessageRequest](docs/ForwardMessageRequest.md)
 - [ForwardMessageRequestFromChatId](docs/ForwardMessageRequestFromChatId.md)
 - [ForwardMessageResponse](docs/ForwardMessageResponse.md)
 - [ForwardMessagesRequest](docs/ForwardMessagesRequest.md)
 - [ForwardMessagesRequestFromChatId](docs/ForwardMessagesRequestFromChatId.md)
 - [ForwardMessagesResponse](docs/ForwardMessagesResponse.md)
 - [Game](docs/Game.md)
 - [GameHighScore](docs/GameHighScore.md)
 - [GetAvailableGiftsResponse](docs/GetAvailableGiftsResponse.md)
 - [GetBusinessAccountGiftsRequest](docs/GetBusinessAccountGiftsRequest.md)
 - [GetBusinessAccountGiftsResponse](docs/GetBusinessAccountGiftsResponse.md)
 - [GetBusinessAccountStarBalanceRequest](docs/GetBusinessAccountStarBalanceRequest.md)
 - [GetBusinessAccountStarBalanceResponse](docs/GetBusinessAccountStarBalanceResponse.md)
 - [GetBusinessConnectionRequest](docs/GetBusinessConnectionRequest.md)
 - [GetBusinessConnectionResponse](docs/GetBusinessConnectionResponse.md)
 - [GetChatAdministratorsRequest](docs/GetChatAdministratorsRequest.md)
 - [GetChatAdministratorsResponse](docs/GetChatAdministratorsResponse.md)
 - [GetChatMemberCountRequest](docs/GetChatMemberCountRequest.md)
 - [GetChatMemberCountResponse](docs/GetChatMemberCountResponse.md)
 - [GetChatMemberRequest](docs/GetChatMemberRequest.md)
 - [GetChatMemberResponse](docs/GetChatMemberResponse.md)
 - [GetChatMenuButtonRequest](docs/GetChatMenuButtonRequest.md)
 - [GetChatMenuButtonResponse](docs/GetChatMenuButtonResponse.md)
 - [GetChatRequest](docs/GetChatRequest.md)
 - [GetChatResponse](docs/GetChatResponse.md)
 - [GetCustomEmojiStickersRequest](docs/GetCustomEmojiStickersRequest.md)
 - [GetCustomEmojiStickersResponse](docs/GetCustomEmojiStickersResponse.md)
 - [GetFileRequest](docs/GetFileRequest.md)
 - [GetFileResponse](docs/GetFileResponse.md)
 - [GetForumTopicIconStickersResponse](docs/GetForumTopicIconStickersResponse.md)
 - [GetGameHighScoresRequest](docs/GetGameHighScoresRequest.md)
 - [GetGameHighScoresResponse](docs/GetGameHighScoresResponse.md)
 - [GetMeResponse](docs/GetMeResponse.md)
 - [GetMyCommandsRequest](docs/GetMyCommandsRequest.md)
 - [GetMyCommandsResponse](docs/GetMyCommandsResponse.md)
 - [GetMyDefaultAdministratorRightsRequest](docs/GetMyDefaultAdministratorRightsRequest.md)
 - [GetMyDefaultAdministratorRightsResponse](docs/GetMyDefaultAdministratorRightsResponse.md)
 - [GetMyDescriptionRequest](docs/GetMyDescriptionRequest.md)
 - [GetMyDescriptionResponse](docs/GetMyDescriptionResponse.md)
 - [GetMyNameRequest](docs/GetMyNameRequest.md)
 - [GetMyNameResponse](docs/GetMyNameResponse.md)
 - [GetMyShortDescriptionRequest](docs/GetMyShortDescriptionRequest.md)
 - [GetMyShortDescriptionResponse](docs/GetMyShortDescriptionResponse.md)
 - [GetMyStarBalanceResponse](docs/GetMyStarBalanceResponse.md)
 - [GetStarTransactionsRequest](docs/GetStarTransactionsRequest.md)
 - [GetStarTransactionsResponse](docs/GetStarTransactionsResponse.md)
 - [GetStickerSetRequest](docs/GetStickerSetRequest.md)
 - [GetStickerSetResponse](docs/GetStickerSetResponse.md)
 - [GetUpdatesRequest](docs/GetUpdatesRequest.md)
 - [GetUpdatesResponse](docs/GetUpdatesResponse.md)
 - [GetUserChatBoostsRequest](docs/GetUserChatBoostsRequest.md)
 - [GetUserChatBoostsRequestChatId](docs/GetUserChatBoostsRequestChatId.md)
 - [GetUserChatBoostsResponse](docs/GetUserChatBoostsResponse.md)
 - [GetUserProfilePhotosRequest](docs/GetUserProfilePhotosRequest.md)
 - [GetUserProfilePhotosResponse](docs/GetUserProfilePhotosResponse.md)
 - [GetWebhookInfoResponse](docs/GetWebhookInfoResponse.md)
 - [Gift](docs/Gift.md)
 - [GiftInfo](docs/GiftInfo.md)
 - [GiftPremiumSubscriptionRequest](docs/GiftPremiumSubscriptionRequest.md)
 - [GiftPremiumSubscriptionResponse](docs/GiftPremiumSubscriptionResponse.md)
 - [Gifts](docs/Gifts.md)
 - [Giveaway](docs/Giveaway.md)
 - [GiveawayCompleted](docs/GiveawayCompleted.md)
 - [GiveawayCreated](docs/GiveawayCreated.md)
 - [GiveawayWinners](docs/GiveawayWinners.md)
 - [HideGeneralForumTopicRequest](docs/HideGeneralForumTopicRequest.md)
 - [HideGeneralForumTopicResponse](docs/HideGeneralForumTopicResponse.md)
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
 - [InputChecklist](docs/InputChecklist.md)
 - [InputChecklistTask](docs/InputChecklistTask.md)
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
 - [LeaveChatRequest](docs/LeaveChatRequest.md)
 - [LeaveChatRequestChatId](docs/LeaveChatRequestChatId.md)
 - [LeaveChatResponse](docs/LeaveChatResponse.md)
 - [LinkPreviewOptions](docs/LinkPreviewOptions.md)
 - [Location](docs/Location.md)
 - [LocationAddress](docs/LocationAddress.md)
 - [LogOutResponse](docs/LogOutResponse.md)
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
 - [PinChatMessageRequest](docs/PinChatMessageRequest.md)
 - [PinChatMessageResponse](docs/PinChatMessageResponse.md)
 - [Poll](docs/Poll.md)
 - [PollAnswer](docs/PollAnswer.md)
 - [PollOption](docs/PollOption.md)
 - [PostStoryRequest](docs/PostStoryRequest.md)
 - [PostStoryResponse](docs/PostStoryResponse.md)
 - [PreCheckoutQuery](docs/PreCheckoutQuery.md)
 - [PreparedInlineMessage](docs/PreparedInlineMessage.md)
 - [PromoteChatMemberRequest](docs/PromoteChatMemberRequest.md)
 - [PromoteChatMemberResponse](docs/PromoteChatMemberResponse.md)
 - [ProximityAlertTriggered](docs/ProximityAlertTriggered.md)
 - [ReactionCount](docs/ReactionCount.md)
 - [ReactionType](docs/ReactionType.md)
 - [ReactionTypeCustomEmoji](docs/ReactionTypeCustomEmoji.md)
 - [ReactionTypeEmoji](docs/ReactionTypeEmoji.md)
 - [ReactionTypePaid](docs/ReactionTypePaid.md)
 - [ReadBusinessMessageRequest](docs/ReadBusinessMessageRequest.md)
 - [ReadBusinessMessageResponse](docs/ReadBusinessMessageResponse.md)
 - [RefundStarPaymentRequest](docs/RefundStarPaymentRequest.md)
 - [RefundStarPaymentResponse](docs/RefundStarPaymentResponse.md)
 - [RefundedPayment](docs/RefundedPayment.md)
 - [RemoveBusinessAccountProfilePhotoRequest](docs/RemoveBusinessAccountProfilePhotoRequest.md)
 - [RemoveBusinessAccountProfilePhotoResponse](docs/RemoveBusinessAccountProfilePhotoResponse.md)
 - [RemoveChatVerificationRequest](docs/RemoveChatVerificationRequest.md)
 - [RemoveChatVerificationResponse](docs/RemoveChatVerificationResponse.md)
 - [RemoveUserVerificationRequest](docs/RemoveUserVerificationRequest.md)
 - [RemoveUserVerificationResponse](docs/RemoveUserVerificationResponse.md)
 - [ReopenForumTopicRequest](docs/ReopenForumTopicRequest.md)
 - [ReopenForumTopicResponse](docs/ReopenForumTopicResponse.md)
 - [ReopenGeneralForumTopicRequest](docs/ReopenGeneralForumTopicRequest.md)
 - [ReopenGeneralForumTopicResponse](docs/ReopenGeneralForumTopicResponse.md)
 - [ReplaceStickerInSetRequest](docs/ReplaceStickerInSetRequest.md)
 - [ReplaceStickerInSetResponse](docs/ReplaceStickerInSetResponse.md)
 - [ReplyKeyboardMarkup](docs/ReplyKeyboardMarkup.md)
 - [ReplyKeyboardRemove](docs/ReplyKeyboardRemove.md)
 - [ReplyParameters](docs/ReplyParameters.md)
 - [ReplyParametersChatId](docs/ReplyParametersChatId.md)
 - [ResponseParameters](docs/ResponseParameters.md)
 - [RestrictChatMemberRequest](docs/RestrictChatMemberRequest.md)
 - [RestrictChatMemberResponse](docs/RestrictChatMemberResponse.md)
 - [RevenueWithdrawalState](docs/RevenueWithdrawalState.md)
 - [RevenueWithdrawalStateFailed](docs/RevenueWithdrawalStateFailed.md)
 - [RevenueWithdrawalStatePending](docs/RevenueWithdrawalStatePending.md)
 - [RevenueWithdrawalStateSucceeded](docs/RevenueWithdrawalStateSucceeded.md)
 - [RevokeChatInviteLinkRequest](docs/RevokeChatInviteLinkRequest.md)
 - [RevokeChatInviteLinkRequestChatId](docs/RevokeChatInviteLinkRequestChatId.md)
 - [RevokeChatInviteLinkResponse](docs/RevokeChatInviteLinkResponse.md)
 - [SavePreparedInlineMessageRequest](docs/SavePreparedInlineMessageRequest.md)
 - [SavePreparedInlineMessageResponse](docs/SavePreparedInlineMessageResponse.md)
 - [SendAnimationRequest](docs/SendAnimationRequest.md)
 - [SendAnimationResponse](docs/SendAnimationResponse.md)
 - [SendAudioRequest](docs/SendAudioRequest.md)
 - [SendAudioResponse](docs/SendAudioResponse.md)
 - [SendChatActionRequest](docs/SendChatActionRequest.md)
 - [SendChatActionResponse](docs/SendChatActionResponse.md)
 - [SendChecklistRequest](docs/SendChecklistRequest.md)
 - [SendChecklistResponse](docs/SendChecklistResponse.md)
 - [SendContactRequest](docs/SendContactRequest.md)
 - [SendContactResponse](docs/SendContactResponse.md)
 - [SendDiceRequest](docs/SendDiceRequest.md)
 - [SendDiceResponse](docs/SendDiceResponse.md)
 - [SendDocumentRequest](docs/SendDocumentRequest.md)
 - [SendDocumentResponse](docs/SendDocumentResponse.md)
 - [SendGameRequest](docs/SendGameRequest.md)
 - [SendGameResponse](docs/SendGameResponse.md)
 - [SendGiftRequest](docs/SendGiftRequest.md)
 - [SendGiftRequestChatId](docs/SendGiftRequestChatId.md)
 - [SendGiftResponse](docs/SendGiftResponse.md)
 - [SendInvoiceRequest](docs/SendInvoiceRequest.md)
 - [SendInvoiceResponse](docs/SendInvoiceResponse.md)
 - [SendLocationRequest](docs/SendLocationRequest.md)
 - [SendLocationResponse](docs/SendLocationResponse.md)
 - [SendMediaGroupRequest](docs/SendMediaGroupRequest.md)
 - [SendMediaGroupRequestMediaInner](docs/SendMediaGroupRequestMediaInner.md)
 - [SendMediaGroupResponse](docs/SendMediaGroupResponse.md)
 - [SendMessageRequest](docs/SendMessageRequest.md)
 - [SendMessageRequestChatId](docs/SendMessageRequestChatId.md)
 - [SendMessageRequestReplyMarkup](docs/SendMessageRequestReplyMarkup.md)
 - [SendMessageResponse](docs/SendMessageResponse.md)
 - [SendPaidMediaRequest](docs/SendPaidMediaRequest.md)
 - [SendPaidMediaRequestChatId](docs/SendPaidMediaRequestChatId.md)
 - [SendPaidMediaResponse](docs/SendPaidMediaResponse.md)
 - [SendPhotoRequest](docs/SendPhotoRequest.md)
 - [SendPhotoResponse](docs/SendPhotoResponse.md)
 - [SendPollRequest](docs/SendPollRequest.md)
 - [SendPollResponse](docs/SendPollResponse.md)
 - [SendStickerRequest](docs/SendStickerRequest.md)
 - [SendStickerResponse](docs/SendStickerResponse.md)
 - [SendVenueRequest](docs/SendVenueRequest.md)
 - [SendVenueResponse](docs/SendVenueResponse.md)
 - [SendVideoNoteRequest](docs/SendVideoNoteRequest.md)
 - [SendVideoNoteResponse](docs/SendVideoNoteResponse.md)
 - [SendVideoRequest](docs/SendVideoRequest.md)
 - [SendVideoResponse](docs/SendVideoResponse.md)
 - [SendVoiceRequest](docs/SendVoiceRequest.md)
 - [SendVoiceResponse](docs/SendVoiceResponse.md)
 - [SentWebAppMessage](docs/SentWebAppMessage.md)
 - [SetBusinessAccountBioRequest](docs/SetBusinessAccountBioRequest.md)
 - [SetBusinessAccountBioResponse](docs/SetBusinessAccountBioResponse.md)
 - [SetBusinessAccountGiftSettingsRequest](docs/SetBusinessAccountGiftSettingsRequest.md)
 - [SetBusinessAccountGiftSettingsResponse](docs/SetBusinessAccountGiftSettingsResponse.md)
 - [SetBusinessAccountNameRequest](docs/SetBusinessAccountNameRequest.md)
 - [SetBusinessAccountNameResponse](docs/SetBusinessAccountNameResponse.md)
 - [SetBusinessAccountProfilePhotoRequest](docs/SetBusinessAccountProfilePhotoRequest.md)
 - [SetBusinessAccountProfilePhotoResponse](docs/SetBusinessAccountProfilePhotoResponse.md)
 - [SetBusinessAccountUsernameRequest](docs/SetBusinessAccountUsernameRequest.md)
 - [SetBusinessAccountUsernameResponse](docs/SetBusinessAccountUsernameResponse.md)
 - [SetChatAdministratorCustomTitleRequest](docs/SetChatAdministratorCustomTitleRequest.md)
 - [SetChatAdministratorCustomTitleResponse](docs/SetChatAdministratorCustomTitleResponse.md)
 - [SetChatDescriptionRequest](docs/SetChatDescriptionRequest.md)
 - [SetChatDescriptionResponse](docs/SetChatDescriptionResponse.md)
 - [SetChatMenuButtonRequest](docs/SetChatMenuButtonRequest.md)
 - [SetChatMenuButtonResponse](docs/SetChatMenuButtonResponse.md)
 - [SetChatPermissionsRequest](docs/SetChatPermissionsRequest.md)
 - [SetChatPermissionsResponse](docs/SetChatPermissionsResponse.md)
 - [SetChatPhotoRequest](docs/SetChatPhotoRequest.md)
 - [SetChatPhotoResponse](docs/SetChatPhotoResponse.md)
 - [SetChatStickerSetRequest](docs/SetChatStickerSetRequest.md)
 - [SetChatStickerSetResponse](docs/SetChatStickerSetResponse.md)
 - [SetChatTitleRequest](docs/SetChatTitleRequest.md)
 - [SetChatTitleResponse](docs/SetChatTitleResponse.md)
 - [SetCustomEmojiStickerSetThumbnailRequest](docs/SetCustomEmojiStickerSetThumbnailRequest.md)
 - [SetCustomEmojiStickerSetThumbnailResponse](docs/SetCustomEmojiStickerSetThumbnailResponse.md)
 - [SetGameScoreRequest](docs/SetGameScoreRequest.md)
 - [SetGameScoreResponse](docs/SetGameScoreResponse.md)
 - [SetMessageReactionRequest](docs/SetMessageReactionRequest.md)
 - [SetMessageReactionResponse](docs/SetMessageReactionResponse.md)
 - [SetMyCommandsRequest](docs/SetMyCommandsRequest.md)
 - [SetMyCommandsResponse](docs/SetMyCommandsResponse.md)
 - [SetMyDefaultAdministratorRightsRequest](docs/SetMyDefaultAdministratorRightsRequest.md)
 - [SetMyDefaultAdministratorRightsResponse](docs/SetMyDefaultAdministratorRightsResponse.md)
 - [SetMyDescriptionRequest](docs/SetMyDescriptionRequest.md)
 - [SetMyDescriptionResponse](docs/SetMyDescriptionResponse.md)
 - [SetMyNameRequest](docs/SetMyNameRequest.md)
 - [SetMyNameResponse](docs/SetMyNameResponse.md)
 - [SetMyShortDescriptionRequest](docs/SetMyShortDescriptionRequest.md)
 - [SetMyShortDescriptionResponse](docs/SetMyShortDescriptionResponse.md)
 - [SetPassportDataErrorsRequest](docs/SetPassportDataErrorsRequest.md)
 - [SetPassportDataErrorsResponse](docs/SetPassportDataErrorsResponse.md)
 - [SetStickerEmojiListRequest](docs/SetStickerEmojiListRequest.md)
 - [SetStickerEmojiListResponse](docs/SetStickerEmojiListResponse.md)
 - [SetStickerKeywordsRequest](docs/SetStickerKeywordsRequest.md)
 - [SetStickerKeywordsResponse](docs/SetStickerKeywordsResponse.md)
 - [SetStickerMaskPositionRequest](docs/SetStickerMaskPositionRequest.md)
 - [SetStickerMaskPositionResponse](docs/SetStickerMaskPositionResponse.md)
 - [SetStickerPositionInSetRequest](docs/SetStickerPositionInSetRequest.md)
 - [SetStickerPositionInSetResponse](docs/SetStickerPositionInSetResponse.md)
 - [SetStickerSetThumbnailRequest](docs/SetStickerSetThumbnailRequest.md)
 - [SetStickerSetThumbnailResponse](docs/SetStickerSetThumbnailResponse.md)
 - [SetStickerSetTitleRequest](docs/SetStickerSetTitleRequest.md)
 - [SetStickerSetTitleResponse](docs/SetStickerSetTitleResponse.md)
 - [SetUserEmojiStatusRequest](docs/SetUserEmojiStatusRequest.md)
 - [SetUserEmojiStatusResponse](docs/SetUserEmojiStatusResponse.md)
 - [SetWebhookRequest](docs/SetWebhookRequest.md)
 - [SetWebhookResponse](docs/SetWebhookResponse.md)
 - [SharedUser](docs/SharedUser.md)
 - [ShippingAddress](docs/ShippingAddress.md)
 - [ShippingOption](docs/ShippingOption.md)
 - [ShippingQuery](docs/ShippingQuery.md)
 - [StarAmount](docs/StarAmount.md)
 - [StarTransaction](docs/StarTransaction.md)
 - [StarTransactions](docs/StarTransactions.md)
 - [Sticker](docs/Sticker.md)
 - [StickerSet](docs/StickerSet.md)
 - [StopMessageLiveLocationRequest](docs/StopMessageLiveLocationRequest.md)
 - [StopMessageLiveLocationResponse](docs/StopMessageLiveLocationResponse.md)
 - [StopPollRequest](docs/StopPollRequest.md)
 - [StopPollResponse](docs/StopPollResponse.md)
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
 - [TransferBusinessAccountStarsRequest](docs/TransferBusinessAccountStarsRequest.md)
 - [TransferBusinessAccountStarsResponse](docs/TransferBusinessAccountStarsResponse.md)
 - [TransferGiftRequest](docs/TransferGiftRequest.md)
 - [TransferGiftResponse](docs/TransferGiftResponse.md)
 - [UnbanChatMemberRequest](docs/UnbanChatMemberRequest.md)
 - [UnbanChatMemberResponse](docs/UnbanChatMemberResponse.md)
 - [UnbanChatSenderChatRequest](docs/UnbanChatSenderChatRequest.md)
 - [UnbanChatSenderChatResponse](docs/UnbanChatSenderChatResponse.md)
 - [UnhideGeneralForumTopicRequest](docs/UnhideGeneralForumTopicRequest.md)
 - [UnhideGeneralForumTopicResponse](docs/UnhideGeneralForumTopicResponse.md)
 - [UniqueGift](docs/UniqueGift.md)
 - [UniqueGiftBackdrop](docs/UniqueGiftBackdrop.md)
 - [UniqueGiftBackdropColors](docs/UniqueGiftBackdropColors.md)
 - [UniqueGiftInfo](docs/UniqueGiftInfo.md)
 - [UniqueGiftModel](docs/UniqueGiftModel.md)
 - [UniqueGiftSymbol](docs/UniqueGiftSymbol.md)
 - [UnpinAllChatMessagesRequest](docs/UnpinAllChatMessagesRequest.md)
 - [UnpinAllChatMessagesResponse](docs/UnpinAllChatMessagesResponse.md)
 - [UnpinAllForumTopicMessagesRequest](docs/UnpinAllForumTopicMessagesRequest.md)
 - [UnpinAllForumTopicMessagesResponse](docs/UnpinAllForumTopicMessagesResponse.md)
 - [UnpinAllGeneralForumTopicMessagesRequest](docs/UnpinAllGeneralForumTopicMessagesRequest.md)
 - [UnpinAllGeneralForumTopicMessagesResponse](docs/UnpinAllGeneralForumTopicMessagesResponse.md)
 - [UnpinChatMessageRequest](docs/UnpinChatMessageRequest.md)
 - [UnpinChatMessageResponse](docs/UnpinChatMessageResponse.md)
 - [Update](docs/Update.md)
 - [UpgradeGiftRequest](docs/UpgradeGiftRequest.md)
 - [UpgradeGiftResponse](docs/UpgradeGiftResponse.md)
 - [UploadStickerFileRequest](docs/UploadStickerFileRequest.md)
 - [UploadStickerFileResponse](docs/UploadStickerFileResponse.md)
 - [User](docs/User.md)
 - [UserChatBoosts](docs/UserChatBoosts.md)
 - [UserProfilePhotos](docs/UserProfilePhotos.md)
 - [UsersShared](docs/UsersShared.md)
 - [Venue](docs/Venue.md)
 - [VerifyChatRequest](docs/VerifyChatRequest.md)
 - [VerifyChatResponse](docs/VerifyChatResponse.md)
 - [VerifyUserRequest](docs/VerifyUserRequest.md)
 - [VerifyUserResponse](docs/VerifyUserResponse.md)
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

gitctrlx@gmail.com

