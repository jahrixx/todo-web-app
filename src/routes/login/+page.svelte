<script lang="ts">
  import { loginUser, registerUser, token, userId } from "$lib/stores/auth";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import Header from "$lib/components/Header.svelte";

  let email = $state("");
  let password = $state("");
  let confirmPassword = $state("");
  let error = $state("");
  let activeForm = $state<"login" | "signup">("login");

  onMount(() => {
    activeForm = "login";
  });

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

  async function handleRegister() {
    error = "";
    if (password !== confirmPassword) {
      error = "Passwords do not match";
      return;
    }
    try {
      await registerUser(email, password);
      goto("/");
    } catch (err) {
      error = "Registration failed";
    }
  }
</script>

<Header />
<section class="forms-section">
  <div class="forms">
    <!-- Login Form -->
    <div class="form-wrapper {activeForm === 'login' ? 'is-active' : ''}">
      <form class="form form-login" onsubmit={handleLogin}>
        <fieldset>
          <legend>Login</legend>
          <div class="input-block">
            <label for="login-email">E-mail</label>
            <input id="login-email" type="email" bind:value={email} required />
          </div>
          <div class="input-block">
            <label for="login-password">Password</label>
            <input id="login-password" type="password" bind:value={password} required />
          </div>
        </fieldset>
        <button type="submit" class="btn-login">Login</button>
        <div>
          <button type="button" class="switcher switcher-login" onclick={() => (activeForm = 'signup')}>
            <p style="color: black;">Dont have an account?</p>
          </button>
        </div>
      </form>
    </div>

    <!-- Signup Form -->
    <div class="form-wrapper {activeForm === 'signup' ? 'is-active' : ''}">
      <form class="form form-signup" onsubmit={handleRegister}>
        <fieldset>
          <legend>Sign Up</legend>
          <div class="input-block">
            <label for="signup-email">E-mail</label>
            <input id="signup-email" type="email" bind:value={email} required />
          </div>
          <div class="input-block">
            <label for="signup-password">Password</label>
            <input id="signup-password" type="password" bind:value={password} required />
          </div>
          <div class="input-block">
            <label for="signup-password-confirm">Confirm Password</label>
            <input id="signup-password-confirm" type="password" bind:value={confirmPassword} required />
          </div>
        </fieldset>
        <button type="submit" class="btn-signup">Register</button>
        <div>  
          <button type="button" class="switcher switcher-signup" onclick={() => (activeForm = 'login')}>
            <p style="color: black;">Already have an account?</p>
          </button>
        </div>
      </form>
    </div>
  </div>

  {#if error}
    <p style="color: red; margin-top: 1rem;">{error}</p>
  {/if}
</section>

<style>
 *,
*::before,
*::after {
  box-sizing: border-box;
}

:global(body) {
  margin: 0;
  font-family: Roboto, -apple-system, 'Helvetica Neue', 'Segoe UI', Arial, sans-serif;
  background: url('../../lib/assets/wallpaper.jpg');
}

.forms-section {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.forms {
  display: flex;
  align-items: flex-start;
  margin-top: 150px;
}

.form-wrapper {
  animation: hideLayer .3s ease-out forwards;
}

.form-wrapper.is-active {
  animation: showLayer .3s ease-in forwards;
}

@keyframes showLayer {
  50% {
    z-index: 1;
  }
  100% {
    z-index: 1;
  }
}

@keyframes hideLayer {
  0% {
    z-index: 1;
  }
  49.999% {
    z-index: 1;
  }
}

.switcher {
  position: relative;
  cursor: pointer;
  display: block;
  margin-right: auto;
  margin-left: auto;
  padding: 0;
  text-transform: uppercase;
  font-family: inherit;
  font-size: 12px;
  letter-spacing: .5px;
  color: #999;
  background-color: transparent;
  border: none;
  outline: none;
  font-style: italic;
  transform: translateX(0);
  transition: all .3s ease-out;
}

.form {
  overflow: hidden;
  min-width: 100px;
  margin-top: 50px;
  padding: 30px 70px;
  border-radius: 5px;
  transform-origin: top;
}

.form-login {
  animation: hideLogin .3s ease-out forwards;
}

.form-wrapper.is-active .form-login {
  animation: showLogin .3s ease-in forwards;
}

@keyframes showLogin {
  0% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(40%, 10px);
  }
  50% {
    transform: translate(0, 0);
  }
  100% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(35%, -20px);
  }
}

@keyframes hideLogin {
  0% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(35%, -20px);
  }
  50% {
    transform: translate(0, 0);
  }
  100% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(40%, 10px);
  }
}

.form-signup {
  animation: hideSignup .3s ease-out forwards;
}

.form-wrapper.is-active .form-signup {
  animation: showSignup .3s ease-in forwards;
}

@keyframes showSignup {
  0% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(-40%, 10px) scaleY(.8);
  }
  50% {
    transform: translate(0, 0) scaleY(.8);
  }
  100% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(-35%, -20px) scaleY(1);
  }
}

@keyframes hideSignup {
  0% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(-35%, -20px) scaleY(1);
  }
  50% {
    transform: translate(0, 0) scaleY(.8);
  }
  100% {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    transform: translate(-40%, 10px) scaleY(.8);
  }
}

.form fieldset {
  position: relative;
  opacity: 0;
  margin: 0;
  padding: 0;
  border: 0;
  transition: all .3s ease-out;
}

.form-login fieldset {
  transform: translateX(-50%);
}

.form-signup fieldset {
  transform: translateX(50%);
}

.form-wrapper.is-active fieldset {
  opacity: 1;
  transform: translateX(0);
  transition: opacity .4s ease-in, transform .35s ease-in;
}

.form legend {
  position: absolute;
  overflow: hidden;
  width: 1px;
  height: 1px;
  clip: rect(0 0 0 0);
}

.input-block {
  margin-bottom: 10px;
}

.input-block label {
  font-size: 14px;
  color: black;
}

.input-block input {
  display: block;
  width: 100%;
  margin-top: 8px;
  padding-right: 15px;
  padding-left: 15px;
  font-size: 16px;
  line-height: 30px;
  color: #3b4465;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: none;
}

.form [type='submit'] {
  opacity: 0;
  display: block;
  min-width: 120px;
  margin: 10px auto 10px;
  font-size: 18px;
  line-height: 40px;
  border-radius: 25px;
  border: none;
  transition: all .3s ease-out;
}

.form-wrapper.is-active .form [type='submit'] {
  opacity: 1;
  transform: translateX(0);
  transition: all .4s ease-in;
}

.btn-login {
  color: #fbfdff;
  background: #1d4ed8;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  transform: translateX(-30%);
}

.btn-signup {
  color: #fbfdff;
  background: #1d4ed8;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  transform: translateX(30%);
}
</style>