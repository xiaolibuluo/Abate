package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10360 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10360 = "registry-87"
    }
}
@JvmName("step10361")
fun step10361(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
data class Pair10362<A, B>(val first: A, val second: B)

fun combine10362(left: Int, right: Int): Pair10362<Int, Int> {
    return Pair10362(left + 101, right + 24)
}
fun <T> List<T>.secondOrNull10363(): T? = if (size > 1) this[1] else null

fun render10363(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
