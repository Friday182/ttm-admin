<template>
  <div>
    <div class="row">
      <div class="col-2 q-mr-sm">
        <el-select
          v-model="province"
          class="myClass"
          filterable
          size="mini"
          placeholder="所属省份"
        >
          <el-option
            v-for="item in provinceOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </div>
      <div class="col-2 q-mr-sm">
        <el-select
          v-model="city"
          filterable
          size="mini"
          placeholder="选择城市"
        >
          <el-option
            v-for="item in cityOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
          <template v-slot:prefix>
            <q-icon name="place" />
          </template>
        </el-select>
      </div>
      <div class="col-2 q-mr-sm">
        <el-input
          v-model="inId"
          class="myClass"
          size="mini"
          placeholder="身份证号码"
          clearable
        />
      </div>
      <div class="col-2 q-mr-sm">
        <el-input
          v-model="block"
          size="mini"
          placeholder="社区名称"
          clearable
        />
      </div>
    </div>
    <div class="row q-my-sm">
      <div class="col-4 q-mr-sm">
        <el-date-picker
          v-model="period"
          size="mini"
          type="daterange"
          align="right"
          unlink-panels
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          :picker-options="pickerOptions"
        />
      </div>
      <q-space />
      <div class="col-2">
        <q-btn
          label="查 询"
          outlined
          dense
          color="primary"
          style="font: 100% Cursive"
          @click="search()"
        />
        <q-btn
          class="q-ml-sm"
          label="重 置"
          outlined
          dense
          color="primary"
          style="font: 100% Cursive"
          @click="clearForm()"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { sendWSPush } from '../../api/websocket'

export default {
  name: 'Error404',
  data () {
    return {
      inId: '',
      city: '',
      province: '',
      block: '',
      deviceid: '',
      provinceOptions: [
        { value: '选项1', label: '新疆' },
        { value: '选项2', label: '甘肃' },
        { value: '选项3', label: '陕西' },
        { value: '选项4', label: '四川' },
        { value: '选项5', label: '河南' }
      ],
      cityOptions: [
        { value: '选项1', label: '乌鲁木齐' },
        { value: '选项2', label: '石河子' },
        { value: '选项3', label: '伊犁' },
        { value: '选项4', label: '库尔勒' },
        { value: '选项5', label: '克拉玛依' }
      ],
      pickerOptions: {
        shortcuts: [{
          text: '最近一周',
          onClick (picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近一个月',
          onClick (picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近三个月',
          onClick (picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
            picker.$emit('pick', [start, end])
          }
        }]
      },
      period: ''
    }
  },
  methods: {
    search () {
      console.log('start search ...', this.period)
      if (sessionStorage.getItem('ws-status') === 'OK') {
        let msg = {
          'sq': 0,
          'type': 'SEARCH',
          'data': {
            'id': this.inId,
            'deviceid': '',
            'block': this.block,
            'period': this.period
          },
          'from': 'ws-id',
          'token': localStorage.getItem('access_token')
        }
        sendWSPush(msg)
      } else {
        // send data through REST API
      }
    },
    clearForm () {
      this.inId = ''
      this.province = ''
      this.city = ''
      this.block = ''
      this.period = ''
    }
  }
}
</script>

<style>
.myClass input.el-input__inner {
    border-radius:15px!important;
}
.dialog-input-text >>> .el-input__inner {
  -web-kit-appearance: none;
  -moz-appearance: none;
  font-size: 1.4em;
  height: 2.9em;
  border-radius: 4px;
  border: 1px solid #b6d8f1;
  color: #6a6f77;
  outline: 0;
}
</style>
<style scope>
</style>
