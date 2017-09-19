/**
 * Consumer 消费
 */
var Consumer = {

    bindFancyBox: function() {

        $('[name="edit"]').each(function () {
            
            $(this).fancybox({
                minWidth: 500,
                minHeight: 480,
                width: '65%',
                height: '57%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('button[name="add_consumer"]').each(function () {
            $(this).fancybox({
                minWidth: 500,
                minHeight: 480,
                width: '65%',
                height: '57%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
    },

    node: function (element) {
        var id = $(element).val();
        location.href='/consumer/list?node_id='+id;
    },

    //定时更新状态
    status: function (url) {

        url = '/consumer/status?node_id=1';

        var nowTime = Date.parse(new Date()) / 1000;
        $.ajax({
            type : 'post',
            url : url,
            data : {'arr':''},
            dataType: "json",
            success : function(response) {
                if(response.code == 1) {
                    var values = response.data;
                    for(var i = 0; i < values.length; i++) {
                        var element = $('#consumer_'+values[i].ID).find(".consumer_status");
                        if(values[i].LastTime != 0) {
                            element.html('<label class="text-success">running('+values[i].Count+')</label>');
                        }else {
                            element.html('<label class="text-danger">stoped('+values[i].Count+')</label>');
                        }
                    }
                } else {
                    console.log(response.message);
                }
            },
            error : function(response) {
                console.log(response);
            }
        });
    }
};