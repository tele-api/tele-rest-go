# ReactionTypeCustomEmoji

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the reaction, always “custom\\_emoji” | [default to "custom_emoji"]
**CustomEmojiId** | **string** | Custom emoji identifier | 

## Methods

### NewReactionTypeCustomEmoji

`func NewReactionTypeCustomEmoji(type_ string, customEmojiId string, ) *ReactionTypeCustomEmoji`

NewReactionTypeCustomEmoji instantiates a new ReactionTypeCustomEmoji object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionTypeCustomEmojiWithDefaults

`func NewReactionTypeCustomEmojiWithDefaults() *ReactionTypeCustomEmoji`

NewReactionTypeCustomEmojiWithDefaults instantiates a new ReactionTypeCustomEmoji object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReactionTypeCustomEmoji) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionTypeCustomEmoji) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionTypeCustomEmoji) SetType(v string)`

SetType sets Type field to given value.


### GetCustomEmojiId

`func (o *ReactionTypeCustomEmoji) GetCustomEmojiId() string`

GetCustomEmojiId returns the CustomEmojiId field if non-nil, zero value otherwise.

### GetCustomEmojiIdOk

`func (o *ReactionTypeCustomEmoji) GetCustomEmojiIdOk() (*string, bool)`

GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiId

`func (o *ReactionTypeCustomEmoji) SetCustomEmojiId(v string)`

SetCustomEmojiId sets CustomEmojiId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


