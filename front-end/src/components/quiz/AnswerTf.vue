<template>
  <div>
    <q-btn
      label="True"
      glossy
      outlined
      color="green"
      class="q-mr-md"
    />
    <q-btn
      label="False"
      glossy
      outlined
      color="red"
    />
  </div>
</template>

<script type="text/javascript">

export default {
  name: 'TrueFalseAnswer',
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      ansCorrect: false
    }
  }
}
</script>

<style>
</style>
