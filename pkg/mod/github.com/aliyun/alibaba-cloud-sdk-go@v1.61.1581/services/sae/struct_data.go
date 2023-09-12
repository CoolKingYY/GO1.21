package sae

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

// Data is a nested struct in sae response
type Data struct {
	LastDisableTime               int64                    `json:"LastDisableTime" xml:"LastDisableTime"`
	VpcName                       string                   `json:"VpcName" xml:"VpcName"`
	TomcatConfig                  string                   `json:"TomcatConfig" xml:"TomcatConfig"`
	Logo                          string                   `json:"Logo" xml:"Logo"`
	VSwitchName                   string                   `json:"VSwitchName" xml:"VSwitchName"`
	LastChangeOrderId             string                   `json:"LastChangeOrderId" xml:"LastChangeOrderId"`
	EdasContainerVersion          string                   `json:"EdasContainerVersion" xml:"EdasContainerVersion"`
	SlsConfigs                    string                   `json:"SlsConfigs" xml:"SlsConfigs"`
	NasId                         string                   `json:"NasId" xml:"NasId"`
	PreStop                       string                   `json:"PreStop" xml:"PreStop"`
	ListenerPort                  int                      `json:"ListenerPort" xml:"ListenerPort"`
	PipelineName                  string                   `json:"PipelineName" xml:"PipelineName"`
	MseApplicationId              string                   `json:"MseApplicationId" xml:"MseApplicationId"`
	NamespaceDescription          string                   `json:"NamespaceDescription" xml:"NamespaceDescription"`
	UpdateTime                    int64                    `json:"UpdateTime" xml:"UpdateTime"`
	JarStartArgs                  string                   `json:"JarStartArgs" xml:"JarStartArgs"`
	Id                            int64                    `json:"Id" xml:"Id"`
	GreyTagRouteId                int64                    `json:"GreyTagRouteId" xml:"GreyTagRouteId"`
	Command                       string                   `json:"Command" xml:"Command"`
	InternetSlbId                 string                   `json:"InternetSlbId" xml:"InternetSlbId"`
	JarStartOptions               string                   `json:"JarStartOptions" xml:"JarStartOptions"`
	PhpExtensions                 string                   `json:"PhpExtensions" xml:"PhpExtensions"`
	PhpConfig                     string                   `json:"PhpConfig" xml:"PhpConfig"`
	Envs                          string                   `json:"Envs" xml:"Envs"`
	RepoType                      string                   `json:"RepoType" xml:"RepoType"`
	RepoTag                       string                   `json:"RepoTag" xml:"RepoTag"`
	SlbId                         string                   `json:"SlbId" xml:"SlbId"`
	TotalSize                     int                      `json:"TotalSize" xml:"TotalSize"`
	NotificationExpired           bool                     `json:"NotificationExpired" xml:"NotificationExpired"`
	UpdateStrategy                string                   `json:"UpdateStrategy" xml:"UpdateStrategy"`
	Name                          string                   `json:"Name" xml:"Name"`
	IntranetSlbId                 string                   `json:"IntranetSlbId" xml:"IntranetSlbId"`
	ChangeOrderId                 string                   `json:"ChangeOrderId" xml:"ChangeOrderId"`
	SecurityGroupId               string                   `json:"SecurityGroupId" xml:"SecurityGroupId"`
	Group                         string                   `json:"Group" xml:"Group"`
	Timezone                      string                   `json:"Timezone" xml:"Timezone"`
	OssAkId                       string                   `json:"OssAkId" xml:"OssAkId"`
	PackageVersion                string                   `json:"PackageVersion" xml:"PackageVersion"`
	PageSize                      int                      `json:"PageSize" xml:"PageSize"`
	MinReadyInstanceRatio         int                      `json:"MinReadyInstanceRatio" xml:"MinReadyInstanceRatio"`
	EnableGreyTagRoute            bool                     `json:"EnableGreyTagRoute" xml:"EnableGreyTagRoute"`
	Replicas                      int                      `json:"Replicas" xml:"Replicas"`
	PhpConfigLocation             string                   `json:"PhpConfigLocation" xml:"PhpConfigLocation"`
	CoStatus                      string                   `json:"CoStatus" xml:"CoStatus"`
	CustomHostAlias               string                   `json:"CustomHostAlias" xml:"CustomHostAlias"`
	TenantId                      string                   `json:"TenantId" xml:"TenantId"`
	NamespaceName                 string                   `json:"NamespaceName" xml:"NamespaceName"`
	NextPipelineId                string                   `json:"NextPipelineId" xml:"NextPipelineId"`
	Memory                        int                      `json:"Memory" xml:"Memory"`
	VSwitchId                     string                   `json:"VSwitchId" xml:"VSwitchId"`
	PhpArmsConfigLocation         string                   `json:"PhpArmsConfigLocation" xml:"PhpArmsConfigLocation"`
	Version                       string                   `json:"Version" xml:"Version"`
	PostStart                     string                   `json:"PostStart" xml:"PostStart"`
	RepoNamespace                 string                   `json:"RepoNamespace" xml:"RepoNamespace"`
	Liveness                      string                   `json:"Liveness" xml:"Liveness"`
	RepoName                      string                   `json:"RepoName" xml:"RepoName"`
	EnableAhas                    string                   `json:"EnableAhas" xml:"EnableAhas"`
	CurrentPage                   int                      `json:"CurrentPage" xml:"CurrentPage"`
	AcrInstanceId                 string                   `json:"AcrInstanceId" xml:"AcrInstanceId"`
	NextToken                     string                   `json:"NextToken" xml:"NextToken"`
	SpringApplicationName         string                   `json:"SpringApplicationName" xml:"SpringApplicationName"`
	AppName                       string                   `json:"AppName" xml:"AppName"`
	PackageUrl                    string                   `json:"PackageUrl" xml:"PackageUrl"`
	DubboApplicationName          string                   `json:"DubboApplicationName" xml:"DubboApplicationName"`
	PackageType                   string                   `json:"PackageType" xml:"PackageType"`
	LastChangeOrderRunning        bool                     `json:"LastChangeOrderRunning" xml:"LastChangeOrderRunning"`
	KafkaConfigs                  string                   `json:"KafkaConfigs" xml:"KafkaConfigs"`
	MountHost                     string                   `json:"MountHost" xml:"MountHost"`
	SlbType                       string                   `json:"SlbType" xml:"SlbType"`
	EdasAppName                   string                   `json:"EdasAppName" xml:"EdasAppName"`
	CurrentPoint                  int                      `json:"CurrentPoint" xml:"CurrentPoint"`
	ImageUrl                      string                   `json:"ImageUrl" xml:"ImageUrl"`
	OssAkSecret                   string                   `json:"OssAkSecret" xml:"OssAkSecret"`
	LastChangeOrderStatus         string                   `json:"LastChangeOrderStatus" xml:"LastChangeOrderStatus"`
	VpcId                         string                   `json:"VpcId" xml:"VpcId"`
	IntranetIp                    string                   `json:"IntranetIp" xml:"IntranetIp"`
	ServiceType                   string                   `json:"ServiceType" xml:"ServiceType"`
	ServiceName                   string                   `json:"ServiceName" xml:"ServiceName"`
	CertId                        string                   `json:"CertId" xml:"CertId"`
	ShowBatch                     bool                     `json:"ShowBatch" xml:"ShowBatch"`
	BatchWaitTime                 int                      `json:"BatchWaitTime" xml:"BatchWaitTime"`
	RepoId                        int                      `json:"RepoId" xml:"RepoId"`
	BelongRegion                  string                   `json:"BelongRegion" xml:"BelongRegion"`
	CrUrl                         string                   `json:"CrUrl" xml:"CrUrl"`
	Jdk                           string                   `json:"Jdk" xml:"Jdk"`
	AppId                         string                   `json:"AppId" xml:"AppId"`
	Metadata                      map[string]interface{}   `json:"Metadata" xml:"Metadata"`
	Php                           string                   `json:"Php" xml:"Php"`
	ScaleRuleName                 string                   `json:"ScaleRuleName" xml:"ScaleRuleName"`
	AcrAssumeRoleArn              string                   `json:"AcrAssumeRoleArn" xml:"AcrAssumeRoleArn"`
	Cpu                           int                      `json:"Cpu" xml:"Cpu"`
	TerminationGracePeriodSeconds int                      `json:"TerminationGracePeriodSeconds" xml:"TerminationGracePeriodSeconds"`
	WebContainer                  string                   `json:"WebContainer" xml:"WebContainer"`
	AssociateEip                  bool                     `json:"AssociateEip" xml:"AssociateEip"`
	RepoOriginType                string                   `json:"RepoOriginType" xml:"RepoOriginType"`
	PipelineStatus                int                      `json:"PipelineStatus" xml:"PipelineStatus"`
	CreateTime                    int64                    `json:"CreateTime" xml:"CreateTime"`
	MinReadyInstances             int                      `json:"MinReadyInstances" xml:"MinReadyInstances"`
	CommandArgs                   string                   `json:"CommandArgs" xml:"CommandArgs"`
	CurrentStageId                string                   `json:"CurrentStageId" xml:"CurrentStageId"`
	AppDescription                string                   `json:"AppDescription" xml:"AppDescription"`
	PhpPECLExtensions             string                   `json:"PhpPECLExtensions" xml:"PhpPECLExtensions"`
	RegionId                      string                   `json:"RegionId" xml:"RegionId"`
	PipelineId                    string                   `json:"PipelineId" xml:"PipelineId"`
	WarStartOptions               string                   `json:"WarStartOptions" xml:"WarStartOptions"`
	UserId                        string                   `json:"UserId" xml:"UserId"`
	Workload                      string                   `json:"Workload" xml:"Workload"`
	Description                   string                   `json:"Description" xml:"Description"`
	ScaleRuleEnabled              bool                     `json:"ScaleRuleEnabled" xml:"ScaleRuleEnabled"`
	InternetIp                    string                   `json:"InternetIp" xml:"InternetIp"`
	AppCount                      int64                    `json:"AppCount" xml:"AppCount"`
	ConfigMapId                   int64                    `json:"ConfigMapId" xml:"ConfigMapId"`
	IsNeedApproval                bool                     `json:"IsNeedApproval" xml:"IsNeedApproval"`
	IngressId                     int64                    `json:"IngressId" xml:"IngressId"`
	Readiness                     string                   `json:"Readiness" xml:"Readiness"`
	ScaleRuleType                 string                   `json:"ScaleRuleType" xml:"ScaleRuleType"`
	NamespaceId                   string                   `json:"NamespaceId" xml:"NamespaceId"`
	Data                          map[string]interface{}   `json:"Data" xml:"Data"`
	Timer                         Timer                    `json:"Timer" xml:"Timer"`
	Order                         Order                    `json:"Order" xml:"Order"`
	DefaultRule                   DefaultRule              `json:"DefaultRule" xml:"DefaultRule"`
	RealTimeRes                   RealTimeRes              `json:"RealTimeRes" xml:"RealTimeRes"`
	Metric                        Metric                   `json:"Metric" xml:"Metric"`
	BagUsage                      BagUsage                 `json:"BagUsage" xml:"BagUsage"`
	Summary                       Summary                  `json:"Summary" xml:"Summary"`
	OssMountDescs                 []OssMountDesc           `json:"OssMountDescs" xml:"OssMountDescs"`
	MountDesc                     []MountDescItem          `json:"MountDesc" xml:"MountDesc"`
	Instances                     []Instance               `json:"Instances" xml:"Instances"`
	Intranet                      []IntranetItem           `json:"Intranet" xml:"Intranet"`
	RelateApps                    []RelateApp              `json:"RelateApps" xml:"RelateApps"`
	TagResources                  []TagResource            `json:"TagResources" xml:"TagResources"`
	Internet                      []InternetItem           `json:"Internet" xml:"Internet"`
	ChangeOrderList               []ChangeOrder            `json:"ChangeOrderList" xml:"ChangeOrderList"`
	Rules                         []Rule                   `json:"Rules" xml:"Rules"`
	LogConfigs                    []LogConfig              `json:"LogConfigs" xml:"LogConfigs"`
	ConfigMapMountDesc            []ConfigMapMountDescItem `json:"ConfigMapMountDesc" xml:"ConfigMapMountDesc"`
	ConfigMaps                    []ConfigMap              `json:"ConfigMaps" xml:"ConfigMaps"`
	DubboRules                    []DubboRule              `json:"DubboRules" xml:"DubboRules"`
	Applications                  []Application            `json:"Applications" xml:"Applications"`
	Tags                          []Tag                    `json:"Tags" xml:"Tags"`
	Methods                       []Method                 `json:"Methods" xml:"Methods"`
	AppEventEntity                []AppEventEntityItem     `json:"AppEventEntity" xml:"AppEventEntity"`
	ApplicationScalingRules       []ApplicationScalingRule `json:"ApplicationScalingRules" xml:"ApplicationScalingRules"`
	ScRules                       []ScRule                 `json:"ScRules" xml:"ScRules"`
	StageList                     []Stage                  `json:"StageList" xml:"StageList"`
	IngressList                   []Ingress                `json:"IngressList" xml:"IngressList"`
	Namespaces                    []Namespace              `json:"Namespaces" xml:"Namespaces"`
}
