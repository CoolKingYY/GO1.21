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

// DescribeStreamURL invokes the vs.DescribeStreamURL API synchronously
func (client *Client) DescribeStreamURL(request *DescribeStreamURLRequest) (response *DescribeStreamURLResponse, err error) {
	response = CreateDescribeStreamURLResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStreamURLWithChan invokes the vs.DescribeStreamURL API asynchronously
func (client *Client) DescribeStreamURLWithChan(request *DescribeStreamURLRequest) (<-chan *DescribeStreamURLResponse, <-chan error) {
	responseChan := make(chan *DescribeStreamURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStreamURL(request)
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

// DescribeStreamURLWithCallback invokes the vs.DescribeStreamURL API asynchronously
func (client *Client) DescribeStreamURLWithCallback(request *DescribeStreamURLRequest, callback func(response *DescribeStreamURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStreamURLResponse
		var err error
		defer close(result)
		response, err = client.DescribeStreamURL(request)
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

// DescribeStreamURLRequest is the request struct for api DescribeStreamURL
type DescribeStreamURLRequest struct {
	*requests.RpcRequest
	AuthKey     string           `position:"Query" name:"AuthKey"`
	Auth        requests.Boolean `position:"Query" name:"Auth"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
	Type        string           `position:"Query" name:"Type"`
	OutHostType string           `position:"Query" name:"OutHostType"`
	Id          string           `position:"Query" name:"Id"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OutProtocol string           `position:"Query" name:"OutProtocol"`
	Transcode   string           `position:"Query" name:"Transcode"`
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	Expire      requests.Integer `position:"Query" name:"Expire"`
	Location    string           `position:"Query" name:"Location"`
}

// DescribeStreamURLResponse is the response struct for api DescribeStreamURL
type DescribeStreamURLResponse struct {
	*responses.BaseResponse
	Url        string `json:"Url" xml:"Url"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ExpireTime int64  `json:"ExpireTime" xml:"ExpireTime"`
}

// CreateDescribeStreamURLRequest creates a request to invoke DescribeStreamURL API
func CreateDescribeStreamURLRequest() (request *DescribeStreamURLRequest) {
	request = &DescribeStreamURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeStreamURL", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStreamURLResponse creates a response to parse from DescribeStreamURL response
func CreateDescribeStreamURLResponse() (response *DescribeStreamURLResponse) {
	response = &DescribeStreamURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}