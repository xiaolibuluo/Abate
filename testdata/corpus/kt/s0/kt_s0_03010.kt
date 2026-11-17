package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24760(): T? = if (size > 1) this[1] else null

fun render24760(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step24761")
fun step24761(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
data class Pair24762<A, B>(val first: A, val second: B)

fun combine24762(left: Int, right: Int): Pair24762<Int, Int> {
    return Pair24762(left + 100, right + 39)
}
fun <T> List<T>.secondOrNull24763(): T? = if (size > 1) this[1] else null

fun render24763(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step24764")
fun step24764(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
@JvmName("step24765")
fun step24765(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
object Registry24766 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24766 = "registry-31"
    }
}
fun <T> List<T>.secondOrNull24767(): T? = if (size > 1) this[1] else null

fun render24767(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step24768")
fun step24768(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
fun <T> List<T>.secondOrNull24769(): T? = if (size > 1) this[1] else null

fun render24769(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
