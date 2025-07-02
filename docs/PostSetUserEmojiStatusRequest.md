# PostSetUserEmojiStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the target user | 
**EmojiStatusCustomEmojiId** | Pointer to **string** | Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status. | [optional] 
**EmojiStatusExpirationDate** | Pointer to **int32** | Expiration date of the emoji status, if any | [optional] 

## Methods

### NewPostSetUserEmojiStatusRequest

`func NewPostSetUserEmojiStatusRequest(userId int32, ) *PostSetUserEmojiStatusRequest`

NewPostSetUserEmojiStatusRequest instantiates a new PostSetUserEmojiStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetUserEmojiStatusRequestWithDefaults

`func NewPostSetUserEmojiStatusRequestWithDefaults() *PostSetUserEmojiStatusRequest`

NewPostSetUserEmojiStatusRequestWithDefaults instantiates a new PostSetUserEmojiStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostSetUserEmojiStatusRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostSetUserEmojiStatusRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostSetUserEmojiStatusRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetEmojiStatusCustomEmojiId

`func (o *PostSetUserEmojiStatusRequest) GetEmojiStatusCustomEmojiId() string`

GetEmojiStatusCustomEmojiId returns the EmojiStatusCustomEmojiId field if non-nil, zero value otherwise.

### GetEmojiStatusCustomEmojiIdOk

`func (o *PostSetUserEmojiStatusRequest) GetEmojiStatusCustomEmojiIdOk() (*string, bool)`

GetEmojiStatusCustomEmojiIdOk returns a tuple with the EmojiStatusCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiStatusCustomEmojiId

`func (o *PostSetUserEmojiStatusRequest) SetEmojiStatusCustomEmojiId(v string)`

SetEmojiStatusCustomEmojiId sets EmojiStatusCustomEmojiId field to given value.

### HasEmojiStatusCustomEmojiId

`func (o *PostSetUserEmojiStatusRequest) HasEmojiStatusCustomEmojiId() bool`

HasEmojiStatusCustomEmojiId returns a boolean if a field has been set.

### GetEmojiStatusExpirationDate

`func (o *PostSetUserEmojiStatusRequest) GetEmojiStatusExpirationDate() int32`

GetEmojiStatusExpirationDate returns the EmojiStatusExpirationDate field if non-nil, zero value otherwise.

### GetEmojiStatusExpirationDateOk

`func (o *PostSetUserEmojiStatusRequest) GetEmojiStatusExpirationDateOk() (*int32, bool)`

GetEmojiStatusExpirationDateOk returns a tuple with the EmojiStatusExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiStatusExpirationDate

`func (o *PostSetUserEmojiStatusRequest) SetEmojiStatusExpirationDate(v int32)`

SetEmojiStatusExpirationDate sets EmojiStatusExpirationDate field to given value.

### HasEmojiStatusExpirationDate

`func (o *PostSetUserEmojiStatusRequest) HasEmojiStatusExpirationDate() bool`

HasEmojiStatusExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


