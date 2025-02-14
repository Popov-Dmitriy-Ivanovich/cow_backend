<template>
    <!-- <div class="all-ranges">
        <div class="column">
            <div class="concret-range">
                <p class="range-label">Максимальный удой, кг</p>
                <div class="range-line">
                    <span>{{ minMilking }}</span>
                    <input type="range" :min="minMilking" :max="maxMilking" v-model="milking" class="range">
                    <span>{{ maxMilking }}</span>
                </div>
                <p class="current-value">{{ milking }}</p>
            </div>
            <div class="concret-range">
                <p class="range-label">Максимальные показатели жира, %</p>
                <div class="range-line">
                    <span>{{ minFat }}</span>
                    <input type="range" :min="minFat" :max="maxFat" v-model="fat" class="range">
                    <span>{{ maxFat }}</span>
                </div>
                <p class="current-value">{{ fat }}</p>
            </div>
            <div class="concret-range">
                <p class="range-label">Максимальные показатели белка, %</p>
                <div class="range-line">
                    <span>{{ minProtein }}</span>
                    <input type="range" :min="minProtein" :max="maxProtein" v-model="protein" class="range">
                    <span>{{ maxProtein }}</span>
                </div>
                <p class="current-value">{{ protein }}</p>
            </div>
            <div class="concret-range">
                <p class="range-label">Максимальное продуктивное долголетие</p>
                <div class="range-line">
                    <span>{{ minLact }}</span>
                    <input type="range" :min="minLact" :max="maxLact" v-model="lact" class="range">
                    <span>{{ maxLact }}</span>
                </div>
                <p class="current-value">{{ lact }}</p>
            </div>
        </div>
        <div class="column">
            <div class="concret-range">
                <p class="range-label">Максимальные показатели фертильности</p>
                <div class="range-line">
                    <span>{{ minOsem }}</span>
                    <input type="range" :min="minOsem" :max="maxOsem" v-model="osem" class="range">
                    <span>{{ maxOsem }}</span>
                </div>
                <p class="current-value">{{ osem }}</p>
            </div>
            <div class="concret-range">
                <p class="range-label">Максимальные значения экстерьера</p>
                <div class="range-line">
                    <span>{{ minExterior }}</span>
                    <input type="range" :min="minExterior" :max="maxExterior" v-model="exterior" class="range">
                    <span>{{ maxExterior }}</span>
                </div>
                <p class="current-value">{{ exterior }}</p>
            </div>
            <div class="concret-range">
                <p class="range-label">Коэффициент инбридинга</p>
                <div class="range-line">
                    <span>{{ minInbriding }}</span>
                    <input type="range" :min="minInbriding" :max="maxInbriding" v-model="inbriding" class="range">
                    <span>{{ maxInbriding }}</span>
                </div>
                <p class="current-value">{{ inbriding }}</p>
            </div>
        </div>
    </div> -->
    <div class="range-table">
        <table>
            <tbody>
                <tr>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Максимальный удой, кг</p>
                            <div class="range-line">
                                <span>{{ minMilking }}</span>
                                <input type="range" :min="minMilking" :max="maxMilking" v-model="milking" class="range">
                                <span>{{ maxMilking }}</span>
                            </div>
                            <p class="current-value">{{ milking }}</p>
                        </div>
                    </td>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Максимальные показатели фертильности</p>
                            <div class="range-line">
                                <span>{{ minOsem }}</span>
                                <input type="range" :min="minOsem" :max="maxOsem" v-model="osem" class="range">
                                <span>{{ maxOsem }}</span>
                            </div>
                            <p class="current-value">{{ osem }}</p>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Максимальные показатели жира, %</p>
                            <div class="range-line">
                                <span>{{ minFat }}</span>
                                <input type="range" :min="minFat" :max="maxFat" v-model="fat" class="range">
                                <span>{{ maxFat }}</span>
                            </div>
                            <p class="current-value">{{ fat }}</p>
                        </div>
                    </td>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Максимальные значения экстерьера</p>
                            <div class="range-line">
                                <span>{{ minExterior }}</span>
                                <input type="range" :min="minExterior" :max="maxExterior" v-model="exterior" class="range">
                                <span>{{ maxExterior }}</span>
                            </div>
                            <p class="current-value">{{ exterior }}</p>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Максимальные показатели белка, %</p>
                            <div class="range-line">
                                <span>{{ minProtein }}</span>
                                <input type="range" :min="minProtein" :max="maxProtein" v-model="protein" class="range">
                                <span>{{ maxProtein }}</span>
                            </div>
                            <p class="current-value">{{ protein }}</p>
                        </div>
                    </td>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Коэффициент инбридинга</p>
                            <div class="range-line">
                                <span>{{ minInbriding }}</span>
                                <input type="range" :min="minInbriding" :max="maxInbriding" v-model="inbriding" class="range">
                                <span>{{ maxInbriding }}</span>
                            </div>
                            <p class="current-value">{{ inbriding }}</p>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td>
                        <div class="concret-range">
                            <p class="range-label">Максимальное продуктивное долголетие</p>
                            <div class="range-line">
                                <span>{{ minLact }}</span>
                                <input type="range" :min="minLact" :max="maxLact" v-model="lact" class="range">
                                <span>{{ maxLact }}</span>
                            </div>
                            <p class="current-value">{{ lact }}</p>
                        </div>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <div>
        <button class="btn-select">Подобрать</button>
        <!-- <button>Отобразить данные по быкам</button>
        <button>Отобразить данные по потомству</button> -->
    </div>
    <div class="match-table">
        <table>
            <thead>
                <tr>
                    <th>РСХН</th>
                    <th>Селекс</th>
                    <th>Инвентарный номер</th>
                    <th>Кличка</th>
                    <th>Дата рождения</th>
                    <th>Удой</th>
                    <th>Жир</th>
                    <th>Белок</th>
                    <th>Долголетие</th>
                    <th>Фертильность</th>
                    <th>Экстерьер</th>
                </tr>
            </thead>
            <tbody>
                <tr>

                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    data() {
        return {
            milking: 10,
            maxMilking: 50,
            minMilking: 5,

            fat: 15,
            maxFat: 50,
            minFat: 5,

            protein: 20,
            maxProtein: 50,
            minProtein: 5,

            lact: 25,
            maxLact: 50,
            minLact: 5,

            osem: 30,
            maxOsem: 50,
            minOsem: 5,

            exterior: 35,
            maxExterior: 50,
            minExterior: 5,

            inbriding: 40,
            maxInbriding: 50,
            minInbriding: 5,
        }
    }
}
</script>

<style scoped>
.all-ranges {
    display: flex;
    margin-bottom: 20px;
}

.range {
    width: 170px;
}

.column {
    margin-right: 50px;
    display: flex;
    flex-direction: column;
}

.concret-range {
    display: flex;
    flex-direction: column;
    width: 280px;
    height: 70px;
    align-items: center;
    justify-content: space-between;
    margin: 10px 40px 10px 0;
}

.range-line {
    display: flex;
    justify-content: space-around;
    width: 100%;
}

.current-value {
    align-self: center;
}

.range-label {
    text-align: center;
    margin: 10px 0;
}

.match-table {
    margin-top: 40px;
    overflow: auto;
}

table {
    border-collapse: collapse;
}

th {
    font-weight: normal;
    padding-left: 8px;
}

.btn-select {
    background-color: rgb(101, 102, 170);
    border: 1px solid rgb(101, 102, 170);
    color: white;
    padding: 5px 9px;
    height: max-content;
    border-radius: 10px;
    width: max-content;
    cursor: pointer;
    margin: 0 15px;
    font-size: 115%;
    transition: 0.3s;
    margin-top: 40px;
}

.btn-select:hover {
    background-color: white;
    color: rgb(101, 102, 170);
}
</style>