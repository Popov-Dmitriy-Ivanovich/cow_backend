<template>
    <div>
        <!-- <div class="prev-chart" @click="toPrev">🠔 Назад</div> -->
        <!-- <ChartYear/> -->
        <apexchart id="analit_click" width="500" type="bar" :options="optionsClick" :series="seriesClick" ref="analit" @dataPointSelection="clickHandler"></apexchart>

    </div>
</template>

<script>
// import ChartYear from '@/components/analyticsComponents/ChartYear.vue';

export default {
    components: {
        // ChartYear
    },
    methods: {
        toPrev() {
            this.$router.push(`/analytics`);
        },
    },
    data() {
        return {
            optionsClick: {
                chart: {
                    id: 'analit_click',
                    stacked: true,
                    zoom: {
                        enabled: false,
                    }
                },
                xaxis: {
                    categories: ['Минимальный индекс', 'Средний индекс', 'Максимальный индекс'],
                    dataLabels: {
                        enabled: true,
                        style: {
                            fontSize: '5px',
                        }
                    },
                    labels: {
                        style: {
                            fontSize: '8px',
                        },
                        hideOverlappingLabels: true,
                        trim: true,
                    }
                },

                colors: ['#63d9cb','#6e5add','#75a2e7'],
                title: {
                    text: `Число голов в `,
                    align: 'center',
                    style: {
                        fontSize:  '12px',
                    },
                },
                tooltip: {
                    enabled: false,
                }
            },
            seriesClick: [{
                data: [JSON.parse(localStorage.getItem('currentHoz')).MinCount, JSON.parse(localStorage.getItem('currentHoz')).AvgCount, JSON.parse(localStorage.getItem('currentHoz')).MaxCount]
            }],
            
            newX: [],
            result: [],
        }
    },
    mounted() {
        this.optionsClick.xaxis.categories = [];

        let name = '';
        if (JSON.parse(localStorage.getItem('currentHoz')).Farm) name = ['Число голов в', JSON.parse(localStorage.getItem('currentHoz')).Farm.Name];
        else name = 'Число голов во всем регионе';
        this.optionsClick.title.text = name;
        
        this.optionsClick.xaxis.categories.push('Минимальный индекс ' + JSON.parse(localStorage.getItem('currentHoz')).MinIndex);
        this.optionsClick.xaxis.categories.push('Средний индекс ' + JSON.parse(localStorage.getItem('currentHoz')).AvgIndex);
        this.optionsClick.xaxis.categories.push('Максимальный индекс ' + JSON.parse(localStorage.getItem('currentHoz')).MaxIndex);
    }
}
</script>

<style scoped>
.prev-chart {
    font-family: Open Sans, sans-serif;
    margin-top: 30px;
    margin-left: 20px;
    color:rgb(10, 113, 75);
    padding-bottom: 20px;
    cursor: pointer;
    transition: 0.3s;
}

.prev-chart:hover {
    color: rgb(63, 205, 132);
    padding-left: 10px;
    width: max-content;
}

</style>