# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.core_catalog_cache_status import CoreCatalogCacheStatus  # noqa: F401,E501
from flyteadmin.models.core_catalog_metadata import CoreCatalogMetadata  # noqa: F401,E501
from flyteadmin.models.flyteidladmin_dynamic_workflow_node_metadata import FlyteidladminDynamicWorkflowNodeMetadata  # noqa: F401,E501


class FlyteidladminTaskNodeMetadata(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'cache_status': 'CoreCatalogCacheStatus',
        'catalog_key': 'CoreCatalogMetadata',
        'dynamic_workflow_node_metadata': 'FlyteidladminDynamicWorkflowNodeMetadata'
    }

    attribute_map = {
        'cache_status': 'cache_status',
        'catalog_key': 'catalog_key',
        'dynamic_workflow_node_metadata': 'dynamic_workflow_node_metadata'
    }

    def __init__(self, cache_status=None, catalog_key=None, dynamic_workflow_node_metadata=None):  # noqa: E501
        """FlyteidladminTaskNodeMetadata - a model defined in Swagger"""  # noqa: E501

        self._cache_status = None
        self._catalog_key = None
        self._dynamic_workflow_node_metadata = None
        self.discriminator = None

        if cache_status is not None:
            self.cache_status = cache_status
        if catalog_key is not None:
            self.catalog_key = catalog_key
        if dynamic_workflow_node_metadata is not None:
            self.dynamic_workflow_node_metadata = dynamic_workflow_node_metadata

    @property
    def cache_status(self):
        """Gets the cache_status of this FlyteidladminTaskNodeMetadata.  # noqa: E501

        Captures the status of caching for this execution.  # noqa: E501

        :return: The cache_status of this FlyteidladminTaskNodeMetadata.  # noqa: E501
        :rtype: CoreCatalogCacheStatus
        """
        return self._cache_status

    @cache_status.setter
    def cache_status(self, cache_status):
        """Sets the cache_status of this FlyteidladminTaskNodeMetadata.

        Captures the status of caching for this execution.  # noqa: E501

        :param cache_status: The cache_status of this FlyteidladminTaskNodeMetadata.  # noqa: E501
        :type: CoreCatalogCacheStatus
        """

        self._cache_status = cache_status

    @property
    def catalog_key(self):
        """Gets the catalog_key of this FlyteidladminTaskNodeMetadata.  # noqa: E501


        :return: The catalog_key of this FlyteidladminTaskNodeMetadata.  # noqa: E501
        :rtype: CoreCatalogMetadata
        """
        return self._catalog_key

    @catalog_key.setter
    def catalog_key(self, catalog_key):
        """Sets the catalog_key of this FlyteidladminTaskNodeMetadata.


        :param catalog_key: The catalog_key of this FlyteidladminTaskNodeMetadata.  # noqa: E501
        :type: CoreCatalogMetadata
        """

        self._catalog_key = catalog_key

    @property
    def dynamic_workflow_node_metadata(self):
        """Gets the dynamic_workflow_node_metadata of this FlyteidladminTaskNodeMetadata.  # noqa: E501

        In the case this task launched a dynamic workflow we capture its structure here.  # noqa: E501

        :return: The dynamic_workflow_node_metadata of this FlyteidladminTaskNodeMetadata.  # noqa: E501
        :rtype: FlyteidladminDynamicWorkflowNodeMetadata
        """
        return self._dynamic_workflow_node_metadata

    @dynamic_workflow_node_metadata.setter
    def dynamic_workflow_node_metadata(self, dynamic_workflow_node_metadata):
        """Sets the dynamic_workflow_node_metadata of this FlyteidladminTaskNodeMetadata.

        In the case this task launched a dynamic workflow we capture its structure here.  # noqa: E501

        :param dynamic_workflow_node_metadata: The dynamic_workflow_node_metadata of this FlyteidladminTaskNodeMetadata.  # noqa: E501
        :type: FlyteidladminDynamicWorkflowNodeMetadata
        """

        self._dynamic_workflow_node_metadata = dynamic_workflow_node_metadata

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(FlyteidladminTaskNodeMetadata, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, FlyteidladminTaskNodeMetadata):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
