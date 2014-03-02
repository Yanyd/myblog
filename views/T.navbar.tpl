{{define "navbar"}}
<div id="intro" style="height:70px;padding-top:20px;background-color:#2C2C2C;">
  <div class="container"><small style="opacity:0.7;color:#dfdfdf;font-size:18px;font-family:Merriweather,serif;">芝兰生于深谷，不以无人而不芳 。君子修身养德，不以穷困而改志。</small>
    <audio id="player"  controls autobuffer class="pull-right">浏览器不支持audio标签.</audio>
  </div>
</div>
<div id ="nav_top" class="navbar navbar-default bk_color">
  <div class="container">
    <div>
      <a class="navbar-brand" href="/">
        <span class="glyphicon glyphicon-home"></span>
        我的博客
      </a>
      <ul class="nav navbar-nav">
        <li {{if .IsIndex}}class="active"{{end}}>
          <a href="/">首页</a>
        </li>
        <li {{if .IsCategory}}class="active"{{end}}>
          <a href="/category">分类</a>
        </li>
        <li {{if .IsTopic}}class="active"{{end}}>
          <a href="/topic">博客</a>
        </li>
        <li {{if .IsPhoto}}class="active"{{end}}>
          <a href="/photo">相册</a>
        </li>
        <li {{if .IsBoard}}class="active"{{end}}>
          <a href="/board">留言</a>
        </li>
        <li {{if .IsTodo}}class="active"{{end}}>
          <a href="/todo">Todo</a>
        </li>
        <li {{if .IsResume}}class="active"{{end}}>
          <a href="/resume">简历</a>
        </li>
      </ul>
    </div>
    <div class="pull-right">
      <ul class="nav navbar-nav">
        {{if .IsLogin}}
        <li>
          <a href="###" style="cursor:default;">您好:{{.Account}}</a>
        </li>
        <li>
          <a href="javascript:window.location='/login?exit=true&time='+new Date().getTime();">退出</a>
        </li>
        {{else}}
        <li>
          <a href="/login">登录</a>
        </li>
        {{end}}
      </ul>
    </div>
  </div>
</div>
{{end}}