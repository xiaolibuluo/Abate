package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24790<A, B>(val first: A, val second: B)

fun combine24790(left: Int, right: Int): Pair24790<Int, Int> {
    return Pair24790(left + 28, right + 38)
}
fun <T, R> List<T>.mapChecked24791(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter24792 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
object Registry24793 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24793 = "registry-49"
    }
}
fun <T> List<T>.secondOrNull24794(): T? = if (size > 1) this[1] else null

fun render24794(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry24795 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24795 = "registry-63"
    }
}
fun <T> List<T>.secondOrNull24796(): T? = if (size > 1) this[1] else null

fun render24796(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24797<A, B>(val first: A, val second: B)

fun combine24797(left: Int, right: Int): Pair24797<Int, Int> {
    return Pair24797(left + 77, right + 54)
}
@JvmName("step24798")
fun step24798(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
