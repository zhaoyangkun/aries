$(function () {
    // getParentSort();

    border_change($(".at-sort-comment-a"), 10);

    //搜索框回车事件
    $('#search-input').bind('keypress', function (event) {
        if (event.keyCode === 13) {
            search_by_key();
        }
    });

    //首页搜索栏聚焦事件
    $("#search-input").focus(function () {
        $(".search-s").addClass("search-on"); //改变边框和icon的颜色
        $(".search-icon").addClass("icon-on");
    });

    //首页搜索栏失去焦点事件
    $("#search-input").blur(function () {
        $(".search-s").removeClass("search-on"); //恢复边框和icon的颜色
        $(".search-icon").removeClass("icon-on");
    });

    //首页搜索栏聚焦事件
    $("#search-input-m").focus(function () {
        $(".search-s").addClass("search-on"); //改变边框和icon的颜色
        $(".search-icon").addClass("icon-on");
    });

    //首页搜索栏失去焦点事件
    $("#search-input-m").blur(function () {
        $(".search-s").removeClass("search-on"); //恢复边框和icon的颜色
        $(".search-icon").removeClass("icon-on");
    });

    $(".a-login").hover(
        function () { //鼠标移入显示下拉栏
            $(this).next(".nav-ul-m").stop(true, false).toggle();
            {
                $(this).addClass('nb-a-hover');
            }
        },
        function () { //鼠标移出隐藏下拉栏
            $(this).next(".nav-ul-m").stop(true, false).toggle();
            {
                $(this).removeClass('nb-a-hover');
            }
        }
    );

    $(".nav-li-m").hover(
        function () { //鼠标移入下拉栏显示下拉栏，并改变下拉栏背景色和字体色
            $(this).parent(".nav-ul-m").stop(true, false).toggle();
            {
                $(this).parent(".nav-ul").prev('.nb-a').addClass('nb-a-hover');
            }
            {
                $(this).addClass('li-hover');
            }
        },
        function () { //鼠标移出下拉栏隐藏下拉栏
            $(this).parent(".nav-ul-m").stop(true, false).toggle();
            {
                $(this).parent(".nav-ul").prev('.nb-a').removeClass('nb-a-hover');
            }
            {
                $(this).removeClass('li-hover');
            }
        }
    );

    //隐藏/显示移动端侧边栏
    $("#mobile_cate").click(function (event) {
        $("#nav-m-list").delay(100).animate({
            right: '0'
        }, 500);
        $(document).one("click", function () { //对document绑定一个影藏Div方法
            $("#nav-m-list").delay(100).animate({
                right: '-250px'
            }, 500);
        });
        event.stopPropagation(); //阻止事件向上冒泡
        $("#nav-m-list").click(function (event) {
            event.stopPropagation(); //阻止事件向上冒泡
        });
    });

    $("#cancel").click(function () {
        $("#nav-m-list").delay(100).animate({
            right: '-250px'
        }, 500);
    });

    //加载进度条
    paceOptions = {
        ajax: false, // disabled
        document: false, // disabled
        eventLag: false, // disabled
        elements: {
            selectors: ['body']
        }
    };
});

//注销
function logout() {
    $.ajax({
        url: "/admin/login",
        type: 'DELETE',
        headers: {
            'X-CSRFToken': getCsrfToken()
        },
        success: function (data) {
            if (data.code === 1) {
                window.location.href = 'register.html';
            }
        },
        error: function (data) {
            console.log(data);
        }
    });

}

//关键字搜索
function search_by_key() {
    if ($("#search-input").val() === "" || $("#search-input").val() == null) {
        swal("请输入关键词！");
        return;
    }
    location.href = "/search?keyword=" + $("#search-input").val();
}

function m_search() {
    if ($("#search-input-m").val() === "" || $("#search-input-m").val() == null) {
        swal("请输入关键词！");
        return;
    }
    location.href = "/search?keyword=" + $("#search-input-m").val();
}

$(".nb-a").hover(
    function () { //鼠标移入显示下拉栏
        $(this).next(".nav-ul").stop(true, false).show();
        $(this).addClass('nb-a-hover');
    },
    function () { //鼠标移出隐藏下拉栏
        $(this).next(".nav-ul").stop(true, false).hide();
        $(this).removeClass('nb-a-hover');
    }
);

$(".nav-li").hover(
    function () { //鼠标移入下拉栏显示下拉栏，并改变下拉栏背景色和字体色
        $(this).parent(".nav-ul").stop(true, false).show();
        $(this).parent(".nav-ul").prev('.nb-a').addClass('nb-a-hover');
        $(this).addClass('li-hover');
    },
    function () { //鼠标移出下拉栏隐藏下拉栏
        $(this).parent(".nav-ul").stop(true, false).hide();
        $(this).parent(".nav-ul").prev('.nb-a').removeClass('nb-a-hover');
        $(this).removeClass('li-hover');
    }
);


/*标题栏底部边框长度变化动画*/
function border_change(obj, length) {
    const old_width = $(obj).width();
    const changed_width = old_width + length;
    //鼠标移入底部边框增大
    $(obj).mouseenter(
        function () {
            $(this).animate({
                width: changed_width
            });
        });
    //鼠标移出恢复原长度
    $(obj).mouseleave(function () {
        $(this).animate({
            width: old_width
        });
    })
}

