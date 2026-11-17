package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18810")
fun step18810(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
object Registry18811 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18811 = "registry-21"
    }
}
fun <T> List<T>.secondOrNull18812(): T? = if (size > 1) this[1] else null

fun render18812(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry18813 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18813 = "registry-35"
    }
}
@JvmName("step18814")
fun step18814(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
data class Pair18815<A, B>(val first: A, val second: B)

fun combine18815(left: Int, right: Int): Pair18815<Int, Int> {
    return Pair18815(left + 49, right + 55)
}
fun <T> List<T>.secondOrNull18816(): T? = if (size > 1) this[1] else null

fun render18816(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18817<A, B>(val first: A, val second: B)

fun combine18817(left: Int, right: Int): Pair18817<Int, Int> {
    return Pair18817(left + 63, right + 16)
}
