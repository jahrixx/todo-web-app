<script lang="ts">
  import { goto } from "$app/navigation";
  let email = '';

  async function submit() {
    const res = await fetch("http://localhost:8080/api/auth/forgot-password", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email }),
    });
    if (res.ok) {
      alert("Check your email for reset instructions");
      goto("/login");
    }
  }
</script>

<input bind:value={email} type="email" placeholder="Your email" />
<button onclick={submit}>Send Reset Link</button>
