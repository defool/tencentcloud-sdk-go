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

package v20190823

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-23"

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


func NewClearTablesRequest() (request *ClearTablesRequest) {
    request = &ClearTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ClearTables")
    return
}

func NewClearTablesResponse() (response *ClearTablesResponse) {
    response = &ClearTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据给定的表信息，清除表数据。
func (c *Client) ClearTables(request *ClearTablesRequest) (response *ClearTablesResponse, err error) {
    if request == nil {
        request = NewClearTablesRequest()
    }
    response = NewClearTablesResponse()
    err = c.Send(request, response)
    return
}

func NewCompareIdlFilesRequest() (request *CompareIdlFilesRequest) {
    request = &CompareIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CompareIdlFiles")
    return
}

func NewCompareIdlFilesResponse() (response *CompareIdlFilesResponse) {
    response = &CompareIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 选中目标表格，上传并校验改表文件，返回是否允许修改表格结构的结果。
func (c *Client) CompareIdlFiles(request *CompareIdlFilesRequest) (response *CompareIdlFilesResponse, err error) {
    if request == nil {
        request = NewCompareIdlFilesRequest()
    }
    response = NewCompareIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateBackup")
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户创建备份任务
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateCluster")
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于创建TcaplusDB集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotsRequest() (request *CreateSnapshotsRequest) {
    request = &CreateSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateSnapshots")
    return
}

func NewCreateSnapshotsResponse() (response *CreateSnapshotsResponse) {
    response = &CreateSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 构造表格过去时间点的快照
func (c *Client) CreateSnapshots(request *CreateSnapshotsRequest) (response *CreateSnapshotsResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotsRequest()
    }
    response = NewCreateSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTableGroupRequest() (request *CreateTableGroupRequest) {
    request = &CreateTableGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateTableGroup")
    return
}

func NewCreateTableGroupResponse() (response *CreateTableGroupResponse) {
    response = &CreateTableGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在TcaplusDB集群下创建表格组
func (c *Client) CreateTableGroup(request *CreateTableGroupRequest) (response *CreateTableGroupResponse, err error) {
    if request == nil {
        request = NewCreateTableGroupRequest()
    }
    response = NewCreateTableGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTablesRequest() (request *CreateTablesRequest) {
    request = &CreateTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateTables")
    return
}

func NewCreateTablesResponse() (response *CreateTablesResponse) {
    response = &CreateTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据选择的IDL文件列表，批量创建表格
func (c *Client) CreateTables(request *CreateTablesRequest) (response *CreateTablesResponse, err error) {
    if request == nil {
        request = NewCreateTablesRequest()
    }
    response = NewCreateTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteCluster")
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除TcaplusDB集群，必须在集群所属所有资源（包括表格组，表）都已经释放的情况下才会成功。
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIdlFilesRequest() (request *DeleteIdlFilesRequest) {
    request = &DeleteIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteIdlFiles")
    return
}

func NewDeleteIdlFilesResponse() (response *DeleteIdlFilesResponse) {
    response = &DeleteIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 指定集群ID和待删除IDL文件的信息，删除目标文件，如果文件正在被表关联则删除失败。
func (c *Client) DeleteIdlFiles(request *DeleteIdlFilesRequest) (response *DeleteIdlFilesResponse, err error) {
    if request == nil {
        request = NewDeleteIdlFilesRequest()
    }
    response = NewDeleteIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteSnapshots")
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除表格的快照
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableGroupRequest() (request *DeleteTableGroupRequest) {
    request = &DeleteTableGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTableGroup")
    return
}

func NewDeleteTableGroupResponse() (response *DeleteTableGroupResponse) {
    response = &DeleteTableGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除表格组
func (c *Client) DeleteTableGroup(request *DeleteTableGroupRequest) (response *DeleteTableGroupResponse, err error) {
    if request == nil {
        request = NewDeleteTableGroupRequest()
    }
    response = NewDeleteTableGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableIndexRequest() (request *DeleteTableIndexRequest) {
    request = &DeleteTableIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTableIndex")
    return
}

func NewDeleteTableIndexResponse() (response *DeleteTableIndexResponse) {
    response = &DeleteTableIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除表格的分布式索引
func (c *Client) DeleteTableIndex(request *DeleteTableIndexRequest) (response *DeleteTableIndexResponse, err error) {
    if request == nil {
        request = NewDeleteTableIndexRequest()
    }
    response = NewDeleteTableIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTablesRequest() (request *DeleteTablesRequest) {
    request = &DeleteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTables")
    return
}

func NewDeleteTablesResponse() (response *DeleteTablesResponse) {
    response = &DeleteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除指定的表,第一次调用此接口代表将表移动至回收站，再次调用代表将此表格从回收站中彻底删除。
func (c *Client) DeleteTables(request *DeleteTablesRequest) (response *DeleteTablesResponse, err error) {
    if request == nil {
        request = NewDeleteTablesRequest()
    }
    response = NewDeleteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeApplications")
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取审批管理的申请单
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterTagsRequest() (request *DescribeClusterTagsRequest) {
    request = &DescribeClusterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeClusterTags")
    return
}

func NewDescribeClusterTagsResponse() (response *DescribeClusterTagsResponse) {
    response = &DescribeClusterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取集群关联的标签列表
func (c *Client) DescribeClusterTags(request *DescribeClusterTagsRequest) (response *DescribeClusterTagsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterTagsRequest()
    }
    response = NewDescribeClusterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeClusters")
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TcaplusDB集群列表，包含集群详细信息。
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdlFileInfosRequest() (request *DescribeIdlFileInfosRequest) {
    request = &DescribeIdlFileInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeIdlFileInfos")
    return
}

func NewDescribeIdlFileInfosResponse() (response *DescribeIdlFileInfosResponse) {
    response = &DescribeIdlFileInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询表描述文件详情
func (c *Client) DescribeIdlFileInfos(request *DescribeIdlFileInfosRequest) (response *DescribeIdlFileInfosResponse, err error) {
    if request == nil {
        request = NewDescribeIdlFileInfosRequest()
    }
    response = NewDescribeIdlFileInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineRequest() (request *DescribeMachineRequest) {
    request = &DescribeMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeMachine")
    return
}

func NewDescribeMachineResponse() (response *DescribeMachineResponse) {
    response = &DescribeMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询独占集群可以申请的剩余机器
func (c *Client) DescribeMachine(request *DescribeMachineRequest) (response *DescribeMachineResponse, err error) {
    if request == nil {
        request = NewDescribeMachineRequest()
    }
    response = NewDescribeMachineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TcaplusDB服务支持的地域列表
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeSnapshots")
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询快照列表
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableGroupTagsRequest() (request *DescribeTableGroupTagsRequest) {
    request = &DescribeTableGroupTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableGroupTags")
    return
}

func NewDescribeTableGroupTagsResponse() (response *DescribeTableGroupTagsResponse) {
    response = &DescribeTableGroupTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取表格组关联的标签列表
func (c *Client) DescribeTableGroupTags(request *DescribeTableGroupTagsRequest) (response *DescribeTableGroupTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTableGroupTagsRequest()
    }
    response = NewDescribeTableGroupTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableGroupsRequest() (request *DescribeTableGroupsRequest) {
    request = &DescribeTableGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableGroups")
    return
}

func NewDescribeTableGroupsResponse() (response *DescribeTableGroupsResponse) {
    response = &DescribeTableGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询表格组列表
func (c *Client) DescribeTableGroups(request *DescribeTableGroupsRequest) (response *DescribeTableGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTableGroupsRequest()
    }
    response = NewDescribeTableGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableTagsRequest() (request *DescribeTableTagsRequest) {
    request = &DescribeTableTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableTags")
    return
}

func NewDescribeTableTagsResponse() (response *DescribeTableTagsResponse) {
    response = &DescribeTableTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取表格标签
func (c *Client) DescribeTableTags(request *DescribeTableTagsRequest) (response *DescribeTableTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTableTagsRequest()
    }
    response = NewDescribeTableTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTables")
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询表详情
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesInRecycleRequest() (request *DescribeTablesInRecycleRequest) {
    request = &DescribeTablesInRecycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTablesInRecycle")
    return
}

func NewDescribeTablesInRecycleResponse() (response *DescribeTablesInRecycleResponse) {
    response = &DescribeTablesInRecycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询回收站中的表详情
func (c *Client) DescribeTablesInRecycle(request *DescribeTablesInRecycleRequest) (response *DescribeTablesInRecycleResponse, err error) {
    if request == nil {
        request = NewDescribeTablesInRecycleRequest()
    }
    response = NewDescribeTablesInRecycleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询任务列表
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUinInWhitelistRequest() (request *DescribeUinInWhitelistRequest) {
    request = &DescribeUinInWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeUinInWhitelist")
    return
}

func NewDescribeUinInWhitelistResponse() (response *DescribeUinInWhitelistResponse) {
    response = &DescribeUinInWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询本用户是否在白名单中，控制是否能创建TDR类型的APP或表
func (c *Client) DescribeUinInWhitelist(request *DescribeUinInWhitelistRequest) (response *DescribeUinInWhitelistResponse, err error) {
    if request == nil {
        request = NewDescribeUinInWhitelistRequest()
    }
    response = NewDescribeUinInWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRestProxyRequest() (request *DisableRestProxyRequest) {
    request = &DisableRestProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DisableRestProxy")
    return
}

func NewDisableRestProxyResponse() (response *DisableRestProxyResponse) {
    response = &DisableRestProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 当restful api为关闭状态时，可以通过此接口关闭restful api
func (c *Client) DisableRestProxy(request *DisableRestProxyRequest) (response *DisableRestProxyResponse, err error) {
    if request == nil {
        request = NewDisableRestProxyRequest()
    }
    response = NewDisableRestProxyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRestProxyRequest() (request *EnableRestProxyRequest) {
    request = &EnableRestProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "EnableRestProxy")
    return
}

func NewEnableRestProxyResponse() (response *EnableRestProxyResponse) {
    response = &EnableRestProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 当restful api为关闭状态时，可以通过此接口开启restful apu
func (c *Client) EnableRestProxy(request *EnableRestProxyRequest) (response *EnableRestProxyResponse, err error) {
    if request == nil {
        request = NewEnableRestProxyRequest()
    }
    response = NewEnableRestProxyResponse()
    err = c.Send(request, response)
    return
}

func NewImportSnapshotsRequest() (request *ImportSnapshotsRequest) {
    request = &ImportSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ImportSnapshots")
    return
}

func NewImportSnapshotsResponse() (response *ImportSnapshotsResponse) {
    response = &ImportSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将快照数据导入到新表或当前表
func (c *Client) ImportSnapshots(request *ImportSnapshotsRequest) (response *ImportSnapshotsResponse, err error) {
    if request == nil {
        request = NewImportSnapshotsRequest()
    }
    response = NewImportSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewMergeTablesDataRequest() (request *MergeTablesDataRequest) {
    request = &MergeTablesDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "MergeTablesData")
    return
}

func NewMergeTablesDataResponse() (response *MergeTablesDataResponse) {
    response = &MergeTablesDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 合并指定表格
func (c *Client) MergeTablesData(request *MergeTablesDataRequest) (response *MergeTablesDataResponse, err error) {
    if request == nil {
        request = NewMergeTablesDataRequest()
    }
    response = NewMergeTablesDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCensorshipRequest() (request *ModifyCensorshipRequest) {
    request = &ModifyCensorshipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyCensorship")
    return
}

func NewModifyCensorshipResponse() (response *ModifyCensorshipResponse) {
    response = &ModifyCensorshipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群审批状态
func (c *Client) ModifyCensorship(request *ModifyCensorshipRequest) (response *ModifyCensorshipResponse, err error) {
    if request == nil {
        request = NewModifyCensorshipRequest()
    }
    response = NewModifyCensorshipResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterMachineRequest() (request *ModifyClusterMachineRequest) {
    request = &ModifyClusterMachineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterMachine")
    return
}

func NewModifyClusterMachineResponse() (response *ModifyClusterMachineResponse) {
    response = &ModifyClusterMachineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改独占集群机器
func (c *Client) ModifyClusterMachine(request *ModifyClusterMachineRequest) (response *ModifyClusterMachineResponse, err error) {
    if request == nil {
        request = NewModifyClusterMachineRequest()
    }
    response = NewModifyClusterMachineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
    request = &ModifyClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterName")
    return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
    response = &ModifyClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定的集群名称
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterPasswordRequest() (request *ModifyClusterPasswordRequest) {
    request = &ModifyClusterPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterPassword")
    return
}

func NewModifyClusterPasswordResponse() (response *ModifyClusterPasswordResponse) {
    response = &ModifyClusterPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定集群的密码，后台将在旧密码失效之前同时支持TcaplusDB SDK使用旧密码和新密码访问数据库。在旧密码失效之前不能提交新的密码修改请求，在旧密码失效之后不能提交修改旧密码过期时间的请求。
func (c *Client) ModifyClusterPassword(request *ModifyClusterPasswordRequest) (response *ModifyClusterPasswordResponse, err error) {
    if request == nil {
        request = NewModifyClusterPasswordRequest()
    }
    response = NewModifyClusterPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterTagsRequest() (request *ModifyClusterTagsRequest) {
    request = &ModifyClusterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterTags")
    return
}

func NewModifyClusterTagsResponse() (response *ModifyClusterTagsResponse) {
    response = &ModifyClusterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改集群标签
func (c *Client) ModifyClusterTags(request *ModifyClusterTagsRequest) (response *ModifyClusterTagsResponse, err error) {
    if request == nil {
        request = NewModifyClusterTagsRequest()
    }
    response = NewModifyClusterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotsRequest() (request *ModifySnapshotsRequest) {
    request = &ModifySnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifySnapshots")
    return
}

func NewModifySnapshotsResponse() (response *ModifySnapshotsResponse) {
    response = &ModifySnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改表格快照的过期时间
func (c *Client) ModifySnapshots(request *ModifySnapshotsRequest) (response *ModifySnapshotsResponse, err error) {
    if request == nil {
        request = NewModifySnapshotsRequest()
    }
    response = NewModifySnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableGroupNameRequest() (request *ModifyTableGroupNameRequest) {
    request = &ModifyTableGroupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableGroupName")
    return
}

func NewModifyTableGroupNameResponse() (response *ModifyTableGroupNameResponse) {
    response = &ModifyTableGroupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改TcaplusDB表格组名称
func (c *Client) ModifyTableGroupName(request *ModifyTableGroupNameRequest) (response *ModifyTableGroupNameResponse, err error) {
    if request == nil {
        request = NewModifyTableGroupNameRequest()
    }
    response = NewModifyTableGroupNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableGroupTagsRequest() (request *ModifyTableGroupTagsRequest) {
    request = &ModifyTableGroupTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableGroupTags")
    return
}

func NewModifyTableGroupTagsResponse() (response *ModifyTableGroupTagsResponse) {
    response = &ModifyTableGroupTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改表格组标签
func (c *Client) ModifyTableGroupTags(request *ModifyTableGroupTagsRequest) (response *ModifyTableGroupTagsResponse, err error) {
    if request == nil {
        request = NewModifyTableGroupTagsRequest()
    }
    response = NewModifyTableGroupTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableMemosRequest() (request *ModifyTableMemosRequest) {
    request = &ModifyTableMemosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableMemos")
    return
}

func NewModifyTableMemosResponse() (response *ModifyTableMemosResponse) {
    response = &ModifyTableMemosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改表备注信息
func (c *Client) ModifyTableMemos(request *ModifyTableMemosRequest) (response *ModifyTableMemosResponse, err error) {
    if request == nil {
        request = NewModifyTableMemosRequest()
    }
    response = NewModifyTableMemosResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableQuotasRequest() (request *ModifyTableQuotasRequest) {
    request = &ModifyTableQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableQuotas")
    return
}

func NewModifyTableQuotasResponse() (response *ModifyTableQuotasResponse) {
    response = &ModifyTableQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 表格扩缩容
func (c *Client) ModifyTableQuotas(request *ModifyTableQuotasRequest) (response *ModifyTableQuotasResponse, err error) {
    if request == nil {
        request = NewModifyTableQuotasRequest()
    }
    response = NewModifyTableQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableTagsRequest() (request *ModifyTableTagsRequest) {
    request = &ModifyTableTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableTags")
    return
}

func NewModifyTableTagsResponse() (response *ModifyTableTagsResponse) {
    response = &ModifyTableTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改表格标签
func (c *Client) ModifyTableTags(request *ModifyTableTagsRequest) (response *ModifyTableTagsResponse, err error) {
    if request == nil {
        request = NewModifyTableTagsRequest()
    }
    response = NewModifyTableTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTablesRequest() (request *ModifyTablesRequest) {
    request = &ModifyTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTables")
    return
}

func NewModifyTablesResponse() (response *ModifyTablesResponse) {
    response = &ModifyTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据用户选定的表定义IDL文件，批量修改指定的表
func (c *Client) ModifyTables(request *ModifyTablesRequest) (response *ModifyTablesResponse, err error) {
    if request == nil {
        request = NewModifyTablesRequest()
    }
    response = NewModifyTablesResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverRecycleTablesRequest() (request *RecoverRecycleTablesRequest) {
    request = &RecoverRecycleTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "RecoverRecycleTables")
    return
}

func NewRecoverRecycleTablesResponse() (response *RecoverRecycleTablesResponse) {
    response = &RecoverRecycleTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复回收站中，用户自行删除的表。对欠费待释放的表无效。
func (c *Client) RecoverRecycleTables(request *RecoverRecycleTablesRequest) (response *RecoverRecycleTablesResponse, err error) {
    if request == nil {
        request = NewRecoverRecycleTablesRequest()
    }
    response = NewRecoverRecycleTablesResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackTablesRequest() (request *RollbackTablesRequest) {
    request = &RollbackTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "RollbackTables")
    return
}

func NewRollbackTablesResponse() (response *RollbackTablesResponse) {
    response = &RollbackTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 表格数据回档
func (c *Client) RollbackTables(request *RollbackTablesRequest) (response *RollbackTablesResponse, err error) {
    if request == nil {
        request = NewRollbackTablesRequest()
    }
    response = NewRollbackTablesResponse()
    err = c.Send(request, response)
    return
}

func NewSetTableIndexRequest() (request *SetTableIndexRequest) {
    request = &SetTableIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "SetTableIndex")
    return
}

func NewSetTableIndexResponse() (response *SetTableIndexResponse) {
    response = &SetTableIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置表格分布式索引
func (c *Client) SetTableIndex(request *SetTableIndexRequest) (response *SetTableIndexResponse, err error) {
    if request == nil {
        request = NewSetTableIndexRequest()
    }
    response = NewSetTableIndexResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApplyRequest() (request *UpdateApplyRequest) {
    request = &UpdateApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "UpdateApply")
    return
}

func NewUpdateApplyResponse() (response *UpdateApplyResponse) {
    response = &UpdateApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新申请单状态
func (c *Client) UpdateApply(request *UpdateApplyRequest) (response *UpdateApplyResponse, err error) {
    if request == nil {
        request = NewUpdateApplyRequest()
    }
    response = NewUpdateApplyResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyIdlFilesRequest() (request *VerifyIdlFilesRequest) {
    request = &VerifyIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "VerifyIdlFiles")
    return
}

func NewVerifyIdlFilesResponse() (response *VerifyIdlFilesResponse) {
    response = &VerifyIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传并校验创建表格文件，返回校验合法的表格定义
func (c *Client) VerifyIdlFiles(request *VerifyIdlFilesRequest) (response *VerifyIdlFilesResponse, err error) {
    if request == nil {
        request = NewVerifyIdlFilesRequest()
    }
    response = NewVerifyIdlFilesResponse()
    err = c.Send(request, response)
    return
}
