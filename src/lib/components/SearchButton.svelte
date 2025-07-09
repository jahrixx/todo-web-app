<script lang="ts">
    import { onMount } from "svelte";
    import Icon from "@iconify/svelte";
    import { searchQuery } from "$lib/stores/tasks";

    let { 
        toggleSearch = $bindable(false), 
        placeholder = 'Search...', 
        searchInput = $bindable(null as HTMLInputElement | null), 
        ...props }: { 
            toggleSearch?: boolean; 
            placeholder?: string; 
            searchInput?: HTMLInputElement | null; 
            [key: string]: any; } = $props();

    let showSearch = $state(false)
    let localQuery = $state('')
    
    $effect(() => {
        searchQuery.set(localQuery);    
    });

    $effect(() => {
        if (toggleSearch) {
            setTimeout(() => {
                showSearch = true;
                searchInput?.focus();
            }, 0);
        } else {
            showSearch = false;
            localQuery = '';
        }
    });

    const handleClickOutside = (event: MouseEvent) => {
        if (showSearch && searchInput && !event.composedPath().includes(searchInput)) {
            showSearch = false;
            toggleSearch = false;
        }
    };

    onMount(() => {
        window.addEventListener('click', handleClickOutside);
        return () => {
            window.removeEventListener('click', handleClickOutside);
        };
    });
</script>

<div class="search-wrapper" {...props}>
    <button class="search-icon" onclick={() => toggleSearch = !toggleSearch} aria-labelledby="search">
        <Icon icon="fa6-solid:magnifying-glass" style="font-size: 1.2rem; margin-top: .25rem;" />
    </button>
    
    <input 
        bind:this={searchInput} 
        bind:value={localQuery}
        class:visible={showSearch} 
        {placeholder} 
        onkeydown={(e) => e.key === 'Escape' && (showSearch = toggleSearch = false)} />
</div>

<style>
    .search-wrapper {
        position: relative;
        display: flex;
        align-items: center;
        gap: .35rem;
        margin-right: .35rem;
    }

    .search-icon {
        background: transparent;
        border: none;
        cursor: pointer;
        padding: .5rem;
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 100;
        /* outline: 2px solid black; */
    }

    input {
        position: absolute;
        right: 0;
        width: 0;
        opacity: 0;
        border: none;
        border-bottom: 2.35px solid #495057;
        padding: .5rem;
        padding-bottom: .1rem;
        margin-right: 1.25rem;
        margin-bottom: 2px;
        font-size: 1.1rem;
        background-color: transparent;
        outline: none;
        transition: 
            width 0.3s ease-in-out,
            opacity 0.5s ease-in-out,
            transform 0.3s ease-in-out;
        transform: translateX(100%); 
    }

    input.visible {
        width: 200px;
        opacity: 1;
        transform: translateX(0);
    }
</style>