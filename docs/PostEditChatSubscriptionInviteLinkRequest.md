# PostEditChatSubscriptionInviteLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**InviteLink** | **string** | The invite link to edit | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 

## Methods

### NewPostEditChatSubscriptionInviteLinkRequest

`func NewPostEditChatSubscriptionInviteLinkRequest(chatId PostSendMessageRequestChatId, inviteLink string, ) *PostEditChatSubscriptionInviteLinkRequest`

NewPostEditChatSubscriptionInviteLinkRequest instantiates a new PostEditChatSubscriptionInviteLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEditChatSubscriptionInviteLinkRequestWithDefaults

`func NewPostEditChatSubscriptionInviteLinkRequestWithDefaults() *PostEditChatSubscriptionInviteLinkRequest`

NewPostEditChatSubscriptionInviteLinkRequestWithDefaults instantiates a new PostEditChatSubscriptionInviteLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostEditChatSubscriptionInviteLinkRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostEditChatSubscriptionInviteLinkRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostEditChatSubscriptionInviteLinkRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetInviteLink

`func (o *PostEditChatSubscriptionInviteLinkRequest) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *PostEditChatSubscriptionInviteLinkRequest) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *PostEditChatSubscriptionInviteLinkRequest) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.


### GetName

`func (o *PostEditChatSubscriptionInviteLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostEditChatSubscriptionInviteLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostEditChatSubscriptionInviteLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostEditChatSubscriptionInviteLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


