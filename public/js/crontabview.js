
$(function() {
    $("#add").click(function() {
        $.ajax({
         url : "/crontab/add",
         data : {Name:$("input[name=Name]").val(), Cron:$("input[name=Cron]").val()},         
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
         url : "/crontab/one",
         data : {Id:$(this).val()},
         type : "post",
         cache : false,
         dataType : "json",
         success: one
         });
    });
    $("#edit").click(function() {
        console.log($("input[name=Status]:checked").val());
        $.ajax({
         url : "/crontab/edit",
         data : {Id:$("#Id").val(), Name:$("#Name").val(), Cron:$("#Cron").val(), Status:$("input[name=Status]:checked").val()},
         type : "post",
         cache : false,
         dataType : "json",
         success: commonInfo
         });
    });
    $("#infoModal .btn").click(function(){
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
        console.log("abc");
        $('#editCrontabModal').modal('toggle');
        if (json.status !== undefined && json.text !== undefined) {
            $("#Status").html(json.text);
        } else {            
            var htmls = [];
            $("#Id").val(json.id);
            $("#Name").val(json.name);
            $("#Cron").val(json.cron);
            if (json.status) {
                htmls.push('<input type="radio" name="Status" value=0 >停用<input type="radio" name="Status" value=1 checked>启用');
            } else {
                htmls.push('<input type="radio" name="Status" value=0 checked>停用<input type="radio" name="Status" value=1 >启用');
            }
            $("#Status").html(htmls.join(""));
        }
}
