/*
构建一个复杂结构,如大楼,首先是打牢地基,搭建架构,然后一层一层盖起来.
就是先构建成一个物体的各个部分,然后分阶段将他们组装起来.

生成器模式将复杂对象的构造与表示分开，以便相同的构建过程可以创建不同的表示形式。
*/
package main


func main(){
	b := NewBuilder()
	lamp_1 := b.Color(BlueColor).Brand(Osram).Build()
	lamp_1.Open()
	lamp_1.ProductionIllustrative()

	lamp_2 := b.Color(GreenColor).Brand(OppleBulb).Build()
	lamp_2.Open()
	lamp_2.ProductionIllustrative()
	}
