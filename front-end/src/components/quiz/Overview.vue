<template>
  <div class="q-pa-md">
    <q-stepper
      ref="stepper"
      v-model="step"
      alternative-labels
      color="primary"
      animated
    >
      <q-step
        :name="1"
        title="Milestone 1"
        icon="settings"
        :done="step > 1"
      >
        The description for step 1
      </q-step>

      <q-step
        :name="2"
        title="Milestone 2"
        caption="Optional"
        icon="create_new_folder"
        :done="step > 2"
      >
        The description for step 2
      </q-step>

      <q-step
        :name="3"
        title="Milestone 3"
        icon="add_comment"
        :done="step > 3"
      >
        The description for step 3
      </q-step>

      <q-step
        :name="4"
        title="Milestone 4"
        icon="add_comment"
        :done="step > 4"
      >
        The description for step 4
      </q-step>

      <q-step
        :name="5"
        title="Milestone 5"
        icon="add_comment"
        :done="step > 5"
      >
        The description for step 5
      </q-step>

      <template v-slot:navigation>
        <q-stepper-navigation>
          <q-btn
            color="primary"
            :label="step === 5 ? 'Finish' : 'Continue'"
            @click="$refs.stepper.next()"
          />
          <q-btn
            v-if="step > 1"
            flat
            color="primary"
            class="q-ml-sm"
            label="Back"
            @click="$refs.stepper.previous()"
          />
        </q-stepper-navigation>
      </template>
    </q-stepper>
  </div>
</template>

<script type="text/javascript">
// import { mapGetters } from 'vuex'

export default {
  name: 'StudentOverview',
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      step: 1
    }
  },
  computed: {
  },
  methods: {
  }
}
</script>

<style>
</style>
