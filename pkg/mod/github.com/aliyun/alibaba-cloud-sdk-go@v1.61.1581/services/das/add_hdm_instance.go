package das

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

// AddHDMInstance invokes the das.AddHDMInstance API synchronously
func (client *Client) AddHDMInstance(request *AddHDMInstanceRequest) (response *AddHDMInstanceResponse, err error) {
	response = CreateAddHDMInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// AddHDMInstanceWithChan invokes the das.AddHDMInstance API asynchronously
func (client *Client) AddHDMInstanceWithChan(request *AddHDMInstanceRequest) (<-chan *AddHDMInstanceResponse, <-chan error) {
	responseChan := make(chan *AddHDMInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddHDMInstance(request)
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

// AddHDMInstanceWithCallback invokes the das.AddHDMInstance API asynchronously
func (client *Client) AddHDMInstanceWithCallback(request *AddHDMInstanceRequest, callback func(response *AddHDMInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddHDMInstanceResponse
		var err error
		defer close(result)
		response, err = client.AddHDMInstance(request)
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

// AddHDMInstanceRequest is the request struct for api AddHDMInstance
type AddHDMInstanceRequest struct {
	*requests.RpcRequest
	NetworkType   string `position:"Query" name:"NetworkType"`
	Password      string `position:"Query" name:"Password"`
	Engine        string `position:"Query" name:"Engine"`
	Context       string `position:"Query" name:"__context"`
	Ip            string `position:"Query" name:"Ip"`
	InstanceAlias string `position:"Query" name:"InstanceAlias"`
	InstanceArea  string `position:"Query" name:"InstanceArea"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	Port          string `position:"Query" name:"Port"`
	FlushAccount  string `position:"Query" name:"FlushAccount"`
	VpcId         string `position:"Query" name:"VpcId"`
	Region        string `position:"Query" name:"Region"`
	Username      string `position:"Query" name:"Username"`
}

// AddHDMInstanceResponse is the response struct for api AddHDMInstance
type AddHDMInstanceResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Synchro   string `json:"Synchro" xml:"Synchro"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddHDMInstanceRequest creates a request to invoke AddHDMInstance API
func CreateAddHDMInstanceRequest() (request *AddHDMInstanceRequest) {
	request = &AddHDMInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "AddHDMInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateAddHDMInstanceResponse creates a response to parse from AddHDMInstance response
func CreateAddHDMInstanceResponse() (response *AddHDMInstanceResponse) {
	response = &AddHDMInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}