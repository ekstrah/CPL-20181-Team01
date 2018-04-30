from django.contrib import admin

from dm_system.models import SingleRow, SingleImage, SingleImageId
# Register your models here.

admin.site.register(SingleRow)
admin.site.register(SingleImage)
admin.site.register(SingleImageId)