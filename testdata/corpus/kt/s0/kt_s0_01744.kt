package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter14290 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
@JvmName("step14291")
fun step14291(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
data class Pair14292<A, B>(val first: A, val second: B)

fun combine14292(left: Int, right: Int): Pair14292<Int, Int> {
    return Pair14292(left + 69, right + 30)
}
fun <T, R> List<T>.mapChecked14293(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
