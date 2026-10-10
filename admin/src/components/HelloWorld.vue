<script setup lang="ts">
import { ref } from 'vue'
import { ping } from '../ping/ping'
import { Button } from './ui/button'

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
    <Button @click="click">Load</Button>
    <div>{{data}}</div>
  </section>
</template>
