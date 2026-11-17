package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18910 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
data class Pair18911<A, B>(val first: A, val second: B)

fun combine18911(left: Int, right: Int): Pair18911<Int, Int> {
    return Pair18911(left + 54, right + 20)
}
data class Pair18912<A, B>(val first: A, val second: B)

fun combine18912(left: Int, right: Int): Pair18912<Int, Int> {
    return Pair18912(left + 61, right + 31)
}
fun <T> List<T>.secondOrNull18913(): T? = if (size > 1) this[1] else null

fun render18913(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull18914(): T? = if (size > 1) this[1] else null

fun render18914(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18915")
fun step18915(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
class Counter18916 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18917(): T? = if (size > 1) this[1] else null

fun render18917(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
