package controllers

import (
	"fmt"
	"iReferral/models"

	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"unicode/utf8"
	//"github.com/astaxie/beego/validation"
	)
type DocController struct {
	MainController
}
func(this *DocController) FindPatient(){
	this.doctor_portal("doctor")
			if this.Ctx.Input.Method() == "GET"{
			value := this.GetString("huduma")
			if utf8.RuneCountInString(value) != 0 {
			o := orm.NewOrm()
			o.Using("default")
			hos := models.Patient_account{HudumaNo: value}
			err := o.Read(&hos, "HudumaNo")
			if err == orm.ErrNoRows {
				fmt.Println(err)
				fmt.Println("incorrect huduma number")
				flash.Error("You've entered incorrect huduma number")
				flash.Store(&this.Controller)
				return
			} else if err != nil {
				fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
				flash.Store(&this.Controller)
				return
			} else if err == nil{
				fmt.Println("Successful searching patient")
				this.Redirect("/report", 302)
			}
			
		}
		
	}
}	
