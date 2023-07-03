<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { onMount } from 'svelte';
	import Note from '../components/Note.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

  onMount(async () => {
    await invalidateAll();
  })
</script>

<div class="flex justify-center items-center pt-16">
	<div class="max-w-7xl flex flex-wrap justify-center gap-3 px-3">
		{#each data.notes as note}
			<Note id={note.id} content={note.content} />
		{/each}
		<span
			class="inset-0 bg-black/30 fixed peer-hover:opacity-100 opacity-0 pointer-events-none transition-all"
		/>
	</div>
</div>
