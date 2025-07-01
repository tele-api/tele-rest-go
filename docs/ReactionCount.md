# ReactionCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ReactionType**](ReactionType.md) |  | 
**TotalCount** | **int32** | Number of times the reaction was added | 

## Methods

### NewReactionCount

`func NewReactionCount(type_ ReactionType, totalCount int32, ) *ReactionCount`

NewReactionCount instantiates a new ReactionCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionCountWithDefaults

`func NewReactionCountWithDefaults() *ReactionCount`

NewReactionCountWithDefaults instantiates a new ReactionCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReactionCount) GetType() ReactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionCount) GetTypeOk() (*ReactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionCount) SetType(v ReactionType)`

SetType sets Type field to given value.


### GetTotalCount

`func (o *ReactionCount) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReactionCount) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReactionCount) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


