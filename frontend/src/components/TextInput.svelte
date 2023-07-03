<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';

	interface $$Props extends HTMLInputAttributes {}
</script>

<div class="flex gap-3 justify-between">
	<input
		{...$$restProps}
		class="w-full px-4 py-3 text-zinc-200 bg-zinc-800 border border-zinc-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600 disabled:bg-zinc-900 disabled:border-zinc-800 disabled:cursor-not-allowed"
	/>
	<slot name="right-button" />
</div>
