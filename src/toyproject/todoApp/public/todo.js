(function ($) {
  'use strict';
  $(function () {
    var todoListItem = $('.todo-list');
    var todoListInput = $('.todo-list-input');
    $('.todo-list-add-btn').on('click', function (event) {
      event.preventDefault();

      var item = $(this).prevAll('.todo-list-input').val();
      // POST로 Item을 보낸 후 서버에서 값이 오면 addItem함수를 호출한다.
      if (item) {
        $.post('/todos', { name: item }, addItem);
        todoListInput.val('');
      }
    });

    var addItem = function (item) {
      if (item.completed) {
        todoListItem.append(
          "<li class='completed'" +
            " id='" +
            item.id +
            "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" +
            item.name +
            "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>"
        );
      } else {
        todoListItem.append(
          '<li ' +
            " id='" +
            item.id +
            "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" +
            item.name +
            "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>"
        );
      }
    };
    // GET /todos : todos 리스트를 불러오는 메서드
    $.get('/todos', function (items) {
      items.forEach((e) => {
        addItem(e);
      });
    });

    // GET /complete-todo/{id} : 체크박스의 상태를 전달하는 메서드
    todoListItem.on('change', '.checkbox', function () {
      var id = $(this).closest('li').attr('id');
      var $self = $(this);
      var complete = true;
      if ($(this).attr('checked')) {
        complete = false;
      }
      $.get('complete-todo/' + id + '?complete=' + complete, function (data) {
        if (complete) {
          $self.attr('checked', 'checked');
        } else {
          $self.removeAttr('checked');
        }

        $self.closest('li').toggleClass('completed');
      });
    });

    // DELETE todos/id todo 리스트 한가지 항목 삭제하는 메서드
    todoListItem.on('click', '.remove', function () {
      var id = $(this).closest('li').attr('id');
      var $self = $(this);
      $.ajax({
        url: '/todos/' + id,
        type: 'DELETE',
        success: function (data) {
          if (data.success) {
            $self.parent().remove();
          }
        },
      });
      console.log('?');
    });
  });
})(jQuery);
