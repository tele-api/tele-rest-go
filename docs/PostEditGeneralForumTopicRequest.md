# PostEditGeneralForumTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**PostRestrictChatMemberRequestChatId**](PostRestrictChatMemberRequestChatId.md) |  | 
**Name** | **string** | New topic name, 1-128 characters | 

## Methods

### NewPostEditGeneralForumTopicRequest

`func NewPostEditGeneralForumTopicRequest(chatId PostRestrictChatMemberRequestChatId, name string, ) *PostEditGeneralForumTopicRequest`

NewPostEditGeneralForumTopicRequest instantiates a new PostEditGeneralForumTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEditGeneralForumTopicRequestWithDefaults

`func NewPostEditGeneralForumTopicRequestWithDefaults() *PostEditGeneralForumTopicRequest`

NewPostEditGeneralForumTopicRequestWithDefaults instantiates a new PostEditGeneralForumTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PostEditGeneralForumTopicRequest) GetChatId() PostRestrictChatMemberRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostEditGeneralForumTopicRequest) GetChatIdOk() (*PostRestrictChatMemberRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostEditGeneralForumTopicRequest) SetChatId(v PostRestrictChatMemberRequestChatId)`

SetChatId sets ChatId field to given value.


### GetName

`func (o *PostEditGeneralForumTopicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostEditGeneralForumTopicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostEditGeneralForumTopicRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


