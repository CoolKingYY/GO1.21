package sddp

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

// InstanceSource is a nested struct in sddp response
type InstanceSource struct {
	LastModifyUserId    string `json:"LastModifyUserId" xml:"LastModifyUserId"`
	PasswordStatus      int    `json:"PasswordStatus" xml:"PasswordStatus"`
	EngineType          string `json:"EngineType" xml:"EngineType"`
	TenantName          string `json:"TenantName" xml:"TenantName"`
	InstanceId          string `json:"InstanceId" xml:"InstanceId"`
	InstanceDescription string `json:"InstanceDescription" xml:"InstanceDescription"`
	DataLimitId         int64  `json:"DataLimitId" xml:"DataLimitId"`
	RegionId            string `json:"RegionId" xml:"RegionId"`
	DbName              string `json:"DbName" xml:"DbName"`
	LastModifyTime      int64  `json:"LastModifyTime" xml:"LastModifyTime"`
	RegionName          string `json:"RegionName" xml:"RegionName"`
	CanModifyUserName   bool   `json:"CanModifyUserName" xml:"CanModifyUserName"`
	LogStoreDay         int    `json:"LogStoreDay" xml:"LogStoreDay"`
	GmtCreate           int64  `json:"GmtCreate" xml:"GmtCreate"`
	AutoScan            int    `json:"AutoScan" xml:"AutoScan"`
	ProductId           int64  `json:"ProductId" xml:"ProductId"`
	InstanceSize        int64  `json:"InstanceSize" xml:"InstanceSize"`
	UserName            string `json:"UserName" xml:"UserName"`
	AuditStatus         int    `json:"AuditStatus" xml:"AuditStatus"`
	Id                  int64  `json:"Id" xml:"Id"`
	TenantId            string `json:"TenantId" xml:"TenantId"`
	Enable              int    `json:"Enable" xml:"Enable"`
	CheckStatus         int    `json:"CheckStatus" xml:"CheckStatus"`
	DatamaskStatus      int    `json:"DatamaskStatus" xml:"DatamaskStatus"`
	ErrorMessage        string `json:"ErrorMessage" xml:"ErrorMessage"`
	SamplingSize        int    `json:"SamplingSize" xml:"SamplingSize"`
}
