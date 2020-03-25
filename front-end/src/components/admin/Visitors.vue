<template>
  <div>
    <search-bar />
    <q-table
      title="访客记录"
      :data="tableData"
      :columns="columns"
      :visible-columns="visibleColumns"
      separator="cell"
      wrap-cells
      row-key="date"
      :table-header-style="{ backgroundColor: 'lightblue' }"
      fixed-center
      bordered
      dense
      :pagination.sync="pagination"
      :loading="loading"
    >
      <template v-slot:top="props">
        <div class="col-2 q-table__title">
          <q-chip
            size="xs"
            text-color="blue"
            style="font: 50% Cursive;"
          >
            访客记录
          </q-chip>
        </div>
        <q-space />
        <q-select
          v-model="visibleColumns"
          multiple
          dense
          options-dense
          display-value="选择显示列"
          emit-value
          map-options
          :options="columns"
          option-value="name"
          style="min-width: 150px"
        />
        <q-btn
          class="q-ml-md"
          flat
          round
          dense
          :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
          @click="props.toggleFullscreen"
        />
      </template>
      <!-- slot name syntax: body-cell-<column_name> -->
      <q-td
        slot="body-cell-action"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          v-if="isDeletable"
          flat
          icon="delete"
          @click="toDeleteLog(props.row.datetime, props.row.deviceid)"
        />
      </q-td>
    </q-table>
    <alert-msg
      :active="alert"
      :msg="alertMsg"
      @cancelAlert="alertClose"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapMutations } from 'vuex'

export default {
  name: 'Visitors',
  components: {
    'alert-msg': require('components/common/AlertMsg.vue').default,
    'search-bar': require('components/hywt/SearchBar.vue').default
  },
  props: {
    studentId: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      isDeletable: true,
      alert: false,
      alertMsg: '',
      pagination: {
        page: 0,
        rowsPerPage: 10
      },
      visibleColumns: ['datetime', 'name', 'gender', 'ethnic', 'id', 'dob', 'block', 'enter', 'deviceid'],
      columns: [
        {
          name: 'datetime',
          align: 'center',
          label: '访问时间',
          field: 'datetime',
          sortable: true
        },
        {
          name: 'name',
          align: 'center',
          label: '姓名',
          field: 'name',
          sortable: true
        },
        {
          name: 'gender',
          align: 'center',
          label: '性别',
          field: 'gender',
          sortable: true
        },
        {
          name: 'ethnic',
          align: 'center',
          label: '民族',
          field: 'ethnic',
          sortable: true
        },
        {
          name: 'id',
          align: 'center',
          label: '身份证号',
          field: 'id',
          sortable: true
        },
        {
          name: 'dob',
          align: 'center',
          label: '生日',
          field: 'dob',
          sortable: true
        },
        {
          name: 'block',
          align: 'center',
          label: '出入社区',
          field: 'block',
          sortable: true
        },
        {
          name: 'enter',
          align: 'center',
          label: '通行类型',
          field: 'enter',
          sortable: true
        },
        {
          name: 'deviceid',
          align: 'center',
          label: '设备号',
          field: 'deviceid',
          sortable: true
        },
        {
          name: 'action',
          align: 'center',
          label: 'action',
          field: 'action',
          sortable: false
        }
      ],
      loading: false
    }
  },
  computed: {
    ...mapGetters('visitor', ['visitor']),
    tableData: function () {
      return (this.visitor.visitorLog)
    }
  },
  mounted () {
    console.log('task log table mounted - ')
    setInterval(this.testFun, 8000)
  },
  destroyed () {
    console.log('task log table destroied')
  },
  methods: {
    ...mapMutations('visitor', ['removeVisitor']),
    toDeleteLog (dt, di) {
      for (let i = 0; i < this.tableData.length; i++) {
        if (this.tableData[i].datetime === dt && this.tableData[i].deviceid === di) {
          this.removeVisitor(this.tableData[i])
        }
      }
    },
    testFun () {
      console.log('current visitorLog length - ', this.visitor.visitorLog.length)
    },
    alertClose () {
      this.alert = false
      this.alertMsg = ''
    }
  }
}
</script>

<style>
.label {
  white-space: pre-wrap;
}
</style>
