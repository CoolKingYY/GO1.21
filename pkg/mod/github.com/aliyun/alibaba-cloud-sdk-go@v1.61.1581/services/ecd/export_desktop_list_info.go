package ecd

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

// ExportDesktopListInfo invokes the ecd.ExportDesktopListInfo API synchronously
func (client *Client) ExportDesktopListInfo(request *ExportDesktopListInfoRequest) (response *ExportDesktopListInfoResponse, err error) {
	response = CreateExportDesktopListInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ExportDesktopListInfoWithChan invokes the ecd.ExportDesktopListInfo API asynchronously
func (client *Client) ExportDesktopListInfoWithChan(request *ExportDesktopListInfoRequest) (<-chan *ExportDesktopListInfoResponse, <-chan error) {
	responseChan := make(chan *ExportDesktopListInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportDesktopListInfo(request)
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

// ExportDesktopListInfoWithCallback invokes the ecd.ExportDesktopListInfo API asynchronously
func (client *Client) ExportDesktopListInfoWithCallback(request *ExportDesktopListInfoRequest, callback func(response *ExportDesktopListInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportDesktopListInfoResponse
		var err error
		defer close(result)
		response, err = client.ExportDesktopListInfo(request)
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

// ExportDesktopListInfoRequest is the request struct for api ExportDesktopListInfo
type ExportDesktopListInfoRequest struct {
	*requests.RpcRequest
	OfficeSiteId  string           `position:"Query" name:"OfficeSiteId"`
	DesktopStatus string           `position:"Query" name:"DesktopStatus"`
	NextToken     string           `position:"Query" name:"NextToken"`
	DirectoryId   string           `position:"Query" name:"DirectoryId"`
	EndUserId     *[]string        `position:"Query" name:"EndUserId"  type:"Repeated"`
	DesktopId     *[]string        `position:"Query" name:"DesktopId"  type:"Repeated"`
	DesktopName   string           `position:"Query" name:"DesktopName"`
	GroupId       string           `position:"Query" name:"GroupId"`
	ExpiredTime   string           `position:"Query" name:"ExpiredTime"`
	MaxResults    requests.Integer `position:"Query" name:"MaxResults"`
	ChargeType    string           `position:"Query" name:"ChargeType"`
	PolicyGroupId string           `position:"Query" name:"PolicyGroupId"`
	UserName      string           `position:"Query" name:"UserName"`
}

// ExportDesktopListInfoResponse is the response struct for api ExportDesktopListInfo
type ExportDesktopListInfoResponse struct {
	*responses.BaseResponse
	Url       string `json:"Url" xml:"Url"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExportDesktopListInfoRequest creates a request to invoke ExportDesktopListInfo API
func CreateExportDesktopListInfoRequest() (request *ExportDesktopListInfoRequest) {
	request = &ExportDesktopListInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ExportDesktopListInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateExportDesktopListInfoResponse creates a response to parse from ExportDesktopListInfo response
func CreateExportDesktopListInfoResponse() (response *ExportDesktopListInfoResponse) {
	response = &ExportDesktopListInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
