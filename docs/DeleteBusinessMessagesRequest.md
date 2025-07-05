# DeleteBusinessMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection on behalf of which to delete the messages | 
**MessageIds** | **[]int32** | A JSON-serialized list of 1-100 identifiers of messages to delete. All messages must be from the same chat. See [deleteMessage](https://core.telegram.org/bots/api/#deletemessage) for limitations on which messages can be deleted | 

## Methods

### NewDeleteBusinessMessagesRequest

`func NewDeleteBusinessMessagesRequest(businessConnectionId string, messageIds []int32, ) *DeleteBusinessMessagesRequest`

NewDeleteBusinessMessagesRequest instantiates a new DeleteBusinessMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBusinessMessagesRequestWithDefaults

`func NewDeleteBusinessMessagesRequestWithDefaults() *DeleteBusinessMessagesRequest`

NewDeleteBusinessMessagesRequestWithDefaults instantiates a new DeleteBusinessMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *DeleteBusinessMessagesRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *DeleteBusinessMessagesRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *DeleteBusinessMessagesRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetMessageIds

`func (o *DeleteBusinessMessagesRequest) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *DeleteBusinessMessagesRequest) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *DeleteBusinessMessagesRequest) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


