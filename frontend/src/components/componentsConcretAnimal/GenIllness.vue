<template>
    <div>
        <div class="vet-title">Моногенные заболевания</div>
        <div class="parent-table">
            <table class="vet-table">
                <thead>
                    <tr class="vet-header">
                        <th>Патология</th>
                        <th>Статус</th>
                        <th>Описание</th>
                    </tr>
                </thead>
                <tbody class="vet-tablebody" v-if="cow_info.length">
                    <tr v-for="item in cow_info" :key="item.Name">
                        <td>{{ item.Name || 'Нет информации'}}</td>
                        <td>{{ item.Status || 'Нет информации'}}</td>
                        <td>{{ item.Description || 'Нет информации'}}</td>
                    </tr>
                </tbody>
                <tbody v-else class="vet-tablebody">
                    <tr>
                        <td>Нет информации</td>
                        <td>Нет информации</td>
                        <td>Нет информации</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            illnesses: {},
            cow_info: [],
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/monogenetic_illnesses`);
        let result = await response.json();
        this.illnesses = result;
        let response1 = await fetch(`/api/cows/${cow_id}/genetic`)
        let result1 = await response1.json();

        for( let i = 0; i < this.illnesses.length; i++) {
            let obj = {
                Name: this.illnesses[i].Name,
                Status: 'Нет информации',
                Description: this.illnesses[i].Description,
            }
            if(result1) {
                for (let j = 0; j < result1.GeneticIllnessesData.length; j++) {
                    if (this.illnesses[i].Name === result1.GeneticIllnessesData[j].Illness.Name) {
                        if(result1.GeneticIllnessesData[j].Status){
                            obj.Status = result1.GeneticIllnessesData[j].Status.Status;
                        } else {
                            obj.Status = result1.GeneticIllnessesData[j].Status;
                        }
                    }
                }
            }
            this.cow_info.push(obj);
        }
    },
}    
</script>
    
<style scoped>
.vet-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding-bottom: 30px;
    width: max-content;
}
    
.parent-table {
    width: 50vw;
    overflow-x: auto;
}
    
.vet-table {
    margin-bottom: 30px;
    text-align: left;
    
}
    
th {
    font-weight: normal;
}

td {
    padding: 5px 15px 5px 0;
}
    
.vet-header {
    color: grey;
}
    
.vet-header th {
    padding-right: 30px;
    padding-bottom: 15px;
}
    
.vet-tablebody {
    text-align: left;
}
    
.parent-table::-webkit-scrollbar {
    height: 12px;
}
    
.parent-table::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}
    
.parent-table::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

</style>