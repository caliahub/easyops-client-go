package model

/**
 * @Description 流程实例model
 * @Author 梁泽华Calia
 * @Date 2023/5/24 16:21
 **/

// GetProcessInstanceResponseBody 获取流程实例详情
type GetProcessInstanceResponseBody struct {
	ModelProcessInstance
	GetProcessInstanceResponseBody_2
}

type ModelProcessInstance struct {
	Category                     string                                             `json:"category,omitempty"`                     //服务分类
	Creator                      string                                             `json:"creator,omitempty"`                      //创建人
	Ctime                        string                                             `json:"ctime,omitempty"`                        //创建时间
	CurrentAssigneeList          []string                                           `json:"currentAssigneeList,omitempty"`          //当前审批人列表(废弃)
	Etime                        string                                             `json:"etime,omitempty"`                        //结束时间
	FlowableInstanceId           string                                             `json:"flowableInstanceId,omitempty"`           //flowable的实例ID
	FocusFields                  []ModelProcessInstanceFocusFields                  `json:"focusFields,omitempty"`                  //工单关注任务字段设置
	HandleWay                    string                                             `json:"handleWay,omitempty"`                    //工单处理方式(工单优先级,根据影响范围和紧急程度计算得来)
	InfluenceScope               string                                             `json:"influenceScope,omitempty"`               //影响范围(ITSC 标准字段)
	InstanceId                   string                                             `json:"instanceId,omitempty"`                   //ID
	IsCancelled                  bool                                               `json:"isCancelled,omitempty"`                  //是否作废
	IsComment                    bool                                               `json:"isComment,omitempty"`                    //是否已评价
	IsDelete                     bool                                               `json:"isDelete,omitempty"`                     //是否已被删除
	IsOldProcessInstance         bool                                               `json:"isOldProcessInstance,omitempty"`         //是否为老流程工单
	IsSubInstance                bool                                               `json:"isSubInstance,omitempty"`                //是否为子实例(子工单)
	IsSuspended                  bool                                               `json:"isSuspended,omitempty"`                  //是否挂起
	IsTimeout                    bool                                               `json:"isTimeout,omitempty"`                    //是否超时
	LastDiscussTime              string                                             `json:"lastDiscussTime,omitempty"`              //最近一次工单论述时间
	Name                         string                                             `json:"name,omitempty"`                         //名称
	OldProcessInstanceStruct     ModelProcessInstanceOldProcessInstanceStruct       `json:"oldProcessInstanceStruct"`               //老流程(流程平台1.0)部分不兼容数据(用一个struct存储)
	OperationTime                string                                             `json:"operationTime,omitempty"`                //操作时间
	OrderNum                     string                                             `json:"orderNum,omitempty"`                     //工单号
	RTime                        string                                             `json:"RTime,omitempty"`                        //剩余时间(sla 超时的剩余时间, - 表示已超时)
	ServiceId                    string                                             `json:"serviceId,omitempty"`                    //服务实例Id
	SlaStatus                    string                                             `json:"slaStatus,omitempty"`                    //sla状态[预警 或 超时], 空为正常状态
	Source                       string                                             `json:"source,omitempty"`                       //工单来源(默认 ITSC)
	Status                       string                                             `json:"status,omitempty"`                       //工单状态(running, done, closed)
	StepIdList                   []string                                           `json:"stepIdList,omitempty"`                   //流程实例（子任务）步骤列表(废弃)
	SubsequentConf               []ModelProcessInstanceSubsequentConf               `json:"subsequentConf,omitempty"`               //工单节点候选人
	SupervisorList               []string                                           `json:"supervisorList,omitempty"`               //督办人列表
	SuspendInfo                  ModelProcessInstanceSuspendInfo                    `json:"suspendInfo"`                            //当前挂起信息
	TimeoutTime                  string                                             `json:"timeoutTime,omitempty"`                  //工单超时时间
	Urgency                      string                                             `json:"urgency,omitempty"`                      //紧急程度(ITSC 标准字段)
	VersionRelevanceSubTaskInfo  []ModelProcessInstanceVersionRelevanceSubTaskInfo  `json:"versionRelevanceSubTaskInfo,omitempty"`  //被发起流程版本关联的子流程任务信息
	VersionRelevanceUserTaskInfo []ModelProcessInstanceVersionRelevanceUserTaskInfo `json:"versionRelevanceUserTaskInfo,omitempty"` //被发起流程版本关联的用户任务信息
	VisibleRange                 string                                             `json:"visibleRange,omitempty"`                 //可见范围(operator-处理人可见,mine仅自己可见,工单发起人),默认operator
}

