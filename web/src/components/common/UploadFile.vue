<!--
 * @Author: yaoyaoyu
 * @Date: 2021-05-27 21:49:56
-->
<template>
<div>
    <el-upload
      class="upload-demo"
      ref="upload"
      drag
      :action="uploadActionUrl"
      :on-change="handleChange"
      :data="uploadData"
      :file-list="fileList"
      :on-success="uploadSuccess"
      :auto-upload="false">
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过5m</div>
    </el-upload>
    <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
</div>
</template>

<script>
export default {
  name:'UploadFile',
  data(){
    return{
      uploadActionUrl:'http://localhost:8080/file/upload',   // 上传路径
      fileList: [ ],                                         // 文件列表
      // 上传附加参数
      uploadData: {
          type: 1,
      }
    }
  },
  methods:{
      // 文件列表变更时相应
      handleChange(file, fileList) {
        console.log('file.type :>> ', fileList[0].raw.type);
        // this.uploadData.type = fileList[0].raw.type;
      },
      // 上传成功的回调
      uploadSuccess(response, file, fileList){
        console.log('response :>> ', response);
      },
      // 上传
      submitUpload(){
         this.$refs.upload.submit();
      }

  }
};
</script>

<style>
</style>