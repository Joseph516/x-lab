
export function checkAccount(userAccount){
  if (userAccount == null || userAccount == "") {
    return {data: false, msg: '请输入账号'};
  }
 
  if (userAccount.length > 20 || userAccount.length < 4) {
    return {data: false, msg: '登录账号长度应为6-20'};
  }
  // 特殊字符判断和汉字
  // /[^\u4e00-\u9fa5]/匹配中文字符，[^\x00-\xff]/匹配双字节字符
  let re = /[^\x00-\xff]/;
  if (/[#\$%\^&\*【】@!！￥?|‘；：”“'。，、？<>+=:]+/g.test(userAccount) || re.test(userAccount)) {
    return {data: false, msg: '账号不允许特殊字符或者中文字符'};
  }
 
  let reg = /^[0-9a-zA-Z_]{6,20}$/g
  if (reg.test(userAccount)) {
    return {data: true, msg: '账号可用'};
  } else {
    return {data: false, msg: '请输入6~20位字符'};
  }
}

export function checkNull(username,password){
  let msg = '';
  if(username === ''|| username=== null){
    msg += '用户名不能为空  ';
  }

  if(password === ''|| password === null){
    msg += '密码不能为空  ';
  }

  return msg;
}

