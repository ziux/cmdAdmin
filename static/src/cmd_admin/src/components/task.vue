<template>
    <div>
        <task-line :data="useData" v-if="isShow" @closeIt="closeIt"></task-line>
        <Output :task="outputTask" v-if="isShowOutput" @closeOut="closeOut"></Output>
        <el-container>
            <el-header></el-header>
            <el-main style="text-align: center">
                <div align="right"></div>
                <div align="right">
                    <el-button @click="init">刷新</el-button>
                    <el-button @click="isShow=true">添加</el-button>
                </div>

                <el-table
                        :data="tableData"
                        style="width: 80%;text-align: center">
                    <el-table-column
                            prop="Name"
                            label="任务名称"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="UpdatedAt"
                            label="更新时间"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="PID"
                            label="PID">
                    </el-table-column>
                    <el-table-column

                            label="状态">
                        <template slot-scope="scope">
                            <el-tag v-show="scope.row.State===0">未开始</el-tag>
                            <el-tag v-show="scope.row.State===1">进行中</el-tag>
                            <el-tag v-show="scope.row.State===2">已完成</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column
                            label="重启策略">
                        <template slot-scope="scope">
                            <el-tag v-show="scope.row.KeepType===1">定时重启{{scope.row.Timing}}</el-tag>
                            <el-tag v-show="scope.row.KeepType===2">保持运行</el-tag>

                        </template>
                    </el-table-column>

                    <el-table-column label="操作">
                        <template slot-scope="scope">
                            <el-button type="text" size="small" @click="start(scope.row)"
                                       v-if="scope.row.State===0||scope.row.State===2">start
                            </el-button>
                            <el-button type="text" size="small" @click="kill(scope.row)" v-if="scope.row.State===1">
                                kill
                            </el-button>
                            <el-button type="text" size="small" @click="Openoutput(scope.row)">查看</el-button>
                            <el-button type="text" size="small" @click="updateTask(scope.row)">编辑</el-button>
                            <el-button type="text" size="small" @click="deleteTask(scope.row)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-main>

        </el-container>
    </div>

</template>

<script>
    import TaskLine from "./TaskLine";
    import axios from 'axios'
    import Output from "./Output";

    export default {
        name: "task",
        components: {Output, TaskLine},
        data() {
            return {
                tableData: [{
                    a: 1, b: 2
                }],
                createData: {
                    id:'',
                    name: "",
                    dir: '',
                    command: '',
                    keepType:'0',
                    timing:0,
                    remark: ''
                },
                useData: {
                    id:'',
                    name: "",
                    dir: '',
                    command: '',
                    keepType:'0',
                    timing:0,
                    remark: ''
                },
                isShow: false,
                isShowOutput: false,
                outputTask: {}
            }
        },
        created() {
            this.init()
        },
        methods: {
            closeIt: function (refresh) {
                this.isShow = false;
                if (refresh) {
                    this.messageSuccess()
                    this.init();
                }

            },
            updateTask:function(task){

                this.useData.id=task.ID
                this.useData.name=task.Name
                this.useData.dir=task.Dir
                this.useData.command=task.Command
                this.useData.keepType=task.KeepType.toString()
                this.useData.timing=task.Timing
                this.useData.remark=task.Remark
                console.log(this.useData)
                this.isShow = true

            },
            closeOut: function () {
                this.isShowOutput = false;

            },
            Openoutput: function (task) {
                this.isShowOutput = true
                this.outputTask = task
            },
            start: function (taskData) {
                let vue = this;
                vue.confirmAction(function () {
                axios.get("http://127.0.0.1:3000/api/task/start?id=" + taskData.ID,).then(function (res) {
                    if (res.data.code == 0) {
                        vue.messageSuccess()
                        vue.init()
                    }else {
                        vue.messageError(res.data.msg)
                    }
                })
                })
            },
            kill: function (taskData) {
                let vue = this;
                vue.confirmAction(function () {
                    axios.get("http://127.0.0.1:3000/api/task/kill?id=" + taskData.ID,).then(function (res) {
                        if (res.data.code == 0) {
                            vue.messageSuccess()
                            vue.init()
                        }else {
                            vue.messageError(res.data.msg)
                        }
                    })
                })
            },
            deleteTask: function (taskData) {
                let vue = this;
                vue.confirmAction(function () {
                axios.delete("http://127.0.0.1:3000/api/task?id=" + taskData.ID,).then(function (res) {
                    if (res.data.code == 0) {
                        vue.messageSuccess()
                        vue.init()
                    }else {
                        vue.messageError(res.data.msg)
                    }
                })
                })
            },
            init: function () {
                let vue = this
                axios.get("http://127.0.0.1:3000/api/task",).then(function (res) {
                    vue.tableData = res.data.data
                })
            },
            messageSuccess: function (msg) {
                if (!msg) {
                    msg = '操作成功'
                }
                this.$message({
                    message: msg,
                    type: 'success'
                });
            },
            messageError: function (msg) {
                if (!msg) {
                    msg = '操作失败'
                }
                this.$message({
                    message: msg,
                    type: 'warning'
                });
            },
            confirmAction: function (callback, content) {
                if (!content) {
                    content = '确认操作?'
                }
                this.$alert(content, '提示', {
                    confirmButtonText: '确定',
                    callback: action => {
                        if (action === 'confirm'){
                            callback()
                        }
                    }
                });
            }
        }

    }
</script>

<style scoped>

</style>