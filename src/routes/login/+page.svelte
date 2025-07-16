<script lang="ts">
  import { loginUser, registerUser, token, userId } from "$lib/stores/auth";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import Header from "$lib/components/Header.svelte";

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

  function loginWithGitthub() {
    window.location.href = "http://localhost:8080/api/auth/github";
  }
</script>

<Header />
<div class="login-page">
  <div class="form" >
    <form class="login-form" onsubmit={handleLogin}>
      <input type="email" placeholder="Email" bind:value={email} required/>
      <input type="password" placeholder="Password" bind:value={password} required/>
      <button type="submit">login</button><br>
      <p>or</p>
      <button onclick={loginWithGitthub} class="btn-login">Sign in with Github</button>
      <p class="message">Not registered? <a href="/register">Create an account</a></p>
    </form>
  </div>
    {#if error}
      <p style="color: red; margin-top: 1rem;">{error}</p>
    {/if}
</div>

<style>
  :global(body) {
    padding: 0;
    margin: 0;
    background: url('../../lib/assets/wallpaper.jpg');
  }
  .login-page {
    width: 360px;
    padding: 8% 0 0;
    margin: auto;
  }
  .form {
    position: relative;
    z-index: 1;
    background: #FFFFFF;
    max-width: 360px;
    margin: 0 auto 100px;
    padding: 45px;
    text-align: center;
    box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
  }
  .form input {
    font-family: "Roboto", sans-serif;
    outline: 0;
    background: #f2f2f2;
    width: 100%;
    border: 0;
    margin: 0 0 15px;
    padding: 15px;
    box-sizing: border-box;
    font-size: 14px;
  }
  .form button {
    font-family: "Roboto", sans-serif;
    text-transform: uppercase;
    outline: 0;
    background: #4CAF50;
    width: 100%;
    border: 0;
    padding: 15px;
    color: #FFFFFF;
    font-size: 14px;
    -webkit-transition: all 0.3 ease;
    transition: all 0.3 ease;
    cursor: pointer;
  }
  .form button:hover,.form button:active,.form button:focus {
    background: #43A047;
  }
  .form .message {
    margin: 15px 0 0;
    color: #b3b3b3;
    font-size: 12px;
  }
  .form .message a {
    color: #4CAF50;
    text-decoration: none;
  }
</style>