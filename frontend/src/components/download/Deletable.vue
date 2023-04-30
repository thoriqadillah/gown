<script setup lang="ts">
import { ref } from 'vue';
import { useDownloads } from '../../store/downloads';

const props = defineProps<{
  filename: string
  id: string
}>()

const downloads = useDownloads()
const dialog = ref(false)
const deleteFromdisk = ref(false)
</script>

<template>
  <v-btn color="red" density="compact" variant="text" icon>
    <v-icon icon="mdi-trash-can-outline" class="tw-text-base"></v-icon>

    <v-dialog v-model="dialog" activator="parent" max-width="450px">
      <v-card>
        <v-card-text>Delete file {{ props.filename }}?</v-card-text>
        <div class="tw-flex tw-justify-between">
          <v-card-actions>
            <v-checkbox v-model="deleteFromdisk" label="Delete from disk" color="red" value="true" hide-details/>
          </v-card-actions>
          <div class="tw-flex tw-flex-row-reverse">
            <v-card-actions>
              <v-btn color="red" block @click="downloads.remove(props.id, deleteFromdisk)">Delete</v-btn>
            </v-card-actions>
            <v-card-actions>
              <v-btn variant="text" block @click="dialog = false">Cancel</v-btn>
            </v-card-actions>
          </div>
        </div>
      </v-card>
    </v-dialog>
  </v-btn>
</template>