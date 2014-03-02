{{define "pagination2"}}
<script>

function page(opt){

  if(!opt.id){return false};

  var nowNum = opt.nowNum || {{.Page.PageNow}};
  var allNum = opt.allNum || {{.Page.TotalPageNum}};
  var time   = opt.time  || '2001-06-01'
  if( nowNum>=4 && allNum>=6 ){
  
    var cloneObj=$("li[name=pg_tpl]").clone();
    $(cloneObj).appendTo("#"+opt.id);
    $(cloneObj).attr({ "name": time+'/1'});
    $(cloneObj).children('a').html('首页')
    $(cloneObj).show();
  }
  
  if(nowNum>=2){
    var cloneObj=$("li[name=pg_tpl]").clone();
    $(cloneObj).appendTo("#"+opt.id);
    $(cloneObj).attr({ "name": time+'/'+ (nowNum - 1)});
    $(cloneObj).children('a').html('上一页')
    $(cloneObj).show();

  }
  
  if(allNum<=5){
    for(var i=1;i<=allNum;i++){
      var cloneObj=$("li[name=pg_tpl]").clone();
      $(cloneObj).appendTo("#"+opt.id);
      $(cloneObj).attr({ "name": time+'/'+i});
      $(cloneObj).children('a').html(i)

      if(nowNum == i){
        $(cloneObj).addClass('active');
      }
      else{
        $(cloneObj).removeClass('active');
      }

      $(cloneObj).show();
    } 
  }
  else{
  
    for(var i=1;i<=5;i++){


      var cloneObj=$("li[name=pg_tpl]").clone();
      $(cloneObj).appendTo("#"+opt.id);
      
      if(nowNum == 1 || nowNum == 2){
        $(cloneObj).attr({ "name": time+'/'+i});
        $(cloneObj).children('a').html(i)

        if(nowNum == i){
          $(cloneObj).addClass('active');
        }
        else{
          $(cloneObj).removeClass('active');
        }
        $(cloneObj).show();
      }
      else if( (allNum - nowNum) == 0 || (allNum - nowNum) == 1 ){
        
        $(cloneObj).attr({ "name": time+'/'+(allNum - 5 + i)});
        
        if((allNum - nowNum) == 0 && i==5){
          $(cloneObj).children('a').html((allNum - 5 + i));
          $(cloneObj).addClass('active');
        }
        else if((allNum - nowNum) == 1 && i==4){
          $(cloneObj).children('a').html((allNum - 5 + i));
          $(cloneObj).addClass('active');
        }
        else{
          $(cloneObj).children('a').html((allNum - 5 + i) );
          $(cloneObj).removeClass('active');
        }
        $(cloneObj).show();
      }
      else{
        $(cloneObj).attr({ "name": time+'/'+(nowNum - 3 + i)});
        
        if(i==3){
          $(cloneObj).children('a').html((nowNum - 3 + i));
          $(cloneObj).addClass('active');
        }
        else{
         $(cloneObj).children('a').html((nowNum - 3 + i));
         $(cloneObj).removeClass('active');
        }
        $(cloneObj).show();
      }
      
    }
  
  }
  
  if( (allNum - nowNum) >= 1 ){

    var cloneObj=$("li[name=pg_tpl]").clone();
    $(cloneObj).appendTo("#"+opt.id);
    $(cloneObj).attr({ "name": time+'/'+(nowNum + 1)});
    $(cloneObj).children('a').html('下一页');
    $(cloneObj).show();
  }
  
  if( (allNum - nowNum) >= 3 && allNum>=6 ){

    var cloneObj=$("li[name=pg_tpl]").clone();
    $(cloneObj).appendTo("#"+opt.id);
    $(cloneObj).attr({ "name":time+'/'+allNum});
    $(cloneObj).children('a').html('尾页');
    $(cloneObj).show();
  }
}
</script>
{{end}}