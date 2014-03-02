{{define "right"}}
<div class="col-md-3 col-md-offset-1" style="background-color:#fff;">
  <div class="ctg_section">
  <h3>文章分类</h3>

       {{range .Categories}}

       {{if eq .TopicCount 0 }}
        <a  class="btn btn-default disabled" style="display:inline-block;margin-bottom:10px;padding-right:6px;margin-right:4px;width:120px;color:#428bd1;">{{.Title}} <span class="badge pull-right" style="color:#777;background-color:#F7F7F7;">{{.TopicCount}}</span></a>

       {{else}}
        <a href="/?id={{.Id}}"  class="btn btn-default" style="display:inline-block;margin-bottom:10px;padding-right:6px;margin-right:4px;width:120px;color:#428bd1;">{{.Title}} <span class="badge pull-right" style="color:#777;background-color:#F7F7F7;">{{.TopicCount}}</span></a>
       {{end}}
    
       {{end}}
  </div>
  <div class="label_section">
    <h3>标 签 云</h3>
    <div id="tagsList">
      {{range .Labels}}
      <a href="/?label={{.}}" title="{{.}}" >{{.}}</a>
      {{end}}
    </div>
  </div>

  <div class="comment_section">
    <h3>最新评论</h3>
    <div id="new_comments">
      <div class="list-group" id="cmt_list">
      </div>
    </div>
  </div>
</div>
{{end}}