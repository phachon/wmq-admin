/**
 * 表单提交类
 * Copyright (c) 2017 panchao
 */
var Form = {

	inPopup : false,

	/**
	 * ajax 提交表单
	 * @param element
	 * @param inPopup
	 * @returns {boolean}
	 */
	ajaxSubmit: function (element, inPopup) {

		var submitButton = $("button[name='submit']");
		if(inPopup) {
			Form.inPopup = true;
		}

		function successAlert(messages, data) {
			// var text = messages.join("\n");
			var text = messages;
			var timer = 2000;
			swal({
				'title' : '操作成功',
				'text' : "<h4>"+text+"</h4>",
				'html' : true,
				'type' : 'success',
				'showConfirmButton' : false,
				'timer' : timer,
				'location' : null
			});
		}

		function successNotify(messages, data) {
			// var text = messages.join("\n");
			var title = '<strong>操作成功：</strong>';
			var text = messages;
			var timer = 2000;
			submitButton.notify(title + text, {
				position: "right",
				className: 'success'
			})
		}

		//错误弹出信息
		function failedAlert(errors, data) {
			var text = errors.join("\n");
			var timer = 2000;
			swal({
				'title' : '操作失败',
				'text' : "<h4>"+text+"</h4>",
				'html' : true,
				'type' : 'error',
				'showConfirmButton' : true
				// 'timer' : timer
			});
		}

		function failedNotify(errors, data) {
			var title = "<strong>操作失败：</strong>";
			var text = errors.split("\n");
			var ul = '<ul>';
			for(var i = 0; i < text.length; i++) {
				ul +='<li>'+ text[i] +'</li>';
			}
			ul += '</ul>';
			submitButton.notify(title + ul, {
				position: "right",
				className: 'error'
			})
		}

		//弹出信息
		function response(result) {
			if(result.code == 0) {
				failedNotify(result.message, result.data);
			}
			if(result.code == 1) {
				successNotify(result.message, result.data);
			}

			//如果设置了跳转
			if(result.redirect) {
				setTimeout(function() {
					if(Form.inPopup) {
						parent.location.href = result.redirect;
					} else {
						location.href = result.redirect;
					}
				}, 3000);
			}
		}

		var options = {
			dataType: 'json',
			success: response
		};

		$(element).ajaxSubmit(options);

		return false;
	}
};