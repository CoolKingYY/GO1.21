package ivision

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

// StartStreamPredict invokes the ivision.StartStreamPredict API synchronously
func (client *Client) StartStreamPredict(request *StartStreamPredictRequest) (response *StartStreamPredictResponse, err error) {
	response = CreateStartStreamPredictResponse()
	err = client.DoAction(request, response)
	return
}

// StartStreamPredictWithChan invokes the ivision.StartStreamPredict API asynchronously
func (client *Client) StartStreamPredictWithChan(request *StartStreamPredictRequest) (<-chan *StartStreamPredictResponse, <-chan error) {
	responseChan := make(chan *StartStreamPredictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartStreamPredict(request)
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

// StartStreamPredictWithCallback invokes the ivision.StartStreamPredict API asynchronously
func (client *Client) StartStreamPredictWithCallback(request *StartStreamPredictRequest, callback func(response *StartStreamPredictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartStreamPredictResponse
		var err error
		defer close(result)
		response, err = client.StartStreamPredict(request)
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

// StartStreamPredictRequest is the request struct for api StartStreamPredict
type StartStreamPredictRequest struct {
	*requests.RpcRequest
	PredictId string           `position:"Query" name:"PredictId"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// StartStreamPredictResponse is the response struct for api StartStreamPredict
type StartStreamPredictResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PredictId string `json:"PredictId" xml:"PredictId"`
}

// CreateStartStreamPredictRequest creates a request to invoke StartStreamPredict API
func CreateStartStreamPredictRequest() (request *StartStreamPredictRequest) {
	request = &StartStreamPredictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "StartStreamPredict", "ivision", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartStreamPredictResponse creates a response to parse from StartStreamPredict response
func CreateStartStreamPredictResponse() (response *StartStreamPredictResponse) {
	response = &StartStreamPredictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
