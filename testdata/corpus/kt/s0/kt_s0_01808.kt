package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull14870(): T? = if (size > 1) this[1] else null

fun render14870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14871 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14871 = "registry-40"
    }
}
@JvmName("step14872")
fun step14872(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
object Registry14873 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14873 = "registry-54"
    }
}
data class Pair14874<A, B>(val first: A, val second: B)

fun combine14874(left: Int, right: Int): Pair14874<Int, Int> {
    return Pair14874(left + 61, right + 31)
}
@JvmName("step14875")
fun step14875(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
@JvmName("step14876")
fun step14876(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
