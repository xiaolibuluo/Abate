package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5170<A, B>(val first: A, val second: B)

fun combine5170(left: Int, right: Int): Pair5170<Int, Int> {
    return Pair5170(left + 33, right + 29)
}
@JvmName("step5171")
fun step5171(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
object Registry5172 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5172 = "registry-47"
    }
}
data class Pair5173<A, B>(val first: A, val second: B)

fun combine5173(left: Int, right: Int): Pair5173<Int, Int> {
    return Pair5173(left + 54, right + 62)
}
class Counter5174 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked5175(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull5176(): T? = if (size > 1) this[1] else null

fun render5176(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
