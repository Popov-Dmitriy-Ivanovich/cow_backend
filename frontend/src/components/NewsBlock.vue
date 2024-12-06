<template>
<div class="news">
    <div class="news-title">Новости</div>
    <div class="news-content">
        <div class="news-item"  v-for="item in news" :key="item[0]">
            <div class="news-icon">✎</div>
            <div class="news-text">
                <div class="piece-news-title">{{ item.Title }}</div>
                <div class="news-description">{{ item.Text }}</div>
            </div>
            <div class="news-date">{{ dateConverter(item.Date) }}</div>
        </div>
    </div>
</div>
</template>

<script>
export default {
    data(){
        return{
            news: [],
        }
    },
    methods: {
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        }
    },
    async created() {
        let id_region = 1;
        const response = await fetch(`/api/regions/${id_region}/news`);
        const news_result = await response.json();
        this.news = news_result;
    }
}
</script>

<style>
.news {
    width: 100%;
    height: max-content;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    font-family: Open Sans, sans-serif;
    color: rgb(37, 0, 132);
}

.news-title {
    font-size: 190%;
    padding: 40px 0 20px 0;
}

.news-content {
    width: 900px;
    height: max-content;
    background-color: white;
    box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
    border-radius: 10px;
    padding: 10px 0;
}

.news-item {
    padding: 15px 20px;
    display: flex;
    align-items: center;
    justify-content: space-around;
}

.news-icon {
    align-self: flex-start;
}

.news-text {
    width: 85%;
}

.piece-news-title {
    font-size: 120%;
    padding-bottom: 10px;
}

.news-description {
    color: black;
    font-size: 90%;
}

.news-date {
    width: 90px;
    font-size: 90%;
}
</style>