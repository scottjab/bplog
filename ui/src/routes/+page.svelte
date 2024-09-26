<script>
    import { onMount } from "svelte";
    let apiData = {}

    onMount(async () => {
        fetch("/ping")
            .then(response => response.json())
            .then(data => {
                console.log(data.message);
                apiData = data;
            })
    })
    let name = 'jim'
</script>


<h1>Welcome to {name}</h1>
<h2> Message {apiData.message} </h2>
<h2> Message {apiData.time_stamp} </h2>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
