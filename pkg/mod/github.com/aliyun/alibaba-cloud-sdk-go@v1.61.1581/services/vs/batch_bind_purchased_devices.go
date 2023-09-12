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

// BatchBindPurchasedDevices invokes the vs.BatchBindPurchasedDevices API synchronously
func (client *Client) BatchBindPurchasedDevices(request *BatchBindPurchasedDevicesRequest) (response *BatchBindPurchasedDevicesResponse, err error) {
	response = CreateBatchBindPurchasedDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// BatchBindPurchasedDevicesWithChan invokes the vs.BatchBindPurchasedDevices API asynchronously
func (client *Client) BatchBindPurchasedDevicesWithChan(request *BatchBindPurchasedDevicesRequest) (<-chan *BatchBindPurchasedDevicesResponse, <-chan error) {
	responseChan := make(chan *BatchBindPurchasedDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchBindPurchasedDevices(request)
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

// BatchBindPurchasedDevicesWithCallback invokes the vs.BatchBindPurchasedDevices API asynchronously
func (client *Client) BatchBindPurchasedDevicesWithCallback(request *BatchBindPurchasedDevicesRequest, callback func(response *BatchBindPurchasedDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchBindPurchasedDevicesResponse
		var err error
		defer close(result)
		response, err = client.BatchBindPurchasedDevices(request)
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

// BatchBindPurchasedDevicesRequest is the request struct for api BatchBindPurchasedDevices
type BatchBindPurchasedDevicesRequest struct {
	*requests.RpcRequest
	ShowLog  string           `position:"Query" name:"ShowLog"`
	GroupId  string           `position:"Query" name:"GroupId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	DeviceId string           `position:"Query" name:"DeviceId"`
	Region   string           `position:"Query" name:"Region"`
}

// BatchBindPurchasedDevicesResponse is the response struct for api BatchBindPurchasedDevices
type BatchBindPurchasedDevicesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateBatchBindPurchasedDevicesRequest creates a request to invoke BatchBindPurchasedDevices API
func CreateBatchBindPurchasedDevicesRequest() (request *BatchBindPurchasedDevicesRequest) {
	request = &BatchBindPurchasedDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchBindPurchasedDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchBindPurchasedDevicesResponse creates a response to parse from BatchBindPurchasedDevices response
func CreateBatchBindPurchasedDevicesResponse() (response *BatchBindPurchasedDevicesResponse) {
	response = &BatchBindPurchasedDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
