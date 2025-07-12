<script lang="ts">
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import { token } from "$lib/stores/auth";
  import { filteredTasks, filter, searchQuery, fetchAllTasks } from "$lib/stores/tasksStore";
  import { get } from "svelte/store";
  import { loading } from "$lib/stores/tasksStore"
  import Header from "$lib/components/Header.svelte";
  import Task from "$lib/components/Task.svelte";

  let toggleNewTask = $state(false);

  onMount(async () => {
    const userToken = get(token);
    if (!userToken){
      window.location.href = '/login';
    } else {
      await fetchAllTasks();
    }
  });

</script>

{#if $token}
  <Header bind:toggleNewTask {loading} />
  <main class="main-app-container">
    {#if $loading}
      <div class="task-lists" in:fade={{ delay: 150 }}>
        <span class="no-tasks">Loading your notes...</span>
      </div>

    {:else if $filteredTasks.length === 0 && !toggleNewTask}
      <div class="task-lists" in:fade={{ delay: 200, duration: 150 }}>
        <span class="no-tasks">
          {#if $filter === 'all'}
            {$searchQuery ? 'No matching notes found' : 'No Notes Created Yet!'}
          {:else if $filter === 'active'}
            {$searchQuery ? 'No matching active notes' : 'No Active Notes!'}
          {:else}
            {$searchQuery ? 'No matching completed notes' : 'No Completed Notes!'}
          {/if}
        </span>
      </div>

    {:else if $filteredTasks.length > 0 && !toggleNewTask}
      <div class="task-grid" in:fade={{ delay: 200 }}>
        {#each $filteredTasks as task (task.id)}
          <Task {task} />
        {/each}
      </div>

    {/if}
  </main>
{/if}

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
        font-size: 2rem;
        font-weight: 600;
        font-style: italic;
        color: #495057;
        margin-top: 20%;
        margin-bottom: 10px;
    }
    .task-grid {
        column-count: auto;
        column-width: 250px;
        column-gap: 1rem;
        padding: 0 1rem;
        grid-auto-flow: dense;
        margin-bottom: 1rem;
        margin-top: 6.5rem;
    }
    .task-lists {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: .5rem;
    }
</style>

