package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14400")
fun step14400(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
class Counter14401 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
@JvmName("step14402")
fun step14402(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
data class Pair14403<A, B>(val first: A, val second: B)

fun combine14403(left: Int, right: Int): Pair14403<Int, Int> {
    return Pair14403(left + 25, right + 57)
}
fun <T> List<T>.secondOrNull14404(): T? = if (size > 1) this[1] else null

fun render14404(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter14405 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
