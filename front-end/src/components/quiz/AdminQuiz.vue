<template>
  <div class="q-mx-lg q-px-lg">
    <q-table
      title="quiz"
      :data="quizData"
      :columns="columns"
      separator="cell"
      row-key="name"
      :table-header-style="{ backgroundColor: 'lightblue' }"
      fixed-center
      bordered
      dense
      no-data-label="No Quiz Available for Now."
      :hide-bottom="noBottom"
      :loading="loading"
    >
      <template v-slot:top="props">
        <div class="col-4 text-blue q-table__title">
          Available Quiz
        </div>
        <q-space />
        <q-btn
          flat
          disable
        >
          Count: {{ mathLevel }}
        </q-btn>
      </template>
      <!-- slot name syntax: body-cell-<column_name> -->
      <q-td
        slot="body-cell-play"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          flat
          icon="play_circle_outline"
          @click="toRunTask(props.row.name)"
        />
      </q-td>
      <q-td
        slot="body-cell-done"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          v-if="props.row.done"
          flat
          color="green"
          icon="done_outline"
        />
        <q-spinner-hourglass
          v-else
          color="grey"
          size="sm"
        />
      </q-td>
    </q-table>
    <quiz-log
      :gid="currentStudent.gid"
      :num-log="numLog"
      :new-log="newLog"
    />
    <runTask
      :active="active"
      :taskid="currentTaskId"
      @completed="processTaskClosed"
    />
    <div
      class="row"
    >
      <q-input
        v-model="text1"
        label="I"
        style="width: 20px"
      />
      <q-input
        v-model="text2"
        label="has"
        style="width: 40px"
      />
      <q-input
        v-model="text4"
        label="an"
        style="width: 30px"
      />
      <q-input
        v-model="text3"
        label="dogdogdogdog."
        style="width: 130px"
      />
    </div>
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'StudentAssignment',
  components: {
    'runTask': require('components/student/RunTask.vue').default,
    'quiz-log': require('components/common/QuizlogTable.vue').default
  },
  data () {
    return {
      devMode: false,
      text1: '',
      text2: '',
      text3: '',
      text4: '',
      active: false,
      currentTaskId: 'NA',
      loading: false,
      newLog: 0,
      numLog: 10,
      columns: [
        {
          name: 'quiz',
          required: true,
          label: 'Quiz',
          align: 'center',
          field: row => row.name,
          format: val => `${val}`,
          sortable: true
        },
        {
          name: 'description',
          align: 'center',
          label: 'Description',
          field: 'description',
          sortable: false
        },
        {
          name: 'play',
          align: 'center',
          label: 'play',
          field: 'play',
          sortable: false
        },
        {
          name: 'done',
          align: 'center',
          label: 'Done',
          field: 'done',
          sortable: false
        }
      ]
    }
  },
  computed: {
    ...mapGetters('student', ['currentStudent']),
    noBottom: function () {
      return (this.currentStudent.quizzes.length > 0)
    },
    quizData: {
      // getter
      get: function () {
        let tmp = []
        for (let i = 0; i < this.currentStudent.quizzes.length; i++) {
          tmp.push({
            name: this.currentStudent.quizzes[i].Kp,
            description: this.currentStudent.quizzes[i].Desc,
            done: this.currentStudent.assignmentDone
          })
        }
        return tmp
      },
      // setter
      set: function (newValue) {
      }
    },
    mathLevel: function () {
      return 3
    }
  },
  mounted () {
    console.log('mounted quizzes component')
  },
  methods: {
    ...mapActions('student', ['updateQuizDone']),

    toRunTask (task) {
      this.$q.fullscreen.request()
      this.currentTaskId = task
      this.active = true
    },
    processTaskClosed (opt) {
      if (opt >= 2) {
        this.updateQuizDone(true)
      }
      this.$q.fullscreen.exit()
      this.currentTaskId = 'NA'
      this.active = false
    }
  }
}
</script>

<style>
</style>
