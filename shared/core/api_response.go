package core

import (
	"net/http"
	"time"
)

// ApiResponse là cấu trúc chuẩn cho tất cả phản hồi của API
type ApiResponse struct {
	Success   bool        `json:"success"`             // true nếu thành công, false nếu thất bại
	Code      int         `json:"code"`                // HTTP status code
	Message   string      `json:"message,omitempty"`   // Mô tả ngắn gọn
	Data      interface{} `json:"data,omitempty"`      // Dữ liệu trả về (nếu có)  // Chi tiết lỗi (nếu có)
	Timestamp string      `json:"timestamp,omitempty"` // Thời gian phản hồi
}

// Success tạo phản hồi thành công
func Success(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// Error tạo phản hồi lỗi
func Error(code int, message string, err interface{}) ApiResponse {
	return ApiResponse{
		Success:   false,
		Code:      code,
		Message:   message,
		Data:      err,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
