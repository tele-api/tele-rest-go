# GetMyDefaultAdministratorRightsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForChannels** | Pointer to **bool** | Pass *True* to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned. | [optional] 

## Methods

### NewGetMyDefaultAdministratorRightsPostRequest

`func NewGetMyDefaultAdministratorRightsPostRequest() *GetMyDefaultAdministratorRightsPostRequest`

NewGetMyDefaultAdministratorRightsPostRequest instantiates a new GetMyDefaultAdministratorRightsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyDefaultAdministratorRightsPostRequestWithDefaults

`func NewGetMyDefaultAdministratorRightsPostRequestWithDefaults() *GetMyDefaultAdministratorRightsPostRequest`

NewGetMyDefaultAdministratorRightsPostRequestWithDefaults instantiates a new GetMyDefaultAdministratorRightsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForChannels

`func (o *GetMyDefaultAdministratorRightsPostRequest) GetForChannels() bool`

GetForChannels returns the ForChannels field if non-nil, zero value otherwise.

### GetForChannelsOk

`func (o *GetMyDefaultAdministratorRightsPostRequest) GetForChannelsOk() (*bool, bool)`

GetForChannelsOk returns a tuple with the ForChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChannels

`func (o *GetMyDefaultAdministratorRightsPostRequest) SetForChannels(v bool)`

SetForChannels sets ForChannels field to given value.

### HasForChannels

`func (o *GetMyDefaultAdministratorRightsPostRequest) HasForChannels() bool`

HasForChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


