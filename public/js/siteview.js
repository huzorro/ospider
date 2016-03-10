
$(function() {
   $("input[name=Url]").change(function () {
       console.log($("input[name=Url]").val());
        // var htmls = [];
        $.ajax({
            url:$(this).val() + "/api/category",
            type:"get",
            cache:false,
            dataType:"json",                        
            success:function(result) {
                var htmls = [];
                // htmls.push($("select[name=Category]").html());
                $.each(result.data, function(key, value){
                    htmls.push("<option value=" + JSON.stringify(value) + ">" + value.title + "</option>");
                });
                $("select[name=Category]").html(htmls.join(""));                
            }           
        });
        
        $.ajax({
            url:$(this).val() + "/api/member",
            type:"get",
            cache:false,
            dataType:"json",                            
            success:function(result) {
                // htmls.push($("select[name=Category]").html());
                var htmls = []; 
                $.each(result.data, function(key, value){
                    htmls.push("<option value=" + JSON.stringify(value) + ">" + value.nickname + "</option>");
                });
                $("select[name=Member]").html(htmls.join(""));                
            }
        });
        $.ajax({
            url:"/api/rules",
            type:"post",
            cache:false,
            dataType:"json",                            
            success:function(result) {
                // htmls.push($("select[name=Category]").html());
                var htmls = []; 
                $.each(result, function(key, value){
                    htmls.push("<option value=" + value.id + ">" + value.name + "</option>");
                });
                $("select[name=RuleId]").html(htmls.join(""));                
            }
        });
        $.ajax({
            url:"/api/crontabs",
            type:"post",
            cache:false,
            dataType:"json",                            
            success:function(result) {
                // htmls.push($("select[name=Category]").html());
                var htmls = []; 
                $.each(result, function(key, value){
                    htmls.push("<option value=" + value.id + ">" + value.name + "</option>");
                });
                $("select[name=CronId]").html(htmls.join(""));                
            }
        });        
   })
   
    // $("select[name=Spiderid]").load(function() {
    //     alert("文档加载完毕!");
    //     console.log("select load....");
    //     // $.ajax({
    //     //     url:"/api/spiders",
    //     //     type:"post",
    //     //     cache:false,
    //     //     dataType:"json",                        
    //     //     success:spiders
    //     // });
    // });
    $("#add").click(function() {
        // console.log("Spiderid" + $("select[name=SpiderId]").val());
        var arr = [];
        $("input[name=Position]:checked").each(function(){
            arr.push(parseInt($(this).val()));
        });
        var category = JSON.parse($("select[name=Category]").val());
        var member = JSON.parse($("select[name=Member]").val());
        
        $.ajax({
         url : "/site/add",
         data : {Name:$("input[name=Name]").val(),
                Url:$("input[name=Url]").val(), Category:category.id,
                CategoryTitle:category.title,
                ModelId:category.model, GroupId:$("input[name=GroupId]").val(),
                Member:member.uid, NickName:member.nickname,
                Position:eval(arr.join('+')),                
                Display:$("input[name=Display]:checked").val(), Check:$("input[name=Check]:checked").val(),
                RuleId:$("select[name=RuleId]").val(), CronId:$("select[name=CronId]").val()
         },
         type : "post",
         cache : false,
         dataType : "json",
         success: commonInfo
         });
    });
//    $("#my-tab-content").delegate("button[name=operate]", "click", function() {
//         console.log($(this).val());
//        $.ajax({
//         url : "/key/one",
//         data : {Id:$(this).val()},
//         type : "post",
//         cache : false,
//         dataType : "json",
//         success: keyone
//         });
//    });
    $("button[name=operate]").click(function() {
         console.log($(this).val() + "Abc");
        $.ajax({
         url : "/site/one",
         data : {Id:$(this).val()},
         type : "post",
         cache : false,
         dataType : "json",
         success: one
         });
    });
    $("#edit").click(function() {
        var arr = [];
        $("input[name=Position_edit]:checked").each(function(){
            arr.push(parseInt($(this).val()));
        });
        var category = JSON.parse($("#Category").val());
        var member = JSON.parse($("#Member").val());
        
        $.ajax({
         url : "/site/edit",
         data : {Id:$("#Id").val(), Name:$("#Name").val(),
                Url:$("#Url").val(), Category:category.id,
                CategoryTitle:category.title,
                ModelId:category.model,
                Member:member.uid, NickName:member.nickname,
                Position:eval(arr.join('+')),                
                Display:$("input[name=Display_edit]:checked").val(), Check:$("input[name=Check_edit]:checked").val(),
                RuleId:$("#RuleId").val(), CronId:$("#CronId").val(),
                Status:$("input[name=Status]:checked").val()
         },
         type : "post",
         cache : false,
         dataType : "json",
         success: commonInfo
         });
    });
    $("#infoModal .btn").click(function(){
        location.reload();
    });
    $("#close").click(function(){
        location.reload();
    });
});


