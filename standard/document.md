# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [standard.proto](#standard.proto)
    - [CancelTaskByIDRequest](#standard.CancelTaskByIDRequest)
    - [CancelTaskByIDResponse](#standard.CancelTaskByIDResponse)
    - [CreateTaskRequest](#standard.CreateTaskRequest)
    - [CreateTaskResponse](#standard.CreateTaskResponse)
    - [QueryLengthByChannelRequest](#standard.QueryLengthByChannelRequest)
    - [QueryLengthByChannelResponse](#standard.QueryLengthByChannelResponse)
    - [QueryTaskByHashCodeRequest](#standard.QueryTaskByHashCodeRequest)
    - [QueryTaskByHashCodeResponse](#standard.QueryTaskByHashCodeResponse)
    - [QueryTaskByIDRequest](#standard.QueryTaskByIDRequest)
    - [QueryTaskByIDResponse](#standard.QueryTaskByIDResponse)
    - [QueryTaskByOwnerRequest](#standard.QueryTaskByOwnerRequest)
    - [QueryTaskByOwnerResponse](#standard.QueryTaskByOwnerResponse)
    - [ReceiveQueueByChannelRequest](#standard.ReceiveQueueByChannelRequest)
    - [ReceiveQueueByChannelResponse](#standard.ReceiveQueueByChannelResponse)
    - [ReportTaskResultRequest](#standard.ReportTaskResultRequest)
    - [ReportTaskResultResponse](#standard.ReportTaskResultResponse)
    - [Task](#standard.Task)
  
  
  
    - [Queues](#standard.Queues)
  

- [Scalar Value Types](#scalar-value-types)



<a name="standard.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## standard.proto



<a name="standard.CancelTaskByIDRequest"></a>

### CancelTaskByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |






<a name="standard.CancelTaskByIDResponse"></a>

### CancelTaskByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Data | [Task](#standard.Task) | repeated |  |






<a name="standard.CreateTaskRequest"></a>

### CreateTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Task | [Task](#standard.Task) |  | 任务 |






<a name="standard.CreateTaskResponse"></a>

### CreateTaskResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Data | [string](#string) |  | 任务的 HashCode |






<a name="standard.QueryLengthByChannelRequest"></a>

### QueryLengthByChannelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Channel | [string](#string) |  | 指定频道 |






<a name="standard.QueryLengthByChannelResponse"></a>

### QueryLengthByChannelResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Data | [uint64](#uint64) |  |  |






<a name="standard.QueryTaskByHashCodeRequest"></a>

### QueryTaskByHashCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| HashCode | [string](#string) |  | 任务的唯一哈希码 |






<a name="standard.QueryTaskByHashCodeResponse"></a>

### QueryTaskByHashCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Data | [Task](#standard.Task) |  |  |






<a name="standard.QueryTaskByIDRequest"></a>

### QueryTaskByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  |  |






<a name="standard.QueryTaskByIDResponse"></a>

### QueryTaskByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Data | [Task](#standard.Task) |  |  |






<a name="standard.QueryTaskByOwnerRequest"></a>

### QueryTaskByOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Owner | [uint64](#uint64) |  | 所属用户 |
| Limit | [uint64](#uint64) |  |  |
| Offset | [uint64](#uint64) |  |  |






<a name="standard.QueryTaskByOwnerResponse"></a>

### QueryTaskByOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Total | [uint64](#uint64) |  |  |
| Data | [Task](#standard.Task) | repeated |  |






<a name="standard.ReceiveQueueByChannelRequest"></a>

### ReceiveQueueByChannelRequest
领取任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Channel | [string](#string) |  | 指定频道 |






<a name="standard.ReceiveQueueByChannelResponse"></a>

### ReceiveQueueByChannelResponse
返回一个任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Tasks | [Task](#standard.Task) |  | 任务 |






<a name="standard.ReportTaskResultRequest"></a>

### ReportTaskResultRequest
汇报任务处理结果


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  | 任务 ID |
| State | [uint64](#uint64) |  | 任务状态 |
| Output | [string](#string) |  | 任务输出 |






<a name="standard.ReportTaskResultResponse"></a>

### ReportTaskResultResponse
汇报任务的响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [uint64](#uint64) |  |  |
| Message | [string](#string) |  |  |
| Data | [uint64](#uint64) |  |  |






<a name="standard.Task"></a>

### Task



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [uint64](#uint64) |  | 当前的任务 ID |
| Next | [uint64](#uint64) |  | 上一个任务 在 组合任务/任务链 里如果为 0 就是首任务 |
| Prior | [uint64](#uint64) |  | 下一个任务 在 组合任务/任务链 里如果为 0 就是尾任务 |
| Owner | [uint64](#uint64) |  | 所属用户 |
| State | [string](#string) |  | 任务状态 |
| Input | [string](#string) |  | 任务输入 |
| Output | [string](#string) |  | 任务输出 |
| Channel | [string](#string) |  | 指定频道 |
| HashCode | [string](#string) |  | 任务的唯一哈希码 |
| RetryCount | [uint64](#uint64) |  | 重试次数 |
| CreateTime | [string](#string) |  | 创建时间 |
| UpdateTime | [string](#string) |  | 更新时间 |
| RetryMaxLimit | [uint64](#uint64) |  | 最大重试次数 |





 

 

 


<a name="standard.Queues"></a>

### Queues


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTask | [CreateTaskRequest](#standard.CreateTaskRequest) | [CreateTaskResponse](#standard.CreateTaskResponse) |  |
| QueryTaskByID | [QueryTaskByIDRequest](#standard.QueryTaskByIDRequest) | [QueryTaskByIDResponse](#standard.QueryTaskByIDResponse) |  |
| CancelTaskByID | [CancelTaskByIDRequest](#standard.CancelTaskByIDRequest) | [CancelTaskByIDResponse](#standard.CancelTaskByIDResponse) |  |
| QueryTaskByOwner | [QueryTaskByOwnerRequest](#standard.QueryTaskByOwnerRequest) | [QueryTaskByOwnerResponse](#standard.QueryTaskByOwnerResponse) |  |
| QueryTaskByHashCode | [QueryTaskByHashCodeRequest](#standard.QueryTaskByHashCodeRequest) | [QueryTaskByHashCodeResponse](#standard.QueryTaskByHashCodeResponse) |  |
| QueryLengthByChannel | [QueryLengthByChannelRequest](#standard.QueryLengthByChannelRequest) | [QueryLengthByChannelResponse](#standard.QueryLengthByChannelResponse) |  |
| ReportTaskResult | [ReportTaskResultRequest](#standard.ReportTaskResultRequest) stream | [ReportTaskResultResponse](#standard.ReportTaskResultResponse) stream |  |
| ReceiveQueueByChannel | [ReceiveQueueByChannelRequest](#standard.ReceiveQueueByChannelRequest) stream | [ReceiveQueueByChannelResponse](#standard.ReceiveQueueByChannelResponse) stream |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

