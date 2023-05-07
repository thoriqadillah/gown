<script setup lang="ts">
import Downloader from '../../services/downloader';
import { ref } from 'vue';

const props = defineProps<{
  active: boolean,
  filename: string,
  statusname: string,
  id: string,
}>()

const downloader = Downloader.service()
const dialog = ref(false)
const disableStop = ref(props.statusname === 'Canceled')

async function stop(id: string) {
  dialog.value = false
  downloader.stop(id)
}
</script>

<template>
  <div class="tw-w-20">
    <div v-if="props.active" class="tw-flex">
      <v-btn v-if="props.statusname == 'Canceled'" @click="downloader.restart(props.id)" density="compact" variant="text" icon class="icon-name tw-opacity-0 tw-ml-5 -tw-mr-3">
        <v-icon class="tw-text-base tw-opacity-70" icon="mdi-replay"></v-icon>    
      </v-btn>
      <v-btn v-else-if="props.statusname == 'Paused'" @click="downloader.resume(props.id)" density="compact" variant="text" icon class="icon-name tw-opacity-0 tw-ml-5 -tw-mr-3">
        <v-icon class="tw-text-base tw-opacity-70" icon="mdi-play"></v-icon>    
      </v-btn>
      <v-btn v-else @click="downloader.pause(props.id)" density="compact" variant="text" icon class="icon-name tw-opacity-0 tw-ml-5 -tw-mr-3">
        <v-icon class="tw-text-base tw-opacity-70" icon="mdi-pause"></v-icon>    
        <!-- TODO: add dialog if the download is not resumable -->
      </v-btn>

      <v-btn density="compact" :disabled="props.statusname === 'Canceled' || disableStop"  :variant="props.statusname === 'Canceled' || disableStop ? 'plain' : 'text'" color="warning" icon class="icon-name tw-opacity-0 tw-ml-5 -tw-mr-3">
        <v-icon icon="mdi-stop" class="tw-text-base tw-opacity-70"></v-icon>

        <v-dialog v-model="dialog" activator="parent" max-width="450px">
          <v-card>
            <v-card-text>Are you sure you want to stop downloading {{ props.filename }}?</v-card-text>
            <div class="tw-flex tw-flex-row-reverse">
              <v-card-actions>
                <v-btn color="warning" block @click="stop(props.id)">Stop</v-btn>
              </v-card-actions>
              <v-card-actions>
                <v-btn variant="text" block @click="dialog = false">Cancel</v-btn>
              </v-card-actions>
            </div>
          </v-card>
        </v-dialog>
      </v-btn>
    </div>
  </div>
</template>