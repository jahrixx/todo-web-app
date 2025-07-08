<script lang="ts">
    import { fade } from "svelte/transition";
    import { tasks } from "$lib/stores/tasks";
    import Header from "$lib/components/Header.svelte";
    import Task from "$lib/components/Task.svelte";

    let toggleNewTask = $state(false)
</script>

<Header bind:toggleNewTask />

<main class="main-app-container">
    {#if $tasks.length === 0 && !toggleNewTask}
        <div class="task-lists" in:fade={{ delay: 200, duration: 150 }}>
            <span class="no-tasks">No Tasks Available!</span>
        </div>
    {:else}
        {#each $tasks as task (task.id)}
            <Task {task} />
        {/each}
    {/if}
</main>

<style lang="scss">
    :global(html, body) {
        height: 100%;
        width: 100%;
    }
    :global(body) {
        background-image: url('../lib/assets/wallpaper.jpg');
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        font-family: system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Open Sans","Helvetica Neue",sans-serif;
        @media screen and (min-width: 400px) {
            max-width: 400px;
            margin: auto;
        }
    }
    .no-tasks {
        display: flex;
        justify-content: center;
        font-size: 1.8em;
        color: #495057;
        margin-bottom: 10px;
    }
</style>

