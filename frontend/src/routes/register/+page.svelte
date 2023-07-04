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

	const invalidUsername = /[^a-zA-Z_]/g;
	const validPassword = /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/g;

	let hiddenPassword = true;
	let hiddenConfirmPassword = true;
	let loading = false;
	let username = '';
	let password = '';
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
			<FormError>{form.error}</FormError>
		{/if}
		<div class="flex flex-col gap-2 w-full">
			<label for="username">Username</label>
			<TextInput
				onChange={(value) => (username = value)}
				disabled={loading}
				type="text"
				id="username"
				name="username"
				required
			/>
			<!-- {#if username.length}
				{#if invalidUsername.exec(username)}
					<FormError>
						Your username contains invalid characters. Use only uppercase and lowercase letters or
						"_"
					</FormError>
				{/if}
				{#if username.length > 15 || username.length < 2}
					<FormError>Username length must be between 2 and 15</FormError>
				{/if}
			{/if} -->
		</div>
		<div class="flex flex-col gap-2 w-full">
			<label for="password">Password</label>
			<TextInput
				onChange={(value) => (password = value)}
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
							'absolute left-[50%] -translate-x-1/2 scale-[2] rotate-90 transition-all -translate-y-1/2',
							hiddenPassword ? 'top-[50%]' : 'top-0'
						)}>🤚</span
					>
				</Button>
			</TextInput>
			<!-- {#if password.length}
				{#if validPassword.exec(password) === null}
					<FormError>
						Your password must includes minimum 8 characters, one upper letter, one lower letter and
						one number
					</FormError>
				{/if}
			{/if} -->
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
			<Button
				disabled={loading}
				type="submit"
				className="w-full font-bold text-lg flex justify-center"
			>
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
