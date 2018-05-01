from django.shortcuts import render
from .models import SingleRow, SingleImage, SingleImageId

# Create your views here.
def index(request):
    return render(request, 'dm_system/index.html')

def allrows(request):
    rows = SingleRow.objects.order_by('id')
    context = {'allrows' : rows}
    return render(request, 'dm_system/allrows.html', context)

def dgm(request):
    imageId = SingleImageId.objects.all().order_by("-id")[0].imageId
    imageId = imageId - 1
    image = SingleImage.objects.order_by('id')[imageId]
    args = {'singleimage' : image}
    return render(request, 'dm_system/dgm.html',args)