type ModelProcessInstanceFocusFields struct {
	FocusFieldsOrder        []string          `json:"focusFieldsOrder,omitempty"`        //字段展示顺序
	FocusUserTaskFormFields []ModelFocusField `json:"focusUserTaskFormFields,omitempty"` //关注的用户任务表单字段
	UserTaskId              string            `json:"userTaskId,omitempty"`              //用户任务ID
}

type ModelFocusField struct {
	FocusFormVersionId string                   `json:"focusFormVersionId,omitempty"` //被关注任务表单版本ID
	FocusUserTaskId    string                   `json:"focusUserTaskId,omitempty"`    //被关注用户任务ID
	FormDefinitionPart []map[string]interface{} `json:"formDefinitionPart,omitempty"` //用户任务表单的部分表单定义（是将完整的表单定义拆出一部分保存）
}

type ModelProcessInstanceOldProcessInstanceStruct struct {
	Category                        string `json:"category,omitempty"`                        //流程分类
	Creator                         string `json:"creator,omitempty"`                         //创建者
	Ctime                           string `json:"ctime,omitempty"`                           //创建时间
	Id                              int    `json:"id,omitempty"`                              //老流程实例表主键id
	OldVersionRelevanceUserTaskInfo string `json:"oldVersionRelevanceUserTaskInfo,omitempty"` //被发起流程版本关联的用户任务信息({userTaskId:formDefinition})
	Pid                             string `json:"pid,omitempty"`                             //流程Id
	Pname                           string `json:"pname,omitempty"`                           //流程名称
}

type ModelProcessInstanceSubsequentConf struct {
	Candidates []string `json:"candidates,omitempty"` //候选人
	UserTaskId string   `json:"userTaskId,omitempty"` //用户任务ID
}

type ModelProcessInstanceSuspendInfo struct {
	StepId          string `json:"stepId,omitempty"`          //挂起节点实例id
	SuspendNickname string `json:"suspendNickname,omitempty"` //挂起人昵称
	SuspendUsername string `json:"suspendUsername,omitempty"` //挂起人名
	UserTaskId      string `json:"userTaskId,omitempty"`      //挂起节点任务id
	UserTaskName    string `json:"userTaskName,omitempty"`    //挂起节点任务名称
}

type ModelProcessInstanceVersionRelevanceSubTaskInfo struct {
	SubProcessVersionId string `json:"subProcessVersionId,omitempty"` //子流程对应的流程版本Id
	SubTaskId           string `json:"subTaskId,omitempty"`           //子流程任务Id
}

type ModelProcessInstanceVersionRelevanceUserTaskInfo struct {
	FormDisplayMode string `json:"formDisplayMode,omitempty"` //通过时表单展示模式,"popup" | "side"
	FormVersionId   string `json:"formVersionId,omitempty"`   //表单版本Id
	UserTaskId      string `json:"userTaskId,omitempty"`      //用户任务Id
}

type GetProcessInstanceResponseBody_2 struct {
	AllowedOp             ModelInstanceAllowedOp                                `json:"allowedOp"`                       //当前用户所拥有的工单权限操作
	Comment               GetProcessInstanceResponseBodyComment                 `json:"comment,omitempty"`               //工单评价
	FinishedStepList      []string                                              `json:"finishedStepList,omitempty"`      //显示当前步骤和之前流转步骤任务Id列表(例a->b-c->b,此时显示a->b), 必填
	Process               GetProcessInstanceResponseBodyProcess                 `json:"process,omitempty"`               //流程定义
	ProcessVersion        GetProcessInstanceResponseBodyProcessVersion          `json:"processVersion,omitempty"`        //流程定义版本
	ServiceInstance       ModelBriefService                                     `json:"serviceInstance,omitempty"`       //工单关联服务实例
	ServiceRelevanceOrder []GetProcessInstanceResponseBodyServiceRelevanceOrder `json:"serviceRelevanceOrder,omitempty"` //工单关联单
	StepAllowedOp         ModelStepIdAllowedOp                                  `json:"stepAllowedOp,omitempty"`         //运行中任务相关信息
	StepList              []ModelBriefStepInfo                                  `json:"stepList,omitempty"`              //流程子步骤列表，必填
	StepOperationRecord   []ModelStepOperationRecord                            `json:"stepOperationRecord,omitempty"`   //流程实例步骤操作详细信息列表
	StopAts               []string                                              `json:"stopAts,omitempty"`               //工单所停留的任务id列表
	SubTaskList           []ModelBpmnSubProcess                                 `json:"subTaskList,omitempty"`           //子流程任务列表
	UserInfoMap           map[string]interface{}                                `json:"userInfoMap,omitempty"`           //用户名-呢称字典
	UserTaskInfo          []ModelTaskOperatorInfo                               `json:"userTaskInfo,omitempty"`          //当前用户任务信息
	UserTaskList          []ModelUserTaskView                                   `json:"userTaskList,omitempty"`          //用户任务列表，必填
}

