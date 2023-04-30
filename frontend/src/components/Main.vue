<script setup lang="ts">
import { provide, ref } from 'vue';

const drawer = ref(false)
provide('drawer', drawer)

const reload = () => location.reload()

</script>

<template>
  <v-card>
    <v-layout >
      <v-navigation-drawer v-model="drawer" :elevation="0" expand-on-hover rail :border="0" location="right">
        <v-list density="compact" nav class="tw-flex tw-flex-col tw-mt-12">
          <v-list-item active-color="primary" prepend-icon="mdi-brightness-4" title="Theme" value="theme"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-bug-outline" title="Log" value="log"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-replay" title="Refresh" value="refresh" @click="reload()"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-cog-outline" title="Settings" value="settings"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-information-outline" title="About" value="about"></v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main class="tw-overflow-scroll tw-h-screen tw-w-fit">
        <slot/>
      </v-main>
    </v-layout>
  </v-card>
</template>