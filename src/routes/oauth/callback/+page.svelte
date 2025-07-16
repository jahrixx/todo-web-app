<script lang="ts">
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { token } from "$lib/stores/auth";

    onMount(() => {
        const url = new URL(window.location.href);
        const jwt = url.searchParams.get("token");

        if(jwt) {
            localStorage.setItem("token", jwt);
            token.set(jwt);
            goto("/");
        } else {
            goto("/login");
        }
    });
</script>

<p>Logging in with Github...</p>