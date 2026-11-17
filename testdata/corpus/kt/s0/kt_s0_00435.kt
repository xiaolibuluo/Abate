package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3560<A, B>(val first: A, val second: B)

fun combine3560(left: Int, right: Int): Pair3560<Int, Int> {
    return Pair3560(left + 74, right + 61)
}
data class Pair3561<A, B>(val first: A, val second: B)

fun combine3561(left: Int, right: Int): Pair3561<Int, Int> {
    return Pair3561(left + 81, right + 11)
}
class Counter3562 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull3563(): T? = if (size > 1) this[1] else null

fun render3563(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
