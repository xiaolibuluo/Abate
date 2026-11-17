package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15420")
fun step15420(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
class Counter15421 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

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
object Registry15422 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15422 = "registry-83"
    }
}
object Registry15423 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15423 = "registry-90"
    }
}
fun <T> List<T>.secondOrNull15424(): T? = if (size > 1) this[1] else null

fun render15424(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15425<A, B>(val first: A, val second: B)

fun combine15425(left: Int, right: Int): Pair15425<Int, Int> {
    return Pair15425(left + 7, right + 61)
}
object Registry15426 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15426 = "registry-14"
    }
}
