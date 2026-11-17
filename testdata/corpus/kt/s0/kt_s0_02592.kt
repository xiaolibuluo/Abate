package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21350<A, B>(val first: A, val second: B)

fun combine21350(left: Int, right: Int): Pair21350<Int, Int> {
    return Pair21350(left + 18, right + 9)
}
class Counter21351 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked21352(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step21353")
fun step21353(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
@JvmName("step21354")
fun step21354(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
fun <T> List<T>.secondOrNull21355(): T? = if (size > 1) this[1] else null

fun render21355(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked21356(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21357<A, B>(val first: A, val second: B)

fun combine21357(left: Int, right: Int): Pair21357<Int, Int> {
    return Pair21357(left + 67, right + 25)
}
fun <T> List<T>.secondOrNull21358(): T? = if (size > 1) this[1] else null

fun render21358(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
