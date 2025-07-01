# ReactionTypeEmoji

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the reaction, always “emoji” | [default to "emoji"]
**Emoji** | **string** | Reaction emoji. Currently, it can be one of \&quot;❤\&quot;, \&quot;👍\&quot;, \&quot;👎\&quot;, \&quot;🔥\&quot;, \&quot;🥰\&quot;, \&quot;👏\&quot;, \&quot;😁\&quot;, \&quot;🤔\&quot;, \&quot;🤯\&quot;, \&quot;😱\&quot;, \&quot;🤬\&quot;, \&quot;😢\&quot;, \&quot;🎉\&quot;, \&quot;🤩\&quot;, \&quot;🤮\&quot;, \&quot;💩\&quot;, \&quot;🙏\&quot;, \&quot;👌\&quot;, \&quot;🕊\&quot;, \&quot;🤡\&quot;, \&quot;🥱\&quot;, \&quot;🥴\&quot;, \&quot;😍\&quot;, \&quot;🐳\&quot;, \&quot;❤‍🔥\&quot;, \&quot;🌚\&quot;, \&quot;🌭\&quot;, \&quot;💯\&quot;, \&quot;🤣\&quot;, \&quot;⚡\&quot;, \&quot;🍌\&quot;, \&quot;🏆\&quot;, \&quot;💔\&quot;, \&quot;🤨\&quot;, \&quot;😐\&quot;, \&quot;🍓\&quot;, \&quot;🍾\&quot;, \&quot;💋\&quot;, \&quot;🖕\&quot;, \&quot;😈\&quot;, \&quot;😴\&quot;, \&quot;😭\&quot;, \&quot;🤓\&quot;, \&quot;👻\&quot;, \&quot;👨‍💻\&quot;, \&quot;👀\&quot;, \&quot;🎃\&quot;, \&quot;🙈\&quot;, \&quot;😇\&quot;, \&quot;😨\&quot;, \&quot;🤝\&quot;, \&quot;✍\&quot;, \&quot;🤗\&quot;, \&quot;🫡\&quot;, \&quot;🎅\&quot;, \&quot;🎄\&quot;, \&quot;☃\&quot;, \&quot;💅\&quot;, \&quot;🤪\&quot;, \&quot;🗿\&quot;, \&quot;🆒\&quot;, \&quot;💘\&quot;, \&quot;🙉\&quot;, \&quot;🦄\&quot;, \&quot;😘\&quot;, \&quot;💊\&quot;, \&quot;🙊\&quot;, \&quot;😎\&quot;, \&quot;👾\&quot;, \&quot;🤷‍♂\&quot;, \&quot;🤷\&quot;, \&quot;🤷‍♀\&quot;, \&quot;😡\&quot; | 

## Methods

### NewReactionTypeEmoji

`func NewReactionTypeEmoji(type_ string, emoji string, ) *ReactionTypeEmoji`

NewReactionTypeEmoji instantiates a new ReactionTypeEmoji object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionTypeEmojiWithDefaults

`func NewReactionTypeEmojiWithDefaults() *ReactionTypeEmoji`

NewReactionTypeEmojiWithDefaults instantiates a new ReactionTypeEmoji object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReactionTypeEmoji) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionTypeEmoji) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionTypeEmoji) SetType(v string)`

SetType sets Type field to given value.


### GetEmoji

`func (o *ReactionTypeEmoji) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ReactionTypeEmoji) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ReactionTypeEmoji) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


