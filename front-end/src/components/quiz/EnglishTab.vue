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
      :pagination.sync="pagination"
      :loading="loading"
    >
      <template v-slot:top="props">
        <div class="col-2 q-table__title">
          Avaialble Quizzes
        </div>
        <q-btn-toggle
          v-model="mathType"
          class="q-mb-md my-custom-toggle"
          push
          no-caps
          outlined
          rounded
          unelevated
          toggle-color="primary"
          color="q-blue-10"
          text-color="primary"
          size="md"
          :options="[
            {label: 'Lists', value: 'lists'},
            {label: 'Details', value: 'details'}
          ]"
        />
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
    <edit-en-quiz
      :active="editQuizActive"
      :quizId="quizId"
      @quizEditCompleted="editQuizDone"
    />
    <q-page-scroller
      position="bottom"
      :scroll-offset="150"
      :offset="[18, 18]"
    >
      <q-btn
        fab
        icon="keyboard_arrow_up"
        color="accent"
      />
    </q-page-scroller>
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapMutations } from 'vuex'
import { GET_QUIZ_QUERY } from '../../graphql/queries'
import { UPDATE_QUIZ_MUTATION } from '../../graphql/mutations'

export default {
  name: 'EnQuizTable',
  components: {
    'edit-en-quiz': require('components/quiz/EditEnQuiz.vue').default,
    'add-quiz': require('components/quiz/AddQuiz.vue').default
  },
  data () {
    return {
      loading: true,
      editQuizActive: false,
      quizId: '',
      pagination: {
        page: 0,
        rowsPerPage: 20
      },
      mathType: 'lists',
      columns: [
        { name: 'QuizId', required: true, align: 'center', label: 'Quiz Id', field: 'QuizId', sortable: true },
        { name: 'Grade', align: 'center', label: 'Grade', field: 'Grade', sortable: true },
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
        },
        { name: 'Ma', align: 'center', label: 'Basic Calculation', field: 'Ma', sortable: true },
        { name: 'Mb', align: 'center', label: 'Decimal & Fraction', field: 'Mb', sortable: true },
        { name: 'Mc', align: 'center', label: 'Algebra', field: 'Mc', sortable: true },
        { name: 'Md', align: 'center', label: 'Measurement', field: 'Md', sortable: true },
        { name: 'Me', align: 'center', label: 'Shape & Space', field: 'Me', sortable: true },
        { name: 'Mf', align: 'center', label: 'Data Handling', field: 'Mf', sortable: true },
        { name: 'Mh', align: 'center', label: 'Misc', field: 'Mh', sortable: true },
        { name: 'Mi', align: 'center', label: 'Time & Date', field: 'Mi', sortable: true }
      ]
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    ...mapGetters('quiz', ['getQuizList']),
    tableData: function () {
      console.log('quiz list = ', this.getQuizList)
      let tmp = []
      for (let i = 0; i < this.getQuizList.length; i++) {
        if (this.getQuizList[i].QuizId.indexOf('EQ_') >= 0) {
          tmp.push({
            QuizId: this.getQuizList[i].QuizId,
            Grade: this.getQuizList[i].Grade,
            Desc: this.getQuizList[i].Desc,
            Operator: this.getQuizList[i].Operator,
            Status: this.getQuizList[i].Status,
            Approver: this.getQuizList[i].Approver,
            Ma: this.getQuizList[i].Details.Ma,
            Mb: this.getQuizList[i].Details.Mb,
            Mc: this.getQuizList[i].Details.Mc,
            Md: this.getQuizList[i].Details.Md,
            Me: this.getQuizList[i].Details.Me,
            Mf: this.getQuizList[i].Details.Mf,
            Mg: this.getQuizList[i].Details.Mg,
            Mh: this.getQuizList[i].Details.Mh,
            Mi: this.getQuizList[i].Details.Mi,
            Mj: this.getQuizList[i].Details.Mj
          })
        }
      }
      return tmp
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
      if (this.mathType === 'lists') {
        if (this.addQuizActive === true) {
          return ['QuizId', 'Grade', 'Desc', 'Operator', 'Status', 'Approver', 'edit', 'delete']
        } else {
          return ['QuizId', 'Grade', 'Desc', 'Operator', 'Status', 'Approver', 'edit']
        }
      } else {
        return ['QuizId', 'Ma', 'Mb', 'Mc', 'Md', 'Me', 'Mf', 'Mh', 'Mi']
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

<style lang="sass" scoped>
.my-custom-toggle
  border: 1px solid #027be3
</style>
