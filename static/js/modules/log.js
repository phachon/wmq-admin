var Log = {
	
	bindFancyBox: function() {
		$('button[name="log_download"]').each(function () {
			$(this).fancybox({
				minWidth: 600,
				minHeight: 420,
				width: '80%',
				height: '45%',
				autoSize: false,
				type: 'iframe',
				href: $(this).attr('data-link')
			});
		});
	}
};
