export type Download = {
  status: {
    name: 'success' | 'failed' | 'paused' | 'queued',
    icon: "mdi-check-circle-outline" | 'mdi-alert-outline' | 'mdi-pause-circle-outline' | 'mdi-tray-full',
    color: 'success' | 'warning' | '' | 'info',
  },
  name: string,
  timeElapsed: number,
  size: number,
  date: string,
  type: {
    name: 'document' | 'video' | 'music' | 'compressed' | 'image',
    icon: 'mdi-file-document' | 'mdi-video' | 'mdi-music-box' | 'mdi-zip-box' | 'mdi-image',
    color: string
  }
}