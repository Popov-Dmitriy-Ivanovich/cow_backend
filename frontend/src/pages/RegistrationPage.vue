<template>
    <div class="rg">
    <form class="registrate">
        <div class="registrate-title">
            <div>Регистрация</div>
            <div class="login-link" @click="$router.push('/login')">Вход</div>
        </div>
        <label>Регион (область)</label><br>
        <MultiselectRegions class="region"/>
        <br>
        <label>Роль</label>
        <div class="role">
            <input type="radio" name="role">Федеральный чиновник<br>
            <input type="radio" name="role">Региональный чиновник<br>
            <input type="radio" name="role">Фермер<br>
        </div>
        <label>ФИО</label><br>
        <input type="text" placeholder="Иванов Иван Иванович" class="registration-field" @keyup="valid" id="fio" autocomplete="off"> <br>
        <label>Название организации</label><br>
        <input type="text"  class="registration-field"> <br>
        <label>Электронная почта</label><br>
        <input type="email" placeholder="example@email.com" class="registration-field"> <br>
        <label>Телефон</label><br>
        <input type="tel" 
        pattern="^(\+7|8)?[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$" 
        placeholder="8 999 999 99 99"
        class="registration-field"> <br>
        <label>Пароль</label><br>
        <input type="password"  class="registration-field"> <br>
        <label>Повторите пароль</label><br>
        <input type="password"  class="registration-field"> <br>
        <label>Ваши ожидания от проекта генетической селекции</label><br>
        <textarea  class="registration-field textarea"></textarea> <br>
        <button type="submit" class="registr-btn">Отправить</button>
    </form></div>
</template>

<script>
import MultiselectRegions from '@/components/MultiselectRegions.vue';

export default {
    components: {
        MultiselectRegions
    },
    methods: {
        valid() {
            document.getElementById('fio').value = document.getElementById('fio').value.replace(/[\d]/g,'');
        }
    }
}
</script>

<style scoped>
    .rg {
        display: flex;
        flex-direction: column;
        margin-top: 130px;
        align-items: center;
    }
    .registrate {
        font-family: Open Sans, sans-serif;
        background-color: white;
        padding: 50px 80px 70px 80px;
        border-radius: 20px;
    }

    .registrate-title {
        display: flex;
        font-size: 170%;
        align-items: flex-end;
        padding-bottom: 40px;
    }

    .login-link {
        font-size: 60%;
        padding-left: 20px;
        color: grey;
        cursor: pointer;
        transition: 0.3s;
    }

    .login-link:hover {
        color: black;
    }

    .registration-field {
        margin: 15px 20px;
        height: 35px;
        width: 300px;
        padding: 0 10px;
        font-size: 14px;
        box-sizing: border-box;
        outline: none;
        border: 3px solid rgb(195, 200, 212);
        border-radius: 10px;
        transition: 0.3s;
    }

    .registration-field:hover {
        border: 3px solid rgb(101, 102, 170);
    }

    .textarea {
        height: 70px;
        padding-top: 10px;
    }

    .role {
        margin: 15px 20px;
    }

    .registr-btn {
        font-size: 100%;
        border: none;
        margin: 15px 20px;
        padding: 7px 0;
        width: 300px;
        color: white;
        background-color: rgb(101, 82, 183);
        border-radius: 10px;
        transition: 0.3s;
        cursor: pointer;
    }

    .registr-btn:hover {
        background-color: rgb(69, 27, 152);
    }

    .region {
        max-width: 300px;
        margin: 15px 20px;
    }
</style>