# PostDeleteBusinessMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection on behalf of which to delete the messages | 
**MessageIds** | **[]int32** | A JSON-serialized list of 1-100 identifiers of messages to delete. All messages must be from the same chat. See [deleteMessage](https://core.telegram.org/bots/api/#deletemessage) for limitations on which messages can be deleted | 

## Methods

### NewPostDeleteBusinessMessagesRequest

`func NewPostDeleteBusinessMessagesRequest(businessConnectionId string, messageIds []int32, ) *PostDeleteBusinessMessagesRequest`

NewPostDeleteBusinessMessagesRequest instantiates a new PostDeleteBusinessMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteBusinessMessagesRequestWithDefaults

`func NewPostDeleteBusinessMessagesRequestWithDefaults() *PostDeleteBusinessMessagesRequest`

NewPostDeleteBusinessMessagesRequestWithDefaults instantiates a new PostDeleteBusinessMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostDeleteBusinessMessagesRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostDeleteBusinessMessagesRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostDeleteBusinessMessagesRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetMessageIds

`func (o *PostDeleteBusinessMessagesRequest) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *PostDeleteBusinessMessagesRequest) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *PostDeleteBusinessMessagesRequest) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


