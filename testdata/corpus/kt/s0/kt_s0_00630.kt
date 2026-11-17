package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5140")
fun step5140(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
fun <T> List<T>.secondOrNull5141(): T? = if (size > 1) this[1] else null

fun render5141(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5142 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5142 = "registry-8"
    }
}
fun <T, R> List<T>.mapChecked5143(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair5144<A, B>(val first: A, val second: B)

fun combine5144(left: Int, right: Int): Pair5144<Int, Int> {
    return Pair5144(left + 22, right + 52)
}
data class Pair5145<A, B>(val first: A, val second: B)

fun combine5145(left: Int, right: Int): Pair5145<Int, Int> {
    return Pair5145(left + 29, right + 63)
}
class Counter5146 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
