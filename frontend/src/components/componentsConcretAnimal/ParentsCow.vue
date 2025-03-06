<template>
    <div>
        <div class="id-title">Родословная</div>
        <table>
            <thead>
                <th></th>
                <th class="id-min-title">Кличка</th>
                <th class="id-min-title inv-number">Инвентарный номер</th>
                <th class="id-min-title">Номер РСХН</th>
                <th class="id-min-title">Номер Селэкс</th>
                <th class="id-min-title">Дата рождения</th>
            </thead>
            <tbody>
                <tr @click="clickFather" :class="{'par-hover': father.Name}">
                    <td class="parent-line">Отец</td>
                    <td>{{ father.Name || 'Нет информации' }}</td>
                    <td>{{ father.InventoryNumber || 'Нет информации' }}</td>
                    <td>{{ father.RSHNNumber || 'Нет информации' }}</td>
                    <td>{{ father.SelecsNumber || 'Нет информации' }}</td>
                    <td v-if="father.BirthDate">{{ dateConverter(father.BirthDate)}}</td><td v-else>Нет информации</td>
                </tr>
                <tr @click="clickMother" :class="{'par-hover': mother.Name}">
                    <td class="parent-line">Мать</td>
                    <td>{{ mother.Name || 'Нет информации' }}</td>
                    <td>{{ mother.InventoryNumber || 'Нет информации' }}</td>
                    <td>{{ mother.RSHNNumber || 'Нет информации' }}</td>
                    <td>{{ mother.SelecsNumber || 'Нет информации' }}</td>
                    <td v-if="mother.BirthDate">{{ dateConverter(mother.BirthDate)}}</td><td v-else>Нет информации</td>
                </tr>
            </tbody>
        </table>
        <div class="idnum-flex">
            <!-- <div class="item-block animal-parent"  @click="clickFather" :class="{'par-hover': father.Name}">
                <div class="id-min-title">Отец</div>
                <div class="link-parent">{{ father.Name || 'Нет информации' }}</div>
            </div>
            <div class="item-block animal-parent" @click="clickMother" :class="{'par-hover': mother.Name}">
                <div class="id-min-title">Мать</div>
                <div class="link-parent">{{ mother.Name || 'Нет информации' }}</div>
            </div> -->
                <div class="item-block">
                <div class="id-min-title">Коэф. инбридинга по родословной</div>
                <div v-if="coeffByFamily || coeffByFamily === 0">{{ coeffByFamily }}</div>
                <div v-else>Нет информации</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Коэф. инбридинга по генотипу</div>
                <div v-if="genetic.InbrindingCoeffByGenotype || genetic.InbrindingCoeffByGenotype === 0">{{ genetic.InbrindingCoeffByGenotype }}</div>
                <div v-else>Нет информации</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    props: {
        mother: {
            type: Object,
            required: true,
        },
        father: {
            type: Object,
            required: true,
        },
        genetic: {
            type: Object,
            required: true,
        },
        coeffByFamily: {
            type: Number,
        }
    },
    methods: {
        clickFather() {
            if (this.father.ID) {
                this.$router.push(`/animals/${this.father.ID}`)
            }
        },
        clickMother() {
            if (this.mother.ID) {
                this.$router.push(`/animals/${this.mother.ID}`)
            }
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        },
    }
}
</script>

<style scoped>
.id-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.idnum-flex {
    display: flex;
    flex-wrap: wrap;
    margin-top: 30px;
}

.item-block {
    width: max-content;
    margin: 0 25px 30px 0;
    padding: 7px;
}

.id-min-title {
    color: grey;
    margin-bottom: 7px;
}

.animal-parent {
    transition: 0.3s;
}

.par-hover {
    cursor: pointer;
    transition: 0.3s;
}

.par-hover:hover {
    background-color: rgb(236, 232, 245);
    border-radius: 10px;
}

table {
    border-collapse: collapse;
}

th {
    width: max-content;
    font-weight: 500;
    text-align: left;
    height: 30px;
}

td {
    padding-right: 30px;
    width: max-content;
    height: 30px;
}

.inv-number {
    width: 180px;
}

.parent-line {
    color: grey;
    padding-left: 10px;
}
</style>