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

// DataInGetRequestDiagnosisResult is a nested struct in das response
type DataInGetRequestDiagnosisResult struct {
	MessageId   string `json:"messageId" xml:"messageId"`
	Uuid        string `json:"uuid" xml:"uuid"`
	AccountId   string `json:"accountId" xml:"accountId"`
	SqlId       string `json:"sqlId" xml:"sqlId"`
	Engine      string `json:"engine" xml:"engine"`
	DbSchema    string `json:"dbSchema" xml:"dbSchema"`
	Param       string `json:"param" xml:"param"`
	State       int    `json:"state" xml:"state"`
	Result      string `json:"result" xml:"result"`
	GmtCreate   string `json:"gmtCreate" xml:"gmtCreate"`
	GmtModified string `json:"gmtModified" xml:"gmtModified"`
}
