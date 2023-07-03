<script lang="ts">
	import { enhance } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { ArrowRight, Loader2 } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { twMerge } from 'tailwind-merge';
	import Button from '../../components/Button.svelte';
	import TextInput from '../../components/TextInput.svelte';
	import type { ActionData } from './$types';

	let hiddenPassword = true;
	let hiddenConfirmPassword = true;
	let loading = false;
	export let form: ActionData;

	function clickPassword() {
		hiddenPassword = !hiddenPassword;
	}

	function clickConfirmPassword() {
		hiddenConfirmPassword = !hiddenConfirmPassword;
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
				form.error = null;
			}
			return async ({ update }) => {
				await update();
				loading = false;
			};
		}}
		action="/register"
		method="POST"
		class="flex flex-col justify-center items-center gap-3 w-[80%] lg:w-[30%]"
	>
		{#if form && form.error}
			<p class="px-2 rounded bg-blue-600/20 text-blue-600 text-center">{form.error}</p>
		{/if}
		<div class="flex flex-col gap-2 w-full">
			<label for="username">Username</label>
			<TextInput disabled={loading} type="text" id="username" name="username" required />
		</div>
		<div class="flex flex-col gap-2 w-full">
			<label for="password">Password</label>
			<TextInput
				disabled={loading}
				type={hiddenPassword ? 'password' : 'text'}
				name="password"
				id="password"
				required
			>
				<Button type="button" slot="right-button" onClick={clickPassword} className="relative">
					<p>👀</p>
					<span
						class={twMerge(
							'absolute left-[50%] -translate-x-1/2 scale-[2] rotate-90 transition-all',
							hiddenPassword ? 'top-[50%] -translate-y-1/2' : 'top-0 -translate-y-1/2 '
						)}>🤚</span
					>
				</Button>
			</TextInput>
		</div>
		<div class="flex flex-col gap-2 w-full">
			<label for="password">Confirm Password</label>
			<TextInput
				disabled={loading}
				type={hiddenConfirmPassword ? 'password' : 'text'}
				name="confirm-password"
				id="confirm-password"
				required
			>
				<Button
					disabled={loading}
					type="button"
					slot="right-button"
					onClick={clickConfirmPassword}
					className="relative"
				>
					<p>👀</p>
					<span
						class={twMerge(
							'absolute left-[50%] -translate-x-1/2 scale-[2] rotate-90 transition-all',
							hiddenConfirmPassword ? 'top-[50%] -translate-y-1/2' : 'top-0 -translate-y-1/2 '
						)}>🤚</span
					>
				</Button>
			</TextInput>
		</div>
		<div class="w-full">
			<Button disabled={loading} type="submit" className="w-full font-bold text-lg flex justify-center">
				{#if loading}
					<Loader2 class="animate-spin" />
				{:else}
					Register
				{/if}
			</Button>
		</div>
		<a
			href="/login"
			class="font-bold flex gap-2 items-center focus:outline-none focus:ring-2 focus:ring-blue-600 rounded text-blue-600"
		>
			Already have an account?
			<ArrowRight size={22} />
		</a>
	</form>
</div>
