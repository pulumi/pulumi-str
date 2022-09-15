# coding=utf-8
# *** WARNING: this file was generated by pulumi-gen-str. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'SplitResult',
    'AwaitableSplitResult',
    'split',
    'split_output',
]

@pulumi.output_type
class SplitResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[str]:
        """
        The result of the string split.
        """
        return pulumi.get(self, "result")


class AwaitableSplitResult(SplitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SplitResult(
            result=self.result)


def split(count: Optional[int] = None,
          on: Optional[str] = None,
          string: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSplitResult:
    """
    Split a string on a regex.


    :param int count: `count` determines the number of substrings to return. 
           If `count` is not provided, it defaults to substrings.
           If `count` is provided then the last substring will be the unsplit remainder.
           It is an error to pass `count < 1`.
    :param str on: The regex to split on.
    :param str string: The string on which to split.
    """
    __args__ = dict()
    __args__['count'] = count
    __args__['on'] = on
    __args__['string'] = string
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('str:regexp:split', __args__, opts=opts, typ=SplitResult).value

    return AwaitableSplitResult(
        result=__ret__.result)


@_utilities.lift_output_func(split)
def split_output(count: Optional[pulumi.Input[Optional[int]]] = None,
                 on: Optional[pulumi.Input[str]] = None,
                 string: Optional[pulumi.Input[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SplitResult]:
    """
    Split a string on a regex.


    :param int count: `count` determines the number of substrings to return. 
           If `count` is not provided, it defaults to substrings.
           If `count` is provided then the last substring will be the unsplit remainder.
           It is an error to pass `count < 1`.
    :param str on: The regex to split on.
    :param str string: The string on which to split.
    """
    ...
