from django.shortcuts import render
from .models import SingleRow, SingleImage

# Create your views here.
def index(request):
    return render(request, 'index.html')

def allrows(request):
    rows = SingleRow.objects.order_by('id')
    context = {'allrows' : rows}
    return render(request, 'dm_system/allrows.html', context)

def dgm(request):
    image = SingleImage.objects.order_by('id')[0]
    args = {'singleimage' : image}
    return render(request, 'dm_system/dgm.html',args)