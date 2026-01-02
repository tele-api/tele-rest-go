# UserRating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **int32** | Current level of the user, indicating their reliability when purchasing digital goods and services. A higher level suggests a more trustworthy customer; a negative level is likely reason for concern. | 
**Rating** | **int32** | Numerical value of the user&#39;s rating; the higher the rating, the better | 
**CurrentLevelRating** | **int32** | The rating value required to get the current level | 
**NextLevelRating** | Pointer to **int32** | *Optional*. The rating value required to get to the next level; omitted if the maximum level was reached | [optional] 

## Methods

### NewUserRating

`func NewUserRating(level int32, rating int32, currentLevelRating int32, ) *UserRating`

NewUserRating instantiates a new UserRating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRatingWithDefaults

`func NewUserRatingWithDefaults() *UserRating`

NewUserRatingWithDefaults instantiates a new UserRating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *UserRating) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UserRating) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UserRating) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetRating

`func (o *UserRating) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *UserRating) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *UserRating) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetCurrentLevelRating

`func (o *UserRating) GetCurrentLevelRating() int32`

GetCurrentLevelRating returns the CurrentLevelRating field if non-nil, zero value otherwise.

### GetCurrentLevelRatingOk

`func (o *UserRating) GetCurrentLevelRatingOk() (*int32, bool)`

GetCurrentLevelRatingOk returns a tuple with the CurrentLevelRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLevelRating

`func (o *UserRating) SetCurrentLevelRating(v int32)`

SetCurrentLevelRating sets CurrentLevelRating field to given value.


### GetNextLevelRating

`func (o *UserRating) GetNextLevelRating() int32`

GetNextLevelRating returns the NextLevelRating field if non-nil, zero value otherwise.

### GetNextLevelRatingOk

`func (o *UserRating) GetNextLevelRatingOk() (*int32, bool)`

GetNextLevelRatingOk returns a tuple with the NextLevelRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLevelRating

`func (o *UserRating) SetNextLevelRating(v int32)`

SetNextLevelRating sets NextLevelRating field to given value.

### HasNextLevelRating

`func (o *UserRating) HasNextLevelRating() bool`

HasNextLevelRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


