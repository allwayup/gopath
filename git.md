# 长期保存密码
git config --global credential.helper store
# 不缓存密码
git config --global --unset credential.helper
# 禁用代理
git config --global --unset http.proxy
git config --global --unset https.proxy
# 忽略ssl认证
git config --global http.sslVerify false
git config --global https.sslVerify false
# 禁用安全目录
git config --global --add safe.directory "*"
# linux查询当前账号
who
# 更改文件夹所属账号
chown -R <username> <filepath>
# 登陆信息
git config --global user.name "allwayup"
git config --global user.email "allwayup@163.com"
# 生成密钥对
ssh-keygen -t rsa -C "allwayup@163.com"
# 检验密钥
ssh -T git@github.com
# 当前的remote url
git remote -v
# 设置当前工程remote url
git remote add origin https://<token>@github.com/allwayup/gopath.git
# 切换git源,访问令牌
git remote set-url origin https://<token>@github.com/allwayup/gopath.git
# 提交
git -c https.sslVerify=false -c http.sslVerify=false -c credential.helper= -c core.quotepath=false -c log.showSignature=false push --progress --porcelain origin refs/heads/main:main
