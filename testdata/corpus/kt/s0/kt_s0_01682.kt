package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull13860(): T? = if (size > 1) this[1] else null

fun render13860(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step13861")
fun step13861(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
class Counter13862 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
data class Pair13863<A, B>(val first: A, val second: B)

fun combine13863(left: Int, right: Int): Pair13863<Int, Int> {
    return Pair13863(left + 99, right + 45)
}
data class Pair13864<A, B>(val first: A, val second: B)

fun combine13864(left: Int, right: Int): Pair13864<Int, Int> {
    return Pair13864(left + 9, right + 56)
}
