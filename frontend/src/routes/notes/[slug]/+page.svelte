<script lang="ts">
	import NoteEditor from '../../../components/NoteEditor.svelte';
	import type { ActionData, PageData } from './$types';

	export let form: ActionData;
	export let data: PageData;
</script>

<div class="flex justify-center h-screen items-center overflow-y-auto">
	<div class="pt-16 max-w-7xl w-full h-full flex flex-wrap justify-center gap-3 px-3">
		{#key form?.success}
			{#if data.note}
				<NoteEditor error={form?.error ?? ''} clearOnSubmit={false} title={data.note.title} content={data.note.content} />
			{:else if data.error}
				<h1>{data.error}</h1>
			{/if}
		{/key}
	</div>
</div>
