# KeyboardButtonRequestChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **int32** | Signed 32-bit identifier of the request, which will be received back in the [ChatShared](https://core.telegram.org/bots/api/#chatshared) object. Must be unique within the message | 
**ChatIsChannel** | **bool** | Pass *True* to request a channel chat, pass *False* to request a group or a supergroup chat. | 
**ChatIsForum** | Pointer to **bool** | *Optional*. Pass *True* to request a forum supergroup, pass *False* to request a non-forum chat. If not specified, no additional restrictions are applied. | [optional] 
**ChatHasUsername** | Pointer to **bool** | *Optional*. Pass *True* to request a supergroup or a channel with a username, pass *False* to request a chat without a username. If not specified, no additional restrictions are applied. | [optional] 
**ChatIsCreated** | Pointer to **bool** | *Optional*. Pass *True* to request a chat owned by the user. Otherwise, no additional restrictions are applied. | [optional] 
**UserAdministratorRights** | Pointer to [**ChatAdministratorRights**](ChatAdministratorRights.md) |  | [optional] 
**BotAdministratorRights** | Pointer to [**ChatAdministratorRights**](ChatAdministratorRights.md) |  | [optional] 
**BotIsMember** | Pointer to **bool** | *Optional*. Pass *True* to request a chat with the bot as a member. Otherwise, no additional restrictions are applied. | [optional] 
**RequestTitle** | Pointer to **bool** | *Optional*. Pass *True* to request the chat&#39;s title | [optional] 
**RequestUsername** | Pointer to **bool** | *Optional*. Pass *True* to request the chat&#39;s username | [optional] 
**RequestPhoto** | Pointer to **bool** | *Optional*. Pass *True* to request the chat&#39;s photo | [optional] 

## Methods

### NewKeyboardButtonRequestChat

`func NewKeyboardButtonRequestChat(requestId int32, chatIsChannel bool, ) *KeyboardButtonRequestChat`

NewKeyboardButtonRequestChat instantiates a new KeyboardButtonRequestChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyboardButtonRequestChatWithDefaults

`func NewKeyboardButtonRequestChatWithDefaults() *KeyboardButtonRequestChat`

NewKeyboardButtonRequestChatWithDefaults instantiates a new KeyboardButtonRequestChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *KeyboardButtonRequestChat) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *KeyboardButtonRequestChat) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *KeyboardButtonRequestChat) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.


### GetChatIsChannel

`func (o *KeyboardButtonRequestChat) GetChatIsChannel() bool`

GetChatIsChannel returns the ChatIsChannel field if non-nil, zero value otherwise.

### GetChatIsChannelOk

`func (o *KeyboardButtonRequestChat) GetChatIsChannelOk() (*bool, bool)`

GetChatIsChannelOk returns a tuple with the ChatIsChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatIsChannel

`func (o *KeyboardButtonRequestChat) SetChatIsChannel(v bool)`

SetChatIsChannel sets ChatIsChannel field to given value.


### GetChatIsForum

`func (o *KeyboardButtonRequestChat) GetChatIsForum() bool`

GetChatIsForum returns the ChatIsForum field if non-nil, zero value otherwise.

### GetChatIsForumOk

`func (o *KeyboardButtonRequestChat) GetChatIsForumOk() (*bool, bool)`

GetChatIsForumOk returns a tuple with the ChatIsForum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatIsForum

`func (o *KeyboardButtonRequestChat) SetChatIsForum(v bool)`

SetChatIsForum sets ChatIsForum field to given value.

### HasChatIsForum

`func (o *KeyboardButtonRequestChat) HasChatIsForum() bool`

HasChatIsForum returns a boolean if a field has been set.

### GetChatHasUsername

`func (o *KeyboardButtonRequestChat) GetChatHasUsername() bool`

GetChatHasUsername returns the ChatHasUsername field if non-nil, zero value otherwise.

### GetChatHasUsernameOk

`func (o *KeyboardButtonRequestChat) GetChatHasUsernameOk() (*bool, bool)`

GetChatHasUsernameOk returns a tuple with the ChatHasUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatHasUsername

`func (o *KeyboardButtonRequestChat) SetChatHasUsername(v bool)`

SetChatHasUsername sets ChatHasUsername field to given value.

### HasChatHasUsername

`func (o *KeyboardButtonRequestChat) HasChatHasUsername() bool`

HasChatHasUsername returns a boolean if a field has been set.

### GetChatIsCreated

`func (o *KeyboardButtonRequestChat) GetChatIsCreated() bool`

GetChatIsCreated returns the ChatIsCreated field if non-nil, zero value otherwise.

### GetChatIsCreatedOk

`func (o *KeyboardButtonRequestChat) GetChatIsCreatedOk() (*bool, bool)`

GetChatIsCreatedOk returns a tuple with the ChatIsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatIsCreated

