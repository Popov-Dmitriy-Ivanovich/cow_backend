<template>
    <div>
        <div class="id-title">Идентификаторы</div>
        <div class="idnum-flex">
            <div class="item-block">
                <div class="id-min-title">Инв. номер</div>
                <div>{{ cow_info.InventoryNumber || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">ID Селекса</div>
                <div>{{ cow_info.SelecsNumber || 'Нет информации' }}</div>
            </div>
            <!-- <div class="item-block">
                <div class="id-min-title">UN</div>
                <div>test</div>
            </div> -->
            <div class="item-block">
                <div class="id-min-title">Хозяйство</div>
                <div>{{ cow_info.HozHame || 'Нет информации' }}</div>
            </div>
            <div class="item-block">
                <div class="id-min-title">Ферма</div>
                <div>{{ cow_info.FarmName || 'Нет информации' }}</div>
            </div>
        </div>
    </div>
</template>
    
<script>
export default {
    props: {
        cow_info: {
            type: Object,
        }
    }
}
</script>

<style scoped>
.id-title {
    font-size: 130%;
    color: rgb(10, 113, 75);
    padding-bottom: 30px;
    width: max-content;
}

.idnum-flex {
    display: flex;
    flex-wrap: wrap;
}

.item-block {
    width: max-content;
    margin: 0 35px 30px 0;
}

.id-min-title {
    color: grey;
    margin-bottom: 7px;
}
</style>