<script lang="ts">
   import type { Snippet } from "svelte";
   let { children, ...props }: { children: Snippet, [key: string]: any} = $props();
</script>

<button class="base-btn" {...props}>{@render children()}</button>

<style>
.base-btn {
    background-color: #2563eb;
    color: white;
    border: none;
    border-radius: 0.5rem;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    font-weight: 500;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.2s ease;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.05);
}

.base-btn:hover {
    background-color: #1d4ed8;
}

.base-btn:active {
    transform: scale(0.98);
}

.base-btn:disabled {
    background-color: #cbd5e1;
    color: #94a3b8;
    cursor: not-allowed;
}
</style>