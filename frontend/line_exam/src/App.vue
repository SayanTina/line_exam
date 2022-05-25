<template>
  <div class="">
    <upload-component @result="updateResult" @selectedFile="selectedFile"></upload-component>
    <upload-progress-component v-if="showUpload"></upload-progress-component>
    <health-check-component v-if="showResult" :resultHealthCheck="resultHealthCheck"></health-check-component>
  </div>
</template>

<script>
import uploadComponent from "./components/uploadComponent.vue";
import uploadProgressComponent from "./components/uploadProgressComponent.vue";
import healthCheckComponent from "./components/healthCheckComponent.vue";

import './assets/tailwind.css'

export default {
  name: "App",
  components: {
    uploadComponent,
    uploadProgressComponent,
    healthCheckComponent,
  },
  data() {
    return {
      showUpload: false,
      showResult: false,
      resultHealthCheck: '',
    }
  },
  methods: {
    updateResult(value) {
      this.showUpload = false
      this.resultHealthCheck = value.data
      this.showResult = true

    },
    selectedFile() {
      this.showUpload = true
      this.showResult = false
    }
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
