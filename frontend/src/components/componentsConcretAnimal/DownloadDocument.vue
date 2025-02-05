<template>
<div class="doc-title">Загрузка документов</div>    
<div>
    <form class="download-doc-form" id="form-document">
        <input type="file" id="Document" name="Document">
        <input type="text" id="CowID" name="CowID" required="" :value="$route.params.id">
        <button type="submit">Загрузить</button>
        <div v-if="uspeh">Успешно загружено!</div>
    </form>

</div>
</template>

<script>
export default {
    data() {
        return {
            uspeh: false,
        }
    },
    mounted() {
        this.uspeh = false;
        let formDocument = document.getElementById('form-document');
        formDocument.addEventListener('submit', this.sendData);
    },
    methods: {
        async sendData(event) {
            event.preventDefault();
            let formDocument = document.getElementById('form-document');
            let formData = new FormData(formDocument);

            let response = await fetch('/api/load/document', {
                method: 'POST',
                headers: {
                    'Authorization': this.getJwt()
                },
                body: formData
            })
            let result = await response.json();
            console.log(result);
            if (result == 'ok') {
                this.uspeh = true;
                this.$emit('changeUspeh', this.uspeh);
            }
        },
        getJwt() {
            let arr = document.cookie.split(';');
            for (let i = 0; i < arr.length; i++) {
                if (arr[i].split('=')[0] == 'jwt') {
                    return arr[i].split('=')[1];
                }
            }
            return null;
        }
    }
}
</script>

<style scoped>
.doc-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding-bottom: 10px;
}
.download-doc-form {
    margin: 20px 0;
}

#CowID {
    display: none;
}
</style>