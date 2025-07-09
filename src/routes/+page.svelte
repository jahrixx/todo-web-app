<script lang="ts">
    import { fade } from "svelte/transition";
    import { filteredTasks, filter } from "$lib/stores/tasks";
    import Header from "$lib/components/Header.svelte";
    import Task from "$lib/components/Task.svelte";

    let toggleNewTask = $state(false)
</script>

<Header bind:toggleNewTask />

<main class="main-app-container">
    {#if $filteredTasks.length === 0 && !toggleNewTask}
        <div class="task-lists" in:fade={{ delay: 200, duration: 150 }}>
           <span class="no-tasks">
                {#if $filter === 'all'}
                    No Tasks Available!
                {:else if $filter === 'active'}
                    No Active Tasks!
                {:else}
                    No Completed Tasks!
                {/if}
            </span>
        </div>
    {:else}
        <div class="task-grid">
            {#each $filteredTasks as task (task.id)}
                <Task {task} />
            {/each}
        </div>
    {/if}
</main>

<style>
    :global(body) {
        background-image: url('../lib/assets/wallpaper.jpg');
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Open Sans","Helvetica Neue",sans-serif;
    }
    .no-tasks {
        display: flex;
        justify-content: center;
        font-size: 1.8em;
        color: #495057;
        margin-bottom: 10px;
    }
    .task-grid {
        column-count: auto;
        column-width: 250px;
        column-gap: 1rem;
        padding: 0 1rem;
        grid-auto-flow: dense;
        margin-bottom: 1rem;
        /* display: grid;
        grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
        gap: 1rem;
        padding: 1rem;
        align-items: start;
        position: relative; */
    }
    .task-lists {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: .5rem;
    }
</style>

