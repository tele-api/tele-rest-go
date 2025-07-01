# ForumTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageThreadId** | **int32** | Unique identifier of the forum topic | 
**Name** | **string** | Name of the topic | 
**IconColor** | **int32** | Color of the topic icon in RGB format | 
**IconCustomEmojiId** | Pointer to **string** | *Optional*. Unique identifier of the custom emoji shown as the topic icon | [optional] 

## Methods

### NewForumTopic

`func NewForumTopic(messageThreadId int32, name string, iconColor int32, ) *ForumTopic`

NewForumTopic instantiates a new ForumTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumTopicWithDefaults

`func NewForumTopicWithDefaults() *ForumTopic`

NewForumTopicWithDefaults instantiates a new ForumTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageThreadId

`func (o *ForumTopic) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *ForumTopic) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *ForumTopic) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.


### GetName

`func (o *ForumTopic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForumTopic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForumTopic) SetName(v string)`

SetName sets Name field to given value.


### GetIconColor

`func (o *ForumTopic) GetIconColor() int32`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *ForumTopic) GetIconColorOk() (*int32, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *ForumTopic) SetIconColor(v int32)`

SetIconColor sets IconColor field to given value.


### GetIconCustomEmojiId

`func (o *ForumTopic) GetIconCustomEmojiId() string`

GetIconCustomEmojiId returns the IconCustomEmojiId field if non-nil, zero value otherwise.

### GetIconCustomEmojiIdOk

`func (o *ForumTopic) GetIconCustomEmojiIdOk() (*string, bool)`

GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCustomEmojiId

`func (o *ForumTopic) SetIconCustomEmojiId(v string)`

SetIconCustomEmojiId sets IconCustomEmojiId field to given value.

### HasIconCustomEmojiId

`func (o *ForumTopic) HasIconCustomEmojiId() bool`

HasIconCustomEmojiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


