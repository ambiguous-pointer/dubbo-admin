// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"time"

	"github.com/google/uuid"
)

// Response defines unified API response format
type Response struct {
	Message   string `json:"message"`        // Response message
	Data      any    `json:"data,omitempty"` // Response data
	RequestID string `json:"request_id"`     // Request ID for tracking
	Timestamp int64  `json:"timestamp"`      // Response timestamp
}

// NewResponse creates a response
func NewResponse(message string, data any) *Response {
	return &Response{
		Message:   message,
		Data:      data,
		RequestID: generateRequestID(),
		Timestamp: time.Now().Unix(),
	}
}

// NewSuccessResponse creates a success response
func NewSuccessResponse(data any) *Response {
	return NewResponse("success", data)
}

// NewErrorResponse creates an error response
func NewErrorResponse(message string) *Response {
	return NewResponse(message, nil)
}

// ChatRequest defines streaming chat request
type ChatRequest struct {
	Message   string `json:"message" binding:"required"`   // User message
	SessionID string `json:"sessionID" binding:"required"` // Session ID
}

// generateRequestID generates a request ID
func generateRequestID() string {
	return "req_" + uuid.New().String()
}
