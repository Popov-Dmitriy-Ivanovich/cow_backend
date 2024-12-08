<template>
    <div>
        <div class="extrat-title">Оценка экстерьера</div>
        <div class="exterior-title">
            <div>Общая оценка экстерьера: {{ cow_info.Rating }}</div>
            <div class="krs-photo">
                <img width="100%" :src="logo">
            </div>
            
        </div>
        <div>
            <div class="ext-9mark-title" @click="show9Mark">> Признаки с 9-бальной оценкой</div>
            <div class="ext-9mark" :class="{'ext-hide': !isVisible9Mark}">
                <div class="column">
                    <div class="mark-with-num">
                        <div class="ext-param">Ширина груди: </div>
                        <div>{{ cow_info.ChestWidth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Ширина таза: </div>
                        <div>{{ cow_info.PelvicWidth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Высота в крестце: </div>
                        <div>{{ cow_info.SacrumHeight }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Глубина туловища: </div>
                        <div>{{ cow_info.BodyDepth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Тип экстерьера: </div>
                        <div>{{ cow_info.ExteriorType }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Качество костяка/скакательный сустав (вид сзади): </div>
                        <div>{{ cow_info.BoneQHockJointRear }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Угол наклона крестца: </div>
                        <div>{{ cow_info.SacrumAngle }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Длина крестца: </div>
                        <div>{{ cow_info.AcrumLength }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Линия верха: </div>
                        <div>{{ cow_info.TopLine }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Угол копыта: </div>
                        <div>{{ cow_info.HoofAngle }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Постановка задних ног (вид сбоку): </div>
                        <div>{{ cow_info.HindLegPosSide }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Постановка задних ног (вид сзади): </div>
                        <div>{{ cow_info.HindLegPosRead }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Постановка передних ног (вид спереди): </div>
                        <div>{{ cow_info.ForeLegPosFront }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Глубина вымени: </div>
                        <div>{{ cow_info.UdderDepth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Баланс вымени: </div>
                        <div>{{ cow_info.UdderBalance }}</div>
                    </div>
                </div>
                <div class="column">
                    <div class="mark-with-num">
                        <div class="ext-param">Высота прикрепления задних долей вымени: </div>
                        <div>{{ cow_info.HeightOfUdderAttach }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Прикрепление передних долей вымени: </div>
                        <div>{{ cow_info.ForeUdderAttach }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Ширина задних долей вымени: </div>
                        <div>{{ cow_info.HindUdderWidth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Длина передних долей вымени: </div>
                        <div>{{ cow_info.ForeUdderWidth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Расположение передних сосков (вид сзади): </div>
                        <div>{{ cow_info.ForeUdderPlcRear }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Расположение задних сосков: </div>
                        <div>{{ cow_info.HindTeatPlc }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Длина передних сосков: </div>
                        <div>{{ cow_info.ForeTeatLendth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Длина задних сосков: </div>
                        <div>{{ cow_info.RearTeatLength }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Диаметр передних сосков: </div>
                        <div>{{ cow_info.ForeTeatDiameter }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Диаметр задних сосков: </div>
                        <div>{{ cow_info.RearTeatDiameter }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Глубина центральной связки: </div>
                        <div>{{ cow_info.CenterLigamentDepth }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Гармоничность движения: </div>
                        <div>{{ cow_info.HarmonyOfMovement }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Упитанность: </div>
                        <div>{{ cow_info.Conditioning }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Выраженность молочных вен: </div>
                        <div> {{ cow_info.ProminenceOfMilkVeins }}</div>
                    </div> 
                </div> 
            </div>
            <div class="ext-100mark-title" @click="show100Mark">> Признаки со 100-бальной оценкой</div>
            <div class="ext-100mark" :class="{'ext-hide': !isVisible100Mark}">
                <div class="column">
                    <div class="mark-with-num">
                        <div class="ext-param">Молочная сила: </div>
                        <div>{{ cow_info.MilkStrength }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Телосложение: </div>
                        <div>{{ cow_info.BodyStructure }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Конечности: </div>
                        <div>{{ cow_info.Limbs }}</div>
                    </div>
                    <div class="mark-with-num">
                        <div class="ext-param">Вымя: </div>
                        <div>{{ cow_info.Udder }}</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    data() {
        return {
            isVisible9Mark: false,
            isVisible100Mark: false,
            cow_info: {},
            img: '',
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/exterior`);
        let result = await response.json();
        this.cow_info = result;
        console.log(this.cow_info);
        if(this.cow_info.PicturePath) {
            this.img = this.cow_info.PicturePath;
            //this.setImg(this.img);
        }
    },
    methods: {
        show9Mark () {
            if (this.isVisible9Mark) {
                this.isVisible9Mark = false;
            } else {
                this.isVisible9Mark = true;
            }
        },
        show100Mark() {
            if (this.isVisible100Mark) {
                this.isVisible100Mark = false;
            } else {
                this.isVisible100Mark = true;
            }
        },
        setImg(path) {
            console.log(path);
            let block = document.getElementsByClassName('krs-photo');
            block.style.backgroundImage = `url('/api/static/exterior/${path}')`;
        }
    },
    computed: {
        logo() {
            return '/api/static/exterior/' + this.img;
        }
    }
}
</script>
    
<style scoped>
.extrat-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding: 0 0 40px 0;
    width: max-content;
}

.exterior-title {
    display: flex;
    align-items: start;
    justify-content: space-between;
    padding-right: 50px;
}

.krs-photo {
    width: 150px;
    height: 150px;
    background-color: none;
}

.ext-9mark-title, .ext-100mark-title {
    font-size: 120%;
    margin: 18px 0;
    cursor: pointer;
    transition: 0.3s;
    width: max-content;
}

.ext-9mark-title:hover, .ext-100mark-title:hover {
    color: grey;
}

.ext-9mark, .ext-100mark {
    display: flex;
    /* flex-direction: column;
    flex-wrap: wrap; */
    justify-content: flex-start;
    height: 450px;
    font-size: 90%;
    margin-bottom: 15px;
}

.ext-100mark {
    height: max-content;
}

.ext-hide {
    display: none;
}

.ext-param {
    margin: 5px 30px 5px 0;
    color: rgb(74, 74, 74);
    width: 250px;
}

.mark-with-num {
    display: flex;
    justify-content: space-between;
    align-items: end;
    margin-right: 30px;
    border-bottom: 1px solid rgb(242, 237, 248);
}

.column {
    width: auto;
}
</style>