# CloseForumTopicPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**RestrictChatMemberPostRequestChatId**](RestrictChatMemberPostRequestChatId.md) |  | 
**MessageThreadId** | **int32** | Unique identifier for the target message thread of the forum topic | 

## Methods

### NewCloseForumTopicPostRequest

`func NewCloseForumTopicPostRequest(chatId RestrictChatMemberPostRequestChatId, messageThreadId int32, ) *CloseForumTopicPostRequest`

NewCloseForumTopicPostRequest instantiates a new CloseForumTopicPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseForumTopicPostRequestWithDefaults

`func NewCloseForumTopicPostRequestWithDefaults() *CloseForumTopicPostRequest`

NewCloseForumTopicPostRequestWithDefaults instantiates a new CloseForumTopicPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *CloseForumTopicPostRequest) GetChatId() RestrictChatMemberPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CloseForumTopicPostRequest) GetChatIdOk() (*RestrictChatMemberPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CloseForumTopicPostRequest) SetChatId(v RestrictChatMemberPostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *CloseForumTopicPostRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *CloseForumTopicPostRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *CloseForumTopicPostRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