type ModelInstanceAllowedOp struct {
	CanCancel  bool `json:"canCancel,omitempty"`  //是否能作废
	CanComment bool `json:"canComment,omitempty"` //是否能评论
	CanSuspend bool `json:"canSuspend,omitempty"` //是否能挂起
}

type GetProcessInstanceResponseBodyComment struct {
	Creator string `json:"creator,omitempty"` //评价人
	Ctime   string `json:"ctime,omitempty"`   //创建时间
	Memo    string `json:"memo,omitempty"`    //评价内容
	Pattern string `json:"pattern,omitempty"` //评价方式
	Rank    int    `json:"rank,omitempty"`    //评价等级
}

type GetProcessInstanceResponseBodyProcess struct {
	Category   string `json:"category,omitempty"`   //流程类别
	InstanceId string `json:"instanceId,omitempty"` //流程定义实例id
	Name       string `json:"name,omitempty"`       //流程定义名称
}
type GetProcessInstanceResponseBodyProcessVersion struct {
	BpmnXML     string `json:"bpmnXML,omitempty"`     //流程BPMN信息(流程详情需要返回,除此之外可以不返回或返回空字符串)
	InstanceId  string `json:"instanceId,omitempty"`  //实例id
	IsJumpable  bool   `json:"isJumpable,omitempty"`  //是否开启直通车
	VersionName string `json:"versionName,omitempty"` //版本名称
}

type ModelBriefService struct {
	Category   string `json:"category,omitempty"`   //分类(如:变更、问题)
	InstanceId string `json:"instanceId,omitempty"` //实例 id
	Name       string `json:"name,omitempty"`       //服务名称
	Owner      string `json:"owner,omitempty"`      //服务负责人工号
}

type GetProcessInstanceResponseBodyServiceRelevanceOrder struct {
	ProcessInstanceId          string `json:"processInstanceId,omitempty"`          //流程实例Id
	ProcessInstanceIsCancelled bool   `json:"processInstanceIsCancelled,omitempty"` //流程实例(工单)是否作废
	ProcessInstanceIsSuspended bool   `json:"processInstanceIsSuspended,omitempty"` //流程实例(工单)是否挂起
	ProcessInstanceName        string `json:"processInstanceName,omitempty"`        //流程实例(工单)名称
	ProcessInstanceNum         string `json:"processInstanceNum,omitempty"`         //流程实例(工单)号
	ProcessInstanceStatus      string `json:"processInstanceStatus,omitempty"`      //流程实例(工单)状态
	RelevanceType              string `json:"relevanceType,omitempty"`              //关联工单的类型：父工单（sup），子工单（sub），普通关联（normal）
	ServiceInstanceCategory    string `json:"serviceInstanceCategory,omitempty"`    //服务实例类型
	ServiceInstanceName        string `json:"serviceInstanceName,omitempty"`        //服务实例名称
	ServiceInstanceId          string `json:"serviceInstanceId,omitempty"`          //服务实例Id
}

type ModelStepIdAllowedOp struct {
	AllowedOp                 ModelAllowedOp                         `json:"allowedOp,omitempty"`                 //当前用户所拥有的权限操作
	EnableFollowUpTaskHandler bool                                   `json:"enableFollowUpTaskHandler,omitempty"` //启用后续任务处理人, true-启用, false-不启用
	ExtraAssigneeType         string                                 `json:"extraAssigneeType,omitempty"`         //任务加签类型
	IsExtraAssignee           bool                                   `json:"isExtraAssignee,omitempty"`           //是否为加签任务
	NextTaskAssignee          ModelProcessInstanceStepAssignee       `json:"nextTaskAssignee,omitempty"`          //下一节点处理对象，不区分是否子流程
	OtherRunningStep          []ModelStepIdAllowedOpOtherRunningStep `json:"otherRunningStep,omitempty"`          //其他运行中节点信息
	StepId                    string                                 `json:"stepId,omitempty"`                    //任务id
	UserTaskId                string                                 `json:"userTaskId,omitempty"`                //用户任务id
	UserTaskName              string                                 `json:"userTaskName,omitempty"`              //当前任务名
}

