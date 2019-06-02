package routers

import (
	"iReferral/controllers"
	"github.com/astaxie/beego"
)
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("healthIssue", &controllers.MainController{}, "get,post:Referral")
	beego.Router("viewreferrals", &controllers.ViewReferralsController{}, "get,post:MyReferrals")
	beego.Router("confirmreferral", &controllers.MainController{}, "get,post:PostDetails")
	beego.Router("patientreferralform", &controllers.ReferralsController{}, "get,post:PatReferral")
		beego.Router("info/admin_regSuccess", &controllers.MainController{}, "get,post:Admin_Reg_Success")
	beego.Router("info/emp_regSuccess", &controllers.MainController{}, "get,post:Emp_Reg_Success")
	beego.Router("info/patient_regSuccess", &controllers.MainController{}, "get,post:Patient_Reg_Success")
		beego.Router("logout", &controllers.AccountController{}, "get,post:Logout")
	beego.Router("hosreg", &controllers.HosregController{}, "get,post:Create")
	beego.Router("checkvalidity", &controllers.HosregController{}, "get,post:CheckReg")
	beego.Router("facility_mgn", &controllers.FacilityController{}, "post:Createemp")
	beego.Router("facility_mgn", &controllers.FacilityController{}, "get:RemoveEmp")
	beego.Router("facilityCreated", &controllers.FacilityController{}, "get,post:Iscreated")
	beego.Router("hosearch", &controllers.FaController{}, "post:Createservice")
	beego.Router("hosearch", &controllers.FaController{}, "get:Removeservice")
	beego.Router("/doctor", &controllers.DocController{}, "get:FindPatient")
	beego.Router("/report", &controllers.DiagnosisController{}, "get,post:UpdateReport")
	beego.Router("/phome", &controllers.MainController{}, "get,post:Phome")
	beego.Router("/mymedicalrecords", &controllers.PatientController{}, "get,post:Preport")
	beego.Router("/myreferrals", &controllers.PatientController{}, "get,post:Preferral")
	beego.Router("/adminhome", &controllers.MainController{}, "get,post:AdminDash")
	beego.Router("unsuccessfulrequest", &controllers.MainController{}, "get,post:Unsuccessful")
	beego.Router("unsuccessfulservicerequest", &controllers.MainController{}, "get,post:UnsuccessfulHos")
	beego.Router("ourcontacts", &controllers.MainController{}, "get,post:Conta")
	beego.Router("referpatient", &controllers.ReferpController{})
	beego.Router("searchservice", &controllers.ServicesController{})
	beego.Router("patient_login", &controllers.PatientLogInController{})
	beego.Router("staff_login", &controllers.StaffLogInController{})
	beego.Router("management_login", &controllers.MgnLogInController{})
	beego.Router("management_signup", &controllers.MgnSignUpController{})
	beego.Router("patient_signup", &controllers.PatSignUpController{})
	beego.Router("staff_signup", &controllers.StaffSignUpController{})
	beego.Router("contact", &controllers.MainController{}, "get,post:Contactus")
	beego.Router("contact_us", &controllers.EmailController{})
	beego.Router("management_authentication", &controllers.MainController{}, "get,post:AdminAuth")
	beego.Router("services", &controllers.ServiceController{}, "get,post:Listservices")
	
}
