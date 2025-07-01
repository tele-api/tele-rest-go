# BusinessOpeningHoursInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpeningMinute** | **int32** | The minute&#39;s sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 \\* 24 \\* 60 | 
**ClosingMinute** | **int32** | The minute&#39;s sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 \\* 24 \\* 60 | 

## Methods

### NewBusinessOpeningHoursInterval

`func NewBusinessOpeningHoursInterval(openingMinute int32, closingMinute int32, ) *BusinessOpeningHoursInterval`

NewBusinessOpeningHoursInterval instantiates a new BusinessOpeningHoursInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOpeningHoursIntervalWithDefaults

`func NewBusinessOpeningHoursIntervalWithDefaults() *BusinessOpeningHoursInterval`

NewBusinessOpeningHoursIntervalWithDefaults instantiates a new BusinessOpeningHoursInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpeningMinute

`func (o *BusinessOpeningHoursInterval) GetOpeningMinute() int32`

GetOpeningMinute returns the OpeningMinute field if non-nil, zero value otherwise.

### GetOpeningMinuteOk

`func (o *BusinessOpeningHoursInterval) GetOpeningMinuteOk() (*int32, bool)`

GetOpeningMinuteOk returns a tuple with the OpeningMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningMinute

`func (o *BusinessOpeningHoursInterval) SetOpeningMinute(v int32)`

SetOpeningMinute sets OpeningMinute field to given value.


### GetClosingMinute

`func (o *BusinessOpeningHoursInterval) GetClosingMinute() int32`

GetClosingMinute returns the ClosingMinute field if non-nil, zero value otherwise.

### GetClosingMinuteOk

`func (o *BusinessOpeningHoursInterval) GetClosingMinuteOk() (*int32, bool)`

GetClosingMinuteOk returns a tuple with the ClosingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingMinute

`func (o *BusinessOpeningHoursInterval) SetClosingMinute(v int32)`

SetClosingMinute sets ClosingMinute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


