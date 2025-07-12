<script lang="ts">
  import { registerUser } from "$lib/stores/auth";
  import { goto } from "$app/navigation";

  let email = $state("");
  let password = $state("");
  let error = $state("");

  async function handleRegister() {
    error = "";
    try {
      await registerUser(email, password);
      goto("/login");
    } catch (err) {
      error = "Registration failed";
    }
  }
</script>

<div class="auth-container">
  <h2>Register</h2>
  {#if error}<p class="error">{error}</p>{/if}
  <form onsubmit={handleRegister}>
    <input type="email" placeholder="Email" bind:value={email} required />
    <input type="password" placeholder="Password" bind:value={password} required />
    <button type="submit">Register</button>
  </form>
</div>
