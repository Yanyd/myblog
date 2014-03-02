{{define "footer2"}}
<footer class="footer">
    <div class="container">
    	<p class="pull-right">创建时间 : 2014-01-01　　博客数量 : 100　　总浏览量 : 100</p>
      <p>Copyright　&copy; 2014　<a href="###" rel="tooltip" data-placement="top" data-original-title="Email:659986134#qq#com">颜少　</a>All Rights Reserved .</p>
    	<p class="pull-right">更新时间 : 2014-01-01　　留言数量 : 100　　总评论量 : 100</p>
      <p>Powered by <a target="_blank" href="http://www.golang.org">Golang</a> & <a target="_blank" href="http://www.mysql.com">MySql</a> & <a target="_blank" href="http://www.beego.me">Beego</a> & <a target="_blank" href="http://www.getbootstrap.com">Bootstrap</a> . 
      </p>

    </div>
</footer>

</body>
<script>
	$(function(){
		//启动提示工具
		$(".container [rel=tooltip]").tooltip();
		//滚动后固定导航条	
		var didScroll = false;
		//滚动条发生变化时触发
		$(window).scroll(function () {
		   didScroll = true;
		});
		var y;
		setInterval(function () {
		   if (didScroll) {
		        didScroll = false;
				y=$(window).scrollTop();
				//70 待思考
				if(y>70){
					$('body').css('padding-top','70px');
					//顶部介绍栏会比导航栏多出20px，有个黑色背景
					$('#intro').css('opacity', '0');
					$('#nav_top').addClass('navbar-fixed-top');
				}else{
					$('body').css('padding-top','0px');
					$('#intro').css('opacity', '1');
					$('#nav_top').removeClass('navbar-fixed-top');
				}
			}
		}, 50);

		//

	   //回到顶部按钮 start
	   var toTopElem=$('<a id="toTop" class="btn btn-default btn-lg"><i class="glyphicon glyphicon-chevron-up"></i></a>').appendTo($('body')).click(function(event) {
	   	      $('html,body').animate({scrollTop:0}, 120);
	   });
	   var toTopFunc=function(){
	   		var st=$(document).scrollTop();
	   		(st>0)?toTopElem.show():toTopElem.hide();
	   }
	   $(window).bind('scroll',toTopFunc);
       $(function(){
       		toTopFunc();
       		$("#toTop").css({
       			"position": "fixed",
       			"_position": "absolute",
       			"right": "32px",
       			"bottom": "30px",
       			"_bottom": "auto",
       			"cursor": "pointer"
       		});
       })
       //回到顶部按钮 end


       //修改导航栏点击和移出样式 start
        $(".navbar-nav .active a").css({'color':'#f04c5c','background-color':'#fff'});
       $(".navbar-nav >li").hover(function() {
	       	if($(this).hasClass('active')){
	           $(this).children('a').css({'background-color': '#fff'});
	       	}else{
	           $(this).children('a').css({'background-color': '#E7E7E7'});
	       	}
       }, function() {
	        $(this).children('a').css({'background-color': '#fff'});
       });
      
       //导航栏样式修改 end
	})

</script>
{{end}}