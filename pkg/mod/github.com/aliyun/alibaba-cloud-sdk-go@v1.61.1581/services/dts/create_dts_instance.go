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

// CreateDtsInstance invokes the dts.CreateDtsInstance API synchronously
func (client *Client) CreateDtsInstance(request *CreateDtsInstanceRequest) (response *CreateDtsInstanceResponse, err error) {
	response = CreateCreateDtsInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDtsInstanceWithChan invokes the dts.CreateDtsInstance API asynchronously
func (client *Client) CreateDtsInstanceWithChan(request *CreateDtsInstanceRequest) (<-chan *CreateDtsInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDtsInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDtsInstance(request)
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

// CreateDtsInstanceWithCallback invokes the dts.CreateDtsInstance API asynchronously
func (client *Client) CreateDtsInstanceWithCallback(request *CreateDtsInstanceRequest, callback func(response *CreateDtsInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDtsInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDtsInstance(request)
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

// CreateDtsInstanceRequest is the request struct for api CreateDtsInstance
type CreateDtsInstanceRequest struct {
	*requests.RpcRequest
	AutoStart                     requests.Boolean `position:"Query" name:"AutoStart"`
	Type                          string           `position:"Query" name:"Type"`
	InstanceClass                 string           `position:"Query" name:"InstanceClass"`
	DatabaseCount                 requests.Integer `position:"Query" name:"DatabaseCount"`
	JobId                         string           `position:"Query" name:"JobId"`
	Du                            requests.Integer `position:"Query" name:"Du"`
	ComputeUnit                   requests.Integer `position:"Query" name:"ComputeUnit"`
	FeeType                       string           `position:"Query" name:"FeeType"`
	DestinationRegion             string           `position:"Query" name:"DestinationRegion"`
	Period                        string           `position:"Query" name:"Period"`
	DestinationEndpointEngineName string           `position:"Query" name:"DestinationEndpointEngineName"`
	Quantity                      requests.Integer `position:"Query" name:"Quantity"`
	AutoPay                       requests.Boolean `position:"Query" name:"AutoPay"`
	UsedTime                      requests.Integer `position:"Query" name:"UsedTime"`
	SyncArchitecture              string           `position:"Query" name:"SyncArchitecture"`
	PayType                       string           `position:"Query" name:"PayType"`
	SourceRegion                  string           `position:"Query" name:"SourceRegion"`
	SourceEndpointEngineName      string           `position:"Query" name:"SourceEndpointEngineName"`
}

// CreateDtsInstanceResponse is the response struct for api CreateDtsInstance
type CreateDtsInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	JobId      string `json:"JobId" xml:"JobId"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
}

// CreateCreateDtsInstanceRequest creates a request to invoke CreateDtsInstance API
func CreateCreateDtsInstanceRequest() (request *CreateDtsInstanceRequest) {
	request = &CreateDtsInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "CreateDtsInstance", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDtsInstanceResponse creates a response to parse from CreateDtsInstance response
func CreateCreateDtsInstanceResponse() (response *CreateDtsInstanceResponse) {
	response = &CreateDtsInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
