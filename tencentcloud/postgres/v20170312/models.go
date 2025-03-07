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

package v20170312

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountInfo struct {

	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 帐号
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 帐号备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 帐号状态。 1-创建中，2-正常，3-修改中，4-密码重置中，-1-删除中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 帐号创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 帐号最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AddDBInstanceToReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *AddDBInstanceToReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("AddDBInstanceToReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddDBInstanceToReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDBInstanceToReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否关闭Ipv6外网，1：是，0：否
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return errors.New("CloseDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 实例唯一标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

func (r *CloseServerlessDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return errors.New("CloseServerlessDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL内核版本，目前支持以下版本：9.3.5、9.5.4、10.4、11.8、12.4 。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-100
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 续费标记：0-正常续费（默认）；1-自动续费；
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名(后续支持)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 实例需要绑定的Tag信息，默认为空
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecCode")
	delete(f, "DBVersion")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AutoRenewFlag")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号列表。每个实例对应一个订单号。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 冻结流水号
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 创建成功的实例ID集合，只在后付费情景下有返回值
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest

	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL内核版本，目前支持：9.3.5、9.5.4、10.4、11.8、12.4五种版本。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-10。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 实例根账号用户名。
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已配置的私有网络中的子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 续费标记：0-正常续费（默认）；1-自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 活动ID。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否。
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 实例需要绑定的Tag信息，默认为空。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`

	// 安全组ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecCode")
	delete(f, "DBVersion")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "Charset")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "ProjectId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AutoRenewFlag")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号列表。每个实例对应一个订单号。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 冻结流水号。
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 创建成功的实例ID集合，只在后付费情景下有返回值。
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL内核版本，目前强制和主实例保持一致
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 一次性购买的实例数量。取值1-100
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值，按量计费模式下该参数传1。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 只读实例的主实例ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例计费类型。目前支持：PREPAID（预付费，即包年包月），POSTPAID_BY_HOUR（后付费，即按量计费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 续费标记：0-正常续费（默认）；1-自动续费；
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 优惠活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 实例名(后续支持)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要支持Ipv6，1：是，0：否
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitempty" name:"NeedSupportIpv6"`

	// 只读组ID。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 实例需要绑定的Tag信息，默认为空
	TagList *Tag `json:"TagList,omitempty" name:"TagList"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *CreateReadOnlyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecCode")
	delete(f, "DBVersion")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "MasterDBInstanceId")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "AutoRenewFlag")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "ReadOnlyGroupId")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateReadOnlyDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单号列表。每个实例对应一个订单号
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// 冻结流水号
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 创建成功的实例ID集合，只在后付费情景下有返回值
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReadOnlyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// 主实例ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 只读组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 延迟时间大小开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 延迟空间大小开关： 0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值，单位ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 延迟空间大小阈值，单位MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟剔除最小保留实例数
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// 安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`
}

func (r *CreateReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MasterDBInstanceId")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLag")
	delete(f, "MaxReplayLatency")
	delete(f, "MinDelayEliminateReserve")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return errors.New("CreateReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 只读组ID
		ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

		// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 可用区ID。公测阶段仅支持ap-shanghai-2、ap-beijing-1,ap-guangzhou-2.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// DB实例名称，同一个账号下该值必须唯一。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// PostgreSQL内核版本，目前只支持：10.4。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// PostgreSQL数据库字符集，目前支持UTF8。
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例需要绑定的标签数组信息
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`
}

func (r *CreateServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBInstanceName")
	delete(f, "DBVersion")
	delete(f, "DBCharset")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "TagList")
	if len(f) > 0 {
		return errors.New("CreateServerlessDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID，该ID全局唯一，如：postgres-xxxxx
		DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBBackup struct {

	// 备份文件唯一标识
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 文件大小(K)
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 策略（0-实例备份；1-多库备份）
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 类型（0-定时）
	Way *int64 `json:"Way,omitempty" name:"Way"`

	// 备份方式（1-完整）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 状态（1-创建中；2-成功；3-失败）
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// DB列表
	DbList []*string `json:"DbList,omitempty" name:"DbList" list`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`
}

type DBInstance struct {

	// 实例所属地域，如: ap-guangzhou，对应RegionSet的Region字段
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所属可用区， 如：ap-guangzhou-3，对应ZoneSet的Zone字段
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// 实例状态，分别为：applying（申请中）、init(待初始化)、initing(初始化中)、running(运行中)、limited run（受限运行）、isolated（已隔离）、recycling（回收中）、recycled（已回收）、job running（任务执行中）、offline（下线）、migrating（迁移中）、expanding（扩容中）、readonly（只读）、restarting（重启中）
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// 实例分配的内存大小，单位：GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory,omitempty" name:"DBInstanceMemory"`

	// 实例分配的存储空间大小，单位：GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage,omitempty" name:"DBInstanceStorage"`

	// 实例分配的CPU数量，单位：个
	DBInstanceCpu *uint64 `json:"DBInstanceCpu,omitempty" name:"DBInstanceCpu"`

	// 售卖规格ID
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" name:"DBInstanceClass"`

	// 实例类型，类型有：1、primary（主实例）；2、readonly（只读实例）；3、guard（灾备实例）；4、temp（临时实例）
	DBInstanceType *string `json:"DBInstanceType,omitempty" name:"DBInstanceType"`

	// 实例版本，目前只支持standard（双机高可用版, 一主一从）
	DBInstanceVersion *string `json:"DBInstanceVersion,omitempty" name:"DBInstanceVersion"`

	// 实例DB字符集
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// PostgreSQL内核版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例执行最后一次更新的时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 实例隔离时间
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 计费模式，1、prepaid（包年包月,预付费）；2、postpaid（按量计费，后付费）
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 是否自动续费，1：自动续费，0：不自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 实例网络连接信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`

	// 机器类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用户的AppId
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 实例的Uid
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 实例是否支持Ipv6，1：支持，0：不支持
	SupportIpv6 *uint64 `json:"SupportIpv6,omitempty" name:"SupportIpv6"`

	// 实例绑定的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`

	// 主实例信息，仅在实例为只读实例时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 只读实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyInstanceNum *int64 `json:"ReadOnlyInstanceNum,omitempty" name:"ReadOnlyInstanceNum"`

	// 只读实例在只读组中的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInReadonlyGroup *string `json:"StatusInReadonlyGroup,omitempty" name:"StatusInReadonlyGroup"`

	// 下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
}

type DBInstanceNetInfo struct {

	// DNS域名
	Address *string `json:"Address,omitempty" name:"Address"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 连接Port地址
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 网络类型，1、inner（基础网络内网地址）；2、private（私有网络内网地址）；3、public（基础网络或私有网络的外网地址）；
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 网络连接状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeleteReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// 待删除只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *DeleteReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("DeleteReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest

	// DB实例名称，实例名和实例ID必须至少传一个，如果同时存在，将只以实例ID为准。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// DB实例ID，实例名和实例ID必须至少传一个，如果同时存在，将只以实例ID为准。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DeleteServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceName")
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DeleteServerlessDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 分页返回，每页最大返回数目，默认20，取值范围为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，返回第几页的用户数据。页码从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据按照创建时间或者用户名排序。取值只能为createTime或者name。createTime-按照创建时间排序；name-按照用户名排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回结果是升序还是降序。取值只能为desc或者asc。desc-降序；asc-升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeAccountsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次调用接口共返回了多少条数据。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号列表详细信息。
		Details []*AccountInfo `json:"Details,omitempty" name:"Details" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 备份方式（1-全量）。目前只支持全量，取值为1。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份列表分页返回，每页返回数量，默认为 20，最小为1，最大值为 100。（当该参数不传或者传0时按默认值处理）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果中的第几页，从第0页开始。默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDBBackupsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回备份列表中备份文件的个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 备份列表
		BackupList []*DBBackup `json:"BackupList,omitempty" name:"BackupList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBErrlogsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-5bq3wfjd
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-01-01 00:00:00，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 搜索关键字
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys" list`

	// 分页返回，每页返回的最大数量。取值为1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，返回第几页的数据，从第0页开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBErrlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "SearchKeys")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDBErrlogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBErrlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次调用返回了多少条数据
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 错误日志列表
		Details []*ErrLogDetail `json:"Details,omitempty" name:"Details" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBErrlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDBInstanceAttributeRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例详细信息。
		DBInstance *DBInstance `json:"DBInstance,omitempty" name:"DBInstance"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个过滤条件进行查询，目前支持的过滤条件有：
	// db-instance-id：按照实例ID过滤，类型为string
	// db-instance-name：按照实例名过滤，类型为string
	// db-project-id：按照项目ID过滤，类型为integer
	// db-pay-mode：按照付费模式过滤，类型为string
	// db-tag-key：按照标签键过滤，类型为string
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 每页显示数量，取值范围为1-100，默认为返回10条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序指标，如实例名、创建时间等，支持DBInstanceId,CreateTime,Name,EndTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 页码偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式，包括升序：asc、降序：desc。
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "Offset")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息集合。
		DBInstanceSet []*DBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowlogsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询起始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 按照何种指标排序，取值为sum_calls或者sum_cost_time。sum_calls-总调用次数；sum_cost_time-总的花费时间
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序规则。desc-降序；asc-升序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 分页返回结果，每页最大返回数量，取值为1-100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回结果，返回结果的第几页，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("DescribeDBSlowlogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次返回多少条数据
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 慢查询日志详情
		Detail *SlowlogDetail `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBXlogsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页返回，表示返回第几页的条目。从第0页开始计数。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回，表示每页有多少条目。取值为1-100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBXlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return errors.New("DescribeDBXlogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBXlogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表示此次返回结果有多少条数据。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Xlog列表
		XlogList []*Xlog `json:"XlogList,omitempty" name:"XlogList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBXlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DescribeDatabasesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据库信息
		Items []*string `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest

	// 订单名集合
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return errors.New("DescribeOrdersRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 订单数组
		Deals []*PgDeal `json:"Deals,omitempty" name:"Deals" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest

	// 可用区名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return errors.New("DescribeProductConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 售卖规格列表。
		SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReadOnlyGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，必须传入主实例ID进行过滤，否则返回值将为空，过滤参数为：db-master-instance-id
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询每一页的条数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询的页码，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询排序依据，目前支持:ROGroupId,CreateTime,Name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 查询排序依据类型，目前支持:desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeReadOnlyGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeReadOnlyGroupsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReadOnlyGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 只读组列表
		ReadOnlyGroupList []*ReadOnlyGroup `json:"ReadOnlyGroupList,omitempty" name:"ReadOnlyGroupList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReadOnlyGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeRegionsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 地域信息集合。
		RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 查询条件
	Filter []*Filter `json:"Filter,omitempty" name:"Filter" list`

	// 查询个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序指标，目前支持实例创建时间CreateTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，包括升序、降序
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeServerlessDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return errors.New("DescribeServerlessDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		DBInstanceSet []*ServerlessDBInstance `json:"DBInstanceSet,omitempty" name:"DBInstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerlessDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeZonesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用区信息集合。
		ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 待下线实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *DestroyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("DestroyDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DestroyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 资源ID列表
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// 包年包月实例解隔离时购买时常 以月为单位
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否使用代金券：true-使用,false-不使用，默认不使用
	AutoVoucher *bool `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券id列表
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *DisIsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return errors.New("DisIsolateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisIsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisIsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrLogDetail struct {

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 数据库名字
	Database *string `json:"Database,omitempty" name:"Database"`

	// 错误发生时间
	ErrTime *string `json:"ErrTime,omitempty" name:"ErrTime"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type Filter struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// 实例根账号用户名。
	AdminName *string `json:"AdminName,omitempty" name:"AdminName"`

	// 实例根账号用户名对应的密码。
	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`

	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset *string `json:"Charset,omitempty" name:"Charset"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "Charset")
	if len(f) > 0 {
		return errors.New("InitDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID集合。
		DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格ID。该参数可以通过调用DescribeProductConfig接口的返回值中的SpecCode字段来获取。
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 存储容量大小，单位：GB。
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Pid")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return errors.New("InquiryPriceCreateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原始价格，单位：分
		OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 折后价格，单位：分
		Price *uint64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 续费周期，按月计算，最大不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *InquiryPriceRenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	if len(f) > 0 {
		return errors.New("InquiryPriceRenewDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总费用，打折前的。比如24650表示246.5元
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际需要付款金额。比如24650表示246.5元
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例的磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例的内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例计费类型，预付费或者后付费。PREPAID-预付费。目前只支持预付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Storage")
	delete(f, "Memory")
	delete(f, "DBInstanceId")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return errors.New("InquiryPriceUpgradeDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总费用，打折前的
		OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 实际需要付款金额
		Price *int64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID集合
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`
}

func (r *IsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	if len(f) > 0 {
		return errors.New("IsolateDBInstancesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户UserName对应的新备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Remark")
	if len(f) > 0 {
		return errors.New("ModifyAccountRemarkRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 数据库实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 新的数据库实例名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return errors.New("ModifyDBInstanceNameRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 当前实例所在只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 实例修改的目标只读组ID
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitempty" name:"NewReadOnlyGroupId"`
}

func (r *ModifyDBInstanceReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "NewReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("ModifyDBInstanceReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest

	// postgresql实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// postgresql实例所属新项目的ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("ModifyDBInstancesProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转移项目成功的实例个数
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReadOnlyGroupConfigRequest struct {
	*tchttp.BaseRequest

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 延迟时间配置开关：0关、1开
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 延迟日志大小配置开关：0关、1开
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟日志大小阈值，单位MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟时间大小阈值，单位ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 自动负载均衡开关：0关、1开
	Rebalance *uint64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// 延迟剔除最小保留实例数
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`
}

func (r *ModifyReadOnlyGroupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLatency")
	delete(f, "MaxReplayLag")
	delete(f, "Rebalance")
	delete(f, "MinDelayEliminateReserve")
	if len(f) > 0 {
		return errors.New("ModifyReadOnlyGroupConfigRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyReadOnlyGroupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyReadOnlyGroupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NormalQueryItem struct {

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 调用次数
	Calls *int64 `json:"Calls,omitempty" name:"Calls"`

	// 粒度点
	CallsGrids []*int64 `json:"CallsGrids,omitempty" name:"CallsGrids" list`

	// 花费总时间
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`

	// 影响的行数
	Rows *int64 `json:"Rows,omitempty" name:"Rows"`

	// 花费最小时间
	MinCostTime *float64 `json:"MinCostTime,omitempty" name:"MinCostTime"`

	// 花费最大时间
	MaxCostTime *float64 `json:"MaxCostTime,omitempty" name:"MaxCostTime"`

	// 最早一条慢SQL时间
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 最晚一条慢SQL时间
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 读共享内存块数
	SharedReadBlks *int64 `json:"SharedReadBlks,omitempty" name:"SharedReadBlks"`

	// 写共享内存块数
	SharedWriteBlks *int64 `json:"SharedWriteBlks,omitempty" name:"SharedWriteBlks"`

	// 读io总耗时
	ReadCostTime *int64 `json:"ReadCostTime,omitempty" name:"ReadCostTime"`

	// 写io总耗时
	WriteCostTime *int64 `json:"WriteCostTime,omitempty" name:"WriteCostTime"`

	// 数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 脱敏后的慢SQL
	NormalQuery *string `json:"NormalQuery,omitempty" name:"NormalQuery"`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否开通Ipv6外网，1：是，0：否
	IsIpv6 *int64 `json:"IsIpv6,omitempty" name:"IsIpv6"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return errors.New("OpenDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest

	// 实例的唯一标识符
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`
}

func (r *OpenServerlessDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return errors.New("OpenServerlessDBExtranetAccessRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PgDeal struct {

	// 订单名
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 所属用户
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 订单涉及多少个实例
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 付费模式。1-预付费；0-后付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`
}

type ReadOnlyGroup struct {

	// 只读组标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 主实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitempty" name:"MasterDBInstanceId"`

	// 最小保留实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinDelayEliminateReserve *int64 `json:"MinDelayEliminateReserve,omitempty" name:"MinDelayEliminateReserve"`

	// 延迟空间大小阈值
	MaxReplayLatency *int64 `json:"MaxReplayLatency,omitempty" name:"MaxReplayLatency"`

	// 延迟大小开关
	ReplayLatencyEliminate *int64 `json:"ReplayLatencyEliminate,omitempty" name:"ReplayLatencyEliminate"`

	// 延迟时间大小阈值
	MaxReplayLag *float64 `json:"MaxReplayLag,omitempty" name:"MaxReplayLag"`

	// 延迟时间开关
	ReplayLagEliminate *int64 `json:"ReplayLagEliminate,omitempty" name:"ReplayLagEliminate"`

	// 虚拟网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 地域id
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地区id
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例详细信息
	ReadOnlyDBInstanceList []*DBInstance `json:"ReadOnlyDBInstanceList,omitempty" name:"ReadOnlyDBInstanceList" list`

	// 自动负载均衡开关
	Rebalance *int64 `json:"Rebalance,omitempty" name:"Rebalance"`

	// 网络信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`
}

type RebalanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RebalanceReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("RebalanceReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RebalanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebalanceReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 该地域对应的英文名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 该地域对应的中文名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 该地域对应的数字编号
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`

	// 该地域是否支持国际站售卖，0：不支持，1：支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportInternational *uint64 `json:"SupportInternational,omitempty" name:"SupportInternational"`
}

type RemoveDBInstanceFromReadOnlyGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RemoveDBInstanceFromReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return errors.New("RemoveDBInstanceFromReadOnlyGroupRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveDBInstanceFromReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveDBInstanceFromReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 续费多少个月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return errors.New("RenewInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单名
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例账户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// UserName账户对应的新密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return errors.New("ResetAccountPasswordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return errors.New("RestartDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步流程ID
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerlessDBAccount struct {

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBUser *string `json:"DBUser,omitempty" name:"DBUser"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBPassword *string `json:"DBPassword,omitempty" name:"DBPassword"`

	// 连接数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBConnLimit *int64 `json:"DBConnLimit,omitempty" name:"DBConnLimit"`
}

type ServerlessDBInstance struct {

	// 实例id，唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceName *string `json:"DBInstanceName,omitempty" name:"DBInstanceName"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" name:"DBInstanceStatus"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 字符集
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBCharset *string `json:"DBCharset,omitempty" name:"DBCharset"`

	// 数据库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceNetInfo []*ServerlessDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" name:"DBInstanceNetInfo" list`

	// 实例账户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBAccountSet []*ServerlessDBAccount `json:"DBAccountSet,omitempty" name:"DBAccountSet" list`

	// 实例下的db信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBDatabaseList []*string `json:"DBDatabaseList,omitempty" name:"DBDatabaseList" list`

	// 实例绑定的标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList" list`
}

type ServerlessDBInstanceNetInfo struct {

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 实例ID数组
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitempty" name:"DBInstanceIdSet" list`

	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return errors.New("SetAutoRenewFlagRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设置成功的实例个数
		Count *int64 `json:"Count,omitempty" name:"Count"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowlogDetail struct {

	// 花费总时间
	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// 调用总次数
	TotalCalls *int64 `json:"TotalCalls,omitempty" name:"TotalCalls"`

	// 脱敏后的慢SQL列表
	NormalQueries []*NormalQueryItem `json:"NormalQueries,omitempty" name:"NormalQueries" list`
}

type SpecInfo struct {

	// 地域英文编码，对应RegionSet的Region字段
	Region *string `json:"Region,omitempty" name:"Region"`

	// 区域英文编码，对应ZoneSet的Zone字段
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 规格详细信息列表
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList,omitempty" name:"SpecItemInfoList" list`
}

type SpecItemInfo struct {

	// 规格ID
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// PostgreSQL的内核版本编号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 内核编号对应的完整版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// CPU核数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位：MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 该规格所支持最大存储容量，单位：GB
	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 该规格所支持最小存储容量，单位：GB
	MinStorage *uint64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 该规格的预估QPS
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 该规格对应的计费ID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 机器类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Tag struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 升级后的实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例ID，形如postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitempty" name:"DBInstanceId"`

	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 活动ID
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "DBInstanceId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return errors.New("UpgradeDBInstanceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易名字。
		DealName *string `json:"DealName,omitempty" name:"DealName"`

		// 冻结流水号
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Xlog struct {

	// 备份文件唯一标识
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// 备份文件大小
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

type ZoneInfo struct {

	// 该可用区的英文名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 该可用区的中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 该可用区对应的数字编号
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	ZoneState *string `json:"ZoneState,omitempty" name:"ZoneState"`

	// 该可用区是否支持Ipv6
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitempty" name:"ZoneSupportIpv6"`
}
