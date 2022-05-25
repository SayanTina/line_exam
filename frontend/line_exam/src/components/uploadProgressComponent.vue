<template>
  <div class="container mx-auto w-1/3 box-border border-2 rounded-xl shadow-lg">
    <div class="container w-full p-5 flex">
      <div
        class="
          flex-initial
          bg-[url('./../../public/csv.png')] bg-no-repeat bg-center
          w-24
          h-24
          bg-[length:80px_80px]
        "
      ></div>
      <div class="flex-grow px-2 flex flex-col justify-center">
        <div class="flex justify-between text-xl text-gray-400">
          <div>{{ fileName }}</div>
          <div class="font-bold ">{{ uploadProgress }}%</div>
        </div>
        <div class="w-full bg-gray-200 rounded-full h-5 dark:bg-gray-700">
          <div
            class="h-5 rounded-full"
            style="width: 50%; background-color: #6749f5"
            :style="{ width: uploadProgress + '%'}"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    uploadProgress() {
      return this.$store.state.uploadProgress || 0
    },
    fileName() {
      return this.$store.state.fileName
    },
  }
};
</script>

<style>
</style>