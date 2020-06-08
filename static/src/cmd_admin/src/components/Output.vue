<template>
    <div>
        <el-dialog :title="task.Name" :visible.sync="isShow" :before-close="bfClose" align="left">
            <p>Dir:{{task.Dir}}</p>
            <p>Command:{{task.Command}}</p>
            <div style="height: 400px;overflow:auto">
                <p v-for="(data,index) in dataList" v-bind:key="index">{{data}}</p>
            </div>
            <div slot="footer" class="dialog-footer">
                <el-button @click="deleteOutput" >清空</el-button>
                <el-button type="primary" @click="init">刷新</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: "Output",
        props: ['task',],
        data() {
            return {
                dataList: [],
                isShow: true
            }
        },
        methods: {
            bfClose: function (done) {
                this.$emit("closeOut")
                done()
            },
            deleteOutput:function(){
                let vue = this
                vue.confirmAction(function () {
                axios.get('http://127.0.0.1:3000/api/task//output/delete?id=' + vue.task.ID,).then(function (res) {
                    if (res.data.code === 0) {
                       vue.messageSuccess()
                        vue.init()
                    }else {
                        vue.messageError(res.data.msg)
                    }
                })
                })
            },
            init: function () {
                // let data = this.dataList
                let vue = this
                axios.get('http://127.0.0.1:3000/api/task/output?id=' + this.task.ID,).then(function (res) {
                    if (res.data.code === 0) {
                        let dataList = []
                        res.data.data.map(function (item) {
                            let s = ""
                            if (item.Type == 2) {
                                s += "info "
                            } else {
                                s += "error "
                            }
                            s += item.UpdatedAt
                            s += ' :     '
                            s += item.Text
                            dataList.push(s)
                        })
                        vue.dataList = dataList
                    }
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
        },

        created() {
            this.init()
        }
    }
</script>

<style scoped>

</style>