<template>
    <div>
        <div class="lineage-title">Родословное дерево</div>
        <div class="parent-lineagetree">
            <div class="column">
                <div class="main-cow animal-block"></div>
            </div>
            <div class="column">
                <div class="main-cow__mother animal-block"></div>
                <div class="white-block-parent"></div>
                <div class="main-cow__father animal-block"></div>
            </div>
            
            <div class="column">
                <div class="main-cow__mother__mother animal-block"></div>
                <div class="main-cow__mother__father animal-block"></div>
                <div class="white-block-grandparent"></div>
                <div class="main-cow__father__mother animal-block"></div>
                <div class="main-cow__father__father animal-block"></div>
            </div>

            <div class="column">
                <div class="main-cow__mother__mother__mother animal-block"></div>
                <div class="main-cow__mother__mother__father animal-block"></div>
                <div class="main-cow__mother__father__mother animal-block"></div>
                <div class="main-cow__mother__father__father animal-block"></div>
                <div class="main-cow__father__mother__mother animal-block"></div>
                <div class="main-cow__father__mother__father animal-block"></div>
                <div class="main-cow__father__father__mother animal-block"></div>
                <div class="main-cow__father__father__father animal-block"></div>
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
    width: 150px;
    height: 90px;
    background-color: rgb(223, 221, 239);
    margin: 15px 15px;
    border-radius: 10px;
}

.white-block-parent {
    width: 150px;
    height: 100px;
}

.white-block-grandparent {
    width: 150px;
    height: 200px;
}

.parent-lineagetree {
    /* display: grid;
    grid-template-columns: repeat(4, 160px);
    grid-template-rows: repeat(9, 100px);
    grid-template-areas: 
    "W W W G"
    "W W C H"
    "W A D I"
    "W W W J"
    "M W W W"
    "W W W K"
    "W B E L"
    "W W F M"
    "W W W N"; */
    display: flex;
    align-items: center;
}

.main-cow {
    grid-area: M;
}
.main-cow__mother {
    grid-area: A;
}
.main-cow__father {
    grid-area: B;
}
.main-cow__mother__mother {
    grid-area: C;
}
.main-cow__mother__father {
    grid-area: D;
}
.main-cow__father__mother {
    grid-area: E;
}
.main-cow__father__father {
    grid-area: F;
}
.main-cow__mother__mother__mother {
    grid-area: G;
}
.main-cow__mother__mother__father {
    grid-area: H;
}
.main-cow__mother__father__mother {
    grid-area: I;
}
.main-cow__mother__father__father {
    grid-area: J;
}
.main-cow__father__mother__mother {
    grid-area: K;
}
.main-cow__father__mother__father {
    grid-area: L;
}
.main-cow__father__father__mother {
    grid-area: M;
}
.main-cow__father__father__father {
    grid-area: N;
}
</style>