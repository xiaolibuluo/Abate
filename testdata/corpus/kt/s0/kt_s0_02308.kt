package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18950")
fun step18950(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
object Registry18951 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18951 = "registry-9"
    }
}
fun <T> List<T>.secondOrNull18952(): T? = if (size > 1) this[1] else null

fun render18952(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull18953(): T? = if (size > 1) this[1] else null

fun render18953(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18954(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter18955 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
class Counter18956 {
    var current: Int by Delegates.observable(44) { _, _, new -> println(new) }

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
data class Pair18957<A, B>(val first: A, val second: B)

fun combine18957(left: Int, right: Int): Pair18957<Int, Int> {
    return Pair18957(left + 51, right + 53)
}
fun <T> List<T>.secondOrNull18958(): T? = if (size > 1) this[1] else null

fun render18958(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
