package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry9090 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9090 = "registry-85"
    }
}
fun <T> List<T>.secondOrNull9091(): T? = if (size > 1) this[1] else null

fun render9091(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9092")
fun step9092(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
object Registry9093 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9093 = "registry-9"
    }
}
object Registry9094 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9094 = "registry-16"
    }
}
@JvmName("step9095")
fun step9095(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
@JvmName("step9096")
fun step9096(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
data class Pair9097<A, B>(val first: A, val second: B)

fun combine9097(left: Int, right: Int): Pair9097<Int, Int> {
    return Pair9097(left + 37, right + 44)
}
