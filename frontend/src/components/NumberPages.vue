<template>
    <div>
        <div class="page-numbers">
            <div v-for="page in getPaginationRange" :key="page">
                <div v-if="page=='...'" class="page-dots">{{ page }}</div>
                <div v-else class="page-number" 
                @click="changeCurrentPage(page)" 
                :class="{'current-page-number':page==this.currentPage}"
                >{{ page }}</div>
            </div>
        </div>
        <div class="page-input">
            <div>Перейти на страницу с номером</div>
            <input type="number" class="input-of-npage" v-model="inputPage" @keyup.enter="changeCurrentPage(inputPage)">
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            new_curr_page: -1,
            inputPage: 1,
        }
    },
    props:{
        currentPage: {
            type: Number
        },
        totalPages: {
            type: Number
        }
    },
    methods: {
        changeCurrentPage(value) {
            this.new_curr_page = value;
            this.$emit('changePage', this.new_curr_page);
        }
    },
    computed: {
        getPaginationRange () {
            const maxPagesToShow = 5;
            let start = Math.max(1, this.currentPage - 2);
            let end = Math.min(this.totalPages, this.currentPage + 2);

            const showStartEllipsis = start > 2;
            const showEndEllipsis = end < this.totalPages - 1;

            if (end - start + 1 < maxPagesToShow) {
                if (this.currentPage <= this.totalPages / 2) {
                end = Math.min(this.totalPages, start + maxPagesToShow - 1);
                } else {
                start = Math.max(1, end - maxPagesToShow + 1);
                }
            }

            const pages = [];
            if (start > 1) pages.push(1);
            if (showStartEllipsis) pages.push("...");

            for (let i = start; i <= end; i++) {
                pages.push(i);
            }

            if (showEndEllipsis) pages.push("...");
            if (end < this.totalPages) pages.push(this.totalPages);

            return pages;
        }
    },
    watch: {
        currentPage(new_val) {
            this.inputPage = new_val;
        }
    }
}
</script>

<style scoped>
.page-numbers {
    display: flex;
    font-size: 130%;
    font-family: Open Sans, sans-serif;
    width: max-content;
}

.page-number {
    margin: 5px 10px;
    padding: 5px 15px;
    border: 2px solid rgb(195, 200, 212);
    border-radius: 8px;
    background-color: white;
    cursor: pointer;
    transition: 0.3s;
}

.page-number:hover {
    border: 2px solid rgb(101, 102, 170);
}

.page-dots {
    margin: 5px 10px;
    padding: 5px 5px;
}

.current-page-number {
    background-color: rgb(239, 236, 248);
}

.page-input {
    display: flex;
    justify-content: center;
    padding: 20px 0;
    font-family: Open Sans, sans-serif;
    align-items: center;
}

.input-of-npage {
    padding: 0 10px;
    font-size: 16px;
    box-sizing: border-box;
    outline: none;
    border: 2px solid rgb(195, 200, 212);
    border-radius: 5px;
    height: 30px;
    width: 80px;
    margin: 0 6px;
    text-align: center;
}
</style>