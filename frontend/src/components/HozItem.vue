<template>
<div class="hoz-item">
    <div class="hoz-text">
        <div class="hoz-title"> 
            {{ farm_item.Name }}
        </div>
        <div class="hoz-description">
            {{ farm_item.Description }}
        </div>
        <div class="hoz-stat">
            <div>
                <div class="stat">{{ farm_item.CowsCount }} шт</div>
                <div>животных в базе</div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
export default {
    props: {
        farm_item: {
            type: Object,
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
    overflow: auto;
}

.hoz-item::-webkit-scrollbar {
    height: 12px;
}

.hoz-item::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.hoz-item::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
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

.hoz-stat {
    display: flex;
    justify-content: space-around;
    width: 100%;
    margin-top: 40px;
}

.stat {
    font-size: 300%;
    color: black;
}
</style>