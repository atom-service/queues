package models

import (
	"database/sql"

	"github.com/grpcbrick/queues/standard"
)

// Task 任务
type Task struct {
	Owner         sql.NullInt64  // 拥有者
	State         sql.NullString // 状态
	Input         sql.NullString // 输入
	Output        sql.NullString // 输出
	Channel       sql.NullString // 任务频道
	HashCode      sql.NullString // 任务 hash
	RetryCount    sql.NullInt64  // 重试次数
	RetryMaxLimit sql.NullInt64  // 最多重试次数
	DeletedTime   sql.NullTime   // 删除时间
	CreatedTime   sql.NullTime   // 创建时间
	UpdatedTime   sql.NullTime   // 更新时间
}

// LoadProtoStruct LoadProtoStruct
func (srv *Task) LoadProtoStruct(task *standard.Task) {

}

// OutProtoStruct OutProtoStruct
func (srv *Task) OutProtoStruct() *standard.Task {
	task := new(standard.Task)

	return task
}
