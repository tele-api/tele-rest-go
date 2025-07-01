# RevokeChatInviteLinkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**RevokeChatInviteLinkPostRequestChatId**](RevokeChatInviteLinkPostRequestChatId.md) |  | 
**InviteLink** | **string** | The invite link to revoke | 

## Methods

### NewRevokeChatInviteLinkPostRequest

`func NewRevokeChatInviteLinkPostRequest(chatId RevokeChatInviteLinkPostRequestChatId, inviteLink string, ) *RevokeChatInviteLinkPostRequest`

NewRevokeChatInviteLinkPostRequest instantiates a new RevokeChatInviteLinkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeChatInviteLinkPostRequestWithDefaults

`func NewRevokeChatInviteLinkPostRequestWithDefaults() *RevokeChatInviteLinkPostRequest`

NewRevokeChatInviteLinkPostRequestWithDefaults instantiates a new RevokeChatInviteLinkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *RevokeChatInviteLinkPostRequest) GetChatId() RevokeChatInviteLinkPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *RevokeChatInviteLinkPostRequest) GetChatIdOk() (*RevokeChatInviteLinkPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *RevokeChatInviteLinkPostRequest) SetChatId(v RevokeChatInviteLinkPostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetInviteLink

`func (o *RevokeChatInviteLinkPostRequest) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *RevokeChatInviteLinkPostRequest) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *RevokeChatInviteLinkPostRequest) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


