<template>
    <div id="chart" />
</template>

<script setup lang="ts">
import ApexCharts from 'apexcharts'
import { onMounted } from 'vue';
import { type BadgeWithValue } from '../services/badges/badges';

const props = defineProps<{ badges: BadgeWithValue[] }>();

var options = {
    series: [{ data: props.badges.map(badge => badge.value) }],
    chart: {
        height: 350,
        type: 'bar',
        zoom: { enabled: false },
    },
    toolbar: { show: false },
    tooltip: { enabled: false },
    plotOptions: {
        bar: {
            borderRadius: 10,
            dataLabels: { position: 'top', show: false },
        }
    },
    dataLabels: {
        enabled: true,
        formatter: (val: number) => "R$ " + val,
        offsetY: 0,
        style: { fontSize: '12px', colors: ["#304758"] }
    },
    grid: {
        xaxis: { lines: { show: false } },
        yaxis: { lines: { show: false } }
    },
    colors: props.badges.map(badge => badge.color),
    xaxis: {
        categories: props.badges.map(badge => badge.name),
        tooltip: { enabled: false },
        axisTicks: { show: false },
    },
    yaxis: {
        axisBorder: { show: true },
        axisTicks: { show: false },
        labels: { show: false }
    },
    title: {
        floating: false,
        offsetY: 330,
        align: 'center',
        style: {
            color: '#cb3cff'
        }
    }
};

onMounted(() => {
    var chart = new ApexCharts(document.querySelector("#chart"), options);
    chart.render();
})
</script>
