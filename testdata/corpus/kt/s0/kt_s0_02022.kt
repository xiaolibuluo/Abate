package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16720")
fun step16720(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
fun <T> List<T>.secondOrNull16721(): T? = if (size > 1) this[1] else null

fun render16721(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16722 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
data class Pair16723<A, B>(val first: A, val second: B)

fun combine16723(left: Int, right: Int): Pair16723<Int, Int> {
    return Pair16723(left + 34, right + 34)
}
