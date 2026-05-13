/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"strings"

	consolectx "github.com/apache/dubbo-admin/pkg/console/context"
	"github.com/apache/dubbo-admin/pkg/console/model"
	"github.com/apache/dubbo-admin/pkg/core/manager"
	meshresource "github.com/apache/dubbo-admin/pkg/core/resource/apis/mesh/v1alpha1"
	coremodel "github.com/apache/dubbo-admin/pkg/core/resource/model"
	"github.com/apache/dubbo-admin/pkg/core/store/index"
)

func ListApplicationEvents(ctx consolectx.Context, req *model.EventQueryReq) (*model.EventListResp, error) {
	conditions := []index.IndexCondition{
		{IndexName: index.ByMeshIndex, Value: req.Mesh, Operator: index.Equals},
	}

	// Filter events whose involved object name has the app name as prefix,
	// matching K8s pod-name patterns like <appName>-<hash>-<hash>.
	if req.AppName != "" {
		conditions = append(conditions, index.IndexCondition{
			IndexName: index.ByK8sEventInvolvedObjName,
			Value:     req.AppName,
			Operator:  index.HasPrefix,
		})
	}

	pageData, err := manager.PageListByIndexes[*meshresource.K8sEventResource](
		ctx.ResourceManager(),
		meshresource.K8sEventKind,
		conditions,
		req.PageReq,
	)
	if err != nil {
		return nil, err
	}

	return toEventListResp(pageData), nil
}

func ListInstanceEvents(ctx consolectx.Context, req *model.EventQueryReq) (*model.EventListResp, error) {
	conditions := []index.IndexCondition{
		{IndexName: index.ByMeshIndex, Value: req.Mesh, Operator: index.Equals},
	}

	// Use only the primary identifier to avoid AND-intersection on the same index.
	// K8s events reference pods by name; registry events may use IP.
	// Prefer instanceName (pod name) over IP to avoid false AND mismatch.
	var ipFilter string
	if req.InstanceName != "" {
		conditions = append(conditions, index.IndexCondition{
			IndexName: index.ByK8sEventInvolvedObjName,
			Value:     req.InstanceName,
			Operator:  index.Equals,
		})
		ipFilter = req.InstanceIP
	} else if req.InstanceIP != "" {
		conditions = append(conditions, index.IndexCondition{
			IndexName: index.ByK8sEventInvolvedObjName,
			Value:     req.InstanceIP,
			Operator:  index.Equals,
		})
	}

	// When both name and IP are provided, use ListByIndexes so we can apply
	// OR semantics (match either name OR IP) via in-memory post-filter.
	if ipFilter != "" && req.InstanceName != "" {
		allConditions := []index.IndexCondition{
			{IndexName: index.ByMeshIndex, Value: req.Mesh, Operator: index.Equals},
		}
		resources, err := manager.ListByIndexes[*meshresource.K8sEventResource](
			ctx.ResourceManager(),
			meshresource.K8sEventKind,
			allConditions,
		)
		if err != nil {
			return nil, err
		}

		filtered := make([]*meshresource.K8sEventResource, 0)
		for _, r := range resources {
			if r.Spec == nil {
				continue
			}
			n := r.Spec.InvolvedObjName
			if n == req.InstanceName || n == ipFilter {
				filtered = append(filtered, r)
			}
		}

		// Apply manual pagination after in-memory filter
		offset := req.PageOffset
		end := offset + req.PageSize
		if offset > len(filtered) {
			offset = len(filtered)
		}
		if end > len(filtered) {
			end = len(filtered)
		}
		paged := filtered[offset:end]

		return toEventListResp(&coremodel.PageData[*meshresource.K8sEventResource]{
			Pagination: coremodel.Pagination{
				Total:      len(filtered),
				PageOffset: req.PageOffset,
				PageSize:   req.PageSize,
			},
			Data: paged,
		}), nil
	}

	pageData, err := manager.PageListByIndexes[*meshresource.K8sEventResource](
		ctx.ResourceManager(),
		meshresource.K8sEventKind,
		conditions,
		req.PageReq,
	)
	if err != nil {
		return nil, err
	}

	return toEventListResp(pageData), nil
}

func ListServiceEvents(ctx consolectx.Context, req *model.EventQueryReq) (*model.EventListResp, error) {
	conditions := []index.IndexCondition{
		{IndexName: index.ByMeshIndex, Value: req.Mesh, Operator: index.Equals},
	}

	if req.ServiceName != "" {
		conditions = append(conditions, index.IndexCondition{
			IndexName: index.ByK8sEventInvolvedObjName,
			Value:     req.ServiceName,
			Operator:  index.Equals,
		})
	}

	pageData, err := manager.PageListByIndexes[*meshresource.K8sEventResource](
		ctx.ResourceManager(),
		meshresource.K8sEventKind,
		conditions,
		req.PageReq,
	)
	if err != nil {
		return nil, err
	}

	return toEventListResp(pageData), nil
}

func toEventListResp(pageData *coremodel.PageData[*meshresource.K8sEventResource]) *model.EventListResp {
	items := pageData.Data
	list := make([]*model.EventItem, 0, len(items))
	for _, eventRes := range items {
		if eventRes.Spec == nil {
			continue
		}

		eventType := "normal"
		if strings.EqualFold(eventRes.Spec.Type, "Warning") {
			eventType = "warning"
		}

		source := eventRes.Spec.SourceComponent
		if source == "" {
			source = eventRes.Spec.EventSource
		}

		list = append(list, &model.EventItem{
			Time:    eventRes.Spec.LastTimestamp,
			Type:    eventType,
			Message: eventRes.Spec.Message,
			Source:  source,
		})
	}

	return &model.EventListResp{
		List:  list,
		Total: pageData.Total,
	}
}
