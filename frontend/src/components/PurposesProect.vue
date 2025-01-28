<template>
    <div class="purposes-main">
        <div class="purpos-title">Цели проекта</div>
        <div class="purpos-cards">
            <div class="purpos-card">
                <div class="card-title">Выявлять</div>
                <div>поголовье с важными генетическими признаками и закреплять их</div>
            </div>
            <div class="purpos-card">
                <div class="card-title">Ускорять</div>
                <div>рост племенной ценности быков-производителей за счёт сокращения интервала смены их поколений</div>
            </div>
            <div class="purpos-card">
                <div class="card-title">Оценивать</div>
                <div>заранее окупаемость животных и увеличить прибыль от животноводства</div>
            </div>
        </div>
    </div>
</template>

<script>

</script>

<style scoped>
.purposes-main {
    background-color: white;
    width: 100%;
    font-family: Open Sans, sans-serif;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding-bottom: 30px;
}

.purpos-title {
    color: rgb(10, 113, 75);
    font-size: 190%;
    padding: 40px 0 20px 0;
}

.purpos-cards {
    display: flex;
    width: 70%;
    justify-content: space-between;
}

.purpos-card {
    width: 300px;
    height: 150px;
    overflow: auto;
    box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
    padding: 15px 20px;
    border-radius: 10px;
    display: flex;
    align-items: center;
    flex-direction: column;
    text-align: center;
}

.card-title {
    font-size: 190%;
    color: rgb(10, 113, 75);
    margin-bottom: 20px;
}
</style>