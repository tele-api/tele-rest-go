# Game

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the game | 
**Description** | **string** | Description of the game | 
**Photo** | [**[]PhotoSize**](PhotoSize.md) | Photo that will be displayed in the game message in chats. | 
**Text** | Pointer to **string** | *Optional*. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls [setGameScore](https://core.telegram.org/bots/api/#setgamescore), or manually edited using [editMessageText](https://core.telegram.org/bots/api/#editmessagetext). 0-4096 characters. | [optional] 
**TextEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in *text*, such as usernames, URLs, bot commands, etc. | [optional] 
**Animation** | Pointer to [**Animation**](Animation.md) |  | [optional] 

## Methods

### NewGame

`func NewGame(title string, description string, photo []PhotoSize, ) *Game`

NewGame instantiates a new Game object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameWithDefaults

`func NewGameWithDefaults() *Game`

NewGameWithDefaults instantiates a new Game object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Game) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Game) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Game) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Game) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Game) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Game) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPhoto

`func (o *Game) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Game) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Game) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.


### GetText

`func (o *Game) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Game) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Game) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Game) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextEntities

`func (o *Game) GetTextEntities() []MessageEntity`

GetTextEntities returns the TextEntities field if non-nil, zero value otherwise.

### GetTextEntitiesOk

`func (o *Game) GetTextEntitiesOk() (*[]MessageEntity, bool)`

GetTextEntitiesOk returns a tuple with the TextEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEntities

`func (o *Game) SetTextEntities(v []MessageEntity)`

SetTextEntities sets TextEntities field to given value.

### HasTextEntities

`func (o *Game) HasTextEntities() bool`

HasTextEntities returns a boolean if a field has been set.

### GetAnimation

`func (o *Game) GetAnimation() Animation`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *Game) GetAnimationOk() (*Animation, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *Game) SetAnimation(v Animation)`

SetAnimation sets Animation field to given value.

### HasAnimation

`func (o *Game) HasAnimation() bool`

HasAnimation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


