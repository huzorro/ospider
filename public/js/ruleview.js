
$(function() {
    $(window).load(function(){
        $.ajax({
            url:"/api/spiders",
            type:"post",
            cache:false,
            dataType:"json",                        
            success:spiders
        });
    });    
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
        console.log("Spiderid" + $("select[name=SpiderId]").val());
        $.ajax({
         url : "/rule/add",
         data : {Name:$("input[name=Name]").val(), SpiderId:$("select[name=SpiderId]").val(),
                Url:$("input[name=Url]").val(), Section:$("input[name=Section]").val(),
                Href:$("input[name=Href]").val(), Title:$("input[name=Title]").val(),
                Content:$("input[name=Content]").val(), Filter:$("input[name=Filter]").val()
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
         url : "/rule/one",
         data : {Id:$(this).val()},
         type : "post",
         cache : false,
         dataType : "json",
         success: one
         });
    });
    $("#edit").click(function() {
        $.ajax({
         url : "/rule/edit",
         data : {Id:$("#Id").val(), Name:$("#Name").val(), 
                 SpiderId:$("#SpiderId").val(), Url:$("#Url").val(),
                 Section:$("#Section").val(), Href:$("#Href").val(),
                 Title:$("#Title").val(), Content:$("#Content").val(),
                 Filter:$("#Filter").val(),               
                Status:$("input[name=Status]:checked").val()},
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
            // $("#Spiderid").val(json.spiderid);
            $("#Url").val(json.url);
            $("#Section").val(json.selector.section);
            $("#Href").val(json.selector.href);
            $("#Title").val(json.selector.title);
            $("#Content").val(json.selector.content);
            $("#Filter").val(json.selector.filter);            
            if (json.status) {
                htmls.push('<input type="radio" name="Status" value=0 >停用<input type="radio" name="Status" value=1 checked>启用');
            } else {
                htmls.push('<input type="radio" name="Status" value=0 checked>停用<input type="radio" name="Status" value=1 >启用');
            }
            $("#Status").html(htmls.join(""));
            //select
            $("#SpiderId>option").each(function(){
                 if($(this).val() == json.spiderid) {
                     $(this).attr("selected","selected");
                 }
            });
        }
}

function spiders(json) {
        console.log(json);
        if (json.status !== undefined && json.text !== undefined) {
            $("#Status").html(json.text);
        } else {            
            var htmls = [];
            $.each(json, function(key, value) {
                console.log(key+value.id);
                htmls.push("<option value=" + value.id + ">" + value.name + "</option>")     
            });
            $("select[name=SpiderId]").html(htmls.join(""));
            $("#SpiderId").html(htmls.join(""))
        }    
} 