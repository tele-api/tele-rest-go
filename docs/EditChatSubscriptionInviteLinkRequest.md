# EditChatSubscriptionInviteLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**InviteLink** | **string** | The invite link to edit | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 

## Methods

### NewEditChatSubscriptionInviteLinkRequest

`func NewEditChatSubscriptionInviteLinkRequest(chatId SendMessageRequestChatId, inviteLink string, ) *EditChatSubscriptionInviteLinkRequest`

NewEditChatSubscriptionInviteLinkRequest instantiates a new EditChatSubscriptionInviteLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditChatSubscriptionInviteLinkRequestWithDefaults

`func NewEditChatSubscriptionInviteLinkRequestWithDefaults() *EditChatSubscriptionInviteLinkRequest`

NewEditChatSubscriptionInviteLinkRequestWithDefaults instantiates a new EditChatSubscriptionInviteLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *EditChatSubscriptionInviteLinkRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditChatSubscriptionInviteLinkRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditChatSubscriptionInviteLinkRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetInviteLink

`func (o *EditChatSubscriptionInviteLinkRequest) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *EditChatSubscriptionInviteLinkRequest) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *EditChatSubscriptionInviteLinkRequest) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.


### GetName

`func (o *EditChatSubscriptionInviteLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditChatSubscriptionInviteLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditChatSubscriptionInviteLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditChatSubscriptionInviteLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


