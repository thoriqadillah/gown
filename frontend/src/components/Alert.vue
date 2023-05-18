<script setup lang="ts">
import { useAlert } from '../store/alert';

const alert = useAlert()
</script>

<template>
  <v-snackbar v-model="alert.status" :timeout="3000" location="left bottom" :color="alert.type === 'danger' ? 'red' : alert.type">
    {{ alert.message }}
    <template v-slot:actions>
      <v-btn variant="text" density="comfortable" @click="alert.status = false" icon="mdi-close"></v-btn>
    </template>
  </v-snackbar>
</template>