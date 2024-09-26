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
    let rawData = []
    onMount(async () => {
        fetch("/api/bp")
            .then(response => response.json())
            .then(d => {
                console.log(d)
                rawData = d.rawData
                systolic.data = d.systolic
                diastolic.data = d.diastolic
                pulse.data = d.pulse
                options.xAxis = {categories: d.dates}
            })
    })



    let name = 'jim'
</script>
<style>
    table, th {
    border: 1px solid;
}
</style>

<h1>Welcome to {name}'s blood pressure graph.</h1>

<Chart {options} />


<table>
    <caption>{name}'s Blood Pressure Data</caption>
    <thead>
    <tr>
        <th>Date</th>
        <th>Systolic</th>
        <th>Diastolic</th>
        <th>Pulse</th>
    </tr>
    </thead>
    <tbody>
    {#each rawData as rd }
        <tr>
            <th>{rd.date}</th>
            <th>{rd.systolic}</th>
            <th>{rd.diastolic}</th>
            <th>{rd.pulse}</th>
        </tr>
    {/each}
    </tbody>
</table>