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

package v20200910

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Advice struct {

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AspectRatio struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *string `json:"Number,omitempty" name:"Number"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type BlockInfo struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Check struct {

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *Desc `json:"Desc,omitempty" name:"Desc"`

	// 结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *Summary `json:"Summary,omitempty" name:"Summary"`
}

type Desc struct {

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 器官
	// 注意：此字段可能返回 null，表示取不到有效值。
	Organ []*Organ `json:"Organ,omitempty" name:"Organ" list`

	// 结节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tuber []*TuberInfo `json:"Tuber,omitempty" name:"Tuber" list`
}

type DiagCert struct {

	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advice *Advice `json:"Advice,omitempty" name:"Advice"`

	// 诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitempty" name:"Diagnosis" list`
}

type DiagCertItem struct {

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value" list`
}

type DischargeDiagnosis struct {

	// 表格位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIndex *int64 `json:"TableIndex,omitempty" name:"TableIndex"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDiagnosis *BlockInfo `json:"OutDiagnosis,omitempty" name:"OutDiagnosis"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseCode *BlockInfo `json:"DiseaseCode,omitempty" name:"DiseaseCode"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	InStatus *BlockInfo `json:"InStatus,omitempty" name:"InStatus"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutStatus *BlockInfo `json:"OutStatus,omitempty" name:"OutStatus"`
}

type DiseaseMedicalHistory struct {

	// 主病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistory *string `json:"MainDiseaseHistory,omitempty" name:"MainDiseaseHistory"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistory *string `json:"AllergyHistory,omitempty" name:"AllergyHistory"`

	// 传染疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfectHistory *string `json:"InfectHistory,omitempty" name:"InfectHistory"`

	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationHistory *string `json:"OperationHistory,omitempty" name:"OperationHistory"`

	// 输血史
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransfusionHistory *string `json:"TransfusionHistory,omitempty" name:"TransfusionHistory"`
}

type Elastic struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *string `json:"Score,omitempty" name:"Score"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type FamilyMedicalHistory struct {

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeHistory *string `json:"RelativeHistory,omitempty" name:"RelativeHistory"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerHistory *string `json:"RelativeCancerHistory,omitempty" name:"RelativeCancerHistory"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticHistory *string `json:"GeneticHistory,omitempty" name:"GeneticHistory"`
}

type FirstPage struct {

	// 出入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis []*DischargeDiagnosis `json:"DischargeDiagnosis,omitempty" name:"DischargeDiagnosis" list`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosis *BlockInfo `json:"PathologicalDiagnosis,omitempty" name:"PathologicalDiagnosis"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClinicalDiagnosis *BlockInfo `json:"ClinicalDiagnosis,omitempty" name:"ClinicalDiagnosis"`
}

type HandleParam struct {

	// ocr引擎
	OcrEngineType *int64 `json:"OcrEngineType,omitempty" name:"OcrEngineType"`

	// 是否返回分行文本内容
	IsReturnText *bool `json:"IsReturnText,omitempty" name:"IsReturnText"`

	// 顺时针旋转角度
	RotateTheAngle *float64 `json:"RotateTheAngle,omitempty" name:"RotateTheAngle"`

	// 自动适配方向,仅支持优图引擎
	AutoFitDirection *bool `json:"AutoFitDirection,omitempty" name:"AutoFitDirection"`

	// 坐标优化
	AutoOptimizeCoordinate *bool `json:"AutoOptimizeCoordinate,omitempty" name:"AutoOptimizeCoordinate"`

	// 是否开启图片压缩，开启时imageOriginalSize必须正确填写
	IsScale *bool `json:"IsScale,omitempty" name:"IsScale"`

	// 原始图片大小(单位byes),用来判断该图片是否需要压缩
	ImageOriginalSize *uint64 `json:"ImageOriginalSize,omitempty" name:"ImageOriginalSize"`

	// 采用后台默认值(2048Kb)
	ScaleTargetSize *uint64 `json:"ScaleTargetSize,omitempty" name:"ScaleTargetSize"`
}

