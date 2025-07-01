# InputVenueMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float32** | Latitude of the venue in degrees | 
**Longitude** | **float32** | Longitude of the venue in degrees | 
**Title** | **string** | Name of the venue | 
**Address** | **string** | Address of the venue | 
**FoursquareId** | Pointer to **string** | *Optional*. Foursquare identifier of the venue, if known | [optional] 
**FoursquareType** | Pointer to **string** | *Optional*. Foursquare type of the venue, if known. (For example, “arts\\_entertainment/default”, “arts\\_entertainment/aquarium” or “food/icecream”.) | [optional] 
**GooglePlaceId** | Pointer to **string** | *Optional*. Google Places identifier of the venue | [optional] 
**GooglePlaceType** | Pointer to **string** | *Optional*. Google Places type of the venue. (See [supported types](https://developers.google.com/places/web-service/supported_types).) | [optional] 

## Methods

### NewInputVenueMessageContent

`func NewInputVenueMessageContent(latitude float32, longitude float32, title string, address string, ) *InputVenueMessageContent`

NewInputVenueMessageContent instantiates a new InputVenueMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputVenueMessageContentWithDefaults

`func NewInputVenueMessageContentWithDefaults() *InputVenueMessageContent`

NewInputVenueMessageContentWithDefaults instantiates a new InputVenueMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *InputVenueMessageContent) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InputVenueMessageContent) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InputVenueMessageContent) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *InputVenueMessageContent) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InputVenueMessageContent) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InputVenueMessageContent) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetTitle

`func (o *InputVenueMessageContent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InputVenueMessageContent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InputVenueMessageContent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAddress

`func (o *InputVenueMessageContent) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InputVenueMessageContent) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InputVenueMessageContent) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFoursquareId

`func (o *InputVenueMessageContent) GetFoursquareId() string`

GetFoursquareId returns the FoursquareId field if non-nil, zero value otherwise.

### GetFoursquareIdOk

`func (o *InputVenueMessageContent) GetFoursquareIdOk() (*string, bool)`

GetFoursquareIdOk returns a tuple with the FoursquareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareId

`func (o *InputVenueMessageContent) SetFoursquareId(v string)`

SetFoursquareId sets FoursquareId field to given value.

### HasFoursquareId

`func (o *InputVenueMessageContent) HasFoursquareId() bool`

HasFoursquareId returns a boolean if a field has been set.

### GetFoursquareType

`func (o *InputVenueMessageContent) GetFoursquareType() string`

GetFoursquareType returns the FoursquareType field if non-nil, zero value otherwise.

### GetFoursquareTypeOk

`func (o *InputVenueMessageContent) GetFoursquareTypeOk() (*string, bool)`

GetFoursquareTypeOk returns a tuple with the FoursquareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoursquareType

`func (o *InputVenueMessageContent) SetFoursquareType(v string)`

SetFoursquareType sets FoursquareType field to given value.

### HasFoursquareType

`func (o *InputVenueMessageContent) HasFoursquareType() bool`

HasFoursquareType returns a boolean if a field has been set.

### GetGooglePlaceId

`func (o *InputVenueMessageContent) GetGooglePlaceId() string`

GetGooglePlaceId returns the GooglePlaceId field if non-nil, zero value otherwise.

### GetGooglePlaceIdOk

`func (o *InputVenueMessageContent) GetGooglePlaceIdOk() (*string, bool)`

GetGooglePlaceIdOk returns a tuple with the GooglePlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceId

`func (o *InputVenueMessageContent) SetGooglePlaceId(v string)`

SetGooglePlaceId sets GooglePlaceId field to given value.

### HasGooglePlaceId

`func (o *InputVenueMessageContent) HasGooglePlaceId() bool`

HasGooglePlaceId returns a boolean if a field has been set.

### GetGooglePlaceType

`func (o *InputVenueMessageContent) GetGooglePlaceType() string`

GetGooglePlaceType returns the GooglePlaceType field if non-nil, zero value otherwise.

### GetGooglePlaceTypeOk

`func (o *InputVenueMessageContent) GetGooglePlaceTypeOk() (*string, bool)`

GetGooglePlaceTypeOk returns a tuple with the GooglePlaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlaceType

`func (o *InputVenueMessageContent) SetGooglePlaceType(v string)`

SetGooglePlaceType sets GooglePlaceType field to given value.

### HasGooglePlaceType

`func (o *InputVenueMessageContent) HasGooglePlaceType() bool`

HasGooglePlaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


