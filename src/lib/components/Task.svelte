<script lang="ts">
    import Icon from "@iconify/svelte";
    import { type Task, addTask, removeTask, updateTask } from "$lib/stores/tasks";
    import CancelButton from "./CancelButton.svelte";
    import DoneButton from "./DoneButton.svelte";
    import FormButtons from "./FormButtons.svelte";
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

<div class="task-wrapper glass" in:slide={{duration: 500}} out:slide={{duration: 500}} {...props} role="region" aria-labelledby="task">
    {#if task && !isEditable}
        <div class="task-container" in:fade={{duration: 500}}>
            <div class="task-content">
                <h2 id="task">{task.title}</h2>
                {#if task.description}
                    <p>{task.description}</p>
                {/if}
            </div>
            <div class="task-actions">
                <FormButtons onclick={() => removeTask(task)} aria-labelledby="remove task">
                    <Icon icon="fa6-solid:trash" />
                </FormButtons>
                <FormButtons onclick={() => (isEditable = true)} aria-labelledby="update task">
                    <Icon icon="fa6-solid:pencil" />
                </FormButtons>
                <FormButtons onclick={() => task && updateTask(task.id, { ...task, completed: !task.completed })} aria-labelledby={task.completed ? "Mark task as done" : "Mark task as not done"}>
                    {#if task.completed}
                        <Icon icon="mdi:checkbox-outline" />
                    {:else}
                        <Icon icon="mdi:checkbox-blank-outline" />
                    {/if}
                </FormButtons>
            </div>
        </div>
        {:else}
        <form class="task-form" action="#" in:fade={{duration: 150}} onsubmit={(e) => e.preventDefault()} aria-labelledby="task-edit">
            <div class="task-form-group">
                <label for="task-title-input" class="form-label">Task Name</label>
                <input id="task-title-input" class="form-input" placeholder="Enter a Task Name" type="text" name="task-title" required bind:value={taskName} />
                
                <label for="task-description-input" class="form-label">Task Description</label>
                <textarea id="task-description-input" class="form-input" name="description" placeholder="Enter a Task Description" bind:value={taskDescription}></textarea>
            </div>
            <div class="task-form-actions">
                <CancelButton
                    onclick={cancelForm}
                    aria-label="Cancel editing task">
                    Cancel
                </CancelButton>

                <DoneButton
                    action="submit"
                    onclick={submitForm}
                    aria-label="Finish editing task">
                    Done
                </DoneButton>
            </div>
        </form>
    {/if}
</div>

<style>
    .glass {
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }

    .careful-text-overflow {
        word-wrap: break-word;
        white-space: pre-wrap;
        word-break: break-word;
    }

    .task-wrapper {
        margin: 0;
        padding: 1rem;
        border-radius: 1rem;
        transition: all 0.3s ease-in-out;
        height: auto;
        width: 90%;
        max-width: unset;
        display: block;
        break-inside: avoid;
        outline: 2px dashed black;
        min-height: 0;
        margin-bottom: 1rem;
    }

    .task-container {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .task-content {
        padding-bottom: 0.5rem;
        border-bottom: 1px solid white;
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

    .careful-text-overflow {
        word-wrap: break-word;
        white-space: pre-wrap;
        word-break: break-word;
    }

    .task-form {
        padding: .5rem;
        border-radius: 1rem;
        max-width: 100%;
        margin: auto 0;
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
        animation: fadeIn 0.3s ease-out;
    }

    .task-form-group {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .form-label {
        font-size: 1.2rem;
        font-weight: 600;
        color: #FAF9F6;
        margin-bottom: 0.2rem;
    }

    .form-input {
        padding: 0.75rem 1rem;
        border: 1px solid #d1d5db;
        border-radius: 0.5rem;
        font-size: 1rem;
        outline: none;
        transition: border-color 0.2s ease, box-shadow 0.2s ease;
    }

    .form-input::placeholder {
        font-size: 0.85rem;
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