package dms_enterprise

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

// User is a nested struct in dms_enterprise response
type User struct {
	Uid              string                `json:"Uid" xml:"Uid"`
	Email            string                `json:"Email" xml:"Email"`
	CurResultCount   int64                 `json:"CurResultCount" xml:"CurResultCount"`
	NotificationMode string                `json:"NotificationMode" xml:"NotificationMode"`
	MaxExecuteCount  int64                 `json:"MaxExecuteCount" xml:"MaxExecuteCount"`
	LastLoginTime    string                `json:"LastLoginTime" xml:"LastLoginTime"`
	State            string                `json:"State" xml:"State"`
	Mobile           string                `json:"Mobile" xml:"Mobile"`
	CurExecuteCount  int64                 `json:"CurExecuteCount" xml:"CurExecuteCount"`
	UserId           string                `json:"UserId" xml:"UserId"`
	ParentUid        int64                 `json:"ParentUid" xml:"ParentUid"`
	SignatureMethod  string                `json:"SignatureMethod" xml:"SignatureMethod"`
	NickName         string                `json:"NickName" xml:"NickName"`
	DingRobot        string                `json:"DingRobot" xml:"DingRobot"`
	MaxResultCount   int64                 `json:"MaxResultCount" xml:"MaxResultCount"`
	Webhook          string                `json:"Webhook" xml:"Webhook"`
	RoleNameList     RoleNameListInGetUser `json:"RoleNameList" xml:"RoleNameList"`
	RoleIdList       RoleIdListInGetUser   `json:"RoleIdList" xml:"RoleIdList"`
}