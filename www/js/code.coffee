me = ""
showing = ""
lclock = 0

seenClock = (c) ->
    if c > lclock
        lclock = c
        console.log("lclock=" + lclock)
    return

listDocs = (data) ->
    ret = JSON.parse(data)
    if ret.Err != ""
        appendError(ret.Err)
        return
 
    docs = $("div#docs")
    docs.empty()

    if ret.Docs == null || ret.Docs.length == 0
        docs.append("No Documents")
        return

    ul = $('<div class="list-group-justified"/>')
    ret.Docs.reverse()

    for doc in ret.Docs
        #seenClock(trib.Clock)
        li = $('<a href="#" class="list-group-item"/>')
        li.append(
            '<span class="glyhpicon glyphicon-file"></span>' +
             doc + 
            '<span class="glyphicon glyphicon-chevron-right"></span>
             <span class="badge badge-primary">14</span>')

        li.find("a.author").click((ev)->
            ev.preventDefault()
            name = $(this).text()
            if name.length > 0 && name.indexOf('@') == 0
                name = name.substring(1)
            _showUser(name)
        )
        li.hover(((ev)->
            if me != ""
                $(this).find("a.retrib").show()
            return
        ), ((ev)->
            $(this).find("a.retrib").hide()
            return
        ))
        #retrib.click((->
            #msg = trib.Message
            #who = trib.User
            #return (ev) ->
            #    ev.preventDefault()
            #    _postRetrib('RT @' + who + ': ' + msg, who)
        #)())
        ul.append(li)
    docs.append(ul)

    return

appendError = (e) ->
    $("div#errors").show()
    $("div#errors").append('<div class="error">Error: ' +
        e + '</div>')

main = ->
    #$("form#adduser").submit(addUser)
    #$("form#post").submit(postTrib)

    $("div#errors").hide()
    #$("div#timeline").hide()

    #$("a#signin").click(signIn)
    #$("a#home").click(showHome)
    #$("a#signout").click(signOut)

    #$("form#post textarea").keydown(->
    #    setTimeout((-> countPostLength()), 1)
    #)
    # $("form#post textarea").keypress(->
        # setTimeout((-> countPostLength()), 1)
    # )
    # $("form#post textarea").keyup(countPostLength)
    # $("form#post textarea").change(countPostLength)

    listDocs()
    return

$(document).ready(main)

