{{define "comment_form"}}
  <form id="comment_form" method="POST" action="/comment/add" style="margin-top:24px;margin-bottom:6px;">
    <input id="hidden_pid" type="hidden" name="pid" value="" />
    <input type="hidden" name="tid" value="{{.Topic.Id}}" />
    <div class="row">
       <div class="col-md-4">
          <div class="form-group" style="margin-top:8px;margin-bottom:30px;">
            <div class="input-group">
            <span class="input-group-addon">显示昵称</span>
            <input type="text" name="nickname" class="form-control" maxlength="20" placeholder="请输入您的昵称.">
            </div>
          </div>    
          <div class="form-group" style="margin-bottom:30px;">
            <div class="input-group">
            <span class="input-group-addon">E m a i l</span>
            <input type="text" name="email" class="form-control" maxlength="50" placeholder="请输入您的邮箱,不会在评论中显示哦.">
            </div>
          </div>
          <div class="form-group">
            <div class="input-group">
            <span class="input-group-addon">您的站点</span>
            <input type="text" name="website" class="form-control" maxlength="50" placeholder="如果您愿意，可以留个地址.">
            </div>
          </div>
      </div>
      <div class="col-md-4">
        <div class="form-group">
          <label class="control-label">评论内容:</label>
          <textarea name="content" id="" cols="30" rows="7" class="form-control" maxlength="300" placeholder="您怎么看."></textarea>
        </div>
        <a id="submit_btn" class="btn btn-default">提交评论</a>　　　
        <a id="cancle_btn"  class="btn btn-default" style="display:none;">取　　消</a>
      </div>
    </div><!--row end-->
    </form>
{{end}}