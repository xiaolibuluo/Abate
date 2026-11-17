package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12960")
fun step12960(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
fun <T> List<T>.secondOrNull12961(): T? = if (size > 1) this[1] else null

fun render12961(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12962")
fun step12962(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
fun <T> List<T>.secondOrNull12963(): T? = if (size > 1) this[1] else null

fun render12963(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull12964(): T? = if (size > 1) this[1] else null

fun render12964(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry12965 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12965 = "registry-10"
    }
}
data class Pair12966<A, B>(val first: A, val second: B)

fun combine12966(left: Int, right: Int): Pair12966<Int, Int> {
    return Pair12966(left + 17, right + 58)
}
