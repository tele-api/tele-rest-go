# PostSetBusinessAccountGiftSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**ShowGiftButton** | **bool** | Pass True, if a button for sending a gift to the user or by the business account must always be shown in the input field | 
**AcceptedGiftTypes** | [**AcceptedGiftTypes**](AcceptedGiftTypes.md) |  | 

## Methods

### NewPostSetBusinessAccountGiftSettingsRequest

`func NewPostSetBusinessAccountGiftSettingsRequest(businessConnectionId string, showGiftButton bool, acceptedGiftTypes AcceptedGiftTypes, ) *PostSetBusinessAccountGiftSettingsRequest`

NewPostSetBusinessAccountGiftSettingsRequest instantiates a new PostSetBusinessAccountGiftSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetBusinessAccountGiftSettingsRequestWithDefaults

`func NewPostSetBusinessAccountGiftSettingsRequestWithDefaults() *PostSetBusinessAccountGiftSettingsRequest`

NewPostSetBusinessAccountGiftSettingsRequestWithDefaults instantiates a new PostSetBusinessAccountGiftSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostSetBusinessAccountGiftSettingsRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostSetBusinessAccountGiftSettingsRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostSetBusinessAccountGiftSettingsRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetShowGiftButton

`func (o *PostSetBusinessAccountGiftSettingsRequest) GetShowGiftButton() bool`

GetShowGiftButton returns the ShowGiftButton field if non-nil, zero value otherwise.

### GetShowGiftButtonOk

`func (o *PostSetBusinessAccountGiftSettingsRequest) GetShowGiftButtonOk() (*bool, bool)`

GetShowGiftButtonOk returns a tuple with the ShowGiftButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowGiftButton

`func (o *PostSetBusinessAccountGiftSettingsRequest) SetShowGiftButton(v bool)`

SetShowGiftButton sets ShowGiftButton field to given value.


### GetAcceptedGiftTypes

`func (o *PostSetBusinessAccountGiftSettingsRequest) GetAcceptedGiftTypes() AcceptedGiftTypes`

GetAcceptedGiftTypes returns the AcceptedGiftTypes field if non-nil, zero value otherwise.

### GetAcceptedGiftTypesOk

`func (o *PostSetBusinessAccountGiftSettingsRequest) GetAcceptedGiftTypesOk() (*AcceptedGiftTypes, bool)`

GetAcceptedGiftTypesOk returns a tuple with the AcceptedGiftTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedGiftTypes

`func (o *PostSetBusinessAccountGiftSettingsRequest) SetAcceptedGiftTypes(v AcceptedGiftTypes)`

SetAcceptedGiftTypes sets AcceptedGiftTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


