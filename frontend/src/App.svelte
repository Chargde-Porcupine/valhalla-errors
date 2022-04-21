
<script>
import { compute_slots } from "svelte/internal";
//TODO: update Verrori list on edit

  import Verror from "./lib/Verror.svelte";
  import NewVerror from "./lib/NewVerror.svelte";
  export let verrori;
  let search = true
  let loading = "loading"
  let searched = ""
  export const sendRequest = async (path) => {
  if (path == "") {
    path = "all"
  } else {
    path = "find/".concat(path);
  }
  const response = await fetch('https://quote-berkeley-desk-edmonton.trycloudflare.com/'.concat(path));
  const data = await response.json();
  if(response.ok){
    return data;
  }
  else{
    throw new Error(data.message);
  }
  }
  function reload(event) {
    verrori = sendRequest(searched)
  }
  function SwitchView(){
    search = !search
    searched = ""
    loading = "loading"
    reload()
    loading = ""
  }
  reload() // won't print anything as verrori is still a promise
</script>
<main>

  <h1>V-Errors!</h1>
  <h2>Github Issues Meets Chief Delphi</h2>


<button on:click = {SwitchView}>{#if search}Report a Verror{:else}Search for Verrors{/if}</button>
<br>
{#if search}
<input bind:value="{searched}"> <button on:click="{reload}">search</button>
<br>
{#await verrori}
<p>{loading}</p>
{:then verrors}
{#each verrors as verror}
  <Verror on:message={reload} title = {verror.Title} desc = {verror.Desc} status = {verror.Status} category = {verror.Category} solution = {verror.Solution} imageurl = {verror.ImageL}></Verror>
{/each}

{:catch error}
  <p>{error.message}</p>
{/await}
{:else}

<NewVerror></NewVerror>

{/if}
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  h1 {
    font-size: 2em;
  }
  h2 {
    font-size: 1.5em;
  }
  p {
    font-size: 1em;
  }
  button {
    font-size: 1em;
  }
  input {
    font-size: 1em;
  }
  select {
    font-size: 1em;
  }
  
</style>