type ModelAllowedOp struct {
	CanAddExtraAssignee bool `json:"canAddExtraAssignee,omitempty"` //是否能加签
	CanAssignee         bool `json:"canAssignee,omitempty"`         //是否能转派
	CanCc               bool `json:"canCc,omitempty"`               //是否能分阅
	CanClaim            bool `json:"canClaim,omitempty"`            //是否能认领
	CanClose            bool `json:"canClose,omitempty"`            //是否可以关单（注意，关单是工单级别的操作）
	CanComment          bool `json:"canComment,omitempty"`          //是否能评价
	CanDistribute       bool `json:"canDistribute,omitempty"`       //是否能派单
	CanDone             bool `json:"canDone,omitempty"`             //是否能完成
	CanNextAssignee     bool `json:"canNextAssignee,omitempty"`     //是否能指派下个环节审批人
	CanSLAChange        bool `json:"canSLAChange,omitempty"`        //是否能SLA变更
	CanSuspend          bool `json:"canSuspend,omitempty"`          //是否能挂起/激活
	CanWithdraw         bool `json:"canWithdraw,omitempty"`         //是否能撤回
}

type ModelProcessInstanceStepAssignee struct {
	Group      []string `json:"group,omitempty"`      //被指派的用户组
	Handling   string   `json:"handling,omitempty"`   //处理方式
	User       []string `json:"user,omitempty"`       //被指派的人
	UserTaskId string   `json:"userTaskId,omitempty"` //用户任务ID
	UserType   string   `json:"userType,omitempty"`   //普通指派用户类型(loginUser当前登录用户,lastExec上一步执行人,specifyUser指定用户,lastExecLeader上一步执行人leader,historyExec历史节点处理人,dutyGroup值班组,userTree行政树)
}

type ModelStepIdAllowedOpOtherRunningStep struct {
	Name     string   `json:"name,omitempty"`     //任务名
	Operator []string `json:"operator,omitempty"` //任务操作人
	StepId   string   `json:"stepId,omitempty"`   //任务id
}

type ModelBriefStepInfo struct {
	Action                   string   `json:"action,omitempty"`                   //操作
	Consignors               []string `json:"consignors,omitempty"`               //委托人列表
	Ctime                    string   `json:"ctime,omitempty"`                    //创建时间/接收时间
	Etime                    string   `json:"etime,omitempty"`                    //结束时间
	FileInfo                 string   `json:"fileInfo,omitempty"`                 //1.0流程任务表单关联的附件信息,json字符串 文件名:check_sum(当isOldProcessInstance为True时,可能有值)
	FormData                 string   `json:"formData,omitempty"`                 //流程实例表单基础信息,json字典
	InstanceId               string   `json:"instanceId,omitempty"`               //ID
	IsExtraAssignee          bool     `json:"isExtraAssignee,omitempty"`          //是否处于加签状态
	IsSubStep                bool     `json:"isSubStep,omitempty"`                //任务是否为子流程步骤
	Memo                     string   `json:"memo,omitempty"`                     //操作说明
	Mtime                    string   `json:"mtime,omitempty"`                    //修改时间
	Operator                 string   `json:"operator,omitempty"`                 //操作人
	Otime                    string   `json:"otime,omitempty"`                    //操作时间
	Status                   string   `json:"status,omitempty"`                   //状态(running当前正在流转步骤, done通过, rejected驳回, withdraw撤回)
	SubProcessInstanceId     string   `json:"subProcessInstanceId,omitempty"`     //子流程工单实例Id(当为子流程步骤时非空字符串)
	SubProcessInstanceStepId string   `json:"subProcessInstanceStepId,omitempty"` //子工单的任务实例Id(默认取生成子工单的第一个任务实例Id,当为子流程步骤时非空字符串)
	TaskName                 string   `json:"taskName,omitempty"`                 //任务名称
	ToolStatus               string   `json:"toolStatus,omitempty"`               //工具执行总状态
	Type                     string   `json:"type,omitempty"`                     //流程实例步骤类型(common->普通任务, countersign->会签任务)
	UserTaskId               string   `json:"userTaskId,omitempty"`               //用户任务/子流程 id
}

