# KeyboardButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed | 
**RequestUsers** | Pointer to [**KeyboardButtonRequestUsers**](KeyboardButtonRequestUsers.md) |  | [optional] 
**RequestChat** | Pointer to [**KeyboardButtonRequestChat**](KeyboardButtonRequestChat.md) |  | [optional] 
**RequestContact** | Pointer to **bool** | *Optional*. If *True*, the user&#39;s phone number will be sent as a contact when the button is pressed. Available in private chats only. | [optional] 
**RequestLocation** | Pointer to **bool** | *Optional*. If *True*, the user&#39;s current location will be sent when the button is pressed. Available in private chats only. | [optional] 
**RequestPoll** | Pointer to [**KeyboardButtonPollType**](KeyboardButtonPollType.md) |  | [optional] 
**WebApp** | Pointer to [**WebAppInfo**](WebAppInfo.md) |  | [optional] 

## Methods

### NewKeyboardButton

`func NewKeyboardButton(text string, ) *KeyboardButton`

NewKeyboardButton instantiates a new KeyboardButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyboardButtonWithDefaults

`func NewKeyboardButtonWithDefaults() *KeyboardButton`

NewKeyboardButtonWithDefaults instantiates a new KeyboardButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *KeyboardButton) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *KeyboardButton) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *KeyboardButton) SetText(v string)`

SetText sets Text field to given value.


### GetRequestUsers

`func (o *KeyboardButton) GetRequestUsers() KeyboardButtonRequestUsers`

GetRequestUsers returns the RequestUsers field if non-nil, zero value otherwise.

### GetRequestUsersOk

`func (o *KeyboardButton) GetRequestUsersOk() (*KeyboardButtonRequestUsers, bool)`

GetRequestUsersOk returns a tuple with the RequestUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUsers

`func (o *KeyboardButton) SetRequestUsers(v KeyboardButtonRequestUsers)`

SetRequestUsers sets RequestUsers field to given value.

### HasRequestUsers

`func (o *KeyboardButton) HasRequestUsers() bool`

HasRequestUsers returns a boolean if a field has been set.

### GetRequestChat

`func (o *KeyboardButton) GetRequestChat() KeyboardButtonRequestChat`

GetRequestChat returns the RequestChat field if non-nil, zero value otherwise.

### GetRequestChatOk

`func (o *KeyboardButton) GetRequestChatOk() (*KeyboardButtonRequestChat, bool)`

GetRequestChatOk returns a tuple with the RequestChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestChat

`func (o *KeyboardButton) SetRequestChat(v KeyboardButtonRequestChat)`

SetRequestChat sets RequestChat field to given value.

### HasRequestChat

`func (o *KeyboardButton) HasRequestChat() bool`

HasRequestChat returns a boolean if a field has been set.

### GetRequestContact

`func (o *KeyboardButton) GetRequestContact() bool`

GetRequestContact returns the RequestContact field if non-nil, zero value otherwise.

### GetRequestContactOk

`func (o *KeyboardButton) GetRequestContactOk() (*bool, bool)`

GetRequestContactOk returns a tuple with the RequestContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContact

`func (o *KeyboardButton) SetRequestContact(v bool)`

SetRequestContact sets RequestContact field to given value.

### HasRequestContact

`func (o *KeyboardButton) HasRequestContact() bool`

HasRequestContact returns a boolean if a field has been set.

### GetRequestLocation

`func (o *KeyboardButton) GetRequestLocation() bool`

GetRequestLocation returns the RequestLocation field if non-nil, zero value otherwise.

### GetRequestLocationOk

`func (o *KeyboardButton) GetRequestLocationOk() (*bool, bool)`

GetRequestLocationOk returns a tuple with the RequestLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLocation

`func (o *KeyboardButton) SetRequestLocation(v bool)`

SetRequestLocation sets RequestLocation field to given value.

### HasRequestLocation

`func (o *KeyboardButton) HasRequestLocation() bool`

HasRequestLocation returns a boolean if a field has been set.

### GetRequestPoll

`func (o *KeyboardButton) GetRequestPoll() KeyboardButtonPollType`

GetRequestPoll returns the RequestPoll field if non-nil, zero value otherwise.

### GetRequestPollOk

`func (o *KeyboardButton) GetRequestPollOk() (*KeyboardButtonPollType, bool)`

GetRequestPollOk returns a tuple with the RequestPoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPoll

`func (o *KeyboardButton) SetRequestPoll(v KeyboardButtonPollType)`

SetRequestPoll sets RequestPoll field to given value.

### HasRequestPoll

`func (o *KeyboardButton) HasRequestPoll() bool`

HasRequestPoll returns a boolean if a field has been set.

### GetWebApp

`func (o *KeyboardButton) GetWebApp() WebAppInfo`

GetWebApp returns the WebApp field if non-nil, zero value otherwise.

### GetWebAppOk

`func (o *KeyboardButton) GetWebAppOk() (*WebAppInfo, bool)`

GetWebAppOk returns a tuple with the WebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApp

`func (o *KeyboardButton) SetWebApp(v WebAppInfo)`

SetWebApp sets WebApp field to given value.

### HasWebApp

`func (o *KeyboardButton) HasWebApp() bool`

HasWebApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


