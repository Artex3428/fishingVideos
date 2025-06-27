<script>
  import { token } from './lib/auth.js';
  import Login from './components/Login.svelte';
  import Upload from './components/Upload.svelte';
  import VideoList from './components/VideoList.svelte';

  let page = 'list';

  $: authed = $token.length > 0;

  function logout() { token.set(""); }
</script>

<nav>
  {#if authed}
    <button on:click={() => page = 'list'}>Videos</button>
    <button on:click={() => page = 'upload'}>Upload</button>
    <button on:click={logout}>Logout</button>
  {/if}
</nav>

{#if !authed}
  <Login />
{:else if page === 'upload'}
  <Upload />
{:else}
  <VideoList />
{/if}