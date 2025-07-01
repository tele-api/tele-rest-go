# ProximityAlertTriggered

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traveler** | [**User**](User.md) |  | 
**Watcher** | [**User**](User.md) |  | 
**Distance** | **int32** | The distance between the users | 

## Methods

### NewProximityAlertTriggered

`func NewProximityAlertTriggered(traveler User, watcher User, distance int32, ) *ProximityAlertTriggered`

NewProximityAlertTriggered instantiates a new ProximityAlertTriggered object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProximityAlertTriggeredWithDefaults

`func NewProximityAlertTriggeredWithDefaults() *ProximityAlertTriggered`

NewProximityAlertTriggeredWithDefaults instantiates a new ProximityAlertTriggered object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraveler

`func (o *ProximityAlertTriggered) GetTraveler() User`

GetTraveler returns the Traveler field if non-nil, zero value otherwise.

### GetTravelerOk

`func (o *ProximityAlertTriggered) GetTravelerOk() (*User, bool)`

GetTravelerOk returns a tuple with the Traveler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraveler

`func (o *ProximityAlertTriggered) SetTraveler(v User)`

SetTraveler sets Traveler field to given value.


### GetWatcher

`func (o *ProximityAlertTriggered) GetWatcher() User`

GetWatcher returns the Watcher field if non-nil, zero value otherwise.

### GetWatcherOk

`func (o *ProximityAlertTriggered) GetWatcherOk() (*User, bool)`

GetWatcherOk returns a tuple with the Watcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcher

`func (o *ProximityAlertTriggered) SetWatcher(v User)`

SetWatcher sets Watcher field to given value.


### GetDistance

`func (o *ProximityAlertTriggered) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *ProximityAlertTriggered) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *ProximityAlertTriggered) SetDistance(v int32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


