# Generated by Django 2.0.4 on 2018-04-25 02:31

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('dm_system', '0006_auto_20180420_1203'),
    ]

    operations = [
        migrations.CreateModel(
            name='SingleImage',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('image', models.ImageField(upload_to='')),
            ],
        ),
        migrations.DeleteModel(
            name='Posting',
        ),
    ]
