<template>
<div class="hoz-item">
    <div class="hoz-text">
        <div class="hoz-title"> 
            {{ farm_item[4] }}
        </div>
        <div class="hoz-description">
            {{ farm_item[11] }}
        </div>
    </div>
</div>
</template>

<script>
export default {
    props: {
        farm_item: {
            type: Array,
            Required: true,
        }
    }
}
</script>

<style>
.hoz-item {
    width: 450px;
    height: 320px;
    background-color: rgb(255, 255, 255);
    border-radius: 30px;
    box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
}

.hoz-text {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 10px 15px;
}

.hoz-title {
    font-size: 150%;
    margin-bottom: 10px;
    padding-bottom: 5px;
    width: 100%;
    text-align: center;
    border-bottom: 1px solid rgb(198, 195, 235);
}

.hoz-description {
    color: black;
    font-size: 90%;
}
</style>