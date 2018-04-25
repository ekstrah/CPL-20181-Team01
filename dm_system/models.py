from django.db import models

# Create your models here.

class SingleRow(models.Model):
    temperature = models.FloatField()
    humidity = models.FloatField()
    numberOfPeople = models.IntegerField()
    def __str__(self):
        return str(self.temperature) + " °C " + str(self.humidity) + " % " + str(self.numberOfPeople) + " people"


class SingleImage(models.Model):
    image = models.ImageField()