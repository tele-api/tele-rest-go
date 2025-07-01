# InlineKeyboardButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Label text on the button | 
**Url** | Pointer to **string** | *Optional*. HTTP or tg:// URL to be opened when the button is pressed. Links &#x60;tg://user?id&#x3D;&lt;user_id&gt;&#x60; can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings. | [optional] 
**CallbackData** | Pointer to **string** | *Optional*. Data to be sent in a [callback query](https://core.telegram.org/bots/api/#callbackquery) to the bot when the button is pressed, 1-64 bytes | [optional] 
**WebApp** | Pointer to [**WebAppInfo**](WebAppInfo.md) |  | [optional] 
**LoginUrl** | Pointer to [**LoginUrl**](LoginUrl.md) |  | [optional] 
**SwitchInlineQuery** | Pointer to **string** | *Optional*. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot&#39;s username and the specified inline query in the input field. May be empty, in which case just the bot&#39;s username will be inserted. Not supported for messages sent on behalf of a Telegram Business account. | [optional] 
**SwitchInlineQueryCurrentChat** | Pointer to **string** | *Optional*. If set, pressing the button will insert the bot&#39;s username and the specified inline query in the current chat&#39;s input field. May be empty, in which case only the bot&#39;s username will be inserted.    This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages sent on behalf of a Telegram Business account. | [optional] 
**SwitchInlineQueryChosenChat** | Pointer to [**SwitchInlineQueryChosenChat**](SwitchInlineQueryChosenChat.md) |  | [optional] 
**CopyText** | Pointer to [**CopyTextButton**](CopyTextButton.md) |  | [optional] 
**CallbackGame** | Pointer to **interface{}** |  | [optional] 
**Pay** | Pointer to **bool** | *Optional*. Specify *True*, to send a [Pay button](https://core.telegram.org/bots/api/#payments). Substrings “⭐” and “XTR” in the buttons&#39;s text will be replaced with a Telegram Star icon.    **NOTE:** This type of button **must** always be the first button in the first row and can only be used in invoice messages. | [optional] 

## Methods

### NewInlineKeyboardButton

`func NewInlineKeyboardButton(text string, ) *InlineKeyboardButton`

NewInlineKeyboardButton instantiates a new InlineKeyboardButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineKeyboardButtonWithDefaults

`func NewInlineKeyboardButtonWithDefaults() *InlineKeyboardButton`

NewInlineKeyboardButtonWithDefaults instantiates a new InlineKeyboardButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *InlineKeyboardButton) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InlineKeyboardButton) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InlineKeyboardButton) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *InlineKeyboardButton) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineKeyboardButton) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineKeyboardButton) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineKeyboardButton) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCallbackData

`func (o *InlineKeyboardButton) GetCallbackData() string`

GetCallbackData returns the CallbackData field if non-nil, zero value otherwise.

### GetCallbackDataOk

`func (o *InlineKeyboardButton) GetCallbackDataOk() (*string, bool)`

GetCallbackDataOk returns a tuple with the CallbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackData

`func (o *InlineKeyboardButton) SetCallbackData(v string)`

SetCallbackData sets CallbackData field to given value.

### HasCallbackData

`func (o *InlineKeyboardButton) HasCallbackData() bool`

HasCallbackData returns a boolean if a field has been set.

### GetWebApp

`func (o *InlineKeyboardButton) GetWebApp() WebAppInfo`

GetWebApp returns the WebApp field if non-nil, zero value otherwise.

### GetWebAppOk

`func (o *InlineKeyboardButton) GetWebAppOk() (*WebAppInfo, bool)`

GetWebAppOk returns a tuple with the WebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApp

`func (o *InlineKeyboardButton) SetWebApp(v WebAppInfo)`

SetWebApp sets WebApp field to given value.

### HasWebApp

`func (o *InlineKeyboardButton) HasWebApp() bool`

HasWebApp returns a boolean if a field has been set.

### GetLoginUrl

`func (o *InlineKeyboardButton) GetLoginUrl() LoginUrl`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *InlineKeyboardButton) GetLoginUrlOk() (*LoginUrl, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *InlineKeyboardButton) SetLoginUrl(v LoginUrl)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *InlineKeyboardButton) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetSwitchInlineQuery

`func (o *InlineKeyboardButton) GetSwitchInlineQuery() string`

GetSwitchInlineQuery returns the SwitchInlineQuery field if non-nil, zero value otherwise.

### GetSwitchInlineQueryOk

`func (o *InlineKeyboardButton) GetSwitchInlineQueryOk() (*string, bool)`

