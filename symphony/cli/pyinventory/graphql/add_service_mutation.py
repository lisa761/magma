#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .customer_fragment import CustomerFragment, QUERY as CustomerFragmentQuery
from .property_fragment import PropertyFragment, QUERY as PropertyFragmentQuery
from .service_create_data_input import ServiceCreateData


QUERY: List[str] = CustomerFragmentQuery + PropertyFragmentQuery + ["""
mutation AddServiceMutation($data: ServiceCreateData!) {
  addService(data: $data) {
    id
    name
    externalId
    customer {
      ...CustomerFragment
    }
    endpoints {
      id
      port {
        id
        properties {
          ...PropertyFragment
        }
        definition {
          id
          name
          portType {
            id
            name
          }
        }
        link {
          id
          properties {
            ...PropertyFragment
          }
          services {
            id
          }
        }
      }
      definition {
        role
      }
    }
    links {
      id
      properties {
        ...PropertyFragment
      }
      services {
        id
      }
    }
  }
}

"""]

@dataclass
class AddServiceMutation(DataClassJsonMixin):
    @dataclass
    class AddServiceMutationData(DataClassJsonMixin):
        @dataclass
        class Service(DataClassJsonMixin):
            @dataclass
            class Customer(CustomerFragment):
                pass

            @dataclass
            class ServiceEndpoint(DataClassJsonMixin):
                @dataclass
                class EquipmentPort(DataClassJsonMixin):
                    @dataclass
                    class Property(PropertyFragment):
                        pass

                    @dataclass
                    class EquipmentPortDefinition(DataClassJsonMixin):
                        @dataclass
                        class EquipmentPortType(DataClassJsonMixin):
                            id: str
                            name: str

                        id: str
                        name: str
                        portType: Optional[EquipmentPortType]

                    @dataclass
                    class Link(DataClassJsonMixin):
                        @dataclass
                        class Property(PropertyFragment):
                            pass

                        @dataclass
                        class Service(DataClassJsonMixin):
                            id: str

                        id: str
                        properties: List[Property]
                        services: List[Service]

                    id: str
                    properties: List[Property]
                    definition: EquipmentPortDefinition
                    link: Optional[Link]

                @dataclass
                class ServiceEndpointDefinition(DataClassJsonMixin):
                    role: Optional[str]

                id: str
                definition: ServiceEndpointDefinition
                port: Optional[EquipmentPort]

            @dataclass
            class Link(DataClassJsonMixin):
                @dataclass
                class Property(PropertyFragment):
                    pass

                @dataclass
                class Service(DataClassJsonMixin):
                    id: str

                id: str
                properties: List[Property]
                services: List[Service]

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str]
            customer: Optional[Customer]

        addService: Service

    data: AddServiceMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, data: ServiceCreateData) -> AddServiceMutationData:
        # fmt: off
        variables = {"data": data}
        response_text = client.call(''.join(set(QUERY)), variables=variables)
        return cls.from_json(response_text).data