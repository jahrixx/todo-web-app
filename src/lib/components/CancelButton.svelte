<script lang="ts">
   import type { Snippet } from "svelte";
   let { children, ...props }: { children: Snippet, [key: string]: any} = $props();
</script>

<button class="cancel-btn glass" {...props}>{@render children()}</button>

<style>
.glass {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}
.cancel-btn {
    color: #D22B2B;
    border: 1px solid #D22B2B;
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
}

.cancel-btn:hover {
    background-color: #D22B2B;
    color: white;
}

.cancel-btn:active {
    transform: scale(0.98);
}

.cancel-btn:disabled {
    background-color: #cbd5e1;
    color: #94a3b8;
    cursor: not-allowed;
}
</style>