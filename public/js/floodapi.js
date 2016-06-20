
$(function() {
    $("#add").click(function() {
        $.ajax({
         url : "/floodapi/add",
         data : {Name:$("input[name=Name]").val(), Api:$("input[name=Api]").val(),
                 Powerlevel:$("input[name=Powerlevel]").val(),
                 Time:$("input[name=Time]").val()},
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
         url : "/floodapi/one",
         data : {Id:$(this).val()},
         type : "post",
         cache : false,
         dataType : "json",
         success: one
         });
    });
    $("#edit").click(function() {
        console.log($("#Task").val());
        console.log($("input[name=Status]:checked").val());
        $.ajax({
         url : "/floodapi/edit",
         data : {Id:$("#Id").val(), Name:$("#Name").val(), Api:$("#Api").val(), Time:$("#Time").val(),
                 Powerlevel:$("#Powerlevel").val(), Status:$("input[name=Status]:checked").val()},
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
        $('#editSpiderModal').modal('toggle');
        if (json.status !== undefined && json.text !== undefined) {
            $("#Status").html(json.text);
        } else {            
            var htmls = [];
            $("#Id").val(json.id);
            $("#Name").val(json.name);
            $("#Api").val(json.api);
            $("#Powerlevel").val(json.powerlevel);
            $("#Time").val(json.time);
            if (json.status) {
                htmls.push('<input type="radio" name="Status" value=0 >停用<input type="radio" name="Status" value=1 checked>启用');
            } else {
                htmls.push('<input type="radio" name="Status" value=0 checked>停用<input type="radio" name="Status" value=1 >启用');
            }
            $("#Status").html(htmls.join(""));
        }
}
