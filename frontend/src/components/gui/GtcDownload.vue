<template>
<div>
    <button class="back-button" onclick="window.history.back();">Назад</button>
    <div class="form-example">
        <h1>Загрузка gtc</h1>
        <form action="https://genmilk.ru/api/load/gtcFile" method="post" enctype="multipart/form-data">
            <label for="csv">Csv файл сопоставления Селекса gtc-файлу:</label>
            <input type="file" id="csv" name="csv" required="">

            <label for="gtc">Список gtc-файлов:</label>
            <input type="file" id="gtc" name="gtc" required="" multiple="">
            
            <input type="submit" value="Загрузить">
        </form>
    </div>
</div>
</template>

<script>

</script>

<style scoped>
h1 {
    text-align: center;
    color: #333;
    margin-bottom: 20px;
}

.form-example {
    background-color: #fff;
    border-radius: 8px;
    padding: 20px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    max-width: 400px;
    width: 100%;
}

label {
    display: block;
    margin-bottom: 10px;
    font-weight: bold;
    color: #555;
}

input[type="file"] {
    display: block;
    margin-bottom: 20px;
}
input[type="text"] {
    display: block;
    margin-bottom: 20px;
}
input[type="submit"] {
    background-color: #6f42c1;  
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease;
    font-size: 16px;
    width: 100%;
}

input[type="submit"]:hover {
    background-color: #5a32a3;  
}

.back-button {
    position: relative;
    margin: 10px;
    background-color: #e0c6f2;  
    color: #333;
    border: none;
    padding: 8px 12px;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
}

.back-button:hover {
    background-color: #d1b7e3;  
}
</style>