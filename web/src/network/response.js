/*
 * @Author: yaoyaoyu
 * @Date: 2021-05-30 17:04:58
 */

export function codeCheck(res){
  if(res.code === 200){
    return true;
  }else{
    return false;
  }
}