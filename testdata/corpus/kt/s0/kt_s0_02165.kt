package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17820<A, B>(val first: A, val second: B)

fun combine17820(left: Int, right: Int): Pair17820<Int, Int> {
    return Pair17820(left + 85, right + 39)
}
fun <T> List<T>.secondOrNull17821(): T? = if (size > 1) this[1] else null

fun render17821(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17822 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17822 = "registry-99"
    }
}
class Counter17823 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked17824(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17825<A, B>(val first: A, val second: B)

fun combine17825(left: Int, right: Int): Pair17825<Int, Int> {
    return Pair17825(left + 23, right + 33)
}
