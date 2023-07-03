<script lang="ts">
	export let id: string;
	export let content: string;
</script>

<div
	class="peer cursor-pointer hover:scale-105 transition-all space-y-3 bg-zinc-900 max-w-[280px] overflow-hidden rounded border border-zinc-800 text-zinc-300 px-3 py-2 shadow-xl hover:z-10"
>
	<h1 class="font-bold">{id}</h1>
	<p class="overflow-clip whitespace-nowrap text-ellipsis text-zinc-500">{content}</p>
</div>
