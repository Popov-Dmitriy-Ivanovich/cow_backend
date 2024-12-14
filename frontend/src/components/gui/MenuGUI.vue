<template>
    <div class="main-main">
        <h1>Выбор действия</h1>
    
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadCow`)">Загрузка животных</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadLactations`)">Загрузка лактаций</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadControlMilking`)">Загрузка контрольных доений</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadEvents`)">Загрузка вет. событий</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadGenotype`)">Загрузка генотипирования</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadRating`)">Загрузка оценок</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadExterior`)">Загрузка экстерьера</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadGtc`)">Загрузка gtc-файла</button> 
        </div>
    </div>
</template>

<script>

</script>

<style scoped>
h1 {
    text-align: center;
    color: #6f42c1;  
    margin-bottom: 20px;
}

.UserChoise {
    background-color: #ffffff;  
    border-radius: 8px;
    padding: 20px 20px 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    text-align: center;  
    min-width: 20%;
    margin: 10px;
}

.UserChoiseButton {
    background-color: #6f42c1;  
    color: #ffffff;  
    border: none;
    border-radius: 5px;
    padding: 10px 20px 5px;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.3s;  
    min-width: 80%;
    min-height: 80%;
}

.UserChoiseButton:hover {
    background-color: #5a32a1;  
}

.message {
    margin-top: 20px;
    color: #666;  
    font-size: 14px;
}

.main-main {
    width: 70%;
    font-family: Open Sans, sans-serif;
}
</style>