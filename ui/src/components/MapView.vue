<template>
  <div id="map" ref="map-root"></div>
</template>

<script>
import Map from 'ol/Map';
import View from 'ol/View';
import Feature from 'ol/Feature';
import { fromLonLat, transform, transformExtent } from 'ol/proj';
import { Point } from 'ol/geom';
import Tile from 'ol/layer/Tile';
import { Vector as VectorLayer } from 'ol/layer';
import { Style, Icon } from 'ol/style';
import OSM from 'ol/source/OSM';
import Vector from 'ol/source/Vector';
import { Zoom } from 'ol/control';

import 'ol/ol.css'

export default {
  name: 'MapContainer',
  components: {},
  props: {},
  data: () => ({
    mapView: null,
    vector: new Vector(),
  }),
  mounted() {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(({ coords: { longitude, latitude } }) => {
        this.loadMap(longitude, latitude);
      }, () => {
        this.loadMap(0, 0);
      });
    }
  },
  methods: {
    loadMap: function(viewLng, viewLat) {
      const mapRef = new Map({
        target: this.$refs['map-root'],
        layers: [
          new Tile({
            source: new OSM()
          }),
          new VectorLayer({
            source: this.vector,
            style: new Style({
              image: new Icon({
                anchor: [0.5, 0.5],
                anchorXUnits: "fraction",
                anchorYUnits: "fraction",
                src: "https://upload.wikimedia.org/wikipedia/commons/e/ec/RedDot.svg"
              })
            })
          })
        ],
        view: new View({
          zoom: 18,
          center: fromLonLat([viewLng, viewLat]),
          constrainResolution: true
        }),
        controls: [ new Zoom() ]
      });

      mapRef.on('moveend', ({ map }) => {
        this.vector.clear();
        const [startLong, startLat, endLong, endLat] = transformExtent(
          map.getView().calculateExtent(map.getSize()),
          'EPSG:3857',
          'EPSG:4326',
        );

        fetch(`http://localhost:8800/record/bounding-box/${startLong}/${startLat}/${endLong}/${endLat}`)
          .then(res => res.json())
          .then(data => data.forEach(coords => {
            const [lng, lat] = coords.geopoint;

            this.vector.addFeature(new Feature({
              geometry: new Point(transform([parseFloat(lng), parseFloat(lat)], 'EPSG:4326', 'EPSG:3857'))
            }));
          }));
      });

      this.mapView = mapRef;
    }
  }
}
</script>

<style scoped>
#map {
  width: 100%;
  height: 100%;
}
</style>
