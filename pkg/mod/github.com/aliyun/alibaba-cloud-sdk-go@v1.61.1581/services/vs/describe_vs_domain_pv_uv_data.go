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

// DescribeVsDomainPvUvData invokes the vs.DescribeVsDomainPvUvData API synchronously
func (client *Client) DescribeVsDomainPvUvData(request *DescribeVsDomainPvUvDataRequest) (response *DescribeVsDomainPvUvDataResponse, err error) {
	response = CreateDescribeVsDomainPvUvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsDomainPvUvDataWithChan invokes the vs.DescribeVsDomainPvUvData API asynchronously
func (client *Client) DescribeVsDomainPvUvDataWithChan(request *DescribeVsDomainPvUvDataRequest) (<-chan *DescribeVsDomainPvUvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVsDomainPvUvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsDomainPvUvData(request)
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

// DescribeVsDomainPvUvDataWithCallback invokes the vs.DescribeVsDomainPvUvData API asynchronously
func (client *Client) DescribeVsDomainPvUvDataWithCallback(request *DescribeVsDomainPvUvDataRequest, callback func(response *DescribeVsDomainPvUvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsDomainPvUvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsDomainPvUvData(request)
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

// DescribeVsDomainPvUvDataRequest is the request struct for api DescribeVsDomainPvUvData
type DescribeVsDomainPvUvDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsDomainPvUvDataResponse is the response struct for api DescribeVsDomainPvUvData
type DescribeVsDomainPvUvDataResponse struct {
	*responses.BaseResponse
	EndTime       string        `json:"EndTime" xml:"EndTime"`
	StartTime     string        `json:"StartTime" xml:"StartTime"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	DomainName    string        `json:"DomainName" xml:"DomainName"`
	DataInterval  string        `json:"DataInterval" xml:"DataInterval"`
	PvUvDataInfos PvUvDataInfos `json:"PvUvDataInfos" xml:"PvUvDataInfos"`
}

// CreateDescribeVsDomainPvUvDataRequest creates a request to invoke DescribeVsDomainPvUvData API
func CreateDescribeVsDomainPvUvDataRequest() (request *DescribeVsDomainPvUvDataRequest) {
	request = &DescribeVsDomainPvUvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsDomainPvUvData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVsDomainPvUvDataResponse creates a response to parse from DescribeVsDomainPvUvData response
func CreateDescribeVsDomainPvUvDataResponse() (response *DescribeVsDomainPvUvDataResponse) {
	response = &DescribeVsDomainPvUvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
