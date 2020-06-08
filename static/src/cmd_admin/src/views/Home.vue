<template>
  <div class="home" >
    <el-row>
      <el-col :span="4"><div > <p></p></div></el-col>
      <el-col :span="16"><div >
        <task />
      </div>

      </el-col>
      <el-col :span="4"><div ><p></p></div></el-col>
    </el-row>

  </div>
</template>

<script>
// @ is an alias to /src
import task from "../components/task";
export default {
  name: 'Home',
  components: {
   task
  }
}
</script>
