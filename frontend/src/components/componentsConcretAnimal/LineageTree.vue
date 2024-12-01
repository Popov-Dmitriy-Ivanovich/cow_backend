<template>
    <div>
        <div class="lineage-title">Родословное дерево</div>
        <div class="parent-lineagetree">
            <div class="lineage-column">
                <div class="animal-block"></div>
            </div>
            <div class="lineage-column">
                <div class="animal-block"></div>
                <div class="animal-block"></div>
            </div>
            <div class="lineage-column">
                <div class="animal-block"></div>
                <div class="animal-block"></div>
                <div class="animalwhite-block"></div>
                <div class="animal-block"></div>
                <div class="animal-block"></div>
            </div>
        </div>
    </div>
</template>
    
<script>

</script>

<style scoped>
.lineage-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.lineage-column {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.animal-block {
    width: 170px;
    height: 100px;
    background-color: rgb(223, 221, 239);
    margin: 15px 15px;
    border-radius: 10px;
}

.animalwhite-block {
    width: 170px;
    height: 70px;
}

.parent-lineagetree {
    display: flex;
}
</style>