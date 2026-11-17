package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull15360(): T? = if (size > 1) this[1] else null

fun render15360(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15361<A, B>(val first: A, val second: B)

fun combine15361(left: Int, right: Int): Pair15361<Int, Int> {
    return Pair15361(left + 95, right + 36)
}
data class Pair15362<A, B>(val first: A, val second: B)

fun combine15362(left: Int, right: Int): Pair15362<Int, Int> {
    return Pair15362(left + 5, right + 47)
}
class Counter15363 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
@JvmName("step15364")
fun step15364(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
@JvmName("step15365")
fun step15365(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
