<template>
  <div class="q-px-sm">
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
          flat
          round
          icon="add"
          color="blue"
          @click="addQuizActive=true"
        />
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
        slot="body-cell-action"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          flat
          icon="edit"
          @click="toEditQuiz(props.row.QuizId)"
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
      :quiz="quizId"
      @quizEditCompleted="editQuizDone"
    />
    <add-quiz
      :active="addQuizActive"
      @quizAddCompleted="addQuizDone"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapMutations } from 'vuex'
import { GET_QUIZ_QUERY } from '../../graphql/queries'

export default {
  name: 'StudentTable',
  components: {
    'edit-quiz': require('components/quiz/EditQuiz.vue').default,
    'add-quiz': require('components/quiz/AddQuiz.vue').default
  },
  data () {
    return {
      loading: true,
      addQuizActive: false,
      editQuizActive: false,
      quizId: '',
      skipQueryGetQuiz: true,
      visibleColumns: ['QuizId', 'Grade', 'Desc', 'Operator', 'Status', 'Approver', 'action'],
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
          name: 'action',
          align: 'center',
          label: 'Edit Quiz',
          field: 'action',
          sortable: false
        }
      ]
    }
  },
  computed: {
    // ...mapGetters('currentUser', ['currentUser']),
    ...mapGetters('quiz', ['getQuizList']),
    tableData: function () {
      console.log('quiz list = ', this.getQuizList)
      return this.getQuizList
    }
  },
  mounted () {
    this.fetchQuiz()
  },
  methods: {
    ...mapMutations('quiz', ['addQuiz', 'setQuizList']),
    toEditQuiz (opt) {
      this.quizId = opt
      this.editQuizActive = true
    },
    editQuizDone () {
      this.editQuizActive = false
    },
    addQuizDone () {
      this.addQuizActive = false
    },
    fetchQuiz () {
      this.skipQueryGetQuiz = false
      this.$apollo.queries.GetQuiz.refetch()
    },
    updateQuiz (newList) {
      console.log('in update quiz length - ', newList.length)
      if (newList.length > 0) {
        this.loading = false
        this.setQuizList([])
        for (let i = 0; i < newList.length; i++) {
          this.addQuiz(newList[i])
        }
      } else {
        console.log('NO enough quiz')
      }
    }
  },
  apollo: {
    GetQuiz: {
      query: GET_QUIZ_QUERY,
      variables () {
        return {}
      },
      error (error) {
        console.error('We\'ve got an error!', error)
      },
      skip () {
        return this.skipQueryGetQuiz
      },
      result (data, key) {
        this.skipQueryGetQuiz = true
        if (data.data.GetQuiz) {
          this.updateQuiz(data.data.GetQuiz)
        } else {
          console.log('Get quiz failed')
        }
      }
    }
  }
}
</script>

<style>
</style>
