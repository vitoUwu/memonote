<script lang="ts">
	import { enhance } from '$app/forms';
	import { Editor } from '@tiptap/core';
	import StarterKit from '@tiptap/starter-kit';
	import { Loader2, Save } from 'lucide-svelte';
	import { onDestroy, onMount } from 'svelte';
	import Button from './Button.svelte';
	import FormError from './form/FormError.svelte';

	let loading = false;
	let element: undefined | Element;
	let editor: undefined | Editor;
	export let clearOnSubmit: boolean = true;
	export let content = '';
	export let title = '';
	export let error = '';

	onMount(() => {
		editor = new Editor({
			element,
			autofocus: true,
			extensions: [StarterKit],
			content: '',
			onTransaction: () => {
				editor = editor;
			},
			onUpdate: ({ editor }) => {
				content = editor.getHTML()
			}
		});

		editor.commands.setContent(content);
	});

	onDestroy(() => {
		if (editor) {
			editor.destroy();
		}
	});
</script>

<form
	method="POST"
	class="w-full h-full flex flex-col gap-3 pb-16"
	use:enhance={({ cancel, formData }) => {
		if (!content.length) {
			error = 'Note content cannot be null';
			cancel();
			return;
		}

		if (!title.length) {
			error = 'Note title cannot be null';
			cancel();
			return;
		}

		loading = true;
		formData.set('title', title);
		formData.set('content', content);

		return async ({ update, result }) => {
			await update();
			loading = false;
			if (result.type === 'success') {
				if (clearOnSubmit) {
					editor?.commands.setContent('');
					title = '';
				}
				error = '';
			}
		};
	}}
>
	<input
		bind:value={title}
		class="w-full font-bold text-4xl bg-zinc-900 placeholder:text-zinc-700 placeholder:italic focus:outline-none"
		placeholder="Title..."
		type="text"
		name="title"
		id="title"
		maxlength="64"
		disabled={loading}
	/>
	<div bind:this={element} class="p-3 h-full w-full rounded border border-zinc-700 bg-zinc-800" />
	<div class="flex justify-end gap-3">
		{#if error.length}
			<FormError>{error}</FormError>
		{/if}
		<Button
			disabled={content.length === 0 || title.length === 0 || loading}
			type="submit"
			className="p-2 w-fit flex gap-1 text-zinc-400 disabled:text-zinc-600 transition-colors items-center justify-centers"
		>
			{#if loading}
				<Loader2 size={18} class="animate-spin" />
			{:else}
				<Save size={18} />
				<span>Save</span>
			{/if}
		</Button>
	</div>
</form>
