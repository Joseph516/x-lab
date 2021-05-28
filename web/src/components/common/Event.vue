<template>
    <div id="page">
        <div class="main">
            <FullCalendar defaultView="dayGridMonth" 
            locale="zh-cn" 
            timeZone="UTC" 
            firstDay="1" 
            weekNumberCalculation="ISO" 
            editable="true"
            droppable="true"
            displayEventEnd="true"
            :eventTimeFormat="evnetTime"
            :header="header"
            :plugins="calendarPlugins"
            :events="calendarEvents"
            @dateClick="handleDateClick" 
            @eventClick="handleEventClick"
            @eventDrop="calendarEventDrop"
            @datesRender="handleDatesRender"
             />
        </div>

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
import interactionPlugin, { Draggable } from '@fullcalendar/interaction';


export default {
    components: {
        FullCalendar
    },
    data() {
        return {
            calendarPlugins: [ 
                dayGridPlugin,
                timeGridPlugin,
                interactionPlugin
            ],
            header: {
                left: 'prev,next today',
                center: 'title',
                right: 'dayGridMonth,timeGridWeek,timeGridDay'
            },
            evnetTime: {
                hour: 'numeric',
                minute: '2-digit',
                hour12: false
            },

            calendarEvents: [
                {title:'111',start:'2021-5-19',end:'2021-5-30'}
            ],
            calendarEventDrop: info => {
                this.dropEvent(info);
            },
            handleDatesRender: arg => {
                this.getEventsList(arg.view)
            },

            dialogFormVisible: false,
            form: {
                title: null,
                start: null,
                end: null
            },
            optTitle: '添加事件',
        }
    },
    mounted() {
        
    },
    created() {
        //
    },
    methods: {
        getEventsList(info) {
            let params = {
                start: info.activeStart,
                end: info.activeEnd
            };
            // this.$get('events.php', params)
            // .then((res) => {
            //     this.calendarEvents = res;
            // });
        },
        handleDateClick(arg) {
            //console.log(arg.date);
            this.dialogFormVisible = true;
            this.optTitle = '新增事件';
            this.form.title = '',
            this.form.id = '',
            this.form.start = arg.date;
            this.form.end = arg.date;
            //console.log(this.calendarEvents);
        },
        handleEventClick(info) {
            info.el.style.borderColor = 'red';
            this.dialogFormVisible = true;
            this.optTitle = '修改事件';
            this.form = {
                id: info.event.id,
                title: info.event.title,
                start: info.event.start,
                end: info.event.end,
            };
        },
        saveEvent() {
            console.log('this.form :>> ', this.form);
          
              
          
        },
        //删除事件
        delEvent() {
           
            console.log('this.form :>> ', this.form);
        },
        //拖动事件
        dropEvent(info) {
            this.form = {
                id: info.event.id,
                title: info.event.title,
                start: info.event.start,
                end: info.event.end
            };
            this.saveEvent();
        }
    }
}
</script>
<style>

</style>