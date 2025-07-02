# PostUnbanChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostBanChatMemberRequestChatId**](PostBanChatMemberRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 
**OnlyIfBanned** | Pointer to **bool** | Do nothing if the user is not banned | [optional] 

## Methods

### NewPostUnbanChatMemberRequest

`func NewPostUnbanChatMemberRequest(chatId PostBanChatMemberRequestChatId, userId int32, ) *PostUnbanChatMemberRequest`

NewPostUnbanChatMemberRequest instantiates a new PostUnbanChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUnbanChatMemberRequestWithDefaults

`func NewPostUnbanChatMemberRequestWithDefaults() *PostUnbanChatMemberRequest`

NewPostUnbanChatMemberRequestWithDefaults instantiates a new PostUnbanChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostUnbanChatMemberRequest) GetChatId() PostBanChatMemberRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostUnbanChatMemberRequest) GetChatIdOk() (*PostBanChatMemberRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostUnbanChatMemberRequest) SetChatId(v PostBanChatMemberRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *PostUnbanChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostUnbanChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostUnbanChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetOnlyIfBanned

`func (o *PostUnbanChatMemberRequest) GetOnlyIfBanned() bool`

GetOnlyIfBanned returns the OnlyIfBanned field if non-nil, zero value otherwise.

### GetOnlyIfBannedOk

`func (o *PostUnbanChatMemberRequest) GetOnlyIfBannedOk() (*bool, bool)`

GetOnlyIfBannedOk returns a tuple with the OnlyIfBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIfBanned

`func (o *PostUnbanChatMemberRequest) SetOnlyIfBanned(v bool)`

SetOnlyIfBanned sets OnlyIfBanned field to given value.

### HasOnlyIfBanned

`func (o *PostUnbanChatMemberRequest) HasOnlyIfBanned() bool`

HasOnlyIfBanned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


