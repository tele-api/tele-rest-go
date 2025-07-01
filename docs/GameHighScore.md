# GameHighScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **int32** | Position in high score table for the game | 
**User** | [**User**](User.md) |  | 
**Score** | **int32** | Score | 

## Methods

### NewGameHighScore

`func NewGameHighScore(position int32, user User, score int32, ) *GameHighScore`

NewGameHighScore instantiates a new GameHighScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameHighScoreWithDefaults

`func NewGameHighScoreWithDefaults() *GameHighScore`

NewGameHighScoreWithDefaults instantiates a new GameHighScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *GameHighScore) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GameHighScore) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GameHighScore) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetUser

`func (o *GameHighScore) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GameHighScore) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GameHighScore) SetUser(v User)`

SetUser sets User field to given value.


### GetScore

`func (o *GameHighScore) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GameHighScore) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GameHighScore) SetScore(v int32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


