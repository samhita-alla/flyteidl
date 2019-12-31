# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import unittest

import flyteadmin
from flyteadmin.api.admin_service_api import AdminServiceApi  # noqa: E501
from flyteadmin.rest import ApiException


class TestAdminServiceApi(unittest.TestCase):
    """AdminServiceApi unit test stubs"""

    def setUp(self):
        self.api = flyteadmin.api.admin_service_api.AdminServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_create_execution(self):
        """Test case for create_execution

        """
        pass

    def test_create_launch_plan(self):
        """Test case for create_launch_plan

        """
        pass

    def test_create_node_event(self):
        """Test case for create_node_event

        """
        pass

    def test_create_task(self):
        """Test case for create_task

        """
        pass

    def test_create_task_event(self):
        """Test case for create_task_event

        """
        pass

    def test_create_workflow(self):
        """Test case for create_workflow

        """
        pass

    def test_create_workflow_event(self):
        """Test case for create_workflow_event

        """
        pass

    def test_delete_project_attributes(self):
        """Test case for delete_project_attributes

        """
        pass

    def test_delete_project_domain_attributes(self):
        """Test case for delete_project_domain_attributes

        """
        pass

    def test_delete_workflow_attributes(self):
        """Test case for delete_workflow_attributes

        """
        pass

    def test_get_active_launch_plan(self):
        """Test case for get_active_launch_plan

        """
        pass

    def test_get_execution(self):
        """Test case for get_execution

        """
        pass

    def test_get_execution_data(self):
        """Test case for get_execution_data

        """
        pass

    def test_get_launch_plan(self):
        """Test case for get_launch_plan

        """
        pass

    def test_get_named_entity(self):
        """Test case for get_named_entity

        """
        pass

    def test_get_node_execution(self):
        """Test case for get_node_execution

        """
        pass

    def test_get_node_execution_data(self):
        """Test case for get_node_execution_data

        """
        pass

    def test_get_project_attributes(self):
        """Test case for get_project_attributes

        """
        pass

    def test_get_project_domain_attributes(self):
        """Test case for get_project_domain_attributes

        """
        pass

    def test_get_task(self):
        """Test case for get_task

        """
        pass

    def test_get_task_execution(self):
        """Test case for get_task_execution

        """
        pass

    def test_get_task_execution_data(self):
        """Test case for get_task_execution_data

        """
        pass

    def test_get_workflow(self):
        """Test case for get_workflow

        """
        pass

    def test_get_workflow_attributes(self):
        """Test case for get_workflow_attributes

        """
        pass

    def test_list_active_launch_plans(self):
        """Test case for list_active_launch_plans

        """
        pass

    def test_list_executions(self):
        """Test case for list_executions

        """
        pass

    def test_list_launch_plan_ids(self):
        """Test case for list_launch_plan_ids

        """
        pass

    def test_list_launch_plans(self):
        """Test case for list_launch_plans

        """
        pass

    def test_list_launch_plans2(self):
        """Test case for list_launch_plans2

        """
        pass

    def test_list_named_entities(self):
        """Test case for list_named_entities

        """
        pass

    def test_list_node_executions(self):
        """Test case for list_node_executions

        """
        pass

    def test_list_node_executions_for_task(self):
        """Test case for list_node_executions_for_task

        """
        pass

    def test_list_projects(self):
        """Test case for list_projects

        """
        pass

    def test_list_task_executions(self):
        """Test case for list_task_executions

        """
        pass

    def test_list_task_ids(self):
        """Test case for list_task_ids

        """
        pass

    def test_list_tasks(self):
        """Test case for list_tasks

        """
        pass

    def test_list_tasks2(self):
        """Test case for list_tasks2

        """
        pass

    def test_list_workflow_ids(self):
        """Test case for list_workflow_ids

        """
        pass

    def test_list_workflows(self):
        """Test case for list_workflows

        """
        pass

    def test_list_workflows2(self):
        """Test case for list_workflows2

        """
        pass

    def test_register_project(self):
        """Test case for register_project

        """
        pass

    def test_relaunch_execution(self):
        """Test case for relaunch_execution

        """
        pass

    def test_terminate_execution(self):
        """Test case for terminate_execution

        """
        pass

    def test_update_launch_plan(self):
        """Test case for update_launch_plan

        """
        pass

    def test_update_named_entity(self):
        """Test case for update_named_entity

        """
        pass

    def test_update_project_attributes(self):
        """Test case for update_project_attributes

        """
        pass

    def test_update_project_domain_attributes(self):
        """Test case for update_project_domain_attributes

        """
        pass

    def test_update_workflow_attributes(self):
        """Test case for update_workflow_attributes

        """
        pass


if __name__ == '__main__':
    unittest.main()
