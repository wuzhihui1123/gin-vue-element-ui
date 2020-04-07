gin + vue + element-ui + swagger
===

### 背景
平常开发的都是一些内部系统，需要我们后端程序员同时编写前端页面和后端服务代码。以前一直用JQuery之类的传统工具，对HTML页面元素（DOM）的操作都需要
自己来编写前端代码实现。当页面数据和交互变得复杂时,这种开发模式效率比较低，后续的维护和扩展的成本也比较高。
由于浏览器环境下JavasScript本身的不足，开发和维护大型或复杂页面还是比较痛的。这几年，随着NodeJs技术的兴起，前端技术千变万化，
出现了很多新的语言(TypeScript, SCSS)，支持“数据驱动视图”思想的开发框架(Vue, React, AngularJs)以及编译打包工具（Webpack, Gulp)。
同时“前后端分离”的出现也影响了前后端开发人员的协作和前后端代码的部署方式。  
为了提高开发效率，降低前端页面的开发和维护成本，能够更充分使用前端的生态资源（eg: Element-UI)， 需要采用现代化的框架来开发前端页面。
Vue的社区非常活跃，生态非常丰富，国内使用率特别高，因此我们围绕VUE生态来进行选择

### 指导原则
- 前端使用现代化的 MVVM 框架，解放对HTML DOM的关注和操作。能够享受到：
  1. 模块化开发： 更高的扩展和维护性
  1. 数据驱动视图/虚拟DOM：解放劳动力，不再需要自己来关心和操作HTML DOM元素，不用担心数据和DOM之间的同步了
  1. 使用更高效的开发平台：包括语言层面（Typescript, SCSS等）/语法层面（ES6)/工具层面
  1. 使用现代化框架下的生态资源
  
- 前后端开发不分工（没有专门的前端同学），开发时不需要在多个工程之间来回切换。 方案：前后端代码放到同一工程的不同目录下

- 前后端代码统一部署：如果前后端独立部署的话，一是增加运维成本，二是需要解决Api调用的跨域问题。
方案：前端代码编译后放到后端服务的页面文件目录下，然后只需部署后端服务即可

- 不要让研发同学关注和学习代码的编译和部署。方案：把编译和部署代码化（使用Docker封装代码的编译，使用GitLab CI封装代码的部署）
  
  

- ### 前端代码目录：front-end
> 修改于： [vue-admin-template](https://github.com/PanJiaChen/vue-admin-template.git) 

- ### 后端代码目录：back-end
> 修改于： [go-gin-example](https://github.com/eddycjy/go-gin-example.git) 