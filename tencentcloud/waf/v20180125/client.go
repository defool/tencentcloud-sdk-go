// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180125

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-01-25"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddCustomRuleRequest() (request *AddCustomRuleRequest) {
    request = &AddCustomRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "AddCustomRule")
    return
}

func NewAddCustomRuleResponse() (response *AddCustomRuleResponse) {
    response = &AddCustomRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加自定义策略
func (c *Client) AddCustomRule(request *AddCustomRuleRequest) (response *AddCustomRuleResponse, err error) {
    if request == nil {
        request = NewAddCustomRuleRequest()
    }
    response = NewAddCustomRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAttackDownloadTaskRequest() (request *CreateAttackDownloadTaskRequest) {
    request = &CreateAttackDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "CreateAttackDownloadTask")
    return
}

func NewCreateAttackDownloadTaskResponse() (response *CreateAttackDownloadTaskResponse) {
    response = &CreateAttackDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建攻击日志下载任务
func (c *Client) CreateAttackDownloadTask(request *CreateAttackDownloadTaskRequest) (response *CreateAttackDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateAttackDownloadTaskRequest()
    }
    response = NewCreateAttackDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttackDownloadRecordRequest() (request *DeleteAttackDownloadRecordRequest) {
    request = &DeleteAttackDownloadRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteAttackDownloadRecord")
    return
}

func NewDeleteAttackDownloadRecordResponse() (response *DeleteAttackDownloadRecordResponse) {
    response = &DeleteAttackDownloadRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除攻击日志下载任务记录
func (c *Client) DeleteAttackDownloadRecord(request *DeleteAttackDownloadRecordRequest) (response *DeleteAttackDownloadRecordResponse, err error) {
    if request == nil {
        request = NewDeleteAttackDownloadRecordRequest()
    }
    response = NewDeleteAttackDownloadRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDownloadRecordRequest() (request *DeleteDownloadRecordRequest) {
    request = &DeleteDownloadRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteDownloadRecord")
    return
}

func NewDeleteDownloadRecordResponse() (response *DeleteDownloadRecordResponse) {
    response = &DeleteDownloadRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除访问日志下载记录
func (c *Client) DeleteDownloadRecord(request *DeleteDownloadRecordRequest) (response *DeleteDownloadRecordResponse, err error) {
    if request == nil {
        request = NewDeleteDownloadRecordRequest()
    }
    response = NewDeleteDownloadRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSessionRequest() (request *DeleteSessionRequest) {
    request = &DeleteSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DeleteSession")
    return
}

func NewDeleteSessionResponse() (response *DeleteSessionResponse) {
    response = &DeleteSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除CC攻击的session设置
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    if request == nil {
        request = NewDeleteSessionRequest()
    }
    response = NewDeleteSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomRulesRequest() (request *DescribeCustomRulesRequest) {
    request = &DescribeCustomRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeCustomRules")
    return
}

func NewDescribeCustomRulesResponse() (response *DescribeCustomRulesResponse) {
    response = &DescribeCustomRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取防护配置中的自定义策略列表
func (c *Client) DescribeCustomRules(request *DescribeCustomRulesRequest) (response *DescribeCustomRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomRulesRequest()
    }
    response = NewDescribeCustomRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserClbWafRegionsRequest() (request *DescribeUserClbWafRegionsRequest) {
    request = &DescribeUserClbWafRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "DescribeUserClbWafRegions")
    return
}

func NewDescribeUserClbWafRegionsResponse() (response *DescribeUserClbWafRegionsResponse) {
    response = &DescribeUserClbWafRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在负载均衡型WAF的添加、编辑域名配置的时候，需要展示负载均衡型WAF（clb-waf)支持的地域列表，通过DescribeUserClbWafRegions既可以获得当前对客户已经开放的地域列表
func (c *Client) DescribeUserClbWafRegions(request *DescribeUserClbWafRegionsRequest) (response *DescribeUserClbWafRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeUserClbWafRegionsRequest()
    }
    response = NewDescribeUserClbWafRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomRuleStatusRequest() (request *ModifyCustomRuleStatusRequest) {
    request = &ModifyCustomRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("waf", APIVersion, "ModifyCustomRuleStatus")
    return
}

func NewModifyCustomRuleStatusResponse() (response *ModifyCustomRuleStatusResponse) {
    response = &ModifyCustomRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开启或禁用自定义策略
func (c *Client) ModifyCustomRuleStatus(request *ModifyCustomRuleStatusRequest) (response *ModifyCustomRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyCustomRuleStatusRequest()
    }
    response = NewModifyCustomRuleStatusResponse()
    err = c.Send(request, response)
    return
}
