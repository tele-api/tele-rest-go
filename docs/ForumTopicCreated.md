# ForumTopicCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the topic | 
**IconColor** | **int32** | Color of the topic icon in RGB format | 
**IconCustomEmojiId** | Pointer to **string** | *Optional*. Unique identifier of the custom emoji shown as the topic icon | [optional] 
**IsNameImplicit** | Pointer to **bool** | *Optional*. *True*, if the name of the topic wasn&#39;t specified explicitly by its creator and likely needs to be changed by the bot | [optional] [default to true]

## Methods

### NewForumTopicCreated

`func NewForumTopicCreated(name string, iconColor int32, ) *ForumTopicCreated`

NewForumTopicCreated instantiates a new ForumTopicCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumTopicCreatedWithDefaults

`func NewForumTopicCreatedWithDefaults() *ForumTopicCreated`

NewForumTopicCreatedWithDefaults instantiates a new ForumTopicCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ForumTopicCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForumTopicCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForumTopicCreated) SetName(v string)`

SetName sets Name field to given value.


### GetIconColor

`func (o *ForumTopicCreated) GetIconColor() int32`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *ForumTopicCreated) GetIconColorOk() (*int32, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *ForumTopicCreated) SetIconColor(v int32)`

SetIconColor sets IconColor field to given value.


### GetIconCustomEmojiId

`func (o *ForumTopicCreated) GetIconCustomEmojiId() string`

GetIconCustomEmojiId returns the IconCustomEmojiId field if non-nil, zero value otherwise.

### GetIconCustomEmojiIdOk

`func (o *ForumTopicCreated) GetIconCustomEmojiIdOk() (*string, bool)`

GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCustomEmojiId

`func (o *ForumTopicCreated) SetIconCustomEmojiId(v string)`

SetIconCustomEmojiId sets IconCustomEmojiId field to given value.

### HasIconCustomEmojiId

`func (o *ForumTopicCreated) HasIconCustomEmojiId() bool`

HasIconCustomEmojiId returns a boolean if a field has been set.

### GetIsNameImplicit

`func (o *ForumTopicCreated) GetIsNameImplicit() bool`

GetIsNameImplicit returns the IsNameImplicit field if non-nil, zero value otherwise.

### GetIsNameImplicitOk

`func (o *ForumTopicCreated) GetIsNameImplicitOk() (*bool, bool)`

GetIsNameImplicitOk returns a tuple with the IsNameImplicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNameImplicit

`func (o *ForumTopicCreated) SetIsNameImplicit(v bool)`

SetIsNameImplicit sets IsNameImplicit field to given value.

### HasIsNameImplicit

`func (o *ForumTopicCreated) HasIsNameImplicit() bool`

HasIsNameImplicit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


