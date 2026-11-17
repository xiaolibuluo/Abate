package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7990<A, B>(val first: A, val second: B)

fun combine7990(left: Int, right: Int): Pair7990<Int, Int> {
    return Pair7990(left + 13, right + 51)
}
fun <T, R> List<T>.mapChecked7991(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step7992")
fun step7992(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
@JvmName("step7993")
fun step7993(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
object Registry7994 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7994 = "registry-41"
    }
}
fun <T> List<T>.secondOrNull7995(): T? = if (size > 1) this[1] else null

fun render7995(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
