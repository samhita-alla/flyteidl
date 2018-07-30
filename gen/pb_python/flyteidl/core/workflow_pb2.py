# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/workflow.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from flyteidl.core import types_pb2 as flyteidl_dot_core_dot_types__pb2
from flyteidl.core import condition_pb2 as flyteidl_dot_core_dot_condition__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/workflow.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_pb=_b('\n\x1c\x66lyteidl/core/workflow.proto\x12\rflyteidl.core\x1a\x1d\x66lyteidl/core/interface.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x19\x66lyteidl/core/types.proto\x1a\x1d\x66lyteidl/core/condition.proto\x1a\x1egoogle/protobuf/duration.proto\"f\n\x07IfBlock\x12\x33\n\tcondition\x18\x01 \x01(\x0b\x32 .flyteidl.core.BooleanExpression\x12&\n\tthen_node\x18\x02 \x01(\x0b\x32\x13.flyteidl.core.Node\"\xb6\x01\n\x0bIfElseBlock\x12$\n\x04\x63\x61se\x18\x01 \x01(\x0b\x32\x16.flyteidl.core.IfBlock\x12%\n\x05other\x18\x02 \x03(\x0b\x32\x16.flyteidl.core.IfBlock\x12(\n\telse_node\x18\x03 \x01(\x0b\x32\x13.flyteidl.core.NodeH\x00\x12%\n\x05\x65rror\x18\x04 \x01(\x0b\x32\x14.flyteidl.core.ErrorH\x00\x42\t\n\x07\x64\x65\x66\x61ult\"9\n\nBranchNode\x12+\n\x07if_else\x18\x02 \x01(\x0b\x32\x1a.flyteidl.core.IfElseBlock\"/\n\x08TaskNode\x12\x16\n\x0creference_id\x18\x01 \x01(\tH\x00\x42\x0b\n\treference\"}\n\x0cWorkflowNode\x12\"\n\x18launch_plan_reference_id\x18\x01 \x01(\tH\x00\x12<\n\x11workflow_template\x18\x02 \x01(\x0b\x32\x1f.flyteidl.core.WorkflowTemplateH\x00\x42\x0b\n\treference\"w\n\x0cNodeMetadata\x12\x0c\n\x04name\x18\x01 \x01(\t\x12*\n\x07timeout\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12-\n\x07retries\x18\x05 \x01(\x0b\x32\x1c.flyteidl.core.RetryStrategy\"#\n\x05\x41lias\x12\x0b\n\x03var\x18\x01 \x01(\t\x12\r\n\x05\x61lias\x18\x02 \x01(\t\"\xd2\x02\n\x04Node\x12\n\n\x02id\x18\x01 \x01(\t\x12-\n\x08metadata\x18\x02 \x01(\x0b\x32\x1b.flyteidl.core.NodeMetadata\x12&\n\x06inputs\x18\x03 \x03(\x0b\x32\x16.flyteidl.core.Binding\x12\x19\n\x11upstream_node_ids\x18\x04 \x03(\t\x12,\n\x0eoutput_aliases\x18\x05 \x03(\x0b\x32\x14.flyteidl.core.Alias\x12,\n\ttask_node\x18\x06 \x01(\x0b\x32\x17.flyteidl.core.TaskNodeH\x00\x12\x34\n\rworkflow_node\x18\x07 \x01(\x0b\x32\x1b.flyteidl.core.WorkflowNodeH\x00\x12\x30\n\x0b\x62ranch_node\x18\x08 \x01(\x0b\x32\x19.flyteidl.core.BranchNodeH\x00\x42\x08\n\x06target\"*\n\x10WorkflowMetadata\x12\x16\n\x0e\x65xecution_role\x18\x01 \x01(\t\"\xfb\x01\n\x10WorkflowTemplate\x12\n\n\x02id\x18\x01 \x01(\t\x12\x31\n\x08metadata\x18\x02 \x01(\x0b\x32\x1f.flyteidl.core.WorkflowMetadata\x12\x30\n\tinterface\x18\x03 \x01(\x0b\x32\x1d.flyteidl.core.TypedInterface\x12\"\n\x05nodes\x18\x04 \x03(\x0b\x32\x13.flyteidl.core.Node\x12\'\n\x07outputs\x18\x05 \x03(\x0b\x32\x16.flyteidl.core.Binding\x12)\n\x0c\x66\x61ilure_node\x18\x06 \x01(\x0b\x32\x13.flyteidl.core.NodeB2Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_interface__pb2.DESCRIPTOR,flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,flyteidl_dot_core_dot_types__pb2.DESCRIPTOR,flyteidl_dot_core_dot_condition__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])




_IFBLOCK = _descriptor.Descriptor(
  name='IfBlock',
  full_name='flyteidl.core.IfBlock',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='condition', full_name='flyteidl.core.IfBlock.condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='then_node', full_name='flyteidl.core.IfBlock.then_node', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=198,
  serialized_end=300,
)


_IFELSEBLOCK = _descriptor.Descriptor(
  name='IfElseBlock',
  full_name='flyteidl.core.IfElseBlock',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='case', full_name='flyteidl.core.IfElseBlock.case', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='other', full_name='flyteidl.core.IfElseBlock.other', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='else_node', full_name='flyteidl.core.IfElseBlock.else_node', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='flyteidl.core.IfElseBlock.error', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='default', full_name='flyteidl.core.IfElseBlock.default',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=303,
  serialized_end=485,
)


_BRANCHNODE = _descriptor.Descriptor(
  name='BranchNode',
  full_name='flyteidl.core.BranchNode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='if_else', full_name='flyteidl.core.BranchNode.if_else', index=0,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=487,
  serialized_end=544,
)


_TASKNODE = _descriptor.Descriptor(
  name='TaskNode',
  full_name='flyteidl.core.TaskNode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='reference_id', full_name='flyteidl.core.TaskNode.reference_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='reference', full_name='flyteidl.core.TaskNode.reference',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=546,
  serialized_end=593,
)


_WORKFLOWNODE = _descriptor.Descriptor(
  name='WorkflowNode',
  full_name='flyteidl.core.WorkflowNode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='launch_plan_reference_id', full_name='flyteidl.core.WorkflowNode.launch_plan_reference_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='workflow_template', full_name='flyteidl.core.WorkflowNode.workflow_template', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='reference', full_name='flyteidl.core.WorkflowNode.reference',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=595,
  serialized_end=720,
)


_NODEMETADATA = _descriptor.Descriptor(
  name='NodeMetadata',
  full_name='flyteidl.core.NodeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='flyteidl.core.NodeMetadata.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timeout', full_name='flyteidl.core.NodeMetadata.timeout', index=1,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='retries', full_name='flyteidl.core.NodeMetadata.retries', index=2,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=722,
  serialized_end=841,
)


_ALIAS = _descriptor.Descriptor(
  name='Alias',
  full_name='flyteidl.core.Alias',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='var', full_name='flyteidl.core.Alias.var', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='alias', full_name='flyteidl.core.Alias.alias', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=843,
  serialized_end=878,
)


_NODE = _descriptor.Descriptor(
  name='Node',
  full_name='flyteidl.core.Node',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.core.Node.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='flyteidl.core.Node.metadata', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs', full_name='flyteidl.core.Node.inputs', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='upstream_node_ids', full_name='flyteidl.core.Node.upstream_node_ids', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output_aliases', full_name='flyteidl.core.Node.output_aliases', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='task_node', full_name='flyteidl.core.Node.task_node', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='workflow_node', full_name='flyteidl.core.Node.workflow_node', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='branch_node', full_name='flyteidl.core.Node.branch_node', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='target', full_name='flyteidl.core.Node.target',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=881,
  serialized_end=1219,
)


_WORKFLOWMETADATA = _descriptor.Descriptor(
  name='WorkflowMetadata',
  full_name='flyteidl.core.WorkflowMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='execution_role', full_name='flyteidl.core.WorkflowMetadata.execution_role', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1221,
  serialized_end=1263,
)


_WORKFLOWTEMPLATE = _descriptor.Descriptor(
  name='WorkflowTemplate',
  full_name='flyteidl.core.WorkflowTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.core.WorkflowTemplate.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='flyteidl.core.WorkflowTemplate.metadata', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interface', full_name='flyteidl.core.WorkflowTemplate.interface', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nodes', full_name='flyteidl.core.WorkflowTemplate.nodes', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='flyteidl.core.WorkflowTemplate.outputs', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='failure_node', full_name='flyteidl.core.WorkflowTemplate.failure_node', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1266,
  serialized_end=1517,
)

_IFBLOCK.fields_by_name['condition'].message_type = flyteidl_dot_core_dot_condition__pb2._BOOLEANEXPRESSION
_IFBLOCK.fields_by_name['then_node'].message_type = _NODE
_IFELSEBLOCK.fields_by_name['case'].message_type = _IFBLOCK
_IFELSEBLOCK.fields_by_name['other'].message_type = _IFBLOCK
_IFELSEBLOCK.fields_by_name['else_node'].message_type = _NODE
_IFELSEBLOCK.fields_by_name['error'].message_type = flyteidl_dot_core_dot_types__pb2._ERROR
_IFELSEBLOCK.oneofs_by_name['default'].fields.append(
  _IFELSEBLOCK.fields_by_name['else_node'])
_IFELSEBLOCK.fields_by_name['else_node'].containing_oneof = _IFELSEBLOCK.oneofs_by_name['default']
_IFELSEBLOCK.oneofs_by_name['default'].fields.append(
  _IFELSEBLOCK.fields_by_name['error'])
_IFELSEBLOCK.fields_by_name['error'].containing_oneof = _IFELSEBLOCK.oneofs_by_name['default']
_BRANCHNODE.fields_by_name['if_else'].message_type = _IFELSEBLOCK
_TASKNODE.oneofs_by_name['reference'].fields.append(
  _TASKNODE.fields_by_name['reference_id'])
_TASKNODE.fields_by_name['reference_id'].containing_oneof = _TASKNODE.oneofs_by_name['reference']
_WORKFLOWNODE.fields_by_name['workflow_template'].message_type = _WORKFLOWTEMPLATE
_WORKFLOWNODE.oneofs_by_name['reference'].fields.append(
  _WORKFLOWNODE.fields_by_name['launch_plan_reference_id'])
_WORKFLOWNODE.fields_by_name['launch_plan_reference_id'].containing_oneof = _WORKFLOWNODE.oneofs_by_name['reference']
_WORKFLOWNODE.oneofs_by_name['reference'].fields.append(
  _WORKFLOWNODE.fields_by_name['workflow_template'])
_WORKFLOWNODE.fields_by_name['workflow_template'].containing_oneof = _WORKFLOWNODE.oneofs_by_name['reference']
_NODEMETADATA.fields_by_name['timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_NODEMETADATA.fields_by_name['retries'].message_type = flyteidl_dot_core_dot_literals__pb2._RETRYSTRATEGY
_NODE.fields_by_name['metadata'].message_type = _NODEMETADATA
_NODE.fields_by_name['inputs'].message_type = flyteidl_dot_core_dot_literals__pb2._BINDING
_NODE.fields_by_name['output_aliases'].message_type = _ALIAS
_NODE.fields_by_name['task_node'].message_type = _TASKNODE
_NODE.fields_by_name['workflow_node'].message_type = _WORKFLOWNODE
_NODE.fields_by_name['branch_node'].message_type = _BRANCHNODE
_NODE.oneofs_by_name['target'].fields.append(
  _NODE.fields_by_name['task_node'])
_NODE.fields_by_name['task_node'].containing_oneof = _NODE.oneofs_by_name['target']
_NODE.oneofs_by_name['target'].fields.append(
  _NODE.fields_by_name['workflow_node'])
_NODE.fields_by_name['workflow_node'].containing_oneof = _NODE.oneofs_by_name['target']
_NODE.oneofs_by_name['target'].fields.append(
  _NODE.fields_by_name['branch_node'])
_NODE.fields_by_name['branch_node'].containing_oneof = _NODE.oneofs_by_name['target']
_WORKFLOWTEMPLATE.fields_by_name['metadata'].message_type = _WORKFLOWMETADATA
_WORKFLOWTEMPLATE.fields_by_name['interface'].message_type = flyteidl_dot_core_dot_interface__pb2._TYPEDINTERFACE
_WORKFLOWTEMPLATE.fields_by_name['nodes'].message_type = _NODE
_WORKFLOWTEMPLATE.fields_by_name['outputs'].message_type = flyteidl_dot_core_dot_literals__pb2._BINDING
_WORKFLOWTEMPLATE.fields_by_name['failure_node'].message_type = _NODE
DESCRIPTOR.message_types_by_name['IfBlock'] = _IFBLOCK
DESCRIPTOR.message_types_by_name['IfElseBlock'] = _IFELSEBLOCK
DESCRIPTOR.message_types_by_name['BranchNode'] = _BRANCHNODE
DESCRIPTOR.message_types_by_name['TaskNode'] = _TASKNODE
DESCRIPTOR.message_types_by_name['WorkflowNode'] = _WORKFLOWNODE
DESCRIPTOR.message_types_by_name['NodeMetadata'] = _NODEMETADATA
DESCRIPTOR.message_types_by_name['Alias'] = _ALIAS
DESCRIPTOR.message_types_by_name['Node'] = _NODE
DESCRIPTOR.message_types_by_name['WorkflowMetadata'] = _WORKFLOWMETADATA
DESCRIPTOR.message_types_by_name['WorkflowTemplate'] = _WORKFLOWTEMPLATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

IfBlock = _reflection.GeneratedProtocolMessageType('IfBlock', (_message.Message,), dict(
  DESCRIPTOR = _IFBLOCK,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.IfBlock)
  ))
_sym_db.RegisterMessage(IfBlock)

IfElseBlock = _reflection.GeneratedProtocolMessageType('IfElseBlock', (_message.Message,), dict(
  DESCRIPTOR = _IFELSEBLOCK,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.IfElseBlock)
  ))
