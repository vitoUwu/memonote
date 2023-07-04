<script lang="ts">
	import { Trash } from 'lucide-svelte';
	import type { Action } from 'svelte/action';
	import Button from './Button.svelte';

	export let content: string;
	export let title: string;
	export let id: string;
	export let onDelete: () => void;
	let showContextMenu = false;
	let pos = { x: 0, y: 0 };
	let menu = { h: 0, w: 0 };
	let browser = { h: 0, w: 0 };

	const getContextMenuDimension: Action<HTMLElement> = (node) => {
		menu = {
			h: node.offsetHeight,
			w: node.offsetWidth
		};
		return {
			destroy() {
				menu = {
					h: 0,
					w: 0
				}
			}
		}
	};

	function onContextMenu(event: MouseEvent) {
		showContextMenu = true;
		browser = {
			w: window.innerWidth,
			h: window.innerHeight
		};
		pos = {
			x: event.clientX,
			y: event.clientY
		};
		if (browser.h - pos.y < menu.h) pos.y = pos.y - menu.h;
		if (browser.w - pos.x < menu.w) pos.x = pos.x - menu.w;
	}

	function onWindowClick() {
		showContextMenu = false;
	}
</script>

<a
	on:contextmenu|preventDefault={onContextMenu}
	href={`/notes/${id}`}
	class="focus:outline-none min-w-[150px] focus:z-10 focus:scale-105 focus:ring-2 focus:shadow-2xl focus:shadow-blue-600/10 focus:ring-blue-600 peer cursor-pointer hover:scale-105 transition-all space-y-3 bg-zinc-900 max-w-[280px] overflow-hidden rounded border border-zinc-800 text-zinc-300 px-3 py-2 shadow-xl hover:z-10 max-h-20 hover:shadow-blue-600/10 hover:shadow-2xl"
>
	<h1 class="font-bold">{title}</h1>
	<p class="overflow-clip whitespace-nowrap text-ellipsis text-zinc-500 select-none">
		{content.replaceAll(/<[^>]*>/gm, '')}
	</p>
</a>
{#if showContextMenu}
<nav use:getContextMenuDimension style="z-index: 11;position: absolute; top:{pos.y}px; left:{pos.x}px">
	<Button onClick={onDelete} className="flex gap-3 text-red-600 items-center font-bold">
		<Trash size={20}/>
		Delete
	</Button>
</nav>
{/if}

<svelte:window on:click={onWindowClick} />
