<template>
    <div class="back">
        <div class="title-text">
            <div class="title">База данных молочного поголовья КРС Воронежской области</div>
            <div class="title-description">
                Программа по геномной селекции крупного рогатого скота молочного направления, 
                реализующая оценку племенной ценности и позволяющая принимать селекционное 
                решение при осеменении и выбраковке животных
            </div>
            <div>
                <button class="but-login" @click="$router.push('/login')">Войти</button>
                <button class="but-reg" @click="$router.push('/registration')">Регистрация</button>
            </div>
        </div>
        
    </div>
</template>

<script>

</script>

<style scoped>
    .back {
        /* width: 100vw; */
        height: 100vh;
        background: radial-gradient(circle at bottom right, #8daec4, transparent 60%), 
        radial-gradient(circle at top left, #423b5d, #ea070700 60%), 
        radial-gradient(circle at center, #581f70, transparent 70%), 
        radial-gradient(circle at center bottom, #7578b1, transparent 90%), 
        linear-gradient(135deg, #4e1698, #5a8bc1);
        text-align: center;
    }

    .title-text {
        color: white;
        font-family: Open Sans, sans-serif;
        position: absolute;
        top: 130px;
        left: 50%;
        transform: translate(-50%, 0);
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 60vw;
    }

    .title {
        font-size: 64px;
        padding-bottom: 40px;
        font-family: Open Sans, sans-serif;
    }

    .title-description {
        font-size: 18px;
        padding-bottom: 30px;
    }

    .but-login, .but-reg {
        padding: 14px 40px;
        border-radius: 4px;
        border: none;
        cursor: pointer;
        font-family: Open Sans, sans-serif;
        font-size: 1.5rem;
        transition: 0.3s;
        margin: 0 20px;
    }

    .but-reg {
        background-color: white;
        color: rgb(83, 47, 175);
    }

    .but-reg:hover {
        color: white;
        background-color: rgb(146, 122, 210);
    }

    .but-login {
        color: white;
        background-color: rgb(69, 4, 123);
        border: 2px solid rgb(69, 4, 123);
    }
    
    .but-login:hover {
        background-color: rgb(96, 11, 166);
    }
</style>