/**
 * Created by myml on 17-6-23.
 */
app=new Vue({
    el: '#main',
    data: {
        vss:{},
        selectVs:"",
        subs:[],
        selectSub:"",
        vlist:{},
        vids:[],
        downURLs:[],
        vlistload:false,
        inited:false,
    },
    methods:{
        init:function () {
            $.get("host",function (data) {
                app.vss=data
                app.sub(data[0].id)
            })
        },
        sub:function (vs) {
            app.vlistload=true
            app.subs=[]
            app.selectVs=vs
            $.get("host/"+vs,function (data) {
                app.subs=data
                app.list(data[0].id)
                app.inited=true
            })
        },
        list:function(sub){
            app.vlistload=true
            if(sub!=undefined)
                app.selectSub=sub
            else
                sub=app.selectSub
            $.get("host/"+app.selectVs+"/sub/"+app.selectSub,function (data) {
                app.vlist={}
                app.vids=[]
                for(i in data){
                    app.vlist[data[i].videoID]=data[i]
                    app.vids.push(data[i].videoID)
                    app.vlistload=false
                }
            })
        },
        fanSelect:function () {
            for(i in this.vlist){
                vid=this.vlist[i].videoID
                indexOf=this.vids.indexOf(vid)
                if(indexOf!=-1){
                    this.vids.splice(indexOf,1)
                }else{
                    this.vids.push(vid)
                }
            }
        },
        play:function (vid) {
            $.get("host/"+app.selectVs+"/video/"+vid,function (data) {
                video=document.getElementById('video')
                video.src="video?src="+encodeURIComponent(data.src)
                video.play()
                $('#play').modal('open')
            })
        },
        stop:function () {
            document.getElementById('video').pause()
        },
        down:function () {
            app.downURLs=[]
            $('#down').modal('open')
            for(i in app.vids){
                var vid=app.vids[i]
                var name=app.vlist[vid].title;
                (function (vid,name) {
                    $.get("host/"+app.selectVs+"/video/"+vid,function (data) {
                        app.downURLs.push(location.href+"video?src="+encodeURIComponent(data.src)+"&name="+encodeURIComponent(name))
                    })
                })(vid,name)
            }
        },
        downAll:function () {
            for(i in this.downURLs){
				setTimeout("window.open('"+this.downURLs[i]+"')",1000*i)
            }
        }
    },
})
$('.modal').modal({
    complete:function () {
        app.stop()
    }
})
app.init()