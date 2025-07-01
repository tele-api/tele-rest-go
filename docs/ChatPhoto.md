# ChatPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmallFileId** | **string** | File identifier of small (160x160) chat photo. This file\\_id can be used only for photo download and only for as long as the photo is not changed. | 
**SmallFileUniqueId** | **string** | Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**BigFileId** | **string** | File identifier of big (640x640) chat photo. This file\\_id can be used only for photo download and only for as long as the photo is not changed. | 
**BigFileUniqueId** | **string** | Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 

## Methods

### NewChatPhoto

`func NewChatPhoto(smallFileId string, smallFileUniqueId string, bigFileId string, bigFileUniqueId string, ) *ChatPhoto`

NewChatPhoto instantiates a new ChatPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPhotoWithDefaults

`func NewChatPhotoWithDefaults() *ChatPhoto`

NewChatPhotoWithDefaults instantiates a new ChatPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmallFileId

`func (o *ChatPhoto) GetSmallFileId() string`

GetSmallFileId returns the SmallFileId field if non-nil, zero value otherwise.

### GetSmallFileIdOk

`func (o *ChatPhoto) GetSmallFileIdOk() (*string, bool)`

GetSmallFileIdOk returns a tuple with the SmallFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallFileId

`func (o *ChatPhoto) SetSmallFileId(v string)`

SetSmallFileId sets SmallFileId field to given value.


### GetSmallFileUniqueId

`func (o *ChatPhoto) GetSmallFileUniqueId() string`

GetSmallFileUniqueId returns the SmallFileUniqueId field if non-nil, zero value otherwise.

### GetSmallFileUniqueIdOk

`func (o *ChatPhoto) GetSmallFileUniqueIdOk() (*string, bool)`

GetSmallFileUniqueIdOk returns a tuple with the SmallFileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallFileUniqueId

`func (o *ChatPhoto) SetSmallFileUniqueId(v string)`

SetSmallFileUniqueId sets SmallFileUniqueId field to given value.


### GetBigFileId

`func (o *ChatPhoto) GetBigFileId() string`

GetBigFileId returns the BigFileId field if non-nil, zero value otherwise.

### GetBigFileIdOk

`func (o *ChatPhoto) GetBigFileIdOk() (*string, bool)`

GetBigFileIdOk returns a tuple with the BigFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigFileId

`func (o *ChatPhoto) SetBigFileId(v string)`

SetBigFileId sets BigFileId field to given value.


### GetBigFileUniqueId

`func (o *ChatPhoto) GetBigFileUniqueId() string`

GetBigFileUniqueId returns the BigFileUniqueId field if non-nil, zero value otherwise.

### GetBigFileUniqueIdOk

`func (o *ChatPhoto) GetBigFileUniqueIdOk() (*string, bool)`

GetBigFileUniqueIdOk returns a tuple with the BigFileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigFileUniqueId

`func (o *ChatPhoto) SetBigFileUniqueId(v string)`

SetBigFileUniqueId sets BigFileUniqueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


