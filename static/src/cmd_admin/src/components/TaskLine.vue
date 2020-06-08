<template>
    <div>
        <el-dialog title="添加任务" :visible.sync="isShow" :before-close="bfClose">
            <el-form :model="data">
                <el-form-item label="名称" :label-width="formLabelWidth">
                    <el-input v-model="data.name" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="路径" :label-width="formLabelWidth">
                    <el-input v-model="data.dir" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="命令" :label-width="formLabelWidth">
                    <el-input v-model="data.command" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="重启策略" :label-width="formLabelWidth">
                    <template>
                        <el-radio v-model="data.keepType" label="0" value="0">没有</el-radio>
                        <el-radio v-model="data.keepType" label="1" value="1">定时重启</el-radio>
                        <el-radio v-model="data.keepType" label="2" value="2">保持运行</el-radio>
                    </template>
<!--                    <el-input v-model="data.keepType" autocomplete="off" type="number"></el-input>-->
                </el-form-item>
                <el-form-item label="定时间隔" :label-width="formLabelWidth">
                    <el-input v-model="data.timing" autocomplete="off" type="number"></el-input>
                </el-form-item>
                <el-form-item label="备注" :label-width="formLabelWidth">
                    <el-input v-model="data.remark" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="$emit('closeIt',false)">取 消</el-button>
                <el-button type="primary" @click="primaryButton">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: "TaskLine",
        data() {
            return {
                formLabelWidth: "200px",
                isShow: true
            }
        },
        props: ['data',],
        methods: {
            primaryButton:function(){
                let data = this.data
                if (data.id){
                this.updateTask()
                }else {
                    this.createTask()
                }
            },
            createTask: function () {
                let data = this.data
                let vue = this
                axios.post('http://127.0.0.1:3000/api/task', {
                    name: data.name,
                    dir: data.dir,
                    command: data.command,
                    remark: data.remark,
                    keepType:parseInt(data.keepType),
                    timing:parseInt(data.timing)
                }).then(function (res) {
                    if (res.data.code===0){
                        console.log(res)
                        vue.messageSuccess()
                        vue.$emit("closeIt",true)
                    }else {
                        vue.messageError(res.data.msg)
                    }

                })
            },
            updateTask: function () {
                let data = this.data
                let vue = this
                axios.put('http://127.0.0.1:3000/api/task?id='+data.id, {
                    dir: data.dir,
                    command: data.command,
                    remark: data.remark,
                    keepType:parseInt(data.keepType),
                    timing:parseInt(data.timing)
                }).then(function (res) {
                    if (res.data.code===0){
                        console.log(res)
                        vue.messageSuccess()
                        vue.$emit("closeIt",true)
                    }else {
                        vue.messageError(res.data.msg)
                    }

                })
            },
            bfClose:function (done) {
                this.$emit("closeIt",false)
                done()


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
        },
    }

</script>

<style scoped>

</style>