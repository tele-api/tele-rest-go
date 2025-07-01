# SharedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means. | 
**FirstName** | Pointer to **string** | *Optional*. First name of the user, if the name was requested by the bot | [optional] 
**LastName** | Pointer to **string** | *Optional*. Last name of the user, if the name was requested by the bot | [optional] 
**Username** | Pointer to **string** | *Optional*. Username of the user, if the username was requested by the bot | [optional] 
**Photo** | Pointer to [**[]PhotoSize**](PhotoSize.md) | *Optional*. Available sizes of the chat photo, if the photo was requested by the bot | [optional] 

## Methods

### NewSharedUser

`func NewSharedUser(userId int32, ) *SharedUser`

NewSharedUser instantiates a new SharedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedUserWithDefaults

`func NewSharedUserWithDefaults() *SharedUser`

NewSharedUserWithDefaults instantiates a new SharedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SharedUser) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SharedUser) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SharedUser) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *SharedUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SharedUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SharedUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SharedUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SharedUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SharedUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SharedUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SharedUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *SharedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SharedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SharedUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SharedUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhoto

`func (o *SharedUser) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *SharedUser) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *SharedUser) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *SharedUser) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


