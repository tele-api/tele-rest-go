# Go API client for tele_rest

The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram.
To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.

## Overview

- API version: 9.0.0
- Package version: 9.0.0
- Build date: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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
 - [AffiliateInfo](docs/AffiliateInfo.md)
 - [Animation](docs/Animation.md)
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
 - [Contact](docs/Contact.md)
 - [CopyTextButton](docs/CopyTextButton.md)
 - [Dice](docs/Dice.md)
 - [Document](docs/Document.md)
 - [EncryptedCredentials](docs/EncryptedCredentials.md)
 - [EncryptedPassportElement](docs/EncryptedPassportElement.md)
 - [Error](docs/Error.md)
 - [ExternalReplyInfo](docs/ExternalReplyInfo.md)
 - [File](docs/File.md)
 - [ForceReply](docs/ForceReply.md)
 - [ForumTopic](docs/ForumTopic.md)
 - [ForumTopicCreated](docs/ForumTopicCreated.md)
 - [ForumTopicEdited](docs/ForumTopicEdited.md)
 - [Game](docs/Game.md)
 - [GameHighScore](docs/GameHighScore.md)
 - [Gift](docs/Gift.md)
 - [GiftInfo](docs/GiftInfo.md)
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
 - [Poll](docs/Poll.md)
 - [PollAnswer](docs/PollAnswer.md)
 - [PollOption](docs/PollOption.md)
 - [PostAnswerCallbackQueryRequest](docs/PostAnswerCallbackQueryRequest.md)
 - [PostAnswerInlineQueryRequest](docs/PostAnswerInlineQueryRequest.md)
 - [PostAnswerPreCheckoutQueryRequest](docs/PostAnswerPreCheckoutQueryRequest.md)
 - [PostAnswerShippingQueryRequest](docs/PostAnswerShippingQueryRequest.md)
 - [PostAnswerWebAppQuery200Response](docs/PostAnswerWebAppQuery200Response.md)
 - [PostAnswerWebAppQueryRequest](docs/PostAnswerWebAppQueryRequest.md)
 - [PostApproveChatJoinRequestRequest](docs/PostApproveChatJoinRequestRequest.md)
 - [PostBanChatMemberRequest](docs/PostBanChatMemberRequest.md)
 - [PostBanChatMemberRequestChatId](docs/PostBanChatMemberRequestChatId.md)
 - [PostBanChatSenderChatRequest](docs/PostBanChatSenderChatRequest.md)
 - [PostCloseForumTopicRequest](docs/PostCloseForumTopicRequest.md)
 - [PostConvertGiftToStarsRequest](docs/PostConvertGiftToStarsRequest.md)
 - [PostCopyMessage200Response](docs/PostCopyMessage200Response.md)
 - [PostCopyMessageRequest](docs/PostCopyMessageRequest.md)
 - [PostCopyMessagesRequest](docs/PostCopyMessagesRequest.md)
 - [PostCreateChatInviteLink200Response](docs/PostCreateChatInviteLink200Response.md)
 - [PostCreateChatInviteLinkRequest](docs/PostCreateChatInviteLinkRequest.md)
 - [PostCreateChatSubscriptionInviteLinkRequest](docs/PostCreateChatSubscriptionInviteLinkRequest.md)
 - [PostCreateChatSubscriptionInviteLinkRequestChatId](docs/PostCreateChatSubscriptionInviteLinkRequestChatId.md)
 - [PostCreateForumTopic200Response](docs/PostCreateForumTopic200Response.md)
 - [PostCreateForumTopicRequest](docs/PostCreateForumTopicRequest.md)
 - [PostCreateInvoiceLinkRequest](docs/PostCreateInvoiceLinkRequest.md)
 - [PostDeleteBusinessMessagesRequest](docs/PostDeleteBusinessMessagesRequest.md)
 - [PostDeleteChatStickerSetRequest](docs/PostDeleteChatStickerSetRequest.md)
 - [PostDeleteMessageRequest](docs/PostDeleteMessageRequest.md)
 - [PostDeleteMessagesRequest](docs/PostDeleteMessagesRequest.md)
 - [PostDeleteMyCommandsRequest](docs/PostDeleteMyCommandsRequest.md)
 - [PostDeleteStickerFromSetRequest](docs/PostDeleteStickerFromSetRequest.md)
 - [PostDeleteStickerSetRequest](docs/PostDeleteStickerSetRequest.md)
 - [PostDeleteStoryRequest](docs/PostDeleteStoryRequest.md)
 - [PostDeleteWebhookRequest](docs/PostDeleteWebhookRequest.md)
 - [PostEditChatInviteLinkRequest](docs/PostEditChatInviteLinkRequest.md)
 - [PostEditChatSubscriptionInviteLinkRequest](docs/PostEditChatSubscriptionInviteLinkRequest.md)
 - [PostEditForumTopicRequest](docs/PostEditForumTopicRequest.md)
 - [PostEditGeneralForumTopicRequest](docs/PostEditGeneralForumTopicRequest.md)
 - [PostEditMessageCaptionRequest](docs/PostEditMessageCaptionRequest.md)
 - [PostEditMessageLiveLocationRequest](docs/PostEditMessageLiveLocationRequest.md)
 - [PostEditMessageReplyMarkupRequest](docs/PostEditMessageReplyMarkupRequest.md)
 - [PostEditMessageText200Response](docs/PostEditMessageText200Response.md)
 - [PostEditMessageText200ResponseResult](docs/PostEditMessageText200ResponseResult.md)
 - [PostEditMessageTextRequest](docs/PostEditMessageTextRequest.md)
 - [PostEditMessageTextRequestChatId](docs/PostEditMessageTextRequestChatId.md)
 - [PostEditUserStarSubscriptionRequest](docs/PostEditUserStarSubscriptionRequest.md)
 - [PostExportChatInviteLink200Response](docs/PostExportChatInviteLink200Response.md)
 - [PostExportChatInviteLinkRequest](docs/PostExportChatInviteLinkRequest.md)
 - [PostForwardMessageRequest](docs/PostForwardMessageRequest.md)
 - [PostForwardMessageRequestFromChatId](docs/PostForwardMessageRequestFromChatId.md)
 - [PostForwardMessages200Response](docs/PostForwardMessages200Response.md)
 - [PostForwardMessagesRequest](docs/PostForwardMessagesRequest.md)
 - [PostForwardMessagesRequestFromChatId](docs/PostForwardMessagesRequestFromChatId.md)
 - [PostGetAvailableGifts200Response](docs/PostGetAvailableGifts200Response.md)
 - [PostGetBusinessAccountGifts200Response](docs/PostGetBusinessAccountGifts200Response.md)
 - [PostGetBusinessAccountGiftsRequest](docs/PostGetBusinessAccountGiftsRequest.md)
 - [PostGetBusinessAccountStarBalance200Response](docs/PostGetBusinessAccountStarBalance200Response.md)
 - [PostGetBusinessConnection200Response](docs/PostGetBusinessConnection200Response.md)
 - [PostGetBusinessConnectionRequest](docs/PostGetBusinessConnectionRequest.md)
 - [PostGetChat200Response](docs/PostGetChat200Response.md)
 - [PostGetChatAdministrators200Response](docs/PostGetChatAdministrators200Response.md)
 - [PostGetChatMember200Response](docs/PostGetChatMember200Response.md)
 - [PostGetChatMemberCount200Response](docs/PostGetChatMemberCount200Response.md)
 - [PostGetChatMemberRequest](docs/PostGetChatMemberRequest.md)
 - [PostGetChatMenuButton200Response](docs/PostGetChatMenuButton200Response.md)
 - [PostGetChatMenuButtonRequest](docs/PostGetChatMenuButtonRequest.md)
 - [PostGetCustomEmojiStickersRequest](docs/PostGetCustomEmojiStickersRequest.md)
 - [PostGetFile200Response](docs/PostGetFile200Response.md)
 - [PostGetFileRequest](docs/PostGetFileRequest.md)
 - [PostGetForumTopicIconStickers200Response](docs/PostGetForumTopicIconStickers200Response.md)
 - [PostGetGameHighScores200Response](docs/PostGetGameHighScores200Response.md)
 - [PostGetGameHighScoresRequest](docs/PostGetGameHighScoresRequest.md)
 - [PostGetMe200Response](docs/PostGetMe200Response.md)
 - [PostGetMyCommands200Response](docs/PostGetMyCommands200Response.md)
 - [PostGetMyCommandsRequest](docs/PostGetMyCommandsRequest.md)
 - [PostGetMyDefaultAdministratorRights200Response](docs/PostGetMyDefaultAdministratorRights200Response.md)
 - [PostGetMyDefaultAdministratorRightsRequest](docs/PostGetMyDefaultAdministratorRightsRequest.md)
 - [PostGetMyDescription200Response](docs/PostGetMyDescription200Response.md)
 - [PostGetMyName200Response](docs/PostGetMyName200Response.md)
 - [PostGetMyNameRequest](docs/PostGetMyNameRequest.md)
 - [PostGetMyShortDescription200Response](docs/PostGetMyShortDescription200Response.md)
 - [PostGetStarTransactions200Response](docs/PostGetStarTransactions200Response.md)
 - [PostGetStarTransactionsRequest](docs/PostGetStarTransactionsRequest.md)
 - [PostGetStickerSet200Response](docs/PostGetStickerSet200Response.md)
 - [PostGetStickerSetRequest](docs/PostGetStickerSetRequest.md)
 - [PostGetUpdates200Response](docs/PostGetUpdates200Response.md)
 - [PostGetUpdatesRequest](docs/PostGetUpdatesRequest.md)
 - [PostGetUserChatBoosts200Response](docs/PostGetUserChatBoosts200Response.md)
 - [PostGetUserChatBoostsRequest](docs/PostGetUserChatBoostsRequest.md)
 - [PostGetUserChatBoostsRequestChatId](docs/PostGetUserChatBoostsRequestChatId.md)
 - [PostGetUserProfilePhotos200Response](docs/PostGetUserProfilePhotos200Response.md)
 - [PostGetUserProfilePhotosRequest](docs/PostGetUserProfilePhotosRequest.md)
 - [PostGetWebhookInfo200Response](docs/PostGetWebhookInfo200Response.md)
 - [PostGiftPremiumSubscriptionRequest](docs/PostGiftPremiumSubscriptionRequest.md)
 - [PostLeaveChatRequest](docs/PostLeaveChatRequest.md)
 - [PostLeaveChatRequestChatId](docs/PostLeaveChatRequestChatId.md)
 - [PostPinChatMessageRequest](docs/PostPinChatMessageRequest.md)
 - [PostPostStory200Response](docs/PostPostStory200Response.md)
 - [PostPromoteChatMemberRequest](docs/PostPromoteChatMemberRequest.md)
 - [PostReadBusinessMessageRequest](docs/PostReadBusinessMessageRequest.md)
 - [PostRefundStarPaymentRequest](docs/PostRefundStarPaymentRequest.md)
 - [PostRemoveBusinessAccountProfilePhotoRequest](docs/PostRemoveBusinessAccountProfilePhotoRequest.md)
 - [PostRemoveUserVerificationRequest](docs/PostRemoveUserVerificationRequest.md)
 - [PostRestrictChatMemberRequest](docs/PostRestrictChatMemberRequest.md)
 - [PostRestrictChatMemberRequestChatId](docs/PostRestrictChatMemberRequestChatId.md)
 - [PostRevokeChatInviteLinkRequest](docs/PostRevokeChatInviteLinkRequest.md)
 - [PostRevokeChatInviteLinkRequestChatId](docs/PostRevokeChatInviteLinkRequestChatId.md)
 - [PostSavePreparedInlineMessage200Response](docs/PostSavePreparedInlineMessage200Response.md)
 - [PostSavePreparedInlineMessageRequest](docs/PostSavePreparedInlineMessageRequest.md)
 - [PostSendAnimationRequestAnimation](docs/PostSendAnimationRequestAnimation.md)
 - [PostSendAudioRequestAudio](docs/PostSendAudioRequestAudio.md)
 - [PostSendAudioRequestThumbnail](docs/PostSendAudioRequestThumbnail.md)
 - [PostSendChatActionRequest](docs/PostSendChatActionRequest.md)
 - [PostSendContactRequest](docs/PostSendContactRequest.md)
 - [PostSendDiceRequest](docs/PostSendDiceRequest.md)
 - [PostSendDocumentRequestDocument](docs/PostSendDocumentRequestDocument.md)
 - [PostSendGameRequest](docs/PostSendGameRequest.md)
 - [PostSendGiftRequest](docs/PostSendGiftRequest.md)
 - [PostSendGiftRequestChatId](docs/PostSendGiftRequestChatId.md)
 - [PostSendInvoiceRequest](docs/PostSendInvoiceRequest.md)
 - [PostSendLocationRequest](docs/PostSendLocationRequest.md)
 - [PostSendMediaGroup200Response](docs/PostSendMediaGroup200Response.md)
 - [PostSendMediaGroupRequestMediaInner](docs/PostSendMediaGroupRequestMediaInner.md)
 - [PostSendMessage200Response](docs/PostSendMessage200Response.md)
 - [PostSendMessageRequest](docs/PostSendMessageRequest.md)
 - [PostSendMessageRequestChatId](docs/PostSendMessageRequestChatId.md)
 - [PostSendMessageRequestReplyMarkup](docs/PostSendMessageRequestReplyMarkup.md)
 - [PostSendPaidMediaRequestChatId](docs/PostSendPaidMediaRequestChatId.md)
 - [PostSendPhotoRequestPhoto](docs/PostSendPhotoRequestPhoto.md)
 - [PostSendPollRequest](docs/PostSendPollRequest.md)
 - [PostSendStickerRequestSticker](docs/PostSendStickerRequestSticker.md)
 - [PostSendVenueRequest](docs/PostSendVenueRequest.md)
 - [PostSendVideoNoteRequestVideoNote](docs/PostSendVideoNoteRequestVideoNote.md)
 - [PostSendVideoRequestCover](docs/PostSendVideoRequestCover.md)
 - [PostSendVideoRequestVideo](docs/PostSendVideoRequestVideo.md)
 - [PostSendVoiceRequestVoice](docs/PostSendVoiceRequestVoice.md)
 - [PostSetBusinessAccountBioRequest](docs/PostSetBusinessAccountBioRequest.md)
 - [PostSetBusinessAccountGiftSettingsRequest](docs/PostSetBusinessAccountGiftSettingsRequest.md)
 - [PostSetBusinessAccountNameRequest](docs/PostSetBusinessAccountNameRequest.md)
 - [PostSetBusinessAccountUsernameRequest](docs/PostSetBusinessAccountUsernameRequest.md)
 - [PostSetChatAdministratorCustomTitleRequest](docs/PostSetChatAdministratorCustomTitleRequest.md)
 - [PostSetChatDescriptionRequest](docs/PostSetChatDescriptionRequest.md)
 - [PostSetChatMenuButtonRequest](docs/PostSetChatMenuButtonRequest.md)
 - [PostSetChatPermissionsRequest](docs/PostSetChatPermissionsRequest.md)
 - [PostSetChatStickerSetRequest](docs/PostSetChatStickerSetRequest.md)
 - [PostSetChatTitleRequest](docs/PostSetChatTitleRequest.md)
 - [PostSetCustomEmojiStickerSetThumbnailRequest](docs/PostSetCustomEmojiStickerSetThumbnailRequest.md)
 - [PostSetGameScoreRequest](docs/PostSetGameScoreRequest.md)
 - [PostSetMessageReactionRequest](docs/PostSetMessageReactionRequest.md)
 - [PostSetMyCommandsRequest](docs/PostSetMyCommandsRequest.md)
 - [PostSetMyDefaultAdministratorRightsRequest](docs/PostSetMyDefaultAdministratorRightsRequest.md)
 - [PostSetMyDescriptionRequest](docs/PostSetMyDescriptionRequest.md)
 - [PostSetMyNameRequest](docs/PostSetMyNameRequest.md)
 - [PostSetMyShortDescriptionRequest](docs/PostSetMyShortDescriptionRequest.md)
 - [PostSetPassportDataErrorsRequest](docs/PostSetPassportDataErrorsRequest.md)
 - [PostSetStickerEmojiListRequest](docs/PostSetStickerEmojiListRequest.md)
 - [PostSetStickerKeywordsRequest](docs/PostSetStickerKeywordsRequest.md)
 - [PostSetStickerMaskPositionRequest](docs/PostSetStickerMaskPositionRequest.md)
 - [PostSetStickerPositionInSetRequest](docs/PostSetStickerPositionInSetRequest.md)
 - [PostSetStickerSetThumbnailRequestThumbnail](docs/PostSetStickerSetThumbnailRequestThumbnail.md)
 - [PostSetStickerSetTitleRequest](docs/PostSetStickerSetTitleRequest.md)
 - [PostSetUserEmojiStatusRequest](docs/PostSetUserEmojiStatusRequest.md)
 - [PostSetWebhook200Response](docs/PostSetWebhook200Response.md)
 - [PostStopMessageLiveLocationRequest](docs/PostStopMessageLiveLocationRequest.md)
 - [PostStopPoll200Response](docs/PostStopPoll200Response.md)
 - [PostStopPollRequest](docs/PostStopPollRequest.md)
 - [PostTransferBusinessAccountStarsRequest](docs/PostTransferBusinessAccountStarsRequest.md)
 - [PostTransferGiftRequest](docs/PostTransferGiftRequest.md)
 - [PostUnbanChatMemberRequest](docs/PostUnbanChatMemberRequest.md)
 - [PostUnpinChatMessageRequest](docs/PostUnpinChatMessageRequest.md)
 - [PostUpgradeGiftRequest](docs/PostUpgradeGiftRequest.md)
 - [PostVerifyChatRequest](docs/PostVerifyChatRequest.md)
 - [PostVerifyUserRequest](docs/PostVerifyUserRequest.md)
 - [PreCheckoutQuery](docs/PreCheckoutQuery.md)
 - [PreparedInlineMessage](docs/PreparedInlineMessage.md)
 - [ProximityAlertTriggered](docs/ProximityAlertTriggered.md)
 - [ReactionCount](docs/ReactionCount.md)
 - [ReactionType](docs/ReactionType.md)
 - [ReactionTypeCustomEmoji](docs/ReactionTypeCustomEmoji.md)
 - [ReactionTypeEmoji](docs/ReactionTypeEmoji.md)
 - [ReactionTypePaid](docs/ReactionTypePaid.md)
 - [RefundedPayment](docs/RefundedPayment.md)
 - [ReplyKeyboardMarkup](docs/ReplyKeyboardMarkup.md)
 - [ReplyKeyboardRemove](docs/ReplyKeyboardRemove.md)
 - [ReplyParameters](docs/ReplyParameters.md)
 - [ReplyParametersChatId](docs/ReplyParametersChatId.md)
 - [ResponseParameters](docs/ResponseParameters.md)
 - [RevenueWithdrawalState](docs/RevenueWithdrawalState.md)
 - [RevenueWithdrawalStateFailed](docs/RevenueWithdrawalStateFailed.md)
 - [RevenueWithdrawalStatePending](docs/RevenueWithdrawalStatePending.md)
 - [RevenueWithdrawalStateSucceeded](docs/RevenueWithdrawalStateSucceeded.md)
 - [SentWebAppMessage](docs/SentWebAppMessage.md)
 - [SharedUser](docs/SharedUser.md)
 - [ShippingAddress](docs/ShippingAddress.md)
 - [ShippingOption](docs/ShippingOption.md)
 - [ShippingQuery](docs/ShippingQuery.md)
 - [StarAmount](docs/StarAmount.md)
 - [StarTransaction](docs/StarTransaction.md)
 - [StarTransactions](docs/StarTransactions.md)
 - [Sticker](docs/Sticker.md)
 - [StickerSet](docs/StickerSet.md)
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
 - [UniqueGift](docs/UniqueGift.md)
 - [UniqueGiftBackdrop](docs/UniqueGiftBackdrop.md)
 - [UniqueGiftBackdropColors](docs/UniqueGiftBackdropColors.md)
 - [UniqueGiftInfo](docs/UniqueGiftInfo.md)
 - [UniqueGiftModel](docs/UniqueGiftModel.md)
 - [UniqueGiftSymbol](docs/UniqueGiftSymbol.md)
 - [Update](docs/Update.md)
 - [User](docs/User.md)
 - [UserChatBoosts](docs/UserChatBoosts.md)
 - [UserProfilePhotos](docs/UserProfilePhotos.md)
 - [UsersShared](docs/UsersShared.md)
 - [Venue](docs/Venue.md)
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

