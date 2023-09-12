package vs

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

// DescribeStream invokes the vs.DescribeStream API synchronously
func (client *Client) DescribeStream(request *DescribeStreamRequest) (response *DescribeStreamResponse, err error) {
	response = CreateDescribeStreamResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStreamWithChan invokes the vs.DescribeStream API asynchronously
func (client *Client) DescribeStreamWithChan(request *DescribeStreamRequest) (<-chan *DescribeStreamResponse, <-chan error) {
	responseChan := make(chan *DescribeStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStream(request)
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

// DescribeStreamWithCallback invokes the vs.DescribeStream API asynchronously
func (client *Client) DescribeStreamWithCallback(request *DescribeStreamRequest, callback func(response *DescribeStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStreamResponse
		var err error
		defer close(result)
		response, err = client.DescribeStream(request)
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

// DescribeStreamRequest is the request struct for api DescribeStream
type DescribeStreamRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeStreamResponse is the response struct for api DescribeStream
type DescribeStreamResponse struct {
	*responses.BaseResponse
	Status      string `json:"Status" xml:"Status"`
	PlayDomain  string `json:"PlayDomain" xml:"PlayDomain"`
	Protocol    string `json:"Protocol" xml:"Protocol"`
	DeviceId    string `json:"DeviceId" xml:"DeviceId"`
	Height      int    `json:"Height" xml:"Height"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	GroupId     string `json:"GroupId" xml:"GroupId"`
	Width       int    `json:"Width" xml:"Width"`
	App         string `json:"App" xml:"App"`
	Enabled     bool   `json:"Enabled" xml:"Enabled"`
	Name        string `json:"Name" xml:"Name"`
	PushDomain  string `json:"PushDomain" xml:"PushDomain"`
	CreatedTime string `json:"CreatedTime" xml:"CreatedTime"`
	Id          string `json:"Id" xml:"Id"`
}

// CreateDescribeStreamRequest creates a request to invoke DescribeStream API
func CreateDescribeStreamRequest() (request *DescribeStreamRequest) {
	request = &DescribeStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeStream", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStreamResponse creates a response to parse from DescribeStream response
func CreateDescribeStreamResponse() (response *DescribeStreamResponse) {
	response = &DescribeStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
