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

// DataItem is a nested struct in quickbi_public response
type DataItem struct {
	ModifyUserAccountName string     `json:"ModifyUserAccountName" xml:"ModifyUserAccountName"`
	ViewCount             int64      `json:"ViewCount" xml:"ViewCount"`
	Auth3rdFlag           int        `json:"Auth3rdFlag" xml:"Auth3rdFlag"`
	ModifiedTime          string     `json:"ModifiedTime" xml:"ModifiedTime"`
	GmtModify             string     `json:"GmtModify" xml:"GmtModify"`
	WorkType              string     `json:"WorkType" xml:"WorkType"`
	Favorite              bool       `json:"Favorite" xml:"Favorite"`
	OrganizationId        string     `json:"OrganizationId" xml:"OrganizationId"`
	Type                  string     `json:"Type" xml:"Type"`
	PublishStatus         int        `json:"PublishStatus" xml:"PublishStatus"`
	OwnerNum              string     `json:"OwnerNum" xml:"OwnerNum"`
	Email                 string     `json:"Email" xml:"Email"`
	OwnerName             string     `json:"OwnerName" xml:"OwnerName"`
	ModifyUser            string     `json:"ModifyUser" xml:"ModifyUser"`
	WorkspaceName         string     `json:"WorkspaceName" xml:"WorkspaceName"`
	RowLevel              bool       `json:"RowLevel" xml:"RowLevel"`
	WorksId               string     `json:"WorksId" xml:"WorksId"`
	AllowPublishOperation bool       `json:"AllowPublishOperation" xml:"AllowPublishOperation"`
	AccountId             string     `json:"AccountId" xml:"AccountId"`
	SecurityLevel         string     `json:"SecurityLevel" xml:"SecurityLevel"`
	WorkspaceId           string     `json:"WorkspaceId" xml:"WorkspaceId"`
	WorkspaceDescription  string     `json:"WorkspaceDescription" xml:"WorkspaceDescription"`
	Name                  string     `json:"Name" xml:"Name"`
	ModifyTime            string     `json:"ModifyTime" xml:"ModifyTime"`
	DatasetId             string     `json:"DatasetId" xml:"DatasetId"`
	CreateTime            string     `json:"CreateTime" xml:"CreateTime"`
	AuthAdminUser         bool       `json:"AuthAdminUser" xml:"AuthAdminUser"`
	TreeId                string     `json:"TreeId" xml:"TreeId"`
	GmtCreate             string     `json:"GmtCreate" xml:"GmtCreate"`
	WorkName              string     `json:"WorkName" xml:"WorkName"`
	UserId                string     `json:"UserId" xml:"UserId"`
	UserType              int        `json:"UserType" xml:"UserType"`
	NickName              string     `json:"NickName" xml:"NickName"`
	AllowShareOperation   bool       `json:"AllowShareOperation" xml:"AllowShareOperation"`
	CreateUser            string     `json:"CreateUser" xml:"CreateUser"`
	OwnerAccountName      string     `json:"OwnerAccountName" xml:"OwnerAccountName"`
	HasViewAuth           bool       `json:"HasViewAuth" xml:"HasViewAuth"`
	Description           string     `json:"Description" xml:"Description"`
	Phone                 string     `json:"Phone" xml:"Phone"`
	CreateUserAccountName string     `json:"CreateUserAccountName" xml:"CreateUserAccountName"`
	Status                int        `json:"Status" xml:"Status"`
	ModifyName            string     `json:"ModifyName" xml:"ModifyName"`
	DatasetName           string     `json:"DatasetName" xml:"DatasetName"`
	GmtModified           string     `json:"GmtModified" xml:"GmtModified"`
	Owner                 string     `json:"Owner" xml:"Owner"`
	AdminUser             bool       `json:"AdminUser" xml:"AdminUser"`
	LatestViewTime        string     `json:"LatestViewTime" xml:"LatestViewTime"`
	AccountName           string     `json:"AccountName" xml:"AccountName"`
	OwnerId               string     `json:"OwnerId" xml:"OwnerId"`
	HasEditAuth           bool       `json:"HasEditAuth" xml:"HasEditAuth"`
	Directory             Directory  `json:"Directory" xml:"Directory"`
	Role                  Role       `json:"Role" xml:"Role"`
	DataSource            DataSource `json:"DataSource" xml:"DataSource"`
}