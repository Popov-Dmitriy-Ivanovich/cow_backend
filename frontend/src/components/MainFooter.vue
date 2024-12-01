<template>
    <footer>
        <div class="footer">
            <div class="footer-links">
                <div @click="$router.push('/help')">Помощь</div>
                <div @click="$router.push('/')">О проекте</div>
                <div>Поддержка</div>
            </div>
        </div>
    </footer>
</template>

<script>

</script>

<style scoped>
    .footer {
        height: 100px;
        width: 100%;
        background-color: white;
        margin-top: 160px;
    }

    .footer-links {
        font-family: Open Sans, sans-serif;
        font-size: 90%;
        list-style: none;
        padding: 50px 0 0 0;
        width: 70%;
        cursor: pointer;
        color: rgb(103, 103, 103);
        display: flex;
        justify-content: space-around;

        position: absolute;
        left: 50%;
        transform: translate(-50%, 0);
    }

    .footer-links div {
        margin-bottom: 10px;
        transition: 0.3s;
    }

    .footer-links div:hover {
        color: black;
    }
    
</style>