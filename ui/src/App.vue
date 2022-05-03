<template>
  <div class="map-view">
    <MapView />
    <CameraButton />
  </div>
</template>

<script>
import MapView from './components/MapView.vue'
import CameraButton from './components/CameraButton.vue';

export default {
  name: 'App',
  components: {
    MapView,
    CameraButton,
  }
}
</script>

<style>
body {
  margin: 0;
}

.map-view {
  width: 100vw;
  height: 100vh;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
}

.ol-zoom {
  /* Override default values */
  top: unset;
  left: unset;

  bottom: .5em;
  right: .5em;
}
</style>