function commonInfo(json) {
    if (json.status !== undefined) {
        $('#infoModal').modal('toggle');
        $('#infoModal p').text(json.text);
//        location.reload();
        return
    }
}

function one(json) {
        console.log(json);
        $('#editRuleModal').modal('toggle');
        if (json.status !== undefined && json.text !== undefined) {
            $("#Status").html(json.text);
        } else {
            var htmls = [];
            $("#Id").val(json.id);
            $("#Name").val(json.name);
            $("#Url").val(json.url);
            
            //
            $.ajax({
                url:json.url + "/api/category",
                type:"get",
                cache:false,
                dataType:"json",                        
                success:function(result) {
                    var htmls = [];
                    // htmls.push($("select[name=Category]").html());
                    $.each(result.data, function(key, value){
                        htmls.push("<option value=" + JSON.stringify(value) + ">" + value.title + "</option>");
                    });
                    $("#Category").html(htmls.join(""));
                    $("#Category>option").each(function() {
                        if($(this).text() == json.document_set.title) {
                            $(this).attr("selected","selected");
                        }
                    });                
                }           
            });            
            $.ajax({
                url:json.url + "/api/member",
                type:"get",
                cache:false,
                dataType:"json",                            
                success:function(result) {
                    // htmls.push($("select[name=Category]").html());
                    var htmls = []; 
                    $.each(result.data, function(key, value){
                        htmls.push("<option value=" + JSON.stringify(value) + ">" + value.nickname + "</option>");
                    });
                    $("#Member").html(htmls.join("")); 
                    $("#Member>option").each(function() {
                        if($(this).text() == json.document_set.nickname) {
                            $(this).attr("selected","selected");
                        }
                    });                                    
                }
            });
            $.ajax({
                url:"/api/rules",
                type:"post",
                cache:false,
                dataType:"json",                            
                success:function(result) {
                    // htmls.push($("select[name=Category]").html());
                    var htmls = []; 
                    $.each(result, function(key, value){
                        htmls.push("<option value=" + value.id + ">" + value.name + "</option>");
                    });
                    $("#RuleId").html(htmls.join(""));
                    $("#RuleId>option").each(function() {
                        if($(this).text() == json.Rule.name) {
                            $(this).attr("selected","selected");
                        }
                    });                                         
                }
            });
            $.ajax({
                url:"/api/crontabs",
                type:"post",
                cache:false,
                dataType:"json",                            
                success:function(result) {
                    // htmls.push($("select[name=Category]").html());
                    var htmls = []; 
                    $.each(result, function(key, value){
                        htmls.push("<option value=" + value.id + ">" + value.name + "</option>");
                    });
                    $("#CronId").html(htmls.join(""));
                    $("#CronId>option").each(function() {
                        if($(this).text() == json.Cron.name) {
                            $(this).attr("selected","selected");
                        }
                    });                                  
                }
            });
            $("input[name=Position_edit]").each(function(){
               if(($(this).val() & json.document_set.position) == $(this).val()) {
                   $(this).attr("checked","checked");
               }
            });             
            $("input[name=Display_edit]").each(function(){
               if($(this).val() == json.document_set.display) {
                   $(this).attr("checked","checked");
               }
            });               
            $("input[name=Check_edit]").each(function(){
               if($(this).val() == json.document_set.check) {
                   $(this).attr("checked","checked");
               }
            });               
            
            if (json.status) {
                htmls.push('<input type="radio" name="Status" value=0 >停用<input type="radio" name="Status" value=1 checked>启用');
            } else {
                htmls.push('<input type="radio" name="Status" value=0 checked>停用<input type="radio" name="Status" value=1 >启用');
            }
            $("#Status").html(htmls.join(""));
        }
}
