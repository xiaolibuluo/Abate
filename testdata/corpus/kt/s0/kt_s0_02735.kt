package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22580(): T? = if (size > 1) this[1] else null

fun render22580(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22581 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22581 = "registry-72"
    }
}
@JvmName("step22582")
fun step22582(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
data class Pair22583<A, B>(val first: A, val second: B)

fun combine22583(left: Int, right: Int): Pair22583<Int, Int> {
    return Pair22583(left + 86, right + 49)
}
fun <T> List<T>.secondOrNull22584(): T? = if (size > 1) this[1] else null

fun render22584(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
