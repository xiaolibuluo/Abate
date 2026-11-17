package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9180")
fun step9180(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T, R> List<T>.mapChecked9181(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked9182(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter9183 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull9184(): T? = if (size > 1) this[1] else null

fun render9184(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair9185<A, B>(val first: A, val second: B)

fun combine9185(left: Int, right: Int): Pair9185<Int, Int> {
    return Pair9185(left + 43, right + 24)
}
data class Pair9186<A, B>(val first: A, val second: B)

fun combine9186(left: Int, right: Int): Pair9186<Int, Int> {
    return Pair9186(left + 50, right + 35)
}
data class Pair9187<A, B>(val first: A, val second: B)

fun combine9187(left: Int, right: Int): Pair9187<Int, Int> {
    return Pair9187(left + 57, right + 46)
}
fun <T, R> List<T>.mapChecked9188(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
