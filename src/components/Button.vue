<template>
  <div @click="clickFunc" class="item" :id="id">{{ value }}</div>
</template>

<script>
export default {
  name: "Button",
  props: {
    value: String,
    id: String
  },
  methods: {
    clickFunc () {
      this.$emit("button-clicked", this.id)
      // console.log("emited");
    }
  }

}
</script>