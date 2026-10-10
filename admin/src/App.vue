<script lang="ts" setup>
import { ref } from 'vue';
import { ping } from './ping/ping';
import { Button } from './components/ui/button';
import { errAsync } from './lib/errors';
import { RouterView } from 'vue-router';

const data = ref("waiting")

const click = async () => {
	const [res, err] = await errAsync(ping())
	if (err) {
		data.value = `error: ${err.message}`
		return
	}
	if (res && res.msg) {
		data.value = res.msg
	}
}

</script>

<!--  -->

<template>
	<div class="bg-amber-300">Hello from cp with tailwind and shadcn</div>
	<Button @click="click">Click me</Button>
	<div class="bg-green-300">{{ data }}</div>

	<div class="pt-8">
	    <RouterView />
	</div>
</template>
