<template>
    <div class="help-page">
        <div class="help-title">
            <div class="help-title-text">Помощь</div>
            <div class="searchhelp-line">
                <input type="text" placeholder="Поиск информации..." class="search-help" autocomplete="off">
                <button class="searchhelp-btn">Найти</button>
            </div>
        </div>
        <div class="help-section">
            <div class="helpsection-title">Животные, поиск и выбор</div>
            <div class="help-flexblocks">
                <HelpBlock/>
                <HelpBlock/>
            </div>
        </div>   
        <div class="help-section">
            <div class="helpsection-title">Аналитика, диаграммы и отчеты</div>
        </div> 
        <div class="help-section">
            <div class="helpsection-title">Загрузка и обновление данных</div>
        </div> 
        <div class="help-section">
            <div class="helpsection-title">Участие в проекте, доступ к данным</div>
        </div> 
    </div>
</template>

<script>
import HelpBlock from '@/components/HelpBlock.vue';

export default {
    components: {
        HelpBlock,
    }
}
</script>

<style scoped>
.help-page {
    margin-top: 110px;
    font-family: Open Sans, sans-serif;
    padding-left: 11%;
}

.help-title-text {
    font-size: 200%;
    padding: 10px 0 30px 0;
}

.searchhelp-line {
    display: flex;
    align-items: center;
}

.search-help {
    width: 75%;
    height: 30px;
    padding: 0 15px;
    font-size: 110%;
    outline: none;
    border-radius: 10px;
    border: 2px solid rgb(210, 208, 218);
    background-color: rgb(255, 255, 255);
    transition: 0.3s;
}

.search-help:focus {
    border: 2px solid rgb(122, 123, 193);
}
    
.search-help::placeholder {
    opacity: 0.5;
    color:rgb(79, 60, 126);
}

.searchhelp-btn {
    background-color: white;
    border: 1px solid rgb(101, 102, 170);
    color: rgb(101, 102, 170);
    padding: 0 7px;
    height: 30px;
    border-radius: 10px;
    width: 70px;
    cursor: pointer;
    margin: 0 15px;
    transition: 0.3s;
}

.searchhelp-btn:hover {
    background-color: rgb(101, 102, 170);
    color: white;
}

.help-section {
    margin: 20px 0;
    padding: 10px 0;
}

.helpsection-title {
    font-size: 140%;
}

.help-flexblocks {
    display: flex;
}
</style>