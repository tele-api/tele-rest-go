# ReactionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the reaction, always “paid” | [default to "paid"]
**Emoji** | **string** | Reaction emoji. Currently, it can be one of \&quot;❤\&quot;, \&quot;👍\&quot;, \&quot;👎\&quot;, \&quot;🔥\&quot;, \&quot;🥰\&quot;, \&quot;👏\&quot;, \&quot;😁\&quot;, \&quot;🤔\&quot;, \&quot;🤯\&quot;, \&quot;😱\&quot;, \&quot;🤬\&quot;, \&quot;😢\&quot;, \&quot;🎉\&quot;, \&quot;🤩\&quot;, \&quot;🤮\&quot;, \&quot;💩\&quot;, \&quot;🙏\&quot;, \&quot;👌\&quot;, \&quot;🕊\&quot;, \&quot;🤡\&quot;, \&quot;🥱\&quot;, \&quot;🥴\&quot;, \&quot;😍\&quot;, \&quot;🐳\&quot;, \&quot;❤‍🔥\&quot;, \&quot;🌚\&quot;, \&quot;🌭\&quot;, \&quot;💯\&quot;, \&quot;🤣\&quot;, \&quot;⚡\&quot;, \&quot;🍌\&quot;, \&quot;🏆\&quot;, \&quot;💔\&quot;, \&quot;🤨\&quot;, \&quot;😐\&quot;, \&quot;🍓\&quot;, \&quot;🍾\&quot;, \&quot;💋\&quot;, \&quot;🖕\&quot;, \&quot;😈\&quot;, \&quot;😴\&quot;, \&quot;😭\&quot;, \&quot;🤓\&quot;, \&quot;👻\&quot;, \&quot;👨‍💻\&quot;, \&quot;👀\&quot;, \&quot;🎃\&quot;, \&quot;🙈\&quot;, \&quot;😇\&quot;, \&quot;😨\&quot;, \&quot;🤝\&quot;, \&quot;✍\&quot;, \&quot;🤗\&quot;, \&quot;🫡\&quot;, \&quot;🎅\&quot;, \&quot;🎄\&quot;, \&quot;☃\&quot;, \&quot;💅\&quot;, \&quot;🤪\&quot;, \&quot;🗿\&quot;, \&quot;🆒\&quot;, \&quot;💘\&quot;, \&quot;🙉\&quot;, \&quot;🦄\&quot;, \&quot;😘\&quot;, \&quot;💊\&quot;, \&quot;🙊\&quot;, \&quot;😎\&quot;, \&quot;👾\&quot;, \&quot;🤷‍♂\&quot;, \&quot;🤷\&quot;, \&quot;🤷‍♀\&quot;, \&quot;😡\&quot; | 
**CustomEmojiId** | **string** | Custom emoji identifier | 

## Methods

### NewReactionType

`func NewReactionType(type_ string, emoji string, customEmojiId string, ) *ReactionType`

NewReactionType instantiates a new ReactionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionTypeWithDefaults

`func NewReactionTypeWithDefaults() *ReactionType`

NewReactionTypeWithDefaults instantiates a new ReactionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReactionType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionType) SetType(v string)`

SetType sets Type field to given value.


### GetEmoji

`func (o *ReactionType) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ReactionType) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ReactionType) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetCustomEmojiId

`func (o *ReactionType) GetCustomEmojiId() string`

GetCustomEmojiId returns the CustomEmojiId field if non-nil, zero value otherwise.

### GetCustomEmojiIdOk

`func (o *ReactionType) GetCustomEmojiIdOk() (*string, bool)`

GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiId

`func (o *ReactionType) SetCustomEmojiId(v string)`

SetCustomEmojiId sets CustomEmojiId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


