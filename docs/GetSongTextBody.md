# GetSongTextBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** |  | 
**Song** | **string** |  | 
**Limit** | Pointer to **int32** | Limit the number of verses returned | [optional] [default to 5]
**Offset** | Pointer to **int32** | Offset for pagination | [optional] [default to 0]

## Methods

### NewGetSongTextBody

`func NewGetSongTextBody(group string, song string, ) *GetSongTextBody`

NewGetSongTextBody instantiates a new GetSongTextBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSongTextBodyWithDefaults

`func NewGetSongTextBodyWithDefaults() *GetSongTextBody`

NewGetSongTextBodyWithDefaults instantiates a new GetSongTextBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GetSongTextBody) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetSongTextBody) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetSongTextBody) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetSong

`func (o *GetSongTextBody) GetSong() string`

GetSong returns the Song field if non-nil, zero value otherwise.

### GetSongOk

`func (o *GetSongTextBody) GetSongOk() (*string, bool)`

GetSongOk returns a tuple with the Song field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSong

`func (o *GetSongTextBody) SetSong(v string)`

SetSong sets Song field to given value.


### GetLimit

`func (o *GetSongTextBody) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetSongTextBody) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetSongTextBody) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetSongTextBody) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetSongTextBody) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetSongTextBody) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetSongTextBody) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetSongTextBody) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


