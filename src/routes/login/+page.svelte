<script lang="ts">
  import { loginUser } from "$lib/stores/auth";
  import { token, userId } from "$lib/stores/auth";
  import { goto } from "$app/navigation";

  let email = $state("");
  let password = $state("");
  let error = $state("");

  async function handleLogin() {
    error = "";
    try {
      const data = await loginUser(email, password);
      token.set(data.token);
      userId.set(data.user_id);
      goto("/");
    } catch (err) {
      error = "Invalid login credentials";
    }
  }
</script>

<div class="auth-container">
  <h2>Login</h2>
  {#if error}<p class="error">{error}</p>{/if}
  <form onsubmit={handleLogin}>
    <input type="email" placeholder="Email" bind:value={email} required />
    <input type="password" placeholder="Password" bind:value={password} required />
    <button type="submit">Login</button>
  </form>
</div>

<style>
  .auth-container {
    max-width: 400px;
    margin: auto;
    padding: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }
  input, button {
    padding: 0.5rem;
    font-size: 1rem;
  }
  .error {
    color: red;
  }
</style>
