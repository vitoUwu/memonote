<script lang="ts">
	import '../app.css';
	import type { PageData } from './login/$types';

	export let data: PageData;
</script>

<nav class="fixed bg-zinc-900 shadow-xl flex justify-between px-10 h-12 items-center left-0 right-0 top-0">
	<div>
		<a class="font-bold text-xl focus:outline-none focus:ring-2 focus:ring-blue-600 rounded" href="/">Memonote</a>
	</div>
	<div class="flex gap-3 items-center max-w-[60%]">
		{#if data.user}
			<div
				class="h-7 w-7 flex justify-center items-center rounded-full bg-zinc-800 border border-zinc-700"
			>
				👨
			</div>
			<p class="overflow-hidden text-ellipsis">{data.user.username}</p>
			<a href="/logout" class="bg-zinc-800 border border-zinc-700 rounded px-2 py-1">Logout</a>
		{/if}
	</div>
</nav>
<main>
	<slot />
</main>

<style lang="postcss">
	:global(html) {
		background-color: theme(colors.zinc.900);
		color: theme(colors.zinc.300);
		-webkit-tap-highlight-color: rgba(0, 0, 0, 0);
		-webkit-tap-highlight-color: transparent;
	}
</style>
