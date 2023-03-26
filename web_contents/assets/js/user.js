const btn_get_user = document.getElementById("btn_get_user");

function GetUserList() {
   fetch("/user_list")
      .then((response) => console.log(response))
      .then((data) => console.log(data));
}

btn_get_user.addEventListener("click", GetUserList);
