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

package payloadlimit

import (
	"context"

	"sofastack.io/sofa-mosn/pkg/api/v2"
	"sofastack.io/sofa-mosn/pkg/config"
	"sofastack.io/sofa-mosn/pkg/filter"
	"sofastack.io/sofa-mosn/pkg/log"
	"sofastack.io/sofa-mosn/pkg/types"
)

func init() {
	filter.RegisterStream(v2.PayloadLimit, CreatePayloadLimitFilterFactory)
}

type FilterConfigFactory struct {
	Config *v2.StreamPayloadLimit
}

func (f *FilterConfigFactory) CreateFilterChain(context context.Context, callbacks types.StreamFilterChainFactoryCallbacks) {
	filter := NewFilter(context, f.Config)
	callbacks.AddStreamReceiverFilter(filter, types.DownFilterAfterRoute)
}

func CreatePayloadLimitFilterFactory(conf map[string]interface{}) (types.StreamFilterChainFactory, error) {
	log.DefaultLogger.Debugf("create payload limit stream filter factory")
	cfg, err := config.ParseStreamPayloadLimitFilter(conf)
	if err != nil {
		return nil, err
	}
	return &FilterConfigFactory{cfg}, nil
}
