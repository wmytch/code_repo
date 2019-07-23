# 用pod安装swiftyJson的一个实例

1. 在项目中创建一个文件Podfile。
2. Github的swiftyJson项目的readme文件中复制下面内容到Podfile中
~~~shell
platform :ios, '8.0'
use_frameworks!

target 'MyApp' do  #MyApp--你的项目的名字
    pod 'SwiftyJSON', '~> 4.0'
end
~~~
3. 运行`pod  update`,有可能这一步接下来会自动执行`pod install`
1. 运行`pod install`,如果上一步已经自动执行了`pod install`,会有提示的，就当这一步是个确认操作吧。
5. 接下来按照提示启动项目即可。