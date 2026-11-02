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
  <div
    class="border rounded-lg p-4 hover:bg-gray-50 cursor-pointer transition-all duration-200 hover:shadow-md"
    @click="handleClick"
  >
    <div class="flex items-center gap-2 mb-2">
      <div :class="suggestion.iconColor">{{ suggestion.icon }}</div>
      <div class="font-medium">{{ suggestion.title }}</div>
    </div>
    <div class="text-gray-500 text-sm">{{ suggestion.content }}</div>
  </div>
</template>

<script setup lang="ts">
import type { ChatSuggestion } from './types'

// 定义props
const props = defineProps<{
  suggestion: ChatSuggestion
}>()

// 定义emits
const emit = defineEmits<{
  (e: 'click', suggestion: string): void
}>()

// 处理点击事件
const handleClick = () => {
  emit('click', props.suggestion.content)
}
</script>
