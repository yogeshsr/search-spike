$(function(){
    $('#autocomplete').autocomplete({
        minChars: 3,
        serviceUrl: "http://localhost:8080/",
        dataType: "json",
        paramName: "search",
        crossDomain: true,
        ajaxSettings: {async:false},
        transformResult: function(data, originalQuery) {
            var suggestions = []
            if (!data) return {suggestions};

            for (var i=0; i<data.length; i++) {
                suggestions.push({ "value": data[i].title,
                    "data": data[i].sponsor})
            }

            return {suggestions};
        },
        onSelect: function (suggestion) {
            var thehtml = '<strong>Sponsor:</strong> ' + suggestion.data + ' <br> <strong>Voucher:</strong> ' + suggestion.value;
            $('#outputcontent').html(thehtml);
        }
    });

});