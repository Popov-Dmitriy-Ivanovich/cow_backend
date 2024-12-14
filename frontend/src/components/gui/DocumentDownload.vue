<template>
<div>
    <button class="back-button" onclick="window.history.back();">Назад</button>
    <div class="form-example" >
        <h1>Загрузка коровы</h1>
        <form id="form-document">
            <label for="Document">Выберите файл:</label>
            <input type="file" id="Document" name="Document" required="">

            <label for="CowID">ID коровы:</label>
            <input type="text" id="CowID" name="CowID" required="">

            <input type="submit" value="Загрузить">
        </form>
    </div>
</div>
</template>

<script>
export default {
    mounted() {
        let formDocument = document.getElementById('form-document');
        formDocument.addEventListener('submit', this.sendData)
    },
    methods: {
        async sendData(event) {
            event.preventDefault();
            let formData = new FormData();
            let file = document.getElementById('Document').value;
            formData.append('Document', file);

            // let response = await fetch('/api/load/document', {
            //     method: 'POST',
            //     headers: {
            //         'Authorization': localStorage.getItem('jwt')
            //     },
            //     body: formData
            // })
            // let result = await response.json();
            console.log(formData);
        }
    }
}
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