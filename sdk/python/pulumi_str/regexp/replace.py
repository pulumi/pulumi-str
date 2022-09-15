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
        """
        The input "string" where each pattern matching "old" is replaced by "new".
        """
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
    A regex based replace on a string.

    The underlying regexp engine is the go "regexp" stdlib package.
    You can see details at for available patterns at https://pkg.go.dev/regexp/syntax.


    :param str new: The new string.
           
           Note: Inside repl, "$" signs are interpreted as an Expand, so for instance $1 represents the text of the first submatch. 
    :param str old: The regular expression to match against.
    :param str string: The string to operate over.
    """
    __args__ = dict()
    __args__['new'] = new
    __args__['old'] = old
    __args__['string'] = string
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('str:regexp:replace', __args__, opts=opts, typ=ReplaceResult).value

    return AwaitableReplaceResult(
        result=__ret__.result)


@_utilities.lift_output_func(replace)
def replace_output(new: Optional[pulumi.Input[str]] = None,
                   old: Optional[pulumi.Input[str]] = None,
                   string: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ReplaceResult]:
    """
    A regex based replace on a string.

    The underlying regexp engine is the go "regexp" stdlib package.
    You can see details at for available patterns at https://pkg.go.dev/regexp/syntax.


    :param str new: The new string.
           
           Note: Inside repl, "$" signs are interpreted as an Expand, so for instance $1 represents the text of the first submatch. 
    :param str old: The regular expression to match against.
    :param str string: The string to operate over.
    """
    ...