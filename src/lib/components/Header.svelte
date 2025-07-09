<script lang="ts">
    import Icon from "@iconify/svelte";
    import BaseButton from "./BaseButton.svelte";
    import Task from "./Task.svelte";
    import { fade } from "svelte/transition";
    import { filter } from "$lib/stores/tasks";

    let { toggleNewTask = $bindable(false), ...props } : { [key: string]: any } = $props(); 
    let showForm = false;

    $effect(() => {
        if (toggleNewTask) {
            setTimeout(() => {
                showForm = true;
            }, 0);
        } else {
            showForm = false;
        }
    });
</script>

<header class="header glass">
    <div class="header-title">
        <h1 class="title">My Task Planner</h1>
        <BaseButton onclick={() => (toggleNewTask = !toggleNewTask)} aria-label="Add task">
            {#if toggleNewTask}
            <div class="fade-in icon" in:fade={{ duration: 150 }}>
                <Icon icon="fa6-solid:xmark" />
            </div>
            {:else}
            <div class="fade-in icon" in:fade={{ duration: 150 }}>
                <Icon icon="fa6-solid:plus" />
            </div>
            {/if}
        </BaseButton>
    </div>
    <div class="filter-controls">
        <button class="glass" class:active={$filter === 'all'} onclick={() => filter.set('all')}>All</button>
        <button class="glass" class:active={$filter === 'active'} onclick={() => filter.set('active')}>Active</button>
        <button class="glass" class:active={$filter === 'completed'} onclick={() => filter.set('completed')}>Completed</button>
    </div>    
</header>

{#if toggleNewTask}
    <div
        class="overlay"
        role="button"
        tabindex="0"
        aria-label="Close modal"
        onclick={() => (toggleNewTask = false)}
        onkeydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { toggleNewTask = false; } }}
    >
        <div
            class="modal"
            role="dialog"
            tabindex="0"
            aria-modal="true"
            onclick={(e) => e.stopPropagation()}
            onkeydown={(e) => e.stopPropagation()}
        >
            <Task bind:isShown={toggleNewTask} task={undefined} />
        </div>
    </div>
{/if}

<style>
.glass {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: .5rem;
    margin-bottom: 1rem;
}

.header-title {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.title {
    margin-top: 15px;
    font-size: 2rem;
    font-weight: 700;
    color: #2d3748;
    opacity: 0.9;
}

.icon {
    width: 20px;
    height: 20px;
    color: white;
    font-size: 1.2rem;
}

/* Animation */
@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(-8px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.fade-in {
    animation: fadeIn 0.3s ease-out;
}

.filter-controls {
    display: flex;
    gap: 0.5rem;
    justify-content: center;
}

.filter-controls button {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin: 0;
    font-size: 14px;
    opacity: 0.9;
}

.filter-controls button.active {
    background: #007acc;
    color: white;
    border-color: #007acc;
}

.overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.75);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    animation: fadeIn 0.25s ease-out;
}

.modal {
    padding: 2rem;
    border-radius: 16px;
    min-width: 320px;
    max-width: 600px;
    width: 90%;
    animation: popIn 0.25s ease-out;
}
</style>