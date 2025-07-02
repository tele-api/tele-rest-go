# SetMyDefaultAdministratorRightsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rights** | Pointer to [**ChatAdministratorRights**](ChatAdministratorRights.md) |  | [optional] 
**ForChannels** | Pointer to **bool** | Pass *True* to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed. | [optional] 

## Methods

### NewSetMyDefaultAdministratorRightsRequest

`func NewSetMyDefaultAdministratorRightsRequest() *SetMyDefaultAdministratorRightsRequest`

NewSetMyDefaultAdministratorRightsRequest instantiates a new SetMyDefaultAdministratorRightsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMyDefaultAdministratorRightsRequestWithDefaults

`func NewSetMyDefaultAdministratorRightsRequestWithDefaults() *SetMyDefaultAdministratorRightsRequest`

NewSetMyDefaultAdministratorRightsRequestWithDefaults instantiates a new SetMyDefaultAdministratorRightsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRights

`func (o *SetMyDefaultAdministratorRightsRequest) GetRights() ChatAdministratorRights`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *SetMyDefaultAdministratorRightsRequest) GetRightsOk() (*ChatAdministratorRights, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *SetMyDefaultAdministratorRightsRequest) SetRights(v ChatAdministratorRights)`

SetRights sets Rights field to given value.

### HasRights

`func (o *SetMyDefaultAdministratorRightsRequest) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetForChannels

`func (o *SetMyDefaultAdministratorRightsRequest) GetForChannels() bool`

GetForChannels returns the ForChannels field if non-nil, zero value otherwise.

### GetForChannelsOk

`func (o *SetMyDefaultAdministratorRightsRequest) GetForChannelsOk() (*bool, bool)`

GetForChannelsOk returns a tuple with the ForChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChannels

`func (o *SetMyDefaultAdministratorRightsRequest) SetForChannels(v bool)`

SetForChannels sets ForChannels field to given value.

### HasForChannels

`func (o *SetMyDefaultAdministratorRightsRequest) HasForChannels() bool`

HasForChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


