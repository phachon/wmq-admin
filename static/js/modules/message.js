/**
 * message 消息
 */
var Message = {

    bindFancyBox: function() {

        $('[name="edit"]').each(function () {
            $(this).fancybox({
                minWidth: 500,
                minHeight: 400,
                width: '65%',
                height: '50%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('button[name="add_message"]').each(function () {
            $(this).fancybox({
                minWidth: 500,
                minHeight: 400,
                width: '65%',
                height: '50%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('[name="consumer"]').each(function () {
            $(this).fancybox({
                minWidth: 500,
                minHeight: 400,
                width: '65%',
                height: '50%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
    },

	/**
	 * node
     * @param element
     */
    node: function (element) {
        var id = $(element).val();
        var text = $('option[value='+id+']').text();
        location.href='/message/list?node_id='+id;
    },

	/**
	 * consumer
     * @param consumers
     */
    consumer: function (consumers) {

    }
};