type ModelStepOperationRecord struct {
	Action               string              `json:"action,omitempty"`               //操作(distribute分配,claim认领, done:完成(通过/驳回), rejected:驳回, "transfer":转派, "cc":分阅, "suspended":挂起, activate:重启, "cancel":作废, "withdraw":撤回, "delete":删除, "relevance":关联工单)
	ExtraAssigneeType    string              `json:"extraAssigneeType,omitempty"`    //加签类型，加签，后置加签
	FormData             string              `json:"formData,omitempty"`             //表单数据（json字符串）
	IsExtraAssignee      bool                `json:"isExtraAssignee,omitempty"`      //是否是加签任务
	IsSubProcess         bool                `json:"isSubProcess,omitempty"`         //步骤是否为子流程
	Memo                 string              `json:"memo,omitempty"`                 //操作说明
	OperationTime        string              `json:"operationTime,omitempty"`        //操作时间（yyyy-mm-dd hh:mm:ss）
	Operator             ModelAccount        `json:"operator,omitempty"`             //操作用户的详细信息
	ProcessInstanceId    string              `json:"processInstanceId,omitempty"`    //流程实例id（CMDB的instanceId）
	RecordCtime          string              `json:"recordCtime,omitempty"`          //记录生成时间
	StepId               string              `json:"stepId,omitempty"`               //任务Id
	SubProcessInstanceId string              `json:"subProcessInstanceId,omitempty"` //子流程ID
	TaskName             string              `json:"taskName,omitempty"`             //步骤任务名称
	ToUser               []ModelAccount      `json:"toUser,omitempty"`               //接收用户的详细信息
	ToUserGroup          []ModelAccountGroup `json:"toUserGroup,omitempty"`          //接收用户组的详细信息
	UserTaskId           string              `json:"userTaskId,omitempty"`           //步骤任务ID
}

type ModelAccount struct {
	InstanceId string `json:"instanceId,omitempty"` //用户实例ID，cmdb中的instanceId
	Nickname   string `json:"nickname,omitempty"`   //昵称
	State      string `json:"state,omitempty"`      //账户状态（invalid 禁用，valid 启用）
	UserIcon   string `json:"userIcon,omitempty"`   //用户头像地址
	Username   string `json:"username,omitempty"`   //用户名
}

type ModelAccountGroup struct {
	GroupName  string `json:"groupName,omitempty"`  //组名
	InstanceId string `json:"instanceId,omitempty"` //用户组ID，cmdb的instanceId，前端有可能传个":instanceId"
}

type ModelBpmnSubProcess struct {
	CalledElement         string                         `json:"calledElement,omitempty"`         //子流程Id(对应流程版本中的processId)
	CalledElementTenantId string                         `json:"calledElementTenantId,omitempty"` //TenantId
	Id                    string                         `json:"id,omitempty"`                    //ID
	Links                 ModelBpmnLinks                 `json:"links,omitempty"`                 //指向链接
	Name                  string                         `json:"name,omitempty"`                  //名称
	ProcessOp             []ModelBpmnSubProcessProcessOp `json:"processOp,omitempty"`             //流程操作
}

type ModelBpmnLinks struct {
	Incoming []string `json:"incoming,omitempty"` //来源列表
	Outgoing []string `json:"outgoing,omitempty"` //指向列表
}

type ModelBpmnSubProcessProcessOp struct {
	IsSubProcess bool   `json:"isSubProcess,omitempty"` //是否为子流程
	Name         string `json:"name,omitempty"`         //操作名称(通过/驳回/结束/xxx)
	TargetTaskId string `json:"targetTaskId,omitempty"` //流转下一目标任务定义Id
}

type ModelTaskOperatorInfo struct {
	Assignee      []string `json:"assignee,omitempty"`      //当前指派的用户
	AssigneeGroup []string `json:"assigneeGroup,omitempty"` //当前指派的用户组
	Role          string   `json:"role,omitempty"`          //当前指派的用户角色
	Status        string   `json:"status,omitempty"`        //当前用户任务的状态
	TaskName      string   `json:"taskName,omitempty"`      //任务名称
	Type          string   `json:"type,omitempty"`          //当前用户任务类型 (assignee->普通任务, assigneeList->会签任务)
	UserTaskId    string   `json:"userTaskId,omitempty"`    //用户任务Id
}

