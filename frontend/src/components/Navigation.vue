<script setup lang="ts">
import { useTheme } from 'vuetify/lib/framework.mjs';
import { ref } from 'vue';

const theme = useTheme()
const themeIcon = ref(theme.global.current.value.dark ? 'mdi-white-balance-sunny' : 'mdi-weather-night')
const themeTooltip = ref(theme.global.current.value.dark ? 'Light Mode' : 'Dark Mode')
function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  themeIcon.value = theme.global.current.value.dark ? 'mdi-white-balance-sunny' : 'mdi-weather-night'
  themeTooltip.value = theme.global.current.value.dark ? 'Light Mode' : 'Dark Mode'
}

let loading = ref(false)
let loaded = ref(false)
function onClick() {
  loading.value = true

  setTimeout(() => {
    loading.value = false
    loaded.value = true
  }, 2000)

  loaded.value = false
}
</script>

<template>
  <nav class="tw-mx-5 tw-my-5 tw-flex tw-justify-between tw-gap-10 tw-items-center">
    <div class="tw-flex tw-gap-5">
      <div class="tw-w-28 dark:tw-bg-neutral-700 tw-text-center tw-rounded-sm"></div>

      <v-tooltip text="Setting" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" icon="mdi-cog-outline" variant="flat"/>
        </template>
      </v-tooltip>
      <v-tooltip :text="themeTooltip" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" :icon="themeIcon" @click="toggleTheme" variant="flat"/>
        </template>
      </v-tooltip>
    </div>
    <v-text-field :loading="loading" density="compact" variant="solo" label="Search" append-inner-icon="mdi-magnify" single-line hide-details @click:append-inner="onClick" v-on:keyup.enter="onClick"/>
    <div class="tw-flex tw-gap-5">
      <v-tooltip text="Resume All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" icon="mdi-play" variant="flat"/>
        </template>
      </v-tooltip>
      <v-tooltip text="Pause All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" icon="mdi-pause" variant="flat"/>
        </template>
      </v-tooltip>
      <v-tooltip text="Stop All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" icon="mdi-stop" variant="flat"/>
        </template>
      </v-tooltip>
      <v-divider vertical></v-divider>
      
      <v-tooltip text="Add into Queue" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" icon="mdi-tray-plus" variant="flat"/>
        </template>
      </v-tooltip>
      <v-tooltip text="New" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="compact" icon="mdi-plus" variant="flat"/>
        </template>
      </v-tooltip>
    </div>
  </nav>
</template>