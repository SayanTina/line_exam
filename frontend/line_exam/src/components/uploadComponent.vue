<template>
  <div class="container mx-auto my-5 w-1/3">
    <div class="text-2xl font-bold text-left py-2">Websites Checker</div>
    <div
      class="
        container
        mx-auto
        box-border
        border-dashed border-4
        rounded-xl
        w-full
        py-5
        box-upload
        text-gray-500
      "
      @dragover.prevent="dragOver"
      @dragleave.prevent="dragLeave"
      @drop.prevent="dropFile($event)"
    >
      <div
        class="
          bg-[url('./../../public/csv.png')] bg-no-repeat bg-center
          w-28
          h-28
          mx-auto
          bg-[length:100px_100px]
        "
      ></div>
      <div class="text-2xl font-bold py-2">
        Drag your .csv file here to start uploading
      </div>
      <div class="w-2/5 mx-auto">
        <div class="relative flex py-5 items-center">
          <div class="flex-grow border-t-4 border-gray-400"></div>
          <span class="flex-shrink mx-4 text-2xl font-bold text-gray-400"
            >OR</span
          >
          <div class="flex-grow border-t-4 border-gray-400"></div>
        </div>
        <div
          class="rounded-lg text-white text-2xl font-bold py-4 cursor-pointer"
          style="background-color: #6749f5"
          @click="$refs.file.click()"
        >
          Browse File
        </div>
        <input
          type="file"
          id="csv_file"
          name="csv_file"
          ref="file"
          style="display: none"
          accept=".csv"
          hidden
          @change="requestUploadFile"
        />
      </div>
    </div>
  </div>
</template>

<script>
import healthCheckService from '@/services/HealthCheckService';
export default {
  name: "upload-component",
  data() {
    return {
      isDragging: false,
      file: "",
    };
  },
  methods: {
    dragOver() {
      this.isDragging = true;
    },
    dragLevave() {
      this.isDragging = false;
    },

    async uploadFile() {
      this.$emit('selectedFile')
      let result = await healthCheckService.healthCheckUpload(this.file)
      this.$emit('result', result)
    },

    requestUploadFile() {
      this.file = this.$refs.file.files[0]
      this.uploadFile()
    },

    dropFile(e) {
      if(e.dataTransfer.files[0].type == "text/csv") {
        this.file = e.dataTransfer.files[0];
        this.uploadFile()
      } else { 
        alert("file type not match")
      }
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.box-upload {
  background: #f5f7fc;
  border-color: #c7cddb;
}
</style>
