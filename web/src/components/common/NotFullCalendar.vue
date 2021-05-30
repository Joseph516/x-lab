<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-28 15:41:44
-->
<template>  
  <div>
    <FullCalendar ref="fullCalendar" :options="calendarOptions"  :events="calendarOptions.events"  locale="zh-cn" />
    <div id="mydraggable"></div>
    
          <!-- 表单控制 -->
        <el-dialog :title="optTitle" :visible.sync="dialogFormVisible">
          <el-form :model="form">
            <el-form-item label="事件名称" label-width="80px">
              <el-input v-model="form.title" auto-complete="off" placeholder="请输入事件名称"></el-input>
            </el-form-item>
            <el-form-item label="开始时间" label-width="80px">
                <el-date-picker
                  v-model="form.start"
                  type="datetime"
                  placeholder="选择日期时间">
                </el-date-picker>
            </el-form-item>
            <el-form-item label="结束时间" label-width="80px">
                <el-date-picker
                  v-model="form.end"
                  type="datetime"
                  placeholder="选择日期时间">
                </el-date-picker>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="warning" @click="delEvent" v-if="form.id" style="float: left;">删 除</el-button>
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="saveEvent">确 定</el-button>
          </div>
        </el-dialog>

  </div>
</template>

<script>
import FullCalendar from '@fullcalendar/vue'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin , { Draggable }from '@fullcalendar/interaction'

export default {
  components: {
    FullCalendar // make the <FullCalendar> tag available
  },
  mounted(){

  },
  data() {
    return {
      calendarOptions: {
        plugins: [ dayGridPlugin,timeGridPlugin, interactionPlugin ],
        initialView: 'dayGridMonth',
        themeSystem:'standard',
        firstDay:1 ,
        timeZone:"UTC" ,
        weekNumberCalculation:"ISO" ,
        editable:true,
        displayEventEnd:true,
        selectable: true,                   // 可选
        droppable: true,                    // 拖入
            headerToolbar: {
                left: 'prev,next today',
                center: 'title',
                right: 'dayGridMonth,timeGridWeek,timeGridDay'
            },
        locale:"zh-cn" ,                    // 中文
        dateClick: this.handleDateClick,    // 日期点击
        eventClick: this.handleEventClick,  // 事件点击
        select:this.handleDateSelect,       // 日期区域选择
        eventDrop:this.dropEvent,           // 事件拖拽
        events: [
          { title: 'event 1', date: '2021-05-20' },
          { title: 'event 2', date: '2021-05-30' },
          { title: 'enve', start:'2021-05-20',end:'2021-05-30'}
        ]
      },

      //表单控制
      dialogFormVisible: false,
      form: {
                title: null,
                start: null,
                end: null,
                endStr:'',
                startStr:'',
            },
      optTitle: '添加事件',
    }
  },
  methods:{

    getAPI(){
      let calendarApi = this.$refs.fullCalendar.getApi()
      console.log('calendarApi :>> ', calendarApi);
      calendarApi.next()
    },

    // 点击日期回调
    handleDateClick: function(arg) {
      console.log(arg)
    },

    // 选择一段时间回调
    handleDateSelect:function(selectionInfo){
      console.log('selectionInfo :>> ', selectionInfo);
      console.log('日期 :>> ', selectionInfo.startStr + "  "+selectionInfo.endStr);
      this.dialogFormVisible = true;
      // 表单参数更新
      this.form.start = selectionInfo.start;
      this.form.end = selectionInfo.end;
      this.form.startStr = selectionInfo.startStr;
      this.form.endStr = selectionInfo.endStr;
    },
    // 删除事件
    delEvent(){

    },
    // 保存事件
    saveEvent(){
      console.log('saveEvent :>> ', {title: this.form.title,  start: this.form.start,  end: this.form.end});
      this.calendarOptions.events.push({title: this.form.title,  start: this.form.start,  end: this.form.end})
      this.dialogFormVisible = false;

    },
    // 更新事件
    updateEvent(){

    },
    // 从其他地方拖拽创建新事件
    calendarEventDrop(info){
      console.log('drop info :>> ', info);
      this.dropEvent(info);
    },
    // 拖拽事件回调，可更新
    dropEvent(info) {
      console.log('dropevent  :>> ', info );
      // 当前拖拽的事件的form更新
      this.form = {
          id: info.event.id,
          title: info.event.title,
          start: info.event.start,
          end: info.event.end,
          startStr:info.event.startStr,
          endStr : info.event.endStr,
      };
      this.updateEvent();
    },

    // 事件点击回调
    handleEventClick(info){
      console.log('eventclick info :>> ', info);
    }


  }
}
</script>

<style>

</style>