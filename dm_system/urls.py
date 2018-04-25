from django.conf.urls import url
from . import views

urlpatterns = [
    #Home/index page
    url(r'^$', views.index, name='index'),

    #Show all data
    url(r'^allrows/$', views.allrows, name='allrows'),

    #Show random picture
    url(r'^dgm', views.dgm, name='dgm')
]