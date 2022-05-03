<template>
  <div v-if="isCameraVisible" class="video-modal">
    <video ref="videoStream" autoplay></video>
  </div>
  <div class="camera-button-container">
    <button @click="touch" class="camera-button" />
  </div>
</template>

<script>
export default {
  name: 'CameraButton',
  data: () => ({
    isCameraVisible: false,
  }),
  methods: {
    touch: async function() {
      this.isCameraVisible = true;

      const videoStream = await navigator.mediaDevices.getUserMedia({
        audio: true,
        video: true
      });
      this.$refs.videoStream.srcObject = videoStream;
      this.$refs.videoStream.play();
    }
  }
}
</script>


<style scoped>
.video-modal {
  position: fixed;
  height: 100vh;
  width: 100vw;

  left: 0;
  top: 0;
  z-index: 99999;
  background: #000;
}

.video-modal>video {
  height: 100%;
  width: 100%;
}

.camera-button-container {
  width: 100%;
  position: absolute;
  bottom: 1em;

  text-align: center;
  pointer-events: none;
}

.camera-button {
  pointer-events: all;
  border-radius: 100%;
  border: 2px solid #000b;
  padding: 2rem;

  background-color: #fffe;
}

.camera-button:hover {
  cursor: pointer;
}
</style>
