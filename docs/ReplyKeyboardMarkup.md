# ReplyKeyboardMarkup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyboard** | [**[][]KeyboardButton**]([]KeyboardButton.md) | Array of button rows, each represented by an Array of [KeyboardButton](https://core.telegram.org/bots/api/#keyboardbutton) objects | 
**IsPersistent** | Pointer to **bool** | *Optional*. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to *false*, in which case the custom keyboard can be hidden and opened with a keyboard icon. | [optional] [default to false]
**ResizeKeyboard** | Pointer to **bool** | *Optional*. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to *false*, in which case the custom keyboard is always of the same height as the app&#39;s standard keyboard. | [optional] [default to false]
**OneTimeKeyboard** | Pointer to **bool** | *Optional*. Requests clients to hide the keyboard as soon as it&#39;s been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to *false*. | [optional] [default to false]
**InputFieldPlaceholder** | Pointer to **string** | *Optional*. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters | [optional] 
**Selective** | Pointer to **bool** | *Optional*. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot&#39;s message is a reply to a message in the same chat and forum topic, sender of the original message.    *Example:* A user requests to change the bot&#39;s language, bot replies to the request with a keyboard to select the new language. Other users in the group don&#39;t see the keyboard. | [optional] 

## Methods

### NewReplyKeyboardMarkup

`func NewReplyKeyboardMarkup(keyboard [][]KeyboardButton, ) *ReplyKeyboardMarkup`

NewReplyKeyboardMarkup instantiates a new ReplyKeyboardMarkup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyKeyboardMarkupWithDefaults

`func NewReplyKeyboardMarkupWithDefaults() *ReplyKeyboardMarkup`

NewReplyKeyboardMarkupWithDefaults instantiates a new ReplyKeyboardMarkup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyboard

`func (o *ReplyKeyboardMarkup) GetKeyboard() [][]KeyboardButton`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *ReplyKeyboardMarkup) GetKeyboardOk() (*[][]KeyboardButton, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *ReplyKeyboardMarkup) SetKeyboard(v [][]KeyboardButton)`

SetKeyboard sets Keyboard field to given value.


### GetIsPersistent

`func (o *ReplyKeyboardMarkup) GetIsPersistent() bool`

GetIsPersistent returns the IsPersistent field if non-nil, zero value otherwise.

### GetIsPersistentOk

`func (o *ReplyKeyboardMarkup) GetIsPersistentOk() (*bool, bool)`

GetIsPersistentOk returns a tuple with the IsPersistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersistent

`func (o *ReplyKeyboardMarkup) SetIsPersistent(v bool)`

SetIsPersistent sets IsPersistent field to given value.

### HasIsPersistent

`func (o *ReplyKeyboardMarkup) HasIsPersistent() bool`

HasIsPersistent returns a boolean if a field has been set.

### GetResizeKeyboard

`func (o *ReplyKeyboardMarkup) GetResizeKeyboard() bool`

GetResizeKeyboard returns the ResizeKeyboard field if non-nil, zero value otherwise.

### GetResizeKeyboardOk

`func (o *ReplyKeyboardMarkup) GetResizeKeyboardOk() (*bool, bool)`

GetResizeKeyboardOk returns a tuple with the ResizeKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeKeyboard

`func (o *ReplyKeyboardMarkup) SetResizeKeyboard(v bool)`

SetResizeKeyboard sets ResizeKeyboard field to given value.

### HasResizeKeyboard

`func (o *ReplyKeyboardMarkup) HasResizeKeyboard() bool`

HasResizeKeyboard returns a boolean if a field has been set.

### GetOneTimeKeyboard

`func (o *ReplyKeyboardMarkup) GetOneTimeKeyboard() bool`

GetOneTimeKeyboard returns the OneTimeKeyboard field if non-nil, zero value otherwise.

### GetOneTimeKeyboardOk

`func (o *ReplyKeyboardMarkup) GetOneTimeKeyboardOk() (*bool, bool)`

GetOneTimeKeyboardOk returns a tuple with the OneTimeKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeKeyboard

`func (o *ReplyKeyboardMarkup) SetOneTimeKeyboard(v bool)`

SetOneTimeKeyboard sets OneTimeKeyboard field to given value.

### HasOneTimeKeyboard

`func (o *ReplyKeyboardMarkup) HasOneTimeKeyboard() bool`

HasOneTimeKeyboard returns a boolean if a field has been set.

### GetInputFieldPlaceholder

`func (o *ReplyKeyboardMarkup) GetInputFieldPlaceholder() string`

GetInputFieldPlaceholder returns the InputFieldPlaceholder field if non-nil, zero value otherwise.

### GetInputFieldPlaceholderOk

`func (o *ReplyKeyboardMarkup) GetInputFieldPlaceholderOk() (*string, bool)`

GetInputFieldPlaceholderOk returns a tuple with the InputFieldPlaceholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFieldPlaceholder

`func (o *ReplyKeyboardMarkup) SetInputFieldPlaceholder(v string)`

SetInputFieldPlaceholder sets InputFieldPlaceholder field to given value.

### HasInputFieldPlaceholder

`func (o *ReplyKeyboardMarkup) HasInputFieldPlaceholder() bool`

HasInputFieldPlaceholder returns a boolean if a field has been set.

### GetSelective

`func (o *ReplyKeyboardMarkup) GetSelective() bool`

GetSelective returns the Selective field if non-nil, zero value otherwise.

### GetSelectiveOk

`func (o *ReplyKeyboardMarkup) GetSelectiveOk() (*bool, bool)`

GetSelectiveOk returns a tuple with the Selective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelective

`func (o *ReplyKeyboardMarkup) SetSelective(v bool)`

SetSelective sets Selective field to given value.

### HasSelective

`func (o *ReplyKeyboardMarkup) HasSelective() bool`

HasSelective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


