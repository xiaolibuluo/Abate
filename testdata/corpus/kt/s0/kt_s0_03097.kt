package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25430")
fun step25430(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
fun <T> List<T>.secondOrNull25431(): T? = if (size > 1) this[1] else null

fun render25431(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25432<A, B>(val first: A, val second: B)

fun combine25432(left: Int, right: Int): Pair25432<Int, Int> {
    return Pair25432(left + 98, right + 20)
}
class Counter25433 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
