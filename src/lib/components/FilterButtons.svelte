<script lang="ts">
   import type { Snippet } from "svelte";

   let { label, value, current, onclick, children, ...props }: { label: string, value: string, current: string, onclick: () => void, children: Snippet, [key: string]: any} = $props();
</script>

<button class="glass" class:active={current === value} onclick={onclick} {...props}>{@render children()}</button>

<style>
.glass {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}
button {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin: 0;
    font-size: 14px;
}

button.active {
    background: #2563eb;
    color: white;
    border-color: #2563eb;
}
</style>