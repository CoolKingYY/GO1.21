package cloudcallcenter

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

// TaobaoAliyunAccount is a nested struct in cloudcallcenter response
type TaobaoAliyunAccount struct {
	AliyunPk       int64  `json:"aliyunPk" xml:"aliyunPk"`
	TaobaoUid      int64  `json:"taobaoUid" xml:"taobaoUid"`
	TaobaoNick     string `json:"taobaoNick" xml:"taobaoNick"`
	Status         int    `json:"status" xml:"status"`
	Type           int    `json:"type" xml:"type"`
	StatusModified int64  `json:"statusModified" xml:"statusModified"`
}
