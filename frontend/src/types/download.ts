export type Download = {
  id: string,
  status: {
    name: 'success' | 'failed' | 'paused' | 'queued' | 'processing',
    icon: "mdi-check-circle-outline" | 'mdi-alert-outline' | 'mdi-pause-circle-outline' | 'mdi-tray-full' | '',
    color: 'success' | 'warning' | '' | 'info',
  },
  name: string,
  timeElapsed: number,
  size: number,
  date: string,
  type: {
    name: 'document' | 'video' | 'audio' | 'compressed' | 'image' | 'other',
    icon: 'mdi-file-document' | 'mdi-video' | 'mdi-music-box' | 'mdi-zip-box' | 'mdi-image' | 'mdi-file-question',
    color: 'blue-accent-2' | 'deep-orange-accent-2' | 'purple-accent-2' | 'yellow-accent-4' | 'red-accent-2' | ''
  }
}