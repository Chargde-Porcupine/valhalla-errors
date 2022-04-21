
<script>
import { statuslist, statustable, uberCat, unterCat } from "./tables";

	import { createEventDispatcher } from 'svelte';


  export let imageurl;
  export let title;
  export let desc;
  export let status;
  export let category;
  export let solution;
  export let editable = false;
  let updatedStatus = 0;
  let updatedSolution = "";

  const dispatch = createEventDispatcher();
  export function SwitchView(){
    editable = !editable;
  }
  export function SubmitSolution(){
     console.log(category)
     let data  = JSON.stringify({
        "Title": title,
        "Desc": desc,
        "Status": updatedStatus,
        "Category": category,
        "Solution": updatedSolution,
        "ImageL": "uploaded.jpg"

    })

    data = data.replaceAll('\\', '');
    const options = {
    method: 'POST',
    body: data,
    };
    fetch("http://localhost:8080/solve/".concat(imageurl), options)
    SwitchView();
    dispatch('message', {text: "reload"})
  }
  export function getCat(number){
    return uberCat[Math.floor(number/10)].concat( "/", unterCat[number % 10]);
  }
</script>

<main>
  {#if !editable}
  <div>

    <button on:click="{SwitchView}">{#if editable}Discard Edit{:else}Edit{/if}</button>
  <h1>{title}</h1>
  <h2>{desc}</h2>
  <p>Category : {getCat(category)}</p>
  <p>Status : {statustable[status]}</p>
  {#if solution != ""}
  <p>Solution : {solution}</p>
  {/if}
  <br>
  <img src="http://localhost:8080/images/{imageurl}" alt={title}>
  </div>
  {:else}
  <div>
    <button on:click="{SwitchView}">{#if editable}Discard Edits{:else}Edit{/if}</button>
    <h1>{title}</h1>
    <h2>{desc}</h2>
    <p>Category : {getCat(category)}</p>
    <br>
    <p>Status</p>
    <select bind:value={updatedStatus}>
        {#each statuslist as status}
            <option value={statuslist.indexOf(status)}>{status}</option>
        {/each}
    
    </select>
    <br>
    <br>
    <p>Solution</p>
    <input bind:value={updatedSolution}>
    <br>
    <br>
    <img src="https://quote-berkeley-desk-edmonton.trycloudflare.com/images/{imageurl}" alt={title} width="300" height="300">
    <button on:click = {SubmitSolution} >Submit</button>
    </div>
  {/if}
</main>

<style>
  div{
      border: 1px solid black ;
      margin: 10px;
      padding: 10px;
  }
  button{
    background-color: black;
    color: yellow;
  }
</style>