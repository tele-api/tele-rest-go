# UnbanChatMemberPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BanChatMemberPostRequestChatId**](BanChatMemberPostRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**OnlyIfBanned** | Pointer to **bool** | Do nothing if the user is not banned | [optional] 

## Methods

### NewUnbanChatMemberPostRequest

`func NewUnbanChatMemberPostRequest(chatId BanChatMemberPostRequestChatId, userId int32, ) *UnbanChatMemberPostRequest`

NewUnbanChatMemberPostRequest instantiates a new UnbanChatMemberPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbanChatMemberPostRequestWithDefaults

`func NewUnbanChatMemberPostRequestWithDefaults() *UnbanChatMemberPostRequest`

NewUnbanChatMemberPostRequestWithDefaults instantiates a new UnbanChatMemberPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *UnbanChatMemberPostRequest) GetChatId() BanChatMemberPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *UnbanChatMemberPostRequest) GetChatIdOk() (*BanChatMemberPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *UnbanChatMemberPostRequest) SetChatId(v BanChatMemberPostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *UnbanChatMemberPostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnbanChatMemberPostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnbanChatMemberPostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetOnlyIfBanned

`func (o *UnbanChatMemberPostRequest) GetOnlyIfBanned() bool`

GetOnlyIfBanned returns the OnlyIfBanned field if non-nil, zero value otherwise.

### GetOnlyIfBannedOk

`func (o *UnbanChatMemberPostRequest) GetOnlyIfBannedOk() (*bool, bool)`

GetOnlyIfBannedOk returns a tuple with the OnlyIfBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIfBanned

`func (o *UnbanChatMemberPostRequest) SetOnlyIfBanned(v bool)`

SetOnlyIfBanned sets OnlyIfBanned field to given value.

### HasOnlyIfBanned

`func (o *UnbanChatMemberPostRequest) HasOnlyIfBanned() bool`

HasOnlyIfBanned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


