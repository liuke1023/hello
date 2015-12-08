<!DOCTYPE html>

<html>
<head>
   <meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
    <title>测试</title>
    <link rel="Stylesheet"  type="text/css" href="../static/css/StyleSheet.css" />
	<script type="text/javascript" src="../static/js/jquery-1.9.1.min.js" ></script>
	<script type="text/javascript">
	$(document).ready(function(){
		$("#btnsubmit").click(function(){
		 if($("#uname").val()=="")
		{
			alert("请填写登陆名");
			return;
		}
		if($("#upw").val()=="")
		{
			alert("密码不能为空");
			return;
		}
		if($("#myyzm").val()=="")
		{
			alert("请填写验证码");
		}
		
		$.ajax({
			type:"post",
			url:"/yzm/judgeyzm",
			data:{
			param:$("#myyzm").val()
			},
			error:function(t1,t2,t3){
				alert("提交数据时发生错误，错误:"+t1.status)
			},
			success:function(data)
			{
				var v=data;
				if(v=="y")
				{
					//alert("验证码正确");
					$.ajax({
						type:"post",
						url:"/user",
						data:{
						uname:$("#uname").val(),
						upw:$("#upw").val()
						},
						error:function(t1,t2,t3){
							alert("提交时发生错误，错误:"+t1.status)
						},
						success:function(data)
						{ 
							var v=data;
							if(v=="yes")
							{
							location.href="/logmain";}
							else
							{
								alert("用户名或密码错误");
								$("#captchacode").attr("src","/yzm?"+Math.random());
							}
						}
					
					});
				}
				else
				{
					alert(v);
					$("#captchacode").attr("src","/yzm?"+Math.random());
				}
				
			}
			
		});
		
		});
		
		$("#btncancel").click(function(){
			$("#uname").val("");
			$("#upw").val("");
			$("#myyzm").val("");
			$("#captchacode").attr("src","/yzm?"+Math.random());
		});
	});
	</script>
</head>
<body>
<div class="loginall">
    <table class="logintable">
    <tr><th colspan="3">用户登陆</th></tr>
    <tr><td>登陆名：</td><td colspan="2"><input type="text" id="uname" class="inputtext150" /></td></tr>
    <tr><td>密码：</td><td colspan="2"><input type="password" id="upw" class="inputtext150" /></td></tr>
    <tr><td>验证码</td><td class="td100"><input type="text" id="myyzm" class="inputtext40" /></td><td class="imgtest"><img id="captchacode" src="/yzm"  onclick="this.src='/yzm?'+Math.random()" /></td></tr>
    <tr><td colspan="3"><input type="button" id="btnsubmit" value="提交" />
            <input type="button" id="btncancel" value="清空" />
    </td></tr>
    </table>
</div>
</body>
</html>
