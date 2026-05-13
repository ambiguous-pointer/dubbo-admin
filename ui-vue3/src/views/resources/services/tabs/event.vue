<!--
  ~ Licensed to the Apache Software Foundation (ASF) under one or more
  ~ contributor license agreements.  See the NOTICE file distributed with
  ~ this work for additional information regarding copyright ownership.
  ~ The ASF licenses this file to You under the Apache License, Version 2.0
  ~ (the "License"); you may not use this file except in compliance with
  ~ the License.  You may obtain a copy of the License at
  ~
  ~     http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
-->

<template>
  <EventTimeline :events="eventList" :loading="loading" />
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useMeshStore } from '@/stores/mesh'
import { listServiceEvent } from '@/api/service/service'
import EventTimeline from '@/components/EventTimeline.vue'
import type { EventItem } from '@/types/api'

const route = useRoute()
const meshStore = useMeshStore()

const eventList = ref<EventItem[]>([])
const loading = ref(false)

onMounted(async () => {
  loading.value = true
  try {
    const serviceName = (route.params.pathId as string) || ''
    const mesh = meshStore.mesh || 'default'
    const res = await listServiceEvent({ serviceName, mesh })
    if (res?.data?.list) {
      eventList.value = res.data.list
    }
  } finally {
    loading.value = false
  }
})
</script>
