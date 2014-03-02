{{define "header"}}
<!doctype html>
<html>
<head>
<meta charset="utf-8">
<meta name="author" content="颜少" />
<meta name="Copyright" content="www.jacobcookie.com" />
<meta name="keywords" content="go,golang,beego,mysql,博客,开源" />
<meta name="description" content="隐匿是个人博客站点,记录成长中的思考,关于人生、学习、生活!" />
<!--禁止缓存 start-->
<meta http-equiv="pragma" content="no-cache"/>
<meta http-equiv="cache-control" content="no-cache"/>
<meta http-equiv="expires" content="0"/>
<!--禁止缓存 end-->
<title>{{.Title}}</title>
<link rel="shortcut icon" href="/static/img/favicon.ico">
<link href="/static/css/bootstrap.min.css" rel="stylesheet" >
<link href="/static/css/tag.css" rel="stylesheet" >
<script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>

<style>
body { 
	padding-top: 60px; background-color: #EEEEEE;
	font-family:Microsoft Yahei,Helvetica Neue,Helvetica,Arial,sans-serif;
	color:#777;font-size:14px;
	Scrollbar-Face-Color:#A4C6F9;
	Scrollbar-Highlight-Color:#A4C6F9;
	Scrollbar-Shadow-Color:#A4C6F9;
	Scrollbar-3Dlight-Color:8080FF;
	Scrollbar-Darkshadow-Color:8080FF;
	Scrollbar-Arrow-Color:FFffff;
	Scrollbar-Track-Color:E6E6FF;
	/* background-image: url(/static/img/bkg/3.jpg); */
}
body .container{
	background-color: #fff;
}
.todyaColor {
	BACKGROUND-COLOR: aqua
}
</style>
</head>
<body>
<!--background:#fff url(/static/img/backImg1.jpg) center repeat fixed!important;-->
{{end}}