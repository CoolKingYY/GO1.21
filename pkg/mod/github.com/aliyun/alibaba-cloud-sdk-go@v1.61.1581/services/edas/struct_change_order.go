package edas

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

// ChangeOrder is a nested struct in edas response
type ChangeOrder struct {
	Status                 int    `json:"Status" xml:"Status"`
	FinishTime             string `json:"FinishTime" xml:"FinishTime"`
	CreateTime             string `json:"CreateTime" xml:"CreateTime"`
	UserId                 string `json:"UserId" xml:"UserId"`
	ChangeOrderDescription string `json:"ChangeOrderDescription" xml:"ChangeOrderDescription"`
	Source                 string `json:"Source" xml:"Source"`
	BatchCount             int    `json:"BatchCount" xml:"BatchCount"`
	CreateUserId           string `json:"CreateUserId" xml:"CreateUserId"`
	CoTypeCode             string `json:"CoTypeCode" xml:"CoTypeCode"`
	ChangeOrderId          string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	BatchType              string `json:"BatchType" xml:"BatchType"`
	GroupId                string `json:"GroupId" xml:"GroupId"`
	CoType                 string `json:"CoType" xml:"CoType"`
	AppId                  string `json:"AppId" xml:"AppId"`
}
