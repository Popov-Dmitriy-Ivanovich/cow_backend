<template>
    <div class="back">
        <div class="title-text">
            <div class="title">База данных молочного поголовья КРС Краснодарского края</div>
            <div class="title-description">
                Программа по геномной селекции крупного рогатого скота молочного направления, 
                реализующая оценку племенной ценности и позволяющая принимать селекционное 
                решение при осеменении и выбраковке животных
            </div>
            <div>
                <button class="but-login" @click="$router.push('/login')" v-if="!token">Войти</button>
                <button class="but-reg" @click="$router.push('/registration')" v-if="!token">Регистрация</button>
                <button class="but-login" @click="$router.push('/animals')" v-if="token">Животные</button>
                <button class="but-reg" @click="$router.push('/analytics')" v-if="token">Аналитика</button>
            </div>
        </div>
        
    </div>
</template>

<script>
export default {
    computed: {
        token(){
            console.log(localStorage.getItem('jwt'));
            return this.$store.state.isLogged || localStorage.getItem('jwt');
        }
    }
}
</script>

<style scoped>
    .back {
        /* width: 100vw; */
        height: 100vh;
        background: radial-gradient(circle at bottom right, #8dc49f, transparent 60%), radial-gradient(circle at top left, #3b4b5d, #ea070700 60%), radial-gradient(circle at center, #267166, transparent 70%), radial-gradient(circle at center bottom, #75b190, transparent 90%), linear-gradient(135deg, #109874, #52af8a);
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
        color: rgb(29, 172, 108);
    }

    .but-reg:hover {
        color: white;
        background-color: rgb(43, 201, 130);
    }

    .but-login {
        color: white;
        background-color: rgb(9, 111, 67);
        /* border: 2px solid rgb(3, 91, 40); */
    }
    
    .but-login:hover {
        background-color: rgb(11, 157, 82);
    }
</style>