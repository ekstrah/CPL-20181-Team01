from tastypie.resources import ModelResource
from dm_system.models import SingleRow
from tastypie.authorization import Authorization

class SingleRowResource(ModelResource):
    class Meta:
        queryset = SingleRow.objects.all()
        resource_name = 'singlerow'
        authorization = Authorization()