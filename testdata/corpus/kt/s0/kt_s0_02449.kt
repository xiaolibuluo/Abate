package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20110<A, B>(val first: A, val second: B)

fun combine20110(left: Int, right: Int): Pair20110<Int, Int> {
    return Pair20110(left + 55, right + 56)
}
fun <T> List<T>.secondOrNull20111(): T? = if (size > 1) this[1] else null

fun render20111(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20112")
fun step20112(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
object Registry20113 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20113 = "registry-76"
    }
}
