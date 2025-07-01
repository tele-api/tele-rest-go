# Venue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**Location**](Location.md) |  | 
**Title** | **string** | Name of the venue | 
**Address** | **string** | Address of the venue | 
**FoursquareId** | Pointer to **string** | *Optional*. Foursquare identifier of the venue | [optional] 
**FoursquareType** | Pointer to **string** | *Optional*. Foursquare type of the venue. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.) | [optional] 
**GooglePlaceId** | Pointer to **string** | *Optional*. Google Places identifier of the venue | [optional] 
**GooglePlaceType** | Pointer to **string** | *Optional*. Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).) | [optional] 

## Methods

### NewVenue

`func NewVenue(location Location, title string, address string, ) *Venue`

NewVenue instantiates a new Venue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVenueWithDefaults

`func NewVenueWithDefaults() *Venue`

NewVenueWithDefaults instantiates a new Venue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *Venue) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Venue) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Venue) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetTitle

`func (o *Venue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Venue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Venue) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAddress

`func (o *Venue) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Venue) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Venue) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFoursquareId

`func (o *Venue) GetFoursquareId() string`

GetFoursquareId returns the FoursquareId field if non-nil, zero value otherwise.

### GetFoursquareIdOk

`func (o *Venue) GetFoursquareIdOk() (*string, bool)`

GetFoursquareIdOk returns a tuple with the FoursquareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareId

`func (o *Venue) SetFoursquareId(v string)`

SetFoursquareId sets FoursquareId field to given value.

### HasFoursquareId

`func (o *Venue) HasFoursquareId() bool`

HasFoursquareId returns a boolean if a field has been set.

### GetFoursquareType

`func (o *Venue) GetFoursquareType() string`

GetFoursquareType returns the FoursquareType field if non-nil, zero value otherwise.

### GetFoursquareTypeOk

`func (o *Venue) GetFoursquareTypeOk() (*string, bool)`

GetFoursquareTypeOk returns a tuple with the FoursquareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareType

`func (o *Venue) SetFoursquareType(v string)`

SetFoursquareType sets FoursquareType field to given value.

### HasFoursquareType

`func (o *Venue) HasFoursquareType() bool`

HasFoursquareType returns a boolean if a field has been set.

### GetGooglePlaceId

`func (o *Venue) GetGooglePlaceId() string`

GetGooglePlaceId returns the GooglePlaceId field if non-nil, zero value otherwise.

### GetGooglePlaceIdOk

`func (o *Venue) GetGooglePlaceIdOk() (*string, bool)`

GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceId

`func (o *Venue) SetGooglePlaceId(v string)`

SetGooglePlaceId sets GooglePlaceId field to given value.

### HasGooglePlaceId

`func (o *Venue) HasGooglePlaceId() bool`

HasGooglePlaceId returns a boolean if a field has been set.

### GetGooglePlaceType

`func (o *Venue) GetGooglePlaceType() string`

GetGooglePlaceType returns the GooglePlaceType field if non-nil, zero value otherwise.

### GetGooglePlaceTypeOk

`func (o *Venue) GetGooglePlaceTypeOk() (*string, bool)`

GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceType

`func (o *Venue) SetGooglePlaceType(v string)`

SetGooglePlaceType sets GooglePlaceType field to given value.

### HasGooglePlaceType

`func (o *Venue) HasGooglePlaceType() bool`

HasGooglePlaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


