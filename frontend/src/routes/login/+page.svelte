<script lang="ts">
	import { enhance } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { ArrowRight, Loader2 } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { twMerge } from 'tailwind-merge';
	import Button from '../../components/Button.svelte';
	import TextInput from '../../components/TextInput.svelte';
	import FormError from '../../components/form/FormError.svelte';
	import type { ActionData } from './$types';

	let hidden = true;
	let loading = false;

	export let form: ActionData;

	function click() {
		hidden = !hidden;
	}

	onMount(async () => {
		await invalidateAll();
	});
</script>

<div class="flex w-screen h-screen justify-center items-center">
	<form
		use:enhance={() => {
			loading = true;
			if (form) {
				form.error = '';
			}
			return async ({ update }) => {
				await update();
				loading = false;
			};
		}}
		action="/login"
		method="POST"
		class="flex flex-col justify-center items-center gap-3 w-[80%] lg:w-[30%]"
	>
		{#if form && form.error !== ''}
			<FormError>{form.error}</FormError>
		{/if}
		<div class="flex flex-col gap-2 w-full">
			<label for="username">Username</label>
			<TextInput disabled={loading} type="text" id="username" name="username" required />
		</div>
		<div class="flex flex-col gap-2 w-full">
			<label for="password">Password</label>
			<TextInput
				disabled={loading}
				type={hidden ? 'password' : 'text'}
				name="password"
				id="password"
				required
			>
				<Button type="button" slot="right-button" onClick={click} className="relative">
					<p>👀</p>
					<span
						class={twMerge(
							'absolute left-[50%] -translate-x-1/2 scale-[2] rotate-90 transition-all',
							hidden ? 'top-[50%] -translate-y-1/2' : 'top-0 -translate-y-1/2 '
						)}>🤚</span
					>
				</Button>
			</TextInput>
		</div>
		<div class="w-full">
			<Button
				disabled={loading}
				type="submit"
				className="w-full font-bold text-lg flex justify-center"
			>
				{#if loading}
					<Loader2 class="animate-spin" />
				{:else}
					Login
				{/if}
			</Button>
		</div>
		<a
			href="/register"
			class="font-bold flex gap-2 items-center focus:outline-none focus:ring-2 focus:ring-blue-600 rounded text-blue-600"
		>
			Create an account
			<ArrowRight size={22} />
		</a>
	</form>
</div>
