<template>
    <div v-if="!isLoading && !isError">
        <table>
            <thead>
                <tr>
                    <th></th>
                    <th colspan="3" class="main-head">Основные</th>
                    <th colspan="3" class="main-head">Резервные</th>
                </tr>
                <tr>
                    <th></th>
                    <th>Ожидаемый EBV</th>
                    <th>Коэффициент инбридинга</th>
                    <th>Имя быка</th>
                    <th>Ожидаемый EBV</th>
                    <th>Коэффициент инбридинга</th>
                    <th>Имя быка</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>Лучший по потенциалу удоя</td>
                    <td v-if="bullsGeneralPicked.length">{{ calcRating(bullsGeneralPicked[0]) }}</td><td v-else></td>
                    <td v-if="bullsGeneralPicked.length">{{ calcInbriding(bullsGeneralPicked[0]) }}</td><td v-else></td>
                    <td v-if="bullsGeneralPicked.length">{{ bullsGeneralPicked[0].name }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ calcRatingReserve(bullsReservedPicked[0]) }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ calcInbridingReserv(bullsReservedPicked[0]) }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ bullsReservedPicked[0].name }}</td><td v-else></td>
                </tr>
                <tr>
                    <td>Лучший по потенциалу % жира</td>
                    <td v-if="bullsGeneralPicked.length">{{ calcRating(bullsGeneralPicked[1]) }}</td><td v-else></td>
                    <td v-if="bullsGeneralPicked.length">{{ calcInbriding(bullsGeneralPicked[1]) }}</td><td v-else></td>
                    <td v-if="bullsGeneralPicked.length">{{ bullsGeneralPicked[1].name }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ calcRatingReserve(bullsReservedPicked[1]) }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ calcInbridingReserv(bullsReservedPicked[1]) }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ bullsReservedPicked[1].name }}</td><td v-else></td>
                </tr>
                <tr>
                    <td>Лучший по потенциалу % белка</td>
                    <td v-if="bullsGeneralPicked.length">{{ calcRating(bullsGeneralPicked[2]) }}</td><td v-else></td>
                    <td v-if="bullsGeneralPicked.length">{{ calcInbriding(bullsGeneralPicked[2]) }}</td><td v-else></td>
                    <td v-if="bullsGeneralPicked.length">{{ bullsGeneralPicked[2].name }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ calcRatingReserve(bullsReservedPicked[2]) }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ calcInbridingReserv(bullsReservedPicked[2]) }}</td><td v-else></td>
                    <td v-if="bullsReservedPicked.length">{{ bullsReservedPicked[2].name }}</td><td v-else></td>
                </tr>
            </tbody>
        </table>
    </div>
    <div v-if="isLoading">Идёт загрузка...</div>
    <div v-if="isError">У животного отсутствуют оценки, подбор пар невозможен.</div>
</template>

<script>
export default {
    data() {
        return {
            bullsGeneral: ['WILLOW-MARSH-CC GABOR-ET', 'KINGS-RANSOM LAURIN EXTRA', 'MAINSTREAM MANIFOLD', 'COPPERTOP DOBERMAN-ET', 'REGANCREST-HHF MAC-ET', 
            'WELCOME PENNTASTIC', 'STANTONS MAYFIELD 8959', 'CLEAR-ECHO ALTAR2-ET', 'LARCREST CONSTRUCTIVE-ET', 'STANTONS LOBO'],
            bullsReserved: ['DINOMI ALTADETROIT-ET', 'MY-JOHN BW MARSHALL ACE-ET', 'END-ROAD O-MAN BRONCO-ET', 'RALMA FOCUS-ET', 'LANGS-TWIN-B CR CASSINO-ET', 
            'RAMA-WAY RAYNOR 218', 'COVE-VIEW MOSCOW APPLE BOY', 'JENNY-LOU MRSHL TOYSTORY-ET', 'R-E-W BUCKEYE-ET', 'RIVER-BRIDGE CO-OP TROY-ET'],
            bullsGeneralPicked: [],
            bullsReservedPicked: [],
            cowRSHN: '',
            cowRating: 0,
            cowInbriding: 0,

            isLoading: false,
            isError: false,
        }
    },
    methods: {
        async getCowRSHN() {
            let id = this.$route.path.split('/')[2];
            let response = await fetch(`/api/cows/${id}`);
            let result = await response.json();
            this.cowRSHN = result.RSHNNumber;
        },
        getBulls() {
            for (let i = -1; i >= -3; i--) {
                let obj = {};
                obj.number = Number(this.cowRSHN.at(i));
                obj.name = this.bullsGeneral[obj.number];
                this.bullsGeneralPicked.push(obj);
            }
            for (let i = -2; i >= -4; i--) {
                let obj = {};
                obj.number = Number(this.cowRSHN.at(i));
                obj.name = this.bullsReserved[obj.number];
                this.bullsReservedPicked.push(obj);
            }
        },
        async getRating() {
            let id = this.$route.path.split('/')[2];
            let response = await fetch(`/api/cows/${id}/grades`);
            let result = await response.json();
            if (result.ByRegion) {
                this.cowRating = Math.round(result.ByRegion.GeneralValue*100)/100;
            } else {
                this.isError = true;
            }
        },
        calcRating(bull) {
            return Math.round((this.cowRating + this.cowRating*(0.025 + 1.2*(bull.number)/100))*100)/100;
        },
        calcRatingReserve(bull) {
            return Math.round((this.cowRating + this.cowRating*(0.025 + (bull.number)/100))*100)/100;
        },
        async getInbriding() {
            let id = this.$route.path.split('/')[2];
            let response = await fetch(`/api/cows/${id}/genetic`);
            let result = await response.json();
            console.log(result);
            if(result) {
                if (result.InbrindingCoeffByGenotype === 0) {
                    this.cowInbriding = 0.128;
                } else {
                    this.cowInbriding = result.InbrindingCoeffByGenotype;
                }
            } else {
                this.cowInbriding = 0;
            }
        },
        calcInbriding(bull) {
            let result;
            if (this.cowInbriding < 0.25) {
                result = Math.round((this.cowInbriding + this.cowInbriding*(0.04 + (bull.number*4)/100))*10000)/10000;
            } else {
                result = Math.round((0.3 - this.cowInbriding*(0.04 + (bull.number*5)/100))*10000)/10000;
            }
            if (result < 0.3) {
                return result;
            } else {
                return 0.29;
            }
        },
        calcInbridingReserv(bull) {
            let result;
            if (this.cowInbriding < 0.25) {
                result = Math.round((this.cowInbriding + this.cowInbriding*(0.05 + (bull.number*5)/100))*10000)/10000;
            } else {
                result = Math.round((0.3 - this.cowInbriding*(0.05 + (bull.number*4)/100))*10000)/10000;
            }
            if (result < 0.3) {
                return result;
            } else {
                return 0.29;
            }
        },
    },
    async mounted() {
        this.isLoading = true;
        await this.getCowRSHN();
        this.getBulls();
        await this.getRating();
        await this.getInbriding();
        this.isLoading = false;
    }
}
</script>

<style scoped>
table {
    font-size: 90%;
    border-collapse: collapse;
}

th {
    font-weight: normal;
    text-align: center;
    padding: 5px 4px;
}

td {
    padding: 5px 0 3px 0;
    text-align: center;
}

th, td {
    border: 1px solid rgb(179, 190, 187);
}

.main-head {
    padding: 16px 0;
}
</style>