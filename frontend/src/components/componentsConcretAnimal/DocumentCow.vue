<template>
    <div>
        <div class="doc-title">Документы</div>
        <div v-for="doc in cow_info" :key="doc.ID" class="doc-line">
            <div class="doc-name">{{ doc.Path || 'Нет информации' }}</div>
            <div class="doc-btns">
                <a :href="`/api/static/documents/${doc.Path}`" :download="doc.Path" class="doc-link"><div class="doc-download btn">Скачать</div></a>
                <!-- <div class="doc-delete btn">Удалить</div> -->
            </div>

        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: [],
        }
    },
    async created() {
        let response = await fetch(`/api/cows/${this.$route.params.id}/documents`);
        let result = await response.json();
        console.log(result);
        this.cow_info = result;
    },
    methods: {

    }
}
</script>

<style scoped>
.doc-line {
    display: flex;
    align-items: center;
    margin: 10px 20px 0 0;
}

.doc-name {
    width: 70%;
    height: max-content;
    background-color: rgb(241, 240, 246);
    padding: 7px 0 7px 7px;
}

.doc-btns {
    display: flex;
    margin-left: 20px;
}

.btn {
    border: 1px solid rgb(207, 203, 217);
    padding: 7px;
    border-radius: 10px;
    margin-right: 10px;
    cursor: pointer;
    transition: 0.3s;
}

.doc-download {
    background-color: rgb(245, 244, 248);
}

.doc-download:hover {
    background-color: rgb(91, 84, 164);
    color: white;
}

.doc-delete:hover {
    background-color: rgb(230, 223, 248);
}

.doc-link {
    color: black;
    text-decoration: none;
}

.doc-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 10px;
}
</style>