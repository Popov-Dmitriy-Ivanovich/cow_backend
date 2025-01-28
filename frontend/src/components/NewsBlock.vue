<template>
<div class="news">
    <div class="news-title">Новости</div>
    <div class="news-content">
        <!-- <div class="news-item"  v-for="item in news" :key="item[0]">
            <div class="news-icon">✎</div>
            <div class="news-text">
                <div class="piece-news-title">{{ item.Title }}</div>
                <div class="news-description"><span>{{ item.Text }}</span></div>
            </div>
            <div class="news-date">{{ dateConverter(item.Date) }}</div>
        </div> -->

        <div class="news-item">
            <div class="news-icon">✎</div>
            <div class="news-text">
                <div class="piece-news-title">В Краснодарском крае в 2025 году на субсидии для производителей молока направят 1,3 млрд. рублей</div>
                <div class="news-description"><span>Краснодарский край полностью обеспечивает молоком собственные потребности. Он выпускает больше 80% всей молочной продукции ЮФО. 
                    Господдержка позволит сельхозпредприятиям в 2025 году произвести более 1,7 миллиона тонн молока, как и в прошлом году. </span></div>
                <div class="news-description"><a href="https://msh.krasnodar.ru/news/common/s/common/e/373874">Источник</a></div>
            </div>
            <div class="news-date">10.01.2025</div>
        </div>

        <div class="news-item">
            <div class="news-icon">✎</div>
            <div class="news-text">
                <div class="piece-news-title">В 2025 году объем производства сельхозпродукции в Краснодарском крае должен превысить 670 миллиардов рублей. </div>
                <div class="news-description"><span>Также ожидается рост инвестиций, который составит почти 3,5 процента или 64 миллиарда 
                    рублей в денежном выражении. Будут запущены животноводческие комплексы по 2 тысячи голов в Усть-Лабинском и Каневском 
                    районах и виноградарско-винодельческий комплекс в Новороссийске. Начнет работу семеноводческий завод в Староминском 
                    районе, масложировые комбинаты в Брюховецком и Курганинском районах. </span></div>
                <div class="news-description"><a href="https://msh.krasnodar.ru/news/common/s/common/e/373626">Источник</a></div>
            </div>
            <div class="news-date">09.01.2025</div>
        </div>

        <div class="news-item">
            <div class="news-icon">✎</div>
            <div class="news-text">
                <div class="piece-news-title">Оксана Лут: Минсельхоз изучает меры поддержки кооперации с холдингами малого агробизнеса</div>
                <div class="news-description"><span>По словам министра, фермерское хозяйство как малое производство от кооперации с 
                    крупным холдингом сможет больше заработать, поэтому министерство изучает возможности отдельной поддержки для такой 
                    кооперации. Еще одной из ниш для малого бизнеса было названо органическое производство. Помимо этого, изучаются 
                    возможности грантов по отдельным нишевым видам животноводства.</span></div>
                <div class="news-description">Согласно материалам министерства, представленным на форуме «Развитие малого агробизнеса России», 
                    по итогам девяти месяцев 2024 года на субъекты малого агробизнеса приходится более 95% от всего количества 
                    зарегистрированных в России сельхозтоваропроизводителей: это более 160 тыс. крестьянских (фермерских) хозяйств и 
                    более 17 млн. личных подсобных хозяйств.</div>
                <div class="news-description">В 2025 году объем поддержки планируется сохранить на уровне 15 млрд. рублей, как и в 2024 году: 
                    в частности, это семь грантов («Агростартап», «Развитие семейных ферм», «Агропрогресс», «Развитие МТБ СПоК», «Агротуризм», 
                    «Развитие начинающих СПоК», «Агромотиватор»), три направления субсидирования — самозанятым, на укрепление МТБ, семейным 
                    фермам и расширение «целёвки».</div>
                <div class="news-description"><a href="https://dairynews.ru/news/oksana-lut-minselkhoz-izuchaet-mery-podderzhki-koo.html">Источник</a></div>
            </div>
            <div class="news-date">24.01.2025</div>
        </div>

        <div class="news-item">
            <div class="news-icon">✎</div>
            <div class="news-text">
                <div class="piece-news-title">Минсельхоз России запустил программу генотипирования животных</div>
                <div class="news-description"><span>Российский Минсельхоз в 2025 году запустил программу генотипирования 
                    животных для выявления их лучших свойств, в ближайшие годы ведомство планирует существенно увеличить 
                    субсидирование этой работы. Об этом глава министерства Оксана Лут рассказала на XVI съезде 
                    Союза производителей молока. На поддержку этого направления в 2025 году предусмотрено 300 млн. рублей.</span></div>
                <div class="news-description"><a href="https://dairynews.ru/news/minselkhoz-rossii-zapustil-programmu-genotipirovan.html">Источник</a></div>
            </div>
            <div class="news-date">23.01.2025</div>
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
        let id_region = 36;
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
    color: rgb(10, 113, 75);
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
    text-indent: 20px;
    padding-bottom: 10px;
    text-align: justify;
}

.news-date {
    width: 90px;
    font-size: 90%;
}
</style>