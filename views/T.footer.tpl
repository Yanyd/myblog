{{define "footer"}}
<div  class="container page_bottom">
	<hr style="border:0;background-color:#d7d7d7;height:1px;">
	<i class="icon-bow"></i>
	<div id="footer">
        	<span>版权所有　&copy; 2014　<a href="###" rel="tooltip" data-placement="top" data-original-title="Email:659986134@qq.com" style="text-decoration:none;">颜少</a>　Powered by <a target="_blank" href="http://www.golang.org">Golang</a> . </span>
	</div>
</div>
</body>
<script>
	$(function(){
		//启动提示工具
		$(".container [rel=tooltip]").tooltip();

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