package quickbi_public

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

// Data is a nested struct in quickbi_public response
type Data struct {
	Name                 string          `json:"Name" xml:"Name"`
	ModifyTime           string          `json:"ModifyTime" xml:"ModifyTime"`
	ShareId              string          `json:"ShareId" xml:"ShareId"`
	ReportId             string          `json:"ReportId" xml:"ReportId"`
	CreateTime           string          `json:"CreateTime" xml:"CreateTime"`
	WorksName            string          `json:"WorksName" xml:"WorksName"`
	ShowOnlyWithAccess   bool            `json:"ShowOnlyWithAccess" xml:"ShowOnlyWithAccess"`
	IdentifiedPath       string          `json:"IdentifiedPath" xml:"IdentifiedPath"`
	UsergroupDesc        string          `json:"UsergroupDesc" xml:"UsergroupDesc"`
	ParentUserGroupName  string          `json:"ParentUserGroupName" xml:"ParentUserGroupName"`
	UsergroupId          string          `json:"UsergroupId" xml:"UsergroupId"`
	ModifiedTime         string          `json:"ModifiedTime" xml:"ModifiedTime"`
	WorkName             string          `json:"WorkName" xml:"WorkName"`
	TagName              string          `json:"TagName" xml:"TagName"`
	AuthPoint            int             `json:"AuthPoint" xml:"AuthPoint"`
	FavoriteId           int             `json:"FavoriteId" xml:"FavoriteId"`
	UserGroupName        string          `json:"UserGroupName" xml:"UserGroupName"`
	CreateUser           string          `json:"CreateUser" xml:"CreateUser"`
	WorkType             string          `json:"WorkType" xml:"WorkType"`
	TagValue             string          `json:"TagValue" xml:"TagValue"`
	Id                   string          `json:"Id" xml:"Id"`
	Description          string          `json:"Description" xml:"Description"`
	IsUserGroup          bool            `json:"IsUserGroup" xml:"IsUserGroup"`
	ShareToId            string          `json:"ShareToId" xml:"ShareToId"`
	ExpireDate           int64           `json:"ExpireDate" xml:"ExpireDate"`
	WorksType            string          `json:"WorksType" xml:"WorksType"`
	Status               int             `json:"Status" xml:"Status"`
	ParentUserGroupId    string          `json:"ParentUserGroupId" xml:"ParentUserGroupId"`
	ModifyName           string          `json:"ModifyName" xml:"ModifyName"`
	ModifyUser           string          `json:"ModifyUser" xml:"ModifyUser"`
	OwnerName            string          `json:"OwnerName" xml:"OwnerName"`
	UserGroupId          string          `json:"UserGroupId" xml:"UserGroupId"`
	WorkspaceName        string          `json:"WorkspaceName" xml:"WorkspaceName"`
	TagId                string          `json:"TagId" xml:"TagId"`
	ShareToName          string          `json:"ShareToName" xml:"ShareToName"`
	ShareToType          int             `json:"ShareToType" xml:"ShareToType"`
	ParentUsergroupId    string          `json:"ParentUsergroupId" xml:"ParentUsergroupId"`
	WorksId              string          `json:"WorksId" xml:"WorksId"`
	UsergroupName        string          `json:"UsergroupName" xml:"UsergroupName"`
	MenuId               string          `json:"MenuId" xml:"MenuId"`
	ThirdPartAuthFlag    int             `json:"ThirdPartAuthFlag" xml:"ThirdPartAuthFlag"`
	SecurityLevel        string          `json:"SecurityLevel" xml:"SecurityLevel"`
	OwnerId              string          `json:"OwnerId" xml:"OwnerId"`
	UserGroupDescription string          `json:"UserGroupDescription" xml:"UserGroupDescription"`
	WorkspaceId          string          `json:"WorkspaceId" xml:"WorkspaceId"`
	ShareType            string          `json:"ShareType" xml:"ShareType"`
	Directory            Directory       `json:"Directory" xml:"Directory"`
	Receivers            []ReceiversItem `json:"Receivers" xml:"Receivers"`
}