GetSwitchInlineQueryOk returns a tuple with the SwitchInlineQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchInlineQuery

`func (o *InlineKeyboardButton) SetSwitchInlineQuery(v string)`

SetSwitchInlineQuery sets SwitchInlineQuery field to given value.

### HasSwitchInlineQuery

`func (o *InlineKeyboardButton) HasSwitchInlineQuery() bool`

HasSwitchInlineQuery returns a boolean if a field has been set.

### GetSwitchInlineQueryCurrentChat

`func (o *InlineKeyboardButton) GetSwitchInlineQueryCurrentChat() string`

GetSwitchInlineQueryCurrentChat returns the SwitchInlineQueryCurrentChat field if non-nil, zero value otherwise.

### GetSwitchInlineQueryCurrentChatOk

`func (o *InlineKeyboardButton) GetSwitchInlineQueryCurrentChatOk() (*string, bool)`

GetSwitchInlineQueryCurrentChatOk returns a tuple with the SwitchInlineQueryCurrentChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchInlineQueryCurrentChat

`func (o *InlineKeyboardButton) SetSwitchInlineQueryCurrentChat(v string)`

SetSwitchInlineQueryCurrentChat sets SwitchInlineQueryCurrentChat field to given value.

### HasSwitchInlineQueryCurrentChat

`func (o *InlineKeyboardButton) HasSwitchInlineQueryCurrentChat() bool`

HasSwitchInlineQueryCurrentChat returns a boolean if a field has been set.

### GetSwitchInlineQueryChosenChat

`func (o *InlineKeyboardButton) GetSwitchInlineQueryChosenChat() SwitchInlineQueryChosenChat`

GetSwitchInlineQueryChosenChat returns the SwitchInlineQueryChosenChat field if non-nil, zero value otherwise.

### GetSwitchInlineQueryChosenChatOk

`func (o *InlineKeyboardButton) GetSwitchInlineQueryChosenChatOk() (*SwitchInlineQueryChosenChat, bool)`

GetSwitchInlineQueryChosenChatOk returns a tuple with the SwitchInlineQueryChosenChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchInlineQueryChosenChat

`func (o *InlineKeyboardButton) SetSwitchInlineQueryChosenChat(v SwitchInlineQueryChosenChat)`

SetSwitchInlineQueryChosenChat sets SwitchInlineQueryChosenChat field to given value.

### HasSwitchInlineQueryChosenChat

`func (o *InlineKeyboardButton) HasSwitchInlineQueryChosenChat() bool`

HasSwitchInlineQueryChosenChat returns a boolean if a field has been set.

### GetCopyText

`func (o *InlineKeyboardButton) GetCopyText() CopyTextButton`

GetCopyText returns the CopyText field if non-nil, zero value otherwise.

### GetCopyTextOk

`func (o *InlineKeyboardButton) GetCopyTextOk() (*CopyTextButton, bool)`

GetCopyTextOk returns a tuple with the CopyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyText

`func (o *InlineKeyboardButton) SetCopyText(v CopyTextButton)`

SetCopyText sets CopyText field to given value.

### HasCopyText

`func (o *InlineKeyboardButton) HasCopyText() bool`

HasCopyText returns a boolean if a field has been set.

### GetCallbackGame

`func (o *InlineKeyboardButton) GetCallbackGame() interface{}`

GetCallbackGame returns the CallbackGame field if non-nil, zero value otherwise.

### GetCallbackGameOk

`func (o *InlineKeyboardButton) GetCallbackGameOk() (*interface{}, bool)`

GetCallbackGameOk returns a tuple with the CallbackGame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackGame

`func (o *InlineKeyboardButton) SetCallbackGame(v interface{})`

SetCallbackGame sets CallbackGame field to given value.

### HasCallbackGame

`func (o *InlineKeyboardButton) HasCallbackGame() bool`

HasCallbackGame returns a boolean if a field has been set.

### SetCallbackGameNil

`func (o *InlineKeyboardButton) SetCallbackGameNil(b bool)`

 SetCallbackGameNil sets the value for CallbackGame to be an explicit nil

### UnsetCallbackGame
`func (o *InlineKeyboardButton) UnsetCallbackGame()`

UnsetCallbackGame ensures that no value is present for CallbackGame, not even an explicit nil
### GetPay

`func (o *InlineKeyboardButton) GetPay() bool`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *InlineKeyboardButton) GetPayOk() (*bool, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *InlineKeyboardButton) SetPay(v bool)`

SetPay sets Pay field to given value.

### HasPay

`func (o *InlineKeyboardButton) HasPay() bool`

HasPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


