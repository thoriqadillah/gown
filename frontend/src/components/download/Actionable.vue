<script setup lang="ts">
import { ref } from 'vue';
import Downloader from '../../services/downloader';

const props = defineProps<{
  active: boolean,
  filename: string,
  id: string,
}>()

const downloader = Downloader.service()
const dialog = ref(false)
const clicked = ref(false)
function resumepause(id: string) {
  clicked.value = !clicked.value
  // clicked.value ? downloader.pause(id) : downloader.resume(id)
}
</script>

<template>
  <div class="tw-w-20">
    <div v-if="props.active" class="tw-flex">
      <!-- TODO: pause/resume implementation -->
      <v-btn @click="resumepause(props.id)" density="compact" variant="text" icon class="icon-name tw-opacity-0 tw-ml-5 -tw-mr-3">
        <v-icon :icon="clicked ? 'mdi-play' : 'mdi-pause'" class="tw-text-base tw-opacity-70"></v-icon>    
      </v-btn>

      <v-btn density="compact" variant="text" color="warning" icon class="icon-name tw-opacity-0 tw-ml-5 -tw-mr-3">
        <v-icon icon="mdi-stop" class="tw-text-base tw-opacity-70"></v-icon>

        <v-dialog v-model="dialog" activator="parent" max-width="450px">
          <v-card>
            <v-card-text>Are you sure you want to stop downloading {{ props.filename }}?</v-card-text>
            <div class="tw-flex tw-justify-between tw-flex-row-reverse">
              <div class="tw-flex tw-flex-row-reverse">
                <v-card-actions>
                  <v-btn color="warning" block @click="downloader.stop(props.id)">Stop</v-btn>
                </v-card-actions>
                <v-card-actions>
                  <v-btn variant="text" block @click="dialog = false">Cancel</v-btn>
                </v-card-actions>
              </div>
            </div>
          </v-card>
        </v-dialog>
      </v-btn>
    </div>
  </div>
</template>