/**
 * node 节点
 */
var Node = {

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
    }
};