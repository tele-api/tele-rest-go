# ReplyKeyboardRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveKeyboard** | **bool** | Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use *one\\_time\\_keyboard* in [ReplyKeyboardMarkup](https://core.telegram.org/bots/api/#replykeyboardmarkup)) | [default to true]
**Selective** | Pointer to **bool** | *Optional*. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the *text* of the [Message](https://core.telegram.org/bots/api/#message) object; 2) if the bot&#39;s message is a reply to a message in the same chat and forum topic, sender of the original message.    *Example:* A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven&#39;t voted yet. | [optional] 

## Methods

### NewReplyKeyboardRemove

`func NewReplyKeyboardRemove(removeKeyboard bool, ) *ReplyKeyboardRemove`

NewReplyKeyboardRemove instantiates a new ReplyKeyboardRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyKeyboardRemoveWithDefaults

`func NewReplyKeyboardRemoveWithDefaults() *ReplyKeyboardRemove`

NewReplyKeyboardRemoveWithDefaults instantiates a new ReplyKeyboardRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveKeyboard

`func (o *ReplyKeyboardRemove) GetRemoveKeyboard() bool`

GetRemoveKeyboard returns the RemoveKeyboard field if non-nil, zero value otherwise.

### GetRemoveKeyboardOk

`func (o *ReplyKeyboardRemove) GetRemoveKeyboardOk() (*bool, bool)`

GetRemoveKeyboardOk returns a tuple with the RemoveKeyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveKeyboard

`func (o *ReplyKeyboardRemove) SetRemoveKeyboard(v bool)`

SetRemoveKeyboard sets RemoveKeyboard field to given value.


### GetSelective

`func (o *ReplyKeyboardRemove) GetSelective() bool`

GetSelective returns the Selective field if non-nil, zero value otherwise.

### GetSelectiveOk

`func (o *ReplyKeyboardRemove) GetSelectiveOk() (*bool, bool)`

GetSelectiveOk returns a tuple with the Selective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelective

`func (o *ReplyKeyboardRemove) SetSelective(v bool)`

SetSelective sets Selective field to given value.

### HasSelective

`func (o *ReplyKeyboardRemove) HasSelective() bool`

HasSelective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


