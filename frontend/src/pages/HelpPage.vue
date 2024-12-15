<template>
    <div class="help-page">
        <div class="help-title">
            <div class="help-title-text">Помощь</div>
        </div>
        <div class="help-section">
            <div class="helpsection-title">Животные, поиск и выбор</div>
            <div class="help-flexblocks">
                <HelpBlock>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как попасть на вкладку животные?
                        </template>
                        <template v-slot:answer>
                            Вверху любой страницы сайта есть горизонтальное меню, после логина пользователя в этом меню появятся вкладки «Животные», «Аналитика», «Загрузка файлов». Для того, чтобы перейти к поиску и выбору животных необходимо перейти на вкладку «Животные».
                        </template>
                    </HelpQuestion>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как найти нужное КРС?
                        </template>
                        <template v-slot:answer>
                            Вкладка животные поддерживает фильтрацию животных по опреленному пользователем признаку. Может показать животных выбранного пользователем пола. Поддерживает поиск животного по селексу, идентификационному номеру, РСХН номеру, кличке.
                        </template>
                    </HelpQuestion>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как узнать полную информацию о КРС?
                        </template>
                        <template v-slot:answer>
                            В блоке отображения животных выберете строку с интересующим вас КРС, сайт перенаправит пользователя на страницу 
                            с информацией о животном. Пользователю доступны следующие блоки информации: «Общая информация» 
                            (основные данные о КРС, такие как порода, кличка, факт генотипирования и т. д.), «Оценки» (данные об оценках ebv коровы 
                            по региону и по хозяйству), «Генетика» (данные о моногенных заболеваниях и gtc файл), «Экстерьер» (данные об экстерьере: 
                            фото и оценки параметров), «Здоровье» (данные о ветеренарных событиях), «Контрольные доения и лактации» (список всех 
                            лактаций и контрольных доений выбранной коровы с возможностью построить граффик), «Родословная» (дерево предков коровы 
                            до прародителей и список потомков), «Документы» (загруженные фермером документы).
                        </template>
                    </HelpQuestion>
                </HelpBlock>
            </div>
        </div>   
        <div class="help-section">
            <div class="helpsection-title">Аналитика, диаграммы и отчеты</div>
            <div class="help-flexblocks">
                <HelpBlock>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как попасть на вкладку Аналитика?
                        </template>
                        <template v-slot:answer>
                            В верху любой страницы сайта есть горизонтальное меню, после логина пользователя в этом меню появятся вкладки «Животные», «Аналитика», 
                            «Загрузка файлов». Для того, чтобы перейти к аналитике необходимо перейти на вкладку «Аналитика».
                        </template>
                    </HelpQuestion>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как собрать аналитику только по интересующим животным?
                        </template>
                        <template v-slot:answer>
                            В правой части вкладки аналитика находятся фильтры, аналитика собирается только по живым КРС, удовлетворяющим условиям фильтров.
                        </template>
                    </HelpQuestion>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как перейти в аналитику по году, региону, району?
                        </template>
                        <template v-slot:answer>
                            При нажатии на столбец аналитики сайт перенаправит вас на аналитику по параметру соответсвующему значению оси X выбранного столбца.
                        </template>
                    </HelpQuestion>
                </HelpBlock>
            </div>
        </div> 
        <div class="help-section">
            <div class="helpsection-title">Загрузка и обновление данных</div>
            <div class="help-flexblocks">
                <HelpBlock>
                    <HelpQuestion>
                        <template v-slot:title>
                            Как попасть во вкладку «Загрузка файлов»?
                        </template>
                        <template v-slot:answer>
                            В верху любой страницы сайта есть горизонтальное меню, после логина пользователя в этом меню появятся вкладки «Животные», «Аналитика», «Загрузка файлов». 
                            Для того, чтобы перейти к загрузке файлов нужно перейти на вкладку «Загрузка файлов».
                        </template>
                    </HelpQuestion>
                </HelpBlock>
            </div>
        </div> 
    </div>
</template>

<script>
import HelpBlock from '@/components/HelpBlock.vue';
import HelpQuestion from '@/components/HelpQuestion.vue';

export default {
    components: {
        HelpBlock, HelpQuestion
    }
}
</script>

<style scoped>
.help-page {
    margin-top: 130px;
    font-family: Open Sans, sans-serif;
    padding-left: 11%;
}

.help-title-text {
    font-size: 200%;
    padding: 10px 0 30px 0;
    color: rgb(37, 0, 132);
}

.help-section {
    margin: 20px 0;
    padding: 10px 0;
}

.helpsection-title {
    font-size: 140%;
}

.help-flexblocks {
    display: flex;
}
</style>