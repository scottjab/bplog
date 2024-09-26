<script lang="ts">
    import { onMount } from "svelte";
    import Highcharts from 'highcharts';
    import { Chart } from '@highcharts/svelte'; // Chart is also exported by default


    let systolic: Highcharts.SeriesLineOptions = {
            type: 'line',
            name: 'Systolic',
            data: [0]
    }

    let diastolic: Highcharts.SeriesLineOptions = {
        type: 'line',
        name: 'Diastolic',
        data: [0]
    }

    let pulse: Highcharts.SeriesLineOptions = {
        type: 'line',
        name: 'Pulse',
        data: [0]
    }
    let items = ['Systolic', 'Diastolic', 'Pulse']

    let options: Highcharts.Options  = {
        title: {
            text: 'Blood Pressure'
        },
        series: [systolic,diastolic,pulse]
    };


    onMount(async () => {
        fetch("/api/bp")
            .then(response => response.json())
            .then(d => {
                console.log(d)
                systolic.data = d.systolic
                diastolic.data = d.diastolic
                pulse.data = d.pulse
                options.xAxis = {categories: d.dates}
            })
    })



    let name = 'jim'
</script>


<h1>Welcome to {name}'s blood pressure graph.</h1>

<Chart {options} />

