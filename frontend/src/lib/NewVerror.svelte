<script>
import { each } from "svelte/internal";
import { statuslist, uberCat, uberlist, unterCat, unterlist } from "./tables";


    //post request, multipart (file and data)
    let title;
    let desc;
    let status;
    let cd1;
    let cd2;
    let solution;
    let imageurl;
    let fileinput;
    let image;
    let error = ""
    let ext = "";
    
  export function submitVerror(){
    if (title == undefined || desc == undefined || status == undefined || cd1 == undefined || cd2 == undefined || ext == undefined){
      error = "Fill out the form my dude"
      return
    }
    let data  = JSON.stringify({
        "Title": title,
        "Desc": desc,
        "Status": status,
        "Category": (cd1*10 + cd2),
        "Solution": solution,
        "ImageL": ext

    })


    data = data.replaceAll('\\', '');

    console.log(data)

    const formData = new FormData();
    formData.append("data", data)
    formData.append("file", image)
    const options = {
    method: 'POST',
    body: formData,
    };
    fetch("http://localhost:8080/create", options).then((res) => {
      if(res.ok) {
        error = "Verror reported."
      } else {
        error = "Verror in Verror reporting(report this one too and see how far you get in Verror recursion)";
      }
    })
  }

  export function onFileSelected(e){
    image = e.target.files[0];
    ext.concat(e.target.files[0].name);
    let reader = new FileReader();
            reader.readAsDataURL(image);
            reader.onload = e => {
                 imageurl = e.target.result
            };
  }

  export function getCat(number){
    return uberCat[Math.floor(number/10)].concat( "/", unterCat[number % 10]);
  }
</script>

<main>
    <p>Report a new Verror</p>
    <p>Title</p>
    <input bind:value={title}>
    <br>
    <p>Description of the Error</p>
    <input bind:value={desc}>
    <br>
    <p>Status</p>
    <select bind:value={status}>
        {#each statuslist as status}
            <option value={statuslist.indexOf(status)}>{status}</option>
        {/each}
    
    </select>
    <br>
    <p>Category</p>
    <select bind:value = {cd1}>
        {#each uberlist as ubc}
        <option value={uberlist.indexOf(ubc)+1}>{ubc}</option>
        {/each}
    </select>
    <p>Sub-Category</p>
    <select bind:value = {cd2}>
        {#each unterlist as ubc}
        <option value={unterlist.indexOf(ubc)+1}>{ubc}</option>
        {/each}
    </select>
    <br>
    <p>Solution</p>
    <input bind:value={solution}>
    <br>
    <p>Photo of Error</p>
    <input type="file" accept=".jpg, .jpeg, .png" bind:this={fileinput} on:change="{(e) => onFileSelected(e)}" >
    <br>
    <img class="image" src={imageurl} alt="Upload an Image">
    <br>
    <button on:click="{submitVerror}">Report</button>
    <p>{error}</p>
</main>