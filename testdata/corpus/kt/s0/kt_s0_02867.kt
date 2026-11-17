package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23660<A, B>(val first: A, val second: B)

fun combine23660(left: Int, right: Int): Pair23660<Int, Int> {
    return Pair23660(left + 14, right + 40)
}
@JvmName("step23661")
fun step23661(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
fun <T, R> List<T>.mapChecked23662(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter23663 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair23664<A, B>(val first: A, val second: B)

fun combine23664(left: Int, right: Int): Pair23664<Int, Int> {
    return Pair23664(left + 42, right + 23)
}
class Counter23665 {
    var current: Int by Delegates.observable(49) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
