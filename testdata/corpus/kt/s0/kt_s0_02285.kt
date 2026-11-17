package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18740(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18741<A, B>(val first: A, val second: B)

fun combine18741(left: Int, right: Int): Pair18741<Int, Int> {
    return Pair18741(left + 27, right + 23)
}
@JvmName("step18742")
fun step18742(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
class Counter18743 {
    var current: Int by Delegates.observable(41) { _, _, new -> println(new) }

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
data class Pair18744<A, B>(val first: A, val second: B)

fun combine18744(left: Int, right: Int): Pair18744<Int, Int> {
    return Pair18744(left + 48, right + 56)
}
fun <T> List<T>.secondOrNull18745(): T? = if (size > 1) this[1] else null

fun render18745(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
