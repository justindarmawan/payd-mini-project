<!-- /src/routes/employee/+page.svelte -->
<script>
  import {
    getAssignedShifts,
    getAvailableShifts,
    requestShift,
  } from "$lib/api/shifts";
  import { onMount } from "svelte";

  let assigned = [];
  let available = [];

  onMount(async () => {
    assigned = await getAssignedShifts();
    available = await getAvailableShifts();
  });

  async function request(id) {
    await requestShift(id);
    available = available.filter((s) => s.id !== id);
  }
</script>

<div class="container">
  <h1>Assigned Shifts</h1>
  {#each assigned as s}
    <div class="shift">
      <p>{s.date} - {s.start_time} to {s.end_time} - {s.role}</p>
    </div>
  {/each}

  <h2>Available Shifts</h2>
  {#each available as s}
    <div class="shift">
      <p>{s.date} - {s.start_time} to {s.end_time} - {s.role}</p>
      <button on:click={() => request(s.id)}>Request</button>
    </div>
  {/each}
</div>

<style lang="scss">
  .container {
    padding: 2rem;
  }
  .shift {
    border: 1px solid #ccc;
    border-radius: 0.5rem;
    padding: 1rem;
    margin-bottom: 1rem;
  }
  h2 {
    margin-top: 2rem;
  }
</style>
