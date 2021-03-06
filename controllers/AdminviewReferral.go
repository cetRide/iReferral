package controllers

import (
	"iReferral/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Referrals struct {
	HudumaNo string
	Service  string
	// Hosname string
	RDate    string
	RTime    string
	Refer_hos string
	Refer_by  string
	Email     string
	Phone     string
}

var Referral []Referrals
var Len int
var Start, End, HosCode string

type ViewReferralsController struct {
	MainController
}

func (this *ViewReferralsController) MyReferrals() {

	//get the id of the logged in user
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID == nil {
		this.Redirect("/management_authentication", 302)
		return
	}
	myId := userID.(string)
	this.viewReferrals("viewreferrals")
		if this.Ctx.Input.Method() == "POST" {
		start := this.GetString("start")
		end := this.GetString("end")
		o := orm.NewOrm()
		o.Using("default")
		me := models.Hospital_account{MgnId: myId}
		err := o.Read(&me, "MgnId")
		if err != nil {
			fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
			flash.Store(&this.Controller)
			return
		}
		mycode := me.Name
		if Len > 0 {
			Referral = nil
		}
		o.Raw("SELECT huduma_no, service, r_date, r_time, refer_hos, refer_by, email, phone FROM referrals WHERE hos_name=? AND r_date BETWEEN ? AND ? ORDER BY r_date", mycode, start, end).QueryRows(&Referral)
		len := len(Referral)
		Len = len
		Start = start
		End = end
		HosCode = me.Code
		this.Redirect("/viewreferrals", 302)

	}
}

// func (this *ViewReferralsController) PatReferrals(){
// 	session := this.StartSession()
// 	userID := session.Get("UserID")
// 	if userID == nil {
// 		this.Redirect("/", 302)
// 		return
// 	}
// 	myId := userID.(string)
// 	this.viewReferrals("viewreferrals")
// 		if this.Ctx.Input.Method() == "POST" {
// 		start := this.GetString("start")
// 		end := this.GetString("end")
// 		o := orm.NewOrm()
// 		o.Using("default")
// 		// me := models.Patient_account{HudumaNo: myId}
// 		// err := o.Read(&me, "HudumaNo")
// 		// if err != nil {
// 		// 	fmt.Println("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
// 		// 	flash.Error("Internal server error - Sorry but we're unable to process your request at the moment. Please try later or contact support.")
// 		// 	flash.Store(&this.Controller)
// 		// 	return
// 		// }
// 		// mycode := me.Name
// 		if Len > 0 {
// 			Referral = nil
// 		}
// 		o.Raw("SELECT service, r_date, r_time, refer_hos, hos_name FROM referrals WHERE huduma_no=? AND r_date BETWEEN ? AND ? ORDER BY r_date", myId, start, end).QueryRows(&Referral)
// 		len := len(Referral)
// 		Len = len
// 		// Start = start
// 		// End = end
// 		// HosCode = me.Code
// 		this.Redirect("/viewreferrals", 302)

// 	}
// }