type ModelUserTaskView struct {
	AssigneeGroups     []string                     `json:"assigneeGroups,omitempty"`     //关联用户组,多个用户组名称
	AssigneeListUser   []string                     `json:"assigneeListUser,omitempty"`   //关联用户,多个用户名称
	FormDefinition     string                       `json:"formDefinition,omitempty"`     //表单定义json
	FormDisplayMode    string                       `json:"formDisplayMode,omitempty"`    //通过时表单展示模式
	FormExpressionName string                       `json:"formExpressionName,omitempty"` //由表格决定多分支流转,需要制定变量名称(变量的key,例如pass==1,name传pass)
	FormName           string                       `json:"formName,omitempty"`           //表单名称
	FormVersionId      string                       `json:"formVersionId,omitempty"`      //表单版本Id
	Handling           string                       `json:"handling,omitempty"`           //节点处理方式(directly->直接处理 send_directly->先派单后处理 claim_directly->先认领后处理 send_claim_directly->先派单,再认领,再处理)
	Id                 string                       `json:"id,omitempty"`                 //ID
	IsFormDecision     string                       `json:"isFormDecision,omitempty"`     //是否为表格决定多分支流转,默认为字符串 "0"表示否, "1"表示由表格决定
	IsNextPar          bool                         `json:"isNextPar,omitempty"`          //是否连接并行网关
	JumpableNodes      []string                     `json:"jumpableNodes,omitempty"`      //可驳回的节点列表，形如`Activity_2jlsdf:1`，后面1或0代表是否可跳回
	LabelViews         []string                     `json:"labelViews,omitempty"`         //任务操作文案设置，key:value格式，key为操作项英文，value为对应的中文文案 任务操作：pass:通过 withdraw:撤回 claim:认领 reject:驳回 jump:直达 assign:转派 distribute:派单
	Name               string                       `json:"name,omitempty"`               //名称
	OperationConf      []ModelBpmnUserTask          `json:"operationConf,omitempty"`      //当前节点操作对应的处理人配置
	ProcessOp          []ModelUserTaskViewProcessOp `json:"processOp,omitempty"`          //流程操作
	StandardFields     []ModelStandardField         `json:"standardFields,omitempty"`     //当前表单所引用标准字段信息
	SubsequentConf     []ModelBpmnUserTask          `json:"subsequentConf,omitempty"`     //后续节点处理人配置
	Type               string                       `json:"type,omitempty"`               //用户任务类型 (assignee->普通任务, assigneeList->会签任务)
}

