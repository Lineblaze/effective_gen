# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfoGet**](DefaultAPI.md#InfoGet) | **Get** /info | 
[**SongsGet**](DefaultAPI.md#SongsGet) | **Get** /songs | 
[**SongsPost**](DefaultAPI.md#SongsPost) | **Post** /songs | 
[**SongsSongIdDelete**](DefaultAPI.md#SongsSongIdDelete) | **Delete** /songs/{songId} | 
[**SongsSongIdPatch**](DefaultAPI.md#SongsSongIdPatch) | **Patch** /songs/{songId} | 
[**SongsTextPost**](DefaultAPI.md#SongsTextPost) | **Post** /songs/text | 



## InfoGet

> SongDetail InfoGet(ctx).Group(group).Song(song).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	group := "group_example" // string | 
	song := "song_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InfoGet(context.Background()).Group(group).Song(song).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfoGet`: SongDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string** |  | 
 **song** | **string** |  | 

### Return type

[**SongDetail**](SongDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SongsGet

> []Song SongsGet(ctx).GetSongsBody(getSongsBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	getSongsBody := *openapiclient.NewGetSongsBody() // GetSongsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SongsGet(context.Background()).GetSongsBody(getSongsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SongsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SongsGet`: []Song
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SongsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSongsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSongsBody** | [**GetSongsBody**](GetSongsBody.md) |  | 

### Return type

[**[]Song**](Song.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SongsPost

> Song SongsPost(ctx).CreateSongBody(createSongBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createSongBody := *openapiclient.NewCreateSongBody("Muse", "Supermassive Black Hole") // CreateSongBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SongsPost(context.Background()).CreateSongBody(createSongBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SongsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SongsPost`: Song
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SongsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSongsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSongBody** | [**CreateSongBody**](CreateSongBody.md) |  | 

### Return type

[**Song**](Song.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SongsSongIdDelete

> SongsSongIdDelete(ctx, songId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	songId := "songId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SongsSongIdDelete(context.Background(), songId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SongsSongIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**songId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSongsSongIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SongsSongIdPatch

> Song SongsSongIdPatch(ctx, songId).UpdateSongBody(updateSongBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	songId := "songId_example" // string | 
	updateSongBody := *openapiclient.NewUpdateSongBody() // UpdateSongBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SongsSongIdPatch(context.Background(), songId).UpdateSongBody(updateSongBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SongsSongIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SongsSongIdPatch`: Song
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SongsSongIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**songId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSongsSongIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSongBody** | [**UpdateSongBody**](UpdateSongBody.md) |  | 

### Return type

[**Song**](Song.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SongsTextPost

> SongsTextPost200Response SongsTextPost(ctx).GetSongTextBody(getSongTextBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	getSongTextBody := *openapiclient.NewGetSongTextBody("Muse", "Supermassive Black Hole") // GetSongTextBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SongsTextPost(context.Background()).GetSongTextBody(getSongTextBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SongsTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SongsTextPost`: SongsTextPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SongsTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSongsTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSongTextBody** | [**GetSongTextBody**](GetSongTextBody.md) |  | 

### Return type

[**SongsTextPost200Response**](SongsTextPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

