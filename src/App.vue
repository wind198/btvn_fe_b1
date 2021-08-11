
<template>
  <div class="container">
    <Monitor class="item" :value="display" />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="C"
      id="buttonClear"
    />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="/"
      id="buttonDivide"
    />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="*"
      id="buttonMultiply"
    />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="-"
      id="buttonMinus"
    />
    <Button @button-clicked="handleClick" class="item" value="7" id="button7" />
    <Button @button-clicked="handleClick" class="item" value="8" id="button8" />
    <Button @button-clicked="handleClick" class="item" value="9" id="button9" />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="+"
      id="buttonPlus"
    />
    <Button @button-clicked="handleClick" class="item" value="4" id="button4" />
    <Button @button-clicked="handleClick" class="item" value="5" id="button5" />
    <Button @button-clicked="handleClick" class="item" value="6" id="button6" />
    <Button @button-clicked="handleClick" class="item" value="1" id="button1" />
    <Button @button-clicked="handleClick" class="item" value="2" id="button2" />
    <Button @button-clicked="handleClick" class="item" value="3" id="button3" />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="="
      id="buttonEval"
    />
    <Button @button-clicked="handleClick" class="item" value="0" id="button0" />
    <Button
      @button-clicked="handleClick"
      class="item"
      value="."
      id="buttonDecimal"
    />
  </div>
</template>

<script>
import Button from './components/Button.vue'
import Monitor from './components/Monitor.vue'
import postToSever from './postToSever.js'
export default {
  name: 'App',
  components: {
    Button,
    Monitor
  },
  data () {
    return {
      display: "",
      expression: "",
      result: ""
    }
  },
  methods: {
    async handleClick (id) {
      switch (id) {
        case "buttonClear":
          this.expression = "";
          this.display = this.expression;
          break;
        case "buttonEval":
          postToSever(this.expression).then((data)=>data).then((data)=>{this.result=data;this.display=this.result;this.expression=""})
          break;
        default:
          this.expression += document.getElementById(id).innerHTML
          this.display = this.expression
      }
    },
  },
  emits: ["button-clicked"]



}
</script>

<style lang="scss">
html {
  font-family: Arial, sans-serif;
}
.container {
  display: inline-grid;
  grid-template-columns: 50px 50px 50px 50px;
  grid-template-rows: 50px 50px 50px 50px 50px 50px;
  row-gap: 10px;
  column-gap: 10px;
  background-color: rgb(231, 231, 231);
  padding: 0.5rem;
  margin: 10px 50px;
  border-radius: 10px;
  border: 2px solid rgb(129, 129, 129, 0.5);
  box-shadow: 5px 5px 5px black;
  .item {
    background: #777;
    box-shadow: 2px 2px 2px #777;
    border-radius: 5px;
    color: white;
    line-height: 50px;
    text-align: center;
    font-size: 1.1rem;
    cursor: pointer;
    &:hover {
      background: rgb(151, 151, 151);
    }
  }
  #buttonPlus {
    grid-column: 4/5;
    grid-row: 3/5;
    line-height: 100px;
  }
  #monitor {
    grid-column: 1/5;
    grid-row: 1/2;
    text-align: right;
    padding-right: 0.5rem;
    font-size: 1.4rem;
    background-color: rgb(33, 45, 219);
    border: 1px solid rgb(168, 167, 167);
    color: black;
    box-shadow: 1px 1px 1px #777;
    cursor: text;
    &:hover {
      background: rgb(33, 45, 219);
    }
  }
  #button0 {
    grid-column: 3/4;
  }
}
</style>
