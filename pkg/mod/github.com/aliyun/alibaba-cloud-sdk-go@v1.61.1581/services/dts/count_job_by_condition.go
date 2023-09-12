package dts

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

// CountJobByCondition invokes the dts.CountJobByCondition API synchronously
func (client *Client) CountJobByCondition(request *CountJobByConditionRequest) (response *CountJobByConditionResponse, err error) {
	response = CreateCountJobByConditionResponse()
	err = client.DoAction(request, response)
	return
}

// CountJobByConditionWithChan invokes the dts.CountJobByCondition API asynchronously
func (client *Client) CountJobByConditionWithChan(request *CountJobByConditionRequest) (<-chan *CountJobByConditionResponse, <-chan error) {
	responseChan := make(chan *CountJobByConditionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CountJobByCondition(request)
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

// CountJobByConditionWithCallback invokes the dts.CountJobByCondition API asynchronously
func (client *Client) CountJobByConditionWithCallback(request *CountJobByConditionRequest, callback func(response *CountJobByConditionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CountJobByConditionResponse
		var err error
		defer close(result)
		response, err = client.CountJobByCondition(request)
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

// CountJobByConditionRequest is the request struct for api CountJobByCondition
type CountJobByConditionRequest struct {
	*requests.RpcRequest
	Type       string `position:"Query" name:"Type"`
	SrcDbType  string `position:"Query" name:"SrcDbType"`
	GroupId    string `position:"Query" name:"GroupId"`
	Params     string `position:"Query" name:"Params"`
	JobType    string `position:"Query" name:"JobType"`
	DestDbType string `position:"Query" name:"DestDbType"`
	Region     string `position:"Query" name:"Region"`
	Status     string `position:"Query" name:"Status"`
}

// CountJobByConditionResponse is the response struct for api CountJobByCondition
type CountJobByConditionResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode   int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrCode          string `json:"ErrCode" xml:"ErrCode"`
	Success          bool   `json:"Success" xml:"Success"`
	ErrMessage       string `json:"ErrMessage" xml:"ErrMessage"`
	DynamicMessage   string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode      string `json:"DynamicCode" xml:"DynamicCode"`
	TotalRecordCount int64  `json:"TotalRecordCount" xml:"TotalRecordCount"`
}

// CreateCountJobByConditionRequest creates a request to invoke CountJobByCondition API
func CreateCountJobByConditionRequest() (request *CountJobByConditionRequest) {
	request = &CountJobByConditionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "CountJobByCondition", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCountJobByConditionResponse creates a response to parse from CountJobByCondition response
func CreateCountJobByConditionResponse() (response *CountJobByConditionResponse) {
	response = &CountJobByConditionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}