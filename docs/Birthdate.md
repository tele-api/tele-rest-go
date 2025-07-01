# Birthdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | **int32** | Day of the user&#39;s birth; 1-31 | 
**Month** | **int32** | Month of the user&#39;s birth; 1-12 | 
**Year** | Pointer to **int32** | *Optional*. Year of the user&#39;s birth | [optional] 

## Methods

### NewBirthdate

`func NewBirthdate(day int32, month int32, ) *Birthdate`

NewBirthdate instantiates a new Birthdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBirthdateWithDefaults

`func NewBirthdateWithDefaults() *Birthdate`

NewBirthdateWithDefaults instantiates a new Birthdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *Birthdate) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Birthdate) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Birthdate) SetDay(v int32)`

SetDay sets Day field to given value.


### GetMonth

`func (o *Birthdate) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Birthdate) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Birthdate) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetYear

`func (o *Birthdate) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Birthdate) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Birthdate) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *Birthdate) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


