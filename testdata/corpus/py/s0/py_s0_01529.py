
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize698_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch698_1(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def traced698_2(value: int) -> int:
    return value + 72


def decorate698_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate698_2
def doubled698_2(value: int) -> int:
    return value + value
def traced698_3(value: int) -> int:
    return value + 79


def decorate698_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate698_3
def doubled698_3(value: int) -> int:
    return value + value
