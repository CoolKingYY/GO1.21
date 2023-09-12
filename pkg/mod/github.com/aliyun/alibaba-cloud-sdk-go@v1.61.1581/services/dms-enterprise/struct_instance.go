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

// Instance is a nested struct in dms_enterprise response
type Instance struct {
	DatabasePassword string                     `json:"DatabasePassword" xml:"DatabasePassword"`
	DbaNickName      string                     `json:"DbaNickName" xml:"DbaNickName"`
	InstanceType     string                     `json:"InstanceType" xml:"InstanceType"`
	DbaId            string                     `json:"DbaId" xml:"DbaId"`
	SafeRuleId       string                     `json:"SafeRuleId" xml:"SafeRuleId"`
	InstanceId       string                     `json:"InstanceId" xml:"InstanceId"`
	VpcId            string                     `json:"VpcId" xml:"VpcId"`
	DatabaseUser     string                     `json:"DatabaseUser" xml:"DatabaseUser"`
	DataLinkName     string                     `json:"DataLinkName" xml:"DataLinkName"`
	Port             int                        `json:"Port" xml:"Port"`
	DdlOnline        int                        `json:"DdlOnline" xml:"DdlOnline"`
	Sid              string                     `json:"Sid" xml:"Sid"`
	QueryTimeout     int                        `json:"QueryTimeout" xml:"QueryTimeout"`
	EcsInstanceId    string                     `json:"EcsInstanceId" xml:"EcsInstanceId"`
	InstanceSource   string                     `json:"InstanceSource" xml:"InstanceSource"`
	Host             string                     `json:"Host" xml:"Host"`
	State            string                     `json:"State" xml:"State"`
	EnvType          string                     `json:"EnvType" xml:"EnvType"`
	UseDsql          int                        `json:"UseDsql" xml:"UseDsql"`
	ExportTimeout    int                        `json:"ExportTimeout" xml:"ExportTimeout"`
	InstanceAlias    string                     `json:"InstanceAlias" xml:"InstanceAlias"`
	EcsRegion        string                     `json:"EcsRegion" xml:"EcsRegion"`
	OwnerIdList      OwnerIdListInGetInstance   `json:"OwnerIdList" xml:"OwnerIdList"`
	OwnerNameList    OwnerNameListInGetInstance `json:"OwnerNameList" xml:"OwnerNameList"`
	StandardGroup    StandardGroup              `json:"StandardGroup" xml:"StandardGroup"`
}