# BusinessOpeningHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZoneName** | **string** | Unique name of the time zone for which the opening hours are defined | 
**OpeningHours** | [**[]BusinessOpeningHoursInterval**](BusinessOpeningHoursInterval.md) | List of time intervals describing business opening hours | 

## Methods

### NewBusinessOpeningHours

`func NewBusinessOpeningHours(timeZoneName string, openingHours []BusinessOpeningHoursInterval, ) *BusinessOpeningHours`

NewBusinessOpeningHours instantiates a new BusinessOpeningHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOpeningHoursWithDefaults

`func NewBusinessOpeningHoursWithDefaults() *BusinessOpeningHours`

NewBusinessOpeningHoursWithDefaults instantiates a new BusinessOpeningHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZoneName

`func (o *BusinessOpeningHours) GetTimeZoneName() string`

GetTimeZoneName returns the TimeZoneName field if non-nil, zero value otherwise.

### GetTimeZoneNameOk

`func (o *BusinessOpeningHours) GetTimeZoneNameOk() (*string, bool)`

GetTimeZoneNameOk returns a tuple with the TimeZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneName

`func (o *BusinessOpeningHours) SetTimeZoneName(v string)`

SetTimeZoneName sets TimeZoneName field to given value.


### GetOpeningHours

`func (o *BusinessOpeningHours) GetOpeningHours() []BusinessOpeningHoursInterval`

GetOpeningHours returns the OpeningHours field if non-nil, zero value otherwise.

### GetOpeningHoursOk

`func (o *BusinessOpeningHours) GetOpeningHoursOk() (*[]BusinessOpeningHoursInterval, bool)`

GetOpeningHoursOk returns a tuple with the OpeningHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningHours

`func (o *BusinessOpeningHours) SetOpeningHours(v []BusinessOpeningHoursInterval)`

SetOpeningHours sets OpeningHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


