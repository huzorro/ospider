
$(function() {
   $("input[name=Url]").change(function () {
       console.log($("input[name=Url]").val());
        // var htmls = [];
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
        // var arr = [];
        // $("input[name=Position]:checked").each(function(){
        //     arr.push(parseInt($(this).val()));
        // });
        // var category = JSON.parse($("select[name=Category]").val());
        // var member = JSON.parse($("select[name=Member]").val());
        
        $.ajax({
         url : "/floodtarget/add",
         data : {Host:$("input[name=Host]").val(),
                Url:$("input[name=Url]").val(),
                CronId:$("select[name=CronId]").val()
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
         url : "/floodtarget/one",
         data : {Id:$(this).val()},
         type : "post",
         cache : false,
         dataType : "json",
         success: one
         });
    });
    $("#edit").click(function() {       
        $.ajax({
         url : "/floodtarget/edit",
         data : {Id:$("#Id").val(), Host:$("#Host").val(),
                Url:$("#Url").val(), CronId:$("#CronId").val(),
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
            $("#Host").val(json.host);
            $("#Url").val(json.url);
            
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
            
            
            if (json.status) {
                htmls.push('<input type="radio" name="Status" value=0 >停用<input type="radio" name="Status" value=1 checked>启用');
            } else {
                htmls.push('<input type="radio" name="Status" value=0 checked>停用<input type="radio" name="Status" value=1 >启用');
            }
            $("#Status").html(htmls.join(""));
        }
}
