
$(function(){ 
  var informurl="studentOffice/inform.html";
    //三类奖学金下拉框的显示
    $(".firstul").on("click",function(){
      var index=$(this).index();
      index=index/2;
      $(".firstul").parent().find("ul").hide();
      $(this).parent().find("ul").eq(index).show();
    });
    //发布通知
    $("#inform").load(informurl,function(){

    })
    //初始默认展开
    showtable(0,"resources/assets/studentOffice/json/stateGrants.json",stategrantscolums);
    //系统相关
    //国家助学金
    $("#collapseOne li").on("click",function(){
      $("li").removeClass("active");
      $(this).addClass("active");
      var code=$(this).attr("code");
      var stateGrantsurl="resources/assets/studentOffice/json/stateGrants.json";
      showtable(code,stateGrantsurl,stategrantscolums);
    })
    //国家励志助学金
    $("#collapseTwo li").on("click",function(){
      $("li").removeClass("active");
      $(this).addClass("active");
      var code=$(this).attr("code");
      var stateGrantsurl="resources/assets/studentOffice/json/nis.json";
      showtable2(code,stateGrantsurl,niscolums);
    })
    //应善良助学金
    $("#collapseThree li").on("click",function(){
      $("li").removeClass("active");
      $(this).addClass("active");
      var code=$(this).attr("code");
      var stateGrantsurl="resources/assets/studentOffice/json/ks.json";
      showtable3(code,stateGrantsurl,kscolums);
    })
    //通知管理
    $("#collapseFour li").on("click",function(){
      $("li").removeClass("active");
      $(this).addClass("active");
      var code=$(this).attr("code");
      var stateGrantsurl="resources/assets/studentOffice/json/ks.json";
      showtable4(code,stateGrantsurl,kscolums);
    })
    //审核
    $("#table").on("click",".examine",function(){
      layer.open({
        type: 1,
        title: false,
        closeBtn: 0,
        area: '516px',
        shadeClose: true,
        content: $('#show-info-check')
      })
    })
    //删除
    $("#table").on("click",".deletedata",function(){
      layer.msg('删除成功');
    })       

});
    //励志奖学金头部
    var niscolums=[{
      field:'num',
      title:'',
      align:'center',
      sortable:true
    },{
      field:'person',
      title:'申请人',
      align:'center',
      sortable:true
    },{
      field:'college',
      title:'隶属学院',
      align:'center',
      sortable:true
    },{
      field:'applyname',
      title:'文件名',
      align:'center',
      sortable:true
    },{
      field:'materials',
      title:'相关证明材料',
      align:'center',
      sortable:true
    },{
      field:'time',
      title:'申请时间',
      align:'center',
      sortable:true
    },{
      field:'status',
      title:'审核状态',
      align:'center',
      sortable:true
    },{
      field:'description',
      title:'审核描述',
      align:'center',
      sortable:true
    },{
      field:'ways',
      title:'操作',
      align:'center',
      sortable:true
    }];
    //助学金头部
    var stategrantscolums=[{
      field:'num',
      title:'',
      align:'center',
      sortable:true
    },{
      field:'person',
      title:'申请人',
      align:'center',
      sortable:true
    },{
      field:'college',
      title:'隶属学院',
      align:'center',
      sortable:true
    },{
      field:'applyname',
      title:'文件名',
      align:'center',
      sortable:true
    },{
      field:'materials',
      title:'相关证明材料',
      align:'center',
      sortable:true
    },{
      field:'rank',
      title:'助学金等级',
      align:'center',
      sortable:true
    },{
      field:'time',
      title:'申请时间',
      align:'center',
      sortable:true
    },{
      field:'status',
      title:'审核状态',
      align:'center',
      sortable:true
    },{
      field:'description',
      title:'审核描述',
      align:'center',
      sortable:true
    },{
      field:'ways',
      title:'操作',
      align:'center',
      sortable:true
    }];
    //应善良奖学金
    var kscolums=[{
      field:'num',
      title:'',
      align:'center',
      sortable:true
    },{
      field:'person',
      title:'申请人',
      align:'center',
      sortable:true
    },{
      field:'college',
      title:'隶属学院',
      align:'center',
      sortable:true
    },{
      field:'applyname',
      title:'文件名',
      align:'center',
      sortable:true
    },{
      field:'materials',
      title:'本人相关证明材料',
      align:'center',
      sortable:true
    },{
      field:'counselor',
      title:'辅导员相关证明材料',
      align:'center',
      sortable:true
    },{
      field:'time',
      title:'申请时间',
      align:'center',
      sortable:true
    },{
      field:'status',
      title:'审核状态',
      align:'center',
      sortable:true
    },{
      field:'description',
      title:'审核描述',
      align:'center',
      sortable:true
    },{
      field:'ways',
      title:'操作',
      align:'center',
      sortable:true
    }];
    //通知管理
    var inform=[{
      field:'num',
      title:'',
      align:'center',
      sortable:true
    },{
      field:'person',
      title:'申请人',
      align:'center',
      sortable:true
    },{
      field:'college',
      title:'隶属学院',
      align:'center',
      sortable:true
    },{
      field:'applyname',
      title:'文件名',
      align:'center',
      sortable:true
    },{
      field:'materials',
      title:'本人相关证明材料',
      align:'center',
      sortable:true
    },{
      field:'counselor',
      title:'辅导员相关证明材料',
      align:'center',
      sortable:true
    },{
      field:'time',
      title:'申请时间',
      align:'center',
      sortable:true
    },{
      field:'status',
      title:'审核状态',
      align:'center',
      sortable:true
    },{
      field:'description',
      title:'审核描述',
      align:'center',
      sortable:true
    },{
      field:'ways',
      title:'操作',
      align:'center',
      sortable:true
    }];
    //表格参数
    var tableOption = {
      height:550,
      search: true,
      pagination: false,
      striped: true,
      sortStable: true,
      heckbox:false,
      toolbar: '#toolbar',
      pagination:true,
      showRefresh:true,
      formatNoMatches:function () {
        return "暂无数据！";
      },
      columns: null,
      data:[]
    };
    function loadTable($tableDom,tableOption) {
      var oldOption = $tableDom.bootstrapTable('getOptions');
      if (!!($('#table')).data("bootstrap.table")) {
        $tableDom.bootstrapTable('refreshOptions',tableOption)
      } else {
        $tableDom.bootstrapTable(tableOption);
      }
    }
    //获取助学金等
    function getdata(url,callback){
      $.ajax({
      url:url,
      type:"get",
      xhrFields: {
        withCredentials: true
      },
      crossDomain: true,
      success:function(data){
        callback(data);
      }
    })
    }
    //助学金
    function showtable(code,url,columns){
      tableOption.columns=columns;
      tableOption.data = [];
      loadTable($('#table'),tableOption);
      var arr=[];
      var arr1=[];
      getdata(url,function(result){
        var item=result.item;
        var rank="";
        var materials="";

        for(var i=0;i<item.length;i++){
          var items=item[i];
          if(items.rank==1){
            rank="一等助学金";
          }else if(items.rank==2){
            rank="二等助学金";
          }else{
            rank="三等助学金";
          }
          for(var j=0;j<items.materials.length;j++){
            materials+="<a>"+items.materials[j]+"&nbsp&nbsp,</a>";

          }

          if(items.status==1){
            var status="待审核";
            arr.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              rank:rank,
              materials:materials,
              time:items.time,
              status:status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='examine'>审核</a>"
            })
            materials="";
          }else{
            arr1.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              rank:rank,
              time:items.time,
              status:items.status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='deletedata'>删除</a>"
            })
            materials="";
          }

        }
        if(code==0){

          $("#table").bootstrapTable("load",arr);
          $("#inform").hide();
          $("#passtable").show();
        }
        if(code==1){
          $("#table").bootstrapTable("load",arr1);
          $("#inform").hide();
          $("#passtable").show();
        }

      })

    }
    //励志奖学金
    function showtable2(code,url,columns){
      tableOption.columns=columns;
      tableOption.data = [];
      loadTable($('#table'),tableOption);
      var arr=[];
      var arr1=[];
      getdata(url,function(result){
        var item=result.item;
        var materials="";
        for(var i=0;i<item.length;i++){
          var items=item[i];
          for(var j=0;j<items.materials.length;j++){
            materials+="<a>"+items.materials[j]+"&nbsp&nbsp,</a>";

          }
          if(items.status==1){
            var status="待审核";
            arr.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              time:items.time,
              status:status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='examine'>审核</a>"
            })
            materials="";
          }else{
            arr1.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              time:items.time,
              status:items.status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='deletedata'>删除</a>"
            })
            materials="";
          }

        }
        if(code==0){

          $("#table").bootstrapTable("load",arr);
          $("#inform").hide();
          $("#passtable").show();
        }
        if(code==1){
          $("#table").bootstrapTable("load",arr1);
          $("#inform").hide();
          $("#passtable").show();
        }

      })

    }
    //应善良助学金
    function showtable3(code,url,columns){
      tableOption.columns=columns;
      tableOption.data = [];
      loadTable($('#table'),tableOption);
      var arr=[];
      var arr1=[];
      getdata(url,function(result){
        var item=result.item;
        var materials="";
        var counselor="";
        for(var i=0;i<item.length;i++){
          var items=item[i];
          for(var j=0;j<items.materials.length;j++){
            materials+="<a>"+items.materials[j]+"&nbsp&nbsp,</a>";
          }
          for(var j=0;j<items.counselor.length;j++){
            counselor+="<a>"+items.counselor[j]+"&nbsp&nbsp,</a>";

          }
          if(items.status==1){
            var status="待审核";
            arr.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              counselor:counselor,
              time:items.time,
              status:status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='examine'>审核</a>"
            })
            materials="";
            counselor="";
          }else{
            arr1.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              counselor:counselor,
              time:items.time,
              status:items.status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='deletedata'>删除</a>"
            })
            materials="";
            counselor="";
          }

        }
        if(code==0){

          $("#table").bootstrapTable("load",arr);
          $("#inform").hide();
          $("#passtable").show();
        }
        if(code==1){
          $("#table").bootstrapTable("load",arr1);
          $("#inform").hide();
          $("#passtable").show();
        }

      })

    }
    //通知管理
    function showtable4(code,url,columns){
      tableOption.columns=columns;
      tableOption.data = [];
      loadTable($('#table'),tableOption);
      var arr=[];
      var arr1=[];
      getdata(url,function(result){
        var item=result.item;
        var materials="";
        var counselor="";
        for(var i=0;i<item.length;i++){
          var items=item[i];
          for(var j=0;j<items.materials.length;j++){
            materials+="<a>"+items.materials[j]+"&nbsp&nbsp,</a>";
          }
          for(var j=0;j<items.counselor.length;j++){
            counselor+="<a>"+items.counselor[j]+"&nbsp&nbsp,</a>";

          }
          if(items.status==1){
            var status="待审核";
            arr.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              counselor:counselor,
              time:items.time,
              status:status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='examine'>审核</a>"
            })
            materials="";
            counselor="";
          }else{
            arr1.push({
              num:i+1,
              person:items.person,
              college:items.college,
              applyname:"<a>"+items.applyname+"</a>",
              materials:materials,
              counselor:counselor,
              time:items.time,
              status:items.status,
              description:items.description,
              ways:"<a style='color:#337ab7;' class='deletedata'>删除</a>"
            })
            materials="";
            counselor="";
          }

        }
        if(code==0){
          $("#inform").show();
          $("#passtable").hide();
        }
        if(code==1){
          $("#table").bootstrapTable("load",arr1);
          $("#inform").hide();
          $("#passtable").show();
        }

      })

    }



