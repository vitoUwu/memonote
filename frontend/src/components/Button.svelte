<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';
	import { twMerge } from 'tailwind-merge';

	interface $$Props extends Omit<HTMLButtonAttributes, 'class'> {
		onClick?: () => void;
		className?: string;
	}

	export let onClick = () => {};
	export let className: string = '';
</script>

<button
	{...$$restProps}
	on:click={onClick}
	class={twMerge(
		'bg-zinc-800 border-zinc-700 border rounded-md px-4 py-3 focus:outline-none focus:ring-2 focus:ring-blue-600',
		'disabled:bg-zinc-900 disabled:border-zinc-800 disabled:cursor-not-allowed',
		className
	)}
>
	<slot />
</button>
