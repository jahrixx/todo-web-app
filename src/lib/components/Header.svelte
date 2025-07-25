<script lang="ts">
  import { fade } from "svelte/transition";
  import { filter } from "$lib/stores/tasksStore";
  import { token } from "$lib/stores/auth";
  import Icon from "@iconify/svelte";
  import Logo from "./Logo.svelte";
  import Task from "./Task.svelte";
  import LogoutButton from "./LogoutButton.svelte";
  import SearchButton from "./SearchButton.svelte";
  import BaseButton from "./BaseButton.svelte";
  import FilterButtons from "./FilterButtons.svelte";
  import LoginButton from "./LoginButton.svelte";
  
  let { toggleNewTask = $bindable(false), ...props }: { [key: string]: any } = $props();
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

<header class="header glass sticky">
  <div class="header-title">
    <Logo />
    {#if !$token}
      <BaseButton style="cursor: default">
        <div class="fade-in icon" style="cursor: default">
          <Icon icon="fa6-solid:plus" />
        </div>
      </BaseButton>
    {/if}
    {#if $token}
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
    {/if}
  </div>
  
  {#if !$token}
    <div>
      <LoginButton />
    </div>
  {/if}
  
  {#if $token}
    <div class="filter-controls">
      <FilterButtons
        label="All"
        value="all"
        current={$filter}
        onclick={() => filter.set("all")}
        >All Notes</FilterButtons
      >
      <FilterButtons
        label="Active"
        value="active"
        current={$filter}
        onclick={() => filter.set("active")}
        >Active Notes</FilterButtons
      >
      <FilterButtons
        label="Completed"
        value="completed"
        current={$filter}
        onclick={() => filter.set("completed")}
        >Completed Notes</FilterButtons
      >
    </div>
    
    <div class="search">
      <SearchButton debounceTime={500} />    
      <LogoutButton />
    </div>
  {/if}
</header>

{#if toggleNewTask}
  <div
    class="overlay"
    role="button"
    tabindex="0"
    aria-label="Close modal"
    onclick={() => (toggleNewTask = false)}
    onkeydown={(e) => {
      if (e.key === "Enter" || e.key === " ") {
        toggleNewTask = false;
      }
    }}
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
.sticky {
    position: fixed;
    top: 0;
    z-index: 100;
}

.glass {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.header {
    height: 80px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: .25rem 1rem;
    margin-bottom: 1rem;
    box-shadow: 0 8px 16px rgba(0, 0, 0, .5);
    width: 100%;
}

.header-title {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.icon {
    width: 20px;
    height: 20px;
    color: white;
    font-size: 1.2rem;
}
.fade-in {
    animation: fadeIn 0.3s ease-out;
}

.filter-controls {
    display: flex;
    gap: 0.5rem;
    justify-content: center;
    margin-right: 2rem;
}

.overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    animation: fadeIn 0.25s ease-out;
}

.modal {
    padding: 1rem;
    border-radius: 16px;
    min-width: 320px;
    max-width: 600px;
    width: 90%;
    animation: popIn 0.25s ease-out;
}

.search {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 1.25rem;
}

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

/* Tablet Responsiveness */
@media (max-width: 1024px) {

}

/* Phones Responsiveness */
@media (max-width: 600px) {

}
</style>