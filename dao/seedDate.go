package dao

import (
	"backend/dao/database"
	"backend/model"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

// SeedData 向数据库中添加初始数据
func SeedData(c *gin.Context) {
	database.DB.AutoMigrate(model.User{}, model.Book{}, model.Borrow{})
	// 创建图书示例s
	books := []model.Book{
		{
			Barcode:            "100001",
			Title:              "Go语言高级编程",
			Classification:     "TP312",
			ClassificationName: utils.GetClassificationName("TP312"),
			Author:             "柴树衫，曹春辉",
			Publisher:          "人民邮电出版社",
			Price:              "89.00",
			ISBN:               "9787115510365",
			Info:               "本书从实践出发讲解Go语言的进阶知识。本书共6章，第1章简单回顾Go语言的发展历史；第2章和第3章系统地介绍CGO编程和Go汇编语言的用法；第4章对RPC和Protobuf技术进行深入介绍，并讲述如何打造一个自己的RPC系统；第5章介绍工业级环境的Web系统的设计和相关技术；第6章介绍Go语言在分布式领域的一些编程技术。书中还涉及CGO和汇编方面的知识，其中CGO能够帮助读者继承的软件遗产，而在深入学习Go运行时，汇编对于理解各种语法设计的底层实现是必不可少的知识。此外，本书还包含一些紧跟潮流的内容，介绍开源界流行的gRPC及其相关应用，讲述Go Web框架中的基本实现原理和大型Web项目中的技术要点，引导读者对Go语言进行更深入的应用。 本书适合对Go语言的应用已经有一些心得，并希望能够深入理解底层实现原理或者是希望能够在Web开发方面结合Go语言来实现进阶学习的技术人员学习和参考。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202004/9787115510365.jpg",
		},
		{
			Barcode:            "100002",
			Title:              "Python数据分析快速上手",
			Classification:     "TP311.561",
			ClassificationName: utils.GetClassificationName("TP311.561"),
			Author:             "王靖、商艳红、张洪波、卢军",
			Publisher:          "清华大学出版社",
			Price:              "89.00",
			ISBN:               "9787302651758",
			Info:               "《《Python数据分析快速上手》通过通俗易懂的语言、丰富多彩的实例，详细介绍了使用Python进行数据分析应该掌握的各方面技术。本书内容包括Python基础，用NumPy进行数据计算，用Pandas进行数据分析，用SciPy进行数据分析，用Scikit-learn进行数据分析、数据预处理、数据可视化，用Matplotlib进行可视化等内容。本书示例丰富，所有涉及的程序代码都给出了详细的注释，读者可以轻松学习，快速提升开发技能。除此之外，本书还附配了教学视频、PPT课件和全书示例源码。《Python数据分析快速上手》适合数据分析的初学者、职场人士和所有对数据分析感兴趣的人员阅读，也适合作为大中专院校相关专业的教学用书。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202402/9787302651758.jpg",
		},
		{
			Barcode:            "100003",
			Title:              "零基础学Python编程一本通",
			Classification:     "TP311.561 ",
			ClassificationName: utils.GetClassificationName("TTP311.561"),
			Author:             "刘雅琼、何公甫、邹荣陞、李宗泽、何鑫",
			Publisher:          "化学工业出版社",
			Price:              "59",
			ISBN:               "9787122421616",
			Info:               "本书通过全彩图解+视频讲解的形式，介绍了Python编程入门及应用的相关知识，主要内容包括：Python编程环境安装与运行、Python中的数字运算、Python中的数据类型、输入输出与文件操作、条件与循环语句、函数与库、Python的OS、Python的命名空间与生命周期，以及Python五子棋项目实例、Python实现简易计算器、Python嵌入式实例—机器视觉等综合案例的开发。本书内容循序渐进，讲解通俗易懂，书中重难点章节配套视频讲解，扫码即可随时观看，同时提供源程序，方便学习实践。本书适合Python初学者、热爱编程的青少年朋友自学使用，也适合中小学信息技术课堂或相关培训机构用作教材。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202310/9787122421616.jpg",
		},
		{
			Barcode:            "100004",
			Title:              "近代中国开端的洞察",
			Classification:     "A164",
			ClassificationName: utils.GetClassificationName("A164"),
			Author:             "陈培永、喻春曦",
			Publisher:          "广东人民出版社",
			Price:              "26",
			ISBN:               "9787218167749",
			Info:               "本书深入剖析《马克思恩格斯论中国》中关于中国问题的文章和书信，以小品文的写作风格、通俗易懂的语言，生动呈现马克思恩格斯的中国观、对中国的总体看法与未来希冀，全书融思想性、理论性、可读性于一体，为深化马克思主义经典研究提供学术借鉴，同时有利于为中国特色社会主义实践建设提供理论参考、思想借鉴与现实指引，对推动马克思主义中国化时代化有着重要意义。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202312/9787218167749.jpg",
		},
		{
			Barcode:            "100005",
			Title:              "教子真有道",
			Classification:     "G78",
			ClassificationName: utils.GetClassificationName("G78"),
			Author:             "于含冰",
			Publisher:          "中华工商联合出版社",
			Price:              "58",
			ISBN:               "9787515837703",
			Info:               "本书以问题为导向，论证了孩子的问题本质上都是父母的问题。同时，从人生终极意义和时代发展的哲学高度，论证了家庭教育和学校教育的关系;从孩子生命具有无限可能的这个特性出发，基于《道德经》道法自然的哲学观，建立了开创性的理论系统和教育系统。提出了如何成为优秀父母的树理论，建立了父母教育课程系统，使父母在家庭教育中有章可循，不再迷茫，少走弯路。提出了如何润化孩子生命的全素质教育理论，明确指引父母在孩子小学毕业之前，一定要给孩子全素质营养，多元尝试多元探索、多元实践，籍此开发大脑，拓宽视野，增长成功经验，并尽早地发现孩子生命潜能的优势方向，给孩子早定位、早规划。提出了如何培养孩子学习动力的生根教育理论，确立了父母和孩子同学共修以主流文化为核心的[3-5-7学动力教育系统]特别强调用红色文化和传统经典文化的高级养料给孩子生根，使孩子产生强大的向上生长的动力。君子务本，本立而道生。本书是父母培养优秀孩子的道法、术，一书在手，父母无忧。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202312/9787515837703.jpg",
		},
		{
			Barcode:            "100006",
			Title:              "大变局下的中国管理3：商学院批判与自我革新",
			Classification:     "F123",
			ClassificationName: utils.GetClassificationName("F123"),
			Author:             "赵向阳",
			Publisher:          "中国人民大学出版社",
			Price:              "89.00",
			ISBN:               "9787300322698",
			Info:               "《在过去的二十多年里，中国管理学界群体性地走上了一条重视定量实证研究、英文论文发表、课题申请和争取人才“帽子”的道路，与此同时，对中国管理实践的实际影响乏善可陈，这些问题引起了各界人士关于中国商学院未来应该走向何方的讨论。在本书中，赵向阳博士针对以上弊端进行了比较深入的批判和反思，并且介绍了过去十来年自己的一些探索。这种反思与批判可能不是成体系的，不是特别完整，但是展现了难能可贵的真诚，以及在孤独中特立独行的勇气。本书适合所有对管理学和商学院感兴趣的读者阅读。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202311/9787300322698.jpg",
		},
		{
			Barcode:            "100007",
			Title:              "大变局下的中国管理3：商学院批判与自我革新",
			Classification:     "F123",
			ClassificationName: utils.GetClassificationName("F123"),
			Author:             "赵向阳",
			Publisher:          "中国人民大学出版社",
			Price:              "89.00",
			ISBN:               "9787300322698",
			Info:               "《在过去的二十多年里，中国管理学界群体性地走上了一条重视定量实证研究、英文论文发表、课题申请和争取人才“帽子”的道路，与此同时，对中国管理实践的实际影响乏善可陈，这些问题引起了各界人士关于中国商学院未来应该走向何方的讨论。在本书中，赵向阳博士针对以上弊端进行了比较深入的批判和反思，并且介绍了过去十来年自己的一些探索。这种反思与批判可能不是成体系的，不是特别完整，但是展现了难能可贵的真诚，以及在孤独中特立独行的勇气。本书适合所有对管理学和商学院感兴趣的读者阅读。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202311/9787300322698.jpg",
		},
		{
			Barcode:            "100008",
			Title:              "大变局下的中国管理3：商学院批判与自我革新",
			Classification:     "F123",
			ClassificationName: utils.GetClassificationName("F123"),
			Author:             "赵向阳",
			Publisher:          "中国人民大学出版社",
			Price:              "89.00",
			ISBN:               "9787300322698",
			Info:               "《在过去的二十多年里，中国管理学界群体性地走上了一条重视定量实证研究、英文论文发表、课题申请和争取人才“帽子”的道路，与此同时，对中国管理实践的实际影响乏善可陈，这些问题引起了各界人士关于中国商学院未来应该走向何方的讨论。在本书中，赵向阳博士针对以上弊端进行了比较深入的批判和反思，并且介绍了过去十来年自己的一些探索。这种反思与批判可能不是成体系的，不是特别完整，但是展现了难能可贵的真诚，以及在孤独中特立独行的勇气。本书适合所有对管理学和商学院感兴趣的读者阅读。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202311/9787300322698.jpg",
		},
		{
			Barcode:            "100009",
			Title:              "教子真有道",
			Classification:     "G78",
			ClassificationName: utils.GetClassificationName("G78"),
			Author:             "于含冰",
			Publisher:          "中华工商联合出版社",
			Price:              "58",
			ISBN:               "9787515837703",
			Info:               "本书以问题为导向，论证了孩子的问题本质上都是父母的问题。同时，从人生终极意义和时代发展的哲学高度，论证了家庭教育和学校教育的关系;从孩子生命具有无限可能的这个特性出发，基于《道德经》道法自然的哲学观，建立了开创性的理论系统和教育系统。提出了如何成为优秀父母的树理论，建立了父母教育课程系统，使父母在家庭教育中有章可循，不再迷茫，少走弯路。提出了如何润化孩子生命的全素质教育理论，明确指引父母在孩子小学毕业之前，一定要给孩子全素质营养，多元尝试多元探索、多元实践，籍此开发大脑，拓宽视野，增长成功经验，并尽早地发现孩子生命潜能的优势方向，给孩子早定位、早规划。提出了如何培养孩子学习动力的生根教育理论，确立了父母和孩子同学共修以主流文化为核心的[3-5-7学动力教育系统]特别强调用红色文化和传统经典文化的高级养料给孩子生根，使孩子产生强大的向上生长的动力。君子务本，本立而道生。本书是父母培养优秀孩子的道法、术，一书在手，父母无忧。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202312/9787515837703.jpg",
		},
		{
			Barcode:            "100010",
			Title:              "教子真有道",
			Classification:     "G78",
			ClassificationName: utils.GetClassificationName("G78"),
			Author:             "于含冰",
			Publisher:          "中华工商联合出版社",
			Price:              "58",
			ISBN:               "9787515837703",
			Info:               "本书以问题为导向，论证了孩子的问题本质上都是父母的问题。同时，从人生终极意义和时代发展的哲学高度，论证了家庭教育和学校教育的关系;从孩子生命具有无限可能的这个特性出发，基于《道德经》道法自然的哲学观，建立了开创性的理论系统和教育系统。提出了如何成为优秀父母的树理论，建立了父母教育课程系统，使父母在家庭教育中有章可循，不再迷茫，少走弯路。提出了如何润化孩子生命的全素质教育理论，明确指引父母在孩子小学毕业之前，一定要给孩子全素质营养，多元尝试多元探索、多元实践，籍此开发大脑，拓宽视野，增长成功经验，并尽早地发现孩子生命潜能的优势方向，给孩子早定位、早规划。提出了如何培养孩子学习动力的生根教育理论，确立了父母和孩子同学共修以主流文化为核心的[3-5-7学动力教育系统]特别强调用红色文化和传统经典文化的高级养料给孩子生根，使孩子产生强大的向上生长的动力。君子务本，本立而道生。本书是父母培养优秀孩子的道法、术，一书在手，父母无忧。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202312/9787515837703.jpg",
		},
		{
			Barcode:            "100011",
			Title:              "Python数据分析快速上手",
			Classification:     "TP311.561",
			ClassificationName: utils.GetClassificationName("TP311.561"),
			Author:             "王靖、商艳红、张洪波、卢军",
			Publisher:          "清华大学出版社",
			Price:              "89.00",
			ISBN:               "9787302651758",
			Info:               "《《Python数据分析快速上手》通过通俗易懂的语言、丰富多彩的实例，详细介绍了使用Python进行数据分析应该掌握的各方面技术。本书内容包括Python基础，用NumPy进行数据计算，用Pandas进行数据分析，用SciPy进行数据分析，用Scikit-learn进行数据分析、数据预处理、数据可视化，用Matplotlib进行可视化等内容。本书示例丰富，所有涉及的程序代码都给出了详细的注释，读者可以轻松学习，快速提升开发技能。除此之外，本书还附配了教学视频、PPT课件和全书示例源码。《Python数据分析快速上手》适合数据分析的初学者、职场人士和所有对数据分析感兴趣的人员阅读，也适合作为大中专院校相关专业的教学用书。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202402/9787302651758.jpg",
		},
		{
			Barcode:            "100012",
			Title:              "Go语言高级编程",
			Classification:     "TP312",
			ClassificationName: utils.GetClassificationName("TP312"),
			Author:             "柴树衫，曹春辉",
			Publisher:          "人民邮电出版社",
			Price:              "89.00",
			ISBN:               "9787115510365",
			Info:               "本书从实践出发讲解Go语言的进阶知识。本书共6章，第1章简单回顾Go语言的发展历史；第2章和第3章系统地介绍CGO编程和Go汇编语言的用法；第4章对RPC和Protobuf技术进行深入介绍，并讲述如何打造一个自己的RPC系统；第5章介绍工业级环境的Web系统的设计和相关技术；第6章介绍Go语言在分布式领域的一些编程技术。书中还涉及CGO和汇编方面的知识，其中CGO能够帮助读者继承的软件遗产，而在深入学习Go运行时，汇编对于理解各种语法设计的底层实现是必不可少的知识。此外，本书还包含一些紧跟潮流的内容，介绍开源界流行的gRPC及其相关应用，讲述Go Web框架中的基本实现原理和大型Web项目中的技术要点，引导读者对Go语言进行更深入的应用。 本书适合对Go语言的应用已经有一些心得，并希望能够深入理解底层实现原理或者是希望能够在Web开发方面结合Go语言来实现进阶学习的技术人员学习和参考。",
			Status:             1,
			Bookimg:            "http://hbzx.tsxcfw.com/uploadfile/202004/9787115510365.jpg",
		},
	}

	for _, book := range books {
		// 检查数据库中是否存在相同的 Barcode
		var existingBook model.Book
		if err := database.DB.Where("barcode = ?", book.Barcode).First(&existingBook).Error; err != nil {
			// 如果不存在相同的 Barcode，则添加该图书
			database.DB.Create(&book)
		}
		// 如果存在相同的 Barcode，则直接跳过，不进行任何操作
	}

	hashedPassword, _ := utils.HashPassword("123456")

	// 创建管理员和读者账号
	users := []model.User{
		{
			UserName:    "zhangsan",
			Name:        "张三",
			PhoneNumber: "13733737241",
			IsAdmin:     0,
			Password:    hashedPassword, // 这里直接放明文密码，实际应用中应该是加密后的密码
		},
		{
			UserName:    "lisi",
			Name:        "李四",
			PhoneNumber: "13733677776",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "wangwu",
			Name:        "王五",
			PhoneNumber: "13733677775",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "baiyan",
			Name:        "白岩",
			PhoneNumber: "13733677774",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "xioaming",
			Name:        "小明",
			PhoneNumber: "13733677766",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "xiaohua",
			Name:        "小华",
			PhoneNumber: "13733677765",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "xioabai",
			Name:        "小白",
			PhoneNumber: "13733677764",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "xiaozhang",
			Name:        "小张",
			PhoneNumber: "13733677763",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "xiaoli",
			Name:        "小丽",
			PhoneNumber: "13733677754",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
		{
			UserName:    "xiaohu",
			Name:        "小胡",
			PhoneNumber: "13733677734",
			IsAdmin:     0,
			Password:    hashedPassword,
		},
	}

	for _, user := range users {
		// 检查数据库中是否存在相同的用户名
		var existingUser model.User
		if err := database.DB.Where("user_name = ?", user.UserName).First(&existingUser).Error; err != nil {
			// 如果不存在相同的用户名，则添加该用户
			database.DB.Create(&user)
		}
		// 如果存在相同的用户名，则直接跳过，不进行任何操作
	}

}
