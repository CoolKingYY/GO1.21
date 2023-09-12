package polardb

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

// Task is a nested struct in polardb response
type Task struct {
	FinishTime         string `json:"FinishTime" xml:"FinishTime"`
	StepsInfo          string `json:"StepsInfo" xml:"StepsInfo"`
	Progress           int    `json:"Progress" xml:"Progress"`
	ExpectedFinishTime string `json:"ExpectedFinishTime" xml:"ExpectedFinishTime"`
	BeginTime          string `json:"BeginTime" xml:"BeginTime"`
	TaskErrorCode      string `json:"TaskErrorCode" xml:"TaskErrorCode"`
	ProgressInfo       string `json:"ProgressInfo" xml:"ProgressInfo"`
	CurrentStepName    string `json:"CurrentStepName" xml:"CurrentStepName"`
	StepProgressInfo   string `json:"StepProgressInfo" xml:"StepProgressInfo"`
	TaskErrorMessage   string `json:"TaskErrorMessage" xml:"TaskErrorMessage"`
	TaskAction         string `json:"TaskAction" xml:"TaskAction"`
	DBName             string `json:"DBName" xml:"DBName"`
	Remain             int    `json:"Remain" xml:"Remain"`
	TaskId             string `json:"TaskId" xml:"TaskId"`
}