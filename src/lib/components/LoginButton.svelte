<script lang="ts">
    function handleLogin() {
        window.location.href = "/login";
    }
    function handleRegister() {
        window.location.href = "/register";
    }
</script>

<div class="login-register-button-container">
    <button type="button" onclick={handleLogin} class="login-action-btns">Sign In</button>
    <button type="button" onclick={handleRegister} class="login-action-btns">Sign Up</button>
</div>

<style>
    .login-register-button-container {
        display: flex;
        justify-content: center;
        gap: 10px;
        margin-right: 30px;
    }

    .login-action-btns {
        display: flex;
        align-items: center;
        border-radius: 6px;
        padding: 0.5rem 1rem;
        font-size: 0.9rem;
        border: none;
        cursor: pointer;
        color: #1d4ed8;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    
</style>
