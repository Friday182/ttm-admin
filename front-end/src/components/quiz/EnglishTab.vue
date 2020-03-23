<template>
  <div class="q-mx-lg q-px-lg">
    <q-table
      title="English"
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
          English
        </div>
        <q-space />
        <q-btn
          flat
          disable
        >
          Level: {{ englishLevel }}
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
      :subject="1"
      :num-log="numLog"
      :new-log="newLog"
    />
    <runTask
      :active="active"
      :taskid="currentTaskId"
      :subject="1"
      @completed="processTaskClosed"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'MathAssignment',
  components: {
    'runTask': require('components/student/RunTask.vue').default,
    'task-log': require('components/common/TasklogTable.vue').default
  },
  data () {
    return {
      pagination: {
        page: 0,
        rowsPerPage: 5
      },
      active: false,
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
        if (this.currentStudent.subjects[1].IsEnabled === false) {
          return tmp
        }
        if (this.currentStudent.subjects[1].AiEnabled === true) {
          for (let i = 0; i < this.currentStudent.subjects[1].Assignment.length; i++) {
            if (i === 0) {
              tmp.push({
                task: this.currentStudent.subjects[1].Assignment[i].Kp,
                description: this.currentStudent.subjects[1].Assignment[i].Desc,
                done: this.currentStudent.subjects[1].Assignment[0].Done
              })
            } else {
              tmp[0].task = tmp[0].task + ',' + this.currentStudent.subjects[1].Assignment[i].Kp
              tmp[0].description = tmp[0].description + ' / ' + this.currentStudent.subjects[1].Assignment[i].Desc
            }
          }
        } else {
          for (let i = 0; i < this.currentStudent.subjects[1].Assignment.length; i++) {
            tmp.push({
              task: this.currentStudent.subjects[1].Assignment[i].Kp,
              description: this.currentStudent.subjects[1].Assignment[i].Desc,
              done: this.currentStudent.subjects[1].Assignment[i].Done
            })
          }
        }
        return tmp
      },
      // setter
      set: function (newValue) {
      }
    },
    englishLevel: function () {
      return this.currentStudent.subjects[1].level
    }
  },
  mounted () {
    console.log('mounted EnglishTab')
  },
  methods: {
    ...mapActions('student', ['updateAssignmentDone']),

    toRunTask (task) {
      /* for (let i = 0; i < this.assignmentData.length; i++) {
        if (this.assignmentData[i].name === task) {
          if (this.assignmentData[i].done === true) {
            // are you sure you want do it again?
          }
        }
      } */
      this.$q.fullscreen.request()
      this.currentTaskId = task
      this.active = true
    },
    processTaskClosed (opt) {
      if (opt >= 2) {
        this.updateAssignmentDone({
          Math: true
        })
      }
      this.$q.fullscreen.exit()
      this.currentTaskId = 'NA'
      this.active = false
    },
    rowSelected (gid) {
    }
  }
}
</script>

<style>
</style>
