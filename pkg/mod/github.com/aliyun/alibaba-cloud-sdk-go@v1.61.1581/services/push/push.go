package push

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

// Push invokes the push.Push API synchronously
func (client *Client) Push(request *PushRequest) (response *PushResponse, err error) {
	response = CreatePushResponse()
	err = client.DoAction(request, response)
	return
}

// PushWithChan invokes the push.Push API asynchronously
func (client *Client) PushWithChan(request *PushRequest) (<-chan *PushResponse, <-chan error) {
	responseChan := make(chan *PushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Push(request)
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

// PushWithCallback invokes the push.Push API asynchronously
func (client *Client) PushWithCallback(request *PushRequest, callback func(response *PushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushResponse
		var err error
		defer close(result)
		response, err = client.Push(request)
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

// PushRequest is the request struct for api Push
type PushRequest struct {
	*requests.RpcRequest
	AndroidNotificationBarType       requests.Integer `position:"Query" name:"AndroidNotificationBarType"`
	SmsSendPolicy                    requests.Integer `position:"Query" name:"SmsSendPolicy"`
	Body                             string           `position:"Query" name:"Body"`
	DeviceType                       string           `position:"Query" name:"DeviceType"`
	PushTime                         string           `position:"Query" name:"PushTime"`
	SendSpeed                        requests.Integer `position:"Query" name:"SendSpeed"`
	AndroidNotificationHuaweiChannel string           `position:"Query" name:"AndroidNotificationHuaweiChannel"`
	AndroidPopupActivity             string           `position:"Query" name:"AndroidPopupActivity"`
	IOSRemindBody                    string           `position:"Query" name:"iOSRemindBody"`
	Trim                             requests.Boolean `position:"Query" name:"Trim"`
	AndroidNotifyType                string           `position:"Query" name:"AndroidNotifyType"`
	AndroidPopupTitle                string           `position:"Query" name:"AndroidPopupTitle"`
	AndroidMessageHuaweiCategory     string           `position:"Query" name:"AndroidMessageHuaweiCategory"`
	IOSMusic                         string           `position:"Query" name:"iOSMusic"`
	IOSApnsEnv                       string           `position:"Query" name:"iOSApnsEnv"`
	IOSMutableContent                requests.Boolean `position:"Query" name:"iOSMutableContent"`
	AndroidNotificationBarPriority   requests.Integer `position:"Query" name:"AndroidNotificationBarPriority"`
	ExpireTime                       string           `position:"Query" name:"ExpireTime"`
	AndroidImageUrl                  string           `position:"Query" name:"AndroidImageUrl"`
	AndroidNotificationVivoChannel   string           `position:"Query" name:"AndroidNotificationVivoChannel"`
	IOSNotificationCategory          string           `position:"Query" name:"iOSNotificationCategory"`
	AndroidNotificationXiaomiChannel string           `position:"Query" name:"AndroidNotificationXiaomiChannel"`
	StoreOffline                     requests.Boolean `position:"Query" name:"StoreOffline"`
	SmsParams                        string           `position:"Query" name:"SmsParams"`
	IOSRelevanceScore                string           `position:"Query" name:"iOSRelevanceScore"`
	AndroidVivoPushMode              requests.Integer `position:"Query" name:"AndroidVivoPushMode"`
	AndroidInboxBody                 string           `position:"Query" name:"AndroidInboxBody"`
	JobKey                           string           `position:"Query" name:"JobKey"`
	AndroidOpenUrl                   string           `position:"Query" name:"AndroidOpenUrl"`
	AndroidXiaoMiNotifyBody          string           `position:"Query" name:"AndroidXiaoMiNotifyBody"`
	IOSSubtitle                      string           `position:"Query" name:"iOSSubtitle"`
	AndroidXiaomiBigPictureUrl       string           `position:"Query" name:"AndroidXiaomiBigPictureUrl"`
	IOSRemind                        requests.Boolean `position:"Query" name:"iOSRemind"`
	IOSNotificationThreadId          string           `position:"Query" name:"iOSNotificationThreadId"`
	AndroidMusic                     string           `position:"Query" name:"AndroidMusic"`
	IOSNotificationCollapseId        string           `position:"Query" name:"iOSNotificationCollapseId"`
	AndroidMessageHuaweiUrgency      string           `position:"Query" name:"AndroidMessageHuaweiUrgency"`
	PushType                         string           `position:"Query" name:"PushType"`
	IOSInterruptionLevel             string           `position:"Query" name:"iOSInterruptionLevel"`
	AndroidExtParameters             string           `position:"Query" name:"AndroidExtParameters"`
	IOSBadge                         requests.Integer `position:"Query" name:"iOSBadge"`
	AndroidBigBody                   string           `position:"Query" name:"AndroidBigBody"`
	IOSBadgeAutoIncrement            requests.Boolean `position:"Query" name:"iOSBadgeAutoIncrement"`
	AndroidOpenType                  string           `position:"Query" name:"AndroidOpenType"`
	Title                            string           `position:"Query" name:"Title"`
	SmsDelaySecs                     requests.Integer `position:"Query" name:"SmsDelaySecs"`
	AndroidRenderStyle               requests.Integer `position:"Query" name:"AndroidRenderStyle"`
	IOSExtParameters                 string           `position:"Query" name:"iOSExtParameters"`
	AndroidXiaomiImageUrl            string           `position:"Query" name:"AndroidXiaomiImageUrl"`
	SmsTemplateName                  string           `position:"Query" name:"SmsTemplateName"`
	AndroidPopupBody                 string           `position:"Query" name:"AndroidPopupBody"`
	AndroidBigPictureUrl             string           `position:"Query" name:"AndroidBigPictureUrl"`
	IOSSilentNotification            requests.Boolean `position:"Query" name:"iOSSilentNotification"`
	SendChannels                     string           `position:"Query" name:"SendChannels"`
	Target                           string           `position:"Query" name:"Target"`
	AndroidBigTitle                  string           `position:"Query" name:"AndroidBigTitle"`
	AndroidNotificationChannel       string           `position:"Query" name:"AndroidNotificationChannel"`
	AndroidRemind                    requests.Boolean `position:"Query" name:"AndroidRemind"`
	AndroidActivity                  string           `position:"Query" name:"AndroidActivity"`
	SmsSignName                      string           `position:"Query" name:"SmsSignName"`
	AndroidNotificationNotifyId      requests.Integer `position:"Query" name:"AndroidNotificationNotifyId"`
	AppKey                           requests.Integer `position:"Query" name:"AppKey"`
	TargetValue                      string           `position:"Query" name:"TargetValue"`
	AndroidXiaoMiActivity            string           `position:"Query" name:"AndroidXiaoMiActivity"`
	AndroidXiaoMiNotifyTitle         string           `position:"Query" name:"AndroidXiaoMiNotifyTitle"`
}

// PushResponse is the response struct for api Push
type PushResponse struct {
	*responses.BaseResponse
	MessageId string `json:"MessageId" xml:"MessageId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePushRequest creates a request to invoke Push API
func CreatePushRequest() (request *PushRequest) {
	request = &PushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "Push", "", "")
	request.Method = requests.POST
	return
}

// CreatePushResponse creates a response to parse from Push response
func CreatePushResponse() (response *PushResponse) {
	response = &PushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
