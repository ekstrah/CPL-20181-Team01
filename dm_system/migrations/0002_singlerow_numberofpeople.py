# Generated by Django 2.0.4 on 2018-04-19 02:19

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('dm_system', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='singlerow',
            name='numberOfPeople',
            field=models.IntegerField(default=0),
            preserve_default=False,
        ),
    ]
