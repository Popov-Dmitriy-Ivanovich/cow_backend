<template>
    <div>
        <div class="animal-line" @click="$router.push(`/animals/${animal_item.ID}`)" :class="{'is-approved': animal_item.Approved==1}">
            <div class="animal-rshn">{{ animal_item.RSHNNumber }}</div>
            <div class="animal-inv">{{ animal_item.InventoryNumber }}</div>
            <div class="animal-name">{{ animal_item.Name }}</div>
            <div class="animal-hoz">{{ animal_item.FarmGroupName }}</div>
            <div class="animal-bdate">{{ dateConverter(bdate(animal_item.BirthDate)) }}</div>
            <div class="animal-genfact">{{ isGen(animal_item.genotyped) }}</div>
            <div v-if="animal_item.DepartDate" class="animal-dateout">{{ dateConverter(animal_item.DepartDate) }}</div>
            <div v-if="animal_item.IsDead===true || animal_item.IsDead===false" class="animal_dead">{{ isGen(animal_item.IsDead) }}</div>
            <div v-if="animal_item.BreedName" class="animal-breed">{{ animal_item.BreedName }}</div>
            <div v-if="animal_item.GenotypingDate" class="animal-dategen">{{ dateConverter(animal_item.GenotypingDate) }}</div>
            <!-- contol milking -->
            <div v-if="animal_item.Exterior" class="animal-exterior">{{ animal_item.Exterior }}</div>
            <!-- <div v-if="animal_item.InsemenationDate" class="animal-dateosem">{{ animal_item.InsemenationDate }}</div> -->
            <!-- <div v-if="animal_item.CalvingDate" class="animal-dateotel">{{ animal_item.calvingDate }}</div> -->
            <div v-if="animal_item.IsTwins===true || animal_item.IsTwins===false" class="animal-genfact">{{ isGen(animal_item.IsTwins) }}</div>
            <div v-if="animal_item.IsTwins===true || animal_item.IsTwins===false" class="animal-genfact">{{ isGen(animal_item.IsStillBorn) }}</div>
            <div v-if="animal_item.IsAborted===true || animal_item.IsAborted===false" class="animal-genfact">{{ isGen(animal_item.IsAborted) }}</div>
            <div v-if="animal_item.BirkingDate" class="animal-datebirk">{{ dateConverter(animal_item.Birkingdate) }}</div>
            <div v-if="animal_item.InbrindingCoeffByFamily" class="animal-krod">{{ animal_item.InbrindingCoeffByFamily }}</div>
            <!-- <div v-if="animal_item." class="animal-kfen">{{ animal_item.INBRID_FENOTYPE }}</div> -->
        </div>
    </div>
</template>

<script>
export default {
    props: {
        animal_item: {
            type: Object,
            Required: true
        },
        filters: {
            type: Object,
        }
    },
    methods: {
        isGen(a) {
            if(a) return 'Да';
            else return 'Нет';
        },
        bdate(birth) {
            return birth.split('T')[0];
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        }
    }
}
</script>

<style scoped>
.animal-line {
    display: flex;
    padding: 10px 0;
    transition: 0.3s;
    cursor: pointer;
    width: max-content;
    color: grey;
}

.is-approved {
    color: black;
}

.animal-line:hover {
    background-color: rgb(235, 233, 245);
}

.animal-line div {
    padding: 0 10px;
}

.animal-selex {
    width: 140px;
}

.animal-inv {
    width: 115px;
}

.animal-rshn {
    width: 100px;
}

.animal-name {
    width: 140px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.animal-hoz {
    width: 250px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.animal-bdate {
    width: 120px;
}

.animal-genfact {
    width: 120px;
}

.animal-dateout {
    width: 120px;
}

.animal-breed, .animal-dategen, .animal-datemilking, 
.animal-dateosem, .animal-dateotel, .animal-datebirk {
    width: 150px;
}

.animal-exterior {
    width: 100px;
}

.animal-krod, .animal-kfen {
    width: 100px;
}
</style>