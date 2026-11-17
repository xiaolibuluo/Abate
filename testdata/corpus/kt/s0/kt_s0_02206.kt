package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18170")
fun step18170(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
object Registry18171 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18171 = "registry-62"
    }
}
fun <T> List<T>.secondOrNull18172(): T? = if (size > 1) this[1] else null

fun render18172(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18173<A, B>(val first: A, val second: B)

fun combine18173(left: Int, right: Int): Pair18173<Int, Int> {
    return Pair18173(left + 76, right + 12)
}
data class Pair18174<A, B>(val first: A, val second: B)

fun combine18174(left: Int, right: Int): Pair18174<Int, Int> {
    return Pair18174(left + 83, right + 23)
}
