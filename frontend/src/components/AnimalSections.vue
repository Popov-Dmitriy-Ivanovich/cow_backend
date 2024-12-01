<template>
<div class="sections">
    <ul class="animal-links">
        <li @click="changeSection('common')">Общая информация</li>
        <li @click="changeSection('rating')">Оценки</li>
        <li @click="changeSection('genetic')">Генетика</li>
        <li @click="changeSection('exterior')">Экстерьер</li>
        <li @click="changeSection('health')">Здоровье</li>
        <li @click="changeSection('milking')">Контрольные доения и лактации</li>
        <li @click="changeSection('lineage')">Родословная</li>
        <li @click="changeSection('documents')">Документы</li> 
    </ul>
</div>
</template>

<script>
export default {
    methods: {
        changeSection(section_name) {
            this.$emit('changeSection', section_name);
        }
    }
}
</script>

<style scoped>
.sections {
    font-family: Open Sans, sans-serif;
    position: sticky;
    top: 100px;
}

.animal-links {
    font-size: 110%;
    color:rgb(37, 0, 132);
    list-style: none;
}

.animal-links li {
    padding: 10px 5px;
    width: max-content;
    cursor: pointer;
    transition: 0.3s;
}

.animal-links li:hover {
    color:rgb(103, 63, 205);
}
</style>