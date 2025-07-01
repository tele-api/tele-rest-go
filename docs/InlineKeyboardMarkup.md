# InlineKeyboardMarkup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineKeyboard** | [**[][]InlineKeyboardButton**]([]InlineKeyboardButton.md) | Array of button rows, each represented by an Array of [InlineKeyboardButton](https://core.telegram.org/bots/api/#inlinekeyboardbutton) objects | 

## Methods

### NewInlineKeyboardMarkup

`func NewInlineKeyboardMarkup(inlineKeyboard [][]InlineKeyboardButton, ) *InlineKeyboardMarkup`

NewInlineKeyboardMarkup instantiates a new InlineKeyboardMarkup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineKeyboardMarkupWithDefaults

`func NewInlineKeyboardMarkupWithDefaults() *InlineKeyboardMarkup`

NewInlineKeyboardMarkupWithDefaults instantiates a new InlineKeyboardMarkup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineKeyboard

`func (o *InlineKeyboardMarkup) GetInlineKeyboard() [][]InlineKeyboardButton`

GetInlineKeyboard returns the InlineKeyboard field if non-nil, zero value otherwise.

### GetInlineKeyboardOk

`func (o *InlineKeyboardMarkup) GetInlineKeyboardOk() (*[][]InlineKeyboardButton, bool)`

GetInlineKeyboardOk returns a tuple with the InlineKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineKeyboard

`func (o *InlineKeyboardMarkup) SetInlineKeyboard(v [][]InlineKeyboardButton)`

SetInlineKeyboard sets InlineKeyboard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


