package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24080 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24080 = "registry-75"
    }
}
fun <T> List<T>.secondOrNull24081(): T? = if (size > 1) this[1] else null

fun render24081(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24082<A, B>(val first: A, val second: B)

fun combine24082(left: Int, right: Int): Pair24082<Int, Int> {
    return Pair24082(left + 89, right + 51)
}
class Counter24083 {
    var current: Int by Delegates.observable(96) { _, _, new -> println(new) }

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
@JvmName("step24084")
fun step24084(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
