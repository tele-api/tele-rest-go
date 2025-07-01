# LoginUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in [Receiving authorization data](https://core.telegram.org/widgets/login#receiving-authorization-data).    **NOTE:** You **must** always check the hash of the received data to verify the authentication and the integrity of the data as described in [Checking authorization](https://core.telegram.org/widgets/login#checking-authorization). | 
**ForwardText** | Pointer to **string** | *Optional*. New text of the button in forwarded messages. | [optional] 
**BotUsername** | Pointer to **string** | *Optional*. Username of a bot, which will be used for user authorization. See [Setting up a bot](https://core.telegram.org/widgets/login#setting-up-a-bot) for more details. If not specified, the current bot&#39;s username will be assumed. The *url*&#39;s domain must be the same as the domain linked with the bot. See [Linking your domain to the bot](https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot) for more details. | [optional] 
**RequestWriteAccess** | Pointer to **bool** | *Optional*. Pass *True* to request the permission for your bot to send messages to the user. | [optional] 

## Methods

### NewLoginUrl

`func NewLoginUrl(url string, ) *LoginUrl`

NewLoginUrl instantiates a new LoginUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginUrlWithDefaults

`func NewLoginUrlWithDefaults() *LoginUrl`

NewLoginUrlWithDefaults instantiates a new LoginUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LoginUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoginUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoginUrl) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetForwardText

`func (o *LoginUrl) GetForwardText() string`

GetForwardText returns the ForwardText field if non-nil, zero value otherwise.

### GetForwardTextOk

`func (o *LoginUrl) GetForwardTextOk() (*string, bool)`

GetForwardTextOk returns a tuple with the ForwardText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardText

`func (o *LoginUrl) SetForwardText(v string)`

SetForwardText sets ForwardText field to given value.

### HasForwardText

`func (o *LoginUrl) HasForwardText() bool`

HasForwardText returns a boolean if a field has been set.

### GetBotUsername

`func (o *LoginUrl) GetBotUsername() string`

GetBotUsername returns the BotUsername field if non-nil, zero value otherwise.

### GetBotUsernameOk

`func (o *LoginUrl) GetBotUsernameOk() (*string, bool)`

GetBotUsernameOk returns a tuple with the BotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUsername

`func (o *LoginUrl) SetBotUsername(v string)`

SetBotUsername sets BotUsername field to given value.

### HasBotUsername

`func (o *LoginUrl) HasBotUsername() bool`

HasBotUsername returns a boolean if a field has been set.

### GetRequestWriteAccess

`func (o *LoginUrl) GetRequestWriteAccess() bool`

GetRequestWriteAccess returns the RequestWriteAccess field if non-nil, zero value otherwise.

### GetRequestWriteAccessOk

`func (o *LoginUrl) GetRequestWriteAccessOk() (*bool, bool)`

GetRequestWriteAccessOk returns a tuple with the RequestWriteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestWriteAccess

`func (o *LoginUrl) SetRequestWriteAccess(v bool)`

SetRequestWriteAccess sets RequestWriteAccess field to given value.

### HasRequestWriteAccess

`func (o *LoginUrl) HasRequestWriteAccess() bool`

HasRequestWriteAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


