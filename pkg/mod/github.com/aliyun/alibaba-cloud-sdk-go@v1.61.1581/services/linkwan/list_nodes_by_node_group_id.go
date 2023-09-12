package linkwan

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

// ListNodesByNodeGroupId invokes the linkwan.ListNodesByNodeGroupId API synchronously
func (client *Client) ListNodesByNodeGroupId(request *ListNodesByNodeGroupIdRequest) (response *ListNodesByNodeGroupIdResponse, err error) {
	response = CreateListNodesByNodeGroupIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodesByNodeGroupIdWithChan invokes the linkwan.ListNodesByNodeGroupId API asynchronously
func (client *Client) ListNodesByNodeGroupIdWithChan(request *ListNodesByNodeGroupIdRequest) (<-chan *ListNodesByNodeGroupIdResponse, <-chan error) {
	responseChan := make(chan *ListNodesByNodeGroupIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodesByNodeGroupId(request)
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

// ListNodesByNodeGroupIdWithCallback invokes the linkwan.ListNodesByNodeGroupId API asynchronously
func (client *Client) ListNodesByNodeGroupIdWithCallback(request *ListNodesByNodeGroupIdRequest, callback func(response *ListNodesByNodeGroupIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodesByNodeGroupIdResponse
		var err error
		defer close(result)
		response, err = client.ListNodesByNodeGroupId(request)
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

// ListNodesByNodeGroupIdRequest is the request struct for api ListNodesByNodeGroupId
type ListNodesByNodeGroupIdRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	FuzzyDevEui   string           `position:"Query" name:"FuzzyDevEui"`
	Limit         requests.Integer `position:"Query" name:"Limit"`
	Offset        requests.Integer `position:"Query" name:"Offset"`
	Ascending     requests.Boolean `position:"Query" name:"Ascending"`
	NodeGroupId   string           `position:"Query" name:"NodeGroupId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	SortingField  string           `position:"Query" name:"SortingField"`
}

// ListNodesByNodeGroupIdResponse is the response struct for api ListNodesByNodeGroupId
type ListNodesByNodeGroupIdResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Success   bool                         `json:"Success" xml:"Success"`
	Data      DataInListNodesByNodeGroupId `json:"Data" xml:"Data"`
}

// CreateListNodesByNodeGroupIdRequest creates a request to invoke ListNodesByNodeGroupId API
func CreateListNodesByNodeGroupIdRequest() (request *ListNodesByNodeGroupIdRequest) {
	request = &ListNodesByNodeGroupIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListNodesByNodeGroupId", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodesByNodeGroupIdResponse creates a response to parse from ListNodesByNodeGroupId response
func CreateListNodesByNodeGroupIdResponse() (response *ListNodesByNodeGroupIdResponse) {
	response = &ListNodesByNodeGroupIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
