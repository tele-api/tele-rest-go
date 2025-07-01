# ChatInviteLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteLink** | **string** | The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”. | 
**Creator** | [**User**](User.md) |  | 
**CreatesJoinRequest** | **bool** | *True*, if users joining the chat via the link need to be approved by chat administrators | 
**IsPrimary** | **bool** | *True*, if the link is primary | 
**IsRevoked** | **bool** | *True*, if the link is revoked | 
**Name** | Pointer to **string** | *Optional*. Invite link name | [optional] 
**ExpireDate** | Pointer to **int32** | *Optional*. Point in time (Unix timestamp) when the link will expire or has been expired | [optional] 
**MemberLimit** | Pointer to **int32** | *Optional*. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999 | [optional] 
**PendingJoinRequestCount** | Pointer to **int32** | *Optional*. Number of pending join requests created using this link | [optional] 
**SubscriptionPeriod** | Pointer to **int32** | *Optional*. The number of seconds the subscription will be active for before the next payment | [optional] 
**SubscriptionPrice** | Pointer to **int32** | *Optional*. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link | [optional] 

## Methods

### NewChatInviteLink

`func NewChatInviteLink(inviteLink string, creator User, createsJoinRequest bool, isPrimary bool, isRevoked bool, ) *ChatInviteLink`

NewChatInviteLink instantiates a new ChatInviteLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInviteLinkWithDefaults

`func NewChatInviteLinkWithDefaults() *ChatInviteLink`

NewChatInviteLinkWithDefaults instantiates a new ChatInviteLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviteLink

`func (o *ChatInviteLink) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *ChatInviteLink) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *ChatInviteLink) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.


### GetCreator

`func (o *ChatInviteLink) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ChatInviteLink) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ChatInviteLink) SetCreator(v User)`

SetCreator sets Creator field to given value.


### GetCreatesJoinRequest

`func (o *ChatInviteLink) GetCreatesJoinRequest() bool`

GetCreatesJoinRequest returns the CreatesJoinRequest field if non-nil, zero value otherwise.

### GetCreatesJoinRequestOk

`func (o *ChatInviteLink) GetCreatesJoinRequestOk() (*bool, bool)`

GetCreatesJoinRequestOk returns a tuple with the CreatesJoinRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatesJoinRequest

`func (o *ChatInviteLink) SetCreatesJoinRequest(v bool)`

SetCreatesJoinRequest sets CreatesJoinRequest field to given value.


### GetIsPrimary

`func (o *ChatInviteLink) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ChatInviteLink) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ChatInviteLink) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetIsRevoked

`func (o *ChatInviteLink) GetIsRevoked() bool`

GetIsRevoked returns the IsRevoked field if non-nil, zero value otherwise.

### GetIsRevokedOk

`func (o *ChatInviteLink) GetIsRevokedOk() (*bool, bool)`

GetIsRevokedOk returns a tuple with the IsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevoked

`func (o *ChatInviteLink) SetIsRevoked(v bool)`

SetIsRevoked sets IsRevoked field to given value.


### GetName

`func (o *ChatInviteLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatInviteLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatInviteLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatInviteLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpireDate

`func (o *ChatInviteLink) GetExpireDate() int32`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *ChatInviteLink) GetExpireDateOk() (*int32, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *ChatInviteLink) SetExpireDate(v int32)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *ChatInviteLink) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetMemberLimit

`func (o *ChatInviteLink) GetMemberLimit() int32`

GetMemberLimit returns the MemberLimit field if non-nil, zero value otherwise.

### GetMemberLimitOk

`func (o *ChatInviteLink) GetMemberLimitOk() (*int32, bool)`

GetMemberLimitOk returns a tuple with the MemberLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLimit

`func (o *ChatInviteLink) SetMemberLimit(v int32)`

SetMemberLimit sets MemberLimit field to given value.

### HasMemberLimit

`func (o *ChatInviteLink) HasMemberLimit() bool`

HasMemberLimit returns a boolean if a field has been set.

### GetPendingJoinRequestCount

`func (o *ChatInviteLink) GetPendingJoinRequestCount() int32`

GetPendingJoinRequestCount returns the PendingJoinRequestCount field if non-nil, zero value otherwise.

### GetPendingJoinRequestCountOk

`func (o *ChatInviteLink) GetPendingJoinRequestCountOk() (*int32, bool)`

GetPendingJoinRequestCountOk returns a tuple with the PendingJoinRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingJoinRequestCount

`func (o *ChatInviteLink) SetPendingJoinRequestCount(v int32)`

SetPendingJoinRequestCount sets PendingJoinRequestCount field to given value.

### HasPendingJoinRequestCount

`func (o *ChatInviteLink) HasPendingJoinRequestCount() bool`

HasPendingJoinRequestCount returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *ChatInviteLink) GetSubscriptionPeriod() int32`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *ChatInviteLink) GetSubscriptionPeriodOk() (*int32, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *ChatInviteLink) SetSubscriptionPeriod(v int32)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *ChatInviteLink) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetSubscriptionPrice

`func (o *ChatInviteLink) GetSubscriptionPrice() int32`

GetSubscriptionPrice returns the SubscriptionPrice field if non-nil, zero value otherwise.

### GetSubscriptionPriceOk

`func (o *ChatInviteLink) GetSubscriptionPriceOk() (*int32, bool)`

GetSubscriptionPriceOk returns a tuple with the SubscriptionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPrice

`func (o *ChatInviteLink) SetSubscriptionPrice(v int32)`

SetSubscriptionPrice sets SubscriptionPrice field to given value.

### HasSubscriptionPrice

`func (o *ChatInviteLink) HasSubscriptionPrice() bool`

HasSubscriptionPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


