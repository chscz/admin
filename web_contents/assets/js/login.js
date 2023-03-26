const login_btn = document.getElementById("btn_login");
const input_id = document.getElementById("user_id");
const input_pw = document.getElementById("user_pw");
const form_login = document.getElementById("form_login");

console.log(login_btn);
function test() {
   alert("login!!!!!");
}

function handleForm() {
   console.log("dd");
   fetch("/user")
      // .then((response) => response.json())
      .then((response) => console.log(response))
      .then((data) => console.log(data));
}

console.log("func test");
login_btn.addEventListener("click", test);

input_id.addEventListener("submit", test);
input_pw.addEventListener("submit", test);
// form_login.addEventListener("submit", handleForm);
console.log("add listener");

function keyPress(event) {
   console.log("엔터 눌");
   console.log(event);
   if (event.keyCode === 13) {
      console.log("엔터 림");
      handleForm();
   }
}
