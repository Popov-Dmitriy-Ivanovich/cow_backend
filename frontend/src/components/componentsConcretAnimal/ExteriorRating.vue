<template>
    <div>
        <div class="extrat-title">Оценка экстерьера</div>
        <div class="exterior-title">
            <div>Общая оценка экстерьера: {{ cow_info.RATING }}</div>
            <div class="krs-photo">Фото КРС</div>
        </div>
        <div>
            <div class="ext-9mark-title" @click="show9Mark">> Признаки с 9-бальной оценкой</div>
            <div class="ext-9mark" :class="{'ext-hide': !isVisible9Mark}">
                <div class="mark-with-num">
                    <div class="ext-param">Ширина груди: </div>
                    <div>{{ cow_info.CHEST_WIDTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Ширина таза: </div>
                    <div>{{ cow_info.PELVIC_WIDTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Рост в крестце: </div>
                    <div>{{ cow_info.HEIGHT_AT_SACRUM }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Глубина туловища: </div>
                    <div>{{ cow_info.BODY_DEPTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Тип экстерьера: </div>
                    <div>{{ cow_info.EXTERIOR_TYPE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Качество костяка/скакательный сустав (вид сзади): </div>
                    <div>{{ cow_info.BONE_Q_HOCK_JOINT_REAR }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Угол наклона крестца: </div>
                    <div>{{ cow_info.SACRUM_ANGLE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Длина крестца: </div>
                    <div>{{ cow_info.ACRUM_LENGTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Линия верха: </div>
                    <div>{{ cow_info.TOPLINE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Угол копыта: </div>
                    <div>{{ cow_info.HOOF_ANGLE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Постановка задних ног (вид сбоку): </div>
                    <div>{{ cow_info.HIND_LEG_POS_SIDE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Постановка задних ног (вид сзади): </div>
                    <div>{{ cow_info.HIND_LEG_POS_REAR }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Постановка передних конечностей (вид спереди): </div>
                    <div>{{ cow_info.FORE_LEG_POS_FRONT }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Глубина вымени: </div>
                    <div>{{ cow_info.UDDER_DEPTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Баланс вымени: </div>
                    <div>{{ cow_info.UDDER_BALANCE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Высота прикрепления задних долей вымени: </div>
                    <div>{{ cow_info.HEIGHT_OF_UDDER_ATTACH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Прикрепление передних долей вымени: </div>
                    <div>{{ cow_info.FORE_UDDER_ATTACH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Ширина задних долей вымени: </div>
                    <div>{{ cow_info.HIND_UDDER_WIDTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Длина передних долей вымени: </div>
                    <div>{{ cow_info.FORE_UDDER_LENGTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Расположение передних сосков (вид сзади): </div>
                    <div>{{ cow_info.FORE_UDDER_PLC_REAR }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Расположение задних сосков: </div>
                    <div>{{ cow_info.HIND_TEAT_PLC }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Длина передних сосков: </div>
                    <div>{{ cow_info.FORE_TEAT_LENGTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Длина задних сосков: </div>
                    <div>{{ cow_info.REAR_TEAT_LENGTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Диаметр передних сосков: </div>
                    <div>{{ cow_info.FORE_TEAT_DIAM }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Диаметр задних сосков: </div>
                    <div>{{ cow_info.REAR_TEAT_DIAM }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Глубина центральной связки: </div>
                    <div>{{ cow_info.CENTR_LIGAMENT_DEPTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Гармоничность движения: </div>
                    <div>{{ cow_info.HARMONY_OF_MOVEMENT }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Упитанность: </div>
                    <div>{{ cow_info.CONDITIONING }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Выраженность молочных вен: </div>
                    <div> {{ cow_info.PROMINENCE_OF_MILK_VEINS }}</div>
                </div>  
            </div>
            <div class="ext-100mark-title" @click="show100Mark">> Признаки со 100-бальной оценкой</div>
            <div class="ext-100mark" :class="{'ext-hide': !isVisible100Mark}">
                <div class="mark-with-num">
                    <div class="ext-param">Молочная сила: </div>
                    <div>{{ cow_info.MILK_STRENGTH }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Телосложение: </div>
                    <div>{{ cow_info.BODY_STRUCTURE }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Конечности: </div>
                    <div>{{ cow_info.LIMBS }}</div>
                </div>
                <div class="mark-with-num">
                    <div class="ext-param">Вымя: </div>
                    <div>{{ cow_info.UDDER }}</div>
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
        }
    },
    async created() {
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`https://genmilk.ru:9050/api/cow_exterior?ID_COW=${cow_id}`);
        let result = await response.json();
        this.cow_info = result;
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
    background-color: rgb(225, 225, 238);
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
    flex-direction: column;
    flex-wrap: wrap;
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
    margin: 5px 0;
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
</style>