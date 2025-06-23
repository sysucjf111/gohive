package tcliservice

import "testing"

func TestTStatusString(t *testing.T) {
	tests := []struct {
		name     string
		status   *TStatus
		expected string
	}{
		{
			name: "完整的TStatus实例",
			status: &TStatus{
				StatusCode:   TStatusCode_SUCCESS_STATUS,
				InfoMessages: []string{"信息1", "信息2"},
				SqlState:     stringPtr("42000"),
				ErrorCode:    int32Ptr(1001),
				ErrorMessage: stringPtr("语法错误"),
			},
			expected: `TStatus{StatusCode: SUCCESS_STATUS(0), InfoMessages: ["信息1", "信息2"], SqlState: "42000", ErrorCode: 1001, ErrorMessage: "语法错误"}`,
		},
		{
			name: "只有必需字段的TStatus实例",
			status: &TStatus{
				StatusCode: TStatusCode_ERROR_STATUS,
			},
			expected: `TStatus{StatusCode: ERROR_STATUS(3), InfoMessages: <nil>, SqlState: <nil>, ErrorCode: <nil>, ErrorMessage: <nil>}`,
		},
		{
			name: "空InfoMessages切片的TStatus实例",
			status: &TStatus{
				StatusCode:   TStatusCode_SUCCESS_WITH_INFO_STATUS,
				InfoMessages: []string{},
			},
			expected: `TStatus{StatusCode: SUCCESS_WITH_INFO_STATUS(1), InfoMessages: [], SqlState: <nil>, ErrorCode: <nil>, ErrorMessage: <nil>}`,
		},
		{
			name: "部分指针字段有值的TStatus实例",
			status: &TStatus{
				StatusCode:   TStatusCode_STILL_EXECUTING_STATUS,
				InfoMessages: []string{"处理中..."},
				ErrorCode:    int32Ptr(2001),
				// SqlState 和 ErrorMessage 为 nil
			},
			expected: `TStatus{StatusCode: STILL_EXECUTING_STATUS(2), InfoMessages: ["处理中..."], SqlState: <nil>, ErrorCode: 2001, ErrorMessage: <nil>}`,
		},
		{
			name: "包含特殊字符的TStatus实例",
			status: &TStatus{
				StatusCode:   TStatusCode_INVALID_HANDLE_STATUS,
				InfoMessages: []string{"包含\"引号\"的消息", "包含\n换行符的消息"},
				SqlState:     stringPtr("HY000"),
				ErrorCode:    int32Ptr(3001),
				ErrorMessage: stringPtr("包含'单引号'和\"双引号\"的错误"),
			},
			expected: `TStatus{StatusCode: INVALID_HANDLE_STATUS(4), InfoMessages: ["包含\"引号\"的消息", "包含\n换行符的消息"], SqlState: "HY000", ErrorCode: 3001, ErrorMessage: "包含'单引号'和\"双引号\"的错误"}`,
		},
		{
			name:     "nil TStatus实例",
			status:   nil,
			expected: "<nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.String()
			if result != tt.expected {
				t.Errorf("TStatus.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 辅助函数：创建字符串指针
func stringPtr(s string) *string {
	return &s
}

// 辅助函数：创建int32指针
func int32Ptr(i int32) *int32 {
	return &i
}
