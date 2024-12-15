<template>
    <div class="cm-main">
        <button class="back-button" @click="$router.push(`/gui`)">Назад</button>
        <div class="form-example">
            <h1>Загрузка контрольных доений</h1>
            <form action="https://genmilk.ru/api/load/checkMilk" method="post" enctype="multipart/form-data">
                <label for="csv">Выберите файл:</label>
                <input type="file" id="csv" name="csv" required="">
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
    background-color: #e0c6f2;  
    color: #333;
    border: none;
    padding: 8px 12px;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
    margin: 10px
}

.back-button:hover {
    background-color: #d1b7e3;  
}
</style>