
    $(window).on('load', 
    function() {
        $.ajax({
            url: "/api/url",
            type: "POST",
            success : function(response){
                if (response.meta.code == 200) {
                    if (response.data != null) {
                        var content = '';
                        for (var i = 0; i < response.data.URLMonitors.length; i++) {
                            if (response.data.URLMonitors[i].StatusOK == true) {
                                content = content + '<tr><td><li>'+ response.data.URLMonitors[i].URL +'</li></td><td><span class="number-green"></span></td></tr>';
                            }else{
                                content = content + '<tr><td><li>'+ response.data.URLMonitors[i].URL +'</li></td><td><span class="number-red"></span></td></tr>';
                            }
                        }
                        $('.url-health').append(content);
                    }
                } else {
                    console.log("error : " + response.meta.message)
                }
            }
        });
    });





    function loadAllURL() {
        $.ajax({
            url: "/api/url",
            type: "POST",
            success : function(response){
                $('.url-health').empty();
                if (response.meta.code == 200) {
                    if (response.data != null) {
                        var content = '';
                        for (var i = 0; i < response.data.URLMonitors.length; i++) {
                            if (response.data.URLMonitors[i].StatusOK == true) {
                                content = content + '<tr><td><li>'+ response.data.URLMonitors[i].URL +'</li></td><td><span class="number-green"></span></td></tr>';
                            }else{
                                content = content + '<tr><td><li>'+ response.data.URLMonitors[i].URL +'</li></td><td><span class="number-red"></span></td></tr>';
                            }
                        }
                        $(".error-message").empty();
                        $('.url-health').append(content);
                    }
                } else {
                    console.log("error : " + response.meta.message)
                    var htmlEM = '<b  style="color:red">'+em+'</b>'
                    $('.error-message').append(htmlEM);
                }
            }
        });
    }

    function addURL(){
        $(".error-message").empty();
         var urlValue = $("#inputURL").val();
         if (urlValue == '') {
            var em = "url input is required";
            var htmlEM = '<b  style="color:red">'+em+'</b>'
            $('.error-message').append(htmlEM);
         }else{
            var myData = {URL:urlValue};
            $.ajax({
                url: "/api/url/add",
                type: "POST",
                data: JSON.stringify(myData),
                contentType: "application/json",
                success : function(response){
                    if (response.meta.code == 200) {
                        console.log("success");
                        loadAllURL();
                    } else {
                        console.log("error : " + response.meta.message)
                        var htmlEM = '<b  style="color:red">'+response.meta.message+'</b>'
                        $('.error-message').append(htmlEM);
                    }
                }
            });
         }
    }

