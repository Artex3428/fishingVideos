<script>
  let videos = [];
  let error = '';

  async function fetchVideos() {
    try {
      const res = await fetch('http://localhost:8080/videos');
      if (res.ok) {
        videos = await res.json();
      } else {
        error = "Failed to load videos";
      }
    } catch (e) {
      error = e;
    }
  }
  fetchVideos();
</script>

{#if error}
  <div>{error}</div>
{/if}
<ul>
  {#each videos as video}
    <li>
      <strong>{video.name}</strong>
      <video src={`http://localhost:8080${video.url}`} controls width="320"></video>
      <a href={`http://localhost:8080${video.url}`} download>Download</a>
    </li>
  {/each}
</ul>