type HistologyLevel struct {

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type HistologyType struct {

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infiltration *string `json:"Infiltration,omitempty" name:"Infiltration"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type IHCInfo struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// “”
	Value *Value `json:"Value,omitempty" name:"Value"`
}

type ImageInfo struct {

	// 图片id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 图片url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 图片base64编码
	Base64 *string `json:"Base64,omitempty" name:"Base64"`
}

type ImageToClassRequest struct {
	*tchttp.BaseRequest

	// 图片列表
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList" list`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 图片类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ImageToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	if len(f) > 0 {
		return errors.New("ImageToClassRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImageToClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageToObjectRequest struct {
	*tchttp.BaseRequest

	// 图片列表
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList" list`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitempty" name:"HandleParam"`

	// 图片类别
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 是否使用分类引擎
	IsUsedClassify *bool `json:"IsUsedClassify,omitempty" name:"IsUsedClassify"`
}

func (r *ImageToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	if len(f) > 0 {
		return errors.New("ImageToObjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ImageToObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 报告结构化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Template *Template `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Indicator struct {

	// 检验指标项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicators []*IndicatorItem `json:"Indicators,omitempty" name:"Indicators" list`
}

type IndicatorItem struct {

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scode *string `json:"Scode,omitempty" name:"Scode"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sname *string `json:"Sname,omitempty" name:"Sname"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *string `json:"Range,omitempty" name:"Range"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arrow *string `json:"Arrow,omitempty" name:"Arrow"`

	// 是否正常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Normal *bool `json:"Normal,omitempty" name:"Normal"`
}

type Invas struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Lymph struct {

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 转移数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *int64 `json:"TransferNum,omitempty" name:"TransferNum"`
}

type MedDoc struct {

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advice *Advice `json:"Advice,omitempty" name:"Advice"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitempty" name:"Diagnosis" list`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseMedicalHistory *DiseaseMedicalHistory `json:"DiseaseMedicalHistory,omitempty" name:"DiseaseMedicalHistory"`

	// “”
	PersonalMedicalHistory *PersonalMedicalHistory `json:"PersonalMedicalHistory,omitempty" name:"PersonalMedicalHistory"`

	// “”
	ObstericalMedicalHistory *ObstericalMedicalHistory `json:"ObstericalMedicalHistory,omitempty" name:"ObstericalMedicalHistory"`

	// “”
	FamilyMedicalHistory *FamilyMedicalHistory `json:"FamilyMedicalHistory,omitempty" name:"FamilyMedicalHistory"`

	// “”
	MenstrualMedicalHistory *MenstrualMedicalHistory `json:"MenstrualMedicalHistory,omitempty" name:"MenstrualMedicalHistory"`

	// “”
	TreatmentRecord *TreatmentRecord `json:"TreatmentRecord,omitempty" name:"TreatmentRecord"`
}

type MenstrualMedicalHistory struct {

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMenstrualPeriod *string `json:"LastMenstrualPeriod,omitempty" name:"LastMenstrualPeriod"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualFlow *string `json:"MenstrualFlow,omitempty" name:"MenstrualFlow"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenarcheAge *string `json:"MenarcheAge,omitempty" name:"MenarcheAge"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstruationOrNot *string `json:"MenstruationOrNot,omitempty" name:"MenstruationOrNot"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualCycles *string `json:"MenstrualCycles,omitempty" name:"MenstrualCycles"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualPeriod *string `json:"MenstrualPeriod,omitempty" name:"MenstrualPeriod"`
}

type Multiple struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type NormPart struct {

	// 部位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitempty" name:"Part"`

	// 部位方向
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartDirection *string `json:"PartDirection,omitempty" name:"PartDirection"`

	// 组织值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tissue *string `json:"Tissue,omitempty" name:"Tissue"`

	// 组织方向
	// 注意：此字段可能返回 null，表示取不到有效值。
	TissueDirection *string `json:"TissueDirection,omitempty" name:"TissueDirection"`

	// 上级部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Upper *string `json:"Upper,omitempty" name:"Upper"`
}

type NormSize struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number []*string `json:"Number,omitempty" name:"Number" list`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type ObstericalMedicalHistory struct {

	// 婚史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriageHistory *string `json:"MarriageHistory,omitempty" name:"MarriageHistory"`

	// 孕史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FertilityHistory *string `json:"FertilityHistory,omitempty" name:"FertilityHistory"`
}

type Organ struct {

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitempty" name:"Size" list`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envelope *BlockInfo `json:"Envelope,omitempty" name:"Envelope"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edge *BlockInfo `json:"Edge,omitempty" name:"Edge"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEcho *BlockInfo `json:"InnerEcho,omitempty" name:"InnerEcho"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gland *BlockInfo `json:"Gland,omitempty" name:"Gland"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shape *BlockInfo `json:"Shape,omitempty" name:"Shape"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Thickness *BlockInfo `json:"Thickness,omitempty" name:"Thickness"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitempty" name:"ShapeAttr"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDFI *BlockInfo `json:"CDFI,omitempty" name:"CDFI"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDesc *BlockInfo `json:"SymDesc,omitempty" name:"SymDesc"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeStatus *BlockInfo `json:"SizeStatus,omitempty" name:"SizeStatus"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline *BlockInfo `json:"Outline,omitempty" name:"Outline"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Structure *BlockInfo `json:"Structure,omitempty" name:"Structure"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Density *BlockInfo `json:"Density,omitempty" name:"Density"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vas *BlockInfo `json:"Vas,omitempty" name:"Vas"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cysticwall *BlockInfo `json:"Cysticwall,omitempty" name:"Cysticwall"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capsule *BlockInfo `json:"Capsule,omitempty" name:"Capsule"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitempty" name:"IsthmusThicknese"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitempty" name:"InnerEchoDistribution"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`
}

type Part struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormPart *NormPart `json:"NormPart,omitempty" name:"NormPart"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PathologyReport struct {

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerPart *Part `json:"CancerPart,omitempty" name:"CancerPart"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerSize []*Size `json:"CancerSize,omitempty" name:"CancerSize" list`

	// 描述文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescText *string `json:"DescText,omitempty" name:"DescText"`

	// 癌症
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyLevel *HistologyLevel `json:"HistologyLevel,omitempty" name:"HistologyLevel"`

	// 扩散
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyType *HistologyType `json:"HistologyType,omitempty" name:"HistologyType"`

	// 淋巴
	// 注意：此字段可能返回 null，表示取不到有效值。
	IHC []*IHCInfo `json:"IHC,omitempty" name:"IHC" list`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfiltrationDepth *BlockInfo `json:"InfiltrationDepth,omitempty" name:"InfiltrationDepth"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invasive []*Invas `json:"Invasive,omitempty" name:"Invasive" list`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphNodes []*Lymph `json:"LymphNodes,omitempty" name:"LymphNodes" list`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *BlockInfo `json:"PTNM,omitempty" name:"PTNM"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalReportType *BlockInfo `json:"PathologicalReportType,omitempty" name:"PathologicalReportType"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportText *string `json:"ReportText,omitempty" name:"ReportText"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *BlockInfo `json:"SampleType,omitempty" name:"SampleType"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryText *string `json:"SummaryText,omitempty" name:"SummaryText"`
}

type PatientInfo struct {

	// 患者姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 患者性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *string `json:"Sex,omitempty" name:"Sex"`

	// 患者年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Age *string `json:"Age,omitempty" name:"Age"`

	// 患者手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 患者地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 患者身份证
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`
}

type PersonalMedicalHistory struct {

	// 出生史
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitempty" name:"BirthPlace"`

	// 居住史
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivePlace *string `json:"LivePlace,omitempty" name:"LivePlace"`

	// 工作史
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *string `json:"Job,omitempty" name:"Job"`

	// 吸烟史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmokeHistory *string `json:"SmokeHistory,omitempty" name:"SmokeHistory"`

	// 饮酒史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlcoholicHistory *string `json:"AlcoholicHistory,omitempty" name:"AlcoholicHistory"`
}

type ReportInfo struct {

	// 医院名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hospital *string `json:"Hospital,omitempty" name:"Hospital"`

	// 科室名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingTime *string `json:"BillingTime,omitempty" name:"BillingTime"`

	// 报告时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTime *string `json:"ReportTime,omitempty" name:"ReportTime"`

	// 检查时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectTime *string `json:"InspectTime,omitempty" name:"InspectTime"`

	// 检查号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckNum *string `json:"CheckNum,omitempty" name:"CheckNum"`

	// 影像号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageNum *string `json:"ImageNum,omitempty" name:"ImageNum"`

	// 放射号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadiationNum *string `json:"RadiationNum,omitempty" name:"RadiationNum"`

	// 检验号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestNum *string `json:"TestNum,omitempty" name:"TestNum"`

	// 门诊号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientNum *string `json:"OutpatientNum,omitempty" name:"OutpatientNum"`

	// 病理号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologyNum *string `json:"PathologyNum,omitempty" name:"PathologyNum"`

	// 住院号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InHospitalNum *string `json:"InHospitalNum,omitempty" name:"InHospitalNum"`

	// 样本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleNum *string `json:"SampleNum,omitempty" name:"SampleNum"`

	// 标本种类
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 病历号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalRecordNum *string `json:"MedicalRecordNum,omitempty" name:"MedicalRecordNum"`

	// 报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`

	// 超声号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UltraNum *string `json:"UltraNum,omitempty" name:"UltraNum"`
}

type Size struct {

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 标准大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormSize *NormSize `json:"NormSize,omitempty" name:"NormSize"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Summary struct {

	// 症状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symptom []*SymptomInfo `json:"Symptom,omitempty" name:"Symptom" list`

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type SymptomInfo struct {

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *BlockInfo `json:"Grade,omitempty" name:"Grade"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symptom *BlockInfo `json:"Symptom,omitempty" name:"Symptom"`

	// 属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attrs []*BlockInfo `json:"Attrs,omitempty" name:"Attrs" list`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Template struct {

	// 患者信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatientInfo *PatientInfo `json:"PatientInfo,omitempty" name:"PatientInfo"`

	// 报告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportInfo *ReportInfo `json:"ReportInfo,omitempty" name:"ReportInfo"`

	// 检查报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Check *Check `json:"Check,omitempty" name:"Check"`

	// 病理报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pathology *PathologyReport `json:"Pathology,omitempty" name:"Pathology"`

	// 出院报告，入院报告，门诊病历
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDoc *MedDoc `json:"MedDoc,omitempty" name:"MedDoc"`

	// 诊断证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagCert *DiagCert `json:"DiagCert,omitempty" name:"DiagCert"`

	// 病案首页
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstPage *FirstPage `json:"FirstPage,omitempty" name:"FirstPage"`

	// 检验报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicator *Indicator `json:"Indicator,omitempty" name:"Indicator"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`
}

type TextToClassRequest struct {
	*tchttp.BaseRequest

	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *TextToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return errors.New("TextToClassRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TextToClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类结果
		TextTypeList []*TextType `json:"TextTypeList,omitempty" name:"TextTypeList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextToObjectRequest struct {
	*tchttp.BaseRequest

	// 报告文本
	Text *string `json:"Text,omitempty" name:"Text"`

	// 报告类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 是否使用分类引擎
	IsUsedClassify *bool `json:"IsUsedClassify,omitempty" name:"IsUsedClassify"`
}

func (r *TextToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	if len(f) > 0 {
		return errors.New("TextToObjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TextToObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 报告结构化结果
		Template *Template `json:"Template,omitempty" name:"Template"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextType struct {

	// 类别Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 类别层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 类别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type TreatmentRecord struct {

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	DmissionCondition *string `json:"DmissionCondition,omitempty" name:"DmissionCondition"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaint *string `json:"ChiefComplaint,omitempty" name:"ChiefComplaint"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresent *string `json:"DiseasePresent,omitempty" name:"DiseasePresent"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSigns *string `json:"SymptomsAndSigns,omitempty" name:"SymptomsAndSigns"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExamination *string `json:"AuxiliaryExamination,omitempty" name:"AuxiliaryExamination"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExamination *string `json:"BodyExamination,omitempty" name:"BodyExamination"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExamination *string `json:"SpecialistExamination,omitempty" name:"SpecialistExamination"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	MentalExamination *string `json:"MentalExamination,omitempty" name:"MentalExamination"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecord *string `json:"CheckRecord,omitempty" name:"CheckRecord"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectResult *string `json:"InspectResult,omitempty" name:"InspectResult"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealing *string `json:"IncisionHealing,omitempty" name:"IncisionHealing"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentSuggestion *string `json:"TreatmentSuggestion,omitempty" name:"TreatmentSuggestion"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUpRequirements *string `json:"FollowUpRequirements,omitempty" name:"FollowUpRequirements"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcess *string `json:"CheckAndTreatmentProcess,omitempty" name:"CheckAndTreatmentProcess"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryCondition *string `json:"SurgeryCondition,omitempty" name:"SurgeryCondition"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionChanges *string `json:"ConditionChanges,omitempty" name:"ConditionChanges"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeCondition *string `json:"DischargeCondition,omitempty" name:"DischargeCondition"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *string `json:"PTNM,omitempty" name:"PTNM"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMM *string `json:"PTNMM,omitempty" name:"PTNMM"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMN *string `json:"PTNMN,omitempty" name:"PTNMN"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMT *string `json:"PTNMT,omitempty" name:"PTNMT"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	ECOG *string `json:"ECOG,omitempty" name:"ECOG"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	NRS *string `json:"NRS,omitempty" name:"NRS"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	KPS *string `json:"KPS,omitempty" name:"KPS"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeathDate *string `json:"DeathDate,omitempty" name:"DeathDate"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelapseDate *string `json:"RelapseDate,omitempty" name:"RelapseDate"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObservationDays *string `json:"ObservationDays,omitempty" name:"ObservationDays"`
}

type TuberInfo struct {

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *BlockInfo `json:"Type,omitempty" name:"Type"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitempty" name:"Part"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitempty" name:"Size" list`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Multiple *Multiple `json:"Multiple,omitempty" name:"Multiple"`

	// 纵横比
	// 注意：此字段可能返回 null，表示取不到有效值。
	AspectRatio *AspectRatio `json:"AspectRatio,omitempty" name:"AspectRatio"`

	// 边缘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edge *BlockInfo `json:"Edge,omitempty" name:"Edge"`

	// 内部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEcho *BlockInfo `json:"InnerEcho,omitempty" name:"InnerEcho"`

	// 外部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	RearEcho *BlockInfo `json:"RearEcho,omitempty" name:"RearEcho"`

	// 质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elastic *Elastic `json:"Elastic,omitempty" name:"Elastic"`

	// 形态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shape *BlockInfo `json:"Shape,omitempty" name:"Shape"`

	// 形态属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitempty" name:"ShapeAttr"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkinMedulla *BlockInfo `json:"SkinMedulla,omitempty" name:"SkinMedulla"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *BlockInfo `json:"Trend,omitempty" name:"Trend"`

	// 钙化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Calcification *BlockInfo `json:"Calcification,omitempty" name:"Calcification"`

	// 包膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envelope *BlockInfo `json:"Envelope,omitempty" name:"Envelope"`

	// 强化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enhancement *BlockInfo `json:"Enhancement,omitempty" name:"Enhancement"`

	// 淋巴结
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitempty" name:"LymphEnlargement"`

	// 淋巴门
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphDoor *BlockInfo `json:"LymphDoor,omitempty" name:"LymphDoor"`

	// 活动度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Activity *BlockInfo `json:"Activity,omitempty" name:"Activity"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *BlockInfo `json:"Operation,omitempty" name:"Operation"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDFI *BlockInfo `json:"CDFI,omitempty" name:"CDFI"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitempty" name:"Index" list`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeStatus *BlockInfo `json:"SizeStatus,omitempty" name:"SizeStatus"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitempty" name:"InnerEchoDistribution"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoType []*BlockInfo `json:"InnerEchoType,omitempty" name:"InnerEchoType" list`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline *BlockInfo `json:"Outline,omitempty" name:"Outline"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Structure *BlockInfo `json:"Structure,omitempty" name:"Structure"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Density *BlockInfo `json:"Density,omitempty" name:"Density"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vas *BlockInfo `json:"Vas,omitempty" name:"Vas"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cysticwall *BlockInfo `json:"Cysticwall,omitempty" name:"Cysticwall"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capsule *BlockInfo `json:"Capsule,omitempty" name:"Capsule"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitempty" name:"IsthmusThicknese"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitempty" name:"Src"`
}

type Value struct {

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// “”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent []*float64 `json:"Percent,omitempty" name:"Percent" list`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitempty" name:"Positive"`
}
