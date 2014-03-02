{{define "autosave"}}
<!--编辑器配置-->
<script charset="utf-8" src="/static/kindeditor/kindeditor.js"></script>
<script charset="utf-8" src="/static/kindeditor/lang/zh_CN.js"></script>
<script charset="utf-8" src="/static/js/jquery.form.min.js"></script>
<script>
/*样式与bootstrap冲突*/
 KindEditor.ready(function(K) {
      window.editor = K.create('#editor_id',{
        allowUpload : true, 
        urlType : 'absolute',
        uploadJson : '/findeditor',
        afterChange:function(){this.sync();},
        afterFocus:function(){this.sync();},
        afterBlur: function(){this.sync();}  //将HTML数据设置到原来的textarea
      });
});

//分清 新增和更新 时的自动保存

var delay;
$(function(){
   //自动保存
   //setTimeout 执行一次  clearTimeout
   //setInterval 执行多次
   var t;
   t =setInterval(autoSave,5000);//5s 
  //发布   保存草稿
  $("#form_add_btn,#form_draft_btn,#form_upd_btn").click(function(event) {
    //停掉定时器
    clearInterval(t);
    var content=$('#editor_id').val();
    var title=$('input[name=title]').val();
    if(notEmpty(content)&&notEmpty(title)){
      var urlStr;
      var isDraft=$(this).hasClass('draft')
      if(isDraft){
        urlStr="/topic/post/draft";
      }else{
        urlStr="/topic"
      }
      $('#topic_form').ajaxSubmit({
        url:urlStr,
        success:function(data){
          $(window)[0].location.href="/topic/view/"+$.parseJSON(data).id;
      }}); 

    }else{
      $('.save_info').removeClass('hidden');
      if(!notEmpty(title)){
         $(".save_info").html("请填写标题.");
         t =setInterval(autoSave,5000);
      }else if(!notEmpty(content)){
         $(".save_info").html("请填写内容.");
         t =setInterval(autoSave,5000);
      }
      clearTimeout(delay);
      delay=setTimeout(function(){
        $('.save_info').addClass('hidden');
      },3000);
    }
  });
  //更新时取消修改
  $('#form_cnl_upd_btn').click(function(event) {
      clearInterval(t);
      //复原 判断是修改为正文还是草稿
      $('input[name=title]').val({{.Topic.Title}});
      $('input[name=label]').val({{.Topic.Label}});
      $('#editor_id').val({{.Topic.Content}});
      $('input:checkbox').each(function () { 
              $(this).removeAttr('checked') 
      });
      {{range .Categories}}
          {{if .Checked}}
              $('#checkbox_'+{{.Id}})[0].checked='checked';
          {{end}}
      {{end}}

      var urlStr='/topic/post/draft';
      {{if not .Topic.Draft}}
         urlStr='/topic/post';
      {{end}}
      $('#topic_form').ajaxSubmit({
        url:urlStr,
        success:function(data){
          {{if not .Topic.Draft}}
              $(window)[0].location.href="/topic?op=manage&pageNow="+{{.PageNow}};
          {{else}}
              $(window)[0].location.href="/topic?op=manage&draft=true";
          {{end}}
      }}); 

  });
  //取消
  $('#form_cnl_btn').click(function(event) {
    clearInterval(t);
    var id=$('#hidden_id').val();
    if(!('0'===id)){
      //删除已经保存的草稿
      $.get('/topic/del/'+id, function(data) {
      });
    }
    $(window)[0].location.href="/topic?op=manage";
  });

})

//每隔50s就自动保存
function autoSave(){
   var content=$('#editor_id').val();
   var title=$('input[name=title]').val();
   if(notEmpty(content)){
    if(!notEmpty(title)){
       $('input[name=title]').val('这是一篇草稿.');
    }
    //提交保存为草稿 修改时要判断是保存为正文还是草稿
    var urlStr='/topic/post/draft';
    {{if .UpdTopic}}
        {{if not .Topic.Draft}}
           urlStr='/topic/post';
        {{end}}
    {{end}}

    $('#topic_form').ajaxSubmit({
      url:urlStr,
      dataType:'json',
      success:function(data){
      //提示已经保存信息
      $('.save_info').removeClass('hidden');
      $(".save_info").html("文章已于 "+(new Date().Format("yyyy-MM-dd hh:mm:ss"))+" 自动保存.");
      $('#hidden_id').val(data.id);
      {{if not .UpdTopic}}
        if('这是一篇草稿.'===$('input[name=title]').val()){
            $('input[name=title]').val('');
        }
      {{end}}
      clearTimeout(delay);
      //隐藏提示信息,时间长一点,否则频繁出现。
      delay=setTimeout(function(){
        $('.save_info').addClass('hidden');
      },10000);
    }}); 
   }
}

</script>
{{end}}