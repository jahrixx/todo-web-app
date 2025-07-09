<script lang="ts">
   import type { Snippet } from "svelte";

   let { label, value, current, onclick, children, ...props }: { label: string, value: string, current: string, onclick: () => void, children: Snippet, [key: string]: any} = $props();
</script>

<a class:active={current === value} onclick={onclick} {...props}>{@render children()}</a>

<style>
a {
    position: relative;
    background: none;
    padding: 0;
    cursor: pointer;
    margin-left: .75rem;
    margin-right: .25rem;
    font-size: 1.15rem;
    font-weight: 600;
    color: #1d4ed8;
    text-decoration: none;
    transition: color 0.5s ease;
}

a::after {
    content: "";
    position: absolute;
    left: 0;
    bottom: -2px;
    width: 100%;
    height: 2.35px;
    background-color: #495057;
    transform: scaleX(0);
    transform-origin: left;
    transition: transform 0.5s ease;
}

a:hover::after,
a.active::after {
    transform: scaleX(1);
}
</style>