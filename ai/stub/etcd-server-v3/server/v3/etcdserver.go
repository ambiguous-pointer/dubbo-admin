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

// Package etcdserver is a stub to avoid importing the full etcd server.
// This stub only provides the minimum types needed for dependency resolution.
// The actual etcd server has incompatible OpenTelemetry dependencies.
package etcdserver

import "context"

// Server is a stub type.
type Server struct{}

// ServerOpts is a stub type.
type ServerOpts struct{}

// Config is a stub type.
type Config struct{}

// V2 is a stub type.
type V2 struct{}

// SnapServer is a stub type.
type SnapServer struct{}

// RAFTServer is a stub type.
type RAFTServer struct{}

// Stub functions
func NewServer() *Server                      { return &Server{} }
func (s *Server) Start(context.Context) error { return nil }
func (s *Server) Stop()                       {}
func (s *Server) Serve() <-chan error         { return nil }
