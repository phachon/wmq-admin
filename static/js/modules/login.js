var Login = {

	errorMessage: "#errorMessage",

	/**
	 * 登录
	 * @param element
	 */
	ajaxSubmit: function (element) {

		var name = $("input[name='name']").val();
		var password = $("input[name='password']").val();
		if(!name || !password) {
			return false
		}


		function success(messages, data) {
			var text = "";
			var failedText = messages;
			$(Login.errorMessage).removeClass('alert-danger');
			$(Login.errorMessage).addClass('alert-success');
			$(Login.errorMessage).removeClass('hidden');
			$(Login.errorMessage + ' strong ').html(text + failedText);
		}

		function failed(messages, data) {
			var text = "登陆失败：";
			var failedText = messages;
			$(Login.errorMessage).removeClass('alert-success');
			$(Login.errorMessage).addClass('alert-danger');
			$(Login.errorMessage).removeClass('hidden');
			$(Login.errorMessage + ' strong ').html(text + failedText);
		}

		function response(result) {
			$(Login.errorMessage).addClass('hidden');
			if(result.code == 0) {
				failed(result.message, result.data);
			}
			if(result.code == 1) {
				success(result.message, result.data);
				if(result.redirect) {
					location.href = result.redirect;
					setTimeout(function() {
						location.reload();
					}, 2000);
				}
			}
		}

		var options = {
			dataType: 'json',
			success: response
		};

		$(element).ajaxSubmit(options);
	}
};