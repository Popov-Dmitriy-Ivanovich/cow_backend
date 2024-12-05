<template>
    <div>
        <div class="lineage-title">Родословное дерево</div>
        <div class="parent-lineagetree">
            <div class="column">
                <div class="main-cow animal-block">
                    <div class="parent-name">{{ current_cow.Name || 'Имя'}}</div>
                    <div>{{ current_cow.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ current_cow.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ current_cow.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ current_cow.BreedName || 'Порода'}}</div>
                </div>
            </div>
            <div class="column">
                <div class="main-cow__mother animal-block">
                    <div class="parent-name">{{ mother.Name || 'Имя'}}</div>
                    <div>{{ mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="white-block-parent"></div>
                <div class="main-cow__father animal-block">
                    <div class="parent-name">{{ father.Name || 'Имя'}}</div>
                    <div>{{ father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ father.BreedName || 'Порода'}}</div>
                </div>
            </div>
            
            <div class="column">
                <div class="main-cow__mother__mother animal-block">
                    <div class="parent-name">{{ grandmother_mother.Name || 'Имя'}}</div>
                    <div>{{ grandmother_mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ grandmother_mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ grandmother_mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ grandmother_mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="white-block-grandparent_mini"></div>
                <div class="main-cow__mother__father animal-block">
                    <div class="parent-name">{{ grandfather_mother.Name || 'Имя'}}</div>
                    <div>{{ grandfather_mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ grandfather_mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ grandfather_mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ grandfather_mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="white-block-grandparent"></div>
                <div class="main-cow__father__mother animal-block">
                    <div class="parent-name">{{ grandmother_father.Name || 'Имя'}}</div>
                    <div>{{ grandmother_father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ grandmother_father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ grandmother_father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ grandmother_father.BreedName || 'Порода'}}</div>
                </div>
                <div class="white-block-grandparent_mini"></div>
                <div class="main-cow__father__father animal-block">
                    <div class="parent-name">{{ grandfather_father.Name || 'Имя'}}</div>
                    <div>{{ grandfather_father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ grandfather_father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ grandfather_father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ grandfather_father.BreedName || 'Порода'}}</div>
                </div>
            </div>

            <div class="column">
                <div class="main-cow__mother__mother__mother animal-block">
                    <div class="parent-name">{{ mother_grandmother_mother.Name || 'Имя'}}</div>
                    <div>{{ mother_grandmother_mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ mother_grandmother_mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ mother_grandmother_mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ mother_grandmother_mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__mother__mother__father animal-block">
                    <div class="parent-name">{{ mother_grandmother_father.Name || 'Имя'}}</div>
                    <div>{{ mother_grandmother_father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ mother_grandmother_father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ mother_grandmother_father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ mother_grandmother_father.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__mother__father__mother animal-block">
                    <div class="parent-name">{{ father_grandmother_mother.Name || 'Имя'}}</div>
                    <div>{{ father_grandmother_mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ father_grandmother_mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ father_grandmother_mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ father_grandmother_mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__mother__father__father animal-block">
                    <div class="parent-name">{{ father_grandmother_father.Name || 'Имя'}}</div>
                    <div>{{ father_grandmother_father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ father_grandmother_father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ father_grandmother_father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ father_grandmother_father.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__father__mother__mother animal-block">
                    <div class="parent-name">{{ mother_grandfather_mother.Name || 'Имя'}}</div>
                    <div>{{ mother_grandfather_mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ mother_grandfather_mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ mother_grandfather_mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ mother_grandfather_mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__father__mother__father animal-block">
                    <div class="parent-name">{{ father_grandfather_mother.Name || 'Имя'}}</div>
                    <div>{{ father_grandfather_mother.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ father_grandfather_mother.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ father_grandfather_mother.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ father_grandfather_mother.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__father__father__mother animal-block">
                    <div class="parent-name">{{ mother_grandfather_father.Name || 'Имя'}}</div>
                    <div>{{ mother_grandfather_father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ mother_grandfather_father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ mother_grandfather_father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ mother_grandfather_father.BreedName || 'Порода'}}</div>
                </div>
                <div class="main-cow__father__father__father animal-block">
                    <div class="parent-name">{{ father_grandfather_father.Name || 'Имя'}}</div>
                    <div>{{ father_grandfather_father.IdentificationNumber || 'Идент. номер'}}</div>
                    <div>{{ father_grandfather_father.RSHNNumber || 'РСХН'}}</div>
                    <div>{{ father_grandfather_father.BirthDate || 'Дата рождения'}}</div>
                    <div>{{ father_grandfather_father.BreedName || 'Порода'}}</div>
                </div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            current_cow: {},

            mother: {},
            father: {},

            grandmother_mother: {},
            grandfather_mother: {},
            grandmother_father: {},
            grandfather_father: {},

            mother_grandmother_mother: {},
            mother_grandmother_father: {},
            father_grandmother_mother: {},
            father_grandmother_father: {},

            mother_grandfather_mother: {},
            father_grandfather_mother: {},
            mother_grandfather_father: {},
            father_grandfather_father: {},
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}`);
        let result = await response.json();
        this.current_cow = result;
        this.mother = this.current_cow.Mother;
        this.father = this.current_cow.Father;

        console.log(this.current_cow);

        let response1 = await fetch(`/api/cows/${this.mother.ID}`);
        let result1 = await response1.json();
        this.grandfather_mother = result1.Father;
        this.grandmother_mother = result1.Mother;

        let response2 = await fetch(`/api/cows/${this.father.ID}`);
        let result2 = await response2.json();
        this.grandfather_father = result2.Father;
        this.grandmother_father = result2.Mother;

        if (this.grandfather_mother && this.grandmother_mother && this.grandfather_father && this.grandmother_father) {
            let response3 = await fetch(`/api/cows/${this.grandmother_mother.ID}`);
            let result3 = await response3.json();
            this.mother_grandmother_mother = result3.Mother;
            this.mother_grandmother_father = result3.Father;

            let response4 = await fetch(`/api/cows/${this.grandmother_father.ID}`);
            let result4 = await response4.json();
            this.father_grandmother_mother = result4.Mother;
            this.father_grandmother_father = result4.Father;

            let response5 = await fetch(`/api/cows/${this.grandfather_mother.ID}`);
            let result5 = await response5.json();
            this.mother_grandfather_mother = result5.Mother;
            this.father_grandfather_mother = result5.Father;

            let response6 = await fetch(`/api/cows/${this.grandfather_father.ID}`);
            let result6 = await response6.json();
            this.mother_grandfather_father = result6.Mother;
            this.father_grandfather_father = result6.Father;
        }
    },
    methods: {
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
.lineage-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.lineage-column {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.animal-block {
    width: 160px;
    height: 90px;
    background-color: rgb(236, 235, 244);
    margin: 12px 15px;
    border-radius: 10px;
    font-size: 80%;
    padding: 6px 7px;
}

.animal-block div {
    padding: 2px 3px;
}

.parent-name {
    color: black;
    border-bottom: 1px solid rgb(188, 185, 194);
}

.white-block-parent {
    width: 150px;
    height: 200px;
}

.white-block-grandparent {
    width: 150px;
    height: 100px;
}

.white-block-grandparent_mini {
    width: 150px;
    height: 50px;
}

.parent-lineagetree {
    /* display: grid;
    grid-template-columns: repeat(4, 160px);
    grid-template-rows: repeat(9, 100px);
    grid-template-areas: 
    "W W W G"
    "W W C H"
    "W A D I"
    "W W W J"
    "M W W W"
    "W W W K"
    "W B E L"
    "W W F M"
    "W W W N"; */
    display: flex;
    align-items: center;
}

.main-cow {
    grid-area: M;
}
.main-cow__mother {
    grid-area: A;
}
.main-cow__father {
    grid-area: B;
}
.main-cow__mother__mother {
    grid-area: C;
}
.main-cow__mother__father {
    grid-area: D;
}
.main-cow__father__mother {
    grid-area: E;
}
.main-cow__father__father {
    grid-area: F;
}
.main-cow__mother__mother__mother {
    grid-area: G;
}
.main-cow__mother__mother__father {
    grid-area: H;
}
.main-cow__mother__father__mother {
    grid-area: I;
}
.main-cow__mother__father__father {
    grid-area: J;
}
.main-cow__father__mother__mother {
    grid-area: K;
}
.main-cow__father__mother__father {
    grid-area: L;
}
.main-cow__father__father__mother {
    grid-area: M;
}
.main-cow__father__father__father {
    grid-area: N;
}
</style>