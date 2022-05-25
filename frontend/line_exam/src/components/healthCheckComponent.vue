<template>
  <div class="container mx-auto w-1/3 box-border border-2 p-5 rounded-xl shadow-lg">
    <div class="box-border rounded-xl w-full py-5 font-bold text-gray-500">
      <div class="text-left text-3xl">Total {{ resultHealthCheck.service_amount}} Websites</div>
      <div class="text-left">(Used {{ ExecutionTime }})</div>
      <div class="flex place-content-evenly mt-4">
        <!-- Up status -->
        <div class="rounded-lg h-28 w-28 p-2 text-white" style="background-color:#3FE364;">
          <div class="text-left">Up</div>
          <div class="text-6xl">{{ resultHealthCheck.up_service_amount  }}</div>
        </div>
        <!-- Down status -->
        <div class="rounded-lg h-28 w-28 p-2" style="background-color:#F5F7FC;">
          <div class="text-left">Down</div>
          <div class="text-6xl">{{ resultHealthCheck.down_service_amount  }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import helper from '@/helpers'

export default {
  props: {
    resultHealthCheck: {}
  },
  computed: {
    ExecutionTime(){

    let seconds = helper.convertDurationToSecond(this.resultHealthCheck.execution_time)

    let dMins = seconds/60
    let dSecs = seconds%60


      return dMins >= 1 ? `${dMins} minute and ${dSecs.toFixed(2)} second` : `${dSecs.toFixed(2)} second`
    }
  }

}
</script>

<style>

</style>