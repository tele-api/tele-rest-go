# PostSendMessageRequestReplyMarkup

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

### NewPostSendMessageRequestReplyMarkup

`func NewPostSendMessageRequestReplyMarkup(inlineKeyboard [][]InlineKeyboardButton, keyboard [][]KeyboardButton, removeKeyboard bool, forceReply bool, ) *PostSendMessageRequestReplyMarkup`

NewPostSendMessageRequestReplyMarkup instantiates a new PostSendMessageRequestReplyMarkup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSendMessageRequestReplyMarkupWithDefaults

`func NewPostSendMessageRequestReplyMarkupWithDefaults() *PostSendMessageRequestReplyMarkup`

NewPostSendMessageRequestReplyMarkupWithDefaults instantiates a new PostSendMessageRequestReplyMarkup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineKeyboard

`func (o *PostSendMessageRequestReplyMarkup) GetInlineKeyboard() [][]InlineKeyboardButton`

GetInlineKeyboard returns the InlineKeyboard field if non-nil, zero value otherwise.

### GetInlineKeyboardOk

`func (o *PostSendMessageRequestReplyMarkup) GetInlineKeyboardOk() (*[][]InlineKeyboardButton, bool)`

GetInlineKeyboardOk returns a tuple with the InlineKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineKeyboard

`func (o *PostSendMessageRequestReplyMarkup) SetInlineKeyboard(v [][]InlineKeyboardButton)`

SetInlineKeyboard sets InlineKeyboard field to given value.


### GetKeyboard

`func (o *PostSendMessageRequestReplyMarkup) GetKeyboard() [][]KeyboardButton`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *PostSendMessageRequestReplyMarkup) GetKeyboardOk() (*[][]KeyboardButton, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *PostSendMessageRequestReplyMarkup) SetKeyboard(v [][]KeyboardButton)`

SetKeyboard sets Keyboard field to given value.


### GetIsPersistent

`func (o *PostSendMessageRequestReplyMarkup) GetIsPersistent() bool`

GetIsPersistent returns the IsPersistent field if non-nil, zero value otherwise.

### GetIsPersistentOk

`func (o *PostSendMessageRequestReplyMarkup) GetIsPersistentOk() (*bool, bool)`

GetIsPersistentOk returns a tuple with the IsPersistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersistent

`func (o *PostSendMessageRequestReplyMarkup) SetIsPersistent(v bool)`

SetIsPersistent sets IsPersistent field to given value.

### HasIsPersistent

`func (o *PostSendMessageRequestReplyMarkup) HasIsPersistent() bool`

HasIsPersistent returns a boolean if a field has been set.

### GetResizeKeyboard

`func (o *PostSendMessageRequestReplyMarkup) GetResizeKeyboard() bool`

GetResizeKeyboard returns the ResizeKeyboard field if non-nil, zero value otherwise.

### GetResizeKeyboardOk

`func (o *PostSendMessageRequestReplyMarkup) GetResizeKeyboardOk() (*bool, bool)`

GetResizeKeyboardOk returns a tuple with the ResizeKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeKeyboard

`func (o *PostSendMessageRequestReplyMarkup) SetResizeKeyboard(v bool)`

SetResizeKeyboard sets ResizeKeyboard field to given value.

### HasResizeKeyboard

`func (o *PostSendMessageRequestReplyMarkup) HasResizeKeyboard() bool`

HasResizeKeyboard returns a boolean if a field has been set.

### GetOneTimeKeyboard

`func (o *PostSendMessageRequestReplyMarkup) GetOneTimeKeyboard() bool`

GetOneTimeKeyboard returns the OneTimeKeyboard field if non-nil, zero value otherwise.

### GetOneTimeKeyboardOk

`func (o *PostSendMessageRequestReplyMarkup) GetOneTimeKeyboardOk() (*bool, bool)`

GetOneTimeKeyboardOk returns a tuple with the OneTimeKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeKeyboard

`func (o *PostSendMessageRequestReplyMarkup) SetOneTimeKeyboard(v bool)`

SetOneTimeKeyboard sets OneTimeKeyboard field to given value.

### HasOneTimeKeyboard

`func (o *PostSendMessageRequestReplyMarkup) HasOneTimeKeyboard() bool`

HasOneTimeKeyboard returns a boolean if a field has been set.

### GetInputFieldPlaceholder

`func (o *PostSendMessageRequestReplyMarkup) GetInputFieldPlaceholder() string`

GetInputFieldPlaceholder returns the InputFieldPlaceholder field if non-nil, zero value otherwise.

### GetInputFieldPlaceholderOk

`func (o *PostSendMessageRequestReplyMarkup) GetInputFieldPlaceholderOk() (*string, bool)`

GetInputFieldPlaceholderOk returns a tuple with the InputFieldPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldPlaceholder

`func (o *PostSendMessageRequestReplyMarkup) SetInputFieldPlaceholder(v string)`

SetInputFieldPlaceholder sets InputFieldPlaceholder field to given value.

### HasInputFieldPlaceholder

`func (o *PostSendMessageRequestReplyMarkup) HasInputFieldPlaceholder() bool`

HasInputFieldPlaceholder returns a boolean if a field has been set.

### GetSelective

`func (o *PostSendMessageRequestReplyMarkup) GetSelective() bool`

GetSelective returns the Selective field if non-nil, zero value otherwise.

### GetSelectiveOk

`func (o *PostSendMessageRequestReplyMarkup) GetSelectiveOk() (*bool, bool)`

GetSelectiveOk returns a tuple with the Selective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelective

`func (o *PostSendMessageRequestReplyMarkup) SetSelective(v bool)`

SetSelective sets Selective field to given value.

### HasSelective

`func (o *PostSendMessageRequestReplyMarkup) HasSelective() bool`

HasSelective returns a boolean if a field has been set.

### GetRemoveKeyboard

`func (o *PostSendMessageRequestReplyMarkup) GetRemoveKeyboard() bool`

GetRemoveKeyboard returns the RemoveKeyboard field if non-nil, zero value otherwise.

### GetRemoveKeyboardOk

`func (o *PostSendMessageRequestReplyMarkup) GetRemoveKeyboardOk() (*bool, bool)`

GetRemoveKeyboardOk returns a tuple with the RemoveKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveKeyboard

`func (o *PostSendMessageRequestReplyMarkup) SetRemoveKeyboard(v bool)`

SetRemoveKeyboard sets RemoveKeyboard field to given value.


### GetForceReply

`func (o *PostSendMessageRequestReplyMarkup) GetForceReply() bool`

GetForceReply returns the ForceReply field if non-nil, zero value otherwise.

### GetForceReplyOk

`func (o *PostSendMessageRequestReplyMarkup) GetForceReplyOk() (*bool, bool)`

GetForceReplyOk returns a tuple with the ForceReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceReply

`func (o *PostSendMessageRequestReplyMarkup) SetForceReply(v bool)`

SetForceReply sets ForceReply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


