<script lang="ts">
    import Icon from "@iconify/svelte";
    import BaseButton from "./BaseButton.svelte";
    import Task from "./Task.svelte";
    import { fade } from "svelte/transition";

    let { toggleNewTask = $bindable(false), ...props } : { [key: string]: any } = $props(); 
</script>

<header class="header">
    <h1 class="title">My Tasks:</h1>
    <BaseButton
        onclick={() => (toggleNewTask = !toggleNewTask)}
        aria-label="Add task"
    >
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
</header>

{#if toggleNewTask}
    <div class="new-task">
        <Task bind:isShown={toggleNewTask} task={undefined} />
    </div>
{/if}

<style>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: #ffffff;
    padding: 1rem 1.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
    margin-bottom: 1rem;
}

.title {
    font-size: 1.75rem;
    font-weight: 600;
    color: #2d3748;
}

.icon {
    width: 20px;
    height: 20px;
    color: white;
}

.new-task {
    background: #ffffff;
    padding: 1.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.06);
    animation: fadeIn 0.3s ease-out;
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
</style>