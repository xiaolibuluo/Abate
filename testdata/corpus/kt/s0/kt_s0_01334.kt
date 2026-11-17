package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11000 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
@JvmName("step11001")
fun step11001(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
object Registry11002 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11002 = "registry-60"
    }
}
fun <T> List<T>.secondOrNull11003(): T? = if (size > 1) this[1] else null

fun render11003(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11004<A, B>(val first: A, val second: B)

fun combine11004(left: Int, right: Int): Pair11004<Int, Int> {
    return Pair11004(left + 74, right + 67)
}
fun <T> List<T>.secondOrNull11005(): T? = if (size > 1) this[1] else null

fun render11005(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
