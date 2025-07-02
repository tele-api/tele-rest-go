# SetBusinessAccountGiftSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**ShowGiftButton** | **bool** | Pass True, if a button for sending a gift to the user or by the business account must always be shown in the input field | 
**AcceptedGiftTypes** | [**AcceptedGiftTypes**](AcceptedGiftTypes.md) |  | 

## Methods

### NewSetBusinessAccountGiftSettingsRequest

`func NewSetBusinessAccountGiftSettingsRequest(businessConnectionId string, showGiftButton bool, acceptedGiftTypes AcceptedGiftTypes, ) *SetBusinessAccountGiftSettingsRequest`

NewSetBusinessAccountGiftSettingsRequest instantiates a new SetBusinessAccountGiftSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBusinessAccountGiftSettingsRequestWithDefaults

`func NewSetBusinessAccountGiftSettingsRequestWithDefaults() *SetBusinessAccountGiftSettingsRequest`

NewSetBusinessAccountGiftSettingsRequestWithDefaults instantiates a new SetBusinessAccountGiftSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *SetBusinessAccountGiftSettingsRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *SetBusinessAccountGiftSettingsRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *SetBusinessAccountGiftSettingsRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetShowGiftButton

`func (o *SetBusinessAccountGiftSettingsRequest) GetShowGiftButton() bool`

GetShowGiftButton returns the ShowGiftButton field if non-nil, zero value otherwise.

### GetShowGiftButtonOk

`func (o *SetBusinessAccountGiftSettingsRequest) GetShowGiftButtonOk() (*bool, bool)`

GetShowGiftButtonOk returns a tuple with the ShowGiftButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowGiftButton

`func (o *SetBusinessAccountGiftSettingsRequest) SetShowGiftButton(v bool)`

SetShowGiftButton sets ShowGiftButton field to given value.


### GetAcceptedGiftTypes

`func (o *SetBusinessAccountGiftSettingsRequest) GetAcceptedGiftTypes() AcceptedGiftTypes`

GetAcceptedGiftTypes returns the AcceptedGiftTypes field if non-nil, zero value otherwise.

### GetAcceptedGiftTypesOk

`func (o *SetBusinessAccountGiftSettingsRequest) GetAcceptedGiftTypesOk() (*AcceptedGiftTypes, bool)`

GetAcceptedGiftTypesOk returns a tuple with the AcceptedGiftTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedGiftTypes

`func (o *SetBusinessAccountGiftSettingsRequest) SetAcceptedGiftTypes(v AcceptedGiftTypes)`

SetAcceptedGiftTypes sets AcceptedGiftTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


