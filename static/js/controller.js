pphrase = "123"
app.controller("ChatArena", function($scope, $http){
	$scope.message_bucket = [];
	$arena_status = "No Message"
	$scope.get_messages = function() {
			$http.get("/recieve").then(function (response) {
			setTimeout(window.scrollTo(0, document.body.scrollHeight), 5)
			if (response.data.length !== $scope.message_bucket.length && response.data.length !==0) {
				// for (i=0; i<response.data.length; i++) {
				// 	response.data[i].message = encrypt.encode(response.data[i].message, pphrase);
				// }
				$scope.arena_status = ""
				$scope.message_bucket = response.data
			}
		})
	}
	window.scrollTo(0, document.body.scrollHeight)
	setInterval($scope.get_messages, 1000)
	$scope.name = "Anon"
	$scope.send_message = function() {
		$scope.time = Date()
		$scope.m_text
		if ($scope.m_text!== "" || $scope.m_text==undefined) {
			// $scope.m_text = encrypt.encode($scope.m_text, pphrase); 
			$http.get("send?m="+$scope.m_text+"&t="+$scope.time+"&n="+$scope.name).then(function() {
				$scope.m_text = ""
				console.log("sent")
				$scope.ping_sound()
			})
		}
	}

    $scope.ping_sound = function() {
        var audio = new Audio('static/audio/brute-force.mp3');
        audio.play();
    }
})	
	
var encrypt = {
	encode: function (s, k) {
		var enc = "";
		var str = "";
		// make sure that input is string
		str = s.toString();
		for (var i = 0; i < s.length; i++) {
			// create block
			var a = s.charCodeAt(i);
			// bitwise XOR
			var b = a ^ k;
			enc = enc + String.fromCharCode(b);
		}
		return enc;
	}
};



