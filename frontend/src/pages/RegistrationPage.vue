<template>
    <div class="rg">
    <form class="registrate" id="form-registrate">
        <div class="registrate-title">
            <div>Регистрация</div>
            <div class="login-link" @click="$router.push('/login')">Вход</div>
        </div>
        <label>Регион (область) <span class="required">*</span></label><br>
        <MultiselectReg class="region" @sendToMain="setIdReg"/>
        <br>
        <label>Хозяйство</label><br>
        <MultiselectHozRegister  class="hoz" @sendToMain="setIdHoz" v-bind:regionId="newUser.regionId"/> <br>
        <!-- <div class="hoz-underline">Если нет нужного хозяйства, вы можете <span class="create-hoz">создать</span> его</div> -->
        <label>Роль <span class="required">*</span></label>
        <div class="role">
            <input type="radio" name="role" :value="3" v-model="newUser.roleId">Федеральный чиновник<br>
            <input type="radio" name="role" :value="2" v-model="newUser.roleId">Региональный чиновник<br>
            <input type="radio" name="role" :value="1" v-model="newUser.roleId">Фермер<br>
        </div>
        <label>ФИО <span class="required">*</span></label><br>
        <input type="text" placeholder="Иванов Иван Иванович" class="registration-field" @keyup="valid" id="fio" autocomplete="off" v-model="newUser.nameSurnamePatronimic"> <br>


        <!-- <div class="modal-createhoz">
            <div class="background-modal">
                <div class="registrate-title">Создание хозяйства</div>
                <label>Холдинг</label><br>
                <input type="text" class="registration-field"><br>
                <label>Номер хозяйства</label><br>
                <input type="number" class="registration-field"><br>
            </div>
        </div> -->

        <label>Электронная почта <span class="required">*</span></label><br>
        <input type="email" placeholder="example@email.com" class="registration-field" v-model="newUser.email"> <br>
        <label>Телефон <span class="required">*</span></label><br>
        <input type="tel" 
        pattern="^(\+7|8)?[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$" 
        placeholder="8 999 999 99 99"
        class="registration-field"
        v-model="newUser.phone"> <br>
        <label>Пароль <span class="required">*</span></label><br>
        <input type="password"  class="registration-field" v-model="newUser.password"> <br>
        <label>Повторите пароль <span class="required">*</span></label><br>
        <input type="password"  class="registration-field" v-model="repeatPassword"> <br>
        <div v-if="notSamePasswords" class="error">Пароли не совпадают</div>
        <label>Ваши ожидания от проекта генетической селекции</label><br>
        <textarea  class="registration-field textarea"></textarea> <br>
        <button type="submit" class="registr-btn">Отправить</button>
        <div v-if="notAllFields" class="error">Пожалуйста, заполните все необходимые поля</div>
        <div v-if="success" class="success">На указанную почту было отправлено письмо. Для подтверждения регистрации перейдите по ссылке, указанной в письме.</div>
    </form></div>
</template>

<script>
import MultiselectReg from '@/components/MultiselectReg.vue';
import MultiselectHozRegister from '@/MultiselectHozRegister.vue';

export default {
    data(){
        return{
            newUser: {
                email: null,
                hozNumber: '0',
                nameSurnamePatronimic: null,
                password: null,
                phone: null,
                regionId: null,
                roleId: null,
            },
            newHoz: null,
            newHold: null,

            repeatPassword: null,

            notAllFields: false,
            notSamePasswords: false,
            success: false,
        }
    },
    components: {
        MultiselectReg, MultiselectHozRegister
    },
    methods: {
        valid() {
            document.getElementById('fio').value = document.getElementById('fio').value.replace(/[\d]/g,'');
        },
        async handleFormSubmit(event){ 
            event.preventDefault();

            this.notAllFields = false;
            this.notSamePasswords = false;
            this.success = false;

            if (this.newUser.hozNumber === null) {
                this.newUser.hozNumber = '0';
            }

            for (let key in this.newUser) {
                if (this.newUser[key] === null) {
                    this.notAllFields = true;
                    return 0;
                }
            }

            if(this.newUser.password !== this.repeatPassword) {
                this.notSamePasswords = true;
                return 0;
            }
            let obj = {};
            obj.newUser = this.newUser;
            obj.newHoz = this.newHoz;
            obj.newHold = this.newHold;
            let response = await fetch(`/api/user/create`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(obj),
            })
            
            if(response.status == 200) {
                this.success = true;
            }

        },
        setIdReg(new_val){
            if (new_val) {
                this.newUser.regionId = new_val;
            } else {
                this.newUser.regionId = null;
            }
        },
        setIdHoz(new_val){
            if (new_val) {
                this.newUser.hozNumber = new_val;
            } else {
                this.newUser.hozNumber = null;
            }
        },

    },
    mounted() {
        let formRegistrate = document.getElementById('form-registrate');
        formRegistrate.addEventListener('submit', this.handleFormSubmit);
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

    .hoz {
        margin: 15px 20px 0 20px;
    }

    .create-hoz {
        text-decoration: underline;
        cursor: pointer;
        transition: 0.3s;
    }
    .create-hoz:hover {
        color: black;
    }

    .hoz-underline {
        font-size: 90%;
        color: grey;
        margin: 0 20px 15px 0;
    }

    .modal-createhoz {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(0,0,0, 0.3);
        z-index: 40;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .background-modal {
        background-color: white;
        border-radius: 20px;
        padding: 50px 80px;
    }

    .required {
        color: red;
    }

    .error {
        font-size: 90%;
        padding: 10px 0;
        color: red;
    }

    .success {
        font-size: 90%;
        padding: 10px 0;
        color: green;
        max-width: 350px;
    }
</style>