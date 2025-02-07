<template>
<div class="common-title">Подбор пар</div>
<div class="general-info">
    <MatchingCow/>
</div>
</template>

<script>
import MatchingCow from '@/components/componentsConcretAnimal/MatchingCow.vue'

export default {
    components: {
        MatchingCow
    }
}
</script>

<style>
.common-title {
    font-size: 190%;
    margin-bottom: 30px;
}

.general-info {
    background-color: white;
    width: 56vw;
    min-width: 800px;
    height: max-content;
    min-height: 500px;
    border-radius: 10px;
    box-shadow: rgba(100, 100, 111, 0.1) 0px 7px 29px 0px;
    padding: 30px 40px;
}
</style>