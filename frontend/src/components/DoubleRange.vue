<template>
<div class="parent-range">
    <div class="price-input">
        <div class="field">
          <!-- <span>Min</span> -->
          <input type="number" class="input-min" v-model="currentValueMin">
        </div>
        <!-- <div class="separator">-</div> -->
        <div class="field">
          <!-- <span>Max</span> -->
          <input type="number" class="input-max" v-model="currentValueMax">
        </div>
    </div>
    <div class="slider">
        <div class="progress" ref="range"></div>
    </div>
    <div class="range-input">
        <input type="range" class="range-min" :min="minValue" :max="maxValue" v-model="currentValueMin">
        <input type="range" class="range-max" :min="minValue" :max="maxValue" v-model="currentValueMax">
    </div>
</div>
</template>

<script>
export default {
    data() {
        return {
            currentValueMin: 0,
            currentValueMax: 2,
            maxValue: 10,
            minValue: 0,
        }
    },
    methods: {
        setRange() {
            const rangeInput = document.querySelectorAll(".range-input input"),
            priceInput = document.querySelectorAll(".price-input input");
            // range = document.querySelector(".progress");
            let priceGap = 1;
            priceInput.forEach(input =>{
                input.addEventListener("input", e =>{  
                    if((this.currentValueMax - this.currentValueMin >= priceGap) && this.currentValueMax <= this.maxValue){
                        if(e.target.className === "input-min"){
                            this.$refs.range.style.left = ((this.currentValueMin / this.maxValue) * 100) + "%";
                        }else{
                            this.$refs.range.style.right = 100 - (this.currentValueMax / this.maxValue) * 100 + "%";
                        }
                    }
                    if (this.currentValueMax >= this.maxValue) {
                        this.currentValueMax = this.maxValue;
                    }
                    if (this.currentValueMin <= this.minValue) {
                        this.currentValueMin = this.minValue;
                    }
                });
            });
            rangeInput.forEach(input =>{
                input.addEventListener("input", e =>{
                    if((this.currentValueMax - this.currentValueMin) < priceGap){
                        if(e.target.className === "range-min"){
                            this.currentValueMin = this.currentValueMax - priceGap
                        }else{
                            this.currentValueMax = this.currentValueMin + priceGap;
                        }
                    }else{
                        this.$refs.range.style.left = ((this.currentValueMin / this.maxValue) * 100) + "%";
                        this.$refs.range.style.right = 100 - (this.currentValueMax / this.maxValue) * 100 + "%";
                    }
                });
            });
        }
    },
    mounted() {
        this.setRange();
    }
}
</script>

<style scoped>
.parent-range {
    width: 230px;
}

.price-input{
  width: 100%;
  display: flex;
  margin: 30px 0 35px;
}
.price-input .field{
  display: flex;
  width: 100%;
  height: 25px;
  align-items: center;
}
.field input{
  width: 100%;
  height: 100%;
  outline: none;
  font-size: 19px;
  margin-left: 12px;
  border-radius: 5px;
  text-align: center;
  border: 1px solid #999;
  -moz-appearance: textfield;
}
input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
  -webkit-appearance: none;
}
.price-input .separator{
  width: 130px;
  display: flex;
  font-size: 19px;
  align-items: center;
  justify-content: center;
}
.slider{
  height: 5px;
  position: relative;
  background: #ddd;
  border-radius: 5px;
}
.slider .progress{
  height: 100%;
  left: 25%;
  right: 25%;
  position: absolute;
  border-radius: 5px;
  background: #8769ac;
}
.range-input{
  position: relative;
}
.range-input input{
  position: absolute;
  width: 100%;
  height: 5px;
  top: -5px;
  background: none;
  pointer-events: none;
  -webkit-appearance: none;
  -moz-appearance: none;
}
input[type="range"]::-webkit-slider-thumb{
  height: 17px;
  width: 17px;
  border-radius: 50%;
  background: #8769ac;
  pointer-events: auto;
  -webkit-appearance: none;
  box-shadow: 0 0 6px rgba(0,0,0,0.05);
}
input[type="range"]::-moz-range-thumb{
  height: 17px;
  width: 17px;
  border: none;
  border-radius: 50%;
  background: #8769ac;
  pointer-events: auto;
  -moz-appearance: none;
  box-shadow: 0 0 6px rgba(0,0,0,0.05);
}
</style>