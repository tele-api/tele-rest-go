# SendMessagePostRequestReplyMarkup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineKeyboard** | [**[][]InlineKeyboardButton**]([]InlineKeyboardButton.md) | Array of button rows, each represented by an Array of [InlineKeyboardButton](https://core.telegram.org/bots/api/#inlinekeyboardbutton) objects | 
**Keyboard** | [**[][]KeyboardButton**]([]KeyboardButton.md) | Array of button rows, each represented by an Array of [KeyboardButton](https://core.telegram.org/bots/api/#keyboardbutton) objects | 
**IsPersistent** | Pointer to **bool** | *Optional*. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to *false*, in which case the custom keyboard can be hidden and opened with a keyboard icon. | [optional] [default to false]
**ResizeKeyboard** | Pointer to **bool** | *Optional*. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to *false*, in which case the custom keyboard is always of the same height as the app&#39;s standard keyboard. | [optional] [default to false]
**OneTimeKeyboard** | Pointer to **bool** | *Optional*. Requests clients to hide the keyboard as soon as it&#39;s been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to *false*. | [optional] [default to false]
**InputFieldPlaceholder** | Pointer to **string** | *Optional*. The placeholder to be shown in the input field when the reply is active; 1-64 characters | [optional] 
**Selective** | Pointer to **bool** | *Optional*. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot&#39;s message is a reply to a message in the same chat and forum topic, sender of the original message. | [optional] 
**RemoveKeyboard** | **bool** | Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use *one\\_time\\_keyboard* in [ReplyKeyboardMarkup](https://core.telegram.org/bots/api/#replykeyboardmarkup)) | [default to true]
**ForceReply** | **bool** | Shows reply interface to the user, as if they manually selected the bot&#39;s message and tapped &#39;Reply&#39; | [default to true]

## Methods

### NewSendMessagePostRequestReplyMarkup

`func NewSendMessagePostRequestReplyMarkup(inlineKeyboard [][]InlineKeyboardButton, keyboard [][]KeyboardButton, removeKeyboard bool, forceReply bool, ) *SendMessagePostRequestReplyMarkup`

NewSendMessagePostRequestReplyMarkup instantiates a new SendMessagePostRequestReplyMarkup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessagePostRequestReplyMarkupWithDefaults

`func NewSendMessagePostRequestReplyMarkupWithDefaults() *SendMessagePostRequestReplyMarkup`

NewSendMessagePostRequestReplyMarkupWithDefaults instantiates a new SendMessagePostRequestReplyMarkup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineKeyboard

`func (o *SendMessagePostRequestReplyMarkup) GetInlineKeyboard() [][]InlineKeyboardButton`

GetInlineKeyboard returns the InlineKeyboard field if non-nil, zero value otherwise.

### GetInlineKeyboardOk

`func (o *SendMessagePostRequestReplyMarkup) GetInlineKeyboardOk() (*[][]InlineKeyboardButton, bool)`

GetInlineKeyboardOk returns a tuple with the InlineKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineKeyboard

`func (o *SendMessagePostRequestReplyMarkup) SetInlineKeyboard(v [][]InlineKeyboardButton)`

SetInlineKeyboard sets InlineKeyboard field to given value.


### GetKeyboard

`func (o *SendMessagePostRequestReplyMarkup) GetKeyboard() [][]KeyboardButton`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *SendMessagePostRequestReplyMarkup) GetKeyboardOk() (*[][]KeyboardButton, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *SendMessagePostRequestReplyMarkup) SetKeyboard(v [][]KeyboardButton)`

SetKeyboard sets Keyboard field to given value.


### GetIsPersistent

`func (o *SendMessagePostRequestReplyMarkup) GetIsPersistent() bool`

GetIsPersistent returns the IsPersistent field if non-nil, zero value otherwise.

### GetIsPersistentOk

`func (o *SendMessagePostRequestReplyMarkup) GetIsPersistentOk() (*bool, bool)`

GetIsPersistentOk returns a tuple with the IsPersistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersistent

`func (o *SendMessagePostRequestReplyMarkup) SetIsPersistent(v bool)`

SetIsPersistent sets IsPersistent field to given value.

### HasIsPersistent

`func (o *SendMessagePostRequestReplyMarkup) HasIsPersistent() bool`

HasIsPersistent returns a boolean if a field has been set.

### GetResizeKeyboard

`func (o *SendMessagePostRequestReplyMarkup) GetResizeKeyboard() bool`

GetResizeKeyboard returns the ResizeKeyboard field if non-nil, zero value otherwise.

### GetResizeKeyboardOk

`func (o *SendMessagePostRequestReplyMarkup) GetResizeKeyboardOk() (*bool, bool)`

GetResizeKeyboardOk returns a tuple with the ResizeKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeKeyboard

`func (o *SendMessagePostRequestReplyMarkup) SetResizeKeyboard(v bool)`

SetResizeKeyboard sets ResizeKeyboard field to given value.

### HasResizeKeyboard

`func (o *SendMessagePostRequestReplyMarkup) HasResizeKeyboard() bool`

HasResizeKeyboard returns a boolean if a field has been set.

### GetOneTimeKeyboard

`func (o *SendMessagePostRequestReplyMarkup) GetOneTimeKeyboard() bool`

GetOneTimeKeyboard returns the OneTimeKeyboard field if non-nil, zero value otherwise.

### GetOneTimeKeyboardOk

`func (o *SendMessagePostRequestReplyMarkup) GetOneTimeKeyboardOk() (*bool, bool)`

GetOneTimeKeyboardOk returns a tuple with the OneTimeKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeKeyboard

`func (o *SendMessagePostRequestReplyMarkup) SetOneTimeKeyboard(v bool)`

SetOneTimeKeyboard sets OneTimeKeyboard field to given value.

### HasOneTimeKeyboard

`func (o *SendMessagePostRequestReplyMarkup) HasOneTimeKeyboard() bool`

HasOneTimeKeyboard returns a boolean if a field has been set.

### GetInputFieldPlaceholder

`func (o *SendMessagePostRequestReplyMarkup) GetInputFieldPlaceholder() string`

GetInputFieldPlaceholder returns the InputFieldPlaceholder field if non-nil, zero value otherwise.

### GetInputFieldPlaceholderOk

`func (o *SendMessagePostRequestReplyMarkup) GetInputFieldPlaceholderOk() (*string, bool)`

GetInputFieldPlaceholderOk returns a tuple with the InputFieldPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldPlaceholder

`func (o *SendMessagePostRequestReplyMarkup) SetInputFieldPlaceholder(v string)`

SetInputFieldPlaceholder sets InputFieldPlaceholder field to given value.

### HasInputFieldPlaceholder

`func (o *SendMessagePostRequestReplyMarkup) HasInputFieldPlaceholder() bool`

HasInputFieldPlaceholder returns a boolean if a field has been set.

### GetSelective

`func (o *SendMessagePostRequestReplyMarkup) GetSelective() bool`

GetSelective returns the Selective field if non-nil, zero value otherwise.

### GetSelectiveOk

`func (o *SendMessagePostRequestReplyMarkup) GetSelectiveOk() (*bool, bool)`

GetSelectiveOk returns a tuple with the Selective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelective

`func (o *SendMessagePostRequestReplyMarkup) SetSelective(v bool)`

SetSelective sets Selective field to given value.

### HasSelective

`func (o *SendMessagePostRequestReplyMarkup) HasSelective() bool`

HasSelective returns a boolean if a field has been set.

### GetRemoveKeyboard

`func (o *SendMessagePostRequestReplyMarkup) GetRemoveKeyboard() bool`

GetRemoveKeyboard returns the RemoveKeyboard field if non-nil, zero value otherwise.

### GetRemoveKeyboardOk

`func (o *SendMessagePostRequestReplyMarkup) GetRemoveKeyboardOk() (*bool, bool)`

GetRemoveKeyboardOk returns a tuple with the RemoveKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveKeyboard

`func (o *SendMessagePostRequestReplyMarkup) SetRemoveKeyboard(v bool)`

SetRemoveKeyboard sets RemoveKeyboard field to given value.


### GetForceReply

`func (o *SendMessagePostRequestReplyMarkup) GetForceReply() bool`

GetForceReply returns the ForceReply field if non-nil, zero value otherwise.

### GetForceReplyOk

`func (o *SendMessagePostRequestReplyMarkup) GetForceReplyOk() (*bool, bool)`

GetForceReplyOk returns a tuple with the ForceReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceReply

`func (o *SendMessagePostRequestReplyMarkup) SetForceReply(v bool)`

SetForceReply sets ForceReply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


