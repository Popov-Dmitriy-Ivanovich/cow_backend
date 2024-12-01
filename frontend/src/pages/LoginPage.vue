<template>
    <div class="lg">
    <form class="login">
        <div class="login-title">
            <div>Вход</div>
            <div class="registration-link" @click="$router.push('/registration')">Регистрация</div>
        </div>
        <label for="login-email" class="login-label">Почта</label><br>
        <input type="email" class="login-input" id="login-email"><br>
        <label for="login-password" class="login-label">Пароль</label><br>
        <input type="password" class="login-input" id="login-password"><br>
        <div class="forgot-password">Забыли пароль?</div>
        <div class="login-btns">
            <button type="submit" class="btn login-btn">Вход</button>
        </div>
    </form></div>
</template>

<script>

</script>

<style scoped>
    .lg {
        display: flex;
        flex-direction: column;
        margin-top: 130px;
        align-items: center;
    }
    .login {
        font-family: Open Sans, sans-serif;
        background-color: white;
        padding: 50px 80px 70px 80px;
        border-radius: 20px;
    }

    .login-title {
        display: flex;
        font-size: 170%;
        align-items: flex-end;
        padding-bottom: 40px;
    }

    .registration-link {
        font-size: 60%;
        padding-left: 20px;
        color: grey;
        cursor: pointer;
        transition: 0.3s;
    }

    .registration-link:hover {
        color: black;
    }

    .login-input {
        height: 35px;
        width: 300px;
        margin: 15px 0 0 0;
        padding: 0 10px;
        font-size: 120%;
        box-sizing: border-box;
        outline: none;
        border: 3px solid rgb(195, 200, 212);
        border-radius: 10px;
        transition: 0.3s;
    }

    .login-input[id=login-email] {
        margin-bottom: 30px;
    }

    .forgot-password {
        font-size: 80%;
        color: grey;
        padding: 10px 0;
        cursor: pointer;
        transition: 0.3s;
    }

    .forgot-password:hover {
        color: black;
    }

    .login-input:hover {
        border: 3px solid rgb(101, 102, 170);
    }

    .btn {
        font-size: 100%;
        border: none;
        margin-right: 20px;
    }

    .login-btn {
        margin-bottom: 20px;
        padding: 7px 0;
        width: 100%;
        color: white;
        background-color: rgb(101, 82, 183);
        border-radius: 10px;
        transition: 0.3s;
        cursor: pointer;
    }

    .login-btn:hover {
        background-color: rgb(69, 27, 152);
    }

    .login-btns {
        position: relative;
        top: 30px;
    }
</style>