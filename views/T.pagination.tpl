{{define "pagination"}}
<script>
	$(function(){
  
  page({
    id : 'pagination',
    nowNum : {{.Page.PageNow}},
    allNum : {{.Page.TotalPageNum}}
  });
  
}) ;

function page(opt){

  if(!opt.id){return false};
  
  var obj = document.getElementById(opt.id);
  
  var nowNum = opt.nowNum || {{.Page.PageNow}};
  var allNum = opt.allNum || {{.Page.TotalPageNum}};
  
  
  if( nowNum>=4 && allNum>=6 ){
  
    var oLi = document.createElement('li');
    var oA = document.createElement('a');
    oA.href = '/{{.Router}}?op=manage&pageNow=1';
    oA.innerHTML = '首页';
    oLi.appendChild(oA);
    obj.appendChild(oLi);
  
  }
  
  if(nowNum>=2){
    var oLi = document.createElement('li');
    var oA = document.createElement('a');
    oA.href = '/{{.Router}}?op=manage&pageNow=' + (nowNum - 1);
    oA.innerHTML = '上一页';
    oLi.appendChild(oA);
    obj.appendChild(oLi);
  }
  
  if(allNum<=5){
    for(var i=1;i<=allNum;i++){
      var oLi = document.createElement('li');
      var oA = document.createElement('a');
      oA.href = '/{{.Router}}?op=manage&pageNow=' + i;
      oA.innerHTML = i;
      if(nowNum == i){
        $(oLi).addClass('active');
      }
      else{
        $(oLi).removeClass('active');
      }
      oLi.appendChild(oA);
      obj.appendChild(oLi);
    } 
  }
  else{
  
    for(var i=1;i<=5;i++){
      var oLi = document.createElement('li');
      var oA = document.createElement('a');
      
      if(nowNum == 1 || nowNum == 2){
        
        oA.href = '/{{.Router}}?op=manage&pageNow=' + i;
        oA.innerHTML = i;
        if(nowNum == i){
          $(oLi).addClass('active');
        }
        else{
          $(oLi).removeClass('active');
        }
        
      }
      else if( (allNum - nowNum) == 0 || (allNum - nowNum) == 1 ){
      
        oA.href = '/{{.Router}}?op=manage&pageNow=' + (allNum - 5 + i);
        
        if((allNum - nowNum) == 0 && i==5){
          oA.innerHTML = (allNum - 5 + i);
          $(oLi).addClass('active');
        }
        else if((allNum - nowNum) == 1 && i==4){
          oA.innerHTML = (allNum - 5 + i);
          $(oLi).addClass('active');
        }
        else{
          oA.innerHTML = (allNum - 5 + i) ;
          $(oLi).removeClass('active');
        }
      
      }
      else{
        oA.href = '/{{.Router}}?op=manage&pageNow=' + (nowNum - 3 + i);
        
        if(i==3){
          oA.innerHTML = (nowNum - 3 + i);
          $(oLi).addClass('active');
        }
        else{
          oA.innerHTML = (nowNum - 3 + i) ;
          $(oLi).removeClass('active');
        }
      }
      oLi.appendChild(oA);
      obj.appendChild(oLi);
      
    }
  
  }
  
  if( (allNum - nowNum) >= 1 ){
    var oLi = document.createElement('li');
    var oA = document.createElement('a');
    oA.href = '/{{.Router}}?op=manage&pageNow=' + (nowNum + 1);
    oA.innerHTML = '下一页';
    oLi.appendChild(oA);
    obj.appendChild(oLi);
  }
  
  if( (allNum - nowNum) >= 3 && allNum>=6 ){
    var oLi = document.createElement('li');
    var oA = document.createElement('a');
    oA.href = '/{{.Router}}?op=manage&pageNow=' + allNum;
    oA.innerHTML = '尾页';
    oLi.appendChild(oA);
    obj.appendChild(oLi);
  
  }
  
}
</script>
{{end}}