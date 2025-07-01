# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. | 
**IsBot** | **bool** | *True*, if this user is a bot | 
**FirstName** | **string** | User&#39;s or bot&#39;s first name | 
**LastName** | Pointer to **string** | *Optional*. User&#39;s or bot&#39;s last name | [optional] 
**Username** | Pointer to **string** | *Optional*. User&#39;s or bot&#39;s username | [optional] 
**LanguageCode** | Pointer to **string** | *Optional*. [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag) of the user&#39;s language | [optional] 
**IsPremium** | Pointer to **bool** | *Optional*. *True*, if this user is a Telegram Premium user | [optional] [default to true]
**AddedToAttachmentMenu** | Pointer to **bool** | *Optional*. *True*, if this user added the bot to the attachment menu | [optional] [default to true]
**CanJoinGroups** | Pointer to **bool** | *Optional*. *True*, if the bot can be invited to groups. Returned only in [getMe](https://core.telegram.org/bots/api/#getme). | [optional] 
**CanReadAllGroupMessages** | Pointer to **bool** | *Optional*. *True*, if [privacy mode](https://core.telegram.org/bots/features#privacy-mode) is disabled for the bot. Returned only in [getMe](https://core.telegram.org/bots/api/#getme). | [optional] 
**SupportsInlineQueries** | Pointer to **bool** | *Optional*. *True*, if the bot supports inline queries. Returned only in [getMe](https://core.telegram.org/bots/api/#getme). | [optional] 
**CanConnectToBusiness** | Pointer to **bool** | *Optional*. *True*, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in [getMe](https://core.telegram.org/bots/api/#getme). | [optional] 
**HasMainWebApp** | Pointer to **bool** | *Optional*. *True*, if the bot has a main Web App. Returned only in [getMe](https://core.telegram.org/bots/api/#getme). | [optional] 

## Methods

### NewUser

`func NewUser(id int32, isBot bool, firstName string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetIsBot

`func (o *User) GetIsBot() bool`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *User) GetIsBotOk() (*bool, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *User) SetIsBot(v bool)`

SetIsBot sets IsBot field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLanguageCode

`func (o *User) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *User) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *User) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *User) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetIsPremium

`func (o *User) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *User) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *User) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *User) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetAddedToAttachmentMenu

`func (o *User) GetAddedToAttachmentMenu() bool`

GetAddedToAttachmentMenu returns the AddedToAttachmentMenu field if non-nil, zero value otherwise.

### GetAddedToAttachmentMenuOk

`func (o *User) GetAddedToAttachmentMenuOk() (*bool, bool)`

GetAddedToAttachmentMenuOk returns a tuple with the AddedToAttachmentMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedToAttachmentMenu

`func (o *User) SetAddedToAttachmentMenu(v bool)`

SetAddedToAttachmentMenu sets AddedToAttachmentMenu field to given value.

### HasAddedToAttachmentMenu

`func (o *User) HasAddedToAttachmentMenu() bool`

HasAddedToAttachmentMenu returns a boolean if a field has been set.

### GetCanJoinGroups

`func (o *User) GetCanJoinGroups() bool`

GetCanJoinGroups returns the CanJoinGroups field if non-nil, zero value otherwise.

### GetCanJoinGroupsOk

`func (o *User) GetCanJoinGroupsOk() (*bool, bool)`

GetCanJoinGroupsOk returns a tuple with the CanJoinGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanJoinGroups

`func (o *User) SetCanJoinGroups(v bool)`

SetCanJoinGroups sets CanJoinGroups field to given value.

### HasCanJoinGroups

`func (o *User) HasCanJoinGroups() bool`

HasCanJoinGroups returns a boolean if a field has been set.

### GetCanReadAllGroupMessages

`func (o *User) GetCanReadAllGroupMessages() bool`

GetCanReadAllGroupMessages returns the CanReadAllGroupMessages field if non-nil, zero value otherwise.

### GetCanReadAllGroupMessagesOk

`func (o *User) GetCanReadAllGroupMessagesOk() (*bool, bool)`

GetCanReadAllGroupMessagesOk returns a tuple with the CanReadAllGroupMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReadAllGroupMessages

`func (o *User) SetCanReadAllGroupMessages(v bool)`

SetCanReadAllGroupMessages sets CanReadAllGroupMessages field to given value.

### HasCanReadAllGroupMessages

`func (o *User) HasCanReadAllGroupMessages() bool`

HasCanReadAllGroupMessages returns a boolean if a field has been set.

### GetSupportsInlineQueries

`func (o *User) GetSupportsInlineQueries() bool`

GetSupportsInlineQueries returns the SupportsInlineQueries field if non-nil, zero value otherwise.

### GetSupportsInlineQueriesOk

`func (o *User) GetSupportsInlineQueriesOk() (*bool, bool)`

GetSupportsInlineQueriesOk returns a tuple with the SupportsInlineQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsInlineQueries

`func (o *User) SetSupportsInlineQueries(v bool)`

SetSupportsInlineQueries sets SupportsInlineQueries field to given value.

### HasSupportsInlineQueries

`func (o *User) HasSupportsInlineQueries() bool`

HasSupportsInlineQueries returns a boolean if a field has been set.

### GetCanConnectToBusiness

`func (o *User) GetCanConnectToBusiness() bool`

GetCanConnectToBusiness returns the CanConnectToBusiness field if non-nil, zero value otherwise.

### GetCanConnectToBusinessOk

`func (o *User) GetCanConnectToBusinessOk() (*bool, bool)`

GetCanConnectToBusinessOk returns a tuple with the CanConnectToBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConnectToBusiness

`func (o *User) SetCanConnectToBusiness(v bool)`

SetCanConnectToBusiness sets CanConnectToBusiness field to given value.

### HasCanConnectToBusiness

`func (o *User) HasCanConnectToBusiness() bool`

HasCanConnectToBusiness returns a boolean if a field has been set.

### GetHasMainWebApp

`func (o *User) GetHasMainWebApp() bool`

GetHasMainWebApp returns the HasMainWebApp field if non-nil, zero value otherwise.

### GetHasMainWebAppOk

`func (o *User) GetHasMainWebAppOk() (*bool, bool)`

GetHasMainWebAppOk returns a tuple with the HasMainWebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMainWebApp

`func (o *User) SetHasMainWebApp(v bool)`

SetHasMainWebApp sets HasMainWebApp field to given value.

### HasHasMainWebApp

`func (o *User) HasHasMainWebApp() bool`

HasHasMainWebApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


