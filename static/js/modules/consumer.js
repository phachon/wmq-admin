/**
 * Consumer 消费
 */
var Consumer = {

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

    node: function (element) {

        var id = $(element).val();
        var text = $('option[value='+id+']').text();
        
        location.href='/consumer/list?node_id='+id;
    }
};