# UnbanChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BanChatMemberRequestChatId**](BanChatMemberRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**OnlyIfBanned** | Pointer to **bool** | Do nothing if the user is not banned | [optional] 

## Methods

### NewUnbanChatMemberRequest

`func NewUnbanChatMemberRequest(chatId BanChatMemberRequestChatId, userId int32, ) *UnbanChatMemberRequest`

NewUnbanChatMemberRequest instantiates a new UnbanChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbanChatMemberRequestWithDefaults

`func NewUnbanChatMemberRequestWithDefaults() *UnbanChatMemberRequest`

NewUnbanChatMemberRequestWithDefaults instantiates a new UnbanChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *UnbanChatMemberRequest) GetChatId() BanChatMemberRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *UnbanChatMemberRequest) GetChatIdOk() (*BanChatMemberRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *UnbanChatMemberRequest) SetChatId(v BanChatMemberRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *UnbanChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnbanChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnbanChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetOnlyIfBanned

`func (o *UnbanChatMemberRequest) GetOnlyIfBanned() bool`

GetOnlyIfBanned returns the OnlyIfBanned field if non-nil, zero value otherwise.

### GetOnlyIfBannedOk

`func (o *UnbanChatMemberRequest) GetOnlyIfBannedOk() (*bool, bool)`

GetOnlyIfBannedOk returns a tuple with the OnlyIfBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIfBanned

`func (o *UnbanChatMemberRequest) SetOnlyIfBanned(v bool)`

SetOnlyIfBanned sets OnlyIfBanned field to given value.

### HasOnlyIfBanned

`func (o *UnbanChatMemberRequest) HasOnlyIfBanned() bool`

HasOnlyIfBanned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


