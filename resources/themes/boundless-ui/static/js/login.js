$(document).ready(function () {

    //切换至注册窗口
    $("#registered").click(function () {
        toRegis();
    });

    //切换至登录窗口
    $("#landing").click(function () {
        toLogin();
    });

    //回车事件
    $(document).keyup(function (event) {
        if (event.keyCode === 13) {
            if ($("#landing-content").css('display') !== 'none') {
                $("#login").click();
            } else {
                $("#register").click();
            }
        }
    });

});

//切换至登录窗口
function toLogin() {
    // $("#registered").removeClass("on");
    // $("#registered").addClass("not-on");
    // $("#landing").addClass("on");
    // $("#landing-content").show(500);
    // $("#registered-content").hide(500);
    window.location.href = 'login.html';
}

//切换至注册窗口
function toRegis() {
    // $("#landing").removeClass("on");
    // $("#landing").addClass("not-on");
    // $("#landing-content").hide(500);
    // $("#registered").addClass("on");
    // $("#registered-content").show(500);
    window.location.href = 'register.html';
}

function send_email() {
    let reg = /^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/; //正则表达式
    let to_email = $("#email_r").val();
    if (!reg.test(to_email)) {
        swal({
            text: "请输入格式正确的邮箱",
            icon: "error",
            timer: 1500
        });
        return;
    }
    $.ajax({
        url: "/front/send_email",
        type: "GET",
        data: {"to_email": to_email},
        datatype: "json",
        success: function (data) {
            if (data.code === 1) {
                swal({
                    text: data.msg,
                    icon: "success",
                    timer: 1500
                }).then((value) => {
                    count_down();
                });
                setTimeout("count_down();", 1500);
            } else if (data.code === 0) {
                swal({
                    text: data.msg,
                    icon: "error",
                    timer: 1500
                });
            }
        },
        error: function (data) {
            swal({
                text: "发送异常",
                icon: "error",
                timer: 1500
            });
        }
    });
}


function count_down() {
    let code = $("#send-code");
    code.attr("disabled", "disabled");
    setTimeout(function () {
        code.css("opacity", "0.6");
    }, 1000);
    let time = 60;
    let set = setInterval(function () {
        code.val("(" + --time + ")秒后重新获取");
    }, 1000);
    setTimeout(function () {
        code.attr("disabled", false).val("发送验证码");
        code.css("opacity", "1");
        clearInterval(set);
    }, 60000);
}



