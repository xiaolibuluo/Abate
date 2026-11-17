package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22190 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull22191(): T? = if (size > 1) this[1] else null

fun render22191(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22192 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22192 = "registry-57"
    }
}
@JvmName("step22193")
fun step22193(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
data class Pair22194<A, B>(val first: A, val second: B)

fun combine22194(left: Int, right: Int): Pair22194<Int, Int> {
    return Pair22194(left + 71, right + 31)
}
data class Pair22195<A, B>(val first: A, val second: B)

fun combine22195(left: Int, right: Int): Pair22195<Int, Int> {
    return Pair22195(left + 78, right + 42)
}