type ModelBpmnUserTask struct {
	AssignStrategy     string                               `json:"assignStrategy,omitempty"`     //指派策略(emptySkip为空跳过, emptyAssign为空转人工,sampleSkip与上节点相同审批人跳过)
	AssignType         string                               `json:"assignType,omitempty"`         //指派类型(assignee普通指派,assigneeList会签)
	AssigneeGroups     []string                             `json:"assigneeGroups,omitempty"`     //关联用户组,多个用户组名称
	AssigneeListUser   []string                             `json:"assigneeListUser,omitempty"`   //关联用户,多个用户名称
	AssigneeValue      string                               `json:"assigneeValue,omitempty"`      //用户指派类型值(historyExec类型值为userTaskId, dutyGroup类型值为值班组id, userTree为用户)
	FormExpressionName string                               `json:"formExpressionName,omitempty"` //由表格决定多分支流转,需要制定变量名称(变量的key,例如pass==1,name传pass)
	Handling           string                               `json:"handling,omitempty"`           //节点处理方式(directly->直接处理 send_directly->先派单后处理 claim_directly->先认领后处理 send_claim_directly->先派单,再认领,再处理)
	Id                 string                               `json:"id,omitempty"`                 //ID
	IsFormDecision     string                               `json:"isFormDecision,omitempty"`     //是否为表格决定多分支流转,默认为字符串 "0"表示否, "1"表示由表格决定
	IsNextPar          bool                                 `json:"isNextPar,omitempty"`          //是否连接并行网关
	JumpableNodes      []string                             `json:"jumpableNodes,omitempty"`      //可驳回的节点列表，形如`Activity_2jlsdf:1`，后面1或0代表是否可跳回
	LabelViews         []string                             `json:"labelViews,omitempty"`         //任务操作文案设置，key:value格式，key为操作项英文，value为对应的中文文案 任务操作：pass:通过 withdraw:撤回 claim:认领 reject:驳回 jump:直达 assign:转派 distribute:派单
	Links              ModelBpmnLinks                       `json:"links"`                        //指向链接
	MultiInstStrategy  []ModelBpmnUserTaskMultiInstStrategy `json:"multiInstStrategy,omitempty"`  //多实例策略
	Name               string                               `json:"name,omitempty"`               //名称
	OperationConf      []ModelBpmnUserTaskOperationConf     `json:"operationConf,omitempty"`      //当前节点操作对应的处理人配置
	OpsAllowed         []string                             `json:"opsAllowed,omitempty"`         //本节点人员操作权限, "done":完成(通过/驳回),"assignee":转派, "cc":分阅, "suspended":挂起,"distribute":派单,"claim":认领 "withdraw":撤回
	ProcessOp          []ModelBpmnUserTaskProcessOp         `json:"processOp,omitempty"`          //流程操作
	Script             []ModelBpmnUserTaskScript            `json:"script,omitempty"`             //用户任务脚本信息
	SetAssignee        bool                                 `json:"setAssignee,omitempty"`        //是否允许本节点人员设置下节点处理人员
	SubsequentConf     []ModelBpmnUserTaskSubsequentConf    `json:"subsequentConf,omitempty"`     //后续节点处理人配置
	TriggerIdList      []string                             `json:"triggerIdList,omitempty"`      //关联触发器,多个触发器实例id
	UserType           string                               `json:"userType,omitempty"`           //普通指派用户类型(loginUser当前登录用户,lastExec上一步执行人,specifyUser指定用户,lastExecLeader上一步执行人leader,historyExec历史节点处理人,dutyGroup值班组,userTree行政树)
}

type ModelBpmnUserTaskMultiInstStrategy struct {
	IsSequential string `json:"IsSequential,omitempty"` //是否是串行
}

type ModelBpmnUserTaskOperationConf struct {
	Candidates []string `json:"candidates,omitempty"` //候选人, eg, ["easyops", "user1"]
	Operation  string   `json:"operation,omitempty"`  //操作类型
}

type ModelBpmnUserTaskProcessOp struct {
	ConditionExpression ModelProcessVariable `json:"conditionExpression,omitempty"` //流程变量
	IsSubProcess        bool                 `json:"isSubProcess,omitempty"`        //是否为子流程
	Name                string               `json:"name,omitempty"`                //操作名称(通过/驳回/结束/xxx)
	TargetTaskId        string               `json:"targetTaskId,omitempty"`        //流转下一目标任务定义Id
}

type ModelProcessVariable struct {
	Name  string `json:"name,omitempty"`  //变量名称
	Value string `json:"value,omitempty"` //变量值
}

type ModelBpmnUserTaskScript struct {
	Desc            string   `json:"desc,omitempty"`            //脚本描述
	Name            string   `json:"name,omitempty"`            //脚本名称
	OpMode          string   `json:"opMode,omitempty"`          //脚本执行方式(async:异步,sync:同步)
	ScriptType      string   `json:"scriptType,omitempty"`      //脚本类型(pre前置脚本,rear后置脚本)
	SupportCategory []string `json:"supportCategory,omitempty"` //脚本触发方式(通过:done,派单:distribute,认领:claim,转派:assignee,分阅:cc,撤回:withdraw)
	ToolIds         []string `json:"toolIds,omitempty"`         //执行工具Id列表
}

type ModelBpmnUserTaskSubsequentConf struct {
	Candidates   []string `json:"candidates,omitempty"`   //候选人, eg, ["easyops", "user1"]
	Label        string   `json:"label,omitempty"`        //标签
	UserTaskId   string   `json:"userTaskId,omitempty"`   //任务Id
	UserTaskName string   `json:"userTaskName,omitempty"` //任务名称
}

type ModelUserTaskViewProcessOp struct {
	ConditionExpression ModelProcessVariable `json:"conditionExpression"`    //流程变量
	Name                string               `json:"name,omitempty"`         //操作名称(通过/驳回/xxx)
	TargetTaskId        string               `json:"targetTaskId,omitempty"` //流转下一目标任务定义Id
}

