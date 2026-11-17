package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11580 = "registry-24"
    }
}
class Counter11581 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull11582(): T? = if (size > 1) this[1] else null

fun render11582(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry11583 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11583 = "registry-45"
    }
}
@JvmName("step11584")
fun step11584(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
fun <T> List<T>.secondOrNull11585(): T? = if (size > 1) this[1] else null

fun render11585(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11586")
fun step11586(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
data class Pair11587<A, B>(val first: A, val second: B)

fun combine11587(left: Int, right: Int): Pair11587<Int, Int> {
    return Pair11587(left + 73, right + 18)
}
