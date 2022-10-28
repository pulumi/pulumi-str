# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ReplaceResult',
    'AwaitableReplaceResult',
    'replace',
    'replace_output',
]

@pulumi.output_type
class ReplaceResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableReplaceResult(ReplaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ReplaceResult(
            result=self.result)


def replace(new: Optional[str] = None,
            old: Optional[str] = None,
            string: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableReplaceResult:
    """
    Replace returns a copy of the string s with all
    non-overlapping instances of old replaced by new.
    If old is empty, it matches at the beginning of the string
    and after each UTF-8 sequence, yielding up to k+1 replacements
    for a k-rune string.
    """
    __args__ = dict()
    __args__['new'] = new
    __args__['old'] = old
    __args__['string'] = string
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('str:index:replace', __args__, opts=opts, typ=ReplaceResult).value

    return AwaitableReplaceResult(
        result=__ret__.result)


@_utilities.lift_output_func(replace)
def replace_output(new: Optional[pulumi.Input[str]] = None,
                   old: Optional[pulumi.Input[str]] = None,
                   string: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ReplaceResult]:
    """
    Replace returns a copy of the string s with all
    non-overlapping instances of old replaced by new.
    If old is empty, it matches at the beginning of the string
    and after each UTF-8 sequence, yielding up to k+1 replacements
    for a k-rune string.
    """
    ...
