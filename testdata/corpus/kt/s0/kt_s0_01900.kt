package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15670<A, B>(val first: A, val second: B)

fun combine15670(left: Int, right: Int): Pair15670<Int, Int> {
    return Pair15670(left + 6, right + 59)
}
fun <T> List<T>.secondOrNull15671(): T? = if (size > 1) this[1] else null

fun render15671(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter15672 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
