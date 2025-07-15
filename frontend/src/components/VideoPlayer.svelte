<script>
  export let videoUrl = '';
  let start = 0;
  let end = 0;
  let videoEl;

  function cut() {
    // This just sets the start/end time in the player
    videoEl.currentTime = start;
    videoEl.ontimeupdate = () => {
      if (videoEl.currentTime >= end && end > 0) {
        videoEl.pause();
      }
    };
    videoEl.play();
  }
</script>

<div>
  <video bind:this={videoEl} src={videoUrl} controls width="400"></video>
  <div>
    <input type="number" bind:value={start} placeholder="Start (sec)">
    <input type="number" bind:value={end} placeholder="End (sec)">
    <button on:click={cut}>Play Cut</button>
  </div>
</div>