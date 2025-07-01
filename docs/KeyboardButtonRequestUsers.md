# KeyboardButtonRequestUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **int32** | Signed 32-bit identifier of the request that will be received back in the [UsersShared](https://core.telegram.org/bots/api/#usersshared) object. Must be unique within the message | 
**UserIsBot** | Pointer to **bool** | *Optional*. Pass *True* to request bots, pass *False* to request regular users. If not specified, no additional restrictions are applied. | [optional] 
**UserIsPremium** | Pointer to **bool** | *Optional*. Pass *True* to request premium users, pass *False* to request non-premium users. If not specified, no additional restrictions are applied. | [optional] 
**MaxQuantity** | Pointer to **int32** | *Optional*. The maximum number of users to be selected; 1-10. Defaults to 1. | [optional] [default to 1]
**RequestName** | Pointer to **bool** | *Optional*. Pass *True* to request the users&#39; first and last names | [optional] 
**RequestUsername** | Pointer to **bool** | *Optional*. Pass *True* to request the users&#39; usernames | [optional] 
**RequestPhoto** | Pointer to **bool** | *Optional*. Pass *True* to request the users&#39; photos | [optional] 

## Methods

### NewKeyboardButtonRequestUsers

`func NewKeyboardButtonRequestUsers(requestId int32, ) *KeyboardButtonRequestUsers`

NewKeyboardButtonRequestUsers instantiates a new KeyboardButtonRequestUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyboardButtonRequestUsersWithDefaults

`func NewKeyboardButtonRequestUsersWithDefaults() *KeyboardButtonRequestUsers`

NewKeyboardButtonRequestUsersWithDefaults instantiates a new KeyboardButtonRequestUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *KeyboardButtonRequestUsers) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *KeyboardButtonRequestUsers) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *KeyboardButtonRequestUsers) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.


### GetUserIsBot

`func (o *KeyboardButtonRequestUsers) GetUserIsBot() bool`

GetUserIsBot returns the UserIsBot field if non-nil, zero value otherwise.

### GetUserIsBotOk

`func (o *KeyboardButtonRequestUsers) GetUserIsBotOk() (*bool, bool)`

GetUserIsBotOk returns a tuple with the UserIsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIsBot

`func (o *KeyboardButtonRequestUsers) SetUserIsBot(v bool)`

SetUserIsBot sets UserIsBot field to given value.

### HasUserIsBot

`func (o *KeyboardButtonRequestUsers) HasUserIsBot() bool`

HasUserIsBot returns a boolean if a field has been set.

### GetUserIsPremium

`func (o *KeyboardButtonRequestUsers) GetUserIsPremium() bool`

GetUserIsPremium returns the UserIsPremium field if non-nil, zero value otherwise.

### GetUserIsPremiumOk

`func (o *KeyboardButtonRequestUsers) GetUserIsPremiumOk() (*bool, bool)`

GetUserIsPremiumOk returns a tuple with the UserIsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIsPremium

`func (o *KeyboardButtonRequestUsers) SetUserIsPremium(v bool)`

SetUserIsPremium sets UserIsPremium field to given value.

### HasUserIsPremium

`func (o *KeyboardButtonRequestUsers) HasUserIsPremium() bool`

HasUserIsPremium returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *KeyboardButtonRequestUsers) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *KeyboardButtonRequestUsers) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *KeyboardButtonRequestUsers) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *KeyboardButtonRequestUsers) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetRequestName

`func (o *KeyboardButtonRequestUsers) GetRequestName() bool`

GetRequestName returns the RequestName field if non-nil, zero value otherwise.

### GetRequestNameOk

`func (o *KeyboardButtonRequestUsers) GetRequestNameOk() (*bool, bool)`

GetRequestNameOk returns a tuple with the RequestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestName

`func (o *KeyboardButtonRequestUsers) SetRequestName(v bool)`

SetRequestName sets RequestName field to given value.

### HasRequestName

`func (o *KeyboardButtonRequestUsers) HasRequestName() bool`

HasRequestName returns a boolean if a field has been set.

### GetRequestUsername

`func (o *KeyboardButtonRequestUsers) GetRequestUsername() bool`

GetRequestUsername returns the RequestUsername field if non-nil, zero value otherwise.

### GetRequestUsernameOk

`func (o *KeyboardButtonRequestUsers) GetRequestUsernameOk() (*bool, bool)`

GetRequestUsernameOk returns a tuple with the RequestUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUsername

`func (o *KeyboardButtonRequestUsers) SetRequestUsername(v bool)`

SetRequestUsername sets RequestUsername field to given value.

### HasRequestUsername

`func (o *KeyboardButtonRequestUsers) HasRequestUsername() bool`

HasRequestUsername returns a boolean if a field has been set.

### GetRequestPhoto

`func (o *KeyboardButtonRequestUsers) GetRequestPhoto() bool`

GetRequestPhoto returns the RequestPhoto field if non-nil, zero value otherwise.

### GetRequestPhotoOk

`func (o *KeyboardButtonRequestUsers) GetRequestPhotoOk() (*bool, bool)`

GetRequestPhotoOk returns a tuple with the RequestPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPhoto

`func (o *KeyboardButtonRequestUsers) SetRequestPhoto(v bool)`

SetRequestPhoto sets RequestPhoto field to given value.

### HasRequestPhoto

`func (o *KeyboardButtonRequestUsers) HasRequestPhoto() bool`

HasRequestPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


