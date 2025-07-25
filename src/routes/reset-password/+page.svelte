<script lang="ts">
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";
  let password = '';
  const token = $page.url.searchParams.get('token');

  async function reset() {
    const res = await fetch("http://localhost:8080/api/auth/reset-password", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ token, password }),
    });
    if (res.ok) {
      alert("Password updated!");
      goto("/login");
    } else {
      alert("Invalid or expired token");
    }
  }
</script>

<input type="password" bind:value={password} placeholder="New password" />
<button onclick={reset}>Reset Password</button>
