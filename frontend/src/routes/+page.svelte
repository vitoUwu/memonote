<script lang="ts">
	import { invalidate, invalidateAll } from '$app/navigation';
	import { onMount } from 'svelte';
	import Note from '../components/Note.svelte';
	import NoteEditor from '../components/NoteEditor.svelte';
	import type { ActionData, PageData } from './$types';

	export let form: ActionData;
	export let data: PageData;

	onMount(async () => {
		await invalidateAll();
	});

	async function handleDelete(id: string) {
		await fetch('/notes/delete', {
			method: 'POST',
			body: JSON.stringify({ id }),
			headers: {
				'Content-Type': 'application/json'
			}
		})
			.then((res) => res.json())
			.then((data) => {
				if (data.success === true) {
					invalidate('/');
				}

				console.error(data);
			})
			.catch(console.error);
	}
</script>

<div class="flex justify-center h-screen items-center overflow-y-auto">
	<div class="pt-16 max-w-7xl w-full h-full flex flex-wrap justify-center gap-3 px-3">
		{#key data.notes}
			{#if data.notes.length}
				{#each data.notes as note}
					<Note
						onDelete={() => handleDelete(note.id)}
						content={note.content}
						title={note.title}
						id={note.id}
					/>
				{/each}
				<span
					class="inset-0 bg-black/30 fixed peer-focus:opacity-100 peer-hover:opacity-100 opacity-0 pointer-events-none transition-all"
				/>
			{:else}
				<NoteEditor error={form?.error ?? ''} />
			{/if}
		{/key}
	</div>
</div>
