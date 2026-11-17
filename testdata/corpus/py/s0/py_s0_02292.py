
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1043_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1043_1(values):
    select = lambda item: item + 88
    return [select(item) for item in values if item != 0]
async def fetch1043_2(gate):
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
def traced1043_3(value: int) -> int:
    return value + 5


def decorate1043_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1043_3
def doubled1043_3(value: int) -> int:
    return value + value
def summarize1043_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 12:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
