package ddoscoo

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListTagKeys invokes the ddoscoo.ListTagKeys API synchronously
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (response *ListTagKeysResponse, err error) {
	response = CreateListTagKeysResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagKeysWithChan invokes the ddoscoo.ListTagKeys API asynchronously
func (client *Client) ListTagKeysWithChan(request *ListTagKeysRequest) (<-chan *ListTagKeysResponse, <-chan error) {
	responseChan := make(chan *ListTagKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagKeys(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListTagKeysWithCallback invokes the ddoscoo.ListTagKeys API asynchronously
func (client *Client) ListTagKeysWithCallback(request *ListTagKeysRequest, callback func(response *ListTagKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagKeysResponse
		var err error
		defer close(result)
		response, err = client.ListTagKeys(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListTagKeysRequest is the request struct for api ListTagKeys
type ListTagKeysRequest struct {
	*requests.RpcRequest
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	ResourceType    string           `position:"Query" name:"ResourceType"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// ListTagKeysResponse is the response struct for api ListTagKeys
type ListTagKeysResponse struct {
	*responses.BaseResponse
	RequestId   string   `json:"RequestId" xml:"RequestId"`
	CurrentPage int      `json:"CurrentPage" xml:"CurrentPage"`
	PageSize    int      `json:"PageSize" xml:"PageSize"`
	TotalCount  int      `json:"TotalCount" xml:"TotalCount"`
	TagKeys     []TagKey `json:"TagKeys" xml:"TagKeys"`
}

// CreateListTagKeysRequest creates a request to invoke ListTagKeys API
func CreateListTagKeysRequest() (request *ListTagKeysRequest) {
	request = &ListTagKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "ListTagKeys", "", "")
	request.Method = requests.POST
	return
}

// CreateListTagKeysResponse creates a response to parse from ListTagKeys response
func CreateListTagKeysResponse() (response *ListTagKeysResponse) {
	response = &ListTagKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
