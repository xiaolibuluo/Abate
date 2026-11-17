
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1673_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1673_1(gate):
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
def traced1673_2(value: int) -> int:
    return value + 40


def decorate1673_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1673_2
def doubled1673_2(value: int) -> int:
    return value + value
def summarize1673_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
