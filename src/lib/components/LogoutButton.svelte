<script lang="ts">
    import { token } from "$lib/stores/auth";
    import Icon from "@iconify/svelte";

    let { toggleLogout = false, ...props }: { toggleLogout?: boolean; [key: string]: any; } = $props();
    
    function logout() {
        if(!confirm('Are you sure you want to logout?')) {
            return;
        } else {
            token.set(null);
            window.location.href = '/login';
        }
    }

    $effect(() => {
        if (toggleLogout) {
            setTimeout(() => {
                logout();
            }, 0);
        }
    });
</script>


<div class="logout-wrapper" {...props}>
    <button class="logout-icon" onclick={() => toggleLogout = !toggleLogout} aria-labelledby="logout">
        <Icon icon="streamline-sharp:logout-2-remix" style="font-size: 1.3rem; margin-top: .25rem;" />
    </button>    
</div>

<style>
    .logout-wrapper {
        margin: 0;
        padding: 0;
    }

    .logout-icon {
        background: transparent;
        border: none;
        cursor: pointer;
        padding: .5rem;
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 100;
    }
</style>