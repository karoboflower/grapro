$(function(){
	    ReloadCounselor();   
	    write_level(); 
	    college();       
			var	options = {
				success  : ResponseCallback
			};

			$("#profileForm").ajaxForm(options);
			//判断表格中的内容是否合格
			$("#register").on("click",function(){
				//判断身份证号
				var identityvalue=/(\d{18} |^\d{17}(\d|X|x)|^\d{15})$/;
				var regphone=/^(13[0-9]\d{8}|15[0-35-9]\d{8}|18[0-9]\d{8}|17[6-8]\d{8}|14[57]\d{8})$/;//手机号码
				var idnumber = $("#idnumber").val();
				var phone = $("#phone").val();
				var name = $("#name").val();
				if(name == null || name == ""){
					layer.tips("请输入姓名",$("#name"), {tips:3});
                	return false;
				}
				 if (idnumber == null || idnumber == "") {
                	layer.tips("请输入身份证号",$("#idnumber"), {tips:3});
                	return false;
                }else if(!identityvalue.test(idnumber)){
                	layer.tips("请输入正确的身份证号",$("#idnumber"), {tips:3});
                	return false;
                }
                if(phone == null || phone == ""){
                	//判断电话号码
                	layer.tips("请输入电话号码",$("#phone"), {tips:3});
                	return false;
                }else if(!regphone.test(phone)){
                	layer.tips("请输入正确的电话号码",$("#phone"), {tips:3});
                	return false;
                }
			
	});

		});

			
			function ResponseCallback(response,status,xhr,$form) {
			if (response.status==1) {
				layer.alert(response.msg, {icon:2});
			}
			if (response.status==0) {
				layer.alert("提交信息成功", {icon:6});
				window.parent.location.href="/auth/3/"+$("#id").val()+"/profile";
			}
		}
			function ReloadCounselor() {
            $.get("/auth/3/"+$("#id").val()+"/counselor", {"grade":$("#grade").val(), "college":$("#college").val()}, function(response){
                if (response.status == 1) {
                    layer.alert(response.msg, {icon:2});
                    return;
                }
                if (response.status==0) {
                    console.log(response.counselors);
                }
            })
        }
        //自动填写最新年纪
        function write_level(){
	var levels=$("#counselorid").find("option");
	var myDate = new Date();
	//填写辅导员年级
	for(var i=0;i<levels.length;i++){
		$("#counselorid").find("option").eq(i).text(myDate.getFullYear()-i+"级");
		$("#counselorid").find("option").eq(i).attr("value",myDate.getFullYear()-i);
	}
	var classes=$("#class").find("option");
	for(var i=0;i<levels.length;i++){
		$("#class").find("option").eq(i).text((i+1)+"班");
		$("#class").find("option").eq(i).attr("value",(i+1));
	}
	var classes=$("#grade").find("option");
	for(var i=0;i<levels.length;i++){
		$("#grade").find("option").eq(i).text(myDate.getFullYear()-i+"级");
		$("#grade").find("option").eq(i).attr("value",myDate.getFullYear()-i);
	}
	
}
function college(){
	var getcollegeurl="/resources/assets/student/json/college.json"
	$.ajax({
		url:getcollegeurl,
		type:"get",
		xhrFields: {
			withCredentials: true
			//可执行跨域名请求
		},
		success:function(data){
			write_college(data);
		}
	})
}
function write_college(data){
	var college=$("#college");
	if(data){
	for(var j=0;j<data.college.length;j++){
		var item=data.college[j].name;
		college.append("<option value='"+item+"'>"+item+"</option>");
	}
	$("#college").change(function(){
				var college=$(this).val();
				var major='';
				for(var j=0;j<data.college.length;j++){
					var item=data.college[j];
					if(item.name==college){
						for(var i=0;i<item.item.length;i++){
							var majoritem=item.item[i];
							var content="<option value='"+majoritem+"'>"+majoritem+"</option>";
							major+=content;	
						}
						
					}
	}
	$("#profession").empty().append(major);

			})
	}
}
