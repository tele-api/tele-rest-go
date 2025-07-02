# SetChatPhotoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**Photo** | **interface{}** |  | 

## Methods

### NewSetChatPhotoRequest

`func NewSetChatPhotoRequest(chatId SendMessageRequestChatId, photo interface{}, ) *SetChatPhotoRequest`

NewSetChatPhotoRequest instantiates a new SetChatPhotoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatPhotoRequestWithDefaults

`func NewSetChatPhotoRequestWithDefaults() *SetChatPhotoRequest`

NewSetChatPhotoRequestWithDefaults instantiates a new SetChatPhotoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatPhotoRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatPhotoRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatPhotoRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetPhoto

`func (o *SetChatPhotoRequest) GetPhoto() interface{}`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *SetChatPhotoRequest) GetPhotoOk() (*interface{}, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *SetChatPhotoRequest) SetPhoto(v interface{})`

SetPhoto sets Photo field to given value.


### SetPhotoNil

`func (o *SetChatPhotoRequest) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *SetChatPhotoRequest) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


