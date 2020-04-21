<template>
  <div class="q-px-sm">
    <add-quiz
      v-if="addQuizActive"
      @addQuizDone="fetchQuiz"
    />
    <q-table
      title="Quiz Table"
      :data="tableData"
      :columns="columns"
      :visible-columns="visibleColumns"
      separator="cell"
      wrap-cells
      row-key="id"
      :table-header-style="{ backgroundColor: 'lightblue' }"
      fixed-center
      bordered
      dense
      :loading="loading"
    >
      <template v-slot:top="props">
        <div class="col-2 q-table__title">
          Avaialble Quizzes
        </div>
        <q-space />
        <q-btn
          class="q-ml-md"
          flat
          round
          dense
          color="blue"
          :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
          @click="props.toggleFullscreen"
        />
      </template>
      <q-td
        slot="body-cell-Approver"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          v-if="props.row.Status=='WIP'"
          no-caps
          dense
          color="blue"
          label="Finish"
          @click="toUpdateQuiz(props.row.QuizId, 'Finish')"
        />
        <q-btn
          v-else-if="props.row.Status=='Finish'"
          no-caps
          dense
          :disable="!addQuizActive"
          color="blue"
          label="Approval"
          @click="toUpdateQuiz(props.row.QuizId, 'Ready')"
        />
        <q-btn
          v-else-if="props.row.Status=='Ready'"
          flat
          no-caps
          dense
          color="blue"
        >
          {{ props.row.Approver }}
        </q-btn>
      </q-td>
      <q-td
        slot="body-cell-edit"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          flat
          icon="edit"
          :disable="props.row.Status=='Ready' || props.row.Status=='Finish'"
          @click="toEditQuiz(props.row.QuizId)"
        />
      </q-td>
      <q-td
        slot="body-cell-delete"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          v-if="props.row.Status=='NA'"
          flat
          icon="delete"
          @click="toUpdateQuiz(props.row.QuizId, 'Delete')"
        />
      </q-td>
      <template v-slot:no-data="{ icon, message, filter }">
        <div class="full-width row flex-center text-accent q-gutter-sm">
          <span>
            You don't have any student, add student by press
          </span>
          <q-icon
            size="2em"
            name="add"
          />
        </div>
      </template>
    </q-table>
    <edit-quiz
      :active="editQuizActive"
      :quizId="quizId"
      @quizEditCompleted="editQuizDone"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapMutations } from 'vuex'
import { GET_QUIZ_QUERY } from '../../graphql/queries'
import { UPDATE_QUIZ_MUTATION } from '../../graphql/mutations'

export default {
  name: 'MathQuizTable',
  components: {
    'edit-quiz': require('components/quiz/EditQuiz.vue').default,
    'add-quiz': require('components/quiz/AddQuiz.vue').default
  },
  data () {
    return {
      loading: true,
      editQuizActive: false,
      quizId: '',
      skipQueryGetQuiz: true,
      columns: [
        {
          name: 'QuizId',
          required: true,
          align: 'center',
          label: 'Quiz Id',
          field: 'QuizId',
          sortable: true
        },
        {
          name: 'Grade',
          required: true,
          align: 'center',
          label: 'Grade',
          field: 'Grade',
          sortable: true
        },
        {
          name: 'Desc',
          align: 'center',
          label: 'Description',
          field: 'Desc',
          sortable: false
        },
        {
          name: 'Operator',
          align: 'center',
          label: 'Operator',
          field: 'Operator',
          sortable: false
        },
        {
          name: 'Status',
          align: 'center',
          label: 'Status',
          field: 'Status',
          sortable: true
        },
        {
          name: 'Approver',
          align: 'center',
          label: 'Approver',
          field: 'Approver',
          sortable: true
        },
        {
          name: 'edit',
          align: 'center',
          label: 'Edit Quiz',
          field: 'edit',
          sortable: false
        },
        {
          name: 'delete',
          align: 'center',
          label: 'Delete Quiz',
          field: 'delete',
          sortable: false
        }
      ]
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    ...mapGetters('quiz', ['getQuizList']),
    tableData: function () {
      console.log('quiz list = ', this.getQuizList)
      return this.getQuizList
    },
    addQuizActive: function () {
      if (this.currentUser.Role) {
        if (this.currentUser.Role === 'admin' || this.currentUser.Role === 'staff') {
          return true
        }
      }
      return false
    },
    visibleColumns: function () {
      if (this.addQuizActive === true) {
        return ['QuizId', 'Grade', 'Desc', 'Operator', 'Status', 'Approver', 'edit', 'delete']
      } else {
        return ['QuizId', 'Grade', 'Desc', 'Operator', 'Status', 'Approver', 'edit']
      }
    }
  },
  mounted () {
    this.fetchQuiz()
  },
  methods: {
    ...mapMutations('quiz', ['addNewQuiz', 'setQuizList', 'removeQuiz']),
    toEditQuiz (opt) {
      console.log('set editquiz true - ', opt)
      this.quizId = opt
      this.editQuizActive = true
    },
    editQuizDone () {
      this.editQuizActive = false
    },
    fetchQuiz () {
      console.log('Get quiz list')
      this.$apollo
        .query({
          query: GET_QUIZ_QUERY,
          variables: {
            gid: this.currentUser.Gid
          }
        })
        .then(response => {
          if (response.data.GetQuiz) {
            this.updateQuiz(response.data.GetQuiz)
            console.log('Read quiz list complete')
          } else {
            console.log('Get quiz failed')
          }
        })
        .catch(error => {
          console.log(error)
        })
    },
    updateQuiz (newList) {
      console.log('in update quiz length - ', newList.length)
      if (newList.length > 0) {
        this.loading = false
        this.setQuizList([])
        for (let i = 0; i < newList.length; i++) {
          this.addNewQuiz(newList[i])
        }
      } else {
        console.log('NO enough quiz')
      }
    },
    toUpdateQuiz (qid, status) {
      this.$apollo
        .mutate({
          mutation: UPDATE_QUIZ_MUTATION,
          variables: {
            Gid: this.currentUser.Gid,
            QuizId: qid,
            Status: status
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.UpdateQuiz) {
              console.log('Update quiz successful')
              if (status === 'Delete') {
                this.removeQuiz({
                  QuizId: qid
                })
              }
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })
    }
  }
}
</script>

<style>
</style>
