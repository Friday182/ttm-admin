<template>
  <div class="q-mx-lg q-px-lg">
    <q-table
      title="Math"
      :data="assignmentData"
      :columns="columns"
      :visible-columns="visibleColumns"
      separator="cell"
      hide-bottom
      row-key="name"
      :table-header-style="{ backgroundColor: 'lightblue' }"
      fixed-center
      bordered
      dense
      rows-per-page-label="Page:"
      :pagination.sync="pagination"
      :loading="loading"
    >
      <template v-slot:top="props">
        <div class="col-4 text-blue q-table__title">
          Math
        </div>
        <q-space />
        <q-btn
          flat
          disable
        >
          Level: {{ mathLevel }}
        </q-btn>
      </template>
      <!-- slot name syntax: body-cell-<column_name> -->
      <template v-slot:body="props">
        <q-tr
          :props="props"
          @mousedown.native.prevent="rowSelected(props.row.gid)"
        >
          <q-td
            key="task"
            :props="props"
          >
            {{ props.row.task }}
          </q-td>
          <q-td
            key="description"
            :props="props"
          >
            {{ props.row.description }}
          </q-td>
          <q-td
            key="play"
            :props="props"
          >
            <q-btn
              flat
              icon="play_circle_outline"
              @click="toRunTask(props.row.task)"
            />
          </q-td>
          <q-td
            key="done"
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
        </q-tr>
      </template>
    </q-table>
    <task-log
      :gid="currentStudent.gid"
      :subject="0"
      :num-log="numLog"
      :new-log="newLog"
    />
    <run-task
      :active="active"
      :taskid="currentTaskId"
      :subject="0"
      @completed="processTaskClosed"
    />
    <run-quiz
      :active="quizActive"
      :taskid="currentTaskId"
      :subject="0"
      @completed="processTaskClosed"
    />
    <div
      v-if="devMode"
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
        label="dog."
        style="width: 40px; font-size: 10pt"
      />
    </div>
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'MathAssignment',
  components: {
    'run-task': require('components/student/RunTask.vue').default,
    'run-quiz': require('components/student/RunQuiz.vue').default,
    'task-log': require('components/common/TasklogTable.vue').default
  },
  data () {
    return {
      devMode: false,
      pagination: {
        page: 0,
        rowsPerPage: 10
      },
      text1: '',
      text2: '',
      text3: '',
      text4: '',
      active: false,
      quizActive: false,
      currentTaskId: 'NA',
      loading: false,
      newLog: 0,
      numLog: 10,
      visibleColumns: ['task', 'description', 'play', 'done'],
      columns: [
        {
          name: 'task',
          required: true,
          label: 'Task',
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
    assignmentData: {
      // getter
      get: function () {
        let tmp = []
        if (this.currentStudent.subjects[0].IsEnabled === false) {
          return tmp
        }
        if (this.currentStudent.subjects[0].AiEnabled === true) {
          for (let i = 0; i < this.currentStudent.subjects[0].Assignment.length; i++) {
            if (i === 0) {
              tmp.push({
                task: this.currentStudent.subjects[0].Assignment[i].Kp,
                description: this.currentStudent.subjects[0].Assignment[i].Desc,
                done: this.currentStudent.subjects[0].Assignment[0].Done
              })
            } else {
              tmp[0].task = tmp[0].task + ',' + this.currentStudent.subjects[0].Assignment[i].Kp
              tmp[0].description = tmp[0].description + ' / ' + this.currentStudent.subjects[0].Assignment[i].Desc
            }
          }
        } else {
          for (let i = 0; i < this.currentStudent.subjects[0].Assignment.length; i++) {
            tmp.push({
              task: this.currentStudent.subjects[0].Assignment[i].Kp,
              description: this.currentStudent.subjects[0].Assignment[i].Desc,
              done: this.currentStudent.subjects[0].Assignment[i].Done
            })
          }
        }
        return tmp
      },
      // setter
      set: function (newValue) {
      }
    },
    mathLevel: function () {
      return this.currentStudent.subjects[0].level
    }
  },
  mounted () {
    console.log('mounted MathTab')
  },
  methods: {
    ...mapActions('student', ['updateAssignmentDone']),

    toRunTask (task) {
      if (task.indexOf('MQ') !== -1) {
        this.$q.fullscreen.request()
        this.currentTaskId = task
        this.quizActive = true
        this.active = false
      } else {
        this.$q.fullscreen.request()
        this.currentTaskId = task
        this.quizActive = false
        this.active = true
      }
    },
    processTaskClosed (opt) {
      if (opt >= 2) {
        this.newLog++
        this.updateAssignmentDone({
          Kp: this.currentTaskId,
          Math: true
        })
      }
      this.$q.fullscreen.exit()
      this.currentTaskId = 'NA'
      this.active = false
      this.quizActive = false
    },
    rowSelected (gid) {
    }
  }
}
</script>

<style>
</style>
