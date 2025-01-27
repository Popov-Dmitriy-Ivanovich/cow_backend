<template>
    <div class="rating-columns">
        <div>
                <div class="rat-title">Оценка КРС по региону</div>
            <div class="rating-item">
                <div class="rating-param">Общая индексная оценка:</div>
                <div v-if="ratings_hoz">{{ ratings_hoz.GeneralValue || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему удою за 305 дней:</div>
                <div>{{ round(ratings_hoz.EbvMilk) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему жиру за 305 дней:</div>
                <div>{{ round(ratings_hoz.EbvFat) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему белку за 305 дней:</div>
                <div>{{ round(ratings_hoz.EbvProtein) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней кратности осеменения:</div>
                <div>{{ round(ratings_hoz.EbvInsemenation) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней длительности сервис-периода:</div>
                <div>{{ round(ratings_hoz.EvbService) || 'Нет информации'}}</div>
            </div>
        </div>

        <div>
            <div class="rat-title">Оценка КРС по стране</div>
            <div class="rating-item">
                <div class="rating-param">Общая индексная оценка:</div>
                <div>{{ round(ratings_reg.GeneralValue) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему удою за 305 дней:</div>
                <div>{{ round(ratings_reg.EbvMilk) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему жиру за 305 дней:</div>
                <div>{{ round(ratings_reg.EbvFat) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по среднему белку за 305 дней:</div>
                <div>{{ round(ratings_reg.EbvProtein) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней кратности осеменения:</div>
                <div>{{ round(ratings_reg.EbvInsemenation) || 'Нет информации'}}</div>
            </div>
            <div class="rating-item">
                <div class="rating-param">EBV по средней длительности сервис-периода:</div>
                <div>{{ round(ratings_reg.EvbService) || 'Нет информации'}}</div>
            </div>
        </div>
        
    </div>
</template>

<script>
export default {
    data() {
        return{ 
            ratings_hoz: {},
            ratings_reg: {},
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/grades`)
        let result = await response.json();
        console.log(result);
        if(result.ByHoz) {
            this.ratings_hoz = result.ByHoz;
        }
        if(result.ByRegion) {
            this.ratings_reg = result.ByRegion;
        }
    },
    methods: {
        round(num) {
            return Math.round(num*100)/100;
        }
    }
}
</script>

<style scoped>
.rat-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding: 0 0 30px 0;
    width: max-content;
}

.sub-title {
    color: red;
}

.rating-item {
    display: flex;
    justify-content: space-between;
    align-items: end;
    margin-right: 30px;
    border-bottom: 1px solid rgb(242, 237, 248);
    width: 370px;
}

.rating-param {
    margin: 5px 0;
    color: rgb(74, 74, 74);
    width: 250px;
}

.rating-columns {
    display: flex;
    font-size: 90%;
}
</style>