_sym_db.RegisterMessage(IfElseBlock)

BranchNode = _reflection.GeneratedProtocolMessageType('BranchNode', (_message.Message,), dict(
  DESCRIPTOR = _BRANCHNODE,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.BranchNode)
  ))
_sym_db.RegisterMessage(BranchNode)

TaskNode = _reflection.GeneratedProtocolMessageType('TaskNode', (_message.Message,), dict(
  DESCRIPTOR = _TASKNODE,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.TaskNode)
  ))
_sym_db.RegisterMessage(TaskNode)

WorkflowNode = _reflection.GeneratedProtocolMessageType('WorkflowNode', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWNODE,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.WorkflowNode)
  ))
_sym_db.RegisterMessage(WorkflowNode)

NodeMetadata = _reflection.GeneratedProtocolMessageType('NodeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _NODEMETADATA,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.NodeMetadata)
  ))
_sym_db.RegisterMessage(NodeMetadata)

Alias = _reflection.GeneratedProtocolMessageType('Alias', (_message.Message,), dict(
  DESCRIPTOR = _ALIAS,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.Alias)
  ))
_sym_db.RegisterMessage(Alias)

Node = _reflection.GeneratedProtocolMessageType('Node', (_message.Message,), dict(
  DESCRIPTOR = _NODE,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.Node)
  ))
_sym_db.RegisterMessage(Node)

WorkflowMetadata = _reflection.GeneratedProtocolMessageType('WorkflowMetadata', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWMETADATA,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.WorkflowMetadata)
  ))
_sym_db.RegisterMessage(WorkflowMetadata)

WorkflowTemplate = _reflection.GeneratedProtocolMessageType('WorkflowTemplate', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWTEMPLATE,
  __module__ = 'flyteidl.core.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.WorkflowTemplate)
  ))
_sym_db.RegisterMessage(WorkflowTemplate)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/core'))
# @@protoc_insertion_point(module_scope)
