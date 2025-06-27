<script>
  import { token } from '../lib/auth.js';
  let file = null;
  let message = '';

  async function upload() {
    message = '';
    if (!file) {
      message = "No file selected!";
      return;
    }
    const formData = new FormData();
    formData.append('video', file);

    // Don't set Content-Type when using FormData!
    try {
      const res = await fetch('http://localhost:8080/upload', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${$token}` // If using JWT, else remove this line
          // Do NOT set 'Content-Type'
        },
        body: formData
      });
      if (res.ok) {
        message = "Upload successful!";
      } else {
        const errorText = await res.text();
        message = "Upload failed: " + errorText;
      }
    } catch (e) {
      message = "Network error: " + e;
    }
  }
</script>

<form on:submit|preventDefault={upload}>
  <input type="file" accept="video/*" on:change={e => file = e.target.files[0]} required>
  <button>Upload Video</button>
  {#if message}<div>{message}</div>{/if}
</form>