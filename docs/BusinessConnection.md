# BusinessConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the business connection | 
**User** | [**User**](User.md) |  | 
**UserChatId** | **int32** | Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. | 
**Date** | **int32** | Date the connection was established in Unix time | 
**Rights** | Pointer to [**BusinessBotRights**](BusinessBotRights.md) |  | [optional] 
**IsEnabled** | **bool** | True, if the connection is active | 

## Methods

### NewBusinessConnection

`func NewBusinessConnection(id string, user User, userChatId int32, date int32, isEnabled bool, ) *BusinessConnection`

NewBusinessConnection instantiates a new BusinessConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessConnectionWithDefaults

`func NewBusinessConnectionWithDefaults() *BusinessConnection`

NewBusinessConnectionWithDefaults instantiates a new BusinessConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessConnection) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *BusinessConnection) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BusinessConnection) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BusinessConnection) SetUser(v User)`

SetUser sets User field to given value.


### GetUserChatId

`func (o *BusinessConnection) GetUserChatId() int32`

GetUserChatId returns the UserChatId field if non-nil, zero value otherwise.

### GetUserChatIdOk

`func (o *BusinessConnection) GetUserChatIdOk() (*int32, bool)`

GetUserChatIdOk returns a tuple with the UserChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserChatId

`func (o *BusinessConnection) SetUserChatId(v int32)`

SetUserChatId sets UserChatId field to given value.


### GetDate

`func (o *BusinessConnection) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BusinessConnection) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BusinessConnection) SetDate(v int32)`

SetDate sets Date field to given value.


### GetRights

`func (o *BusinessConnection) GetRights() BusinessBotRights`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *BusinessConnection) GetRightsOk() (*BusinessBotRights, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *BusinessConnection) SetRights(v BusinessBotRights)`

SetRights sets Rights field to given value.

### HasRights

`func (o *BusinessConnection) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetIsEnabled

`func (o *BusinessConnection) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BusinessConnection) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BusinessConnection) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


