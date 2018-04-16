var applyurl="student/apply.html";
var waitpassyurl="student/pass.html";
var updateinfourl="student/profileForm.html"
$(function(){
	$("#apply").load(applyurl,function(){
})
	$("#pass").load(waitpassyurl,function(){
})
	$("#updateinfo").load(updateinfourl,function(){
		$("#register").html("修改");
})

	$(".firstul").on("click",function(){
		var index=$(this).index();
		index=index/2;
		$(".firstul").parent().find("ul").hide();
		$(this).parent().find("ul").eq(index).show();
	});
	var li=$(".dropdown").find("li");
	li.on("click",function(){
		var index=$(this).index();
		li.removeClass("active");
		$(this).addClass("active");
		if((index%2)==0){
		$("#apply").show();
		$("#pass").hide();
		$("#updateinfo").hide();
	}
	if((index%2)!=0){
		$("#apply").hide();
		$("#pass").show();
		$("#updateinfo").hide();
	}
	})
	$(".updateinfo").on("click",function(){
		$("#apply").hide();
		$("#pass").hide();
		$("#updateinfo").show();

	})
	
});


