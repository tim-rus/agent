<script setup lang="ts">
import { ref } from 'vue'
import { ping } from '../ping/ping'

const data = ref("waiting")

const click = async () => {
	const res = await ping()
	console.log(res);
	if (res && res.msg) {
		data.value = res.msg
	}
}
</script>

<template>
  <section id="center">
    <div>
      <h1>Test Connect RPC</h1>
    </div>
    <button type="button" class="counter" @click="click">
      Load
    </button>
    <div>{{data}}</div>
  </section>
</template>
