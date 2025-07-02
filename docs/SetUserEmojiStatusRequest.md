# SetUserEmojiStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the target user | 
**EmojiStatusCustomEmojiId** | Pointer to **string** | Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status. | [optional] 
**EmojiStatusExpirationDate** | Pointer to **int32** | Expiration date of the emoji status, if any | [optional] 

## Methods

### NewSetUserEmojiStatusRequest

`func NewSetUserEmojiStatusRequest(userId int32, ) *SetUserEmojiStatusRequest`

NewSetUserEmojiStatusRequest instantiates a new SetUserEmojiStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetUserEmojiStatusRequestWithDefaults

`func NewSetUserEmojiStatusRequestWithDefaults() *SetUserEmojiStatusRequest`

NewSetUserEmojiStatusRequestWithDefaults instantiates a new SetUserEmojiStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SetUserEmojiStatusRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SetUserEmojiStatusRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SetUserEmojiStatusRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetEmojiStatusCustomEmojiId

`func (o *SetUserEmojiStatusRequest) GetEmojiStatusCustomEmojiId() string`

GetEmojiStatusCustomEmojiId returns the EmojiStatusCustomEmojiId field if non-nil, zero value otherwise.

### GetEmojiStatusCustomEmojiIdOk

`func (o *SetUserEmojiStatusRequest) GetEmojiStatusCustomEmojiIdOk() (*string, bool)`

GetEmojiStatusCustomEmojiIdOk returns a tuple with the EmojiStatusCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiStatusCustomEmojiId

`func (o *SetUserEmojiStatusRequest) SetEmojiStatusCustomEmojiId(v string)`

SetEmojiStatusCustomEmojiId sets EmojiStatusCustomEmojiId field to given value.

### HasEmojiStatusCustomEmojiId

`func (o *SetUserEmojiStatusRequest) HasEmojiStatusCustomEmojiId() bool`

HasEmojiStatusCustomEmojiId returns a boolean if a field has been set.

### GetEmojiStatusExpirationDate

`func (o *SetUserEmojiStatusRequest) GetEmojiStatusExpirationDate() int32`

GetEmojiStatusExpirationDate returns the EmojiStatusExpirationDate field if non-nil, zero value otherwise.

### GetEmojiStatusExpirationDateOk

`func (o *SetUserEmojiStatusRequest) GetEmojiStatusExpirationDateOk() (*int32, bool)`

GetEmojiStatusExpirationDateOk returns a tuple with the EmojiStatusExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiStatusExpirationDate

`func (o *SetUserEmojiStatusRequest) SetEmojiStatusExpirationDate(v int32)`

SetEmojiStatusExpirationDate sets EmojiStatusExpirationDate field to given value.

### HasEmojiStatusExpirationDate

`func (o *SetUserEmojiStatusRequest) HasEmojiStatusExpirationDate() bool`

HasEmojiStatusExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


