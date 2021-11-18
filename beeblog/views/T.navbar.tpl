{{define "navbar"}}
<!--置顶-->
<a class="navbar-brand" href="/">我的博客</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .IsHome}}class="active" {{end}}><a href="/">首页</a></li>
        <li {{if .IsCategory}}class="active" {{end}}><a href="/category">分类</a></li>
        <li {{if .IsTopic}}class="active" {{end}}><a href="/topic">文章</a></li>
    </ul>
</div>

<!--登录按钮-->
<div class="pull-right"> <!--登录按钮位于页面靠右侧-->
    <ul class="nav navbar-nav">
        <!--判断：如果已经登录，则显示为“退出”；如果是未登录，则显示为“管理员登录”-->
        {{if .IsLogin}}
        <li><a href="/login?exit=true">退出</a></li>
        {{else}}
        <li><a href="/login">管理员登录</a></li>
        {{end}}
    </ul>
</div>
{{end}}