package provider

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/yinxulai/goutils/random"
	"github.com/grpcbrick/queues/models"
	"github.com/grpcbrick/queues/standard"
)

// NewService NewService
func NewService() *Service {
	godotenv.Load()
	service := new(Service)
	return service
}

// Service Service
type Service struct {
}

// CreateTask CreateTask
func (srv *Service) CreateTask(ctx context.Context, req *standard.CreateTaskRequest) (resp *standard.CreateTaskResponse, err error) {
	var count int
	HashCode := random.String(20)
	resp = new(standard.CreateTaskResponse)

	if req.Task.Channel == "" {
		resp.State = standard.State_CHANNEL_INVALID
		resp.Message = "不支持处理的的任务"
		return resp, nil
	}

	// 生成 HashCode
	for { // 不停的检查是否重复、如果重复就重新生成
		err = countTaskByHashCodeStmt.GetContext(ctx, &count, map[string]interface{}{"HashCode": HashCode})
		if err != nil {
			resp.State = standard.State_DB_OPERATION_FATLURE
			resp.Message = err.Error()
			return resp, nil
		}

		if count > 0 {
			HashCode = random.String(20)
			continue // 再来一次
		}

		break // 推出
	}

	// 执行
	req.Task.HashCode = HashCode
	_, err = insertTaskByChannelStmt.ExecContext(ctx, req.Task)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"
	resp.Data = HashCode

	return resp, nil
}

// QueryTaskByID QueryTaskByID
func (srv *Service) QueryTaskByID(ctx context.Context, req *standard.QueryTaskByIDRequest) (resp *standard.QueryTaskByIDResponse, err error) {
	tasks := []*models.Task{}
	resp = new(standard.QueryTaskByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	rows, err := queryTaskByIDStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localTask models.Task
		err = rows.StructScan(&localTask)
		if err == nil {
			tasks = append(tasks, &localTask)
		}
	}

	if len(tasks) <= 0 {
		resp.State = standard.State_NOT_EXIST
		resp.Message = "该任务不存在"
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = tasks[0].OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

// CancelTaskByID CancelTaskByID
func (srv *Service) CancelTaskByID(ctx context.Context, req *standard.CancelTaskByIDRequest) (resp *standard.CancelTaskByIDResponse, err error) {
	tasks := []*models.Task{}
	resp = new(standard.CancelTaskByIDResponse)

	if req.ID == 0 {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 ID"
		return resp, nil
	}

	// 查询任务存在
	rows, err := queryTaskByIDStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localTask models.Task
		err = rows.StructScan(&localTask)
		if err == nil {
			tasks = append(tasks, &localTask)
		}
	}

	if len(tasks) <= 0 {
		resp.State = standard.State_NOT_EXIST
		resp.Message = "该任务不存在"
		return resp, nil
	}

	//执行请求
	tasks[0].State = "CANCEL"
	_, err = updateTaskByIDStmt.ExecContext(ctx, tasks[0])
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "取消任务成功"
	return resp, nil
}

// QueryTaskByOwner QueryTaskByOwner
func (srv *Service) QueryTaskByOwner(ctx context.Context, req *standard.QueryTaskByOwnerRequest) (resp *standard.QueryTaskByOwnerResponse, err error) {
	var count uint64
	tasks := []*models.Task{}
	stdtasks := []*standard.Task{}
	resp = new(standard.QueryTaskByOwnerResponse)

	err = countTaskByOwnerStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	rows, err := queryTaskByOwnerStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localTask models.Task
		err = rows.StructScan(&localTask)
		if err == nil {
			tasks = append(tasks, &localTask)
		}
	}

	for _, task := range tasks {
		stdtasks = append(stdtasks, task.OutProtoStruct())
	}

	resp.State = standard.State_SUCCESS
	resp.Data = stdtasks
	resp.Total = count
	resp.Message = "查询成功"

	return resp, nil
}

// QueryTaskByHashCode QueryTaskByHashCode
func (srv *Service) QueryTaskByHashCode(ctx context.Context, req *standard.QueryTaskByHashCodeRequest) (resp *standard.QueryTaskByHashCodeResponse, err error) {
	tasks := []*models.Task{}
	resp = new(standard.QueryTaskByHashCodeResponse)

	if req.HashCode == "" {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = "无效的 HashCode"
		return resp, nil
	}

	rows, err := queryTaskByHashCodeStmt.QueryxContext(ctx, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	for rows.Next() {
		var localTask models.Task
		err = rows.StructScan(&localTask)
		if err == nil {
			tasks = append(tasks, &localTask)
		}
	}

	if len(tasks) <= 0 {
		resp.State = standard.State_NOT_EXIST
		resp.Message = "该任务不存在"
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Data = tasks[0].OutProtoStruct()
	resp.Message = "查询成功"
	return resp, nil
}

// QueryLengthByChannel QueryLengthByChannel
func (srv *Service) QueryLengthByChannel(ctx context.Context, req *standard.QueryLengthByChannelRequest) (resp *standard.QueryLengthByChannelResponse, err error) {
	var count uint64
	resp = new(standard.QueryLengthByChannelResponse)

	err = countTaskByChannelStmt.GetContext(ctx, &count, req)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "查询成功"
	resp.Data = count

	return nil, nil
}

// ReportTaskResult 报告任务结果
func (srv *Service) ReportTaskResult(standard.Queues_ReportTaskResultServer) error {

	return nil
}

// ReceiveQueueByChannel 领取一个任务
func (srv *Service) ReceiveQueueByChannel(standard.Queues_ReceiveQueueByChannelServer) error {

	return nil
}
