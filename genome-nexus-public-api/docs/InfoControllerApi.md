# \InfoControllerAPI

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVersionGET**](InfoControllerAPI.md#FetchVersionGET) | **Get** /version | Retrieve Genome Nexus Version



## FetchVersionGET

> AggregateSourceInfo FetchVersionGET(ctx).Execute()

Retrieve Genome Nexus Version

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfoControllerAPI.FetchVersionGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoControllerAPI.FetchVersionGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchVersionGET`: AggregateSourceInfo
	fmt.Fprintf(os.Stdout, "Response from `InfoControllerAPI.FetchVersionGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVersionGETRequest struct via the builder pattern


### Return type

[**AggregateSourceInfo**](AggregateSourceInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

