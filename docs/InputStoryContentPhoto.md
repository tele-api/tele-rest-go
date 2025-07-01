# InputStoryContentPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the content, must be *photo* | [default to "photo"]
**Photo** | **string** | The photo to post as a story. The photo must be of the size 1080x1920 and must not exceed 10 MB. The photo can&#39;t be reused and can only be uploaded as a new file, so you can pass “attach://\\&lt;file\\_attach\\_name\\&gt;” if the photo was uploaded using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt;. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 

## Methods

### NewInputStoryContentPhoto

`func NewInputStoryContentPhoto(type_ string, photo string, ) *InputStoryContentPhoto`

NewInputStoryContentPhoto instantiates a new InputStoryContentPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputStoryContentPhotoWithDefaults

`func NewInputStoryContentPhotoWithDefaults() *InputStoryContentPhoto`

NewInputStoryContentPhotoWithDefaults instantiates a new InputStoryContentPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputStoryContentPhoto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputStoryContentPhoto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputStoryContentPhoto) SetType(v string)`

SetType sets Type field to given value.


### GetPhoto

`func (o *InputStoryContentPhoto) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *InputStoryContentPhoto) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *InputStoryContentPhoto) SetPhoto(v string)`

SetPhoto sets Photo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