`func (o *KeyboardButtonRequestChat) SetChatIsCreated(v bool)`

SetChatIsCreated sets ChatIsCreated field to given value.

### HasChatIsCreated

`func (o *KeyboardButtonRequestChat) HasChatIsCreated() bool`

HasChatIsCreated returns a boolean if a field has been set.

### GetUserAdministratorRights

`func (o *KeyboardButtonRequestChat) GetUserAdministratorRights() ChatAdministratorRights`

GetUserAdministratorRights returns the UserAdministratorRights field if non-nil, zero value otherwise.

### GetUserAdministratorRightsOk

`func (o *KeyboardButtonRequestChat) GetUserAdministratorRightsOk() (*ChatAdministratorRights, bool)`

GetUserAdministratorRightsOk returns a tuple with the UserAdministratorRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAdministratorRights

`func (o *KeyboardButtonRequestChat) SetUserAdministratorRights(v ChatAdministratorRights)`

SetUserAdministratorRights sets UserAdministratorRights field to given value.

### HasUserAdministratorRights

`func (o *KeyboardButtonRequestChat) HasUserAdministratorRights() bool`

HasUserAdministratorRights returns a boolean if a field has been set.

### GetBotAdministratorRights

`func (o *KeyboardButtonRequestChat) GetBotAdministratorRights() ChatAdministratorRights`

GetBotAdministratorRights returns the BotAdministratorRights field if non-nil, zero value otherwise.

### GetBotAdministratorRightsOk

`func (o *KeyboardButtonRequestChat) GetBotAdministratorRightsOk() (*ChatAdministratorRights, bool)`

GetBotAdministratorRightsOk returns a tuple with the BotAdministratorRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotAdministratorRights

`func (o *KeyboardButtonRequestChat) SetBotAdministratorRights(v ChatAdministratorRights)`

SetBotAdministratorRights sets BotAdministratorRights field to given value.

### HasBotAdministratorRights

`func (o *KeyboardButtonRequestChat) HasBotAdministratorRights() bool`

HasBotAdministratorRights returns a boolean if a field has been set.

### GetBotIsMember

`func (o *KeyboardButtonRequestChat) GetBotIsMember() bool`

GetBotIsMember returns the BotIsMember field if non-nil, zero value otherwise.

### GetBotIsMemberOk

`func (o *KeyboardButtonRequestChat) GetBotIsMemberOk() (*bool, bool)`

GetBotIsMemberOk returns a tuple with the BotIsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotIsMember

`func (o *KeyboardButtonRequestChat) SetBotIsMember(v bool)`

SetBotIsMember sets BotIsMember field to given value.

### HasBotIsMember

`func (o *KeyboardButtonRequestChat) HasBotIsMember() bool`

HasBotIsMember returns a boolean if a field has been set.

### GetRequestTitle

`func (o *KeyboardButtonRequestChat) GetRequestTitle() bool`

GetRequestTitle returns the RequestTitle field if non-nil, zero value otherwise.

### GetRequestTitleOk

`func (o *KeyboardButtonRequestChat) GetRequestTitleOk() (*bool, bool)`

GetRequestTitleOk returns a tuple with the RequestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTitle

`func (o *KeyboardButtonRequestChat) SetRequestTitle(v bool)`

SetRequestTitle sets RequestTitle field to given value.

### HasRequestTitle

`func (o *KeyboardButtonRequestChat) HasRequestTitle() bool`

HasRequestTitle returns a boolean if a field has been set.

### GetRequestUsername

`func (o *KeyboardButtonRequestChat) GetRequestUsername() bool`

GetRequestUsername returns the RequestUsername field if non-nil, zero value otherwise.

### GetRequestUsernameOk

`func (o *KeyboardButtonRequestChat) GetRequestUsernameOk() (*bool, bool)`

GetRequestUsernameOk returns a tuple with the RequestUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUsername

`func (o *KeyboardButtonRequestChat) SetRequestUsername(v bool)`

SetRequestUsername sets RequestUsername field to given value.

### HasRequestUsername

`func (o *KeyboardButtonRequestChat) HasRequestUsername() bool`

HasRequestUsername returns a boolean if a field has been set.

### GetRequestPhoto

`func (o *KeyboardButtonRequestChat) GetRequestPhoto() bool`

GetRequestPhoto returns the RequestPhoto field if non-nil, zero value otherwise.

### GetRequestPhotoOk

`func (o *KeyboardButtonRequestChat) GetRequestPhotoOk() (*bool, bool)`

GetRequestPhotoOk returns a tuple with the RequestPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPhoto

`func (o *KeyboardButtonRequestChat) SetRequestPhoto(v bool)`

SetRequestPhoto sets RequestPhoto field to given value.

### HasRequestPhoto

`func (o *KeyboardButtonRequestChat) HasRequestPhoto() bool`

HasRequestPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