/*//获取主类和子类
function getParentSort() {
    $.getJSON('/front/getParentCategory', function (data) {
        $.each(data.parent, function (i, val) {
            $("#category-ul").append('<li class="nav-li ta-c">' +
                '<a href="category.html/' + val.category_name + '">' + val.category_name + '</a>' +
                '</li>');
        });

        $(".nb-a").hover(
            function () {   //鼠标移入显示下拉栏
                $(this).next(".nav-ul").stop(true, false).show();
                $(this).addClass('nb-a-hover');
            }, function () {    //鼠标移出隐藏下拉栏
                $(this).next(".nav-ul").stop(true, false).hide();
                $(this).removeClass('nb-a-hover');
            }
        );

        $(".nav-li").hover(
            function () {   //鼠标移入下拉栏显示下拉栏，并改变下拉栏背景色和字体色
                $(this).parent(".nav-ul").stop(true, false).show();
                $(this).parent(".nav-ul").prev('.nb-a').addClass('nb-a-hover');
                $(this).addClass('li-hover');
            }, function () {    //鼠标移出下拉栏隐藏下拉栏
                $(this).parent(".nav-ul").stop(true, false).hide();
                $(this).parent(".nav-ul").prev('.nb-a').removeClass('nb-a-hover');
                $(this).removeClass('li-hover');
            }
        );

    })
}*/

class LazyImage {
    constructor(selector) {
        // 懒记载图片列表，将伪数组转为数组，以便可以使用数组的api
        this.lazyImages = Array.prototype.slice.call(document.querySelectorAll(selector));
        this.init()
    }

    inViewShow() {
        // 不支持IntersectionObserver api的情况下判断图片是否出现在可视区域内
        let len = this.lazyImages.length;
        for (let i = 0; i < len; i++) {
            let lazyImage = this.lazyImages[i];
            const rect = lazyImage.getBoundingClientRect();
            // 出现在视野的时候加载图片
            if (rect.top < document.documentElement.clientHeight) {
                lazyImage.src = lazyImage.dataset.src;
                // 移除掉已经显示的
                this.lazyImages.splice(i, 1);
                len--;
                i--;
                if (this.lazyImages.length === 0) {
                    // 如果全部都加载完 则去掉滚动事件监听
                    document.removeEventListener('scroll', this._throttleFn)
                }
            }
        }
    }

    throttle(fn, delay = 600, mustRun = 1200) {
        let t_start = null;
        let timer = null;
        let context = this;
        return function () {
            let t_current = +(new Date());
            let args = Array.prototype.slice.call(arguments);
            clearTimeout(timer);
            if (!t_start) {
                t_start = t_current
            }
            if (t_current - t_start > mustRun) {
                fn.apply(context, args);
                t_start = t_current
            } else {
                timer = setTimeout(() => {
                    fn.apply(context, args)
                }, delay)
            }
        }
    }

    init() {
        // 通过IntersectionObserver api判断图片是否出现在可视区域内，不需要监听Scroll来判断
        if ("IntersectionObserver" in window) {
            let lazyImageObserver = new IntersectionObserver((entries, observer) => {
                entries.forEach((entry, index) => {
                    // 如果元素可见
                    if (entry.isIntersecting) {
                        let lazyImage = entry.target;
                        lazyImage.src = lazyImage.dataset.src;
                        lazyImageObserver.unobserve(lazyImage)
                        // this.lazyImages.splice(index, 1)
                    }
                })
            });
            this.lazyImages.forEach(function (lazyImage) {
                lazyImageObserver.observe(lazyImage);
            })
        } else {
            this.inViewShow();
            this._throttleFn = this.throttle(this.inViewShow);
            document.addEventListener('scroll', this._throttleFn)
        }

    }
}

(function (window) {
    var a = function (as) {
        // default value
        var ps = {
            w: 40,
            h: 40,
            dImg: "http://olv6wm3nj.bkt.clouddn.com/18-3-22/16215533.jpg",
            hImg: "http://olv6wm3nj.bkt.clouddn.com/18-3-22/74337023.jpg",
            bt: 40,
            rg: 30,
            s: 300,
            th: 300
        }
        // merge value 
        if (as !== "undefined") {
            for (var key in as) {
                ps[key] = as[key];
            }
        }

        this.as = ps;
        this.init();
    }
    a.prototype.init = function () {
        var data = this.as;
        var d = document.createElement("div");
        d.style.width = data.w + "px";
        d.style.height = data.h + "px";
        d.style.position = "fixed";
        d.style.bottom = data.bt + "px"
        d.style.right = data.rg + "px"
        d.style.cursor = "pointer";
        d.style.backgroundImage = "url(" + data.dImg + ")";
        d.style.backgroundSize = "100%";
        d.style.display = "none";
        d.onmouseenter = function () {
            d.style.backgroundImage = "url(" + data.hImg + ")";
        }
        d.onmouseout = function () {
            d.style.backgroundImage = "url(" + data.dImg + ")";
        }
        document.body.onscroll = function () {
            if (document.documentElement.scrollTop >= data.th) {
                d.style.display = "inline-block";
            } else {
                d.style.display = "none";
            }
        }

        d.onclick = function () {
            var timer = requestAnimationFrame(function fn() {
                var oTop = document.body.scrollTop || document.documentElement.scrollTop;
                if (oTop > 0) {
                    document.body.scrollTop = document.documentElement.scrollTop = oTop - data.s;
                    timer = requestAnimationFrame(fn);
                } else {
                    d.style.display = "none";
                    d.style.backgroundImage = "url(" + data.dImg + ")";
                    cancelAnimationFrame(timer);
                }
            });
        }
        document.body.appendChild(d);
    }
    return window.Top = a;
})(window);

//返回顶部
const top_to = new Top({
    dImg: "../img/up.svg",
    hImg: "../img/up-on.svg"
});