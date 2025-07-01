# EditChatSubscriptionInviteLinkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**InviteLink** | **string** | The invite link to edit | 
**Name** | Pointer to **string** | Invite link name; 0-32 characters | [optional] 

## Methods

### NewEditChatSubscriptionInviteLinkPostRequest

`func NewEditChatSubscriptionInviteLinkPostRequest(chatId SendMessagePostRequestChatId, inviteLink string, ) *EditChatSubscriptionInviteLinkPostRequest`

NewEditChatSubscriptionInviteLinkPostRequest instantiates a new EditChatSubscriptionInviteLinkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditChatSubscriptionInviteLinkPostRequestWithDefaults

`func NewEditChatSubscriptionInviteLinkPostRequestWithDefaults() *EditChatSubscriptionInviteLinkPostRequest`

NewEditChatSubscriptionInviteLinkPostRequestWithDefaults instantiates a new EditChatSubscriptionInviteLinkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *EditChatSubscriptionInviteLinkPostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditChatSubscriptionInviteLinkPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditChatSubscriptionInviteLinkPostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetInviteLink

`func (o *EditChatSubscriptionInviteLinkPostRequest) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *EditChatSubscriptionInviteLinkPostRequest) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *EditChatSubscriptionInviteLinkPostRequest) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.


### GetName

`func (o *EditChatSubscriptionInviteLinkPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditChatSubscriptionInviteLinkPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditChatSubscriptionInviteLinkPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditChatSubscriptionInviteLinkPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


