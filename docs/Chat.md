# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier. | 
**Type** | **string** | Type of the chat, can be either “private”, “group”, “supergroup” or “channel” | 
**Title** | Pointer to **string** | *Optional*. Title, for supergroups, channels and group chats | [optional] 
**Username** | Pointer to **string** | *Optional*. Username, for private chats, supergroups and channels if available | [optional] 
**FirstName** | Pointer to **string** | *Optional*. First name of the other party in a private chat | [optional] 
**LastName** | Pointer to **string** | *Optional*. Last name of the other party in a private chat | [optional] 
**IsForum** | Pointer to **bool** | *Optional*. *True*, if the supergroup chat is a forum (has [topics](https://telegram.org/blog/topics-in-groups-collectible-usernames#topics-in-groups) enabled) | [optional] [default to true]

## Methods

### NewChat

`func NewChat(id int32, type_ string, ) *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Chat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Chat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Chat) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *Chat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Chat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Chat) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *Chat) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Chat) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Chat) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Chat) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsername

`func (o *Chat) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Chat) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Chat) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Chat) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *Chat) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Chat) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Chat) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Chat) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Chat) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Chat) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Chat) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Chat) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetIsForum

`func (o *Chat) GetIsForum() bool`

GetIsForum returns the IsForum field if non-nil, zero value otherwise.

### GetIsForumOk

`func (o *Chat) GetIsForumOk() (*bool, bool)`

GetIsForumOk returns a tuple with the IsForum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForum

`func (o *Chat) SetIsForum(v bool)`

SetIsForum sets IsForum field to given value.

### HasIsForum

`func (o *Chat) HasIsForum() bool`

HasIsForum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


