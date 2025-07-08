<script lang="ts">
    import Icon from "@iconify/svelte";
    import { type Task, addTask, removeTask, updateTask } from "$lib/stores/tasks";
    import BaseButton from "./BaseButton.svelte";
    import { fade, slide } from "svelte/transition";

    let { task, isShown = $bindable(false), ...props } : { task?: Task, isNotCancelled?: boolean, [key: string]: any} = $props();

    let taskName = $state(task?.title || '');
    let taskDescription = $state(task?.description || '');
    let isEditable = $state(false);

    function cancelForm() {
        if(isEditable) {
            taskName = task!.title;
            taskDescription = task!.description || taskDescription;
            isEditable = false;
        } else {
            isShown = false;
        }
    }

    function submitForm() {
        if(isEditable && taskName !== "") {
            updateTask(task!.id, {
                id: task!.id,
                title: taskName,
                description: taskDescription,
                completed: task!.completed,
            });
            isEditable = false;
        } else if (taskName !== "") {
            addTask({
                id: Math.random(),
                title: taskName,
                description: taskDescription,
                completed: false,
            });
            isShown = false;
        }
    }
</script>

<div class="task-wrapper" in:slide={{duration: 500}} out:slide={{duration: 500}} {...props} role="region" aria-labelledby="task">
    {#if task && !isEditable}
        <div class="task-container" in:fade={{duration: 500}}>
            <div class="task-content">
                <h2 id="task">{task.title}</h2>
                {#if task.description}
                    <p>{task.description}</p>
                {/if}
            </div>
            <div class="task-actions">
                <BaseButton onclick={() => removeTask(task)} aria-labelledby="remove task">
                    <Icon icon="fa6-solid:trash" />
                </BaseButton>
                <BaseButton onclick={() => (isEditable = true)} aria-labelledby="update task">
                    <Icon icon="fa6-solid:pencil" />
                </BaseButton>
                <BaseButton onclick={() => task && updateTask(task.id, { ...task, completed: !task.completed })} aria-labelledby={task.completed ? "Mark task as done" : "Mark task as not done"}>
                    {#if task.completed}
                        <Icon icon="mdi:checkbox-outline" />
                    {:else}
                        <Icon icon="mdi:checkbox-blank-outline" />
                    {/if}
                </BaseButton>
            </div>
        </div>
        {:else}
        <form class="task-form" action="#" in:fade={{duration: 150}} onsubmit={(e) => e.preventDefault()} aria-labelledby="task-edit">
            <div class="task-form-group">
                <label for="task-title-input" class="form-label">Task name</label>
                <input id="task-title-input" class="form-input" placeholder="Enter a Task Name" type="text" name="task-title" required bind:value={taskName} />
                
                <label for="task-description-input" class="form-label">Task description</label>
                <textarea id="task-description-input" class="form-input" name="description" placeholder="Enter a Task Description" bind:value={taskDescription}></textarea>
            </div>
            <div class="task-form-actions">
                <BaseButton
                    onclick={cancelForm}
                    aria-label="Cancel editing task">
                    Cancel
                </BaseButton>

                <BaseButton
                    action="submit"
                    onclick={submitForm}
                    aria-label="Finish editing task">
                    Done
                </BaseButton>
            </div>
        </form>
    {/if}
</div>

<style>
    .careful-text-overflow {
        word-wrap: break-word;
        white-space: pre-wrap;
        word-break: break-word;
    }

    .task-wrapper {
        margin: 1rem auto;
        padding: 1rem;
        max-width: 600px;
        border-radius: 1rem;
        background: #ffffff;
        box-shadow: 0 6px 16px rgba(0, 0, 0, 0.1);
        transition: all 0.3s ease-in-out;
    }

    .task-container {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .task-content {
        padding-bottom: 0.5rem;
        border-bottom: 1px solid #e0e0e0;
    }

    .task-content h2 {
        margin: 0;
        font-size: 1.25rem;
        font-weight: 600;
        color: #333;
    }

    .task-content p {
        margin-top: 0.5rem;
        color: #555;
        line-height: 1.5;
        font-size: 0.95rem;
    }

    .task-actions {
        display: flex;
        gap: 0.75rem;
        justify-content: flex-end;
    }

    .btn {
        background-color: #f3f4f6;
        border: none;
        padding: 0.5rem 0.75rem;
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .btn:hover {
        background-color: #e5e7eb;
    }

    .btn.danger {
        background-color: #fee2e2;
    }

    .btn.danger:hover {
        background-color: #fecaca;
    }

    .careful-text-overflow {
        word-wrap: break-word;
        white-space: pre-wrap;
        word-break: break-word;
    }

    .task-form {
        background: #ffffff;
        padding: 2rem;
        border-radius: 1rem;
        box-shadow: 0 6px 18px rgba(0, 0, 0, 0.06);
        max-width: 600px;
        margin: auto;
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
        animation: fadeIn 0.3s ease-out;
    }

    .task-form-group {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .form-label {
        font-weight: 600;
        color: #374151;
        margin-bottom: 0.25rem;
    }

    .form-input {
        padding: 0.75rem 1rem;
        border: 1px solid #d1d5db;
        border-radius: 0.5rem;
        font-size: 1rem;
        outline: none;
        transition: border-color 0.2s ease, box-shadow 0.2s ease;
    }

    .form-input:focus {
        border-color: #2563eb;
        box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15);
    }

    .task-form-actions {
        display: flex;
        justify-content: flex-end;
        gap: 1rem;
    }

    /* Fade in animation */
    @keyframes fadeIn {
        from {
            opacity: 0;
            transform: translateY(-10px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }
</style>