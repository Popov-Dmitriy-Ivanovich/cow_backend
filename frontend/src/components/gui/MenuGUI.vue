<template>
    <h1>Загрузка файлов</h1>
    <div class="main-main">
        <!-- <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadCow`)">Загрузка животных</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadLactations`)">Загрузка лактаций</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadControlMilking`)">Загрузка контрольных доений</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadEvents`)">Загрузка ветеринарных событий</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadGenotype`)">Загрузка генотипирования</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadRating`)">Загрузка оценок</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/exteriorDownload`)">Загрузка файлов экстерьера</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadExterior`)">Загрузка фотографий экстерьера</button> 
        </div>
        <div class="UserChoise">
            <button class="UserChoiseButton" @click="$router.push(`/gui/downloadGtc`)">Загрузка gtc-файла</button> 
        </div> -->

        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Животные</h1>
                <form action="https://genmilk.ru/api/load/cow" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Выберите файл:</label><br> -->
                    <input type="file" id="csv" name="csv" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Лактации</h1>
                <form action="https://genmilk.ru/api/load/lactation" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Выберите файл:</label><br> -->
                    <input type="file" id="csv" name="csv" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Контрольные доения</h1>
                <form action="https://genmilk.ru/api/load/checkMilk" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Выберите файл:</label><br> -->
                    <input type="file" id="csv" name="csv" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Ветеринарные события</h1>
                <form action="https://genmilk.ru/api/load/event" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Выберите файл:</label><br> -->
                    <input type="file" id="csv" name="csv" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Информация о генотипировании</h1>
                <form action="https://genmilk.ru/api/load/genetic" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Выберите файл:</label><br> -->
                    <input type="file" id="csv" name="csv" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Оценки</h1>
                <form action="https://genmilk.ru/api/load/grade" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Выберите файл:</label><br> -->
                    <input type="file" id="csv" name="csv" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Файлы экстерьера</h1>
                <form method="post" enctype="multipart/form-data">            
                    <!-- <label for="CsvFile">Csv-файл:</label><br> -->
                    <input type="file" id="CsvFile" name="CsvFile" required=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
        <!-- <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Загрузка фотографий экстерьера</h1>
                <form action="https://genmilk.ru/api/load/exterior" method="post" enctype="multipart/form-data">            
                    <label for="CsvFile">Csv-файл сопоставления Селекса файлу фотографии:</label><br>
                    <input type="file" id="CsvFile" name="CsvFile" required=""><br>
                    <label for="Pictures">Список файлов фотографий животных:</label><br>
                    <input type="file" id="Pictures" name="Pictures" required="" multiple=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
            </div>
        </div> -->
        <div class="parent-for-form">
            <div class="form-example">
                <h1 class="card-title">Gtc-файлы</h1>
                <form action="https://genmilk.ru/api/load/gtcFile" method="post" enctype="multipart/form-data">
                    <!-- <label for="csv">Csv файл сопоставления Селекса gtc-файлу:</label><br>
                    <input type="file" id="csv" name="csv" required=""><br>
                    <label for="gtc">Список gtc-файлов:</label><br> -->
                    <input type="file" id="gtc" name="gtc" required="" multiple=""><br>
                    <input type="submit" value="Загрузить" class="download-btn">
                </form>
                <div class="date-of-download">Загрузка завершена: </div>
            </div>
        </div>
    </div>
</template>

<script>

</script>

<style scoped>
h1 {
    text-align: center;
    font-size: 185%;
    font-family: Open Sans, sans-serif;
    color:rgb(37, 0, 132);
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
    font-family: Open Sans, sans-serif;
    display: grid;
    grid-template-columns: auto auto auto;
}

.parent-for-form {
    padding: 20px;
}

.form-example {
    background-color: #fff;
    border-radius: 8px;
    padding: 20px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    width: 320px;
}

.card-title {
    color: black;
    font-size: 120%;
}

.download-btn {
    background-color: #6f42c1;  
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
    font-size: 16px;
    width: 100%;
    margin-top: 20px;
}

.download-btn:hover {
    background-color: #5a32a3;  
}

.date-of-download {
    background-color: #e7e4ee;
    margin-top: 20px;
    padding: 8px 10px;
    border-radius: 5px;
}
</style>