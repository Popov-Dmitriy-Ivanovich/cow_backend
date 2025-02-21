<template>
    <div class="lineage-title">Родословное дерево</div>
    <div v-if="!isLoading" class="lineage">
        <div class="lineage-back"></div>
        <div class="parent-lineagetree">
            <div class="column">
                <div class="main-cow animal-block">
                    <div class="name-flex"><div class="parent-name">{{ current_cow.Name || ' '}}</div>
                        <div v-if="current_cow.SexId===4 || current_cow.SexId===2">&#9792;</div>
                        <div v-else>♂</div>
                    </div>
                    <div>{{ current_cow.IdentificationNumber || ' '}}</div>
                    <div>{{ current_cow.RSHNNumber || ' '}}</div>
                    <div>{{ current_cow.BirthDate || ' '}}</div>
                    <div>{{ current_cow.BreedName || ' '}}</div>
                </div>
            </div>
            <div class="column">
                <div class="main-cow__father animal-block isParent" v-if="Object.keys(father).length" @click="clickToParent(father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ father.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ father.IdentificationNumber || ' '}}</div>
                    <div>{{ father.RSHNNumber || ' '}}</div>
                    <div>{{ father.BirthDate || ' '}}</div>
                    <div>{{ father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="white-block-parent"></div>
                <div class="main-cow__mother animal-block isParent" v-if="Object.keys(mother).length" @click="clickToParent(mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ mother.Name || ' '}}</div><div>&#9792;</div></div>
                    <div>{{ mother.IdentificationNumber || ' '}}</div>
                    <div>{{ mother.RSHNNumber || ' '}}</div>
                    <div>{{ mother.BirthDate || ' '}}</div>
                    <div>{{ mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
            </div>
            
            <div class="column">
                <div class="main-cow__father__father animal-block isParent" v-if="Object.keys(grandfather_father).length"  @click="clickToParent(grandfather_father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ grandfather_father.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ grandfather_father.IdentificationNumber || ' '}}</div>
                    <div>{{ grandfather_father.RSHNNumber || ' '}}</div>
                    <div>{{ grandfather_father.BirthDate || ' '}}</div>
                    <div>{{ grandfather_father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="white-block-grandparent_mini"></div>
                <div class="main-cow__father__mother animal-block isParent" v-if="Object.keys(grandmother_father).length"  @click="clickToParent(grandmother_father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ grandmother_father.Name || ' '}}</div><div>♀</div></div>
                    <div>{{ grandmother_father.IdentificationNumber || ' '}}</div>
                    <div>{{ grandmother_father.RSHNNumber || ' '}}</div>
                    <div>{{ grandmother_father.BirthDate || ' '}}</div>
                    <div>{{ grandmother_father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="white-block-grandparent"></div> 
                <div class="main-cow__mother__father animal-block isParent" v-if="Object.keys(grandfather_mother).length"  @click="clickToParent(grandfather_mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ grandfather_mother.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ grandfather_mother.IdentificationNumber || ' '}}</div>
                    <div>{{ grandfather_mother.RSHNNumber || ' '}}</div>
                    <div>{{ grandfather_mother.BirthDate || ' '}}</div>
                    <div>{{ grandfather_mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="white-block-grandparent_mini"></div>
                <div class="main-cow__mother__mother animal-block isParent" v-if="Object.keys(grandmother_mother).length"  @click="clickToParent(grandmother_mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ grandmother_mother.Name || ' '}}</div><div>♀</div></div>
                    <div>{{ grandmother_mother.IdentificationNumber || ' '}}</div>
                    <div>{{ grandmother_mother.RSHNNumber || ' '}}</div>
                    <div>{{ grandmother_mother.BirthDate || ' '}}</div>
                    <div>{{ grandmother_mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
            </div>

            <div class="column">
                <div class="main-cow__father__father__father animal-block isParent" v-if="Object.keys(father_grandfather_father).length" @click="clickToParent(father_grandfather_father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ father_grandfather_father.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ father_grandfather_father.IdentificationNumber || ' '}}</div>
                    <div>{{ father_grandfather_father.RSHNNumber || ' '}}</div>
                    <div>{{ father_grandfather_father.BirthDate || ' '}}</div>
                    <div>{{ father_grandfather_father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__father__father__mother animal-block isParent" v-if="Object.keys(mother_grandfather_father).length" @click="clickToParent(mother_grandfather_father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ mother_grandfather_father.Name || ' '}}</div><div>♀</div></div>
                    <div>{{ mother_grandfather_father.IdentificationNumber || ' '}}</div>
                    <div>{{ mother_grandfather_father.RSHNNumber || ' '}}</div>
                    <div>{{ mother_grandfather_father.BirthDate || ' '}}</div>
                    <div>{{ mother_grandfather_father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__mother__father__father animal-block isParent" v-if="Object.keys(father_grandmother_father).length" @click="clickToParent(father_grandmother_father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ father_grandmother_father.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ father_grandmother_father.IdentificationNumber || ' '}}</div>
                    <div>{{ father_grandmother_father.RSHNNumber || ' '}}</div>
                    <div>{{ father_grandmother_father.BirthDate || ' '}}</div>
                    <div>{{ father_grandmother_father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__mother__father__mother animal-block isParent" v-if="Object.keys(father_grandmother_mother).length" @click="clickToParent(father_grandmother_mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ father_grandmother_mother.Name || ' '}}</div><div>♀</div></div>
                    <div>{{ father_grandmother_mother.IdentificationNumber || ' '}}</div>
                    <div>{{ father_grandmother_mother.RSHNNumber || ' '}}</div>
                    <div>{{ father_grandmother_mother.BirthDate || ' '}}</div>
                    <div>{{ father_grandmother_mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__father__mother__father animal-block isParent" v-if="Object.keys(father_grandfather_mother).length" @click="clickToParent(father_grandfather_mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ father_grandfather_mother.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ father_grandfather_mother.IdentificationNumber || ' '}}</div>
                    <div>{{ father_grandfather_mother.RSHNNumber || ' '}}</div>
                    <div>{{ father_grandfather_mother.BirthDate || ' '}}</div>
                    <div>{{ father_grandfather_mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__father__mother__mother animal-block isParent" v-if="Object.keys(mother_grandfather_mother).length" @click="clickToParent(mother_grandfather_mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ mother_grandfather_mother.Name || ' '}}</div><div>♀</div></div>
                    <div>{{ mother_grandfather_mother.IdentificationNumber || ' '}}</div>
                    <div>{{ mother_grandfather_mother.RSHNNumber || ' '}}</div>
                    <div>{{ mother_grandfather_mother.BirthDate || ' '}}</div>
                    <div>{{ mother_grandfather_mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__mother__mother__father animal-block isParent" v-if="Object.keys(mother_grandmother_father).length" @click="clickToParent(mother_grandmother_father.ID)">
                    <div class="name-flex"><div class="parent-name">{{ mother_grandmother_father.Name || ' '}}</div><div>♂</div></div>
                    <div>{{ mother_grandmother_father.IdentificationNumber || ' '}}</div>
                    <div>{{ mother_grandmother_father.RSHNNumber || ' '}}</div>
                    <div>{{ mother_grandmother_father.BirthDate || ' '}}</div>
                    <div>{{ mother_grandmother_father.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
                <div class="main-cow__mother__mother__mother animal-block isParent" v-if="Object.keys(mother_grandmother_mother).length" @click="clickToParent(mother_grandmother_mother.ID)">
                    <div class="name-flex"><div class="parent-name">{{ mother_grandmother_mother.Name || ' '}}</div><div>♀</div></div>
                    <div>{{ mother_grandmother_mother.IdentificationNumber || ' '}}</div>
                    <div>{{ mother_grandmother_mother.RSHNNumber || ' '}}</div>
                    <div>{{ mother_grandmother_mother.BirthDate || ' '}}</div>
                    <div>{{ mother_grandmother_mother.BreedName || ' '}}</div>
                </div>
                <div v-else class="animal-block animal-block-noinfo">
                    <div>Нет информации</div>
                </div>
            </div>
        </div>
    </div>
    <div v-else>Идёт загрузка...</div>
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

            isLoading: false,
        }
    },
    async created() {
        this.isLoading = true;
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        await this.fetchInfo(cow_id);
        this.isLoading = false;
    },
    methods: {
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        },
        async fetchInfo(param) {
            this.current_cow = {};

            this.mother = {};
            this.father = {};

            this.grandmother_mother = {};
            this.grandfather_mother = {};
            this.grandmother_father = {};
            this.grandfather_father = {};

            this.mother_grandmother_mother = {};
            this.mother_grandmother_father = {};
            this.father_grandmother_mother = {};
            this.father_grandmother_father = {};

            this.mother_grandfather_mother = {};
            this.father_grandfather_mother = {};
            this.mother_grandfather_father = {};
            this.father_grandfather_father = {};

            let response = await fetch(`/api/cows/${param}`);
            let result = await response.json();
            this.current_cow = result;

            await this.fetchMother();
            await this.fetchFather();

            if (Object.keys(this.grandmother_mother).length && this.grandmother_mother.ID) {  //бабушка по матери
                let response3 = await fetch(`/api/cows/${this.grandmother_mother.ID}`);
                let result3 = await response3.json();
                this.grandmother_mother.BreedName = result3.BreedName;

                if(result3.Mother && result3.Mother.ID) {
                    this.mother_grandmother_mother = result3.Mother;
                    await this.fetchBreed(this.mother_grandmother_mother.BreedId, this.mother_grandmother_mother);
                }
                if(result3.Father && result3.Father.ID) {
                    this.mother_grandmother_father = result3.Father;
                    await this.fetchBreed(this.mother_grandmother_father.BreedId, this.mother_grandmother_father)
                }
            }
            if (Object.keys(this.grandmother_father).length && this.grandmother_father.ID) {  //бабушка по отцу
                let response4 = await fetch(`/api/cows/${this.grandmother_father.ID}`);
                let result4 = await response4.json();
                this.grandmother_father.BreedName = result4.BreedName;

                if(result4.Mother && result4.Mother.ID) {
                    this.father_grandmother_mother = result4.Mother;
                    await this.fetchBreed(this.father_grandmother_mother.BreedId, this.father_grandmother_mother);
                }
                if(result4.Father && result4.Father.ID) {
                    this.father_grandmother_father = result4.Father;
                    await this.fetchBreed(this.father_grandmother_father.BreedId, this.father_grandmother_father);
                }
            }

            if (Object.keys(this.grandfather_mother).length &&this.grandfather_mother.ID) {  //дедушка по матери
                let response5 = await fetch(`/api/cows/${this.grandfather_mother.ID}`);
                let result5 = await response5.json();
                this.grandfather_mother.BreedName = result5.BreedName;

                if(result5.Mother && result5.Mother.ID) {
                    this.mother_grandfather_mother = result5.Mother;
                    await this.fetchBreed(this.mother_grandfather_mother.BreedId, this.mother_grandfather_mother);
                }
                if(result5.Father && result5.Father.ID) {
                    this.father_grandfather_mother = result5.Father;
                    await this.fetchBreed(this.father_grandfather_mother.BreedId, this.father_grandfather_mother);
                }
            }

            if (Object.keys(this.grandfather_father).length && this.grandfather_father.ID) {  //дедушка по отцу
                let response6 = await fetch(`/api/cows/${this.grandfather_father.ID}`);
                let result6 = await response6.json();
                this.grandfather_father.BreedName = result6.BreedName;

                if(result6.Mother && result6.Mother.ID) {
                    this.mother_grandfather_father = result6.Mother;
                    await this.fetchBreed(this.mother_grandfather_father.BreedId, this.mother_grandfather_father);
                }
                if(result6.Father && result6.Father.ID) {
                    this.father_grandfather_father = result6.Father;
                    await this.fetchBreed(this.father_grandfather_father.BreedId, this.father_grandfather_father);
                }
            } 
        },
        async fetchMother() {
            if (this.current_cow.Mother && this.current_cow.Mother.ID) {
                this.mother = this.current_cow.Mother;

                let response1 = await fetch(`/api/cows/${this.mother.ID}`);
                let result1 = await response1.json();
                this.mother.BreedName = result1.BreedName;
                if(result1.Father && result1.Father.ID) this.grandfather_mother = result1.Father;
                if(result1.Mother && result1.Mother.ID) this.grandmother_mother = result1.Mother;
            }
        },
        async fetchFather() {
            if (this.current_cow.Father && this.current_cow.Father.ID) {
                this.father = this.current_cow.Father;

                let response2 = await fetch(`/api/cows/${this.father.ID}`);
                let result2 = await response2.json();
                this.father.BreedName = result2.BreedName;
                if(result2.Father && result2.Father.ID) this.grandfather_father = result2.Father;
                if (result2.Mother && result2.Mother.ID) this.grandmother_father = result2.Mother;
            }
        },
        async fetchBreed(breedID, who) {
            if(breedID) {
                let response = await fetch(`/api/breeds/${breedID}`);
                let result = await response.json();
                who.BreedName = result.Name;
            }
        },
        clickToParent(id) {
            if(id) this.$router.push(`/animals/${id}`);
        },
    },
    watch: {
        async $route(new_val) {
            await this.fetchInfo(new_val.params.id);
        }
    }
}
</script>

<style scoped>
.lineage {
    position: relative;
}

.lineage-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.lineage-back {
    position: absolute;
    width: 830px;
    height: 100%;
    background-image: url('../../../static/lineage.png');
    background-size: no-repeat;
    background-position: center;
    background-size: cover;
}

.lineage-column {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.animal-block, .animal-block-noinfo {
    width: 160px;
    height: 100px;
    background-color: rgb(236, 235, 244);
    border: 1px solid rgb(236, 235, 244);
    margin: 12px 15px;
    border-radius: 10px;
    font-size: 80%;
    padding: 10px 12px;
    transition: 0.3s;
}

.animal-block-noinfo {
    background-color: rgb(249, 248, 253);
    display: flex;
    justify-content: center;
    align-items: center;
}

.isParent {
    cursor: pointer;
}

.isParent:hover {
    background-color: rgb(250, 250, 254);
    border: 1px solid rgb(184, 180, 208);
}

.animal-block div {
    padding: 2px 3px;
}

.parent-name {
    color: black;
    font-size: 105%;
    font-weight: bold;
    text-overflow: ellipsis;
    overflow: hidden;
}

.white-block-parent {
    width: 150px;
    height: 390px;
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
    display: flex;
    align-items: center;
}

.name-flex {
    display: flex;
    border-bottom: 1px solid rgb(188, 185, 194);
}

/* .main-cow__father__father {
    margin-bottom: 90px;
}

.main-cow__father__mother {
    margin-bottom: 60px;
}

.main-cow__mother__father {
    margin-bottom: 80px;
} */

.column {
    z-index: 4;
}
</style>