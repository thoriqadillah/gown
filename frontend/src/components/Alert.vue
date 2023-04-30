<script setup lang="ts">
import { ref } from 'vue';
import { useAlert } from '../store/alert';

const alert = useAlert()
const timeout = ref(3000)
</script>

<template>
  <div class="text-center">
    <v-btn color="orange-darken-2" @click="alert.status = true">
      Open Snackbar
    </v-btn>

    <v-snackbar v-model="alert.status" :timeout="timeout" location="left bottom" color="red">
      {{ alert.message }}

      <template v-slot:actions>
        <v-btn color="red" variant="text" density="comfortable" @click="alert.status = false" icon="mdi-close"></v-btn>
      </template>
    </v-snackbar>
  </div>
</template>