type ModelStandardField struct {
	Builtin      bool   `json:"builtin,omitempty"`      //内置标识 (true 为内置自动)
	Creator      string `json:"creator,omitempty"`      //创建人 (cmdb 内置字段)
	Ctime        string `json:"ctime,omitempty"`        //创建时间 (cmdb 内置字段)
	Default      string `json:"default,omitempty"`      //字段默认值
	Desc         string `json:"desc,omitempty"`         //字段说明
	InstanceId   string `json:"instanceId,omitempty"`   //标准字段实例Id
	Key          string `json:"key,omitempty"`          //唯一标识(英文名, 不可重复)
	Kind         string `json:"kind,omitempty"`         //字段类型(枚举)
	Name         string `json:"name,omitempty"`         //名称(中文名, 可重复)
	Readonly     bool   `json:"readonly,omitempty"`     //只读 (true 只读)
	Required     bool   `json:"required,omitempty"`     //必填 (true 必填)
	SourceConfig string `json:"sourceConfig,omitempty"` //数据源配置（JSON字符串）
	SourceType   string `json:"sourceType,omitempty"`   //数据源类型
}

// StartProcessInstanceRequestBody 发起流程实例
type StartProcessInstanceRequestBody struct {
	ITSCInfluenceScope  string                                              `json:"ITSC_influenceScope,omitempty"` //影响范围(ITSC 标准字段)
	ITSCUrgency         string                                              `json:"ITSC_urgency,omitempty"`        //紧急程度(ITSC 标准字段)
	AssigneeGroups      []string                                            `json:"assigneeGroups,omitempty"`      //下一步指派组
	AssigneeUsers       []string                                            `json:"assigneeUsers,omitempty"`       //下一步指派人
	DepartmentId        string                                              `json:"departmentId,omitempty"`        //部门id
	FormData            string                                              `json:"formData,omitempty"`            //流程实例表单基础信息,json字典
	HandleWay           string                                              `json:"handleWay,omitempty"`           //工单处理方式(工单优先级,根据影响范围和紧急程度计算得来)
	IsSupervision       bool                                                `json:"isSupervision,omitempty"`       //是否督办
	Name                string                                              `json:"name,omitempty"`                //流程实例名称,必填
	RelevanceInstanceId string                                              `json:"relevanceInstanceId,omitempty"` //关联工单Id(由关联单发起时传,直接由服务发起不用传)
	ServiceId           string                                              `json:"serviceId,omitempty"`           //服务Id,必填
	Source              string                                              `json:"source,omitempty"`              //工单来源(默认-ITSC, easyops告警平台-EasyMonitor)
	SubsequentAssignee  []StartProcessInstanceRequestBodySubsequentAssignee `json:"subsequentAssignee,omitempty"`  //后续节点处理人
	SupervisorUserList  []string                                            `json:"supervisorUserList,omitempty"`  //督办用户列表
	VariableName        string                                              `json:"variableName,omitempty"`        //条件变量名称（如pass)
	VariableValue       string                                              `json:"variableValue,omitempty"`       //条件变量值(如0,1)
	VisibleRange        string                                              `json:"visibleRange,omitempty"`        //可见范围(operator-处理人可见,mine仅自己可见,工单发起人),默认operator,必填
}

type StartProcessInstanceRequestBodySubsequentAssignee struct {
	Candidates []string `json:"candidates,omitempty"` //后续节点处理人, eg, ["easyops", "user1"]
	UserTaskId string   `json:"userTaskId,omitempty"` //任务Id
}

type StartProcessInstanceResponseBody struct {
	ModelProcessInstance
	StartProcessInstanceResponseBody_2
}

type StartProcessInstanceResponseBody_2 struct {
	StepList     []StartProcessInstanceResponseBodyStepList     `json:"stepList"`     //流程子步骤列表，必填
	UserTaskList []StartProcessInstanceResponseBodyUserTaskList `json:"userTaskList"` //用户任务列表，必填
}

type StartProcessInstanceResponseBodyStepList struct {
	Assignee   string `json:"assignee,omitempty"`   //流程负责人
	FormData   string `json:"formData,omitempty"`   //流程实例表单基础信息,json字典
	UserTaskId string `json:"userTaskId,omitempty"` //用户任务ID
}

type StartProcessInstanceResponseBodyUserTaskList struct {
	FormDefinition string `json:"formDefinition,omitempty"` //表单定义json
	Id             string `json:"id,omitempty"`             //用户任务ID
	Name           string `json:"name,omitempty"`           //用户任务名称
}
