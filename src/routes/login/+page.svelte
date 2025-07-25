<script lang="ts">
  import { loginUser, token, userId } from "$lib/stores/auth";
  import { goto } from "$app/navigation";
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
    window.location.href = import.meta.env.VITE_GITHUB_LOGIN_URL;
  }
</script>

<Header />
<div class="login-page">
  <div class="form" >
    <div class="login-subtitle">Login Using A Social Media Account</div>
    <div class="social-buttons">
      <button onclick={loginWithGitthub} class="social-btn github-btn">
        <svg fill="white" width="16" height="16" viewBox="0 0 24 24"><path d="M12 .3a12 12 0 00-3.8 23.4c.6.1.8-.3.8-.6v-2.2c-3.3.7-4-1.6-4-1.6-.5-1.3-1.2-1.7-1.2-1.7-1-.7.1-.7.1-.7 1.1.1 1.7 1.2 1.7 1.2 1 .1 1.6-.8 1.6-.8.2-.8.4-1.1.7-1.4-2.6-.3-5.3-1.3-5.3-5.9 0-1.3.5-2.4 1.2-3.2-.1-.3-.5-1.5.1-3.1 0 0 1-.3 3.3 1.3a11.5 11.5 0 016 0c2.2-1.6 3.3-1.3 3.3-1.3.6 1.6.2 2.8.1 3.1.7.8 1.2 1.9 1.2 3.2 0 4.7-2.7 5.5-5.3 5.8.4.3.8 1 .8 2v3c0 .3.2.7.8.6A12 12 0 0012 .3" /></svg>
        Sign in with Github
      </button>
    </div>

    <div class="divider">Or</div>

    <form class="login-form" onsubmit={handleLogin}>
      <input type="email" placeholder="Email" bind:value={email} required/>
      <input type="password" placeholder="Password" bind:value={password} required/>
      <button class="login-btn" type="submit">login</button><br>
    
      <p class="message"><a href="/forgot-password">Forgot Password?</a></p>
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
  .login-subtitle {
    color: #888;
    font-size: 14px;
    margin-bottom: 1.5rem;
  }

  .social-buttons {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-bottom: 1rem;
  }

  .social-btn {
    display: flex;
    align-items: center;
    gap: 6px;
    border: 1px solid #ddd;
    border-radius: 6px;
    padding: 0.5rem 1rem;
    font-size: 0.9rem;
    cursor: pointer;
    background-color: white;
  }
  .github-btn {
    background-color: #24292f;
    color: white;
    border: none;
  }
  .login-page {
    width: 360px;
    padding: 11% 0 0;
    margin: auto;
  }
  .form {
    position: relative;
    z-index: 1;
    max-width: 360px;
    margin: 0 auto 100px;
    padding: 45px;
    text-align: center;
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  }
  .form input {
    font-family: "Roboto", sans-serif;
    outline: 0;
    width: 100%;
    border: 0;
    margin: 0 0 15px;
    padding: 15px;
    box-sizing: border-box;
    font-size: 14px;
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  }
  .login-btn {
    font-family: "Roboto", sans-serif;
    text-transform: uppercase;
    outline: 0;
    width: 50%;
    border: 0;
    padding: 10px;
    color: #888;
    font-size: 14px;
    -webkit-transition: all 0.3 ease;
    transition: all 0.3 ease;
    cursor: pointer;
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  }
  .login-btn:hover {
    background: rgba(255, 255, 255, 0.5);
  }
  .form .message {
    margin: 15px 0 0;
    color: #b3b3b3;
    font-size: 12px;
  }
  .form .message a {
    color: #888;
    text-decoration: none;
  }

  .form .message a:hover {
    text-decoration: underline;
    color: #777;
  }

  .divider {
    display: flex;
    align-items: center;
    text-align: center;
    margin-bottom: 1rem;
  }

  .divider::before,
  .divider::after {
    content: "";
    flex: 1;
    height: 1px;
    background: #888;
  }

  .divider:not(:empty)::before {
    margin-right: 0.75em;
  }

  .divider:not(:empty)::after {
    margin-left: 0.75em;
  }
  
  @media (min-width: 360px) and (max-width: 500px) {
    .login-page {
      margin-top: 25%;
    }
  }
  @media (min-width: 501px) and (max-width: 599px) {
    .login-page {
      margin-top: 15%;
    }
  }
  @media (min-width: 600px) and (max-width: 767px) {
    .login-page {
      margin-top: 15%;
    }
  }
  @media (min-width: 768px) and (max-width: 819px) {
    .login-page {
        margin-top: 11.5%;
    }
  }
  @media (min-width: 820px) and (max-width: 912px)  {
    .login-page {
        margin-top: 10%;
    }
  }
</style>