
<script>
import { compute_slots } from "svelte/internal";
//TODO: update Verrori list on edit

  import Verror from "./lib/Verror.svelte";
  import NewVerror from "./lib/NewVerror.svelte";
  export let verrori;
  let search = true
  let searched = ""
  export const sendRequest = async (path) => {
  if (path == "") {
    path = "all"
  } else {
    path = "find/".concat(path);
  }
  const response = await fetch('http://localhost:8080/'.concat(path));
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
    reload()
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
<p>loading</p>
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
</style>
