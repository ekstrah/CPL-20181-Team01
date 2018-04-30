from django.db import models
from django.db.models.signals import post_save
# Create your models here.

class SingleRow(models.Model):
    temperature = models.FloatField()
    humidity = models.FloatField()
    numberOfPeople = models.IntegerField()
    def __str__(self):
        return str(self.temperature) + " °C " + str(self.humidity) + " % " + str(self.numberOfPeople) + " people"


class SingleImage(models.Model):
    image = models.ImageField()

class SingleImageId(models.Model):
    imageId = models.IntegerField()

def create_single_id(sender, **kwargs):
    singleRows = SingleRow.objects.all()
    #there may be bugs:
    #-this method is being called before new data is saved via post (probably)
    #-this is doing almost nothing to calculate final value, only adds everything

    constructorImageId = 0
    for singleRow in singleRows:
        constructorImageId = constructorImageId +\
                             int(singleRow.temperature) +\
                             int(singleRow.humidity) +\
                             int(singleRow.numberOfPeople)
    maxIdObject = SingleImageId.objects.all().order_by("-id")[0]
    maxId = maxIdObject.id
    constructorImageId = (constructorImageId % maxId)+1
    singleId = SingleImageId.objects.create(imageId=constructorImageId)

post_save.connect(create_single_id, sender=SingleRow)