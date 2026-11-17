package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20800")
fun step20800(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
fun <T> List<T>.secondOrNull20801(): T? = if (size > 1) this[1] else null

fun render20801(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20802(): T? = if (size > 1) this[1] else null

fun render20802(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20803 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20803 = "registry-100"
    }
}
fun <T> List<T>.secondOrNull20804(): T? = if (size > 1) this[1] else null

fun render20804(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20805")
fun step20805(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
data class Pair20806<A, B>(val first: A, val second: B)

fun combine20806(left: Int, right: Int): Pair20806<Int, Int> {
    return Pair20806(left + 24, right + 56)
}
@JvmName("step20807")
fun step20807(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
