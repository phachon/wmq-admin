/**
 * 实例
 */
var Service = {

    bindFancyBox: function() {

        $('[name="source_info"]').each(function () {
            $(this).fancybox({
                minWidth: 700,
                minHeight: 460,
                width: '70%',
                height: '50%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('[name="edit"]').each(function () {
            $(this).fancybox({
                minWidth: 700,
                minHeight: 550,
                width: '70%',
                height: '50%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });

        $('[name="service_config"]').each(function () {
            $(this).fancybox({
                minWidth: 800,
                minHeight: 560,
                width: '80%',
                height: '50%',
                autoSize: false,
                type: 'iframe',
                href: $(this).attr('data-link')
            });
        });
    },

    Config: {

        column: function (defaultColumns) {
	        //默认选择
	        if(defaultColumns) {
				for(var i = 0; i < defaultColumns.length; i++) {
					var column = defaultColumns[i].column;
					var columnAttr = defaultColumns[i].column_attr;
					var isIdColumn = defaultColumns[i].is_id_column;
					$("input[name='column_select'][value='"+column+"']").prop('checked', true);
					$("input[name='column_select'][value='"+column+"']").attr('checked', true);
					$("select[data-target='"+column+"']").val(columnAttr);
					if(isIdColumn == 1) {
						$("input[name='is_id_column'][data-target='"+column+"']").attr('checked', true);
					}
					$("input[name='sql_condition']").val(defaultColumns[i].sql_condition);
					// console.log(defaultColumns[i])
				}
	        }
            var submitUrl = $('form[name="column_form"]').attr('action');
            var allSelect = $('input[name="column_all_select"]');
            var columnSelect = $('input[name="column_select"]');
            var submitButton = $('button[name="column_submit"]');

            //全选
            allSelect.bind('click', function () {
                var checked = $(this).is(':checked');
                if(checked) {
                    columnSelect.prop('checked', true);
                    columnSelect.attr('checked', true);
                }else {
                    columnSelect.prop('checked', false);
                    columnSelect.attr('checked', false);
                    columnSelect.removeAttr('checked');
                }
            });

            //提交
            submitButton.bind('click', function () {
                var sqlCondition = $("input[name='sql_condition']").val();
                var serviceId = $("input[name='service_id']").val();
                var columnsData = [];
                columnSelect.each(function () {
                    var checked = $(this).is(':checked');
                    if(checked) {
                        var column = $(this).val();
                        var columnAttr = $("select[data-target='"+column+"']").find("option:selected").val();
                        var isIdCheck = $("input[name='is_id_column'][data-target='"+column+"']").is(':checked');
                        var isIdColumn = isIdCheck ? 1 : 0;
                        columnsData.push({column:column, column_attr: columnAttr, is_id_column: isIdColumn});
                    }
                });
                console.log(columnsData);

                $.ajax({
                    type : 'post',
                    url : submitUrl,
                    data : {columns_data: columnsData, sql_condition: sqlCondition, service_id: serviceId},
                    dataType: "json",
                    success : function(response) {
                        var message = response.message;
                        if(response.code == 0) {
                            var title = "<strong>操作失败：</strong>";
                            var text = message.split("\n");
                            var ul = '<ul>';
                            for(var i = 0; i < text.length; i++) {
                                ul +='<li>'+ text[i] +'</li>';
                            }
                            ul += '</ul>';
                            submitButton.notify(title + ul, {
                                position: "right",
                                className: 'error'
                            })
                        } else {
                            title = '<strong>操作成功：</strong>';
                            submitButton.notify(title + message, {
                                position: "right",
                                className: 'success'
                            })
                        }
                        if(response.redirect) {
                            setTimeout(function() {
                                location.href = response.redirect;
                            }, 2000);
                        }
                    },
                    error : function(response) {
                        console.log(response)
                    }
                });
            });
        },

        indexer: function () {

        },

        searchd: function () {
            
        }
    }
};