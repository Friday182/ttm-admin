<template>
  <div class="fixed-center text-center">
    <p class="text-faded">
      Under constructure ...
    </p>
    <q-btn
      color="secondary"
      style="width:200px;"
      @click="$router.push('/')"
    >
      Go back
    </q-btn>
  </div>
</template>

<script>
export default {
  name: 'Dashboard'
}
</script>
